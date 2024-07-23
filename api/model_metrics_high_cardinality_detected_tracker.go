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

// checks if the MetricsHighCardinalityDetectedTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsHighCardinalityDetectedTracker{}

// MetricsHighCardinalityDetectedTracker struct for MetricsHighCardinalityDetectedTracker
type MetricsHighCardinalityDetectedTracker struct {
	TrackerIdentity
	// The retention of metrics that approached the limit.
	Retention *string `json:"retention,omitempty"`
}

type _MetricsHighCardinalityDetectedTracker MetricsHighCardinalityDetectedTracker

// NewMetricsHighCardinalityDetectedTracker instantiates a new MetricsHighCardinalityDetectedTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsHighCardinalityDetectedTracker(trackerId string, error_ string, description string) *MetricsHighCardinalityDetectedTracker {
	this := MetricsHighCardinalityDetectedTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewMetricsHighCardinalityDetectedTrackerWithDefaults instantiates a new MetricsHighCardinalityDetectedTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsHighCardinalityDetectedTrackerWithDefaults() *MetricsHighCardinalityDetectedTracker {
	this := MetricsHighCardinalityDetectedTracker{}
	return &this
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *MetricsHighCardinalityDetectedTracker) GetRetention() string {
	if o == nil || IsNil(o.Retention) {
		var ret string
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsHighCardinalityDetectedTracker) GetRetentionOk() (*string, bool) {
	if o == nil || IsNil(o.Retention) {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *MetricsHighCardinalityDetectedTracker) HasRetention() bool {
	if o != nil && !IsNil(o.Retention) {
		return true
	}

	return false
}

// SetRetention gets a reference to the given string and assigns it to the Retention field.
func (o *MetricsHighCardinalityDetectedTracker) SetRetention(v string) {
	o.Retention = &v
}

func (o MetricsHighCardinalityDetectedTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsHighCardinalityDetectedTracker) ToMap() (map[string]interface{}, error) {
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

func (o *MetricsHighCardinalityDetectedTracker) UnmarshalJSON(data []byte) (err error) {
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

	varMetricsHighCardinalityDetectedTracker := _MetricsHighCardinalityDetectedTracker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetricsHighCardinalityDetectedTracker)

	if err != nil {
		return err
	}

	*o = MetricsHighCardinalityDetectedTracker(varMetricsHighCardinalityDetectedTracker)

	return err
}

type NullableMetricsHighCardinalityDetectedTracker struct {
	value *MetricsHighCardinalityDetectedTracker
	isSet bool
}

func (v NullableMetricsHighCardinalityDetectedTracker) Get() *MetricsHighCardinalityDetectedTracker {
	return v.value
}

func (v *NullableMetricsHighCardinalityDetectedTracker) Set(val *MetricsHighCardinalityDetectedTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsHighCardinalityDetectedTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsHighCardinalityDetectedTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsHighCardinalityDetectedTracker(val *MetricsHighCardinalityDetectedTracker) *NullableMetricsHighCardinalityDetectedTracker {
	return &NullableMetricsHighCardinalityDetectedTracker{value: val, isSet: true}
}

func (v NullableMetricsHighCardinalityDetectedTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsHighCardinalityDetectedTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


