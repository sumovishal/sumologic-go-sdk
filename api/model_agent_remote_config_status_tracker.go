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

// checks if the AgentRemoteConfigStatusTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRemoteConfigStatusTracker{}

// AgentRemoteConfigStatusTracker struct for AgentRemoteConfigStatusTracker
type AgentRemoteConfigStatusTracker struct {
	TrackerIdentity
}

// NewAgentRemoteConfigStatusTracker instantiates a new AgentRemoteConfigStatusTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRemoteConfigStatusTracker(trackerId string, error_ string, description string) *AgentRemoteConfigStatusTracker {
	this := AgentRemoteConfigStatusTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewAgentRemoteConfigStatusTrackerWithDefaults instantiates a new AgentRemoteConfigStatusTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRemoteConfigStatusTrackerWithDefaults() *AgentRemoteConfigStatusTracker {
	this := AgentRemoteConfigStatusTracker{}
	return &this
}

func (o AgentRemoteConfigStatusTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRemoteConfigStatusTracker) ToMap() (map[string]interface{}, error) {
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

type NullableAgentRemoteConfigStatusTracker struct {
	value *AgentRemoteConfigStatusTracker
	isSet bool
}

func (v NullableAgentRemoteConfigStatusTracker) Get() *AgentRemoteConfigStatusTracker {
	return v.value
}

func (v *NullableAgentRemoteConfigStatusTracker) Set(val *AgentRemoteConfigStatusTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRemoteConfigStatusTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRemoteConfigStatusTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRemoteConfigStatusTracker(val *AgentRemoteConfigStatusTracker) *NullableAgentRemoteConfigStatusTracker {
	return &NullableAgentRemoteConfigStatusTracker{value: val, isSet: true}
}

func (v NullableAgentRemoteConfigStatusTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRemoteConfigStatusTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


