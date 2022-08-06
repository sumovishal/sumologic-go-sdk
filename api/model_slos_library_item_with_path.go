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

// SlosLibraryItemWithPath struct for SlosLibraryItemWithPath
type SlosLibraryItemWithPath struct {
	Item SlosLibraryBaseResponse `json:"item"`
	// Path of the slo or folder.
	Path string `json:"path"`
}

// NewSlosLibraryItemWithPath instantiates a new SlosLibraryItemWithPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibraryItemWithPath(item SlosLibraryBaseResponse, path string) *SlosLibraryItemWithPath {
	this := SlosLibraryItemWithPath{}
	this.Item = item
	this.Path = path
	return &this
}

// NewSlosLibraryItemWithPathWithDefaults instantiates a new SlosLibraryItemWithPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibraryItemWithPathWithDefaults() *SlosLibraryItemWithPath {
	this := SlosLibraryItemWithPath{}
	return &this
}

// GetItem returns the Item field value
func (o *SlosLibraryItemWithPath) GetItem() SlosLibraryBaseResponse {
	if o == nil {
		var ret SlosLibraryBaseResponse
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryItemWithPath) GetItemOk() (*SlosLibraryBaseResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *SlosLibraryItemWithPath) SetItem(v SlosLibraryBaseResponse) {
	o.Item = v
}

// GetPath returns the Path field value
func (o *SlosLibraryItemWithPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryItemWithPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *SlosLibraryItemWithPath) SetPath(v string) {
	o.Path = v
}

func (o SlosLibraryItemWithPath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableSlosLibraryItemWithPath struct {
	value *SlosLibraryItemWithPath
	isSet bool
}

func (v NullableSlosLibraryItemWithPath) Get() *SlosLibraryItemWithPath {
	return v.value
}

func (v *NullableSlosLibraryItemWithPath) Set(val *SlosLibraryItemWithPath) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibraryItemWithPath) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibraryItemWithPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibraryItemWithPath(val *SlosLibraryItemWithPath) *NullableSlosLibraryItemWithPath {
	return &NullableSlosLibraryItemWithPath{value: val, isSet: true}
}

func (v NullableSlosLibraryItemWithPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibraryItemWithPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


