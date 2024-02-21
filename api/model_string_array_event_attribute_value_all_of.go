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

// checks if the StringArrayEventAttributeValueAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringArrayEventAttributeValueAllOf{}

// StringArrayEventAttributeValueAllOf struct for StringArrayEventAttributeValueAllOf
type StringArrayEventAttributeValueAllOf struct {
	Values []string `json:"values"`
}

// NewStringArrayEventAttributeValueAllOf instantiates a new StringArrayEventAttributeValueAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringArrayEventAttributeValueAllOf(values []string) *StringArrayEventAttributeValueAllOf {
	this := StringArrayEventAttributeValueAllOf{}
	this.Values = values
	return &this
}

// NewStringArrayEventAttributeValueAllOfWithDefaults instantiates a new StringArrayEventAttributeValueAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringArrayEventAttributeValueAllOfWithDefaults() *StringArrayEventAttributeValueAllOf {
	this := StringArrayEventAttributeValueAllOf{}
	return &this
}

// GetValues returns the Values field value
func (o *StringArrayEventAttributeValueAllOf) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *StringArrayEventAttributeValueAllOf) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *StringArrayEventAttributeValueAllOf) SetValues(v []string) {
	o.Values = v
}

func (o StringArrayEventAttributeValueAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringArrayEventAttributeValueAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableStringArrayEventAttributeValueAllOf struct {
	value *StringArrayEventAttributeValueAllOf
	isSet bool
}

func (v NullableStringArrayEventAttributeValueAllOf) Get() *StringArrayEventAttributeValueAllOf {
	return v.value
}

func (v *NullableStringArrayEventAttributeValueAllOf) Set(val *StringArrayEventAttributeValueAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStringArrayEventAttributeValueAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStringArrayEventAttributeValueAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringArrayEventAttributeValueAllOf(val *StringArrayEventAttributeValueAllOf) *NullableStringArrayEventAttributeValueAllOf {
	return &NullableStringArrayEventAttributeValueAllOf{value: val, isSet: true}
}

func (v NullableStringArrayEventAttributeValueAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringArrayEventAttributeValueAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


