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

// checks if the MonitorTemplatesLibraryItemWithPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorTemplatesLibraryItemWithPath{}

// MonitorTemplatesLibraryItemWithPath struct for MonitorTemplatesLibraryItemWithPath
type MonitorTemplatesLibraryItemWithPath struct {
	Item MonitorTemplatesLibraryBaseResponse `json:"item"`
	// Path of the monitortemplate or folder.
	Path string `json:"path"`
}

// NewMonitorTemplatesLibraryItemWithPath instantiates a new MonitorTemplatesLibraryItemWithPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorTemplatesLibraryItemWithPath(item MonitorTemplatesLibraryBaseResponse, path string) *MonitorTemplatesLibraryItemWithPath {
	this := MonitorTemplatesLibraryItemWithPath{}
	this.Item = item
	this.Path = path
	return &this
}

// NewMonitorTemplatesLibraryItemWithPathWithDefaults instantiates a new MonitorTemplatesLibraryItemWithPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorTemplatesLibraryItemWithPathWithDefaults() *MonitorTemplatesLibraryItemWithPath {
	this := MonitorTemplatesLibraryItemWithPath{}
	return &this
}

// GetItem returns the Item field value
func (o *MonitorTemplatesLibraryItemWithPath) GetItem() MonitorTemplatesLibraryBaseResponse {
	if o == nil {
		var ret MonitorTemplatesLibraryBaseResponse
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryItemWithPath) GetItemOk() (*MonitorTemplatesLibraryBaseResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *MonitorTemplatesLibraryItemWithPath) SetItem(v MonitorTemplatesLibraryBaseResponse) {
	o.Item = v
}

// GetPath returns the Path field value
func (o *MonitorTemplatesLibraryItemWithPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryItemWithPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *MonitorTemplatesLibraryItemWithPath) SetPath(v string) {
	o.Path = v
}

func (o MonitorTemplatesLibraryItemWithPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorTemplatesLibraryItemWithPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["item"] = o.Item
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

type NullableMonitorTemplatesLibraryItemWithPath struct {
	value *MonitorTemplatesLibraryItemWithPath
	isSet bool
}

func (v NullableMonitorTemplatesLibraryItemWithPath) Get() *MonitorTemplatesLibraryItemWithPath {
	return v.value
}

func (v *NullableMonitorTemplatesLibraryItemWithPath) Set(val *MonitorTemplatesLibraryItemWithPath) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorTemplatesLibraryItemWithPath) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorTemplatesLibraryItemWithPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorTemplatesLibraryItemWithPath(val *MonitorTemplatesLibraryItemWithPath) *NullableMonitorTemplatesLibraryItemWithPath {
	return &NullableMonitorTemplatesLibraryItemWithPath{value: val, isSet: true}
}

func (v NullableMonitorTemplatesLibraryItemWithPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorTemplatesLibraryItemWithPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


