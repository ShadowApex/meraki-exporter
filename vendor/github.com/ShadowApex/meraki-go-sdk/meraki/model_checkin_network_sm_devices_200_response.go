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

// CheckinNetworkSmDevices200Response struct for CheckinNetworkSmDevices200Response
type CheckinNetworkSmDevices200Response struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
}

// NewCheckinNetworkSmDevices200Response instantiates a new CheckinNetworkSmDevices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckinNetworkSmDevices200Response() *CheckinNetworkSmDevices200Response {
	this := CheckinNetworkSmDevices200Response{}
	return &this
}

// NewCheckinNetworkSmDevices200ResponseWithDefaults instantiates a new CheckinNetworkSmDevices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckinNetworkSmDevices200ResponseWithDefaults() *CheckinNetworkSmDevices200Response {
	this := CheckinNetworkSmDevices200Response{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CheckinNetworkSmDevices200Response) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckinNetworkSmDevices200Response) GetIdsOk() ([]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CheckinNetworkSmDevices200Response) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CheckinNetworkSmDevices200Response) SetIds(v []string) {
	o.Ids = v
}

func (o CheckinNetworkSmDevices200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableCheckinNetworkSmDevices200Response struct {
	value *CheckinNetworkSmDevices200Response
	isSet bool
}

func (v NullableCheckinNetworkSmDevices200Response) Get() *CheckinNetworkSmDevices200Response {
	return v.value
}

func (v *NullableCheckinNetworkSmDevices200Response) Set(val *CheckinNetworkSmDevices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckinNetworkSmDevices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckinNetworkSmDevices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckinNetworkSmDevices200Response(val *CheckinNetworkSmDevices200Response) *NullableCheckinNetworkSmDevices200Response {
	return &NullableCheckinNetworkSmDevices200Response{value: val, isSet: true}
}

func (v NullableCheckinNetworkSmDevices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckinNetworkSmDevices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


