/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MetricsQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsQueryResponse{}

// MetricsQueryResponse struct for MetricsQueryResponse
type MetricsQueryResponse struct {
	// A list of the time series returned by metric query.
	QueryResult []TimeSeriesRow `json:"queryResult,omitempty"`
	// Errors, warnings, and information logged for the query.
	Errors ErrorResponse `json:"errors"`
}

type _MetricsQueryResponse MetricsQueryResponse

// NewMetricsQueryResponse instantiates a new MetricsQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsQueryResponse(errors ErrorResponse) *MetricsQueryResponse {
	this := MetricsQueryResponse{}
	this.Errors = errors
	return &this
}

// NewMetricsQueryResponseWithDefaults instantiates a new MetricsQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsQueryResponseWithDefaults() *MetricsQueryResponse {
	this := MetricsQueryResponse{}
	return &this
}

// GetQueryResult returns the QueryResult field value if set, zero value otherwise.
func (o *MetricsQueryResponse) GetQueryResult() []TimeSeriesRow {
	if o == nil || IsNil(o.QueryResult) {
		var ret []TimeSeriesRow
		return ret
	}
	return o.QueryResult
}

// GetQueryResultOk returns a tuple with the QueryResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponse) GetQueryResultOk() ([]TimeSeriesRow, bool) {
	if o == nil || IsNil(o.QueryResult) {
		return nil, false
	}
	return o.QueryResult, true
}

// HasQueryResult returns a boolean if a field has been set.
func (o *MetricsQueryResponse) HasQueryResult() bool {
	if o != nil && !IsNil(o.QueryResult) {
		return true
	}

	return false
}

// SetQueryResult gets a reference to the given []TimeSeriesRow and assigns it to the QueryResult field.
func (o *MetricsQueryResponse) SetQueryResult(v []TimeSeriesRow) {
	o.QueryResult = v
}

// GetErrors returns the Errors field value
func (o *MetricsQueryResponse) GetErrors() ErrorResponse {
	if o == nil {
		var ret ErrorResponse
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponse) GetErrorsOk() (*ErrorResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *MetricsQueryResponse) SetErrors(v ErrorResponse) {
	o.Errors = v
}

func (o MetricsQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueryResult) {
		toSerialize["queryResult"] = o.QueryResult
	}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

func (o *MetricsQueryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"errors",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMetricsQueryResponse := _MetricsQueryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetricsQueryResponse)

	if err != nil {
		return err
	}

	*o = MetricsQueryResponse(varMetricsQueryResponse)

	return err
}

type NullableMetricsQueryResponse struct {
	value *MetricsQueryResponse
	isSet bool
}

func (v NullableMetricsQueryResponse) Get() *MetricsQueryResponse {
	return v.value
}

func (v *NullableMetricsQueryResponse) Set(val *MetricsQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsQueryResponse(val *MetricsQueryResponse) *NullableMetricsQueryResponse {
	return &NullableMetricsQueryResponse{value: val, isSet: true}
}

func (v NullableMetricsQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


