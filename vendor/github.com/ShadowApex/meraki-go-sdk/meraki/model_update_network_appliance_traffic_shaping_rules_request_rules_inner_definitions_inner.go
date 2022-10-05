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

// UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner struct for UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner
type UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner struct {
	// The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	Type string `json:"type"`
	//     If \"type\" is 'host', 'port', 'ipRange' or 'localNet', then \"value\" must be a string, matching either     a hostname (e.g. \"somesite.com\"), a port (e.g. 8080), or an IP range (\"192.1.0.0\",     \"192.1.0.0/16\", or \"10.1.0.0/16:80\"). 'localNet' also supports CIDR notation, excluding     custom ports.      If \"type\" is 'application' or 'applicationCategory', then \"value\" must be an object     with the structure { \"id\": \"meraki:layer7/...\" }, where \"id\" is the application category or     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories     endpoint). 
	Value string `json:"value"`
}

// NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner(type_ string, value string) *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner {
	this := UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInnerWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInnerWithDefaults() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner {
	this := UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) SetValue(v string) {
	o.Value = v
}

func (o UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner struct {
	value *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) Get() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) Set(val *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner(val *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner {
	return &NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


