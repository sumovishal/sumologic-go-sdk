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

// checks if the MetricsMissingDataConditionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsMissingDataConditionAllOf{}

// MetricsMissingDataConditionAllOf A rule that defines how metrics monitors should evaluate missing data and trigger notifications.
type MetricsMissingDataConditionAllOf struct {
	// Determines which time series from queries to use for Metrics MissingData and ResolvedMissingData triggers Valid values:   1. `AllTimeSeries`: Evaluate the condition against all time series. (NOTE: This option is only valid if monitorType is `Metrics`)   2. `AnyTimeSeries`: Evaluate the condition against any time series. (NOTE: This option is only valid if monitorType is `Metrics`)   3. `AllResults`: Evaluate the condition against results from all queries. (NOTE: This option is only valid if monitorType is `Logs`)
	TriggerSource string `json:"triggerSource"`
	// The relative time range of the monitor. Valid values of time ranges are `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, or `-24h`.
	TimeRange string `json:"timeRange"`
}

// NewMetricsMissingDataConditionAllOf instantiates a new MetricsMissingDataConditionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsMissingDataConditionAllOf(triggerSource string, timeRange string) *MetricsMissingDataConditionAllOf {
	this := MetricsMissingDataConditionAllOf{}
	this.TriggerSource = triggerSource
	this.TimeRange = timeRange
	return &this
}

// NewMetricsMissingDataConditionAllOfWithDefaults instantiates a new MetricsMissingDataConditionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsMissingDataConditionAllOfWithDefaults() *MetricsMissingDataConditionAllOf {
	this := MetricsMissingDataConditionAllOf{}
	return &this
}

// GetTriggerSource returns the TriggerSource field value
func (o *MetricsMissingDataConditionAllOf) GetTriggerSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerSource
}

// GetTriggerSourceOk returns a tuple with the TriggerSource field value
// and a boolean to check if the value has been set.
func (o *MetricsMissingDataConditionAllOf) GetTriggerSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerSource, true
}

// SetTriggerSource sets field value
func (o *MetricsMissingDataConditionAllOf) SetTriggerSource(v string) {
	o.TriggerSource = v
}

// GetTimeRange returns the TimeRange field value
func (o *MetricsMissingDataConditionAllOf) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *MetricsMissingDataConditionAllOf) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *MetricsMissingDataConditionAllOf) SetTimeRange(v string) {
	o.TimeRange = v
}

func (o MetricsMissingDataConditionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsMissingDataConditionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["triggerSource"] = o.TriggerSource
	toSerialize["timeRange"] = o.TimeRange
	return toSerialize, nil
}

type NullableMetricsMissingDataConditionAllOf struct {
	value *MetricsMissingDataConditionAllOf
	isSet bool
}

func (v NullableMetricsMissingDataConditionAllOf) Get() *MetricsMissingDataConditionAllOf {
	return v.value
}

func (v *NullableMetricsMissingDataConditionAllOf) Set(val *MetricsMissingDataConditionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsMissingDataConditionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsMissingDataConditionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsMissingDataConditionAllOf(val *MetricsMissingDataConditionAllOf) *NullableMetricsMissingDataConditionAllOf {
	return &NullableMetricsMissingDataConditionAllOf{value: val, isSet: true}
}

func (v NullableMetricsMissingDataConditionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsMissingDataConditionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


