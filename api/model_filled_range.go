/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// checks if the FilledRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilledRange{}

// FilledRange Range of timestamps already filled since the autoview has been created.
type FilledRange struct {
	// Start of the timestamp for each unit of filled ranges, expressed in UTC.
	StartTime time.Time `json:"startTime"`
	// End of the timestamp for each unit of filled ranges, expressed in UTC.
	EndTime time.Time `json:"endTime"`
}

// NewFilledRange instantiates a new FilledRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilledRange(startTime time.Time, endTime time.Time) *FilledRange {
	this := FilledRange{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewFilledRangeWithDefaults instantiates a new FilledRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilledRangeWithDefaults() *FilledRange {
	this := FilledRange{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *FilledRange) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *FilledRange) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *FilledRange) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *FilledRange) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *FilledRange) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *FilledRange) SetEndTime(v time.Time) {
	o.EndTime = v
}

func (o FilledRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilledRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	return toSerialize, nil
}

type NullableFilledRange struct {
	value *FilledRange
	isSet bool
}

func (v NullableFilledRange) Get() *FilledRange {
	return v.value
}

func (v *NullableFilledRange) Set(val *FilledRange) {
	v.value = val
	v.isSet = true
}

func (v NullableFilledRange) IsSet() bool {
	return v.isSet
}

func (v *NullableFilledRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilledRange(val *FilledRange) *NullableFilledRange {
	return &NullableFilledRange{value: val, isSet: true}
}

func (v NullableFilledRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilledRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


