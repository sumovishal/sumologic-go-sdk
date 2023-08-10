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

// checks if the CollectionAwsInventoryThrottledTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionAwsInventoryThrottledTracker{}

// CollectionAwsInventoryThrottledTracker struct for CollectionAwsInventoryThrottledTracker
type CollectionAwsInventoryThrottledTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
}

// NewCollectionAwsInventoryThrottledTracker instantiates a new CollectionAwsInventoryThrottledTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionAwsInventoryThrottledTracker(trackerId string, error_ string, description string) *CollectionAwsInventoryThrottledTracker {
	this := CollectionAwsInventoryThrottledTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCollectionAwsInventoryThrottledTrackerWithDefaults instantiates a new CollectionAwsInventoryThrottledTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionAwsInventoryThrottledTrackerWithDefaults() *CollectionAwsInventoryThrottledTracker {
	this := CollectionAwsInventoryThrottledTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CollectionAwsInventoryThrottledTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionAwsInventoryThrottledTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CollectionAwsInventoryThrottledTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CollectionAwsInventoryThrottledTracker) SetEventType(v string) {
	o.EventType = &v
}

func (o CollectionAwsInventoryThrottledTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionAwsInventoryThrottledTracker) ToMap() (map[string]interface{}, error) {
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

type NullableCollectionAwsInventoryThrottledTracker struct {
	value *CollectionAwsInventoryThrottledTracker
	isSet bool
}

func (v NullableCollectionAwsInventoryThrottledTracker) Get() *CollectionAwsInventoryThrottledTracker {
	return v.value
}

func (v *NullableCollectionAwsInventoryThrottledTracker) Set(val *CollectionAwsInventoryThrottledTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionAwsInventoryThrottledTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionAwsInventoryThrottledTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionAwsInventoryThrottledTracker(val *CollectionAwsInventoryThrottledTracker) *NullableCollectionAwsInventoryThrottledTracker {
	return &NullableCollectionAwsInventoryThrottledTracker{value: val, isSet: true}
}

func (v NullableCollectionAwsInventoryThrottledTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionAwsInventoryThrottledTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


