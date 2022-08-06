/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// WebhookDefinition struct for WebhookDefinition
type WebhookDefinition struct {
	ConnectionDefinition
	// URL for the webhook connection.
	Url string `json:"url"`
	// List of access authorization headers.
	Headers []Header `json:"headers,omitempty"`
	// List of custom webhook headers.
	CustomHeaders []Header `json:"customHeaders,omitempty"`
	// Default payload of the webhook.
	DefaultPayload string `json:"defaultPayload"`
	// Type of webhook connection. Valid values are `AWSLambda`, `Azure`, `Datadog`, `HipChat`, `Jira`, `NewRelic`, `Opsgenie`, `PagerDuty`, `Slack`, `MicrosoftTeams`, `ServiceNow`, `SumoCloudSOAR` and `Webhook`.
	WebhookType *string `json:"webhookType,omitempty"`
	// The subtype of the connection. Valid values are `Event` or `Incident`.
	ConnectionSubtype *string `json:"connectionSubtype,omitempty"`
}

// NewWebhookDefinition instantiates a new WebhookDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDefinition(url string, defaultPayload string, type_ string, name string) *WebhookDefinition {
	this := WebhookDefinition{}
	this.Type = type_
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Url = url
	this.DefaultPayload = defaultPayload
	return &this
}

// NewWebhookDefinitionWithDefaults instantiates a new WebhookDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDefinitionWithDefaults() *WebhookDefinition {
	this := WebhookDefinition{}
	return &this
}

// GetUrl returns the Url field value
func (o *WebhookDefinition) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookDefinition) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookDefinition) SetUrl(v string) {
	o.Url = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *WebhookDefinition) GetHeaders() []Header {
	if o == nil || o.Headers == nil {
		var ret []Header
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDefinition) GetHeadersOk() ([]Header, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *WebhookDefinition) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Header and assigns it to the Headers field.
func (o *WebhookDefinition) SetHeaders(v []Header) {
	o.Headers = v
}

// GetCustomHeaders returns the CustomHeaders field value if set, zero value otherwise.
func (o *WebhookDefinition) GetCustomHeaders() []Header {
	if o == nil || o.CustomHeaders == nil {
		var ret []Header
		return ret
	}
	return o.CustomHeaders
}

// GetCustomHeadersOk returns a tuple with the CustomHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDefinition) GetCustomHeadersOk() ([]Header, bool) {
	if o == nil || o.CustomHeaders == nil {
		return nil, false
	}
	return o.CustomHeaders, true
}

// HasCustomHeaders returns a boolean if a field has been set.
func (o *WebhookDefinition) HasCustomHeaders() bool {
	if o != nil && o.CustomHeaders != nil {
		return true
	}

	return false
}

// SetCustomHeaders gets a reference to the given []Header and assigns it to the CustomHeaders field.
func (o *WebhookDefinition) SetCustomHeaders(v []Header) {
	o.CustomHeaders = v
}

// GetDefaultPayload returns the DefaultPayload field value
func (o *WebhookDefinition) GetDefaultPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultPayload
}

// GetDefaultPayloadOk returns a tuple with the DefaultPayload field value
// and a boolean to check if the value has been set.
func (o *WebhookDefinition) GetDefaultPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPayload, true
}

// SetDefaultPayload sets field value
func (o *WebhookDefinition) SetDefaultPayload(v string) {
	o.DefaultPayload = v
}

// GetWebhookType returns the WebhookType field value if set, zero value otherwise.
func (o *WebhookDefinition) GetWebhookType() string {
	if o == nil || o.WebhookType == nil {
		var ret string
		return ret
	}
	return *o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDefinition) GetWebhookTypeOk() (*string, bool) {
	if o == nil || o.WebhookType == nil {
		return nil, false
	}
	return o.WebhookType, true
}

// HasWebhookType returns a boolean if a field has been set.
func (o *WebhookDefinition) HasWebhookType() bool {
	if o != nil && o.WebhookType != nil {
		return true
	}

	return false
}

// SetWebhookType gets a reference to the given string and assigns it to the WebhookType field.
func (o *WebhookDefinition) SetWebhookType(v string) {
	o.WebhookType = &v
}

// GetConnectionSubtype returns the ConnectionSubtype field value if set, zero value otherwise.
func (o *WebhookDefinition) GetConnectionSubtype() string {
	if o == nil || o.ConnectionSubtype == nil {
		var ret string
		return ret
	}
	return *o.ConnectionSubtype
}

// GetConnectionSubtypeOk returns a tuple with the ConnectionSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDefinition) GetConnectionSubtypeOk() (*string, bool) {
	if o == nil || o.ConnectionSubtype == nil {
		return nil, false
	}
	return o.ConnectionSubtype, true
}

// HasConnectionSubtype returns a boolean if a field has been set.
func (o *WebhookDefinition) HasConnectionSubtype() bool {
	if o != nil && o.ConnectionSubtype != nil {
		return true
	}

	return false
}

// SetConnectionSubtype gets a reference to the given string and assigns it to the ConnectionSubtype field.
func (o *WebhookDefinition) SetConnectionSubtype(v string) {
	o.ConnectionSubtype = &v
}

func (o WebhookDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectionDefinition, errConnectionDefinition := json.Marshal(o.ConnectionDefinition)
	if errConnectionDefinition != nil {
		return []byte{}, errConnectionDefinition
	}
	errConnectionDefinition = json.Unmarshal([]byte(serializedConnectionDefinition), &toSerialize)
	if errConnectionDefinition != nil {
		return []byte{}, errConnectionDefinition
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.CustomHeaders != nil {
		toSerialize["customHeaders"] = o.CustomHeaders
	}
	if true {
		toSerialize["defaultPayload"] = o.DefaultPayload
	}
	if o.WebhookType != nil {
		toSerialize["webhookType"] = o.WebhookType
	}
	if o.ConnectionSubtype != nil {
		toSerialize["connectionSubtype"] = o.ConnectionSubtype
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookDefinition struct {
	value *WebhookDefinition
	isSet bool
}

func (v NullableWebhookDefinition) Get() *WebhookDefinition {
	return v.value
}

func (v *NullableWebhookDefinition) Set(val *WebhookDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDefinition(val *WebhookDefinition) *NullableWebhookDefinition {
	return &NullableWebhookDefinition{value: val, isSet: true}
}

func (v NullableWebhookDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


