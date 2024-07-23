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

// checks if the RowUpdateDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RowUpdateDefinition{}

// RowUpdateDefinition Lookup table data to be uploaded.
type RowUpdateDefinition struct {
	// A list of all the field identifiers and their corresponding values.
	Row []TableRow `json:"row"`
}

type _RowUpdateDefinition RowUpdateDefinition

// NewRowUpdateDefinition instantiates a new RowUpdateDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRowUpdateDefinition(row []TableRow) *RowUpdateDefinition {
	this := RowUpdateDefinition{}
	this.Row = row
	return &this
}

// NewRowUpdateDefinitionWithDefaults instantiates a new RowUpdateDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRowUpdateDefinitionWithDefaults() *RowUpdateDefinition {
	this := RowUpdateDefinition{}
	return &this
}

// GetRow returns the Row field value
func (o *RowUpdateDefinition) GetRow() []TableRow {
	if o == nil {
		var ret []TableRow
		return ret
	}

	return o.Row
}

// GetRowOk returns a tuple with the Row field value
// and a boolean to check if the value has been set.
func (o *RowUpdateDefinition) GetRowOk() ([]TableRow, bool) {
	if o == nil {
		return nil, false
	}
	return o.Row, true
}

// SetRow sets field value
func (o *RowUpdateDefinition) SetRow(v []TableRow) {
	o.Row = v
}

func (o RowUpdateDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RowUpdateDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["row"] = o.Row
	return toSerialize, nil
}

func (o *RowUpdateDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"row",
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

	varRowUpdateDefinition := _RowUpdateDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRowUpdateDefinition)

	if err != nil {
		return err
	}

	*o = RowUpdateDefinition(varRowUpdateDefinition)

	return err
}

type NullableRowUpdateDefinition struct {
	value *RowUpdateDefinition
	isSet bool
}

func (v NullableRowUpdateDefinition) Get() *RowUpdateDefinition {
	return v.value
}

func (v *NullableRowUpdateDefinition) Set(val *RowUpdateDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableRowUpdateDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableRowUpdateDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRowUpdateDefinition(val *RowUpdateDefinition) *NullableRowUpdateDefinition {
	return &NullableRowUpdateDefinition{value: val, isSet: true}
}

func (v NullableRowUpdateDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRowUpdateDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


