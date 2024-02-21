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

// checks if the ListRoleModelsResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRoleModelsResponseV2{}

// ListRoleModelsResponseV2 struct for ListRoleModelsResponseV2
type ListRoleModelsResponseV2 struct {
	// List of roles.
	Data []GetRoleDefinitionV2 `json:"data"`
	// Next continuation token.
	Next *string `json:"next,omitempty"`
}

// NewListRoleModelsResponseV2 instantiates a new ListRoleModelsResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRoleModelsResponseV2(data []GetRoleDefinitionV2) *ListRoleModelsResponseV2 {
	this := ListRoleModelsResponseV2{}
	this.Data = data
	return &this
}

// NewListRoleModelsResponseV2WithDefaults instantiates a new ListRoleModelsResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRoleModelsResponseV2WithDefaults() *ListRoleModelsResponseV2 {
	this := ListRoleModelsResponseV2{}
	return &this
}

// GetData returns the Data field value
func (o *ListRoleModelsResponseV2) GetData() []GetRoleDefinitionV2 {
	if o == nil {
		var ret []GetRoleDefinitionV2
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListRoleModelsResponseV2) GetDataOk() ([]GetRoleDefinitionV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListRoleModelsResponseV2) SetData(v []GetRoleDefinitionV2) {
	o.Data = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListRoleModelsResponseV2) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRoleModelsResponseV2) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListRoleModelsResponseV2) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListRoleModelsResponseV2) SetNext(v string) {
	o.Next = &v
}

func (o ListRoleModelsResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRoleModelsResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableListRoleModelsResponseV2 struct {
	value *ListRoleModelsResponseV2
	isSet bool
}

func (v NullableListRoleModelsResponseV2) Get() *ListRoleModelsResponseV2 {
	return v.value
}

func (v *NullableListRoleModelsResponseV2) Set(val *ListRoleModelsResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableListRoleModelsResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableListRoleModelsResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRoleModelsResponseV2(val *ListRoleModelsResponseV2) *NullableListRoleModelsResponseV2 {
	return &NullableListRoleModelsResponseV2{value: val, isSet: true}
}

func (v NullableListRoleModelsResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRoleModelsResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


