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

// TextPanelAllOf A panel that has text.
type TextPanelAllOf struct {
	// Text to display in the panel.
	Text string `json:"text"`
}

// NewTextPanelAllOf instantiates a new TextPanelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextPanelAllOf(text string) *TextPanelAllOf {
	this := TextPanelAllOf{}
	this.Text = text
	return &this
}

// NewTextPanelAllOfWithDefaults instantiates a new TextPanelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextPanelAllOfWithDefaults() *TextPanelAllOf {
	this := TextPanelAllOf{}
	return &this
}

// GetText returns the Text field value
func (o *TextPanelAllOf) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TextPanelAllOf) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *TextPanelAllOf) SetText(v string) {
	o.Text = v
}

func (o TextPanelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableTextPanelAllOf struct {
	value *TextPanelAllOf
	isSet bool
}

func (v NullableTextPanelAllOf) Get() *TextPanelAllOf {
	return v.value
}

func (v *NullableTextPanelAllOf) Set(val *TextPanelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTextPanelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTextPanelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextPanelAllOf(val *TextPanelAllOf) *NullableTextPanelAllOf {
	return &NullableTextPanelAllOf{value: val, isSet: true}
}

func (v NullableTextPanelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextPanelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


