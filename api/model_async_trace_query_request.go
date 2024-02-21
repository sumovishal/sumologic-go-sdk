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

// checks if the AsyncTraceQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncTraceQueryRequest{}

// AsyncTraceQueryRequest struct for AsyncTraceQueryRequest
type AsyncTraceQueryRequest struct {
	// A list of trace queries.
	QueryRows []AsyncTraceQueryRow `json:"queryRows"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
}

// NewAsyncTraceQueryRequest instantiates a new AsyncTraceQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncTraceQueryRequest(queryRows []AsyncTraceQueryRow, timeRange ResolvableTimeRange) *AsyncTraceQueryRequest {
	this := AsyncTraceQueryRequest{}
	this.QueryRows = queryRows
	this.TimeRange = timeRange
	return &this
}

// NewAsyncTraceQueryRequestWithDefaults instantiates a new AsyncTraceQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncTraceQueryRequestWithDefaults() *AsyncTraceQueryRequest {
	this := AsyncTraceQueryRequest{}
	return &this
}

// GetQueryRows returns the QueryRows field value
func (o *AsyncTraceQueryRequest) GetQueryRows() []AsyncTraceQueryRow {
	if o == nil {
		var ret []AsyncTraceQueryRow
		return ret
	}

	return o.QueryRows
}

// GetQueryRowsOk returns a tuple with the QueryRows field value
// and a boolean to check if the value has been set.
func (o *AsyncTraceQueryRequest) GetQueryRowsOk() ([]AsyncTraceQueryRow, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryRows, true
}

// SetQueryRows sets field value
func (o *AsyncTraceQueryRequest) SetQueryRows(v []AsyncTraceQueryRow) {
	o.QueryRows = v
}

// GetTimeRange returns the TimeRange field value
func (o *AsyncTraceQueryRequest) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *AsyncTraceQueryRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *AsyncTraceQueryRequest) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

func (o AsyncTraceQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncTraceQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryRows"] = o.QueryRows
	toSerialize["timeRange"] = o.TimeRange
	return toSerialize, nil
}

type NullableAsyncTraceQueryRequest struct {
	value *AsyncTraceQueryRequest
	isSet bool
}

func (v NullableAsyncTraceQueryRequest) Get() *AsyncTraceQueryRequest {
	return v.value
}

func (v *NullableAsyncTraceQueryRequest) Set(val *AsyncTraceQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncTraceQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncTraceQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncTraceQueryRequest(val *AsyncTraceQueryRequest) *NullableAsyncTraceQueryRequest {
	return &NullableAsyncTraceQueryRequest{value: val, isSet: true}
}

func (v NullableAsyncTraceQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncTraceQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


