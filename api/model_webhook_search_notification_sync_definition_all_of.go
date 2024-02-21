/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the WebhookSearchNotificationSyncDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSearchNotificationSyncDefinitionAllOf{}

// WebhookSearchNotificationSyncDefinitionAllOf struct for WebhookSearchNotificationSyncDefinitionAllOf
type WebhookSearchNotificationSyncDefinitionAllOf struct {
	// Identifier of the webhook connection.
	WebhookId string `json:"webhookId"`
	// A JSON object in the format required by the target WebHook URL. For details on variables that can be used as parameters within your JSON object, please refer to Sumo Logic Doc Hub.
	Payload *string `json:"payload,omitempty"`
	// If this field is set to true, one webhook per result will be sent when the trigger conditions are met
	ItemizeAlerts *bool `json:"itemizeAlerts,omitempty"`
	// The maximum number of results for which we send separate alerts. This value should be between 1 and 100.
	MaxItemizedAlerts *int32 `json:"maxItemizedAlerts,omitempty"`
}

// NewWebhookSearchNotificationSyncDefinitionAllOf instantiates a new WebhookSearchNotificationSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSearchNotificationSyncDefinitionAllOf(webhookId string) *WebhookSearchNotificationSyncDefinitionAllOf {
	this := WebhookSearchNotificationSyncDefinitionAllOf{}
	this.WebhookId = webhookId
	var itemizeAlerts bool = false
	this.ItemizeAlerts = &itemizeAlerts
	return &this
}

// NewWebhookSearchNotificationSyncDefinitionAllOfWithDefaults instantiates a new WebhookSearchNotificationSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSearchNotificationSyncDefinitionAllOfWithDefaults() *WebhookSearchNotificationSyncDefinitionAllOf {
	this := WebhookSearchNotificationSyncDefinitionAllOf{}
	var itemizeAlerts bool = false
	this.ItemizeAlerts = &itemizeAlerts
	return &this
}

// GetWebhookId returns the WebhookId field value
func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetWebhookId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value
// and a boolean to check if the value has been set.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetWebhookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookId, true
}

// SetWebhookId sets field value
func (o *WebhookSearchNotificationSyncDefinitionAllOf) SetWebhookId(v string) {
	o.WebhookId = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) SetPayload(v string) {
	o.Payload = &v
}

// GetItemizeAlerts returns the ItemizeAlerts field value if set, zero value otherwise.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetItemizeAlerts() bool {
	if o == nil || IsNil(o.ItemizeAlerts) {
		var ret bool
		return ret
	}
	return *o.ItemizeAlerts
}

// GetItemizeAlertsOk returns a tuple with the ItemizeAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetItemizeAlertsOk() (*bool, bool) {
	if o == nil || IsNil(o.ItemizeAlerts) {
		return nil, false
	}
	return o.ItemizeAlerts, true
}

// HasItemizeAlerts returns a boolean if a field has been set.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) HasItemizeAlerts() bool {
	if o != nil && !IsNil(o.ItemizeAlerts) {
		return true
	}

	return false
}

// SetItemizeAlerts gets a reference to the given bool and assigns it to the ItemizeAlerts field.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) SetItemizeAlerts(v bool) {
	o.ItemizeAlerts = &v
}

// GetMaxItemizedAlerts returns the MaxItemizedAlerts field value if set, zero value otherwise.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetMaxItemizedAlerts() int32 {
	if o == nil || IsNil(o.MaxItemizedAlerts) {
		var ret int32
		return ret
	}
	return *o.MaxItemizedAlerts
}

// GetMaxItemizedAlertsOk returns a tuple with the MaxItemizedAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) GetMaxItemizedAlertsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxItemizedAlerts) {
		return nil, false
	}
	return o.MaxItemizedAlerts, true
}

// HasMaxItemizedAlerts returns a boolean if a field has been set.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) HasMaxItemizedAlerts() bool {
	if o != nil && !IsNil(o.MaxItemizedAlerts) {
		return true
	}

	return false
}

// SetMaxItemizedAlerts gets a reference to the given int32 and assigns it to the MaxItemizedAlerts field.
func (o *WebhookSearchNotificationSyncDefinitionAllOf) SetMaxItemizedAlerts(v int32) {
	o.MaxItemizedAlerts = &v
}

func (o WebhookSearchNotificationSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSearchNotificationSyncDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["webhookId"] = o.WebhookId
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.ItemizeAlerts) {
		toSerialize["itemizeAlerts"] = o.ItemizeAlerts
	}
	if !IsNil(o.MaxItemizedAlerts) {
		toSerialize["maxItemizedAlerts"] = o.MaxItemizedAlerts
	}
	return toSerialize, nil
}

type NullableWebhookSearchNotificationSyncDefinitionAllOf struct {
	value *WebhookSearchNotificationSyncDefinitionAllOf
	isSet bool
}

func (v NullableWebhookSearchNotificationSyncDefinitionAllOf) Get() *WebhookSearchNotificationSyncDefinitionAllOf {
	return v.value
}

func (v *NullableWebhookSearchNotificationSyncDefinitionAllOf) Set(val *WebhookSearchNotificationSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSearchNotificationSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSearchNotificationSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSearchNotificationSyncDefinitionAllOf(val *WebhookSearchNotificationSyncDefinitionAllOf) *NullableWebhookSearchNotificationSyncDefinitionAllOf {
	return &NullableWebhookSearchNotificationSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWebhookSearchNotificationSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSearchNotificationSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


