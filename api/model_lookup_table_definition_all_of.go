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

// checks if the LookupTableDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupTableDefinitionAllOf{}

// LookupTableDefinitionAllOf struct for LookupTableDefinitionAllOf
type LookupTableDefinitionAllOf struct {
	// The name of the lookup table.
	Name *string `json:"name,omitempty"`
	// The parent-folder-path identifier of the lookup table in the Library.
	ParentFolderId *string `json:"parentFolderId,omitempty"`
}

// NewLookupTableDefinitionAllOf instantiates a new LookupTableDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupTableDefinitionAllOf() *LookupTableDefinitionAllOf {
	this := LookupTableDefinitionAllOf{}
	return &this
}

// NewLookupTableDefinitionAllOfWithDefaults instantiates a new LookupTableDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupTableDefinitionAllOfWithDefaults() *LookupTableDefinitionAllOf {
	this := LookupTableDefinitionAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LookupTableDefinitionAllOf) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTableDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LookupTableDefinitionAllOf) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LookupTableDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetParentFolderId returns the ParentFolderId field value if set, zero value otherwise.
func (o *LookupTableDefinitionAllOf) GetParentFolderId() string {
	if o == nil || IsNil(o.ParentFolderId) {
		var ret string
		return ret
	}
	return *o.ParentFolderId
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTableDefinitionAllOf) GetParentFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentFolderId) {
		return nil, false
	}
	return o.ParentFolderId, true
}

// HasParentFolderId returns a boolean if a field has been set.
func (o *LookupTableDefinitionAllOf) HasParentFolderId() bool {
	if o != nil && !IsNil(o.ParentFolderId) {
		return true
	}

	return false
}

// SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.
func (o *LookupTableDefinitionAllOf) SetParentFolderId(v string) {
	o.ParentFolderId = &v
}

func (o LookupTableDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupTableDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ParentFolderId) {
		toSerialize["parentFolderId"] = o.ParentFolderId
	}
	return toSerialize, nil
}

type NullableLookupTableDefinitionAllOf struct {
	value *LookupTableDefinitionAllOf
	isSet bool
}

func (v NullableLookupTableDefinitionAllOf) Get() *LookupTableDefinitionAllOf {
	return v.value
}

func (v *NullableLookupTableDefinitionAllOf) Set(val *LookupTableDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupTableDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupTableDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupTableDefinitionAllOf(val *LookupTableDefinitionAllOf) *NullableLookupTableDefinitionAllOf {
	return &NullableLookupTableDefinitionAllOf{value: val, isSet: true}
}

func (v NullableLookupTableDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupTableDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


