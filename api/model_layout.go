/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Layout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Layout{}

// Layout struct for Layout
type Layout struct {
	// The type of panel layout on the Dashboard. For example, Grid, Tabs, or Hierarchical. Currently supports `Grid` only.
	LayoutType string `json:"layoutType"`
	// Layout structures for the panel childen.
	LayoutStructures []LayoutStructure `json:"layoutStructures"`
}

type _Layout Layout

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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Layout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["layoutType"] = o.LayoutType
	toSerialize["layoutStructures"] = o.LayoutStructures
	return toSerialize, nil
}

func (o *Layout) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"layoutType",
		"layoutStructures",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLayout := _Layout{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLayout)

	if err != nil {
		return err
	}

	*o = Layout(varLayout)

	return err
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


