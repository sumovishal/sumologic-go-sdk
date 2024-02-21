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

// checks if the MutingSchedulesLibraryBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutingSchedulesLibraryBase{}

// MutingSchedulesLibraryBase struct for MutingSchedulesLibraryBase
type MutingSchedulesLibraryBase struct {
	// Name of the mutingschedule or folder.
	Name string `json:"name"`
	// Description of the mutingschedule or folder.
	Description *string `json:"description,omitempty"`
	// Type of the object model. Valid values:   1) MutingSchedulesLibraryMutingschedule   2) MutingSchedulesLibraryFolder
	Type string `json:"type"`
}

// NewMutingSchedulesLibraryBase instantiates a new MutingSchedulesLibraryBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutingSchedulesLibraryBase(name string, type_ string) *MutingSchedulesLibraryBase {
	this := MutingSchedulesLibraryBase{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Type = type_
	return &this
}

// NewMutingSchedulesLibraryBaseWithDefaults instantiates a new MutingSchedulesLibraryBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutingSchedulesLibraryBaseWithDefaults() *MutingSchedulesLibraryBase {
	this := MutingSchedulesLibraryBase{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetName returns the Name field value
func (o *MutingSchedulesLibraryBase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryBase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MutingSchedulesLibraryBase) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MutingSchedulesLibraryBase) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryBase) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MutingSchedulesLibraryBase) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MutingSchedulesLibraryBase) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *MutingSchedulesLibraryBase) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryBase) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MutingSchedulesLibraryBase) SetType(v string) {
	o.Type = v
}

func (o MutingSchedulesLibraryBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutingSchedulesLibraryBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMutingSchedulesLibraryBase struct {
	value *MutingSchedulesLibraryBase
	isSet bool
}

func (v NullableMutingSchedulesLibraryBase) Get() *MutingSchedulesLibraryBase {
	return v.value
}

func (v *NullableMutingSchedulesLibraryBase) Set(val *MutingSchedulesLibraryBase) {
	v.value = val
	v.isSet = true
}

func (v NullableMutingSchedulesLibraryBase) IsSet() bool {
	return v.isSet
}

func (v *NullableMutingSchedulesLibraryBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutingSchedulesLibraryBase(val *MutingSchedulesLibraryBase) *NullableMutingSchedulesLibraryBase {
	return &NullableMutingSchedulesLibraryBase{value: val, isSet: true}
}

func (v NullableMutingSchedulesLibraryBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutingSchedulesLibraryBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


