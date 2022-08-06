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

// AlertsLibraryFolderUpdate struct for AlertsLibraryFolderUpdate
type AlertsLibraryFolderUpdate struct {
	AlertsLibraryBaseUpdate
}

// NewAlertsLibraryFolderUpdate instantiates a new AlertsLibraryFolderUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryFolderUpdate(name string, version int64, type_ string) *AlertsLibraryFolderUpdate {
	this := AlertsLibraryFolderUpdate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Version = version
	this.Type = type_
	return &this
}

// NewAlertsLibraryFolderUpdateWithDefaults instantiates a new AlertsLibraryFolderUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryFolderUpdateWithDefaults() *AlertsLibraryFolderUpdate {
	this := AlertsLibraryFolderUpdate{}
	return &this
}

func (o AlertsLibraryFolderUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAlertsLibraryBaseUpdate, errAlertsLibraryBaseUpdate := json.Marshal(o.AlertsLibraryBaseUpdate)
	if errAlertsLibraryBaseUpdate != nil {
		return []byte{}, errAlertsLibraryBaseUpdate
	}
	errAlertsLibraryBaseUpdate = json.Unmarshal([]byte(serializedAlertsLibraryBaseUpdate), &toSerialize)
	if errAlertsLibraryBaseUpdate != nil {
		return []byte{}, errAlertsLibraryBaseUpdate
	}
	return json.Marshal(toSerialize)
}

type NullableAlertsLibraryFolderUpdate struct {
	value *AlertsLibraryFolderUpdate
	isSet bool
}

func (v NullableAlertsLibraryFolderUpdate) Get() *AlertsLibraryFolderUpdate {
	return v.value
}

func (v *NullableAlertsLibraryFolderUpdate) Set(val *AlertsLibraryFolderUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryFolderUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryFolderUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryFolderUpdate(val *AlertsLibraryFolderUpdate) *NullableAlertsLibraryFolderUpdate {
	return &NullableAlertsLibraryFolderUpdate{value: val, isSet: true}
}

func (v NullableAlertsLibraryFolderUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryFolderUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


