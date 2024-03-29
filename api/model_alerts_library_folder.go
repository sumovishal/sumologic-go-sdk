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

// checks if the AlertsLibraryFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsLibraryFolder{}

// AlertsLibraryFolder struct for AlertsLibraryFolder
type AlertsLibraryFolder struct {
	AlertsLibraryBase
}

// NewAlertsLibraryFolder instantiates a new AlertsLibraryFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryFolder(name string, type_ string) *AlertsLibraryFolder {
	this := AlertsLibraryFolder{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Type = type_
	var isLocked bool = false
	this.IsLocked = &isLocked
	return &this
}

// NewAlertsLibraryFolderWithDefaults instantiates a new AlertsLibraryFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryFolderWithDefaults() *AlertsLibraryFolder {
	this := AlertsLibraryFolder{}
	return &this
}

func (o AlertsLibraryFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsLibraryFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAlertsLibraryBase, errAlertsLibraryBase := json.Marshal(o.AlertsLibraryBase)
	if errAlertsLibraryBase != nil {
		return map[string]interface{}{}, errAlertsLibraryBase
	}
	errAlertsLibraryBase = json.Unmarshal([]byte(serializedAlertsLibraryBase), &toSerialize)
	if errAlertsLibraryBase != nil {
		return map[string]interface{}{}, errAlertsLibraryBase
	}
	return toSerialize, nil
}

type NullableAlertsLibraryFolder struct {
	value *AlertsLibraryFolder
	isSet bool
}

func (v NullableAlertsLibraryFolder) Get() *AlertsLibraryFolder {
	return v.value
}

func (v *NullableAlertsLibraryFolder) Set(val *AlertsLibraryFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryFolder(val *AlertsLibraryFolder) *NullableAlertsLibraryFolder {
	return &NullableAlertsLibraryFolder{value: val, isSet: true}
}

func (v NullableAlertsLibraryFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


