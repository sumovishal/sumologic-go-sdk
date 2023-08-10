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

// checks if the ArrayTracingValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArrayTracingValue{}

// ArrayTracingValue struct for ArrayTracingValue
type ArrayTracingValue struct {
	TracingValue
	Values []TracingValue `json:"values"`
}

// NewArrayTracingValue instantiates a new ArrayTracingValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArrayTracingValue(values []TracingValue, type_ string) *ArrayTracingValue {
	this := ArrayTracingValue{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewArrayTracingValueWithDefaults instantiates a new ArrayTracingValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArrayTracingValueWithDefaults() *ArrayTracingValue {
	this := ArrayTracingValue{}
	return &this
}

// GetValues returns the Values field value
func (o *ArrayTracingValue) GetValues() []TracingValue {
	if o == nil {
		var ret []TracingValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ArrayTracingValue) GetValuesOk() ([]TracingValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ArrayTracingValue) SetValues(v []TracingValue) {
	o.Values = v
}

func (o ArrayTracingValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArrayTracingValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTracingValue, errTracingValue := json.Marshal(o.TracingValue)
	if errTracingValue != nil {
		return map[string]interface{}{}, errTracingValue
	}
	errTracingValue = json.Unmarshal([]byte(serializedTracingValue), &toSerialize)
	if errTracingValue != nil {
		return map[string]interface{}{}, errTracingValue
	}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableArrayTracingValue struct {
	value *ArrayTracingValue
	isSet bool
}

func (v NullableArrayTracingValue) Get() *ArrayTracingValue {
	return v.value
}

func (v *NullableArrayTracingValue) Set(val *ArrayTracingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableArrayTracingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableArrayTracingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArrayTracingValue(val *ArrayTracingValue) *NullableArrayTracingValue {
	return &NullableArrayTracingValue{value: val, isSet: true}
}

func (v NullableArrayTracingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArrayTracingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


