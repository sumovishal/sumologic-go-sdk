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

// checks if the ChangeEmailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeEmailRequest{}

// ChangeEmailRequest struct for ChangeEmailRequest
type ChangeEmailRequest struct {
	// New email address of the user.
	Email string `json:"email"`
}

// NewChangeEmailRequest instantiates a new ChangeEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeEmailRequest(email string) *ChangeEmailRequest {
	this := ChangeEmailRequest{}
	this.Email = email
	return &this
}

// NewChangeEmailRequestWithDefaults instantiates a new ChangeEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeEmailRequestWithDefaults() *ChangeEmailRequest {
	this := ChangeEmailRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *ChangeEmailRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ChangeEmailRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ChangeEmailRequest) SetEmail(v string) {
	o.Email = v
}

func (o ChangeEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeEmailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableChangeEmailRequest struct {
	value *ChangeEmailRequest
	isSet bool
}

func (v NullableChangeEmailRequest) Get() *ChangeEmailRequest {
	return v.value
}

func (v *NullableChangeEmailRequest) Set(val *ChangeEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeEmailRequest(val *ChangeEmailRequest) *NullableChangeEmailRequest {
	return &NullableChangeEmailRequest{value: val, isSet: true}
}

func (v NullableChangeEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


