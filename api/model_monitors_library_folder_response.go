/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MonitorsLibraryFolderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorsLibraryFolderResponse{}

// MonitorsLibraryFolderResponse struct for MonitorsLibraryFolderResponse
type MonitorsLibraryFolderResponse struct {
	MonitorsLibraryBaseResponse
	// Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint.
	Permissions []string `json:"permissions"`
	// Children of the folder. NOTE: Permissions field will not be filled (empty list) for children.
	Children []MonitorsLibraryBaseResponse `json:"children"`
}

// NewMonitorsLibraryFolderResponse instantiates a new MonitorsLibraryFolderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorsLibraryFolderResponse(permissions []string, children []MonitorsLibraryBaseResponse, id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool) *MonitorsLibraryFolderResponse {
	this := MonitorsLibraryFolderResponse{}
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
	this.Permissions = permissions
	this.Children = children
	return &this
}

// NewMonitorsLibraryFolderResponseWithDefaults instantiates a new MonitorsLibraryFolderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorsLibraryFolderResponseWithDefaults() *MonitorsLibraryFolderResponse {
	this := MonitorsLibraryFolderResponse{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *MonitorsLibraryFolderResponse) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryFolderResponse) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *MonitorsLibraryFolderResponse) SetPermissions(v []string) {
	o.Permissions = v
}

// GetChildren returns the Children field value
func (o *MonitorsLibraryFolderResponse) GetChildren() []MonitorsLibraryBaseResponse {
	if o == nil {
		var ret []MonitorsLibraryBaseResponse
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryFolderResponse) GetChildrenOk() ([]MonitorsLibraryBaseResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *MonitorsLibraryFolderResponse) SetChildren(v []MonitorsLibraryBaseResponse) {
	o.Children = v
}

func (o MonitorsLibraryFolderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorsLibraryFolderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMonitorsLibraryBaseResponse, errMonitorsLibraryBaseResponse := json.Marshal(o.MonitorsLibraryBaseResponse)
	if errMonitorsLibraryBaseResponse != nil {
		return map[string]interface{}{}, errMonitorsLibraryBaseResponse
	}
	errMonitorsLibraryBaseResponse = json.Unmarshal([]byte(serializedMonitorsLibraryBaseResponse), &toSerialize)
	if errMonitorsLibraryBaseResponse != nil {
		return map[string]interface{}{}, errMonitorsLibraryBaseResponse
	}
	toSerialize["permissions"] = o.Permissions
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

type NullableMonitorsLibraryFolderResponse struct {
	value *MonitorsLibraryFolderResponse
	isSet bool
}

func (v NullableMonitorsLibraryFolderResponse) Get() *MonitorsLibraryFolderResponse {
	return v.value
}

func (v *NullableMonitorsLibraryFolderResponse) Set(val *MonitorsLibraryFolderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorsLibraryFolderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorsLibraryFolderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorsLibraryFolderResponse(val *MonitorsLibraryFolderResponse) *NullableMonitorsLibraryFolderResponse {
	return &NullableMonitorsLibraryFolderResponse{value: val, isSet: true}
}

func (v NullableMonitorsLibraryFolderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorsLibraryFolderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


