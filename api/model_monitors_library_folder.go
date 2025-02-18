/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MonitorsLibraryFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorsLibraryFolder{}

// MonitorsLibraryFolder struct for MonitorsLibraryFolder
type MonitorsLibraryFolder struct {
	MonitorsLibraryBase
}

type _MonitorsLibraryFolder MonitorsLibraryFolder

// NewMonitorsLibraryFolder instantiates a new MonitorsLibraryFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorsLibraryFolder(name string, type_ string) *MonitorsLibraryFolder {
	this := MonitorsLibraryFolder{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Type = type_
	return &this
}

// NewMonitorsLibraryFolderWithDefaults instantiates a new MonitorsLibraryFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorsLibraryFolderWithDefaults() *MonitorsLibraryFolder {
	this := MonitorsLibraryFolder{}
	return &this
}

func (o MonitorsLibraryFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorsLibraryFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMonitorsLibraryBase, errMonitorsLibraryBase := json.Marshal(o.MonitorsLibraryBase)
	if errMonitorsLibraryBase != nil {
		return map[string]interface{}{}, errMonitorsLibraryBase
	}
	errMonitorsLibraryBase = json.Unmarshal([]byte(serializedMonitorsLibraryBase), &toSerialize)
	if errMonitorsLibraryBase != nil {
		return map[string]interface{}{}, errMonitorsLibraryBase
	}
	return toSerialize, nil
}

func (o *MonitorsLibraryFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMonitorsLibraryFolder := _MonitorsLibraryFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMonitorsLibraryFolder)

	if err != nil {
		return err
	}

	*o = MonitorsLibraryFolder(varMonitorsLibraryFolder)

	return err
}

type NullableMonitorsLibraryFolder struct {
	value *MonitorsLibraryFolder
	isSet bool
}

func (v NullableMonitorsLibraryFolder) Get() *MonitorsLibraryFolder {
	return v.value
}

func (v *NullableMonitorsLibraryFolder) Set(val *MonitorsLibraryFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorsLibraryFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorsLibraryFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorsLibraryFolder(val *MonitorsLibraryFolder) *NullableMonitorsLibraryFolder {
	return &NullableMonitorsLibraryFolder{value: val, isSet: true}
}

func (v NullableMonitorsLibraryFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorsLibraryFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


