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

// OrgIdentity struct for OrgIdentity
type OrgIdentity struct {
	ResourceIdentity
}

// NewOrgIdentity instantiates a new OrgIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgIdentity(id string, type_ string) *OrgIdentity {
	this := OrgIdentity{}
	this.Id = id
	var name string = "Unknown"
	this.Name = &name
	this.Type = type_
	return &this
}

// NewOrgIdentityWithDefaults instantiates a new OrgIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgIdentityWithDefaults() *OrgIdentity {
	this := OrgIdentity{}
	return &this
}

func (o OrgIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedResourceIdentity, errResourceIdentity := json.Marshal(o.ResourceIdentity)
	if errResourceIdentity != nil {
		return []byte{}, errResourceIdentity
	}
	errResourceIdentity = json.Unmarshal([]byte(serializedResourceIdentity), &toSerialize)
	if errResourceIdentity != nil {
		return []byte{}, errResourceIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableOrgIdentity struct {
	value *OrgIdentity
	isSet bool
}

func (v NullableOrgIdentity) Get() *OrgIdentity {
	return v.value
}

func (v *NullableOrgIdentity) Set(val *OrgIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgIdentity(val *OrgIdentity) *NullableOrgIdentity {
	return &NullableOrgIdentity{value: val, isSet: true}
}

func (v NullableOrgIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


