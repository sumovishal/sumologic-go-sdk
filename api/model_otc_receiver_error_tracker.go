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

// checks if the OTCReceiverErrorTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCReceiverErrorTracker{}

// OTCReceiverErrorTracker struct for OTCReceiverErrorTracker
type OTCReceiverErrorTracker struct {
	// Event type.
	EventType *string `json:"eventType,omitempty"`
}

// NewOTCReceiverErrorTracker instantiates a new OTCReceiverErrorTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCReceiverErrorTracker() *OTCReceiverErrorTracker {
	this := OTCReceiverErrorTracker{}
	return &this
}

// NewOTCReceiverErrorTrackerWithDefaults instantiates a new OTCReceiverErrorTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCReceiverErrorTrackerWithDefaults() *OTCReceiverErrorTracker {
	this := OTCReceiverErrorTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *OTCReceiverErrorTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverErrorTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *OTCReceiverErrorTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *OTCReceiverErrorTracker) SetEventType(v string) {
	o.EventType = &v
}

func (o OTCReceiverErrorTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCReceiverErrorTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return toSerialize, nil
}

type NullableOTCReceiverErrorTracker struct {
	value *OTCReceiverErrorTracker
	isSet bool
}

func (v NullableOTCReceiverErrorTracker) Get() *OTCReceiverErrorTracker {
	return v.value
}

func (v *NullableOTCReceiverErrorTracker) Set(val *OTCReceiverErrorTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCReceiverErrorTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCReceiverErrorTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCReceiverErrorTracker(val *OTCReceiverErrorTracker) *NullableOTCReceiverErrorTracker {
	return &NullableOTCReceiverErrorTracker{value: val, isSet: true}
}

func (v NullableOTCReceiverErrorTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCReceiverErrorTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


