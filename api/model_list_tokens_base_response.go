/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// ListTokensBaseResponse struct for ListTokensBaseResponse
type ListTokensBaseResponse struct {
	// List of tokens.
	Data []TokenBaseResponse `json:"data"`
}

// NewListTokensBaseResponse instantiates a new ListTokensBaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTokensBaseResponse(data []TokenBaseResponse) *ListTokensBaseResponse {
	this := ListTokensBaseResponse{}
	this.Data = data
	return &this
}

// NewListTokensBaseResponseWithDefaults instantiates a new ListTokensBaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTokensBaseResponseWithDefaults() *ListTokensBaseResponse {
	this := ListTokensBaseResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListTokensBaseResponse) GetData() []TokenBaseResponse {
	if o == nil {
		var ret []TokenBaseResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListTokensBaseResponse) GetDataOk() ([]TokenBaseResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListTokensBaseResponse) SetData(v []TokenBaseResponse) {
	o.Data = v
}

func (o ListTokensBaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListTokensBaseResponse struct {
	value *ListTokensBaseResponse
	isSet bool
}

func (v NullableListTokensBaseResponse) Get() *ListTokensBaseResponse {
	return v.value
}

func (v *NullableListTokensBaseResponse) Set(val *ListTokensBaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokensBaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokensBaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokensBaseResponse(val *ListTokensBaseResponse) *NullableListTokensBaseResponse {
	return &NullableListTokensBaseResponse{value: val, isSet: true}
}

func (v NullableListTokensBaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokensBaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


