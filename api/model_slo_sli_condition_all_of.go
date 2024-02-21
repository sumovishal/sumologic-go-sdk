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

// checks if the SloSliConditionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloSliConditionAllOf{}

// SloSliConditionAllOf A rule that defines how SLO monitors should evaluate remaining error budget and trigger notifications.
type SloSliConditionAllOf struct {
	// The remaining SLI error budget threshold percentage.
	SliThreshold float64 `json:"sliThreshold"`
}

// NewSloSliConditionAllOf instantiates a new SloSliConditionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloSliConditionAllOf(sliThreshold float64) *SloSliConditionAllOf {
	this := SloSliConditionAllOf{}
	this.SliThreshold = sliThreshold
	return &this
}

// NewSloSliConditionAllOfWithDefaults instantiates a new SloSliConditionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloSliConditionAllOfWithDefaults() *SloSliConditionAllOf {
	this := SloSliConditionAllOf{}
	return &this
}

// GetSliThreshold returns the SliThreshold field value
func (o *SloSliConditionAllOf) GetSliThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SliThreshold
}

// GetSliThresholdOk returns a tuple with the SliThreshold field value
// and a boolean to check if the value has been set.
func (o *SloSliConditionAllOf) GetSliThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SliThreshold, true
}

// SetSliThreshold sets field value
func (o *SloSliConditionAllOf) SetSliThreshold(v float64) {
	o.SliThreshold = v
}

func (o SloSliConditionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloSliConditionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sliThreshold"] = o.SliThreshold
	return toSerialize, nil
}

type NullableSloSliConditionAllOf struct {
	value *SloSliConditionAllOf
	isSet bool
}

func (v NullableSloSliConditionAllOf) Get() *SloSliConditionAllOf {
	return v.value
}

func (v *NullableSloSliConditionAllOf) Set(val *SloSliConditionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSloSliConditionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSloSliConditionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloSliConditionAllOf(val *SloSliConditionAllOf) *NullableSloSliConditionAllOf {
	return &NullableSloSliConditionAllOf{value: val, isSet: true}
}

func (v NullableSloSliConditionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloSliConditionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


