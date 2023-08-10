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

// checks if the Iso8601TimeRangeBoundaryAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Iso8601TimeRangeBoundaryAllOf{}

// Iso8601TimeRangeBoundaryAllOf struct for Iso8601TimeRangeBoundaryAllOf
type Iso8601TimeRangeBoundaryAllOf struct {
	// Starting point in time as a string in ISO 8601 format. For example `2018-10-01T11:10:20.52+01:00`
	Iso8601Time time.Time `json:"iso8601Time"`
}

// NewIso8601TimeRangeBoundaryAllOf instantiates a new Iso8601TimeRangeBoundaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIso8601TimeRangeBoundaryAllOf(iso8601Time time.Time) *Iso8601TimeRangeBoundaryAllOf {
	this := Iso8601TimeRangeBoundaryAllOf{}
	this.Iso8601Time = iso8601Time
	return &this
}

// NewIso8601TimeRangeBoundaryAllOfWithDefaults instantiates a new Iso8601TimeRangeBoundaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIso8601TimeRangeBoundaryAllOfWithDefaults() *Iso8601TimeRangeBoundaryAllOf {
	this := Iso8601TimeRangeBoundaryAllOf{}
	return &this
}

// GetIso8601Time returns the Iso8601Time field value
func (o *Iso8601TimeRangeBoundaryAllOf) GetIso8601Time() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Iso8601Time
}

// GetIso8601TimeOk returns a tuple with the Iso8601Time field value
// and a boolean to check if the value has been set.
func (o *Iso8601TimeRangeBoundaryAllOf) GetIso8601TimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iso8601Time, true
}

// SetIso8601Time sets field value
func (o *Iso8601TimeRangeBoundaryAllOf) SetIso8601Time(v time.Time) {
	o.Iso8601Time = v
}

func (o Iso8601TimeRangeBoundaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Iso8601TimeRangeBoundaryAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iso8601Time"] = o.Iso8601Time
	return toSerialize, nil
}

type NullableIso8601TimeRangeBoundaryAllOf struct {
	value *Iso8601TimeRangeBoundaryAllOf
	isSet bool
}

func (v NullableIso8601TimeRangeBoundaryAllOf) Get() *Iso8601TimeRangeBoundaryAllOf {
	return v.value
}

func (v *NullableIso8601TimeRangeBoundaryAllOf) Set(val *Iso8601TimeRangeBoundaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIso8601TimeRangeBoundaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIso8601TimeRangeBoundaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIso8601TimeRangeBoundaryAllOf(val *Iso8601TimeRangeBoundaryAllOf) *NullableIso8601TimeRangeBoundaryAllOf {
	return &NullableIso8601TimeRangeBoundaryAllOf{value: val, isSet: true}
}

func (v NullableIso8601TimeRangeBoundaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIso8601TimeRangeBoundaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


