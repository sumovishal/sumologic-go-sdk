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

// checks if the CollectionPathAccessDeniedTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionPathAccessDeniedTracker{}

// CollectionPathAccessDeniedTracker struct for CollectionPathAccessDeniedTracker
type CollectionPathAccessDeniedTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The path to the file.
	Path *string `json:"path,omitempty"`
}

// NewCollectionPathAccessDeniedTracker instantiates a new CollectionPathAccessDeniedTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionPathAccessDeniedTracker(trackerId string, error_ string, description string) *CollectionPathAccessDeniedTracker {
	this := CollectionPathAccessDeniedTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCollectionPathAccessDeniedTrackerWithDefaults instantiates a new CollectionPathAccessDeniedTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionPathAccessDeniedTrackerWithDefaults() *CollectionPathAccessDeniedTracker {
	this := CollectionPathAccessDeniedTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CollectionPathAccessDeniedTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPathAccessDeniedTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CollectionPathAccessDeniedTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CollectionPathAccessDeniedTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CollectionPathAccessDeniedTracker) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPathAccessDeniedTracker) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CollectionPathAccessDeniedTracker) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CollectionPathAccessDeniedTracker) SetPath(v string) {
	o.Path = &v
}

func (o CollectionPathAccessDeniedTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionPathAccessDeniedTracker) ToMap() (map[string]interface{}, error) {
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

type NullableCollectionPathAccessDeniedTracker struct {
	value *CollectionPathAccessDeniedTracker
	isSet bool
}

func (v NullableCollectionPathAccessDeniedTracker) Get() *CollectionPathAccessDeniedTracker {
	return v.value
}

func (v *NullableCollectionPathAccessDeniedTracker) Set(val *CollectionPathAccessDeniedTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPathAccessDeniedTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPathAccessDeniedTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPathAccessDeniedTracker(val *CollectionPathAccessDeniedTracker) *NullableCollectionPathAccessDeniedTracker {
	return &NullableCollectionPathAccessDeniedTracker{value: val, isSet: true}
}

func (v NullableCollectionPathAccessDeniedTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPathAccessDeniedTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


