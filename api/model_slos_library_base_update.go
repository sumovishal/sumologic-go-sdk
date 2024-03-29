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

// checks if the SlosLibraryBaseUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlosLibraryBaseUpdate{}

// SlosLibraryBaseUpdate struct for SlosLibraryBaseUpdate
type SlosLibraryBaseUpdate struct {
	// The name of the slo or folder.
	Name string `json:"name"`
	// The description of the slo or folder.
	Description *string `json:"description,omitempty"`
	// The version of the slo or folder.
	Version int64 `json:"version"`
	// Type of the object model.
	Type string `json:"type"`
}

// NewSlosLibraryBaseUpdate instantiates a new SlosLibraryBaseUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibraryBaseUpdate(name string, version int64, type_ string) *SlosLibraryBaseUpdate {
	this := SlosLibraryBaseUpdate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Version = version
	this.Type = type_
	return &this
}

// NewSlosLibraryBaseUpdateWithDefaults instantiates a new SlosLibraryBaseUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibraryBaseUpdateWithDefaults() *SlosLibraryBaseUpdate {
	this := SlosLibraryBaseUpdate{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetName returns the Name field value
func (o *SlosLibraryBaseUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlosLibraryBaseUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SlosLibraryBaseUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SlosLibraryBaseUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SlosLibraryBaseUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value
func (o *SlosLibraryBaseUpdate) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseUpdate) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SlosLibraryBaseUpdate) SetVersion(v int64) {
	o.Version = v
}

// GetType returns the Type field value
func (o *SlosLibraryBaseUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlosLibraryBaseUpdate) SetType(v string) {
	o.Type = v
}

func (o SlosLibraryBaseUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlosLibraryBaseUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["version"] = o.Version
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSlosLibraryBaseUpdate struct {
	value *SlosLibraryBaseUpdate
	isSet bool
}

func (v NullableSlosLibraryBaseUpdate) Get() *SlosLibraryBaseUpdate {
	return v.value
}

func (v *NullableSlosLibraryBaseUpdate) Set(val *SlosLibraryBaseUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibraryBaseUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibraryBaseUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibraryBaseUpdate(val *SlosLibraryBaseUpdate) *NullableSlosLibraryBaseUpdate {
	return &NullableSlosLibraryBaseUpdate{value: val, isSet: true}
}

func (v NullableSlosLibraryBaseUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibraryBaseUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


