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

// checks if the WarningDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarningDescription{}

// WarningDescription Warning description
type WarningDescription struct {
	// Description of the warning.
	Message string `json:"message"`
	// An optional cause of this warning.
	Cause *string `json:"cause,omitempty"`
}

// NewWarningDescription instantiates a new WarningDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarningDescription(message string) *WarningDescription {
	this := WarningDescription{}
	this.Message = message
	return &this
}

// NewWarningDescriptionWithDefaults instantiates a new WarningDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningDescriptionWithDefaults() *WarningDescription {
	this := WarningDescription{}
	return &this
}

// GetMessage returns the Message field value
func (o *WarningDescription) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *WarningDescription) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *WarningDescription) SetMessage(v string) {
	o.Message = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *WarningDescription) GetCause() string {
	if o == nil || IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningDescription) GetCauseOk() (*string, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *WarningDescription) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *WarningDescription) SetCause(v string) {
	o.Cause = &v
}

func (o WarningDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarningDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	return toSerialize, nil
}

type NullableWarningDescription struct {
	value *WarningDescription
	isSet bool
}

func (v NullableWarningDescription) Get() *WarningDescription {
	return v.value
}

func (v *NullableWarningDescription) Set(val *WarningDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableWarningDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableWarningDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarningDescription(val *WarningDescription) *NullableWarningDescription {
	return &NullableWarningDescription{value: val, isSet: true}
}

func (v NullableWarningDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarningDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


