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

// checks if the Grid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Grid{}

// Grid struct for Grid
type Grid struct {
	Layout
}

type _Grid Grid

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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Grid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedLayout, errLayout := json.Marshal(o.Layout)
	if errLayout != nil {
		return map[string]interface{}{}, errLayout
	}
	errLayout = json.Unmarshal([]byte(serializedLayout), &toSerialize)
	if errLayout != nil {
		return map[string]interface{}{}, errLayout
	}
	return toSerialize, nil
}

func (o *Grid) UnmarshalJSON(data []byte) (err error) {
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

	varGrid := _Grid{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGrid)

	if err != nil {
		return err
	}

	*o = Grid(varGrid)

	return err
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


