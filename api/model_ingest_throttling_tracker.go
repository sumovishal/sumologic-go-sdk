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

// IngestThrottlingTracker struct for IngestThrottlingTracker
type IngestThrottlingTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The type of data for which the rate limit was enabled. The possible values are `LogIngest` and `MetricsIngest`.
	DataType *string `json:"dataType,omitempty"`
}

// NewIngestThrottlingTracker instantiates a new IngestThrottlingTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestThrottlingTracker(trackerId string, error_ string, description string) *IngestThrottlingTracker {
	this := IngestThrottlingTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewIngestThrottlingTrackerWithDefaults instantiates a new IngestThrottlingTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestThrottlingTrackerWithDefaults() *IngestThrottlingTracker {
	this := IngestThrottlingTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *IngestThrottlingTracker) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestThrottlingTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *IngestThrottlingTracker) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *IngestThrottlingTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *IngestThrottlingTracker) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestThrottlingTracker) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *IngestThrottlingTracker) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *IngestThrottlingTracker) SetDataType(v string) {
	o.DataType = &v
}

func (o IngestThrottlingTracker) MarshalJSON() ([]byte, error) {
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
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	return json.Marshal(toSerialize)
}

type NullableIngestThrottlingTracker struct {
	value *IngestThrottlingTracker
	isSet bool
}

func (v NullableIngestThrottlingTracker) Get() *IngestThrottlingTracker {
	return v.value
}

func (v *NullableIngestThrottlingTracker) Set(val *IngestThrottlingTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestThrottlingTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestThrottlingTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestThrottlingTracker(val *IngestThrottlingTracker) *NullableIngestThrottlingTracker {
	return &NullableIngestThrottlingTracker{value: val, isSet: true}
}

func (v NullableIngestThrottlingTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestThrottlingTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


