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

// LockNetworkSmDevicesRequest struct for LockNetworkSmDevicesRequest
type LockNetworkSmDevicesRequest struct {
	// The wifiMacs of the devices to be locked.
	WifiMacs []string `json:"wifiMacs,omitempty"`
	// The ids of the devices to be locked.
	Ids []string `json:"ids,omitempty"`
	// The serials of the devices to be locked.
	Serials []string `json:"serials,omitempty"`
	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be wiped.
	Scope []string `json:"scope,omitempty"`
	// The pin number for locking macOS devices (a six digit number). Required only for macOS devices.
	Pin *int32 `json:"pin,omitempty"`
}

// NewLockNetworkSmDevicesRequest instantiates a new LockNetworkSmDevicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockNetworkSmDevicesRequest() *LockNetworkSmDevicesRequest {
	this := LockNetworkSmDevicesRequest{}
	return &this
}

// NewLockNetworkSmDevicesRequestWithDefaults instantiates a new LockNetworkSmDevicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockNetworkSmDevicesRequestWithDefaults() *LockNetworkSmDevicesRequest {
	this := LockNetworkSmDevicesRequest{}
	return &this
}

// GetWifiMacs returns the WifiMacs field value if set, zero value otherwise.
func (o *LockNetworkSmDevicesRequest) GetWifiMacs() []string {
	if o == nil || o.WifiMacs == nil {
		var ret []string
		return ret
	}
	return o.WifiMacs
}

// GetWifiMacsOk returns a tuple with the WifiMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockNetworkSmDevicesRequest) GetWifiMacsOk() ([]string, bool) {
	if o == nil || o.WifiMacs == nil {
		return nil, false
	}
	return o.WifiMacs, true
}

// HasWifiMacs returns a boolean if a field has been set.
func (o *LockNetworkSmDevicesRequest) HasWifiMacs() bool {
	if o != nil && o.WifiMacs != nil {
		return true
	}

	return false
}

// SetWifiMacs gets a reference to the given []string and assigns it to the WifiMacs field.
func (o *LockNetworkSmDevicesRequest) SetWifiMacs(v []string) {
	o.WifiMacs = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *LockNetworkSmDevicesRequest) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockNetworkSmDevicesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *LockNetworkSmDevicesRequest) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *LockNetworkSmDevicesRequest) SetIds(v []string) {
	o.Ids = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *LockNetworkSmDevicesRequest) GetSerials() []string {
	if o == nil || o.Serials == nil {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockNetworkSmDevicesRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || o.Serials == nil {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *LockNetworkSmDevicesRequest) HasSerials() bool {
	if o != nil && o.Serials != nil {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *LockNetworkSmDevicesRequest) SetSerials(v []string) {
	o.Serials = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *LockNetworkSmDevicesRequest) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockNetworkSmDevicesRequest) GetScopeOk() ([]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *LockNetworkSmDevicesRequest) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *LockNetworkSmDevicesRequest) SetScope(v []string) {
	o.Scope = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *LockNetworkSmDevicesRequest) GetPin() int32 {
	if o == nil || o.Pin == nil {
		var ret int32
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockNetworkSmDevicesRequest) GetPinOk() (*int32, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *LockNetworkSmDevicesRequest) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given int32 and assigns it to the Pin field.
func (o *LockNetworkSmDevicesRequest) SetPin(v int32) {
	o.Pin = &v
}

func (o LockNetworkSmDevicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WifiMacs != nil {
		toSerialize["wifiMacs"] = o.WifiMacs
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Serials != nil {
		toSerialize["serials"] = o.Serials
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Pin != nil {
		toSerialize["pin"] = o.Pin
	}
	return json.Marshal(toSerialize)
}

type NullableLockNetworkSmDevicesRequest struct {
	value *LockNetworkSmDevicesRequest
	isSet bool
}

func (v NullableLockNetworkSmDevicesRequest) Get() *LockNetworkSmDevicesRequest {
	return v.value
}

func (v *NullableLockNetworkSmDevicesRequest) Set(val *LockNetworkSmDevicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLockNetworkSmDevicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLockNetworkSmDevicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockNetworkSmDevicesRequest(val *LockNetworkSmDevicesRequest) *NullableLockNetworkSmDevicesRequest {
	return &NullableLockNetworkSmDevicesRequest{value: val, isSet: true}
}

func (v NullableLockNetworkSmDevicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockNetworkSmDevicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


