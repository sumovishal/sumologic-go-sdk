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

// CriticalPathServiceBreakdownSummary struct for CriticalPathServiceBreakdownSummary
type CriticalPathServiceBreakdownSummary struct {
	// List of the elements representing the critical path service duration breakdown - contains the first few services with the longest overall duration of the spans contributing to the critical path.
	Elements []CriticalPathServiceBreakdownElementBase `json:"elements"`
	// Overall processing time in nanoseconds consumed by the rest of the spans in the critical path (a sum of the duration times of the spans' critical path segments).
	OtherServicesDuration int64 `json:"otherServicesDuration"`
}

// NewCriticalPathServiceBreakdownSummary instantiates a new CriticalPathServiceBreakdownSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCriticalPathServiceBreakdownSummary(elements []CriticalPathServiceBreakdownElementBase, otherServicesDuration int64) *CriticalPathServiceBreakdownSummary {
	this := CriticalPathServiceBreakdownSummary{}
	this.Elements = elements
	this.OtherServicesDuration = otherServicesDuration
	return &this
}

// NewCriticalPathServiceBreakdownSummaryWithDefaults instantiates a new CriticalPathServiceBreakdownSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCriticalPathServiceBreakdownSummaryWithDefaults() *CriticalPathServiceBreakdownSummary {
	this := CriticalPathServiceBreakdownSummary{}
	return &this
}

// GetElements returns the Elements field value
func (o *CriticalPathServiceBreakdownSummary) GetElements() []CriticalPathServiceBreakdownElementBase {
	if o == nil {
		var ret []CriticalPathServiceBreakdownElementBase
		return ret
	}

	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value
// and a boolean to check if the value has been set.
func (o *CriticalPathServiceBreakdownSummary) GetElementsOk() ([]CriticalPathServiceBreakdownElementBase, bool) {
	if o == nil {
		return nil, false
	}
	return o.Elements, true
}

// SetElements sets field value
func (o *CriticalPathServiceBreakdownSummary) SetElements(v []CriticalPathServiceBreakdownElementBase) {
	o.Elements = v
}

// GetOtherServicesDuration returns the OtherServicesDuration field value
func (o *CriticalPathServiceBreakdownSummary) GetOtherServicesDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OtherServicesDuration
}

// GetOtherServicesDurationOk returns a tuple with the OtherServicesDuration field value
// and a boolean to check if the value has been set.
func (o *CriticalPathServiceBreakdownSummary) GetOtherServicesDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtherServicesDuration, true
}

// SetOtherServicesDuration sets field value
func (o *CriticalPathServiceBreakdownSummary) SetOtherServicesDuration(v int64) {
	o.OtherServicesDuration = v
}

func (o CriticalPathServiceBreakdownSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["elements"] = o.Elements
	}
	if true {
		toSerialize["otherServicesDuration"] = o.OtherServicesDuration
	}
	return json.Marshal(toSerialize)
}

type NullableCriticalPathServiceBreakdownSummary struct {
	value *CriticalPathServiceBreakdownSummary
	isSet bool
}

func (v NullableCriticalPathServiceBreakdownSummary) Get() *CriticalPathServiceBreakdownSummary {
	return v.value
}

func (v *NullableCriticalPathServiceBreakdownSummary) Set(val *CriticalPathServiceBreakdownSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCriticalPathServiceBreakdownSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCriticalPathServiceBreakdownSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCriticalPathServiceBreakdownSummary(val *CriticalPathServiceBreakdownSummary) *NullableCriticalPathServiceBreakdownSummary {
	return &NullableCriticalPathServiceBreakdownSummary{value: val, isSet: true}
}

func (v NullableCriticalPathServiceBreakdownSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCriticalPathServiceBreakdownSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


