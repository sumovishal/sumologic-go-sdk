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

// checks if the MonitorsLibraryItemWithPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorsLibraryItemWithPath{}

// MonitorsLibraryItemWithPath struct for MonitorsLibraryItemWithPath
type MonitorsLibraryItemWithPath struct {
	Item MonitorsLibraryBaseResponse `json:"item"`
	// Path of the monitor or folder.
	Path string `json:"path"`
}

// NewMonitorsLibraryItemWithPath instantiates a new MonitorsLibraryItemWithPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorsLibraryItemWithPath(item MonitorsLibraryBaseResponse, path string) *MonitorsLibraryItemWithPath {
	this := MonitorsLibraryItemWithPath{}
	this.Item = item
	this.Path = path
	return &this
}

// NewMonitorsLibraryItemWithPathWithDefaults instantiates a new MonitorsLibraryItemWithPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorsLibraryItemWithPathWithDefaults() *MonitorsLibraryItemWithPath {
	this := MonitorsLibraryItemWithPath{}
	return &this
}

// GetItem returns the Item field value
func (o *MonitorsLibraryItemWithPath) GetItem() MonitorsLibraryBaseResponse {
	if o == nil {
		var ret MonitorsLibraryBaseResponse
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryItemWithPath) GetItemOk() (*MonitorsLibraryBaseResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *MonitorsLibraryItemWithPath) SetItem(v MonitorsLibraryBaseResponse) {
	o.Item = v
}

// GetPath returns the Path field value
func (o *MonitorsLibraryItemWithPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryItemWithPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *MonitorsLibraryItemWithPath) SetPath(v string) {
	o.Path = v
}

func (o MonitorsLibraryItemWithPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorsLibraryItemWithPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["item"] = o.Item
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

type NullableMonitorsLibraryItemWithPath struct {
	value *MonitorsLibraryItemWithPath
	isSet bool
}

func (v NullableMonitorsLibraryItemWithPath) Get() *MonitorsLibraryItemWithPath {
	return v.value
}

func (v *NullableMonitorsLibraryItemWithPath) Set(val *MonitorsLibraryItemWithPath) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorsLibraryItemWithPath) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorsLibraryItemWithPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorsLibraryItemWithPath(val *MonitorsLibraryItemWithPath) *NullableMonitorsLibraryItemWithPath {
	return &NullableMonitorsLibraryItemWithPath{value: val, isSet: true}
}

func (v NullableMonitorsLibraryItemWithPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorsLibraryItemWithPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


