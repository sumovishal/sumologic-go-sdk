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

// TraceFieldsResponse struct for TraceFieldsResponse
type TraceFieldsResponse struct {
	// List of filter fields.
	Fields []TraceFieldDetail `json:"fields"`
}

// NewTraceFieldsResponse instantiates a new TraceFieldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceFieldsResponse(fields []TraceFieldDetail) *TraceFieldsResponse {
	this := TraceFieldsResponse{}
	this.Fields = fields
	return &this
}

// NewTraceFieldsResponseWithDefaults instantiates a new TraceFieldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceFieldsResponseWithDefaults() *TraceFieldsResponse {
	this := TraceFieldsResponse{}
	return &this
}

// GetFields returns the Fields field value
func (o *TraceFieldsResponse) GetFields() []TraceFieldDetail {
	if o == nil {
		var ret []TraceFieldDetail
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *TraceFieldsResponse) GetFieldsOk() ([]TraceFieldDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *TraceFieldsResponse) SetFields(v []TraceFieldDetail) {
	o.Fields = v
}

func (o TraceFieldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableTraceFieldsResponse struct {
	value *TraceFieldsResponse
	isSet bool
}

func (v NullableTraceFieldsResponse) Get() *TraceFieldsResponse {
	return v.value
}

func (v *NullableTraceFieldsResponse) Set(val *TraceFieldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceFieldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceFieldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceFieldsResponse(val *TraceFieldsResponse) *NullableTraceFieldsResponse {
	return &NullableTraceFieldsResponse{value: val, isSet: true}
}

func (v NullableTraceFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceFieldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


