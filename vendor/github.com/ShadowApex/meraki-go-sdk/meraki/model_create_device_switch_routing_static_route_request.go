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

// CreateDeviceSwitchRoutingStaticRouteRequest struct for CreateDeviceSwitchRoutingStaticRouteRequest
type CreateDeviceSwitchRoutingStaticRouteRequest struct {
	// Name or description for layer 3 static route
	Name *string `json:"name,omitempty"`
	// The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
	Subnet string `json:"subnet"`
	// IP address of the next hop device to which the device sends its traffic for the subnet
	NextHopIp string `json:"nextHopIp"`
	// Option to advertise static route via OSPF
	AdvertiseViaOspfEnabled *bool `json:"advertiseViaOspfEnabled,omitempty"`
	// Option to prefer static route over OSPF routes
	PreferOverOspfRoutesEnabled *bool `json:"preferOverOspfRoutesEnabled,omitempty"`
}

// NewCreateDeviceSwitchRoutingStaticRouteRequest instantiates a new CreateDeviceSwitchRoutingStaticRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceSwitchRoutingStaticRouteRequest(subnet string, nextHopIp string) *CreateDeviceSwitchRoutingStaticRouteRequest {
	this := CreateDeviceSwitchRoutingStaticRouteRequest{}
	this.Subnet = subnet
	this.NextHopIp = nextHopIp
	return &this
}

// NewCreateDeviceSwitchRoutingStaticRouteRequestWithDefaults instantiates a new CreateDeviceSwitchRoutingStaticRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceSwitchRoutingStaticRouteRequestWithDefaults() *CreateDeviceSwitchRoutingStaticRouteRequest {
	this := CreateDeviceSwitchRoutingStaticRouteRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetSubnet(v string) {
	o.Subnet = v
}

// GetNextHopIp returns the NextHopIp field value
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetNextHopIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetNextHopIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextHopIp, true
}

// SetNextHopIp sets field value
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetNextHopIp(v string) {
	o.NextHopIp = v
}

// GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetAdvertiseViaOspfEnabled() bool {
	if o == nil || o.AdvertiseViaOspfEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AdvertiseViaOspfEnabled
}

// GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetAdvertiseViaOspfEnabledOk() (*bool, bool) {
	if o == nil || o.AdvertiseViaOspfEnabled == nil {
		return nil, false
	}
	return o.AdvertiseViaOspfEnabled, true
}

// HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) HasAdvertiseViaOspfEnabled() bool {
	if o != nil && o.AdvertiseViaOspfEnabled != nil {
		return true
	}

	return false
}

// SetAdvertiseViaOspfEnabled gets a reference to the given bool and assigns it to the AdvertiseViaOspfEnabled field.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetAdvertiseViaOspfEnabled(v bool) {
	o.AdvertiseViaOspfEnabled = &v
}

// GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field value if set, zero value otherwise.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetPreferOverOspfRoutesEnabled() bool {
	if o == nil || o.PreferOverOspfRoutesEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PreferOverOspfRoutesEnabled
}

// GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) GetPreferOverOspfRoutesEnabledOk() (*bool, bool) {
	if o == nil || o.PreferOverOspfRoutesEnabled == nil {
		return nil, false
	}
	return o.PreferOverOspfRoutesEnabled, true
}

// HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) HasPreferOverOspfRoutesEnabled() bool {
	if o != nil && o.PreferOverOspfRoutesEnabled != nil {
		return true
	}

	return false
}

// SetPreferOverOspfRoutesEnabled gets a reference to the given bool and assigns it to the PreferOverOspfRoutesEnabled field.
func (o *CreateDeviceSwitchRoutingStaticRouteRequest) SetPreferOverOspfRoutesEnabled(v bool) {
	o.PreferOverOspfRoutesEnabled = &v
}

func (o CreateDeviceSwitchRoutingStaticRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["subnet"] = o.Subnet
	}
	if true {
		toSerialize["nextHopIp"] = o.NextHopIp
	}
	if o.AdvertiseViaOspfEnabled != nil {
		toSerialize["advertiseViaOspfEnabled"] = o.AdvertiseViaOspfEnabled
	}
	if o.PreferOverOspfRoutesEnabled != nil {
		toSerialize["preferOverOspfRoutesEnabled"] = o.PreferOverOspfRoutesEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDeviceSwitchRoutingStaticRouteRequest struct {
	value *CreateDeviceSwitchRoutingStaticRouteRequest
	isSet bool
}

func (v NullableCreateDeviceSwitchRoutingStaticRouteRequest) Get() *CreateDeviceSwitchRoutingStaticRouteRequest {
	return v.value
}

func (v *NullableCreateDeviceSwitchRoutingStaticRouteRequest) Set(val *CreateDeviceSwitchRoutingStaticRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceSwitchRoutingStaticRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceSwitchRoutingStaticRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceSwitchRoutingStaticRouteRequest(val *CreateDeviceSwitchRoutingStaticRouteRequest) *NullableCreateDeviceSwitchRoutingStaticRouteRequest {
	return &NullableCreateDeviceSwitchRoutingStaticRouteRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceSwitchRoutingStaticRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceSwitchRoutingStaticRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


