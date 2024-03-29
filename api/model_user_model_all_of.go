/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// checks if the UserModelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModelAllOf{}

// UserModelAllOf struct for UserModelAllOf
type UserModelAllOf struct {
	// Unique identifier for the user.
	Id string `json:"id"`
	// True if the user is active.
	IsActive *bool `json:"isActive,omitempty"`
	// This has the value `true` if the user's account has been locked. If a user tries to log into their account several times and fails, his or her account will be locked for security reasons.
	IsLocked *bool `json:"isLocked,omitempty"`
	// True if multi factor authentication is enabled for the user.
	IsMfaEnabled *bool `json:"isMfaEnabled,omitempty"`
	// Timestamp of the last login for the user in UTC. Will be null if the user has never logged in.
	LastLoginTimestamp *time.Time `json:"lastLoginTimestamp,omitempty"`
}

// NewUserModelAllOf instantiates a new UserModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModelAllOf(id string) *UserModelAllOf {
	this := UserModelAllOf{}
	this.Id = id
	return &this
}

// NewUserModelAllOfWithDefaults instantiates a new UserModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelAllOfWithDefaults() *UserModelAllOf {
	this := UserModelAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *UserModelAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserModelAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserModelAllOf) SetId(v string) {
	o.Id = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UserModelAllOf) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelAllOf) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UserModelAllOf) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UserModelAllOf) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *UserModelAllOf) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelAllOf) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *UserModelAllOf) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *UserModelAllOf) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsMfaEnabled returns the IsMfaEnabled field value if set, zero value otherwise.
func (o *UserModelAllOf) GetIsMfaEnabled() bool {
	if o == nil || IsNil(o.IsMfaEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMfaEnabled
}

// GetIsMfaEnabledOk returns a tuple with the IsMfaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelAllOf) GetIsMfaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMfaEnabled) {
		return nil, false
	}
	return o.IsMfaEnabled, true
}

// HasIsMfaEnabled returns a boolean if a field has been set.
func (o *UserModelAllOf) HasIsMfaEnabled() bool {
	if o != nil && !IsNil(o.IsMfaEnabled) {
		return true
	}

	return false
}

// SetIsMfaEnabled gets a reference to the given bool and assigns it to the IsMfaEnabled field.
func (o *UserModelAllOf) SetIsMfaEnabled(v bool) {
	o.IsMfaEnabled = &v
}

// GetLastLoginTimestamp returns the LastLoginTimestamp field value if set, zero value otherwise.
func (o *UserModelAllOf) GetLastLoginTimestamp() time.Time {
	if o == nil || IsNil(o.LastLoginTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginTimestamp
}

// GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelAllOf) GetLastLoginTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastLoginTimestamp) {
		return nil, false
	}
	return o.LastLoginTimestamp, true
}

// HasLastLoginTimestamp returns a boolean if a field has been set.
func (o *UserModelAllOf) HasLastLoginTimestamp() bool {
	if o != nil && !IsNil(o.LastLoginTimestamp) {
		return true
	}

	return false
}

// SetLastLoginTimestamp gets a reference to the given time.Time and assigns it to the LastLoginTimestamp field.
func (o *UserModelAllOf) SetLastLoginTimestamp(v time.Time) {
	o.LastLoginTimestamp = &v
}

func (o UserModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.IsMfaEnabled) {
		toSerialize["isMfaEnabled"] = o.IsMfaEnabled
	}
	if !IsNil(o.LastLoginTimestamp) {
		toSerialize["lastLoginTimestamp"] = o.LastLoginTimestamp
	}
	return toSerialize, nil
}

type NullableUserModelAllOf struct {
	value *UserModelAllOf
	isSet bool
}

func (v NullableUserModelAllOf) Get() *UserModelAllOf {
	return v.value
}

func (v *NullableUserModelAllOf) Set(val *UserModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModelAllOf(val *UserModelAllOf) *NullableUserModelAllOf {
	return &NullableUserModelAllOf{value: val, isSet: true}
}

func (v NullableUserModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


