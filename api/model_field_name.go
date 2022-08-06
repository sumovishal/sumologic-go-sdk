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

// FieldName struct for FieldName
type FieldName struct {
	// Field name.
	FieldName string `json:"fieldName"`
}

// NewFieldName instantiates a new FieldName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldName(fieldName string) *FieldName {
	this := FieldName{}
	this.FieldName = fieldName
	return &this
}

// NewFieldNameWithDefaults instantiates a new FieldName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldNameWithDefaults() *FieldName {
	this := FieldName{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *FieldName) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *FieldName) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *FieldName) SetFieldName(v string) {
	o.FieldName = v
}

func (o FieldName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fieldName"] = o.FieldName
	}
	return json.Marshal(toSerialize)
}

type NullableFieldName struct {
	value *FieldName
	isSet bool
}

func (v NullableFieldName) Get() *FieldName {
	return v.value
}

func (v *NullableFieldName) Set(val *FieldName) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldName) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldName(val *FieldName) *NullableFieldName {
	return &NullableFieldName{value: val, isSet: true}
}

func (v NullableFieldName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


