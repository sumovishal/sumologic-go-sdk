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

// MetricsQueryResultInfo struct for MetricsQueryResultInfo
type MetricsQueryResultInfo struct {
	// Metrics Query row id.
	RowId *string `json:"rowId,omitempty"`
	ResultContext *MetricsQueryResultContext `json:"resultContext,omitempty"`
}

// NewMetricsQueryResultInfo instantiates a new MetricsQueryResultInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsQueryResultInfo() *MetricsQueryResultInfo {
	this := MetricsQueryResultInfo{}
	return &this
}

// NewMetricsQueryResultInfoWithDefaults instantiates a new MetricsQueryResultInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsQueryResultInfoWithDefaults() *MetricsQueryResultInfo {
	this := MetricsQueryResultInfo{}
	return &this
}

// GetRowId returns the RowId field value if set, zero value otherwise.
func (o *MetricsQueryResultInfo) GetRowId() string {
	if o == nil || o.RowId == nil {
		var ret string
		return ret
	}
	return *o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResultInfo) GetRowIdOk() (*string, bool) {
	if o == nil || o.RowId == nil {
		return nil, false
	}
	return o.RowId, true
}

// HasRowId returns a boolean if a field has been set.
func (o *MetricsQueryResultInfo) HasRowId() bool {
	if o != nil && o.RowId != nil {
		return true
	}

	return false
}

// SetRowId gets a reference to the given string and assigns it to the RowId field.
func (o *MetricsQueryResultInfo) SetRowId(v string) {
	o.RowId = &v
}

// GetResultContext returns the ResultContext field value if set, zero value otherwise.
func (o *MetricsQueryResultInfo) GetResultContext() MetricsQueryResultContext {
	if o == nil || o.ResultContext == nil {
		var ret MetricsQueryResultContext
		return ret
	}
	return *o.ResultContext
}

// GetResultContextOk returns a tuple with the ResultContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResultInfo) GetResultContextOk() (*MetricsQueryResultContext, bool) {
	if o == nil || o.ResultContext == nil {
		return nil, false
	}
	return o.ResultContext, true
}

// HasResultContext returns a boolean if a field has been set.
func (o *MetricsQueryResultInfo) HasResultContext() bool {
	if o != nil && o.ResultContext != nil {
		return true
	}

	return false
}

// SetResultContext gets a reference to the given MetricsQueryResultContext and assigns it to the ResultContext field.
func (o *MetricsQueryResultInfo) SetResultContext(v MetricsQueryResultContext) {
	o.ResultContext = &v
}

func (o MetricsQueryResultInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RowId != nil {
		toSerialize["rowId"] = o.RowId
	}
	if o.ResultContext != nil {
		toSerialize["resultContext"] = o.ResultContext
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsQueryResultInfo struct {
	value *MetricsQueryResultInfo
	isSet bool
}

func (v NullableMetricsQueryResultInfo) Get() *MetricsQueryResultInfo {
	return v.value
}

func (v *NullableMetricsQueryResultInfo) Set(val *MetricsQueryResultInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsQueryResultInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsQueryResultInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsQueryResultInfo(val *MetricsQueryResultInfo) *NullableMetricsQueryResultInfo {
	return &NullableMetricsQueryResultInfo{value: val, isSet: true}
}

func (v NullableMetricsQueryResultInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsQueryResultInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


