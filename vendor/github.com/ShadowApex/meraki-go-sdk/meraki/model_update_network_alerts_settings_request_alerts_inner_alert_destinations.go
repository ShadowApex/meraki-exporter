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

// UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations A hash of destinations for this specific alert
type UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations struct {
	// A list of emails that will receive information about the alert
	Emails []string `json:"emails,omitempty"`
	// If true, then all network admins will receive emails for this alert
	AllAdmins *bool `json:"allAdmins,omitempty"`
	// If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network
	Snmp *bool `json:"snmp,omitempty"`
	// A list of HTTP server IDs to send a Webhook to for this alert
	HttpServerIds []string `json:"httpServerIds,omitempty"`
}

// NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations instantiates a new UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations() *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations {
	this := UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations{}
	return &this
}

// NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinationsWithDefaults instantiates a new UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinationsWithDefaults() *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations {
	this := UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetEmails() []string {
	if o == nil || o.Emails == nil {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetEmailsOk() ([]string, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) SetEmails(v []string) {
	o.Emails = v
}

// GetAllAdmins returns the AllAdmins field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetAllAdmins() bool {
	if o == nil || o.AllAdmins == nil {
		var ret bool
		return ret
	}
	return *o.AllAdmins
}

// GetAllAdminsOk returns a tuple with the AllAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetAllAdminsOk() (*bool, bool) {
	if o == nil || o.AllAdmins == nil {
		return nil, false
	}
	return o.AllAdmins, true
}

// HasAllAdmins returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) HasAllAdmins() bool {
	if o != nil && o.AllAdmins != nil {
		return true
	}

	return false
}

// SetAllAdmins gets a reference to the given bool and assigns it to the AllAdmins field.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) SetAllAdmins(v bool) {
	o.AllAdmins = &v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetSnmp() bool {
	if o == nil || o.Snmp == nil {
		var ret bool
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetSnmpOk() (*bool, bool) {
	if o == nil || o.Snmp == nil {
		return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) HasSnmp() bool {
	if o != nil && o.Snmp != nil {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given bool and assigns it to the Snmp field.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) SetSnmp(v bool) {
	o.Snmp = &v
}

// GetHttpServerIds returns the HttpServerIds field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetHttpServerIds() []string {
	if o == nil || o.HttpServerIds == nil {
		var ret []string
		return ret
	}
	return o.HttpServerIds
}

// GetHttpServerIdsOk returns a tuple with the HttpServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) GetHttpServerIdsOk() ([]string, bool) {
	if o == nil || o.HttpServerIds == nil {
		return nil, false
	}
	return o.HttpServerIds, true
}

// HasHttpServerIds returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) HasHttpServerIds() bool {
	if o != nil && o.HttpServerIds != nil {
		return true
	}

	return false
}

// SetHttpServerIds gets a reference to the given []string and assigns it to the HttpServerIds field.
func (o *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) SetHttpServerIds(v []string) {
	o.HttpServerIds = v
}

func (o UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.AllAdmins != nil {
		toSerialize["allAdmins"] = o.AllAdmins
	}
	if o.Snmp != nil {
		toSerialize["snmp"] = o.Snmp
	}
	if o.HttpServerIds != nil {
		toSerialize["httpServerIds"] = o.HttpServerIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations struct {
	value *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations
	isSet bool
}

func (v NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) Get() *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations {
	return v.value
}

func (v *NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) Set(val *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations(val *UpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) *NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations {
	return &NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations{value: val, isSet: true}
}

func (v NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkAlertsSettingsRequestAlertsInnerAlertDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

