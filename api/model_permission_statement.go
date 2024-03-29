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

// checks if the PermissionStatement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionStatement{}

// PermissionStatement struct for PermissionStatement
type PermissionStatement struct {
	// List of permissions.
	Permissions []string `json:"permissions"`
	// Type of subject for the permission. Valid values are: `role` or `org`.
	SubjectType string `json:"subjectType"`
	// The identifier that belongs to the subject type chosen above. For e.g. if the subjectType is set to `role`, subjectId should be the identifier of a role.  Similarly, if the subjectType is `org`, the subjectId should be the identifier of the same org,  which owns the resource target.
	SubjectId string `json:"subjectId"`
	// The identifier that belongs to the resource this permission assignment applies to.
	TargetId string `json:"targetId"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
}

// NewPermissionStatement instantiates a new PermissionStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionStatement(permissions []string, subjectType string, subjectId string, targetId string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *PermissionStatement {
	this := PermissionStatement{}
	this.Permissions = permissions
	this.SubjectType = subjectType
	this.SubjectId = subjectId
	this.TargetId = targetId
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	return &this
}

// NewPermissionStatementWithDefaults instantiates a new PermissionStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionStatementWithDefaults() *PermissionStatement {
	this := PermissionStatement{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *PermissionStatement) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *PermissionStatement) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *PermissionStatement) SetPermissions(v []string) {
	o.Permissions = v
}

// GetSubjectType returns the SubjectType field value
func (o *PermissionStatement) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *PermissionStatement) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value
func (o *PermissionStatement) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetSubjectId returns the SubjectId field value
func (o *PermissionStatement) GetSubjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value
// and a boolean to check if the value has been set.
func (o *PermissionStatement) GetSubjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectId, true
}

// SetSubjectId sets field value
func (o *PermissionStatement) SetSubjectId(v string) {
	o.SubjectId = v
}

// GetTargetId returns the TargetId field value
func (o *PermissionStatement) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *PermissionStatement) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *PermissionStatement) SetTargetId(v string) {
	o.TargetId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PermissionStatement) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PermissionStatement) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PermissionStatement) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *PermissionStatement) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *PermissionStatement) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *PermissionStatement) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *PermissionStatement) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *PermissionStatement) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *PermissionStatement) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *PermissionStatement) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *PermissionStatement) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *PermissionStatement) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

func (o PermissionStatement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionStatement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	toSerialize["subjectType"] = o.SubjectType
	toSerialize["subjectId"] = o.SubjectId
	toSerialize["targetId"] = o.TargetId
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	return toSerialize, nil
}

type NullablePermissionStatement struct {
	value *PermissionStatement
	isSet bool
}

func (v NullablePermissionStatement) Get() *PermissionStatement {
	return v.value
}

func (v *NullablePermissionStatement) Set(val *PermissionStatement) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionStatement) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionStatement(val *PermissionStatement) *NullablePermissionStatement {
	return &NullablePermissionStatement{value: val, isSet: true}
}

func (v NullablePermissionStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


