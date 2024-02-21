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

// checks if the PermissionStatements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionStatements{}

// PermissionStatements struct for PermissionStatements
type PermissionStatements struct {
	// A list of permission statements.
	PermissionStatements []PermissionStatement `json:"permissionStatements"`
}

// NewPermissionStatements instantiates a new PermissionStatements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionStatements(permissionStatements []PermissionStatement) *PermissionStatements {
	this := PermissionStatements{}
	this.PermissionStatements = permissionStatements
	return &this
}

// NewPermissionStatementsWithDefaults instantiates a new PermissionStatements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionStatementsWithDefaults() *PermissionStatements {
	this := PermissionStatements{}
	return &this
}

// GetPermissionStatements returns the PermissionStatements field value
func (o *PermissionStatements) GetPermissionStatements() []PermissionStatement {
	if o == nil {
		var ret []PermissionStatement
		return ret
	}

	return o.PermissionStatements
}

// GetPermissionStatementsOk returns a tuple with the PermissionStatements field value
// and a boolean to check if the value has been set.
func (o *PermissionStatements) GetPermissionStatementsOk() ([]PermissionStatement, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionStatements, true
}

// SetPermissionStatements sets field value
func (o *PermissionStatements) SetPermissionStatements(v []PermissionStatement) {
	o.PermissionStatements = v
}

func (o PermissionStatements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionStatements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissionStatements"] = o.PermissionStatements
	return toSerialize, nil
}

type NullablePermissionStatements struct {
	value *PermissionStatements
	isSet bool
}

func (v NullablePermissionStatements) Get() *PermissionStatements {
	return v.value
}

func (v *NullablePermissionStatements) Set(val *PermissionStatements) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionStatements) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionStatements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionStatements(val *PermissionStatements) *NullablePermissionStatements {
	return &NullablePermissionStatements{value: val, isSet: true}
}

func (v NullablePermissionStatements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionStatements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


