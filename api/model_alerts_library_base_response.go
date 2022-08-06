/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// AlertsLibraryBaseResponse struct for AlertsLibraryBaseResponse
type AlertsLibraryBaseResponse struct {
	// Identifier of the alert or folder.
	Id string `json:"id"`
	// Identifier of the alert or folder.
	Name string `json:"name"`
	// Description of the alert or folder.
	Description string `json:"description"`
	// Version of the alert or folder.
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
	// Type of the content. Valid values:   1) Alert   2) Folder
	ContentType string `json:"contentType"`
	// Type of the object model.
	Type string `json:"type"`
	// Whether the object is locked.
	IsLocked *bool `json:"isLocked,omitempty"`
	// System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can't be updated.
	IsSystem bool `json:"isSystem"`
	// Immutable objects are \"READ-ONLY\".
	IsMutable bool `json:"isMutable"`
}

// NewAlertsLibraryBaseResponse instantiates a new AlertsLibraryBaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryBaseResponse(id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool) *AlertsLibraryBaseResponse {
	this := AlertsLibraryBaseResponse{}
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

// NewAlertsLibraryBaseResponseWithDefaults instantiates a new AlertsLibraryBaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryBaseResponseWithDefaults() *AlertsLibraryBaseResponse {
	this := AlertsLibraryBaseResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AlertsLibraryBaseResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AlertsLibraryBaseResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AlertsLibraryBaseResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlertsLibraryBaseResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AlertsLibraryBaseResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AlertsLibraryBaseResponse) SetDescription(v string) {
	o.Description = v
}

// GetVersion returns the Version field value
func (o *AlertsLibraryBaseResponse) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AlertsLibraryBaseResponse) SetVersion(v int64) {
	o.Version = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AlertsLibraryBaseResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AlertsLibraryBaseResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *AlertsLibraryBaseResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *AlertsLibraryBaseResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *AlertsLibraryBaseResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *AlertsLibraryBaseResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *AlertsLibraryBaseResponse) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *AlertsLibraryBaseResponse) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetParentId returns the ParentId field value
func (o *AlertsLibraryBaseResponse) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *AlertsLibraryBaseResponse) SetParentId(v string) {
	o.ParentId = v
}

// GetContentType returns the ContentType field value
func (o *AlertsLibraryBaseResponse) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *AlertsLibraryBaseResponse) SetContentType(v string) {
	o.ContentType = v
}

// GetType returns the Type field value
func (o *AlertsLibraryBaseResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AlertsLibraryBaseResponse) SetType(v string) {
	o.Type = v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *AlertsLibraryBaseResponse) GetIsLocked() bool {
	if o == nil || o.IsLocked == nil {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetIsLockedOk() (*bool, bool) {
	if o == nil || o.IsLocked == nil {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *AlertsLibraryBaseResponse) HasIsLocked() bool {
	if o != nil && o.IsLocked != nil {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *AlertsLibraryBaseResponse) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsSystem returns the IsSystem field value
func (o *AlertsLibraryBaseResponse) GetIsSystem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetIsSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystem, true
}

// SetIsSystem sets field value
func (o *AlertsLibraryBaseResponse) SetIsSystem(v bool) {
	o.IsSystem = v
}

// GetIsMutable returns the IsMutable field value
func (o *AlertsLibraryBaseResponse) GetIsMutable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseResponse) GetIsMutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMutable, true
}

// SetIsMutable sets field value
func (o *AlertsLibraryBaseResponse) SetIsMutable(v bool) {
	o.IsMutable = v
}

func (o AlertsLibraryBaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if true {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if true {
		toSerialize["parentId"] = o.ParentId
	}
	if true {
		toSerialize["contentType"] = o.ContentType
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.IsLocked != nil {
		toSerialize["isLocked"] = o.IsLocked
	}
	if true {
		toSerialize["isSystem"] = o.IsSystem
	}
	if true {
		toSerialize["isMutable"] = o.IsMutable
	}
	return json.Marshal(toSerialize)
}

type NullableAlertsLibraryBaseResponse struct {
	value *AlertsLibraryBaseResponse
	isSet bool
}

func (v NullableAlertsLibraryBaseResponse) Get() *AlertsLibraryBaseResponse {
	return v.value
}

func (v *NullableAlertsLibraryBaseResponse) Set(val *AlertsLibraryBaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryBaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryBaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryBaseResponse(val *AlertsLibraryBaseResponse) *NullableAlertsLibraryBaseResponse {
	return &NullableAlertsLibraryBaseResponse{value: val, isSet: true}
}

func (v NullableAlertsLibraryBaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryBaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


