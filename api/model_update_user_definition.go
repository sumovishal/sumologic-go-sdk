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

// checks if the UpdateUserDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserDefinition{}

// UpdateUserDefinition struct for UpdateUserDefinition
type UpdateUserDefinition struct {
	// First name of the user. If the caller has `manageUsersAndRoles` capability, this field can be updated for any user. If the caller does NOT have `manageUsersAndRoles` capability, then only the calling user's firstName can be updated.
	FirstName string `json:"firstName"`
	// Last name of the user. If the caller has `manageUsersAndRoles` capability, this field can be updated for any user. If the caller does NOT have `manageUsersAndRoles` capability, then only the calling user's lastName can be updated.
	LastName string `json:"lastName"`
	// This has the value `true` if the user is active and `false` if they have been deactivated. To modify this field you must have the `manageUserAndRoles` capability.
	IsActive *bool `json:"isActive,omitempty"`
	// List of role identifiers associated with the user. To modify this field you must have the `manageUserAndRoles` capability.
	RoleIds []string `json:"roleIds,omitempty"`
}

// NewUpdateUserDefinition instantiates a new UpdateUserDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserDefinition(firstName string, lastName string) *UpdateUserDefinition {
	this := UpdateUserDefinition{}
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewUpdateUserDefinitionWithDefaults instantiates a new UpdateUserDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserDefinitionWithDefaults() *UpdateUserDefinition {
	this := UpdateUserDefinition{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *UpdateUserDefinition) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UpdateUserDefinition) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UpdateUserDefinition) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UpdateUserDefinition) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UpdateUserDefinition) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UpdateUserDefinition) SetLastName(v string) {
	o.LastName = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UpdateUserDefinition) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserDefinition) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateUserDefinition) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UpdateUserDefinition) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *UpdateUserDefinition) GetRoleIds() []string {
	if o == nil || IsNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserDefinition) GetRoleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *UpdateUserDefinition) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *UpdateUserDefinition) SetRoleIds(v []string) {
	o.RoleIds = v
}

func (o UpdateUserDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	return toSerialize, nil
}

type NullableUpdateUserDefinition struct {
	value *UpdateUserDefinition
	isSet bool
}

func (v NullableUpdateUserDefinition) Get() *UpdateUserDefinition {
	return v.value
}

func (v *NullableUpdateUserDefinition) Set(val *UpdateUserDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserDefinition(val *UpdateUserDefinition) *NullableUpdateUserDefinition {
	return &NullableUpdateUserDefinition{value: val, isSet: true}
}

func (v NullableUpdateUserDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


