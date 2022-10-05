/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 July, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.23.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"encoding/json"
	"time"
)

// GetAdministeredIdentitiesMe200Response struct for GetAdministeredIdentitiesMe200Response
type GetAdministeredIdentitiesMe200Response struct {
	// Username
	Name *string `json:"name,omitempty"`
	// User email
	Email *string `json:"email,omitempty"`
	// Last seen active on Dashboard UI
	LastUsedDashboardAt *time.Time `json:"lastUsedDashboardAt,omitempty"`
	Authentication *GetAdministeredIdentitiesMe200ResponseAuthentication `json:"authentication,omitempty"`
}

// NewGetAdministeredIdentitiesMe200Response instantiates a new GetAdministeredIdentitiesMe200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredIdentitiesMe200Response() *GetAdministeredIdentitiesMe200Response {
	this := GetAdministeredIdentitiesMe200Response{}
	return &this
}

// NewGetAdministeredIdentitiesMe200ResponseWithDefaults instantiates a new GetAdministeredIdentitiesMe200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredIdentitiesMe200ResponseWithDefaults() *GetAdministeredIdentitiesMe200Response {
	this := GetAdministeredIdentitiesMe200Response{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200Response) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200Response) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAdministeredIdentitiesMe200Response) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200Response) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200Response) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200Response) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetAdministeredIdentitiesMe200Response) SetEmail(v string) {
	o.Email = &v
}

// GetLastUsedDashboardAt returns the LastUsedDashboardAt field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200Response) GetLastUsedDashboardAt() time.Time {
	if o == nil || o.LastUsedDashboardAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedDashboardAt
}

// GetLastUsedDashboardAtOk returns a tuple with the LastUsedDashboardAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200Response) GetLastUsedDashboardAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedDashboardAt == nil {
		return nil, false
	}
	return o.LastUsedDashboardAt, true
}

// HasLastUsedDashboardAt returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200Response) HasLastUsedDashboardAt() bool {
	if o != nil && o.LastUsedDashboardAt != nil {
		return true
	}

	return false
}

// SetLastUsedDashboardAt gets a reference to the given time.Time and assigns it to the LastUsedDashboardAt field.
func (o *GetAdministeredIdentitiesMe200Response) SetLastUsedDashboardAt(v time.Time) {
	o.LastUsedDashboardAt = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200Response) GetAuthentication() GetAdministeredIdentitiesMe200ResponseAuthentication {
	if o == nil || o.Authentication == nil {
		var ret GetAdministeredIdentitiesMe200ResponseAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200Response) GetAuthenticationOk() (*GetAdministeredIdentitiesMe200ResponseAuthentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200Response) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given GetAdministeredIdentitiesMe200ResponseAuthentication and assigns it to the Authentication field.
func (o *GetAdministeredIdentitiesMe200Response) SetAuthentication(v GetAdministeredIdentitiesMe200ResponseAuthentication) {
	o.Authentication = &v
}

func (o GetAdministeredIdentitiesMe200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.LastUsedDashboardAt != nil {
		toSerialize["lastUsedDashboardAt"] = o.LastUsedDashboardAt
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableGetAdministeredIdentitiesMe200Response struct {
	value *GetAdministeredIdentitiesMe200Response
	isSet bool
}

func (v NullableGetAdministeredIdentitiesMe200Response) Get() *GetAdministeredIdentitiesMe200Response {
	return v.value
}

func (v *NullableGetAdministeredIdentitiesMe200Response) Set(val *GetAdministeredIdentitiesMe200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredIdentitiesMe200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredIdentitiesMe200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredIdentitiesMe200Response(val *GetAdministeredIdentitiesMe200Response) *NullableGetAdministeredIdentitiesMe200Response {
	return &NullableGetAdministeredIdentitiesMe200Response{value: val, isSet: true}
}

func (v NullableGetAdministeredIdentitiesMe200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredIdentitiesMe200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


