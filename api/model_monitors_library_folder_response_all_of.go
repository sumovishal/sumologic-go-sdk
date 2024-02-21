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

// checks if the MonitorsLibraryFolderResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorsLibraryFolderResponseAllOf{}

// MonitorsLibraryFolderResponseAllOf struct for MonitorsLibraryFolderResponseAllOf
type MonitorsLibraryFolderResponseAllOf struct {
	// Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint.
	Permissions []string `json:"permissions"`
	// Children of the folder. NOTE: Permissions field will not be filled (empty list) for children.
	Children []MonitorsLibraryBaseResponse `json:"children"`
}

// NewMonitorsLibraryFolderResponseAllOf instantiates a new MonitorsLibraryFolderResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorsLibraryFolderResponseAllOf(permissions []string, children []MonitorsLibraryBaseResponse) *MonitorsLibraryFolderResponseAllOf {
	this := MonitorsLibraryFolderResponseAllOf{}
	this.Permissions = permissions
	this.Children = children
	return &this
}

// NewMonitorsLibraryFolderResponseAllOfWithDefaults instantiates a new MonitorsLibraryFolderResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorsLibraryFolderResponseAllOfWithDefaults() *MonitorsLibraryFolderResponseAllOf {
	this := MonitorsLibraryFolderResponseAllOf{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *MonitorsLibraryFolderResponseAllOf) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryFolderResponseAllOf) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *MonitorsLibraryFolderResponseAllOf) SetPermissions(v []string) {
	o.Permissions = v
}

// GetChildren returns the Children field value
func (o *MonitorsLibraryFolderResponseAllOf) GetChildren() []MonitorsLibraryBaseResponse {
	if o == nil {
		var ret []MonitorsLibraryBaseResponse
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryFolderResponseAllOf) GetChildrenOk() ([]MonitorsLibraryBaseResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *MonitorsLibraryFolderResponseAllOf) SetChildren(v []MonitorsLibraryBaseResponse) {
	o.Children = v
}

func (o MonitorsLibraryFolderResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorsLibraryFolderResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

type NullableMonitorsLibraryFolderResponseAllOf struct {
	value *MonitorsLibraryFolderResponseAllOf
	isSet bool
}

func (v NullableMonitorsLibraryFolderResponseAllOf) Get() *MonitorsLibraryFolderResponseAllOf {
	return v.value
}

func (v *NullableMonitorsLibraryFolderResponseAllOf) Set(val *MonitorsLibraryFolderResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorsLibraryFolderResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorsLibraryFolderResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorsLibraryFolderResponseAllOf(val *MonitorsLibraryFolderResponseAllOf) *NullableMonitorsLibraryFolderResponseAllOf {
	return &NullableMonitorsLibraryFolderResponseAllOf{value: val, isSet: true}
}

func (v NullableMonitorsLibraryFolderResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorsLibraryFolderResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


