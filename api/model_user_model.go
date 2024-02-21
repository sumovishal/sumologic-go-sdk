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

// checks if the UserModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModel{}

// UserModel struct for UserModel
type UserModel struct {
	// First name of the user.
	FirstName string `json:"firstName"`
	// Last name of the user.
	LastName string `json:"lastName"`
	// Email address of the user.
	Email string `json:"email"`
	// List of roleIds associated with the user.
	RoleIds []string `json:"roleIds"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
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

// NewUserModel instantiates a new UserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModel(firstName string, lastName string, email string, roleIds []string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *UserModel {
	this := UserModel{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.RoleIds = roleIds
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	return &this
}

// NewUserModelWithDefaults instantiates a new UserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelWithDefaults() *UserModel {
	this := UserModel{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *UserModel) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserModel) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UserModel) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserModel) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *UserModel) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserModel) SetEmail(v string) {
	o.Email = v
}

// GetRoleIds returns the RoleIds field value
func (o *UserModel) GetRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetRoleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// SetRoleIds sets field value
func (o *UserModel) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserModel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *UserModel) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *UserModel) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *UserModel) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *UserModel) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *UserModel) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *UserModel) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *UserModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserModel) SetId(v string) {
	o.Id = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UserModel) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UserModel) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UserModel) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *UserModel) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *UserModel) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *UserModel) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsMfaEnabled returns the IsMfaEnabled field value if set, zero value otherwise.
func (o *UserModel) GetIsMfaEnabled() bool {
	if o == nil || IsNil(o.IsMfaEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMfaEnabled
}

// GetIsMfaEnabledOk returns a tuple with the IsMfaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetIsMfaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMfaEnabled) {
		return nil, false
	}
	return o.IsMfaEnabled, true
}

// HasIsMfaEnabled returns a boolean if a field has been set.
func (o *UserModel) HasIsMfaEnabled() bool {
	if o != nil && !IsNil(o.IsMfaEnabled) {
		return true
	}

	return false
}

// SetIsMfaEnabled gets a reference to the given bool and assigns it to the IsMfaEnabled field.
func (o *UserModel) SetIsMfaEnabled(v bool) {
	o.IsMfaEnabled = &v
}

// GetLastLoginTimestamp returns the LastLoginTimestamp field value if set, zero value otherwise.
func (o *UserModel) GetLastLoginTimestamp() time.Time {
	if o == nil || IsNil(o.LastLoginTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginTimestamp
}

// GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetLastLoginTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastLoginTimestamp) {
		return nil, false
	}
	return o.LastLoginTimestamp, true
}

// HasLastLoginTimestamp returns a boolean if a field has been set.
func (o *UserModel) HasLastLoginTimestamp() bool {
	if o != nil && !IsNil(o.LastLoginTimestamp) {
		return true
	}

	return false
}

// SetLastLoginTimestamp gets a reference to the given time.Time and assigns it to the LastLoginTimestamp field.
func (o *UserModel) SetLastLoginTimestamp(v time.Time) {
	o.LastLoginTimestamp = &v
}

func (o UserModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["email"] = o.Email
	toSerialize["roleIds"] = o.RoleIds
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
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

type NullableUserModel struct {
	value *UserModel
	isSet bool
}

func (v NullableUserModel) Get() *UserModel {
	return v.value
}

func (v *NullableUserModel) Set(val *UserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModel(val *UserModel) *NullableUserModel {
	return &NullableUserModel{value: val, isSet: true}
}

func (v NullableUserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


