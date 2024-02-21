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

// checks if the MetricsSavedSearchSyncDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsSavedSearchSyncDefinitionAllOf{}

// MetricsSavedSearchSyncDefinitionAllOf struct for MetricsSavedSearchSyncDefinitionAllOf
type MetricsSavedSearchSyncDefinitionAllOf struct {
	// Item description in the content library.
	Description *string `json:"description,omitempty"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// Query used to add an overlay to the chart.
	LogQuery *string `json:"logQuery,omitempty"`
	// Metrics queries.
	MetricsQueries []MetricsSavedSearchQuerySyncDefinition `json:"metricsQueries"`
	// Desired quantization in seconds.
	DesiredQuantizationInSecs int32 `json:"desiredQuantizationInSecs"`
	// Chart properties. This field is optional.
	Properties *string `json:"properties,omitempty"`
}

// NewMetricsSavedSearchSyncDefinitionAllOf instantiates a new MetricsSavedSearchSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsSavedSearchSyncDefinitionAllOf(timeRange ResolvableTimeRange, metricsQueries []MetricsSavedSearchQuerySyncDefinition, desiredQuantizationInSecs int32) *MetricsSavedSearchSyncDefinitionAllOf {
	this := MetricsSavedSearchSyncDefinitionAllOf{}
	this.TimeRange = timeRange
	this.MetricsQueries = metricsQueries
	this.DesiredQuantizationInSecs = desiredQuantizationInSecs
	return &this
}

// NewMetricsSavedSearchSyncDefinitionAllOfWithDefaults instantiates a new MetricsSavedSearchSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsSavedSearchSyncDefinitionAllOfWithDefaults() *MetricsSavedSearchSyncDefinitionAllOf {
	this := MetricsSavedSearchSyncDefinitionAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricsSavedSearchSyncDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetTimeRange returns the TimeRange field value
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *MetricsSavedSearchSyncDefinitionAllOf) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetLogQuery() string {
	if o == nil || IsNil(o.LogQuery) {
		var ret string
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetLogQueryOk() (*string, bool) {
	if o == nil || IsNil(o.LogQuery) {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) HasLogQuery() bool {
	if o != nil && !IsNil(o.LogQuery) {
		return true
	}

	return false
}

// SetLogQuery gets a reference to the given string and assigns it to the LogQuery field.
func (o *MetricsSavedSearchSyncDefinitionAllOf) SetLogQuery(v string) {
	o.LogQuery = &v
}

// GetMetricsQueries returns the MetricsQueries field value
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetMetricsQueries() []MetricsSavedSearchQuerySyncDefinition {
	if o == nil {
		var ret []MetricsSavedSearchQuerySyncDefinition
		return ret
	}

	return o.MetricsQueries
}

// GetMetricsQueriesOk returns a tuple with the MetricsQueries field value
// and a boolean to check if the value has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetMetricsQueriesOk() ([]MetricsSavedSearchQuerySyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricsQueries, true
}

// SetMetricsQueries sets field value
func (o *MetricsSavedSearchSyncDefinitionAllOf) SetMetricsQueries(v []MetricsSavedSearchQuerySyncDefinition) {
	o.MetricsQueries = v
}

// GetDesiredQuantizationInSecs returns the DesiredQuantizationInSecs field value
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetDesiredQuantizationInSecs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DesiredQuantizationInSecs
}

// GetDesiredQuantizationInSecsOk returns a tuple with the DesiredQuantizationInSecs field value
// and a boolean to check if the value has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetDesiredQuantizationInSecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredQuantizationInSecs, true
}

// SetDesiredQuantizationInSecs sets field value
func (o *MetricsSavedSearchSyncDefinitionAllOf) SetDesiredQuantizationInSecs(v int32) {
	o.DesiredQuantizationInSecs = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetProperties() string {
	if o == nil || IsNil(o.Properties) {
		var ret string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) GetPropertiesOk() (*string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MetricsSavedSearchSyncDefinitionAllOf) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given string and assigns it to the Properties field.
func (o *MetricsSavedSearchSyncDefinitionAllOf) SetProperties(v string) {
	o.Properties = &v
}

func (o MetricsSavedSearchSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsSavedSearchSyncDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.LogQuery) {
		toSerialize["logQuery"] = o.LogQuery
	}
	toSerialize["metricsQueries"] = o.MetricsQueries
	toSerialize["desiredQuantizationInSecs"] = o.DesiredQuantizationInSecs
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableMetricsSavedSearchSyncDefinitionAllOf struct {
	value *MetricsSavedSearchSyncDefinitionAllOf
	isSet bool
}

func (v NullableMetricsSavedSearchSyncDefinitionAllOf) Get() *MetricsSavedSearchSyncDefinitionAllOf {
	return v.value
}

func (v *NullableMetricsSavedSearchSyncDefinitionAllOf) Set(val *MetricsSavedSearchSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsSavedSearchSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsSavedSearchSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsSavedSearchSyncDefinitionAllOf(val *MetricsSavedSearchSyncDefinitionAllOf) *NullableMetricsSavedSearchSyncDefinitionAllOf {
	return &NullableMetricsSavedSearchSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableMetricsSavedSearchSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsSavedSearchSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


