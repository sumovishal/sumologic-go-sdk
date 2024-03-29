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

// checks if the OTCollectorLimitApproachingTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCollectorLimitApproachingTracker{}

// OTCollectorLimitApproachingTracker struct for OTCollectorLimitApproachingTracker
type OTCollectorLimitApproachingTracker struct {
	TrackerIdentity
}

// NewOTCollectorLimitApproachingTracker instantiates a new OTCollectorLimitApproachingTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCollectorLimitApproachingTracker(trackerId string, error_ string, description string) *OTCollectorLimitApproachingTracker {
	this := OTCollectorLimitApproachingTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewOTCollectorLimitApproachingTrackerWithDefaults instantiates a new OTCollectorLimitApproachingTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCollectorLimitApproachingTrackerWithDefaults() *OTCollectorLimitApproachingTracker {
	this := OTCollectorLimitApproachingTracker{}
	return &this
}

func (o OTCollectorLimitApproachingTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCollectorLimitApproachingTracker) ToMap() (map[string]interface{}, error) {
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

type NullableOTCollectorLimitApproachingTracker struct {
	value *OTCollectorLimitApproachingTracker
	isSet bool
}

func (v NullableOTCollectorLimitApproachingTracker) Get() *OTCollectorLimitApproachingTracker {
	return v.value
}

func (v *NullableOTCollectorLimitApproachingTracker) Set(val *OTCollectorLimitApproachingTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCollectorLimitApproachingTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCollectorLimitApproachingTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCollectorLimitApproachingTracker(val *OTCollectorLimitApproachingTracker) *NullableOTCollectorLimitApproachingTracker {
	return &NullableOTCollectorLimitApproachingTracker{value: val, isSet: true}
}

func (v NullableOTCollectorLimitApproachingTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCollectorLimitApproachingTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


