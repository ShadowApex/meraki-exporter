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

// UpdateNetworkTrafficAnalysisRequest struct for UpdateNetworkTrafficAnalysisRequest
type UpdateNetworkTrafficAnalysisRequest struct {
	//     The traffic analysis mode for the network. Can be one of 'disabled' (do not collect traffic types),     'basic' (collect generic traffic categories), or 'detailed' (collect destination hostnames). 
	Mode *string `json:"mode,omitempty"`
	// The list of items that make up the custom pie chart for traffic reporting.
	CustomPieChartItems []UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner `json:"customPieChartItems,omitempty"`
}

// NewUpdateNetworkTrafficAnalysisRequest instantiates a new UpdateNetworkTrafficAnalysisRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkTrafficAnalysisRequest() *UpdateNetworkTrafficAnalysisRequest {
	this := UpdateNetworkTrafficAnalysisRequest{}
	return &this
}

// NewUpdateNetworkTrafficAnalysisRequestWithDefaults instantiates a new UpdateNetworkTrafficAnalysisRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkTrafficAnalysisRequestWithDefaults() *UpdateNetworkTrafficAnalysisRequest {
	this := UpdateNetworkTrafficAnalysisRequest{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *UpdateNetworkTrafficAnalysisRequest) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkTrafficAnalysisRequest) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *UpdateNetworkTrafficAnalysisRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *UpdateNetworkTrafficAnalysisRequest) SetMode(v string) {
	o.Mode = &v
}

// GetCustomPieChartItems returns the CustomPieChartItems field value if set, zero value otherwise.
func (o *UpdateNetworkTrafficAnalysisRequest) GetCustomPieChartItems() []UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner {
	if o == nil || o.CustomPieChartItems == nil {
		var ret []UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner
		return ret
	}
	return o.CustomPieChartItems
}

// GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkTrafficAnalysisRequest) GetCustomPieChartItemsOk() ([]UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner, bool) {
	if o == nil || o.CustomPieChartItems == nil {
		return nil, false
	}
	return o.CustomPieChartItems, true
}

// HasCustomPieChartItems returns a boolean if a field has been set.
func (o *UpdateNetworkTrafficAnalysisRequest) HasCustomPieChartItems() bool {
	if o != nil && o.CustomPieChartItems != nil {
		return true
	}

	return false
}

// SetCustomPieChartItems gets a reference to the given []UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner and assigns it to the CustomPieChartItems field.
func (o *UpdateNetworkTrafficAnalysisRequest) SetCustomPieChartItems(v []UpdateNetworkTrafficAnalysisRequestCustomPieChartItemsInner) {
	o.CustomPieChartItems = v
}

func (o UpdateNetworkTrafficAnalysisRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.CustomPieChartItems != nil {
		toSerialize["customPieChartItems"] = o.CustomPieChartItems
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkTrafficAnalysisRequest struct {
	value *UpdateNetworkTrafficAnalysisRequest
	isSet bool
}

func (v NullableUpdateNetworkTrafficAnalysisRequest) Get() *UpdateNetworkTrafficAnalysisRequest {
	return v.value
}

func (v *NullableUpdateNetworkTrafficAnalysisRequest) Set(val *UpdateNetworkTrafficAnalysisRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkTrafficAnalysisRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkTrafficAnalysisRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkTrafficAnalysisRequest(val *UpdateNetworkTrafficAnalysisRequest) *NullableUpdateNetworkTrafficAnalysisRequest {
	return &NullableUpdateNetworkTrafficAnalysisRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkTrafficAnalysisRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkTrafficAnalysisRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


