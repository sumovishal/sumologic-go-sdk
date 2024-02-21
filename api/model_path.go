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

// checks if the Path type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Path{}

// Path struct for Path
type Path struct {
	// Elements of the path.
	PathItems []PathItem `json:"pathItems"`
	// String representation of the path.
	Path string `json:"path"`
}

// NewPath instantiates a new Path object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPath(pathItems []PathItem, path string) *Path {
	this := Path{}
	this.PathItems = pathItems
	this.Path = path
	return &this
}

// NewPathWithDefaults instantiates a new Path object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathWithDefaults() *Path {
	this := Path{}
	return &this
}

// GetPathItems returns the PathItems field value
func (o *Path) GetPathItems() []PathItem {
	if o == nil {
		var ret []PathItem
		return ret
	}

	return o.PathItems
}

// GetPathItemsOk returns a tuple with the PathItems field value
// and a boolean to check if the value has been set.
func (o *Path) GetPathItemsOk() ([]PathItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.PathItems, true
}

// SetPathItems sets field value
func (o *Path) SetPathItems(v []PathItem) {
	o.PathItems = v
}

// GetPath returns the Path field value
func (o *Path) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Path) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *Path) SetPath(v string) {
	o.Path = v
}

func (o Path) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Path) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pathItems"] = o.PathItems
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

type NullablePath struct {
	value *Path
	isSet bool
}

func (v NullablePath) Get() *Path {
	return v.value
}

func (v *NullablePath) Set(val *Path) {
	v.value = val
	v.isSet = true
}

func (v NullablePath) IsSet() bool {
	return v.isSet
}

func (v *NullablePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePath(val *Path) *NullablePath {
	return &NullablePath{value: val, isSet: true}
}

func (v NullablePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


