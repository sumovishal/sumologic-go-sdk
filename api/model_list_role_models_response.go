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

// checks if the ListRoleModelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRoleModelsResponse{}

// ListRoleModelsResponse struct for ListRoleModelsResponse
type ListRoleModelsResponse struct {
	// List of roles.
	Data []RoleModel `json:"data"`
	// Next continuation token.
	Next *string `json:"next,omitempty"`
}

// NewListRoleModelsResponse instantiates a new ListRoleModelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRoleModelsResponse(data []RoleModel) *ListRoleModelsResponse {
	this := ListRoleModelsResponse{}
	this.Data = data
	return &this
}

// NewListRoleModelsResponseWithDefaults instantiates a new ListRoleModelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRoleModelsResponseWithDefaults() *ListRoleModelsResponse {
	this := ListRoleModelsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListRoleModelsResponse) GetData() []RoleModel {
	if o == nil {
		var ret []RoleModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListRoleModelsResponse) GetDataOk() ([]RoleModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListRoleModelsResponse) SetData(v []RoleModel) {
	o.Data = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListRoleModelsResponse) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRoleModelsResponse) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListRoleModelsResponse) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListRoleModelsResponse) SetNext(v string) {
	o.Next = &v
}

func (o ListRoleModelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRoleModelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableListRoleModelsResponse struct {
	value *ListRoleModelsResponse
	isSet bool
}

func (v NullableListRoleModelsResponse) Get() *ListRoleModelsResponse {
	return v.value
}

func (v *NullableListRoleModelsResponse) Set(val *ListRoleModelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRoleModelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRoleModelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRoleModelsResponse(val *ListRoleModelsResponse) *NullableListRoleModelsResponse {
	return &NullableListRoleModelsResponse{value: val, isSet: true}
}

func (v NullableListRoleModelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRoleModelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


