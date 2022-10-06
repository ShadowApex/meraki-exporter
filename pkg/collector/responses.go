package collector

// GetOrganizationDevicesStatuses200Response struct for GetOrganizationDevicesStatuses200Response
type GetOrganizationDevicesStatuses200Response struct {
	// Device Name
	Name *string `json:"name,omitempty"`
	// Device Serial Number
	Serial *string `json:"serial,omitempty"`
	// MAC Address
	Mac *string `json:"mac,omitempty"`
	// Public IP Address
	PublicIp *string `json:"publicIp,omitempty"`
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Device Status
	Status *string `json:"status,omitempty"`
	// Device Last Reported Location
	LastReportedAt *string `json:"lastReportedAt,omitempty"`
	// LAN IP Address
	LanIp *string `json:"lanIp,omitempty"`
	// IP Gateway
	Gateway *string `json:"gateway,omitempty"`
	// IP Type
	IpType *string `json:"ipType,omitempty"`
	// Primary DNS
	PrimaryDns *string `json:"primaryDns,omitempty"`
	// Secondary DNS
	SecondaryDns *string `json:"secondaryDns,omitempty"`
	// Product Type
	ProductType *string                `json:"productType,omitempty"`
	Components  map[string]interface{} `json:"components,omitempty"`
	// Model
	Model *string `json:"model,omitempty"`
	// Tags
	Tags []string `json:"tags,omitempty"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationDevicesStatuses200Response) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationDevicesStatuses200Response) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationDevicesStatuses200Response) SetMac(v string) {
	o.Mac = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *GetOrganizationDevicesStatuses200Response) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetNetworkId() string {
	if o == nil || o.NetworkId == nil {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetNetworkIdOk() (*string, bool) {
	if o == nil || o.NetworkId == nil {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasNetworkId() bool {
	if o != nil && o.NetworkId != nil {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationDevicesStatuses200Response) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationDevicesStatuses200Response) SetStatus(v string) {
	o.Status = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetLastReportedAt() string {
	if o == nil || o.LastReportedAt == nil {
		var ret string
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetLastReportedAtOk() (*string, bool) {
	if o == nil || o.LastReportedAt == nil {
		return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasLastReportedAt() bool {
	if o != nil && o.LastReportedAt != nil {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given string and assigns it to the LastReportedAt field.
func (o *GetOrganizationDevicesStatuses200Response) SetLastReportedAt(v string) {
	o.LastReportedAt = &v
}

// GetLanIp returns the LanIp field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetLanIp() string {
	if o == nil || o.LanIp == nil {
		var ret string
		return ret
	}
	return *o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetLanIpOk() (*string, bool) {
	if o == nil || o.LanIp == nil {
		return nil, false
	}
	return o.LanIp, true
}

// HasLanIp returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasLanIp() bool {
	if o != nil && o.LanIp != nil {
		return true
	}

	return false
}

// SetLanIp gets a reference to the given string and assigns it to the LanIp field.
func (o *GetOrganizationDevicesStatuses200Response) SetLanIp(v string) {
	o.LanIp = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetOrganizationDevicesStatuses200Response) SetGateway(v string) {
	o.Gateway = &v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetIpType() string {
	if o == nil || o.IpType == nil {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetIpTypeOk() (*string, bool) {
	if o == nil || o.IpType == nil {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasIpType() bool {
	if o != nil && o.IpType != nil {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *GetOrganizationDevicesStatuses200Response) SetIpType(v string) {
	o.IpType = &v
}

// GetPrimaryDns returns the PrimaryDns field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetPrimaryDns() string {
	if o == nil || o.PrimaryDns == nil {
		var ret string
		return ret
	}
	return *o.PrimaryDns
}

// GetPrimaryDnsOk returns a tuple with the PrimaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetPrimaryDnsOk() (*string, bool) {
	if o == nil || o.PrimaryDns == nil {
		return nil, false
	}
	return o.PrimaryDns, true
}

// HasPrimaryDns returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasPrimaryDns() bool {
	if o != nil && o.PrimaryDns != nil {
		return true
	}

	return false
}

// SetPrimaryDns gets a reference to the given string and assigns it to the PrimaryDns field.
func (o *GetOrganizationDevicesStatuses200Response) SetPrimaryDns(v string) {
	o.PrimaryDns = &v
}

// GetSecondaryDns returns the SecondaryDns field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetSecondaryDns() string {
	if o == nil || o.SecondaryDns == nil {
		var ret string
		return ret
	}
	return *o.SecondaryDns
}

// GetSecondaryDnsOk returns a tuple with the SecondaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetSecondaryDnsOk() (*string, bool) {
	if o == nil || o.SecondaryDns == nil {
		return nil, false
	}
	return o.SecondaryDns, true
}

// HasSecondaryDns returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasSecondaryDns() bool {
	if o != nil && o.SecondaryDns != nil {
		return true
	}

	return false
}

// SetSecondaryDns gets a reference to the given string and assigns it to the SecondaryDns field.
func (o *GetOrganizationDevicesStatuses200Response) SetSecondaryDns(v string) {
	o.SecondaryDns = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetProductType() string {
	if o == nil || o.ProductType == nil {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetProductTypeOk() (*string, bool) {
	if o == nil || o.ProductType == nil {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasProductType() bool {
	if o != nil && o.ProductType != nil {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetOrganizationDevicesStatuses200Response) SetProductType(v string) {
	o.ProductType = &v
}

// HasComponents returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationDevicesStatuses200Response) SetModel(v string) {
	o.Model = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200Response) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200Response) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200Response) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetOrganizationDevicesStatuses200Response) SetTags(v []string) {
	o.Tags = v
}
