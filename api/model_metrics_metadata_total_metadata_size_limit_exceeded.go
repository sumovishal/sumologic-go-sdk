/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// MetricsMetadataTotalMetadataSizeLimitExceeded struct for MetricsMetadataTotalMetadataSizeLimitExceeded
type MetricsMetadataTotalMetadataSizeLimitExceeded struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
}

// NewMetricsMetadataTotalMetadataSizeLimitExceeded instantiates a new MetricsMetadataTotalMetadataSizeLimitExceeded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsMetadataTotalMetadataSizeLimitExceeded(trackerId string, error_ string, description string) *MetricsMetadataTotalMetadataSizeLimitExceeded {
	this := MetricsMetadataTotalMetadataSizeLimitExceeded{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewMetricsMetadataTotalMetadataSizeLimitExceededWithDefaults instantiates a new MetricsMetadataTotalMetadataSizeLimitExceeded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsMetadataTotalMetadataSizeLimitExceededWithDefaults() *MetricsMetadataTotalMetadataSizeLimitExceeded {
	this := MetricsMetadataTotalMetadataSizeLimitExceeded{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MetricsMetadataTotalMetadataSizeLimitExceeded) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetadataTotalMetadataSizeLimitExceeded) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MetricsMetadataTotalMetadataSizeLimitExceeded) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MetricsMetadataTotalMetadataSizeLimitExceeded) SetEventType(v string) {
	o.EventType = &v
}

func (o MetricsMetadataTotalMetadataSizeLimitExceeded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return []byte{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return []byte{}, errTrackerIdentity
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsMetadataTotalMetadataSizeLimitExceeded struct {
	value *MetricsMetadataTotalMetadataSizeLimitExceeded
	isSet bool
}

func (v NullableMetricsMetadataTotalMetadataSizeLimitExceeded) Get() *MetricsMetadataTotalMetadataSizeLimitExceeded {
	return v.value
}

func (v *NullableMetricsMetadataTotalMetadataSizeLimitExceeded) Set(val *MetricsMetadataTotalMetadataSizeLimitExceeded) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsMetadataTotalMetadataSizeLimitExceeded) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsMetadataTotalMetadataSizeLimitExceeded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsMetadataTotalMetadataSizeLimitExceeded(val *MetricsMetadataTotalMetadataSizeLimitExceeded) *NullableMetricsMetadataTotalMetadataSizeLimitExceeded {
	return &NullableMetricsMetadataTotalMetadataSizeLimitExceeded{value: val, isSet: true}
}

func (v NullableMetricsMetadataTotalMetadataSizeLimitExceeded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsMetadataTotalMetadataSizeLimitExceeded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


