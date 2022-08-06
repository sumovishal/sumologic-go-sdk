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

// Grid struct for Grid
type Grid struct {
	Layout
}

// NewGrid instantiates a new Grid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrid(layoutType string, layoutStructures []LayoutStructure) *Grid {
	this := Grid{}
	this.LayoutType = layoutType
	this.LayoutStructures = layoutStructures
	return &this
}

// NewGridWithDefaults instantiates a new Grid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridWithDefaults() *Grid {
	this := Grid{}
	return &this
}

func (o Grid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedLayout, errLayout := json.Marshal(o.Layout)
	if errLayout != nil {
		return []byte{}, errLayout
	}
	errLayout = json.Unmarshal([]byte(serializedLayout), &toSerialize)
	if errLayout != nil {
		return []byte{}, errLayout
	}
	return json.Marshal(toSerialize)
}

type NullableGrid struct {
	value *Grid
	isSet bool
}

func (v NullableGrid) Get() *Grid {
	return v.value
}

func (v *NullableGrid) Set(val *Grid) {
	v.value = val
	v.isSet = true
}

func (v NullableGrid) IsSet() bool {
	return v.isSet
}

func (v *NullableGrid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrid(val *Grid) *NullableGrid {
	return &NullableGrid{value: val, isSet: true}
}

func (v NullableGrid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


