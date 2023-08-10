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

// checks if the MonitorTemplatesLibraryFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorTemplatesLibraryFolder{}

// MonitorTemplatesLibraryFolder struct for MonitorTemplatesLibraryFolder
type MonitorTemplatesLibraryFolder struct {
	MonitorTemplatesLibraryBase
}

// NewMonitorTemplatesLibraryFolder instantiates a new MonitorTemplatesLibraryFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorTemplatesLibraryFolder(name string, type_ string) *MonitorTemplatesLibraryFolder {
	this := MonitorTemplatesLibraryFolder{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Type = type_
	return &this
}

// NewMonitorTemplatesLibraryFolderWithDefaults instantiates a new MonitorTemplatesLibraryFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorTemplatesLibraryFolderWithDefaults() *MonitorTemplatesLibraryFolder {
	this := MonitorTemplatesLibraryFolder{}
	return &this
}

func (o MonitorTemplatesLibraryFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorTemplatesLibraryFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMonitorTemplatesLibraryBase, errMonitorTemplatesLibraryBase := json.Marshal(o.MonitorTemplatesLibraryBase)
	if errMonitorTemplatesLibraryBase != nil {
		return map[string]interface{}{}, errMonitorTemplatesLibraryBase
	}
	errMonitorTemplatesLibraryBase = json.Unmarshal([]byte(serializedMonitorTemplatesLibraryBase), &toSerialize)
	if errMonitorTemplatesLibraryBase != nil {
		return map[string]interface{}{}, errMonitorTemplatesLibraryBase
	}
	return toSerialize, nil
}

type NullableMonitorTemplatesLibraryFolder struct {
	value *MonitorTemplatesLibraryFolder
	isSet bool
}

func (v NullableMonitorTemplatesLibraryFolder) Get() *MonitorTemplatesLibraryFolder {
	return v.value
}

func (v *NullableMonitorTemplatesLibraryFolder) Set(val *MonitorTemplatesLibraryFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorTemplatesLibraryFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorTemplatesLibraryFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorTemplatesLibraryFolder(val *MonitorTemplatesLibraryFolder) *NullableMonitorTemplatesLibraryFolder {
	return &NullableMonitorTemplatesLibraryFolder{value: val, isSet: true}
}

func (v NullableMonitorTemplatesLibraryFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorTemplatesLibraryFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


