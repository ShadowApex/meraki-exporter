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

// CreateOrganizationBrandingPolicyRequestCustomLogo Properties describing the custom logo attached to the branding policy.
type CreateOrganizationBrandingPolicyRequestCustomLogo struct {
	// Whether or not there is a custom logo enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Properties for setting the image.
	Image map[string]interface{} `json:"image,omitempty"`
	// The file contents (a base 64 encoded string) of your new logo.
	Contents *string `json:"contents,omitempty"`
	// The format of the encoded contents.  Supported formats are 'png', 'gif', and jpg'. Note that all images are saved as PNG files, regardless of the format they are uploaded in.
	Format *string `json:"format,omitempty"`
}

// NewCreateOrganizationBrandingPolicyRequestCustomLogo instantiates a new CreateOrganizationBrandingPolicyRequestCustomLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationBrandingPolicyRequestCustomLogo() *CreateOrganizationBrandingPolicyRequestCustomLogo {
	this := CreateOrganizationBrandingPolicyRequestCustomLogo{}
	return &this
}

// NewCreateOrganizationBrandingPolicyRequestCustomLogoWithDefaults instantiates a new CreateOrganizationBrandingPolicyRequestCustomLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationBrandingPolicyRequestCustomLogoWithDefaults() *CreateOrganizationBrandingPolicyRequestCustomLogo {
	this := CreateOrganizationBrandingPolicyRequestCustomLogo{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetImage() map[string]interface{} {
	if o == nil || o.Image == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetImageOk() (map[string]interface{}, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given map[string]interface{} and assigns it to the Image field.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetImage(v map[string]interface{}) {
	o.Image = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetContents() string {
	if o == nil || o.Contents == nil {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetContentsOk() (*string, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetContents(v string) {
	o.Contents = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreateOrganizationBrandingPolicyRequestCustomLogo) SetFormat(v string) {
	o.Format = &v
}

func (o CreateOrganizationBrandingPolicyRequestCustomLogo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationBrandingPolicyRequestCustomLogo struct {
	value *CreateOrganizationBrandingPolicyRequestCustomLogo
	isSet bool
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogo) Get() *CreateOrganizationBrandingPolicyRequestCustomLogo {
	return v.value
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogo) Set(val *CreateOrganizationBrandingPolicyRequestCustomLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationBrandingPolicyRequestCustomLogo(val *CreateOrganizationBrandingPolicyRequestCustomLogo) *NullableCreateOrganizationBrandingPolicyRequestCustomLogo {
	return &NullableCreateOrganizationBrandingPolicyRequestCustomLogo{value: val, isSet: true}
}

func (v NullableCreateOrganizationBrandingPolicyRequestCustomLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationBrandingPolicyRequestCustomLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


