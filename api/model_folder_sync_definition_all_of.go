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

// FolderSyncDefinitionAllOf struct for FolderSyncDefinitionAllOf
type FolderSyncDefinitionAllOf struct {
	// An optional description for the folder.
	Description *string `json:"description,omitempty"`
	// The items in the folder, a list of Dashboard and/or Folder items.
	Children []ContentSyncDefinition `json:"children"`
}

// NewFolderSyncDefinitionAllOf instantiates a new FolderSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderSyncDefinitionAllOf(children []ContentSyncDefinition) *FolderSyncDefinitionAllOf {
	this := FolderSyncDefinitionAllOf{}
	this.Children = children
	return &this
}

// NewFolderSyncDefinitionAllOfWithDefaults instantiates a new FolderSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderSyncDefinitionAllOfWithDefaults() *FolderSyncDefinitionAllOf {
	this := FolderSyncDefinitionAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FolderSyncDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderSyncDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FolderSyncDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FolderSyncDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetChildren returns the Children field value
func (o *FolderSyncDefinitionAllOf) GetChildren() []ContentSyncDefinition {
	if o == nil {
		var ret []ContentSyncDefinition
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *FolderSyncDefinitionAllOf) GetChildrenOk() ([]ContentSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *FolderSyncDefinitionAllOf) SetChildren(v []ContentSyncDefinition) {
	o.Children = v
}

func (o FolderSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["children"] = o.Children
	}
	return json.Marshal(toSerialize)
}

type NullableFolderSyncDefinitionAllOf struct {
	value *FolderSyncDefinitionAllOf
	isSet bool
}

func (v NullableFolderSyncDefinitionAllOf) Get() *FolderSyncDefinitionAllOf {
	return v.value
}

func (v *NullableFolderSyncDefinitionAllOf) Set(val *FolderSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderSyncDefinitionAllOf(val *FolderSyncDefinitionAllOf) *NullableFolderSyncDefinitionAllOf {
	return &NullableFolderSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableFolderSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


