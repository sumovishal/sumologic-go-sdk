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

// checks if the ExtractionRuleDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtractionRuleDefinition{}

// ExtractionRuleDefinition struct for ExtractionRuleDefinition
type ExtractionRuleDefinition struct {
	// Name of the field extraction rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Describes the fields to be parsed.
	ParseExpression string `json:"parseExpression"`
	// Is the field extraction rule enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

type _ExtractionRuleDefinition ExtractionRuleDefinition

// NewExtractionRuleDefinition instantiates a new ExtractionRuleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractionRuleDefinition(name string, scope string, parseExpression string) *ExtractionRuleDefinition {
	this := ExtractionRuleDefinition{}
	this.Name = name
	this.Scope = scope
	this.ParseExpression = parseExpression
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewExtractionRuleDefinitionWithDefaults instantiates a new ExtractionRuleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractionRuleDefinitionWithDefaults() *ExtractionRuleDefinition {
	this := ExtractionRuleDefinition{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value
func (o *ExtractionRuleDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtractionRuleDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtractionRuleDefinition) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *ExtractionRuleDefinition) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ExtractionRuleDefinition) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *ExtractionRuleDefinition) SetScope(v string) {
	o.Scope = v
}

// GetParseExpression returns the ParseExpression field value
func (o *ExtractionRuleDefinition) GetParseExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParseExpression
}

// GetParseExpressionOk returns a tuple with the ParseExpression field value
// and a boolean to check if the value has been set.
func (o *ExtractionRuleDefinition) GetParseExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParseExpression, true
}

// SetParseExpression sets field value
func (o *ExtractionRuleDefinition) SetParseExpression(v string) {
	o.ParseExpression = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ExtractionRuleDefinition) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractionRuleDefinition) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ExtractionRuleDefinition) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ExtractionRuleDefinition) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ExtractionRuleDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtractionRuleDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["scope"] = o.Scope
	toSerialize["parseExpression"] = o.ParseExpression
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

func (o *ExtractionRuleDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"scope",
		"parseExpression",
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

	varExtractionRuleDefinition := _ExtractionRuleDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtractionRuleDefinition)

	if err != nil {
		return err
	}

	*o = ExtractionRuleDefinition(varExtractionRuleDefinition)

	return err
}

type NullableExtractionRuleDefinition struct {
	value *ExtractionRuleDefinition
	isSet bool
}

func (v NullableExtractionRuleDefinition) Get() *ExtractionRuleDefinition {
	return v.value
}

func (v *NullableExtractionRuleDefinition) Set(val *ExtractionRuleDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractionRuleDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractionRuleDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractionRuleDefinition(val *ExtractionRuleDefinition) *NullableExtractionRuleDefinition {
	return &NullableExtractionRuleDefinition{value: val, isSet: true}
}

func (v NullableExtractionRuleDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractionRuleDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


