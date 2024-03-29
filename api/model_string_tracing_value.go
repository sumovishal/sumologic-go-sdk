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

// checks if the StringTracingValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringTracingValue{}

// StringTracingValue struct for StringTracingValue
type StringTracingValue struct {
	TracingValue
	Value string `json:"value"`
}

// NewStringTracingValue instantiates a new StringTracingValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringTracingValue(value string, type_ string) *StringTracingValue {
	this := StringTracingValue{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewStringTracingValueWithDefaults instantiates a new StringTracingValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringTracingValueWithDefaults() *StringTracingValue {
	this := StringTracingValue{}
	return &this
}

// GetValue returns the Value field value
func (o *StringTracingValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StringTracingValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StringTracingValue) SetValue(v string) {
	o.Value = v
}

func (o StringTracingValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringTracingValue) ToMap() (map[string]interface{}, error) {
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

type NullableStringTracingValue struct {
	value *StringTracingValue
	isSet bool
}

func (v NullableStringTracingValue) Get() *StringTracingValue {
	return v.value
}

func (v *NullableStringTracingValue) Set(val *StringTracingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableStringTracingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableStringTracingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringTracingValue(val *StringTracingValue) *NullableStringTracingValue {
	return &NullableStringTracingValue{value: val, isSet: true}
}

func (v NullableStringTracingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringTracingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


