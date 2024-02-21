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

// checks if the AlertsLibraryFolderResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsLibraryFolderResponseAllOf{}

// AlertsLibraryFolderResponseAllOf struct for AlertsLibraryFolderResponseAllOf
type AlertsLibraryFolderResponseAllOf struct {
	// Children of the folder. NOTE: Permissions field will not be filled (empty list) for children.
	Children []AlertsLibraryBaseResponse `json:"children"`
}

// NewAlertsLibraryFolderResponseAllOf instantiates a new AlertsLibraryFolderResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryFolderResponseAllOf(children []AlertsLibraryBaseResponse) *AlertsLibraryFolderResponseAllOf {
	this := AlertsLibraryFolderResponseAllOf{}
	this.Children = children
	return &this
}

// NewAlertsLibraryFolderResponseAllOfWithDefaults instantiates a new AlertsLibraryFolderResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryFolderResponseAllOfWithDefaults() *AlertsLibraryFolderResponseAllOf {
	this := AlertsLibraryFolderResponseAllOf{}
	return &this
}

// GetChildren returns the Children field value
func (o *AlertsLibraryFolderResponseAllOf) GetChildren() []AlertsLibraryBaseResponse {
	if o == nil {
		var ret []AlertsLibraryBaseResponse
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryFolderResponseAllOf) GetChildrenOk() ([]AlertsLibraryBaseResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *AlertsLibraryFolderResponseAllOf) SetChildren(v []AlertsLibraryBaseResponse) {
	o.Children = v
}

func (o AlertsLibraryFolderResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsLibraryFolderResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

type NullableAlertsLibraryFolderResponseAllOf struct {
	value *AlertsLibraryFolderResponseAllOf
	isSet bool
}

func (v NullableAlertsLibraryFolderResponseAllOf) Get() *AlertsLibraryFolderResponseAllOf {
	return v.value
}

func (v *NullableAlertsLibraryFolderResponseAllOf) Set(val *AlertsLibraryFolderResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryFolderResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryFolderResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryFolderResponseAllOf(val *AlertsLibraryFolderResponseAllOf) *NullableAlertsLibraryFolderResponseAllOf {
	return &NullableAlertsLibraryFolderResponseAllOf{value: val, isSet: true}
}

func (v NullableAlertsLibraryFolderResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryFolderResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


