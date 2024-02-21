/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CollectionInvalidFilePathTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionInvalidFilePathTracker{}

// CollectionInvalidFilePathTracker struct for CollectionInvalidFilePathTracker
type CollectionInvalidFilePathTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The path to the file.
	Path *string `json:"path,omitempty"`
}

// NewCollectionInvalidFilePathTracker instantiates a new CollectionInvalidFilePathTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionInvalidFilePathTracker(trackerId string, error_ string, description string) *CollectionInvalidFilePathTracker {
	this := CollectionInvalidFilePathTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCollectionInvalidFilePathTrackerWithDefaults instantiates a new CollectionInvalidFilePathTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionInvalidFilePathTrackerWithDefaults() *CollectionInvalidFilePathTracker {
	this := CollectionInvalidFilePathTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CollectionInvalidFilePathTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionInvalidFilePathTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CollectionInvalidFilePathTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CollectionInvalidFilePathTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CollectionInvalidFilePathTracker) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionInvalidFilePathTracker) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CollectionInvalidFilePathTracker) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CollectionInvalidFilePathTracker) SetPath(v string) {
	o.Path = &v
}

func (o CollectionInvalidFilePathTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionInvalidFilePathTracker) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableCollectionInvalidFilePathTracker struct {
	value *CollectionInvalidFilePathTracker
	isSet bool
}

func (v NullableCollectionInvalidFilePathTracker) Get() *CollectionInvalidFilePathTracker {
	return v.value
}

func (v *NullableCollectionInvalidFilePathTracker) Set(val *CollectionInvalidFilePathTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionInvalidFilePathTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionInvalidFilePathTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionInvalidFilePathTracker(val *CollectionInvalidFilePathTracker) *NullableCollectionInvalidFilePathTracker {
	return &NullableCollectionInvalidFilePathTracker{value: val, isSet: true}
}

func (v NullableCollectionInvalidFilePathTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionInvalidFilePathTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


