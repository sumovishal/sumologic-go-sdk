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

// SpansFilterStandaloneKey struct for SpansFilterStandaloneKey
type SpansFilterStandaloneKey struct {
	SpansFilter
}

// NewSpansFilterStandaloneKey instantiates a new SpansFilterStandaloneKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpansFilterStandaloneKey(type_ string, fieldName string) *SpansFilterStandaloneKey {
	this := SpansFilterStandaloneKey{}
	this.Type = type_
	this.FieldName = fieldName
	return &this
}

// NewSpansFilterStandaloneKeyWithDefaults instantiates a new SpansFilterStandaloneKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpansFilterStandaloneKeyWithDefaults() *SpansFilterStandaloneKey {
	this := SpansFilterStandaloneKey{}
	return &this
}

func (o SpansFilterStandaloneKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSpansFilter, errSpansFilter := json.Marshal(o.SpansFilter)
	if errSpansFilter != nil {
		return []byte{}, errSpansFilter
	}
	errSpansFilter = json.Unmarshal([]byte(serializedSpansFilter), &toSerialize)
	if errSpansFilter != nil {
		return []byte{}, errSpansFilter
	}
	return json.Marshal(toSerialize)
}

type NullableSpansFilterStandaloneKey struct {
	value *SpansFilterStandaloneKey
	isSet bool
}

func (v NullableSpansFilterStandaloneKey) Get() *SpansFilterStandaloneKey {
	return v.value
}

func (v *NullableSpansFilterStandaloneKey) Set(val *SpansFilterStandaloneKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSpansFilterStandaloneKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSpansFilterStandaloneKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpansFilterStandaloneKey(val *SpansFilterStandaloneKey) *NullableSpansFilterStandaloneKey {
	return &NullableSpansFilterStandaloneKey{value: val, isSet: true}
}

func (v NullableSpansFilterStandaloneKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpansFilterStandaloneKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


