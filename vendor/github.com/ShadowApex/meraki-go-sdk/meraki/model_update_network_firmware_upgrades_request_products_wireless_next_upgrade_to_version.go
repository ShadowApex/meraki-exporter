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

// UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion The version to be updated to
type UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion struct {
	// The version ID
	Id *string `json:"id,omitempty"`
}

// NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion instantiates a new UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	this := UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion{}
	return &this
}

// NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersionWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersionWithDefaults() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	this := UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

func (o UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion struct {
	value *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) Get() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) Set(val *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion(val *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	return &NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


