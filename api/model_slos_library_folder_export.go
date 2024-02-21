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

// checks if the SlosLibraryFolderExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlosLibraryFolderExport{}

// SlosLibraryFolderExport struct for SlosLibraryFolderExport
type SlosLibraryFolderExport struct {
	SlosLibraryBaseExport
	// The items in the folder. A multi-type list of types slo or folder.
	Children []SlosLibraryBaseExport `json:"children,omitempty"`
}

// NewSlosLibraryFolderExport instantiates a new SlosLibraryFolderExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibraryFolderExport(name string, type_ string) *SlosLibraryFolderExport {
	this := SlosLibraryFolderExport{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewSlosLibraryFolderExportWithDefaults instantiates a new SlosLibraryFolderExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibraryFolderExportWithDefaults() *SlosLibraryFolderExport {
	this := SlosLibraryFolderExport{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *SlosLibraryFolderExport) GetChildren() []SlosLibraryBaseExport {
	if o == nil || IsNil(o.Children) {
		var ret []SlosLibraryBaseExport
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibraryFolderExport) GetChildrenOk() ([]SlosLibraryBaseExport, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *SlosLibraryFolderExport) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []SlosLibraryBaseExport and assigns it to the Children field.
func (o *SlosLibraryFolderExport) SetChildren(v []SlosLibraryBaseExport) {
	o.Children = v
}

func (o SlosLibraryFolderExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlosLibraryFolderExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSlosLibraryBaseExport, errSlosLibraryBaseExport := json.Marshal(o.SlosLibraryBaseExport)
	if errSlosLibraryBaseExport != nil {
		return map[string]interface{}{}, errSlosLibraryBaseExport
	}
	errSlosLibraryBaseExport = json.Unmarshal([]byte(serializedSlosLibraryBaseExport), &toSerialize)
	if errSlosLibraryBaseExport != nil {
		return map[string]interface{}{}, errSlosLibraryBaseExport
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableSlosLibraryFolderExport struct {
	value *SlosLibraryFolderExport
	isSet bool
}

func (v NullableSlosLibraryFolderExport) Get() *SlosLibraryFolderExport {
	return v.value
}

func (v *NullableSlosLibraryFolderExport) Set(val *SlosLibraryFolderExport) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibraryFolderExport) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibraryFolderExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibraryFolderExport(val *SlosLibraryFolderExport) *NullableSlosLibraryFolderExport {
	return &NullableSlosLibraryFolderExport{value: val, isSet: true}
}

func (v NullableSlosLibraryFolderExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibraryFolderExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


