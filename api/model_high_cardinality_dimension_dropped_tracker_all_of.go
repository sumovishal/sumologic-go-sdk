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

// checks if the HighCardinalityDimensionDroppedTrackerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighCardinalityDimensionDroppedTrackerAllOf{}

// HighCardinalityDimensionDroppedTrackerAllOf struct for HighCardinalityDimensionDroppedTrackerAllOf
type HighCardinalityDimensionDroppedTrackerAllOf struct {
	// The dropped high cardinality dimension.
	Dimension *string `json:"dimension,omitempty"`
}

// NewHighCardinalityDimensionDroppedTrackerAllOf instantiates a new HighCardinalityDimensionDroppedTrackerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighCardinalityDimensionDroppedTrackerAllOf() *HighCardinalityDimensionDroppedTrackerAllOf {
	this := HighCardinalityDimensionDroppedTrackerAllOf{}
	return &this
}

// NewHighCardinalityDimensionDroppedTrackerAllOfWithDefaults instantiates a new HighCardinalityDimensionDroppedTrackerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighCardinalityDimensionDroppedTrackerAllOfWithDefaults() *HighCardinalityDimensionDroppedTrackerAllOf {
	this := HighCardinalityDimensionDroppedTrackerAllOf{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *HighCardinalityDimensionDroppedTrackerAllOf) GetDimension() string {
	if o == nil || IsNil(o.Dimension) {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighCardinalityDimensionDroppedTrackerAllOf) GetDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *HighCardinalityDimensionDroppedTrackerAllOf) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *HighCardinalityDimensionDroppedTrackerAllOf) SetDimension(v string) {
	o.Dimension = &v
}

func (o HighCardinalityDimensionDroppedTrackerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighCardinalityDimensionDroppedTrackerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimension) {
		toSerialize["dimension"] = o.Dimension
	}
	return toSerialize, nil
}

type NullableHighCardinalityDimensionDroppedTrackerAllOf struct {
	value *HighCardinalityDimensionDroppedTrackerAllOf
	isSet bool
}

func (v NullableHighCardinalityDimensionDroppedTrackerAllOf) Get() *HighCardinalityDimensionDroppedTrackerAllOf {
	return v.value
}

func (v *NullableHighCardinalityDimensionDroppedTrackerAllOf) Set(val *HighCardinalityDimensionDroppedTrackerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHighCardinalityDimensionDroppedTrackerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHighCardinalityDimensionDroppedTrackerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighCardinalityDimensionDroppedTrackerAllOf(val *HighCardinalityDimensionDroppedTrackerAllOf) *NullableHighCardinalityDimensionDroppedTrackerAllOf {
	return &NullableHighCardinalityDimensionDroppedTrackerAllOf{value: val, isSet: true}
}

func (v NullableHighCardinalityDimensionDroppedTrackerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighCardinalityDimensionDroppedTrackerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


