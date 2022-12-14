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

// NoneAutoCompleteSyncDefinition struct for NoneAutoCompleteSyncDefinition
type NoneAutoCompleteSyncDefinition struct {
	LogSearchParameterAutoCompleteSyncDefinition
}

// NewNoneAutoCompleteSyncDefinition instantiates a new NoneAutoCompleteSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoneAutoCompleteSyncDefinition(autoCompleteType string) *NoneAutoCompleteSyncDefinition {
	this := NoneAutoCompleteSyncDefinition{}
	this.AutoCompleteType = autoCompleteType
	return &this
}

// NewNoneAutoCompleteSyncDefinitionWithDefaults instantiates a new NoneAutoCompleteSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoneAutoCompleteSyncDefinitionWithDefaults() *NoneAutoCompleteSyncDefinition {
	this := NoneAutoCompleteSyncDefinition{}
	return &this
}

func (o NoneAutoCompleteSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedLogSearchParameterAutoCompleteSyncDefinition, errLogSearchParameterAutoCompleteSyncDefinition := json.Marshal(o.LogSearchParameterAutoCompleteSyncDefinition)
	if errLogSearchParameterAutoCompleteSyncDefinition != nil {
		return []byte{}, errLogSearchParameterAutoCompleteSyncDefinition
	}
	errLogSearchParameterAutoCompleteSyncDefinition = json.Unmarshal([]byte(serializedLogSearchParameterAutoCompleteSyncDefinition), &toSerialize)
	if errLogSearchParameterAutoCompleteSyncDefinition != nil {
		return []byte{}, errLogSearchParameterAutoCompleteSyncDefinition
	}
	return json.Marshal(toSerialize)
}

type NullableNoneAutoCompleteSyncDefinition struct {
	value *NoneAutoCompleteSyncDefinition
	isSet bool
}

func (v NullableNoneAutoCompleteSyncDefinition) Get() *NoneAutoCompleteSyncDefinition {
	return v.value
}

func (v *NullableNoneAutoCompleteSyncDefinition) Set(val *NoneAutoCompleteSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableNoneAutoCompleteSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableNoneAutoCompleteSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoneAutoCompleteSyncDefinition(val *NoneAutoCompleteSyncDefinition) *NullableNoneAutoCompleteSyncDefinition {
	return &NullableNoneAutoCompleteSyncDefinition{value: val, isSet: true}
}

func (v NullableNoneAutoCompleteSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoneAutoCompleteSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


