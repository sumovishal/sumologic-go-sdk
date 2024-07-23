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

// checks if the CollectorLimitApproachingTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectorLimitApproachingTracker{}

// CollectorLimitApproachingTracker struct for CollectorLimitApproachingTracker
type CollectorLimitApproachingTracker struct {
	TrackerIdentity
}

type _CollectorLimitApproachingTracker CollectorLimitApproachingTracker

// NewCollectorLimitApproachingTracker instantiates a new CollectorLimitApproachingTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectorLimitApproachingTracker(trackerId string, error_ string, description string) *CollectorLimitApproachingTracker {
	this := CollectorLimitApproachingTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCollectorLimitApproachingTrackerWithDefaults instantiates a new CollectorLimitApproachingTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectorLimitApproachingTrackerWithDefaults() *CollectorLimitApproachingTracker {
	this := CollectorLimitApproachingTracker{}
	return &this
}

func (o CollectorLimitApproachingTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectorLimitApproachingTracker) ToMap() (map[string]interface{}, error) {
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

func (o *CollectorLimitApproachingTracker) UnmarshalJSON(data []byte) (err error) {
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

	varCollectorLimitApproachingTracker := _CollectorLimitApproachingTracker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectorLimitApproachingTracker)

	if err != nil {
		return err
	}

	*o = CollectorLimitApproachingTracker(varCollectorLimitApproachingTracker)

	return err
}

type NullableCollectorLimitApproachingTracker struct {
	value *CollectorLimitApproachingTracker
	isSet bool
}

func (v NullableCollectorLimitApproachingTracker) Get() *CollectorLimitApproachingTracker {
	return v.value
}

func (v *NullableCollectorLimitApproachingTracker) Set(val *CollectorLimitApproachingTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectorLimitApproachingTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectorLimitApproachingTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectorLimitApproachingTracker(val *CollectorLimitApproachingTracker) *NullableCollectorLimitApproachingTracker {
	return &NullableCollectorLimitApproachingTracker{value: val, isSet: true}
}

func (v NullableCollectorLimitApproachingTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectorLimitApproachingTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


