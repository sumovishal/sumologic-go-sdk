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

// checks if the ServiceAccountModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountModel{}

// ServiceAccountModel struct for ServiceAccountModel
type ServiceAccountModel struct {
	// Name of the service account.
	Name string `json:"name"`
	// Email address of the service account.
	Email string `json:"email"`
	// List of roleIds associated with the service account.
	RoleIds []string `json:"roleIds"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier for the service account.
	Id string `json:"id"`
	// True if the service account is active.
	IsActive *bool `json:"isActive,omitempty"`
}

// NewServiceAccountModel instantiates a new ServiceAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountModel(name string, email string, roleIds []string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *ServiceAccountModel {
	this := ServiceAccountModel{}
	this.Name = name
	this.Email = email
	this.RoleIds = roleIds
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	return &this
}

// NewServiceAccountModelWithDefaults instantiates a new ServiceAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountModelWithDefaults() *ServiceAccountModel {
	this := ServiceAccountModel{}
	return &this
}

// GetName returns the Name field value
func (o *ServiceAccountModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceAccountModel) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *ServiceAccountModel) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ServiceAccountModel) SetEmail(v string) {
	o.Email = v
}

// GetRoleIds returns the RoleIds field value
func (o *ServiceAccountModel) GetRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetRoleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// SetRoleIds sets field value
func (o *ServiceAccountModel) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceAccountModel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceAccountModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ServiceAccountModel) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ServiceAccountModel) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *ServiceAccountModel) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *ServiceAccountModel) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *ServiceAccountModel) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *ServiceAccountModel) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *ServiceAccountModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAccountModel) SetId(v string) {
	o.Id = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ServiceAccountModel) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountModel) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ServiceAccountModel) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ServiceAccountModel) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o ServiceAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
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
	return toSerialize, nil
}

type NullableServiceAccountModel struct {
	value *ServiceAccountModel
	isSet bool
}

func (v NullableServiceAccountModel) Get() *ServiceAccountModel {
	return v.value
}

func (v *NullableServiceAccountModel) Set(val *ServiceAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountModel(val *ServiceAccountModel) *NullableServiceAccountModel {
	return &NullableServiceAccountModel{value: val, isSet: true}
}

func (v NullableServiceAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


