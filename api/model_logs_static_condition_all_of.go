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

// LogsStaticConditionAllOf A rule that defines how logs monitor should evaluate static data and trigger notifications.
type LogsStaticConditionAllOf struct {
	// The relative time range of the monitor. Valid values of time ranges are `5m`, `10m`, `15m`, `30m`, `1h`, `3h`, `6h`, `12h`, or `24h`.
	TimeRange string `json:"timeRange"`
	// The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for `MissingData` and `ResolvedMissingData` triggerTypes and will be ignored if specified.
	Threshold float64 `json:"threshold"`
	// The comparison type for the `threshold` evaluation. This defines how you want the data value compared. Valid values:   1. `LessThan`: Less than than the configured threshold.   2. `GreaterThan`: Greater than the configured threshold.   3. `LessThanOrEqual`: Less than or equal to the configured threshold.   4. `GreaterThanOrEqual`: Greater than or equal to the configured threshold. ThresholdType value is not applicable for `MissingData` and `ResolvedMissingData` triggerTypes and will be ignored if specified.
	ThresholdType string `json:"thresholdType"`
	// The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If `field` is not specified, monitor would default to result count instead.
	Field *string `json:"field,omitempty"`
}

// NewLogsStaticConditionAllOf instantiates a new LogsStaticConditionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsStaticConditionAllOf(timeRange string, threshold float64, thresholdType string) *LogsStaticConditionAllOf {
	this := LogsStaticConditionAllOf{}
	this.TimeRange = timeRange
	this.Threshold = threshold
	this.ThresholdType = thresholdType
	return &this
}

// NewLogsStaticConditionAllOfWithDefaults instantiates a new LogsStaticConditionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsStaticConditionAllOfWithDefaults() *LogsStaticConditionAllOf {
	this := LogsStaticConditionAllOf{}
	var threshold float64 = 0.0
	this.Threshold = threshold
	var thresholdType string = "GreaterThanOrEqual"
	this.ThresholdType = thresholdType
	return &this
}

// GetTimeRange returns the TimeRange field value
func (o *LogsStaticConditionAllOf) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *LogsStaticConditionAllOf) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *LogsStaticConditionAllOf) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetThreshold returns the Threshold field value
func (o *LogsStaticConditionAllOf) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *LogsStaticConditionAllOf) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *LogsStaticConditionAllOf) SetThreshold(v float64) {
	o.Threshold = v
}

// GetThresholdType returns the ThresholdType field value
func (o *LogsStaticConditionAllOf) GetThresholdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThresholdType
}

// GetThresholdTypeOk returns a tuple with the ThresholdType field value
// and a boolean to check if the value has been set.
func (o *LogsStaticConditionAllOf) GetThresholdTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThresholdType, true
}

// SetThresholdType sets field value
func (o *LogsStaticConditionAllOf) SetThresholdType(v string) {
	o.ThresholdType = v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *LogsStaticConditionAllOf) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStaticConditionAllOf) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *LogsStaticConditionAllOf) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *LogsStaticConditionAllOf) SetField(v string) {
	o.Field = &v
}

func (o LogsStaticConditionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timeRange"] = o.TimeRange
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["thresholdType"] = o.ThresholdType
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	return json.Marshal(toSerialize)
}

type NullableLogsStaticConditionAllOf struct {
	value *LogsStaticConditionAllOf
	isSet bool
}

func (v NullableLogsStaticConditionAllOf) Get() *LogsStaticConditionAllOf {
	return v.value
}

func (v *NullableLogsStaticConditionAllOf) Set(val *LogsStaticConditionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsStaticConditionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsStaticConditionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsStaticConditionAllOf(val *LogsStaticConditionAllOf) *NullableLogsStaticConditionAllOf {
	return &NullableLogsStaticConditionAllOf{value: val, isSet: true}
}

func (v NullableLogsStaticConditionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsStaticConditionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


