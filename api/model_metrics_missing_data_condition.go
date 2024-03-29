/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// checks if the MetricsMissingDataCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsMissingDataCondition{}

// MetricsMissingDataCondition struct for MetricsMissingDataCondition
type MetricsMissingDataCondition struct {
	TriggerCondition
	// Determines which time series from queries to use for Metrics MissingData and ResolvedMissingData triggers Valid values:   1. `AllTimeSeries`: Evaluate the condition against all time series. (NOTE: This option is only valid if monitorType is `Metrics`)   2. `AnyTimeSeries`: Evaluate the condition against any time series. (NOTE: This option is only valid if monitorType is `Metrics`)   3. `AllResults`: Evaluate the condition against results from all queries. (NOTE: This option is only valid if monitorType is `Logs`)
	TriggerSource string `json:"triggerSource"`
	// The relative time range of the monitor. Valid values of time ranges are `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, or `-24h`.
	TimeRange string `json:"timeRange"`
}

// NewMetricsMissingDataCondition instantiates a new MetricsMissingDataCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsMissingDataCondition(triggerSource string, timeRange string, triggerType string) *MetricsMissingDataCondition {
	this := MetricsMissingDataCondition{}
	var detectionMethod string = "StaticCondition"
	this.DetectionMethod = &detectionMethod
	this.TriggerType = triggerType
	this.TriggerSource = triggerSource
	this.TimeRange = timeRange
	return &this
}

// NewMetricsMissingDataConditionWithDefaults instantiates a new MetricsMissingDataCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsMissingDataConditionWithDefaults() *MetricsMissingDataCondition {
	this := MetricsMissingDataCondition{}
	return &this
}

// GetTriggerSource returns the TriggerSource field value
func (o *MetricsMissingDataCondition) GetTriggerSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerSource
}

// GetTriggerSourceOk returns a tuple with the TriggerSource field value
// and a boolean to check if the value has been set.
func (o *MetricsMissingDataCondition) GetTriggerSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerSource, true
}

// SetTriggerSource sets field value
func (o *MetricsMissingDataCondition) SetTriggerSource(v string) {
	o.TriggerSource = v
}

// GetTimeRange returns the TimeRange field value
func (o *MetricsMissingDataCondition) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *MetricsMissingDataCondition) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *MetricsMissingDataCondition) SetTimeRange(v string) {
	o.TimeRange = v
}

func (o MetricsMissingDataCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsMissingDataCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTriggerCondition, errTriggerCondition := json.Marshal(o.TriggerCondition)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	errTriggerCondition = json.Unmarshal([]byte(serializedTriggerCondition), &toSerialize)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	toSerialize["triggerSource"] = o.TriggerSource
	toSerialize["timeRange"] = o.TimeRange
	return toSerialize, nil
}

type NullableMetricsMissingDataCondition struct {
	value *MetricsMissingDataCondition
	isSet bool
}

func (v NullableMetricsMissingDataCondition) Get() *MetricsMissingDataCondition {
	return v.value
}

func (v *NullableMetricsMissingDataCondition) Set(val *MetricsMissingDataCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsMissingDataCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsMissingDataCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsMissingDataCondition(val *MetricsMissingDataCondition) *NullableMetricsMissingDataCondition {
	return &NullableMetricsMissingDataCondition{value: val, isSet: true}
}

func (v NullableMetricsMissingDataCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsMissingDataCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


