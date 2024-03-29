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

// checks if the TransformationRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformationRuleRequest{}

// TransformationRuleRequest A request for creating or updating a transformation rule.
type TransformationRuleRequest struct {
	RuleDefinition TransformationRuleDefinition `json:"ruleDefinition"`
	// True if the rule is enabled.
	Enabled bool `json:"enabled"`
}

// NewTransformationRuleRequest instantiates a new TransformationRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformationRuleRequest(ruleDefinition TransformationRuleDefinition, enabled bool) *TransformationRuleRequest {
	this := TransformationRuleRequest{}
	this.RuleDefinition = ruleDefinition
	this.Enabled = enabled
	return &this
}

// NewTransformationRuleRequestWithDefaults instantiates a new TransformationRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformationRuleRequestWithDefaults() *TransformationRuleRequest {
	this := TransformationRuleRequest{}
	return &this
}

// GetRuleDefinition returns the RuleDefinition field value
func (o *TransformationRuleRequest) GetRuleDefinition() TransformationRuleDefinition {
	if o == nil {
		var ret TransformationRuleDefinition
		return ret
	}

	return o.RuleDefinition
}

// GetRuleDefinitionOk returns a tuple with the RuleDefinition field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleRequest) GetRuleDefinitionOk() (*TransformationRuleDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleDefinition, true
}

// SetRuleDefinition sets field value
func (o *TransformationRuleRequest) SetRuleDefinition(v TransformationRuleDefinition) {
	o.RuleDefinition = v
}

// GetEnabled returns the Enabled field value
func (o *TransformationRuleRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TransformationRuleRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o TransformationRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformationRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ruleDefinition"] = o.RuleDefinition
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableTransformationRuleRequest struct {
	value *TransformationRuleRequest
	isSet bool
}

func (v NullableTransformationRuleRequest) Get() *TransformationRuleRequest {
	return v.value
}

func (v *NullableTransformationRuleRequest) Set(val *TransformationRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationRuleRequest(val *TransformationRuleRequest) *NullableTransformationRuleRequest {
	return &NullableTransformationRuleRequest{value: val, isSet: true}
}

func (v NullableTransformationRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


