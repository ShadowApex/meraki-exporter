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

// CreateNetworkWebhooksPayloadTemplateRequest struct for CreateNetworkWebhooksPayloadTemplateRequest
type CreateNetworkWebhooksPayloadTemplateRequest struct {
	// The name of the new template
	Name string `json:"name"`
	// The liquid template used for the body of the webhook message. Either `body` or `bodyFile` must be specified.
	Body *string `json:"body,omitempty"`
	// The liquid template used with the webhook headers.
	Headers *string `json:"headers,omitempty"`
	// A file containing liquid template used for the body of the webhook message. Either `body` or `bodyFile` must be specified.
	BodyFile *string `json:"bodyFile,omitempty"`
	// A file containing the liquid template used with the webhook headers.
	HeadersFile *string `json:"headersFile,omitempty"`
}

// NewCreateNetworkWebhooksPayloadTemplateRequest instantiates a new CreateNetworkWebhooksPayloadTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWebhooksPayloadTemplateRequest(name string) *CreateNetworkWebhooksPayloadTemplateRequest {
	this := CreateNetworkWebhooksPayloadTemplateRequest{}
	this.Name = name
	return &this
}

// NewCreateNetworkWebhooksPayloadTemplateRequestWithDefaults instantiates a new CreateNetworkWebhooksPayloadTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWebhooksPayloadTemplateRequestWithDefaults() *CreateNetworkWebhooksPayloadTemplateRequest {
	this := CreateNetworkWebhooksPayloadTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetBody(v string) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetHeaders() string {
	if o == nil || o.Headers == nil {
		var ret string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetHeadersOk() (*string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given string and assigns it to the Headers field.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetHeaders(v string) {
	o.Headers = &v
}

// GetBodyFile returns the BodyFile field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetBodyFile() string {
	if o == nil || o.BodyFile == nil {
		var ret string
		return ret
	}
	return *o.BodyFile
}

// GetBodyFileOk returns a tuple with the BodyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetBodyFileOk() (*string, bool) {
	if o == nil || o.BodyFile == nil {
		return nil, false
	}
	return o.BodyFile, true
}

// HasBodyFile returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) HasBodyFile() bool {
	if o != nil && o.BodyFile != nil {
		return true
	}

	return false
}

// SetBodyFile gets a reference to the given string and assigns it to the BodyFile field.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetBodyFile(v string) {
	o.BodyFile = &v
}

// GetHeadersFile returns the HeadersFile field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetHeadersFile() string {
	if o == nil || o.HeadersFile == nil {
		var ret string
		return ret
	}
	return *o.HeadersFile
}

// GetHeadersFileOk returns a tuple with the HeadersFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetHeadersFileOk() (*string, bool) {
	if o == nil || o.HeadersFile == nil {
		return nil, false
	}
	return o.HeadersFile, true
}

// HasHeadersFile returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) HasHeadersFile() bool {
	if o != nil && o.HeadersFile != nil {
		return true
	}

	return false
}

// SetHeadersFile gets a reference to the given string and assigns it to the HeadersFile field.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetHeadersFile(v string) {
	o.HeadersFile = &v
}

func (o CreateNetworkWebhooksPayloadTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.BodyFile != nil {
		toSerialize["bodyFile"] = o.BodyFile
	}
	if o.HeadersFile != nil {
		toSerialize["headersFile"] = o.HeadersFile
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkWebhooksPayloadTemplateRequest struct {
	value *CreateNetworkWebhooksPayloadTemplateRequest
	isSet bool
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequest) Get() *CreateNetworkWebhooksPayloadTemplateRequest {
	return v.value
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequest) Set(val *CreateNetworkWebhooksPayloadTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWebhooksPayloadTemplateRequest(val *CreateNetworkWebhooksPayloadTemplateRequest) *NullableCreateNetworkWebhooksPayloadTemplateRequest {
	return &NullableCreateNetworkWebhooksPayloadTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


