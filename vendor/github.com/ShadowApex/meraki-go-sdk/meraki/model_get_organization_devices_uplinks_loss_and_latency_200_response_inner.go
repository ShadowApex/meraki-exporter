/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 July, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.23.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
)

// GetOrganizationDevicesUplinksLossAndLatency200ResponseInner struct for GetOrganizationDevicesUplinksLossAndLatency200ResponseInner
type GetOrganizationDevicesUplinksLossAndLatency200ResponseInner struct {
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Serial of MX device
	Serial *string `json:"serial,omitempty"`
	// Uplink interface (wan1, wan2, or cellular)
	Uplink *string `json:"uplink,omitempty"`
	// IP address of uplink
	Ip *string `json:"ip,omitempty"`
	// Loss and latency timeseries data
	TimeSeries []GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner `json:"timeSeries,omitempty"`
}

// NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInner instantiates a new GetOrganizationDevicesUplinksLossAndLatency200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInner() *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner {
	this := GetOrganizationDevicesUplinksLossAndLatency200ResponseInner{}
	return &this
}

// NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesUplinksLossAndLatency200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesUplinksLossAndLatency200ResponseInnerWithDefaults() *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner {
	this := GetOrganizationDevicesUplinksLossAndLatency200ResponseInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetNetworkId() string {
	if o == nil || o.NetworkId == nil {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || o.NetworkId == nil {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasNetworkId() bool {
	if o != nil && o.NetworkId != nil {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetUplink() string {
	if o == nil || o.Uplink == nil {
		var ret string
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetUplinkOk() (*string, bool) {
	if o == nil || o.Uplink == nil {
		return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasUplink() bool {
	if o != nil && o.Uplink != nil {
		return true
	}

	return false
}

// SetUplink gets a reference to the given string and assigns it to the Uplink field.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetUplink(v string) {
	o.Uplink = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetIp(v string) {
	o.Ip = &v
}

// GetTimeSeries returns the TimeSeries field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetTimeSeries() []GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner {
	if o == nil || o.TimeSeries == nil {
		var ret []GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner
		return ret
	}
	return o.TimeSeries
}

// GetTimeSeriesOk returns a tuple with the TimeSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) GetTimeSeriesOk() ([]GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner, bool) {
	if o == nil || o.TimeSeries == nil {
		return nil, false
	}
	return o.TimeSeries, true
}

// HasTimeSeries returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) HasTimeSeries() bool {
	if o != nil && o.TimeSeries != nil {
		return true
	}

	return false
}

// SetTimeSeries gets a reference to the given []GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner and assigns it to the TimeSeries field.
func (o *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) SetTimeSeries(v []GetOrganizationDevicesUplinksLossAndLatency200ResponseInnerTimeSeriesInner) {
	o.TimeSeries = v
}

func (o GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkId != nil {
		toSerialize["networkId"] = o.NetworkId
	}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if o.Uplink != nil {
		toSerialize["uplink"] = o.Uplink
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.TimeSeries != nil {
		toSerialize["timeSeries"] = o.TimeSeries
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner struct {
	value *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner) Get() *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner) Set(val *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner(val *GetOrganizationDevicesUplinksLossAndLatency200ResponseInner) *NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner {
	return &NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesUplinksLossAndLatency200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


