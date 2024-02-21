/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the AllowlistedUserResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllowlistedUserResult{}

// AllowlistedUserResult struct for AllowlistedUserResult
type AllowlistedUserResult struct {
	// Unique identifier of the user.
	UserId string `json:"userId"`
	// First name of the user.
	FirstName string `json:"firstName"`
	// Last name of the user.
	LastName string `json:"lastName"`
	// Email of the user.
	Email string `json:"email"`
	// If the user can manage SAML Configurations.
	CanManageSaml bool `json:"canManageSaml"`
	// Checks if the user is active.
	IsActive bool `json:"isActive"`
	// Timestamp of the last login of the user.
	LastLogin time.Time `json:"lastLogin"`
}

// NewAllowlistedUserResult instantiates a new AllowlistedUserResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowlistedUserResult(userId string, firstName string, lastName string, email string, canManageSaml bool, isActive bool, lastLogin time.Time) *AllowlistedUserResult {
	this := AllowlistedUserResult{}
	this.UserId = userId
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.CanManageSaml = canManageSaml
	this.IsActive = isActive
	this.LastLogin = lastLogin
	return &this
}

// NewAllowlistedUserResultWithDefaults instantiates a new AllowlistedUserResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowlistedUserResultWithDefaults() *AllowlistedUserResult {
	this := AllowlistedUserResult{}
	return &this
}

// GetUserId returns the UserId field value
func (o *AllowlistedUserResult) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AllowlistedUserResult) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AllowlistedUserResult) SetUserId(v string) {
	o.UserId = v
}

// GetFirstName returns the FirstName field value
func (o *AllowlistedUserResult) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *AllowlistedUserResult) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *AllowlistedUserResult) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *AllowlistedUserResult) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *AllowlistedUserResult) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *AllowlistedUserResult) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *AllowlistedUserResult) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AllowlistedUserResult) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AllowlistedUserResult) SetEmail(v string) {
	o.Email = v
}

// GetCanManageSaml returns the CanManageSaml field value
func (o *AllowlistedUserResult) GetCanManageSaml() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanManageSaml
}

// GetCanManageSamlOk returns a tuple with the CanManageSaml field value
// and a boolean to check if the value has been set.
func (o *AllowlistedUserResult) GetCanManageSamlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanManageSaml, true
}

// SetCanManageSaml sets field value
func (o *AllowlistedUserResult) SetCanManageSaml(v bool) {
	o.CanManageSaml = v
}

// GetIsActive returns the IsActive field value
func (o *AllowlistedUserResult) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *AllowlistedUserResult) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *AllowlistedUserResult) SetIsActive(v bool) {
	o.IsActive = v
}

// GetLastLogin returns the LastLogin field value
func (o *AllowlistedUserResult) GetLastLogin() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value
// and a boolean to check if the value has been set.
func (o *AllowlistedUserResult) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastLogin, true
}

// SetLastLogin sets field value
func (o *AllowlistedUserResult) SetLastLogin(v time.Time) {
	o.LastLogin = v
}

func (o AllowlistedUserResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllowlistedUserResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["email"] = o.Email
	toSerialize["canManageSaml"] = o.CanManageSaml
	toSerialize["isActive"] = o.IsActive
	toSerialize["lastLogin"] = o.LastLogin
	return toSerialize, nil
}

type NullableAllowlistedUserResult struct {
	value *AllowlistedUserResult
	isSet bool
}

func (v NullableAllowlistedUserResult) Get() *AllowlistedUserResult {
	return v.value
}

func (v *NullableAllowlistedUserResult) Set(val *AllowlistedUserResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowlistedUserResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowlistedUserResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowlistedUserResult(val *AllowlistedUserResult) *NullableAllowlistedUserResult {
	return &NullableAllowlistedUserResult{value: val, isSet: true}
}

func (v NullableAllowlistedUserResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowlistedUserResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


