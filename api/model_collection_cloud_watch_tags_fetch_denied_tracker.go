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

// checks if the CollectionCloudWatchTagsFetchDeniedTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionCloudWatchTagsFetchDeniedTracker{}

// CollectionCloudWatchTagsFetchDeniedTracker struct for CollectionCloudWatchTagsFetchDeniedTracker
type CollectionCloudWatchTagsFetchDeniedTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
}

// NewCollectionCloudWatchTagsFetchDeniedTracker instantiates a new CollectionCloudWatchTagsFetchDeniedTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionCloudWatchTagsFetchDeniedTracker(trackerId string, error_ string, description string) *CollectionCloudWatchTagsFetchDeniedTracker {
	this := CollectionCloudWatchTagsFetchDeniedTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCollectionCloudWatchTagsFetchDeniedTrackerWithDefaults instantiates a new CollectionCloudWatchTagsFetchDeniedTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionCloudWatchTagsFetchDeniedTrackerWithDefaults() *CollectionCloudWatchTagsFetchDeniedTracker {
	this := CollectionCloudWatchTagsFetchDeniedTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CollectionCloudWatchTagsFetchDeniedTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionCloudWatchTagsFetchDeniedTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CollectionCloudWatchTagsFetchDeniedTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CollectionCloudWatchTagsFetchDeniedTracker) SetEventType(v string) {
	o.EventType = &v
}

func (o CollectionCloudWatchTagsFetchDeniedTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionCloudWatchTagsFetchDeniedTracker) ToMap() (map[string]interface{}, error) {
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

type NullableCollectionCloudWatchTagsFetchDeniedTracker struct {
	value *CollectionCloudWatchTagsFetchDeniedTracker
	isSet bool
}

func (v NullableCollectionCloudWatchTagsFetchDeniedTracker) Get() *CollectionCloudWatchTagsFetchDeniedTracker {
	return v.value
}

func (v *NullableCollectionCloudWatchTagsFetchDeniedTracker) Set(val *CollectionCloudWatchTagsFetchDeniedTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionCloudWatchTagsFetchDeniedTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionCloudWatchTagsFetchDeniedTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionCloudWatchTagsFetchDeniedTracker(val *CollectionCloudWatchTagsFetchDeniedTracker) *NullableCollectionCloudWatchTagsFetchDeniedTracker {
	return &NullableCollectionCloudWatchTagsFetchDeniedTracker{value: val, isSet: true}
}

func (v NullableCollectionCloudWatchTagsFetchDeniedTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionCloudWatchTagsFetchDeniedTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


