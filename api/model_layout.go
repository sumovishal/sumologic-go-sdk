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

// Layout struct for Layout
type Layout struct {
	// The type of panel layout on the Dashboard. For example, Grid, Tabs, or Hierarchical. Currently supports `Grid` only.
	LayoutType string `json:"layoutType"`
	// Layout structures for the panel childen.
	LayoutStructures []LayoutStructure `json:"layoutStructures"`
}

// NewLayout instantiates a new Layout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayout(layoutType string, layoutStructures []LayoutStructure) *Layout {
	this := Layout{}
	this.LayoutType = layoutType
	this.LayoutStructures = layoutStructures
	return &this
}

// NewLayoutWithDefaults instantiates a new Layout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutWithDefaults() *Layout {
	this := Layout{}
	return &this
}

// GetLayoutType returns the LayoutType field value
func (o *Layout) GetLayoutType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LayoutType
}

// GetLayoutTypeOk returns a tuple with the LayoutType field value
// and a boolean to check if the value has been set.
func (o *Layout) GetLayoutTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LayoutType, true
}

// SetLayoutType sets field value
func (o *Layout) SetLayoutType(v string) {
	o.LayoutType = v
}

// GetLayoutStructures returns the LayoutStructures field value
func (o *Layout) GetLayoutStructures() []LayoutStructure {
	if o == nil {
		var ret []LayoutStructure
		return ret
	}

	return o.LayoutStructures
}

// GetLayoutStructuresOk returns a tuple with the LayoutStructures field value
// and a boolean to check if the value has been set.
func (o *Layout) GetLayoutStructuresOk() ([]LayoutStructure, bool) {
	if o == nil {
		return nil, false
	}
	return o.LayoutStructures, true
}

// SetLayoutStructures sets field value
func (o *Layout) SetLayoutStructures(v []LayoutStructure) {
	o.LayoutStructures = v
}

func (o Layout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["layoutType"] = o.LayoutType
	}
	if true {
		toSerialize["layoutStructures"] = o.LayoutStructures
	}
	return json.Marshal(toSerialize)
}

type NullableLayout struct {
	value *Layout
	isSet bool
}

func (v NullableLayout) Get() *Layout {
	return v.value
}

func (v *NullableLayout) Set(val *Layout) {
	v.value = val
	v.isSet = true
}

func (v NullableLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayout(val *Layout) *NullableLayout {
	return &NullableLayout{value: val, isSet: true}
}

func (v NullableLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


