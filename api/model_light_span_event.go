/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the LightSpanEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LightSpanEvent{}

// LightSpanEvent Light version of Span Event, without the attributes.
type LightSpanEvent struct {
	// Time when an event happened in the [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format.
	Timestamp time.Time `json:"timestamp"`
	// Name of the event.
	Name string `json:"name"`
}

// NewLightSpanEvent instantiates a new LightSpanEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLightSpanEvent(timestamp time.Time, name string) *LightSpanEvent {
	this := LightSpanEvent{}
	this.Timestamp = timestamp
	this.Name = name
	return &this
}

// NewLightSpanEventWithDefaults instantiates a new LightSpanEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLightSpanEventWithDefaults() *LightSpanEvent {
	this := LightSpanEvent{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *LightSpanEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *LightSpanEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *LightSpanEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetName returns the Name field value
func (o *LightSpanEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LightSpanEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LightSpanEvent) SetName(v string) {
	o.Name = v
}

func (o LightSpanEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LightSpanEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableLightSpanEvent struct {
	value *LightSpanEvent
	isSet bool
}

func (v NullableLightSpanEvent) Get() *LightSpanEvent {
	return v.value
}

func (v *NullableLightSpanEvent) Set(val *LightSpanEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableLightSpanEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableLightSpanEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLightSpanEvent(val *LightSpanEvent) *NullableLightSpanEvent {
	return &NullableLightSpanEvent{value: val, isSet: true}
}

func (v NullableLightSpanEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLightSpanEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


