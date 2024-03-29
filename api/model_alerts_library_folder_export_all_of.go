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

// checks if the AlertsLibraryFolderExportAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsLibraryFolderExportAllOf{}

// AlertsLibraryFolderExportAllOf struct for AlertsLibraryFolderExportAllOf
type AlertsLibraryFolderExportAllOf struct {
	// The items in the folder. A multi-type list of types alert or folder.
	Children []AlertsLibraryBaseExport `json:"children,omitempty"`
}

// NewAlertsLibraryFolderExportAllOf instantiates a new AlertsLibraryFolderExportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryFolderExportAllOf() *AlertsLibraryFolderExportAllOf {
	this := AlertsLibraryFolderExportAllOf{}
	return &this
}

// NewAlertsLibraryFolderExportAllOfWithDefaults instantiates a new AlertsLibraryFolderExportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryFolderExportAllOfWithDefaults() *AlertsLibraryFolderExportAllOf {
	this := AlertsLibraryFolderExportAllOf{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *AlertsLibraryFolderExportAllOf) GetChildren() []AlertsLibraryBaseExport {
	if o == nil || IsNil(o.Children) {
		var ret []AlertsLibraryBaseExport
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryFolderExportAllOf) GetChildrenOk() ([]AlertsLibraryBaseExport, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *AlertsLibraryFolderExportAllOf) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []AlertsLibraryBaseExport and assigns it to the Children field.
func (o *AlertsLibraryFolderExportAllOf) SetChildren(v []AlertsLibraryBaseExport) {
	o.Children = v
}

func (o AlertsLibraryFolderExportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsLibraryFolderExportAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableAlertsLibraryFolderExportAllOf struct {
	value *AlertsLibraryFolderExportAllOf
	isSet bool
}

func (v NullableAlertsLibraryFolderExportAllOf) Get() *AlertsLibraryFolderExportAllOf {
	return v.value
}

func (v *NullableAlertsLibraryFolderExportAllOf) Set(val *AlertsLibraryFolderExportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryFolderExportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryFolderExportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryFolderExportAllOf(val *AlertsLibraryFolderExportAllOf) *NullableAlertsLibraryFolderExportAllOf {
	return &NullableAlertsLibraryFolderExportAllOf{value: val, isSet: true}
}

func (v NullableAlertsLibraryFolderExportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryFolderExportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


