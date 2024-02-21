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

// checks if the QueriesParametersResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueriesParametersResult{}

// QueriesParametersResult Queries validation and extracted parameters result.
type QueriesParametersResult struct {
	// Whether or not if queries are valid.
	IsValid *bool `json:"isValid,omitempty"`
	// Error messages from validation.
	Errors []string `json:"errors,omitempty"`
	LogsOutlier *LogsOutlier `json:"logsOutlier,omitempty"`
	MetricsOutlier *MetricsOutlier `json:"metricsOutlier,omitempty"`
}

// NewQueriesParametersResult instantiates a new QueriesParametersResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueriesParametersResult() *QueriesParametersResult {
	this := QueriesParametersResult{}
	return &this
}

// NewQueriesParametersResultWithDefaults instantiates a new QueriesParametersResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueriesParametersResultWithDefaults() *QueriesParametersResult {
	this := QueriesParametersResult{}
	return &this
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *QueriesParametersResult) GetIsValid() bool {
	if o == nil || IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesParametersResult) GetIsValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *QueriesParametersResult) HasIsValid() bool {
	if o != nil && !IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *QueriesParametersResult) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *QueriesParametersResult) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesParametersResult) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *QueriesParametersResult) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *QueriesParametersResult) SetErrors(v []string) {
	o.Errors = v
}

// GetLogsOutlier returns the LogsOutlier field value if set, zero value otherwise.
func (o *QueriesParametersResult) GetLogsOutlier() LogsOutlier {
	if o == nil || IsNil(o.LogsOutlier) {
		var ret LogsOutlier
		return ret
	}
	return *o.LogsOutlier
}

// GetLogsOutlierOk returns a tuple with the LogsOutlier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesParametersResult) GetLogsOutlierOk() (*LogsOutlier, bool) {
	if o == nil || IsNil(o.LogsOutlier) {
		return nil, false
	}
	return o.LogsOutlier, true
}

// HasLogsOutlier returns a boolean if a field has been set.
func (o *QueriesParametersResult) HasLogsOutlier() bool {
	if o != nil && !IsNil(o.LogsOutlier) {
		return true
	}

	return false
}

// SetLogsOutlier gets a reference to the given LogsOutlier and assigns it to the LogsOutlier field.
func (o *QueriesParametersResult) SetLogsOutlier(v LogsOutlier) {
	o.LogsOutlier = &v
}

// GetMetricsOutlier returns the MetricsOutlier field value if set, zero value otherwise.
func (o *QueriesParametersResult) GetMetricsOutlier() MetricsOutlier {
	if o == nil || IsNil(o.MetricsOutlier) {
		var ret MetricsOutlier
		return ret
	}
	return *o.MetricsOutlier
}

// GetMetricsOutlierOk returns a tuple with the MetricsOutlier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueriesParametersResult) GetMetricsOutlierOk() (*MetricsOutlier, bool) {
	if o == nil || IsNil(o.MetricsOutlier) {
		return nil, false
	}
	return o.MetricsOutlier, true
}

// HasMetricsOutlier returns a boolean if a field has been set.
func (o *QueriesParametersResult) HasMetricsOutlier() bool {
	if o != nil && !IsNil(o.MetricsOutlier) {
		return true
	}

	return false
}

// SetMetricsOutlier gets a reference to the given MetricsOutlier and assigns it to the MetricsOutlier field.
func (o *QueriesParametersResult) SetMetricsOutlier(v MetricsOutlier) {
	o.MetricsOutlier = &v
}

func (o QueriesParametersResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueriesParametersResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsValid) {
		toSerialize["isValid"] = o.IsValid
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.LogsOutlier) {
		toSerialize["logsOutlier"] = o.LogsOutlier
	}
	if !IsNil(o.MetricsOutlier) {
		toSerialize["metricsOutlier"] = o.MetricsOutlier
	}
	return toSerialize, nil
}

type NullableQueriesParametersResult struct {
	value *QueriesParametersResult
	isSet bool
}

func (v NullableQueriesParametersResult) Get() *QueriesParametersResult {
	return v.value
}

func (v *NullableQueriesParametersResult) Set(val *QueriesParametersResult) {
	v.value = val
	v.isSet = true
}

func (v NullableQueriesParametersResult) IsSet() bool {
	return v.isSet
}

func (v *NullableQueriesParametersResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueriesParametersResult(val *QueriesParametersResult) *NullableQueriesParametersResult {
	return &NullableQueriesParametersResult{value: val, isSet: true}
}

func (v NullableQueriesParametersResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueriesParametersResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


