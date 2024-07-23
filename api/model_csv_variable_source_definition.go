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

// checks if the CsvVariableSourceDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsvVariableSourceDefinition{}

// CsvVariableSourceDefinition struct for CsvVariableSourceDefinition
type CsvVariableSourceDefinition struct {
	VariableSourceDefinition
	// Comma separated values for the variable.
	Values string `json:"values"`
}

type _CsvVariableSourceDefinition CsvVariableSourceDefinition

// NewCsvVariableSourceDefinition instantiates a new CsvVariableSourceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsvVariableSourceDefinition(values string, variableSourceType string) *CsvVariableSourceDefinition {
	this := CsvVariableSourceDefinition{}
	this.VariableSourceType = variableSourceType
	this.Values = values
	return &this
}

// NewCsvVariableSourceDefinitionWithDefaults instantiates a new CsvVariableSourceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsvVariableSourceDefinitionWithDefaults() *CsvVariableSourceDefinition {
	this := CsvVariableSourceDefinition{}
	return &this
}

// GetValues returns the Values field value
func (o *CsvVariableSourceDefinition) GetValues() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *CsvVariableSourceDefinition) GetValuesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *CsvVariableSourceDefinition) SetValues(v string) {
	o.Values = v
}

func (o CsvVariableSourceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsvVariableSourceDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVariableSourceDefinition, errVariableSourceDefinition := json.Marshal(o.VariableSourceDefinition)
	if errVariableSourceDefinition != nil {
		return map[string]interface{}{}, errVariableSourceDefinition
	}
	errVariableSourceDefinition = json.Unmarshal([]byte(serializedVariableSourceDefinition), &toSerialize)
	if errVariableSourceDefinition != nil {
		return map[string]interface{}{}, errVariableSourceDefinition
	}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *CsvVariableSourceDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"values",
		"variableSourceType",
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

	varCsvVariableSourceDefinition := _CsvVariableSourceDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCsvVariableSourceDefinition)

	if err != nil {
		return err
	}

	*o = CsvVariableSourceDefinition(varCsvVariableSourceDefinition)

	return err
}

type NullableCsvVariableSourceDefinition struct {
	value *CsvVariableSourceDefinition
	isSet bool
}

func (v NullableCsvVariableSourceDefinition) Get() *CsvVariableSourceDefinition {
	return v.value
}

func (v *NullableCsvVariableSourceDefinition) Set(val *CsvVariableSourceDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCsvVariableSourceDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCsvVariableSourceDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsvVariableSourceDefinition(val *CsvVariableSourceDefinition) *NullableCsvVariableSourceDefinition {
	return &NullableCsvVariableSourceDefinition{value: val, isSet: true}
}

func (v NullableCsvVariableSourceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsvVariableSourceDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


