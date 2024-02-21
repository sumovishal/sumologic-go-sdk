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

// checks if the SlosLibraryBaseExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlosLibraryBaseExport{}

// SlosLibraryBaseExport struct for SlosLibraryBaseExport
type SlosLibraryBaseExport struct {
	// Name of the slo or folder.
	Name string `json:"name"`
	// Description of the slo or folder.
	Description *string `json:"description,omitempty"`
	// Type of the object model.
	Type string `json:"type"`
}

// NewSlosLibraryBaseExport instantiates a new SlosLibraryBaseExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibraryBaseExport(name string, type_ string) *SlosLibraryBaseExport {
	this := SlosLibraryBaseExport{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewSlosLibraryBaseExportWithDefaults instantiates a new SlosLibraryBaseExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibraryBaseExportWithDefaults() *SlosLibraryBaseExport {
	this := SlosLibraryBaseExport{}
	return &this
}

// GetName returns the Name field value
func (o *SlosLibraryBaseExport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseExport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlosLibraryBaseExport) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SlosLibraryBaseExport) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseExport) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SlosLibraryBaseExport) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SlosLibraryBaseExport) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *SlosLibraryBaseExport) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseExport) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlosLibraryBaseExport) SetType(v string) {
	o.Type = v
}

func (o SlosLibraryBaseExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlosLibraryBaseExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSlosLibraryBaseExport struct {
	value *SlosLibraryBaseExport
	isSet bool
}

func (v NullableSlosLibraryBaseExport) Get() *SlosLibraryBaseExport {
	return v.value
}

func (v *NullableSlosLibraryBaseExport) Set(val *SlosLibraryBaseExport) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibraryBaseExport) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibraryBaseExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibraryBaseExport(val *SlosLibraryBaseExport) *NullableSlosLibraryBaseExport {
	return &NullableSlosLibraryBaseExport{value: val, isSet: true}
}

func (v NullableSlosLibraryBaseExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibraryBaseExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


