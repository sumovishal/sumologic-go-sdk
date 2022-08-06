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

// BulkErrorResponse struct for BulkErrorResponse
type BulkErrorResponse struct {
	// HTTP status code of individual request
	Status int32 `json:"status"`
	ErrorResponse ErrorResponse `json:"errorResponse"`
}

// NewBulkErrorResponse instantiates a new BulkErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkErrorResponse(status int32, errorResponse ErrorResponse) *BulkErrorResponse {
	this := BulkErrorResponse{}
	this.Status = status
	this.ErrorResponse = errorResponse
	return &this
}

// NewBulkErrorResponseWithDefaults instantiates a new BulkErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkErrorResponseWithDefaults() *BulkErrorResponse {
	this := BulkErrorResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *BulkErrorResponse) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkErrorResponse) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkErrorResponse) SetStatus(v int32) {
	o.Status = v
}

// GetErrorResponse returns the ErrorResponse field value
func (o *BulkErrorResponse) GetErrorResponse() ErrorResponse {
	if o == nil {
		var ret ErrorResponse
		return ret
	}

	return o.ErrorResponse
}

// GetErrorResponseOk returns a tuple with the ErrorResponse field value
// and a boolean to check if the value has been set.
func (o *BulkErrorResponse) GetErrorResponseOk() (*ErrorResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorResponse, true
}

// SetErrorResponse sets field value
func (o *BulkErrorResponse) SetErrorResponse(v ErrorResponse) {
	o.ErrorResponse = v
}

func (o BulkErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["errorResponse"] = o.ErrorResponse
	}
	return json.Marshal(toSerialize)
}

type NullableBulkErrorResponse struct {
	value *BulkErrorResponse
	isSet bool
}

func (v NullableBulkErrorResponse) Get() *BulkErrorResponse {
	return v.value
}

func (v *NullableBulkErrorResponse) Set(val *BulkErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkErrorResponse(val *BulkErrorResponse) *NullableBulkErrorResponse {
	return &NullableBulkErrorResponse{value: val, isSet: true}
}

func (v NullableBulkErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


