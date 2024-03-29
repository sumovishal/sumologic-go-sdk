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

// checks if the SpansCountVisualization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpansCountVisualization{}

// SpansCountVisualization struct for SpansCountVisualization
type SpansCountVisualization struct {
	SpansVisualization
	// A field by which the spans need to be counted.
	DistinctBy *string `json:"distinctBy,omitempty"`
}

// NewSpansCountVisualization instantiates a new SpansCountVisualization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpansCountVisualization(type_ string, name string) *SpansCountVisualization {
	this := SpansCountVisualization{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewSpansCountVisualizationWithDefaults instantiates a new SpansCountVisualization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpansCountVisualizationWithDefaults() *SpansCountVisualization {
	this := SpansCountVisualization{}
	return &this
}

// GetDistinctBy returns the DistinctBy field value if set, zero value otherwise.
func (o *SpansCountVisualization) GetDistinctBy() string {
	if o == nil || IsNil(o.DistinctBy) {
		var ret string
		return ret
	}
	return *o.DistinctBy
}

// GetDistinctByOk returns a tuple with the DistinctBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansCountVisualization) GetDistinctByOk() (*string, bool) {
	if o == nil || IsNil(o.DistinctBy) {
		return nil, false
	}
	return o.DistinctBy, true
}

// HasDistinctBy returns a boolean if a field has been set.
func (o *SpansCountVisualization) HasDistinctBy() bool {
	if o != nil && !IsNil(o.DistinctBy) {
		return true
	}

	return false
}

// SetDistinctBy gets a reference to the given string and assigns it to the DistinctBy field.
func (o *SpansCountVisualization) SetDistinctBy(v string) {
	o.DistinctBy = &v
}

func (o SpansCountVisualization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpansCountVisualization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSpansVisualization, errSpansVisualization := json.Marshal(o.SpansVisualization)
	if errSpansVisualization != nil {
		return map[string]interface{}{}, errSpansVisualization
	}
	errSpansVisualization = json.Unmarshal([]byte(serializedSpansVisualization), &toSerialize)
	if errSpansVisualization != nil {
		return map[string]interface{}{}, errSpansVisualization
	}
	if !IsNil(o.DistinctBy) {
		toSerialize["distinctBy"] = o.DistinctBy
	}
	return toSerialize, nil
}

type NullableSpansCountVisualization struct {
	value *SpansCountVisualization
	isSet bool
}

func (v NullableSpansCountVisualization) Get() *SpansCountVisualization {
	return v.value
}

func (v *NullableSpansCountVisualization) Set(val *SpansCountVisualization) {
	v.value = val
	v.isSet = true
}

func (v NullableSpansCountVisualization) IsSet() bool {
	return v.isSet
}

func (v *NullableSpansCountVisualization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpansCountVisualization(val *SpansCountVisualization) *NullableSpansCountVisualization {
	return &NullableSpansCountVisualization{value: val, isSet: true}
}

func (v NullableSpansCountVisualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpansCountVisualization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


