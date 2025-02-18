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

// checks if the DynamicRuleDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicRuleDefinition{}

// DynamicRuleDefinition struct for DynamicRuleDefinition
type DynamicRuleDefinition struct {
	// Name of the dynamic parsing rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the dynamic parsing rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Is the dynamic parsing rule enabled.
	Enabled bool `json:"enabled"`
}

type _DynamicRuleDefinition DynamicRuleDefinition

// NewDynamicRuleDefinition instantiates a new DynamicRuleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicRuleDefinition(name string, scope string, enabled bool) *DynamicRuleDefinition {
	this := DynamicRuleDefinition{}
	this.Name = name
	this.Scope = scope
	this.Enabled = enabled
	return &this
}

// NewDynamicRuleDefinitionWithDefaults instantiates a new DynamicRuleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicRuleDefinitionWithDefaults() *DynamicRuleDefinition {
	this := DynamicRuleDefinition{}
	var enabled bool = true
	this.Enabled = enabled
	return &this
}

// GetName returns the Name field value
func (o *DynamicRuleDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicRuleDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicRuleDefinition) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *DynamicRuleDefinition) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *DynamicRuleDefinition) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *DynamicRuleDefinition) SetScope(v string) {
	o.Scope = v
}

// GetEnabled returns the Enabled field value
func (o *DynamicRuleDefinition) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DynamicRuleDefinition) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DynamicRuleDefinition) SetEnabled(v bool) {
	o.Enabled = v
}

func (o DynamicRuleDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicRuleDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["scope"] = o.Scope
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *DynamicRuleDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"scope",
		"enabled",
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

	varDynamicRuleDefinition := _DynamicRuleDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDynamicRuleDefinition)

	if err != nil {
		return err
	}

	*o = DynamicRuleDefinition(varDynamicRuleDefinition)

	return err
}

type NullableDynamicRuleDefinition struct {
	value *DynamicRuleDefinition
	isSet bool
}

func (v NullableDynamicRuleDefinition) Get() *DynamicRuleDefinition {
	return v.value
}

func (v *NullableDynamicRuleDefinition) Set(val *DynamicRuleDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicRuleDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicRuleDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicRuleDefinition(val *DynamicRuleDefinition) *NullableDynamicRuleDefinition {
	return &NullableDynamicRuleDefinition{value: val, isSet: true}
}

func (v NullableDynamicRuleDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicRuleDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


