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

// checks if the ListDroppedFieldsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDroppedFieldsResponse{}

// ListDroppedFieldsResponse struct for ListDroppedFieldsResponse
type ListDroppedFieldsResponse struct {
	// List of dropped fields.
	Data []DroppedField `json:"data"`
}

// NewListDroppedFieldsResponse instantiates a new ListDroppedFieldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDroppedFieldsResponse(data []DroppedField) *ListDroppedFieldsResponse {
	this := ListDroppedFieldsResponse{}
	this.Data = data
	return &this
}

// NewListDroppedFieldsResponseWithDefaults instantiates a new ListDroppedFieldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDroppedFieldsResponseWithDefaults() *ListDroppedFieldsResponse {
	this := ListDroppedFieldsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListDroppedFieldsResponse) GetData() []DroppedField {
	if o == nil {
		var ret []DroppedField
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListDroppedFieldsResponse) GetDataOk() ([]DroppedField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListDroppedFieldsResponse) SetData(v []DroppedField) {
	o.Data = v
}

func (o ListDroppedFieldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDroppedFieldsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListDroppedFieldsResponse struct {
	value *ListDroppedFieldsResponse
	isSet bool
}

func (v NullableListDroppedFieldsResponse) Get() *ListDroppedFieldsResponse {
	return v.value
}

func (v *NullableListDroppedFieldsResponse) Set(val *ListDroppedFieldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDroppedFieldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDroppedFieldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDroppedFieldsResponse(val *ListDroppedFieldsResponse) *NullableListDroppedFieldsResponse {
	return &NullableListDroppedFieldsResponse{value: val, isSet: true}
}

func (v NullableListDroppedFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDroppedFieldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


