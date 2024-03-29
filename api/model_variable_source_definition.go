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

// checks if the VariableSourceDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableSourceDefinition{}

// VariableSourceDefinition struct for VariableSourceDefinition
type VariableSourceDefinition struct {
	// Source type of the variable values.
	VariableSourceType string `json:"variableSourceType"`
}

// NewVariableSourceDefinition instantiates a new VariableSourceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableSourceDefinition(variableSourceType string) *VariableSourceDefinition {
	this := VariableSourceDefinition{}
	this.VariableSourceType = variableSourceType
	return &this
}

// NewVariableSourceDefinitionWithDefaults instantiates a new VariableSourceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableSourceDefinitionWithDefaults() *VariableSourceDefinition {
	this := VariableSourceDefinition{}
	return &this
}

// GetVariableSourceType returns the VariableSourceType field value
func (o *VariableSourceDefinition) GetVariableSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariableSourceType
}

// GetVariableSourceTypeOk returns a tuple with the VariableSourceType field value
// and a boolean to check if the value has been set.
func (o *VariableSourceDefinition) GetVariableSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariableSourceType, true
}

// SetVariableSourceType sets field value
func (o *VariableSourceDefinition) SetVariableSourceType(v string) {
	o.VariableSourceType = v
}

func (o VariableSourceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableSourceDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variableSourceType"] = o.VariableSourceType
	return toSerialize, nil
}

type NullableVariableSourceDefinition struct {
	value *VariableSourceDefinition
	isSet bool
}

func (v NullableVariableSourceDefinition) Get() *VariableSourceDefinition {
	return v.value
}

func (v *NullableVariableSourceDefinition) Set(val *VariableSourceDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableSourceDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableSourceDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableSourceDefinition(val *VariableSourceDefinition) *NullableVariableSourceDefinition {
	return &NullableVariableSourceDefinition{value: val, isSet: true}
}

func (v NullableVariableSourceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableSourceDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


