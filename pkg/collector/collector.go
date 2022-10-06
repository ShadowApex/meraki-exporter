package collector

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/ShadowApex/meraki-go-sdk/meraki"
	"github.com/prometheus/client_golang/prometheus"
)

const baseURI = "https://api.meraki.com/api/v1"
const header = "X-Cisco-Meraki-API-Key"

// MerakiCollectorConfig defines configuration options for the Meraki Collector
type MerakiCollectorConfig struct {
	Host   string
	ApiKey string
	Debug  bool
	Scheme string
}

type MerakiCollector struct {
	// Meraki client
	client  *meraki.APIClient
	metrics map[string]Settable
}

// Declare your exporter metrics here. Referred to as "collectors"
func NewMerakiCollector(config ...func(*MerakiCollectorConfig)) *MerakiCollector {
	// Build the collector configuration
	options := &MerakiCollectorConfig{
		Host:   "api.meraki.com",
		Scheme: "https",
		ApiKey: os.Getenv("MERAKI_KEY"),
		Debug:  os.Getenv("DEBUG") == "1",
	}
	for _, applyConf := range config {
		applyConf(options)
	}

	// Build the meraki config
	conf := meraki.NewConfiguration()
	conf.Host = options.Host
	conf.Debug = options.Debug
	conf.Scheme = options.Scheme
	conf.DefaultHeader[header] = options.ApiKey

	// Build the meraki API client
	client := meraki.NewAPIClient(conf)

	return &MerakiCollector{
		client: client,
		metrics: map[string]Settable{
			"orgCount": NewGauge(prometheus.NewDesc("meraki_organization_count",
				"Number of organizations tied to the Meraki account",
				nil, nil,
			)),
			"networkCount": NewGauge(prometheus.NewDesc("meraki_network_count",
				"Number of networks in a given organization",
				[]string{"org"}, nil,
			)),
			"deviceStatus": NewGauge(prometheus.NewDesc("meraki_device_status",
				"Whether the device is in a good state",
				[]string{"org", "network", "device", "serial", "model"}, nil,
			)),
			"deviceLastContact": NewGauge(prometheus.NewDesc("meraki_device_last_contact",
				"Last time the given device was contacted",
				[]string{"org", "network", "device", "serial", "model"}, nil,
			)),
			"portEnabled": NewGauge(prometheus.NewDesc("meraki_switch_port_enabled",
				"Whether the given switch port is enabled",
				[]string{"org", "network", "switch", "port"}, nil,
			)),
			"portIsUplink": NewGauge(prometheus.NewDesc("meraki_switch_port_uplink",
				"Whether the given switch port is an uplink",
				[]string{"org", "network", "switch", "port"}, nil,
			)),
			"portClientCount": NewGauge(prometheus.NewDesc("meraki_switch_port_client_count",
				"Number of clients on the given port",
				[]string{"org", "network", "switch", "port"}, nil,
			)),
			"portBytesSent": NewCounter(prometheus.NewDesc("meraki_switch_port_bytes_sent",
				"Number of bytes sent on the given port",
				[]string{"org", "network", "switch", "port"}, nil,
			)),
			"portBytesRecv": NewCounter(prometheus.NewDesc("meraki_switch_port_bytes_recv",
				"Number of bytes received on the given port",
				[]string{"org", "network", "switch", "port"}, nil,
			)),
		},
	}
}

func (m *MerakiCollector) Describe(ch chan<- *prometheus.Desc) {
	// Add one of these lines for each of your collectors declared above
	//ch <- m.systemTime
	//ch <- m.otherThing
}

// This fuction runs when Prometheus scrapes the exporter. It will set a new value for the metric(s)
// I have no idea how it works, but it does.
func (m *MerakiCollector) Collect(ch chan<- prometheus.Metric) {
	// Fetch all organizations in the Meraki account
	ctx := context.Background()
	orgs, err := m.GetOrganizations(ctx)
	if err != nil {
		panic(err)
	}
	m.metrics["orgCount"].Set(ch, float64(len(orgs)))

	// Fetch all networks for each org
	for _, org := range orgs {
		m.collectNetworks(ctx, ch, org)
	}
}

func (m *MerakiCollector) collectNetworks(ctx context.Context, ch chan<- prometheus.Metric, org meraki.GetOrganizations200ResponseInner) {
	// Get all networks
	networks, err := m.GetNetworks(ctx, org)
	if err != nil {
		panic(err)
	}
	m.metrics["networkCount"].Set(ch, float64(len(networks)), *org.Name)

	// Map our networks by network ID
	networkByID := map[string]meraki.UnbindNetwork200Response{}
	for _, network := range networks {
		networkByID[network.GetId()] = network
	}

	// Get the status of all devices in the org
	devStatuses, err := m.GetDevicesStatus(ctx, org)
	if err != nil {
		panic(err)
	}

	// Map our devices by network and product type
	deviceStatusByNetwork := map[string]map[string][]GetOrganizationDevicesStatuses200Response{}
	for _, device := range devStatuses {
		networkId := device.GetNetworkId()
		if _, ok := deviceStatusByNetwork[networkId]; !ok {
			deviceStatusByNetwork[networkId] = map[string][]GetOrganizationDevicesStatuses200Response{}
		}
		productType := device.GetProductType()
		if _, ok := deviceStatusByNetwork[networkId][productType]; !ok {
			deviceStatusByNetwork[networkId][productType] = []GetOrganizationDevicesStatuses200Response{}
		}
		deviceStatusByNetwork[networkId][productType] = append(deviceStatusByNetwork[networkId][productType], device)
	}

	// Loop through all networks
	for networkId, network := range networkByID {
		networkName := network.GetName()
		fmt.Println("Network: ", networkName)
		fmt.Println("NetworkID: ", networkId)

		// Loop through all devices
		for _, devices := range deviceStatusByNetwork[networkId] {
			for _, device := range devices {
				devName := device.GetName()
				devSerial := device.GetSerial()
				var status float64 = 1
				if device.GetStatus() != "online" {
					status = 0
				}
				lastContactStr := device.GetLastReportedAt()
				lastContact, _ := time.Parse(time.RFC3339Nano, lastContactStr)

				m.metrics["deviceStatus"].Set(ch, status, *org.Name, networkName, devName, devSerial, device.GetModel())
				m.metrics["deviceLastContact"].Set(ch, float64(lastContact.Unix()), *org.Name, networkName, devName, devSerial, device.GetModel())
			}
		}

		// Loop through all devices on the network
		switches := deviceStatusByNetwork[networkId]["switch"]
		for _, device := range switches {
			serial := device.GetSerial()
			devName := serial
			if name, ok := device.GetNameOk(); ok {
				devName = *name
			}

			// Fetch the port statuses for this switch
			ports, err := m.GetPortsStatus(ctx, serial)
			if err != nil {
				fmt.Println("Warning:", err)
				continue
			}

			// Loop through all switch ports
			for _, port := range ports {
				portId := port["portId"].(string)
				enabled := boolToFloat64(port["enabled"].(bool))
				isUplink := boolToFloat64(port["isUplink"].(bool))
				clientCount := port["clientCount"].(float64)
				usageInKb := port["usageInKb"].(map[string]interface{})
				sent := usageInKb["sent"].(float64)
				recv := usageInKb["recv"].(float64)

				m.metrics["portEnabled"].Set(ch, enabled, *org.Name, networkName, devName, portId)
				m.metrics["portIsUplink"].Set(ch, isUplink, *org.Name, networkName, devName, portId)
				m.metrics["portClientCount"].Set(ch, clientCount, *org.Name, networkName, devName, portId)
				m.metrics["portBytesRecv"].Set(ch, recv*1024, *org.Name, networkName, devName, portId)
				m.metrics["portBytesSent"].Set(ch, sent*1024, *org.Name, networkName, devName, portId)
			}
		}
	}

}

// Fetch all organizations
func (m *MerakiCollector) GetOrganizations(ctx context.Context) ([]meraki.GetOrganizations200ResponseInner, error) {
	req := m.client.OrganizationsApi.GetOrganizations(ctx)
	orgs, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

// GetNetworks will return all networks for the given organization
func (m *MerakiCollector) GetNetworks(ctx context.Context, org meraki.GetOrganizations200ResponseInner) ([]meraki.UnbindNetwork200Response, error) {
	orgId := org.GetId()
	req := m.client.OrganizationsApi.GetOrganizationNetworks(ctx, orgId)
	networks, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return networks, nil
}

// GetPortsStatus will return switch port statuses
func (m *MerakiCollector) GetPortsStatus(ctx context.Context, serial string) ([]map[string]interface{}, error) {
	req := m.client.SwitchApi.GetDeviceSwitchPortsStatuses(ctx, serial)
	statuses, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

// Get all devices for a network
func (m *MerakiCollector) GetDevices(ctx context.Context, network meraki.UnbindNetwork200Response) ([]map[string]interface{}, error) {
	req := m.client.NetworksApi.GetNetworkDevices(ctx, network.GetId())
	devices, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return devices, nil
}

// GetDevicesStatus will return the status of all devices for an organization
func (m *MerakiCollector) GetDevicesStatus(ctx context.Context, org meraki.GetOrganizations200ResponseInner) ([]GetOrganizationDevicesStatuses200Response, error) {
	orgId := org.GetId()
	req := m.client.OrganizationsApi.GetOrganizationDevicesStatuses(ctx, orgId)
	// Manually parse response
	_, res, _ := req.Execute()
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	// Manually unmarshal into a custom struct because of broken OpenAPI def
	var statuses []GetOrganizationDevicesStatuses200Response
	err = json.Unmarshal(body, &statuses)
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

func boolToFloat64(v bool) float64 {
	if v {
		return 1
	}
	return 0
}
