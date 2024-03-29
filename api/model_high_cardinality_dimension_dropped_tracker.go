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

// checks if the HighCardinalityDimensionDroppedTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighCardinalityDimensionDroppedTracker{}

// HighCardinalityDimensionDroppedTracker struct for HighCardinalityDimensionDroppedTracker
type HighCardinalityDimensionDroppedTracker struct {
	TrackerIdentity
	// The dropped high cardinality dimension.
	Dimension *string `json:"dimension,omitempty"`
}

// NewHighCardinalityDimensionDroppedTracker instantiates a new HighCardinalityDimensionDroppedTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighCardinalityDimensionDroppedTracker(trackerId string, error_ string, description string) *HighCardinalityDimensionDroppedTracker {
	this := HighCardinalityDimensionDroppedTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewHighCardinalityDimensionDroppedTrackerWithDefaults instantiates a new HighCardinalityDimensionDroppedTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighCardinalityDimensionDroppedTrackerWithDefaults() *HighCardinalityDimensionDroppedTracker {
	this := HighCardinalityDimensionDroppedTracker{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *HighCardinalityDimensionDroppedTracker) GetDimension() string {
	if o == nil || IsNil(o.Dimension) {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighCardinalityDimensionDroppedTracker) GetDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *HighCardinalityDimensionDroppedTracker) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *HighCardinalityDimensionDroppedTracker) SetDimension(v string) {
	o.Dimension = &v
}

func (o HighCardinalityDimensionDroppedTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighCardinalityDimensionDroppedTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	if !IsNil(o.Dimension) {
		toSerialize["dimension"] = o.Dimension
	}
	return toSerialize, nil
}

type NullableHighCardinalityDimensionDroppedTracker struct {
	value *HighCardinalityDimensionDroppedTracker
	isSet bool
}

func (v NullableHighCardinalityDimensionDroppedTracker) Get() *HighCardinalityDimensionDroppedTracker {
	return v.value
}

func (v *NullableHighCardinalityDimensionDroppedTracker) Set(val *HighCardinalityDimensionDroppedTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableHighCardinalityDimensionDroppedTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableHighCardinalityDimensionDroppedTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighCardinalityDimensionDroppedTracker(val *HighCardinalityDimensionDroppedTracker) *NullableHighCardinalityDimensionDroppedTracker {
	return &NullableHighCardinalityDimensionDroppedTracker{value: val, isSet: true}
}

func (v NullableHighCardinalityDimensionDroppedTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighCardinalityDimensionDroppedTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


