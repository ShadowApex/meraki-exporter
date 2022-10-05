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

// GetNetworkFirmwareUpgrades200Response struct for GetNetworkFirmwareUpgrades200Response
type GetNetworkFirmwareUpgrades200Response struct {
	UpgradeWindow *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow `json:"upgradeWindow,omitempty"`
	// The timezone for the network
	Timezone *string `json:"timezone,omitempty"`
	Products *GetNetworkFirmwareUpgrades200ResponseProducts `json:"products,omitempty"`
}

// NewGetNetworkFirmwareUpgrades200Response instantiates a new GetNetworkFirmwareUpgrades200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgrades200Response() *GetNetworkFirmwareUpgrades200Response {
	this := GetNetworkFirmwareUpgrades200Response{}
	return &this
}

// NewGetNetworkFirmwareUpgrades200ResponseWithDefaults instantiates a new GetNetworkFirmwareUpgrades200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgrades200ResponseWithDefaults() *GetNetworkFirmwareUpgrades200Response {
	this := GetNetworkFirmwareUpgrades200Response{}
	return &this
}

// GetUpgradeWindow returns the UpgradeWindow field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200Response) GetUpgradeWindow() GetNetworkFirmwareUpgrades200ResponseUpgradeWindow {
	if o == nil || o.UpgradeWindow == nil {
		var ret GetNetworkFirmwareUpgrades200ResponseUpgradeWindow
		return ret
	}
	return *o.UpgradeWindow
}

// GetUpgradeWindowOk returns a tuple with the UpgradeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200Response) GetUpgradeWindowOk() (*GetNetworkFirmwareUpgrades200ResponseUpgradeWindow, bool) {
	if o == nil || o.UpgradeWindow == nil {
		return nil, false
	}
	return o.UpgradeWindow, true
}

// HasUpgradeWindow returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200Response) HasUpgradeWindow() bool {
	if o != nil && o.UpgradeWindow != nil {
		return true
	}

	return false
}

// SetUpgradeWindow gets a reference to the given GetNetworkFirmwareUpgrades200ResponseUpgradeWindow and assigns it to the UpgradeWindow field.
func (o *GetNetworkFirmwareUpgrades200Response) SetUpgradeWindow(v GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) {
	o.UpgradeWindow = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200Response) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200Response) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200Response) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *GetNetworkFirmwareUpgrades200Response) SetTimezone(v string) {
	o.Timezone = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200Response) GetProducts() GetNetworkFirmwareUpgrades200ResponseProducts {
	if o == nil || o.Products == nil {
		var ret GetNetworkFirmwareUpgrades200ResponseProducts
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200Response) GetProductsOk() (*GetNetworkFirmwareUpgrades200ResponseProducts, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200Response) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given GetNetworkFirmwareUpgrades200ResponseProducts and assigns it to the Products field.
func (o *GetNetworkFirmwareUpgrades200Response) SetProducts(v GetNetworkFirmwareUpgrades200ResponseProducts) {
	o.Products = &v
}

func (o GetNetworkFirmwareUpgrades200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpgradeWindow != nil {
		toSerialize["upgradeWindow"] = o.UpgradeWindow
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkFirmwareUpgrades200Response struct {
	value *GetNetworkFirmwareUpgrades200Response
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgrades200Response) Get() *GetNetworkFirmwareUpgrades200Response {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgrades200Response) Set(val *GetNetworkFirmwareUpgrades200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgrades200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgrades200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgrades200Response(val *GetNetworkFirmwareUpgrades200Response) *NullableGetNetworkFirmwareUpgrades200Response {
	return &NullableGetNetworkFirmwareUpgrades200Response{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgrades200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgrades200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


