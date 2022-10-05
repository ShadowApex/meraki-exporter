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

// UpdateOrganizationBrandingPoliciesPrioritiesRequest struct for UpdateOrganizationBrandingPoliciesPrioritiesRequest
type UpdateOrganizationBrandingPoliciesPrioritiesRequest struct {
	// A list of branding policy IDs arranged in ascending priority order (IDs later in the array have higher priority).
	BrandingPolicyIds []string `json:"brandingPolicyIds"`
}

// NewUpdateOrganizationBrandingPoliciesPrioritiesRequest instantiates a new UpdateOrganizationBrandingPoliciesPrioritiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationBrandingPoliciesPrioritiesRequest(brandingPolicyIds []string) *UpdateOrganizationBrandingPoliciesPrioritiesRequest {
	this := UpdateOrganizationBrandingPoliciesPrioritiesRequest{}
	this.BrandingPolicyIds = brandingPolicyIds
	return &this
}

// NewUpdateOrganizationBrandingPoliciesPrioritiesRequestWithDefaults instantiates a new UpdateOrganizationBrandingPoliciesPrioritiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationBrandingPoliciesPrioritiesRequestWithDefaults() *UpdateOrganizationBrandingPoliciesPrioritiesRequest {
	this := UpdateOrganizationBrandingPoliciesPrioritiesRequest{}
	return &this
}

// GetBrandingPolicyIds returns the BrandingPolicyIds field value
func (o *UpdateOrganizationBrandingPoliciesPrioritiesRequest) GetBrandingPolicyIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BrandingPolicyIds
}

// GetBrandingPolicyIdsOk returns a tuple with the BrandingPolicyIds field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationBrandingPoliciesPrioritiesRequest) GetBrandingPolicyIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandingPolicyIds, true
}

// SetBrandingPolicyIds sets field value
func (o *UpdateOrganizationBrandingPoliciesPrioritiesRequest) SetBrandingPolicyIds(v []string) {
	o.BrandingPolicyIds = v
}

func (o UpdateOrganizationBrandingPoliciesPrioritiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["brandingPolicyIds"] = o.BrandingPolicyIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest struct {
	value *UpdateOrganizationBrandingPoliciesPrioritiesRequest
	isSet bool
}

func (v NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) Get() *UpdateOrganizationBrandingPoliciesPrioritiesRequest {
	return v.value
}

func (v *NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) Set(val *UpdateOrganizationBrandingPoliciesPrioritiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationBrandingPoliciesPrioritiesRequest(val *UpdateOrganizationBrandingPoliciesPrioritiesRequest) *NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest {
	return &NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationBrandingPoliciesPrioritiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


