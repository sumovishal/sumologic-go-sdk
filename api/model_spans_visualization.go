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

// SpansVisualization struct for SpansVisualization
type SpansVisualization struct {
	// The visualization type.
	Type string `json:"type"`
	// A unique name of the visualization.
	Name string `json:"name"`
}

// NewSpansVisualization instantiates a new SpansVisualization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpansVisualization(type_ string, name string) *SpansVisualization {
	this := SpansVisualization{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewSpansVisualizationWithDefaults instantiates a new SpansVisualization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpansVisualizationWithDefaults() *SpansVisualization {
	this := SpansVisualization{}
	return &this
}

// GetType returns the Type field value
func (o *SpansVisualization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpansVisualization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpansVisualization) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *SpansVisualization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpansVisualization) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpansVisualization) SetName(v string) {
	o.Name = v
}

func (o SpansVisualization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSpansVisualization struct {
	value *SpansVisualization
	isSet bool
}

func (v NullableSpansVisualization) Get() *SpansVisualization {
	return v.value
}

func (v *NullableSpansVisualization) Set(val *SpansVisualization) {
	v.value = val
	v.isSet = true
}

func (v NullableSpansVisualization) IsSet() bool {
	return v.isSet
}

func (v *NullableSpansVisualization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpansVisualization(val *SpansVisualization) *NullableSpansVisualization {
	return &NullableSpansVisualization{value: val, isSet: true}
}

func (v NullableSpansVisualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpansVisualization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


