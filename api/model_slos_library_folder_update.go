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

// checks if the SlosLibraryFolderUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlosLibraryFolderUpdate{}

// SlosLibraryFolderUpdate struct for SlosLibraryFolderUpdate
type SlosLibraryFolderUpdate struct {
	SlosLibraryBaseUpdate
}

// NewSlosLibraryFolderUpdate instantiates a new SlosLibraryFolderUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibraryFolderUpdate(name string, version int64, type_ string) *SlosLibraryFolderUpdate {
	this := SlosLibraryFolderUpdate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Version = version
	this.Type = type_
	return &this
}

// NewSlosLibraryFolderUpdateWithDefaults instantiates a new SlosLibraryFolderUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibraryFolderUpdateWithDefaults() *SlosLibraryFolderUpdate {
	this := SlosLibraryFolderUpdate{}
	return &this
}

func (o SlosLibraryFolderUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlosLibraryFolderUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSlosLibraryBaseUpdate, errSlosLibraryBaseUpdate := json.Marshal(o.SlosLibraryBaseUpdate)
	if errSlosLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errSlosLibraryBaseUpdate
	}
	errSlosLibraryBaseUpdate = json.Unmarshal([]byte(serializedSlosLibraryBaseUpdate), &toSerialize)
	if errSlosLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errSlosLibraryBaseUpdate
	}
	return toSerialize, nil
}

type NullableSlosLibraryFolderUpdate struct {
	value *SlosLibraryFolderUpdate
	isSet bool
}

func (v NullableSlosLibraryFolderUpdate) Get() *SlosLibraryFolderUpdate {
	return v.value
}

func (v *NullableSlosLibraryFolderUpdate) Set(val *SlosLibraryFolderUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibraryFolderUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibraryFolderUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibraryFolderUpdate(val *SlosLibraryFolderUpdate) *NullableSlosLibraryFolderUpdate {
	return &NullableSlosLibraryFolderUpdate{value: val, isSet: true}
}

func (v NullableSlosLibraryFolderUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibraryFolderUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


