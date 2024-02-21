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

// checks if the MutingSchedulesLibraryFolderExportAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutingSchedulesLibraryFolderExportAllOf{}

// MutingSchedulesLibraryFolderExportAllOf struct for MutingSchedulesLibraryFolderExportAllOf
type MutingSchedulesLibraryFolderExportAllOf struct {
	// The items in the folder. A multi-type list of types mutingschedule or folder.
	Children []MutingSchedulesLibraryBaseExport `json:"children,omitempty"`
}

// NewMutingSchedulesLibraryFolderExportAllOf instantiates a new MutingSchedulesLibraryFolderExportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutingSchedulesLibraryFolderExportAllOf() *MutingSchedulesLibraryFolderExportAllOf {
	this := MutingSchedulesLibraryFolderExportAllOf{}
	return &this
}

// NewMutingSchedulesLibraryFolderExportAllOfWithDefaults instantiates a new MutingSchedulesLibraryFolderExportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutingSchedulesLibraryFolderExportAllOfWithDefaults() *MutingSchedulesLibraryFolderExportAllOf {
	this := MutingSchedulesLibraryFolderExportAllOf{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *MutingSchedulesLibraryFolderExportAllOf) GetChildren() []MutingSchedulesLibraryBaseExport {
	if o == nil || IsNil(o.Children) {
		var ret []MutingSchedulesLibraryBaseExport
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryFolderExportAllOf) GetChildrenOk() ([]MutingSchedulesLibraryBaseExport, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *MutingSchedulesLibraryFolderExportAllOf) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []MutingSchedulesLibraryBaseExport and assigns it to the Children field.
func (o *MutingSchedulesLibraryFolderExportAllOf) SetChildren(v []MutingSchedulesLibraryBaseExport) {
	o.Children = v
}

func (o MutingSchedulesLibraryFolderExportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutingSchedulesLibraryFolderExportAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableMutingSchedulesLibraryFolderExportAllOf struct {
	value *MutingSchedulesLibraryFolderExportAllOf
	isSet bool
}

func (v NullableMutingSchedulesLibraryFolderExportAllOf) Get() *MutingSchedulesLibraryFolderExportAllOf {
	return v.value
}

func (v *NullableMutingSchedulesLibraryFolderExportAllOf) Set(val *MutingSchedulesLibraryFolderExportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMutingSchedulesLibraryFolderExportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMutingSchedulesLibraryFolderExportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutingSchedulesLibraryFolderExportAllOf(val *MutingSchedulesLibraryFolderExportAllOf) *NullableMutingSchedulesLibraryFolderExportAllOf {
	return &NullableMutingSchedulesLibraryFolderExportAllOf{value: val, isSet: true}
}

func (v NullableMutingSchedulesLibraryFolderExportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutingSchedulesLibraryFolderExportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


