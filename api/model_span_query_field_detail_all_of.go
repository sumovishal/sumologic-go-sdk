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

// SpanQueryFieldDetailAllOf struct for SpanQueryFieldDetailAllOf
type SpanQueryFieldDetailAllOf struct {
	// Indicates whether the field is available in the schema.
	InSchema bool `json:"inSchema"`
}

// NewSpanQueryFieldDetailAllOf instantiates a new SpanQueryFieldDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanQueryFieldDetailAllOf(inSchema bool) *SpanQueryFieldDetailAllOf {
	this := SpanQueryFieldDetailAllOf{}
	this.InSchema = inSchema
	return &this
}

// NewSpanQueryFieldDetailAllOfWithDefaults instantiates a new SpanQueryFieldDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanQueryFieldDetailAllOfWithDefaults() *SpanQueryFieldDetailAllOf {
	this := SpanQueryFieldDetailAllOf{}
	return &this
}

// GetInSchema returns the InSchema field value
func (o *SpanQueryFieldDetailAllOf) GetInSchema() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InSchema
}

// GetInSchemaOk returns a tuple with the InSchema field value
// and a boolean to check if the value has been set.
func (o *SpanQueryFieldDetailAllOf) GetInSchemaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InSchema, true
}

// SetInSchema sets field value
func (o *SpanQueryFieldDetailAllOf) SetInSchema(v bool) {
	o.InSchema = v
}

func (o SpanQueryFieldDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inSchema"] = o.InSchema
	}
	return json.Marshal(toSerialize)
}

type NullableSpanQueryFieldDetailAllOf struct {
	value *SpanQueryFieldDetailAllOf
	isSet bool
}

func (v NullableSpanQueryFieldDetailAllOf) Get() *SpanQueryFieldDetailAllOf {
	return v.value
}

func (v *NullableSpanQueryFieldDetailAllOf) Set(val *SpanQueryFieldDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanQueryFieldDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanQueryFieldDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanQueryFieldDetailAllOf(val *SpanQueryFieldDetailAllOf) *NullableSpanQueryFieldDetailAllOf {
	return &NullableSpanQueryFieldDetailAllOf{value: val, isSet: true}
}

func (v NullableSpanQueryFieldDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanQueryFieldDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


