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

// checks if the MetricNameErrorTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricNameErrorTracker{}

// MetricNameErrorTracker struct for MetricNameErrorTracker
type MetricNameErrorTracker struct {
	// Event type.
	EventType *string `json:"eventType,omitempty"`
}

// NewMetricNameErrorTracker instantiates a new MetricNameErrorTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricNameErrorTracker() *MetricNameErrorTracker {
	this := MetricNameErrorTracker{}
	return &this
}

// NewMetricNameErrorTrackerWithDefaults instantiates a new MetricNameErrorTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricNameErrorTrackerWithDefaults() *MetricNameErrorTracker {
	this := MetricNameErrorTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MetricNameErrorTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricNameErrorTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MetricNameErrorTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MetricNameErrorTracker) SetEventType(v string) {
	o.EventType = &v
}

func (o MetricNameErrorTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricNameErrorTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return toSerialize, nil
}

type NullableMetricNameErrorTracker struct {
	value *MetricNameErrorTracker
	isSet bool
}

func (v NullableMetricNameErrorTracker) Get() *MetricNameErrorTracker {
	return v.value
}

func (v *NullableMetricNameErrorTracker) Set(val *MetricNameErrorTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricNameErrorTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricNameErrorTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricNameErrorTracker(val *MetricNameErrorTracker) *NullableMetricNameErrorTracker {
	return &NullableMetricNameErrorTracker{value: val, isSet: true}
}

func (v NullableMetricNameErrorTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricNameErrorTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


