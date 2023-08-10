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

// checks if the LogSearchParameterAutoCompleteSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogSearchParameterAutoCompleteSyncDefinition{}

// LogSearchParameterAutoCompleteSyncDefinition struct for LogSearchParameterAutoCompleteSyncDefinition
type LogSearchParameterAutoCompleteSyncDefinition struct {
	// The autocomplete parameter type.
	AutoCompleteType string `json:"autoCompleteType"`
}

// NewLogSearchParameterAutoCompleteSyncDefinition instantiates a new LogSearchParameterAutoCompleteSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchParameterAutoCompleteSyncDefinition(autoCompleteType string) *LogSearchParameterAutoCompleteSyncDefinition {
	this := LogSearchParameterAutoCompleteSyncDefinition{}
	this.AutoCompleteType = autoCompleteType
	return &this
}

// NewLogSearchParameterAutoCompleteSyncDefinitionWithDefaults instantiates a new LogSearchParameterAutoCompleteSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchParameterAutoCompleteSyncDefinitionWithDefaults() *LogSearchParameterAutoCompleteSyncDefinition {
	this := LogSearchParameterAutoCompleteSyncDefinition{}
	return &this
}

// GetAutoCompleteType returns the AutoCompleteType field value
func (o *LogSearchParameterAutoCompleteSyncDefinition) GetAutoCompleteType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoCompleteType
}

// GetAutoCompleteTypeOk returns a tuple with the AutoCompleteType field value
// and a boolean to check if the value has been set.
func (o *LogSearchParameterAutoCompleteSyncDefinition) GetAutoCompleteTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCompleteType, true
}

// SetAutoCompleteType sets field value
func (o *LogSearchParameterAutoCompleteSyncDefinition) SetAutoCompleteType(v string) {
	o.AutoCompleteType = v
}

func (o LogSearchParameterAutoCompleteSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogSearchParameterAutoCompleteSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["autoCompleteType"] = o.AutoCompleteType
	return toSerialize, nil
}

type NullableLogSearchParameterAutoCompleteSyncDefinition struct {
	value *LogSearchParameterAutoCompleteSyncDefinition
	isSet bool
}

func (v NullableLogSearchParameterAutoCompleteSyncDefinition) Get() *LogSearchParameterAutoCompleteSyncDefinition {
	return v.value
}

func (v *NullableLogSearchParameterAutoCompleteSyncDefinition) Set(val *LogSearchParameterAutoCompleteSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchParameterAutoCompleteSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchParameterAutoCompleteSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchParameterAutoCompleteSyncDefinition(val *LogSearchParameterAutoCompleteSyncDefinition) *NullableLogSearchParameterAutoCompleteSyncDefinition {
	return &NullableLogSearchParameterAutoCompleteSyncDefinition{value: val, isSet: true}
}

func (v NullableLogSearchParameterAutoCompleteSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchParameterAutoCompleteSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


