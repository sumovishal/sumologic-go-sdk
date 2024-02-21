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

// checks if the AlertsLibraryFolderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsLibraryFolderResponse{}

// AlertsLibraryFolderResponse struct for AlertsLibraryFolderResponse
type AlertsLibraryFolderResponse struct {
	AlertsLibraryBaseResponse
	// Children of the folder. NOTE: Permissions field will not be filled (empty list) for children.
	Children []AlertsLibraryBaseResponse `json:"children"`
}

// NewAlertsLibraryFolderResponse instantiates a new AlertsLibraryFolderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryFolderResponse(children []AlertsLibraryBaseResponse, id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool) *AlertsLibraryFolderResponse {
	this := AlertsLibraryFolderResponse{}
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
	this.Children = children
	return &this
}

// NewAlertsLibraryFolderResponseWithDefaults instantiates a new AlertsLibraryFolderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryFolderResponseWithDefaults() *AlertsLibraryFolderResponse {
	this := AlertsLibraryFolderResponse{}
	return &this
}

// GetChildren returns the Children field value
func (o *AlertsLibraryFolderResponse) GetChildren() []AlertsLibraryBaseResponse {
	if o == nil {
		var ret []AlertsLibraryBaseResponse
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryFolderResponse) GetChildrenOk() ([]AlertsLibraryBaseResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *AlertsLibraryFolderResponse) SetChildren(v []AlertsLibraryBaseResponse) {
	o.Children = v
}

func (o AlertsLibraryFolderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsLibraryFolderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAlertsLibraryBaseResponse, errAlertsLibraryBaseResponse := json.Marshal(o.AlertsLibraryBaseResponse)
	if errAlertsLibraryBaseResponse != nil {
		return map[string]interface{}{}, errAlertsLibraryBaseResponse
	}
	errAlertsLibraryBaseResponse = json.Unmarshal([]byte(serializedAlertsLibraryBaseResponse), &toSerialize)
	if errAlertsLibraryBaseResponse != nil {
		return map[string]interface{}{}, errAlertsLibraryBaseResponse
	}
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

type NullableAlertsLibraryFolderResponse struct {
	value *AlertsLibraryFolderResponse
	isSet bool
}

func (v NullableAlertsLibraryFolderResponse) Get() *AlertsLibraryFolderResponse {
	return v.value
}

func (v *NullableAlertsLibraryFolderResponse) Set(val *AlertsLibraryFolderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryFolderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryFolderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryFolderResponse(val *AlertsLibraryFolderResponse) *NullableAlertsLibraryFolderResponse {
	return &NullableAlertsLibraryFolderResponse{value: val, isSet: true}
}

func (v NullableAlertsLibraryFolderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryFolderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


