/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateExtractionRuleDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateExtractionRuleDefinition{}

// UpdateExtractionRuleDefinition struct for UpdateExtractionRuleDefinition
type UpdateExtractionRuleDefinition struct {
	// Name of the field extraction rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Describes the fields to be parsed.
	ParseExpression string `json:"parseExpression"`
	// Is the field extraction rule enabled.
	Enabled bool `json:"enabled"`
}

// NewUpdateExtractionRuleDefinition instantiates a new UpdateExtractionRuleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExtractionRuleDefinition(name string, scope string, parseExpression string, enabled bool) *UpdateExtractionRuleDefinition {
	this := UpdateExtractionRuleDefinition{}
	this.Name = name
	this.Scope = scope
	this.ParseExpression = parseExpression
	this.Enabled = enabled
	return &this
}

// NewUpdateExtractionRuleDefinitionWithDefaults instantiates a new UpdateExtractionRuleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExtractionRuleDefinitionWithDefaults() *UpdateExtractionRuleDefinition {
	this := UpdateExtractionRuleDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateExtractionRuleDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateExtractionRuleDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateExtractionRuleDefinition) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *UpdateExtractionRuleDefinition) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *UpdateExtractionRuleDefinition) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *UpdateExtractionRuleDefinition) SetScope(v string) {
	o.Scope = v
}

// GetParseExpression returns the ParseExpression field value
func (o *UpdateExtractionRuleDefinition) GetParseExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParseExpression
}

// GetParseExpressionOk returns a tuple with the ParseExpression field value
// and a boolean to check if the value has been set.
func (o *UpdateExtractionRuleDefinition) GetParseExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParseExpression, true
}

// SetParseExpression sets field value
func (o *UpdateExtractionRuleDefinition) SetParseExpression(v string) {
	o.ParseExpression = v
}

// GetEnabled returns the Enabled field value
func (o *UpdateExtractionRuleDefinition) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateExtractionRuleDefinition) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateExtractionRuleDefinition) SetEnabled(v bool) {
	o.Enabled = v
}

func (o UpdateExtractionRuleDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateExtractionRuleDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["scope"] = o.Scope
	toSerialize["parseExpression"] = o.ParseExpression
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableUpdateExtractionRuleDefinition struct {
	value *UpdateExtractionRuleDefinition
	isSet bool
}

func (v NullableUpdateExtractionRuleDefinition) Get() *UpdateExtractionRuleDefinition {
	return v.value
}

func (v *NullableUpdateExtractionRuleDefinition) Set(val *UpdateExtractionRuleDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExtractionRuleDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExtractionRuleDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExtractionRuleDefinition(val *UpdateExtractionRuleDefinition) *NullableUpdateExtractionRuleDefinition {
	return &NullableUpdateExtractionRuleDefinition{value: val, isSet: true}
}

func (v NullableUpdateExtractionRuleDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExtractionRuleDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


