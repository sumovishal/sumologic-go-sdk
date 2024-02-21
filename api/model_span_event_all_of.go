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

// checks if the SpanEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanEventAllOf{}

// SpanEventAllOf struct for SpanEventAllOf
type SpanEventAllOf struct {
	// Span event attributes.
	Attributes []SpanEventAttribute `json:"attributes,omitempty"`
}

// NewSpanEventAllOf instantiates a new SpanEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanEventAllOf() *SpanEventAllOf {
	this := SpanEventAllOf{}
	return &this
}

// NewSpanEventAllOfWithDefaults instantiates a new SpanEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanEventAllOfWithDefaults() *SpanEventAllOf {
	this := SpanEventAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SpanEventAllOf) GetAttributes() []SpanEventAttribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []SpanEventAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanEventAllOf) GetAttributesOk() ([]SpanEventAttribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SpanEventAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []SpanEventAttribute and assigns it to the Attributes field.
func (o *SpanEventAllOf) SetAttributes(v []SpanEventAttribute) {
	o.Attributes = v
}

func (o SpanEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableSpanEventAllOf struct {
	value *SpanEventAllOf
	isSet bool
}

func (v NullableSpanEventAllOf) Get() *SpanEventAllOf {
	return v.value
}

func (v *NullableSpanEventAllOf) Set(val *SpanEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanEventAllOf(val *SpanEventAllOf) *NullableSpanEventAllOf {
	return &NullableSpanEventAllOf{value: val, isSet: true}
}

func (v NullableSpanEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


