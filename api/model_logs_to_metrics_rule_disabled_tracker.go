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

// checks if the LogsToMetricsRuleDisabledTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogsToMetricsRuleDisabledTracker{}

// LogsToMetricsRuleDisabledTracker struct for LogsToMetricsRuleDisabledTracker
type LogsToMetricsRuleDisabledTracker struct {
	TrackerIdentity
}

type _LogsToMetricsRuleDisabledTracker LogsToMetricsRuleDisabledTracker

// NewLogsToMetricsRuleDisabledTracker instantiates a new LogsToMetricsRuleDisabledTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsToMetricsRuleDisabledTracker(trackerId string, error_ string, description string) *LogsToMetricsRuleDisabledTracker {
	this := LogsToMetricsRuleDisabledTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewLogsToMetricsRuleDisabledTrackerWithDefaults instantiates a new LogsToMetricsRuleDisabledTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsToMetricsRuleDisabledTrackerWithDefaults() *LogsToMetricsRuleDisabledTracker {
	this := LogsToMetricsRuleDisabledTracker{}
	return &this
}

func (o LogsToMetricsRuleDisabledTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogsToMetricsRuleDisabledTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	return toSerialize, nil
}

func (o *LogsToMetricsRuleDisabledTracker) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"trackerId",
		"error",
		"description",
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

	varLogsToMetricsRuleDisabledTracker := _LogsToMetricsRuleDisabledTracker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogsToMetricsRuleDisabledTracker)

	if err != nil {
		return err
	}

	*o = LogsToMetricsRuleDisabledTracker(varLogsToMetricsRuleDisabledTracker)

	return err
}

type NullableLogsToMetricsRuleDisabledTracker struct {
	value *LogsToMetricsRuleDisabledTracker
	isSet bool
}

func (v NullableLogsToMetricsRuleDisabledTracker) Get() *LogsToMetricsRuleDisabledTracker {
	return v.value
}

func (v *NullableLogsToMetricsRuleDisabledTracker) Set(val *LogsToMetricsRuleDisabledTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsToMetricsRuleDisabledTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsToMetricsRuleDisabledTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsToMetricsRuleDisabledTracker(val *LogsToMetricsRuleDisabledTracker) *NullableLogsToMetricsRuleDisabledTracker {
	return &NullableLogsToMetricsRuleDisabledTracker{value: val, isSet: true}
}

func (v NullableLogsToMetricsRuleDisabledTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsToMetricsRuleDisabledTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


