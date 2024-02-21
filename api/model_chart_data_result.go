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

// checks if the ChartDataResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChartDataResult{}

// ChartDataResult Response for monitor chart data visualization.
type ChartDataResult struct {
	// Execution warnings of queries.
	Warnings []ErrorDescription `json:"warnings,omitempty"`
	// List of time series of the monitor chart data.
	Series []SeriesData `json:"series,omitempty"`
}

// NewChartDataResult instantiates a new ChartDataResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartDataResult() *ChartDataResult {
	this := ChartDataResult{}
	return &this
}

// NewChartDataResultWithDefaults instantiates a new ChartDataResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartDataResultWithDefaults() *ChartDataResult {
	this := ChartDataResult{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChartDataResult) GetWarnings() []ErrorDescription {
	if o == nil || IsNil(o.Warnings) {
		var ret []ErrorDescription
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartDataResult) GetWarningsOk() ([]ErrorDescription, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChartDataResult) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []ErrorDescription and assigns it to the Warnings field.
func (o *ChartDataResult) SetWarnings(v []ErrorDescription) {
	o.Warnings = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *ChartDataResult) GetSeries() []SeriesData {
	if o == nil || IsNil(o.Series) {
		var ret []SeriesData
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartDataResult) GetSeriesOk() ([]SeriesData, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *ChartDataResult) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []SeriesData and assigns it to the Series field.
func (o *ChartDataResult) SetSeries(v []SeriesData) {
	o.Series = v
}

func (o ChartDataResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChartDataResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	return toSerialize, nil
}

type NullableChartDataResult struct {
	value *ChartDataResult
	isSet bool
}

func (v NullableChartDataResult) Get() *ChartDataResult {
	return v.value
}

func (v *NullableChartDataResult) Set(val *ChartDataResult) {
	v.value = val
	v.isSet = true
}

func (v NullableChartDataResult) IsSet() bool {
	return v.isSet
}

func (v *NullableChartDataResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartDataResult(val *ChartDataResult) *NullableChartDataResult {
	return &NullableChartDataResult{value: val, isSet: true}
}

func (v NullableChartDataResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartDataResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


