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

// checks if the LogSearchQueryParameterSyncDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogSearchQueryParameterSyncDefinitionAllOf{}

// LogSearchQueryParameterSyncDefinitionAllOf struct for LogSearchQueryParameterSyncDefinitionAllOf
type LogSearchQueryParameterSyncDefinitionAllOf struct {
	AutoComplete LogSearchParameterAutoCompleteSyncDefinition `json:"autoComplete"`
}

// NewLogSearchQueryParameterSyncDefinitionAllOf instantiates a new LogSearchQueryParameterSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchQueryParameterSyncDefinitionAllOf(autoComplete LogSearchParameterAutoCompleteSyncDefinition) *LogSearchQueryParameterSyncDefinitionAllOf {
	this := LogSearchQueryParameterSyncDefinitionAllOf{}
	this.AutoComplete = autoComplete
	return &this
}

// NewLogSearchQueryParameterSyncDefinitionAllOfWithDefaults instantiates a new LogSearchQueryParameterSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchQueryParameterSyncDefinitionAllOfWithDefaults() *LogSearchQueryParameterSyncDefinitionAllOf {
	this := LogSearchQueryParameterSyncDefinitionAllOf{}
	return &this
}

// GetAutoComplete returns the AutoComplete field value
func (o *LogSearchQueryParameterSyncDefinitionAllOf) GetAutoComplete() LogSearchParameterAutoCompleteSyncDefinition {
	if o == nil {
		var ret LogSearchParameterAutoCompleteSyncDefinition
		return ret
	}

	return o.AutoComplete
}

// GetAutoCompleteOk returns a tuple with the AutoComplete field value
// and a boolean to check if the value has been set.
func (o *LogSearchQueryParameterSyncDefinitionAllOf) GetAutoCompleteOk() (*LogSearchParameterAutoCompleteSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoComplete, true
}

// SetAutoComplete sets field value
func (o *LogSearchQueryParameterSyncDefinitionAllOf) SetAutoComplete(v LogSearchParameterAutoCompleteSyncDefinition) {
	o.AutoComplete = v
}

func (o LogSearchQueryParameterSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogSearchQueryParameterSyncDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["autoComplete"] = o.AutoComplete
	return toSerialize, nil
}

type NullableLogSearchQueryParameterSyncDefinitionAllOf struct {
	value *LogSearchQueryParameterSyncDefinitionAllOf
	isSet bool
}

func (v NullableLogSearchQueryParameterSyncDefinitionAllOf) Get() *LogSearchQueryParameterSyncDefinitionAllOf {
	return v.value
}

func (v *NullableLogSearchQueryParameterSyncDefinitionAllOf) Set(val *LogSearchQueryParameterSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchQueryParameterSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchQueryParameterSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchQueryParameterSyncDefinitionAllOf(val *LogSearchQueryParameterSyncDefinitionAllOf) *NullableLogSearchQueryParameterSyncDefinitionAllOf {
	return &NullableLogSearchQueryParameterSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableLogSearchQueryParameterSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchQueryParameterSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


