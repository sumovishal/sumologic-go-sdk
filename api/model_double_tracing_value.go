/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// checks if the DoubleTracingValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DoubleTracingValue{}

// DoubleTracingValue struct for DoubleTracingValue
type DoubleTracingValue struct {
	TracingValue
	Value float64 `json:"value"`
}

// NewDoubleTracingValue instantiates a new DoubleTracingValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDoubleTracingValue(value float64, type_ string) *DoubleTracingValue {
	this := DoubleTracingValue{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewDoubleTracingValueWithDefaults instantiates a new DoubleTracingValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDoubleTracingValueWithDefaults() *DoubleTracingValue {
	this := DoubleTracingValue{}
	return &this
}

// GetValue returns the Value field value
func (o *DoubleTracingValue) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DoubleTracingValue) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DoubleTracingValue) SetValue(v float64) {
	o.Value = v
}

func (o DoubleTracingValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DoubleTracingValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTracingValue, errTracingValue := json.Marshal(o.TracingValue)
	if errTracingValue != nil {
		return map[string]interface{}{}, errTracingValue
	}
	errTracingValue = json.Unmarshal([]byte(serializedTracingValue), &toSerialize)
	if errTracingValue != nil {
		return map[string]interface{}{}, errTracingValue
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableDoubleTracingValue struct {
	value *DoubleTracingValue
	isSet bool
}

func (v NullableDoubleTracingValue) Get() *DoubleTracingValue {
	return v.value
}

func (v *NullableDoubleTracingValue) Set(val *DoubleTracingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDoubleTracingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDoubleTracingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDoubleTracingValue(val *DoubleTracingValue) *NullableDoubleTracingValue {
	return &NullableDoubleTracingValue{value: val, isSet: true}
}

func (v NullableDoubleTracingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDoubleTracingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


