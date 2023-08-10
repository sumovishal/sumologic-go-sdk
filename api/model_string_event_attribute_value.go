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

// checks if the StringEventAttributeValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringEventAttributeValue{}

// StringEventAttributeValue struct for StringEventAttributeValue
type StringEventAttributeValue struct {
	EventAttributeValue
	Value string `json:"value"`
}

// NewStringEventAttributeValue instantiates a new StringEventAttributeValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringEventAttributeValue(value string, type_ string) *StringEventAttributeValue {
	this := StringEventAttributeValue{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewStringEventAttributeValueWithDefaults instantiates a new StringEventAttributeValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringEventAttributeValueWithDefaults() *StringEventAttributeValue {
	this := StringEventAttributeValue{}
	return &this
}

// GetValue returns the Value field value
func (o *StringEventAttributeValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StringEventAttributeValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StringEventAttributeValue) SetValue(v string) {
	o.Value = v
}

func (o StringEventAttributeValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringEventAttributeValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEventAttributeValue, errEventAttributeValue := json.Marshal(o.EventAttributeValue)
	if errEventAttributeValue != nil {
		return map[string]interface{}{}, errEventAttributeValue
	}
	errEventAttributeValue = json.Unmarshal([]byte(serializedEventAttributeValue), &toSerialize)
	if errEventAttributeValue != nil {
		return map[string]interface{}{}, errEventAttributeValue
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableStringEventAttributeValue struct {
	value *StringEventAttributeValue
	isSet bool
}

func (v NullableStringEventAttributeValue) Get() *StringEventAttributeValue {
	return v.value
}

func (v *NullableStringEventAttributeValue) Set(val *StringEventAttributeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableStringEventAttributeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableStringEventAttributeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringEventAttributeValue(val *StringEventAttributeValue) *NullableStringEventAttributeValue {
	return &NullableStringEventAttributeValue{value: val, isSet: true}
}

func (v NullableStringEventAttributeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringEventAttributeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


