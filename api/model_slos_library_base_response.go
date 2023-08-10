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

// checks if the SlosLibraryBaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlosLibraryBaseResponse{}

// SlosLibraryBaseResponse struct for SlosLibraryBaseResponse
type SlosLibraryBaseResponse struct {
	// Identifier of the slo or folder.
	Id string `json:"id"`
	// Identifier of the slo or folder.
	Name string `json:"name"`
	// Description of the slo or folder.
	Description string `json:"description"`
	// Version of the slo or folder.
	Version int64 `json:"version"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Identifier of the parent folder.
	ParentId string `json:"parentId"`
	// Type of the content. Valid values:   1) Slo   2) Folder
	ContentType string `json:"contentType"`
	// Type of the object model.
	Type string `json:"type"`
	// System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can't be updated.
	IsSystem bool `json:"isSystem"`
	// Immutable objects are \"READ-ONLY\".
	IsMutable bool `json:"isMutable"`
	// Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint.
	Permissions []string `json:"permissions,omitempty"`
}

// NewSlosLibraryBaseResponse instantiates a new SlosLibraryBaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibraryBaseResponse(id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool) *SlosLibraryBaseResponse {
	this := SlosLibraryBaseResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Version = version
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.ParentId = parentId
	this.ContentType = contentType
	this.Type = type_
	this.IsSystem = isSystem
	this.IsMutable = isMutable
	return &this
}

// NewSlosLibraryBaseResponseWithDefaults instantiates a new SlosLibraryBaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibraryBaseResponseWithDefaults() *SlosLibraryBaseResponse {
	this := SlosLibraryBaseResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SlosLibraryBaseResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SlosLibraryBaseResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SlosLibraryBaseResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlosLibraryBaseResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *SlosLibraryBaseResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SlosLibraryBaseResponse) SetDescription(v string) {
	o.Description = v
}

// GetVersion returns the Version field value
func (o *SlosLibraryBaseResponse) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SlosLibraryBaseResponse) SetVersion(v int64) {
	o.Version = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SlosLibraryBaseResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SlosLibraryBaseResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SlosLibraryBaseResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SlosLibraryBaseResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *SlosLibraryBaseResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *SlosLibraryBaseResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *SlosLibraryBaseResponse) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *SlosLibraryBaseResponse) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetParentId returns the ParentId field value
func (o *SlosLibraryBaseResponse) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *SlosLibraryBaseResponse) SetParentId(v string) {
	o.ParentId = v
}

// GetContentType returns the ContentType field value
func (o *SlosLibraryBaseResponse) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *SlosLibraryBaseResponse) SetContentType(v string) {
	o.ContentType = v
}

// GetType returns the Type field value
func (o *SlosLibraryBaseResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlosLibraryBaseResponse) SetType(v string) {
	o.Type = v
}

// GetIsSystem returns the IsSystem field value
func (o *SlosLibraryBaseResponse) GetIsSystem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetIsSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystem, true
}

// SetIsSystem sets field value
func (o *SlosLibraryBaseResponse) SetIsSystem(v bool) {
	o.IsSystem = v
}

// GetIsMutable returns the IsMutable field value
func (o *SlosLibraryBaseResponse) GetIsMutable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetIsMutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMutable, true
}

// SetIsMutable sets field value
func (o *SlosLibraryBaseResponse) SetIsMutable(v bool) {
	o.IsMutable = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *SlosLibraryBaseResponse) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibraryBaseResponse) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *SlosLibraryBaseResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *SlosLibraryBaseResponse) SetPermissions(v []string) {
	o.Permissions = v
}

func (o SlosLibraryBaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlosLibraryBaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["version"] = o.Version
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["parentId"] = o.ParentId
	toSerialize["contentType"] = o.ContentType
	toSerialize["type"] = o.Type
	toSerialize["isSystem"] = o.IsSystem
	toSerialize["isMutable"] = o.IsMutable
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableSlosLibraryBaseResponse struct {
	value *SlosLibraryBaseResponse
	isSet bool
}

func (v NullableSlosLibraryBaseResponse) Get() *SlosLibraryBaseResponse {
	return v.value
}

func (v *NullableSlosLibraryBaseResponse) Set(val *SlosLibraryBaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibraryBaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibraryBaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibraryBaseResponse(val *SlosLibraryBaseResponse) *NullableSlosLibraryBaseResponse {
	return &NullableSlosLibraryBaseResponse{value: val, isSet: true}
}

func (v NullableSlosLibraryBaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibraryBaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


