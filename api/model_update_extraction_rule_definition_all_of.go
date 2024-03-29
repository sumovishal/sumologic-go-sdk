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

// checks if the UpdateExtractionRuleDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateExtractionRuleDefinitionAllOf{}

// UpdateExtractionRuleDefinitionAllOf struct for UpdateExtractionRuleDefinitionAllOf
type UpdateExtractionRuleDefinitionAllOf struct {
	// Is the field extraction rule enabled.
	Enabled bool `json:"enabled"`
}

// NewUpdateExtractionRuleDefinitionAllOf instantiates a new UpdateExtractionRuleDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExtractionRuleDefinitionAllOf(enabled bool) *UpdateExtractionRuleDefinitionAllOf {
	this := UpdateExtractionRuleDefinitionAllOf{}
	this.Enabled = enabled
	return &this
}

// NewUpdateExtractionRuleDefinitionAllOfWithDefaults instantiates a new UpdateExtractionRuleDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExtractionRuleDefinitionAllOfWithDefaults() *UpdateExtractionRuleDefinitionAllOf {
	this := UpdateExtractionRuleDefinitionAllOf{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *UpdateExtractionRuleDefinitionAllOf) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateExtractionRuleDefinitionAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateExtractionRuleDefinitionAllOf) SetEnabled(v bool) {
	o.Enabled = v
}

func (o UpdateExtractionRuleDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateExtractionRuleDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableUpdateExtractionRuleDefinitionAllOf struct {
	value *UpdateExtractionRuleDefinitionAllOf
	isSet bool
}

func (v NullableUpdateExtractionRuleDefinitionAllOf) Get() *UpdateExtractionRuleDefinitionAllOf {
	return v.value
}

func (v *NullableUpdateExtractionRuleDefinitionAllOf) Set(val *UpdateExtractionRuleDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExtractionRuleDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExtractionRuleDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExtractionRuleDefinitionAllOf(val *UpdateExtractionRuleDefinitionAllOf) *NullableUpdateExtractionRuleDefinitionAllOf {
	return &NullableUpdateExtractionRuleDefinitionAllOf{value: val, isSet: true}
}

func (v NullableUpdateExtractionRuleDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExtractionRuleDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


