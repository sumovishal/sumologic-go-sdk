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

// checks if the MetricsMetadataTotalMetadataSizeLimitExceededTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsMetadataTotalMetadataSizeLimitExceededTracker{}

// MetricsMetadataTotalMetadataSizeLimitExceededTracker struct for MetricsMetadataTotalMetadataSizeLimitExceededTracker
type MetricsMetadataTotalMetadataSizeLimitExceededTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
}

// NewMetricsMetadataTotalMetadataSizeLimitExceededTracker instantiates a new MetricsMetadataTotalMetadataSizeLimitExceededTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsMetadataTotalMetadataSizeLimitExceededTracker(trackerId string, error_ string, description string) *MetricsMetadataTotalMetadataSizeLimitExceededTracker {
	this := MetricsMetadataTotalMetadataSizeLimitExceededTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewMetricsMetadataTotalMetadataSizeLimitExceededTrackerWithDefaults instantiates a new MetricsMetadataTotalMetadataSizeLimitExceededTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsMetadataTotalMetadataSizeLimitExceededTrackerWithDefaults() *MetricsMetadataTotalMetadataSizeLimitExceededTracker {
	this := MetricsMetadataTotalMetadataSizeLimitExceededTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MetricsMetadataTotalMetadataSizeLimitExceededTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetadataTotalMetadataSizeLimitExceededTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MetricsMetadataTotalMetadataSizeLimitExceededTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MetricsMetadataTotalMetadataSizeLimitExceededTracker) SetEventType(v string) {
	o.EventType = &v
}

func (o MetricsMetadataTotalMetadataSizeLimitExceededTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsMetadataTotalMetadataSizeLimitExceededTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return toSerialize, nil
}

type NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker struct {
	value *MetricsMetadataTotalMetadataSizeLimitExceededTracker
	isSet bool
}

func (v NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker) Get() *MetricsMetadataTotalMetadataSizeLimitExceededTracker {
	return v.value
}

func (v *NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker) Set(val *MetricsMetadataTotalMetadataSizeLimitExceededTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsMetadataTotalMetadataSizeLimitExceededTracker(val *MetricsMetadataTotalMetadataSizeLimitExceededTracker) *NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker {
	return &NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker{value: val, isSet: true}
}

func (v NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsMetadataTotalMetadataSizeLimitExceededTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


