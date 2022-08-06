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

// TracingValue struct for TracingValue
type TracingValue struct {
	// Type of the value model.
	Type string `json:"type"`
}

// NewTracingValue instantiates a new TracingValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTracingValue(type_ string) *TracingValue {
	this := TracingValue{}
	this.Type = type_
	return &this
}

// NewTracingValueWithDefaults instantiates a new TracingValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTracingValueWithDefaults() *TracingValue {
	this := TracingValue{}
	return &this
}

// GetType returns the Type field value
func (o *TracingValue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TracingValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TracingValue) SetType(v string) {
	o.Type = v
}

func (o TracingValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTracingValue struct {
	value *TracingValue
	isSet bool
}

func (v NullableTracingValue) Get() *TracingValue {
	return v.value
}

func (v *NullableTracingValue) Set(val *TracingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTracingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTracingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTracingValue(val *TracingValue) *NullableTracingValue {
	return &NullableTracingValue{value: val, isSet: true}
}

func (v NullableTracingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTracingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


