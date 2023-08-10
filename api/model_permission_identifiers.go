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

// checks if the PermissionIdentifiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionIdentifiers{}

// PermissionIdentifiers struct for PermissionIdentifiers
type PermissionIdentifiers struct {
	// List of permission identifiers.
	PermissionIdentifiers []PermissionIdentifier `json:"permissionIdentifiers"`
}

// NewPermissionIdentifiers instantiates a new PermissionIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionIdentifiers(permissionIdentifiers []PermissionIdentifier) *PermissionIdentifiers {
	this := PermissionIdentifiers{}
	this.PermissionIdentifiers = permissionIdentifiers
	return &this
}

// NewPermissionIdentifiersWithDefaults instantiates a new PermissionIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionIdentifiersWithDefaults() *PermissionIdentifiers {
	this := PermissionIdentifiers{}
	return &this
}

// GetPermissionIdentifiers returns the PermissionIdentifiers field value
func (o *PermissionIdentifiers) GetPermissionIdentifiers() []PermissionIdentifier {
	if o == nil {
		var ret []PermissionIdentifier
		return ret
	}

	return o.PermissionIdentifiers
}

// GetPermissionIdentifiersOk returns a tuple with the PermissionIdentifiers field value
// and a boolean to check if the value has been set.
func (o *PermissionIdentifiers) GetPermissionIdentifiersOk() ([]PermissionIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionIdentifiers, true
}

// SetPermissionIdentifiers sets field value
func (o *PermissionIdentifiers) SetPermissionIdentifiers(v []PermissionIdentifier) {
	o.PermissionIdentifiers = v
}

func (o PermissionIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionIdentifiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissionIdentifiers"] = o.PermissionIdentifiers
	return toSerialize, nil
}

type NullablePermissionIdentifiers struct {
	value *PermissionIdentifiers
	isSet bool
}

func (v NullablePermissionIdentifiers) Get() *PermissionIdentifiers {
	return v.value
}

func (v *NullablePermissionIdentifiers) Set(val *PermissionIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionIdentifiers(val *PermissionIdentifiers) *NullablePermissionIdentifiers {
	return &NullablePermissionIdentifiers{value: val, isSet: true}
}

func (v NullablePermissionIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


