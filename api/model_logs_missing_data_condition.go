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

// checks if the LogsMissingDataCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogsMissingDataCondition{}

// LogsMissingDataCondition struct for LogsMissingDataCondition
type LogsMissingDataCondition struct {
	TriggerCondition
	// The relative time range of the monitor. Valid values of time ranges are `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, or `-24h`.
	TimeRange string `json:"timeRange"`
}

// NewLogsMissingDataCondition instantiates a new LogsMissingDataCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsMissingDataCondition(timeRange string, triggerType string) *LogsMissingDataCondition {
	this := LogsMissingDataCondition{}
	var detectionMethod string = "StaticCondition"
	this.DetectionMethod = &detectionMethod
	this.TriggerType = triggerType
	this.TimeRange = timeRange
	return &this
}

// NewLogsMissingDataConditionWithDefaults instantiates a new LogsMissingDataCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsMissingDataConditionWithDefaults() *LogsMissingDataCondition {
	this := LogsMissingDataCondition{}
	return &this
}

// GetTimeRange returns the TimeRange field value
func (o *LogsMissingDataCondition) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *LogsMissingDataCondition) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *LogsMissingDataCondition) SetTimeRange(v string) {
	o.TimeRange = v
}

func (o LogsMissingDataCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogsMissingDataCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTriggerCondition, errTriggerCondition := json.Marshal(o.TriggerCondition)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	errTriggerCondition = json.Unmarshal([]byte(serializedTriggerCondition), &toSerialize)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	toSerialize["timeRange"] = o.TimeRange
	return toSerialize, nil
}

type NullableLogsMissingDataCondition struct {
	value *LogsMissingDataCondition
	isSet bool
}

func (v NullableLogsMissingDataCondition) Get() *LogsMissingDataCondition {
	return v.value
}

func (v *NullableLogsMissingDataCondition) Set(val *LogsMissingDataCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsMissingDataCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsMissingDataCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsMissingDataCondition(val *LogsMissingDataCondition) *NullableLogsMissingDataCondition {
	return &NullableLogsMissingDataCondition{value: val, isSet: true}
}

func (v NullableLogsMissingDataCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsMissingDataCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


