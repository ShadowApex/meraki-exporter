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

// GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment Systems Manager sentry enrollment splash settings.
type GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment struct {
	SystemsManagerNetwork *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollmentSystemsManagerNetwork `json:"systemsManagerNetwork,omitempty"`
	// The strength of the enforcement of selected system types.
	Strength *string `json:"strength,omitempty"`
	// The system types that the Sentry enforces.
	EnforcedSystems []string `json:"enforcedSystems,omitempty"`
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment instantiates a new GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment() *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment {
	this := GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment{}
	return &this
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollmentWithDefaults instantiates a new GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollmentWithDefaults() *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment {
	this := GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment{}
	return &this
}

// GetSystemsManagerNetwork returns the SystemsManagerNetwork field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) GetSystemsManagerNetwork() GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollmentSystemsManagerNetwork {
	if o == nil || o.SystemsManagerNetwork == nil {
		var ret GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollmentSystemsManagerNetwork
		return ret
	}
	return *o.SystemsManagerNetwork
}

// GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) GetSystemsManagerNetworkOk() (*GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollmentSystemsManagerNetwork, bool) {
	if o == nil || o.SystemsManagerNetwork == nil {
		return nil, false
	}
	return o.SystemsManagerNetwork, true
}

// HasSystemsManagerNetwork returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) HasSystemsManagerNetwork() bool {
	if o != nil && o.SystemsManagerNetwork != nil {
		return true
	}

	return false
}

// SetSystemsManagerNetwork gets a reference to the given GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollmentSystemsManagerNetwork and assigns it to the SystemsManagerNetwork field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) SetSystemsManagerNetwork(v GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollmentSystemsManagerNetwork) {
	o.SystemsManagerNetwork = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) GetStrength() string {
	if o == nil || o.Strength == nil {
		var ret string
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) GetStrengthOk() (*string, bool) {
	if o == nil || o.Strength == nil {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) HasStrength() bool {
	if o != nil && o.Strength != nil {
		return true
	}

	return false
}

// SetStrength gets a reference to the given string and assigns it to the Strength field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) SetStrength(v string) {
	o.Strength = &v
}

// GetEnforcedSystems returns the EnforcedSystems field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) GetEnforcedSystems() []string {
	if o == nil || o.EnforcedSystems == nil {
		var ret []string
		return ret
	}
	return o.EnforcedSystems
}

// GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) GetEnforcedSystemsOk() ([]string, bool) {
	if o == nil || o.EnforcedSystems == nil {
		return nil, false
	}
	return o.EnforcedSystems, true
}

// HasEnforcedSystems returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) HasEnforcedSystems() bool {
	if o != nil && o.EnforcedSystems != nil {
		return true
	}

	return false
}

// SetEnforcedSystems gets a reference to the given []string and assigns it to the EnforcedSystems field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) SetEnforcedSystems(v []string) {
	o.EnforcedSystems = v
}

func (o GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SystemsManagerNetwork != nil {
		toSerialize["systemsManagerNetwork"] = o.SystemsManagerNetwork
	}
	if o.Strength != nil {
		toSerialize["strength"] = o.Strength
	}
	if o.EnforcedSystems != nil {
		toSerialize["enforcedSystems"] = o.EnforcedSystems
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment struct {
	value *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment
	isSet bool
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) Get() *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) Set(val *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment(val *GetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) *NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment {
	return &NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSentryEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


