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

// CreateCpcQueryResponse struct for CreateCpcQueryResponse
type CreateCpcQueryResponse struct {
	// The id of the created query.
	QueryId string `json:"queryId"`
}

// NewCreateCpcQueryResponse instantiates a new CreateCpcQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCpcQueryResponse(queryId string) *CreateCpcQueryResponse {
	this := CreateCpcQueryResponse{}
	this.QueryId = queryId
	return &this
}

// NewCreateCpcQueryResponseWithDefaults instantiates a new CreateCpcQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCpcQueryResponseWithDefaults() *CreateCpcQueryResponse {
	this := CreateCpcQueryResponse{}
	return &this
}

// GetQueryId returns the QueryId field value
func (o *CreateCpcQueryResponse) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *CreateCpcQueryResponse) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *CreateCpcQueryResponse) SetQueryId(v string) {
	o.QueryId = v
}

func (o CreateCpcQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryId"] = o.QueryId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCpcQueryResponse struct {
	value *CreateCpcQueryResponse
	isSet bool
}

func (v NullableCreateCpcQueryResponse) Get() *CreateCpcQueryResponse {
	return v.value
}

func (v *NullableCreateCpcQueryResponse) Set(val *CreateCpcQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCpcQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCpcQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCpcQueryResponse(val *CreateCpcQueryResponse) *NullableCreateCpcQueryResponse {
	return &NullableCreateCpcQueryResponse{value: val, isSet: true}
}

func (v NullableCreateCpcQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCpcQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


