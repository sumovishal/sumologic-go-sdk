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

// IntegerTracingValue struct for IntegerTracingValue
type IntegerTracingValue struct {
	TracingValue
	Value int64 `json:"value"`
}

// NewIntegerTracingValue instantiates a new IntegerTracingValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegerTracingValue(value int64, type_ string) *IntegerTracingValue {
	this := IntegerTracingValue{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewIntegerTracingValueWithDefaults instantiates a new IntegerTracingValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegerTracingValueWithDefaults() *IntegerTracingValue {
	this := IntegerTracingValue{}
	return &this
}

// GetValue returns the Value field value
func (o *IntegerTracingValue) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IntegerTracingValue) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *IntegerTracingValue) SetValue(v int64) {
	o.Value = v
}

func (o IntegerTracingValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTracingValue, errTracingValue := json.Marshal(o.TracingValue)
	if errTracingValue != nil {
		return []byte{}, errTracingValue
	}
	errTracingValue = json.Unmarshal([]byte(serializedTracingValue), &toSerialize)
	if errTracingValue != nil {
		return []byte{}, errTracingValue
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableIntegerTracingValue struct {
	value *IntegerTracingValue
	isSet bool
}

func (v NullableIntegerTracingValue) Get() *IntegerTracingValue {
	return v.value
}

func (v *NullableIntegerTracingValue) Set(val *IntegerTracingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegerTracingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegerTracingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegerTracingValue(val *IntegerTracingValue) *NullableIntegerTracingValue {
	return &NullableIntegerTracingValue{value: val, isSet: true}
}

func (v NullableIntegerTracingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegerTracingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


