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

// TextEntriesAutoCompleteSyncDefinitionAllOf struct for TextEntriesAutoCompleteSyncDefinitionAllOf
type TextEntriesAutoCompleteSyncDefinitionAllOf struct {
	// The autocomplete key to be used to fetch autocomplete values.
	AutoCompleteKey string `json:"autoCompleteKey"`
}

// NewTextEntriesAutoCompleteSyncDefinitionAllOf instantiates a new TextEntriesAutoCompleteSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextEntriesAutoCompleteSyncDefinitionAllOf(autoCompleteKey string) *TextEntriesAutoCompleteSyncDefinitionAllOf {
	this := TextEntriesAutoCompleteSyncDefinitionAllOf{}
	this.AutoCompleteKey = autoCompleteKey
	return &this
}

// NewTextEntriesAutoCompleteSyncDefinitionAllOfWithDefaults instantiates a new TextEntriesAutoCompleteSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextEntriesAutoCompleteSyncDefinitionAllOfWithDefaults() *TextEntriesAutoCompleteSyncDefinitionAllOf {
	this := TextEntriesAutoCompleteSyncDefinitionAllOf{}
	return &this
}

// GetAutoCompleteKey returns the AutoCompleteKey field value
func (o *TextEntriesAutoCompleteSyncDefinitionAllOf) GetAutoCompleteKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoCompleteKey
}

// GetAutoCompleteKeyOk returns a tuple with the AutoCompleteKey field value
// and a boolean to check if the value has been set.
func (o *TextEntriesAutoCompleteSyncDefinitionAllOf) GetAutoCompleteKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCompleteKey, true
}

// SetAutoCompleteKey sets field value
func (o *TextEntriesAutoCompleteSyncDefinitionAllOf) SetAutoCompleteKey(v string) {
	o.AutoCompleteKey = v
}

func (o TextEntriesAutoCompleteSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["autoCompleteKey"] = o.AutoCompleteKey
	}
	return json.Marshal(toSerialize)
}

type NullableTextEntriesAutoCompleteSyncDefinitionAllOf struct {
	value *TextEntriesAutoCompleteSyncDefinitionAllOf
	isSet bool
}

func (v NullableTextEntriesAutoCompleteSyncDefinitionAllOf) Get() *TextEntriesAutoCompleteSyncDefinitionAllOf {
	return v.value
}

func (v *NullableTextEntriesAutoCompleteSyncDefinitionAllOf) Set(val *TextEntriesAutoCompleteSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTextEntriesAutoCompleteSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTextEntriesAutoCompleteSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextEntriesAutoCompleteSyncDefinitionAllOf(val *TextEntriesAutoCompleteSyncDefinitionAllOf) *NullableTextEntriesAutoCompleteSyncDefinitionAllOf {
	return &NullableTextEntriesAutoCompleteSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableTextEntriesAutoCompleteSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextEntriesAutoCompleteSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


