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

// checks if the MetricsCardinalityLimitExceededTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsCardinalityLimitExceededTracker{}

// MetricsCardinalityLimitExceededTracker struct for MetricsCardinalityLimitExceededTracker
type MetricsCardinalityLimitExceededTracker struct {
	TrackerIdentity
	// The retention of metrics that exceeded the limit.
	Retention *string `json:"retention,omitempty"`
}

// NewMetricsCardinalityLimitExceededTracker instantiates a new MetricsCardinalityLimitExceededTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsCardinalityLimitExceededTracker(trackerId string, error_ string, description string) *MetricsCardinalityLimitExceededTracker {
	this := MetricsCardinalityLimitExceededTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewMetricsCardinalityLimitExceededTrackerWithDefaults instantiates a new MetricsCardinalityLimitExceededTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsCardinalityLimitExceededTrackerWithDefaults() *MetricsCardinalityLimitExceededTracker {
	this := MetricsCardinalityLimitExceededTracker{}
	return &this
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *MetricsCardinalityLimitExceededTracker) GetRetention() string {
	if o == nil || IsNil(o.Retention) {
		var ret string
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsCardinalityLimitExceededTracker) GetRetentionOk() (*string, bool) {
	if o == nil || IsNil(o.Retention) {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *MetricsCardinalityLimitExceededTracker) HasRetention() bool {
	if o != nil && !IsNil(o.Retention) {
		return true
	}

	return false
}

// SetRetention gets a reference to the given string and assigns it to the Retention field.
func (o *MetricsCardinalityLimitExceededTracker) SetRetention(v string) {
	o.Retention = &v
}

func (o MetricsCardinalityLimitExceededTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsCardinalityLimitExceededTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	if !IsNil(o.Retention) {
		toSerialize["retention"] = o.Retention
	}
	return toSerialize, nil
}

type NullableMetricsCardinalityLimitExceededTracker struct {
	value *MetricsCardinalityLimitExceededTracker
	isSet bool
}

func (v NullableMetricsCardinalityLimitExceededTracker) Get() *MetricsCardinalityLimitExceededTracker {
	return v.value
}

func (v *NullableMetricsCardinalityLimitExceededTracker) Set(val *MetricsCardinalityLimitExceededTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsCardinalityLimitExceededTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsCardinalityLimitExceededTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsCardinalityLimitExceededTracker(val *MetricsCardinalityLimitExceededTracker) *NullableMetricsCardinalityLimitExceededTracker {
	return &NullableMetricsCardinalityLimitExceededTracker{value: val, isSet: true}
}

func (v NullableMetricsCardinalityLimitExceededTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsCardinalityLimitExceededTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


