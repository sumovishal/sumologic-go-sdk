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

// checks if the DisableMfaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisableMfaRequest{}

// DisableMfaRequest struct for DisableMfaRequest
type DisableMfaRequest struct {
	// Email of user whose mfa is being disabled.
	Email string `json:"email"`
	// Password of user whose mfa is being disabled.
	Password string `json:"password"`
}

// NewDisableMfaRequest instantiates a new DisableMfaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableMfaRequest(email string, password string) *DisableMfaRequest {
	this := DisableMfaRequest{}
	this.Email = email
	this.Password = password
	return &this
}

// NewDisableMfaRequestWithDefaults instantiates a new DisableMfaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableMfaRequestWithDefaults() *DisableMfaRequest {
	this := DisableMfaRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *DisableMfaRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *DisableMfaRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *DisableMfaRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *DisableMfaRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *DisableMfaRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *DisableMfaRequest) SetPassword(v string) {
	o.Password = v
}

func (o DisableMfaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableMfaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableDisableMfaRequest struct {
	value *DisableMfaRequest
	isSet bool
}

func (v NullableDisableMfaRequest) Get() *DisableMfaRequest {
	return v.value
}

func (v *NullableDisableMfaRequest) Set(val *DisableMfaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableMfaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableMfaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableMfaRequest(val *DisableMfaRequest) *NullableDisableMfaRequest {
	return &NullableDisableMfaRequest{value: val, isSet: true}
}

func (v NullableDisableMfaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableMfaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


