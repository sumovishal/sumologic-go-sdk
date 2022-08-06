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

// DimensionTransformation Base class of all transformation types.
type DimensionTransformation struct {
	// This is the base type of all dimension transformations.
	TransformationType string `json:"transformationType"`
}

// NewDimensionTransformation instantiates a new DimensionTransformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionTransformation(transformationType string) *DimensionTransformation {
	this := DimensionTransformation{}
	this.TransformationType = transformationType
	return &this
}

// NewDimensionTransformationWithDefaults instantiates a new DimensionTransformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionTransformationWithDefaults() *DimensionTransformation {
	this := DimensionTransformation{}
	return &this
}

// GetTransformationType returns the TransformationType field value
func (o *DimensionTransformation) GetTransformationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransformationType
}

// GetTransformationTypeOk returns a tuple with the TransformationType field value
// and a boolean to check if the value has been set.
func (o *DimensionTransformation) GetTransformationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransformationType, true
}

// SetTransformationType sets field value
func (o *DimensionTransformation) SetTransformationType(v string) {
	o.TransformationType = v
}

func (o DimensionTransformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transformationType"] = o.TransformationType
	}
	return json.Marshal(toSerialize)
}

type NullableDimensionTransformation struct {
	value *DimensionTransformation
	isSet bool
}

func (v NullableDimensionTransformation) Get() *DimensionTransformation {
	return v.value
}

func (v *NullableDimensionTransformation) Set(val *DimensionTransformation) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionTransformation) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionTransformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionTransformation(val *DimensionTransformation) *NullableDimensionTransformation {
	return &NullableDimensionTransformation{value: val, isSet: true}
}

func (v NullableDimensionTransformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimensionTransformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


