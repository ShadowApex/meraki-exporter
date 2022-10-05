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

// UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner struct for UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner
type UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner struct {
	// Starting IP included in the reserved range of IPs
	Start string `json:"start"`
	// Ending IP included in the reserved range of IPs
	End string `json:"end"`
	// Comment explaining the reserved IP range
	Comment string `json:"comment"`
}

// NewUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner instantiates a new UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner(start string, end string, comment string) *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner {
	this := UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner{}
	this.Start = start
	this.End = end
	this.Comment = comment
	return &this
}

// NewUpdateDeviceCellularGatewayLanRequestReservedIpRangesInnerWithDefaults instantiates a new UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCellularGatewayLanRequestReservedIpRangesInnerWithDefaults() *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner {
	this := UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner{}
	return &this
}

// GetStart returns the Start field value
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) SetEnd(v string) {
	o.End = v
}

// GetComment returns the Comment field value
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) SetComment(v string) {
	o.Comment = v
}

func (o UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	if true {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner struct {
	value *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner
	isSet bool
}

func (v NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) Get() *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner {
	return v.value
}

func (v *NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) Set(val *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner(val *UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) *NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner {
	return &NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner{value: val, isSet: true}
}

func (v NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


