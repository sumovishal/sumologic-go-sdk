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

// checks if the TransformationRuleDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformationRuleDefinition{}

// TransformationRuleDefinition The properties that define a transformation rule.
type TransformationRuleDefinition struct {
	// Name of the transformation rule.
	Name string `json:"name"`
	// Selector of the transformation rule.
	Selector string `json:"selector"`
	// Dimension transformations of the transformation rule.
	DimensionTransformations []DimensionTransformation `json:"dimensionTransformations,omitempty"`
	// Retention period in days for the transformed metrics that are generated by this rule. The supported retention periods for transformed metrics are 8 days, and 400 days. If no dimension transformations are defined, this value will be set to 0.
	TransformedMetricsRetention *int64 `json:"transformedMetricsRetention,omitempty"`
	// Retention period in days for the metrics that are selected by the selector. The supported retention periods for selected metrics are 8 days, 400 days, and 0 (Do not store) if this rule contains dimension transformation.
	Retention int64 `json:"retention"`
}

type _TransformationRuleDefinition TransformationRuleDefinition

// NewTransformationRuleDefinition instantiates a new TransformationRuleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformationRuleDefinition(name string, selector string, retention int64) *TransformationRuleDefinition {
	this := TransformationRuleDefinition{}
	this.Name = name
	this.Selector = selector
	var transformedMetricsRetention int64 = 0
	this.TransformedMetricsRetention = &transformedMetricsRetention
	this.Retention = retention
	return &this
}

// NewTransformationRuleDefinitionWithDefaults instantiates a new TransformationRuleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformationRuleDefinitionWithDefaults() *TransformationRuleDefinition {
	this := TransformationRuleDefinition{}
	var transformedMetricsRetention int64 = 0
	this.TransformedMetricsRetention = &transformedMetricsRetention
	var retention int64 = 400
	this.Retention = retention
	return &this
}

// GetName returns the Name field value
func (o *TransformationRuleDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TransformationRuleDefinition) SetName(v string) {
	o.Name = v
}

// GetSelector returns the Selector field value
func (o *TransformationRuleDefinition) GetSelector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleDefinition) GetSelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *TransformationRuleDefinition) SetSelector(v string) {
	o.Selector = v
}

// GetDimensionTransformations returns the DimensionTransformations field value if set, zero value otherwise.
func (o *TransformationRuleDefinition) GetDimensionTransformations() []DimensionTransformation {
	if o == nil || IsNil(o.DimensionTransformations) {
		var ret []DimensionTransformation
		return ret
	}
	return o.DimensionTransformations
}

// GetDimensionTransformationsOk returns a tuple with the DimensionTransformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationRuleDefinition) GetDimensionTransformationsOk() ([]DimensionTransformation, bool) {
	if o == nil || IsNil(o.DimensionTransformations) {
		return nil, false
	}
	return o.DimensionTransformations, true
}

// HasDimensionTransformations returns a boolean if a field has been set.
func (o *TransformationRuleDefinition) HasDimensionTransformations() bool {
	if o != nil && !IsNil(o.DimensionTransformations) {
		return true
	}

	return false
}

// SetDimensionTransformations gets a reference to the given []DimensionTransformation and assigns it to the DimensionTransformations field.
func (o *TransformationRuleDefinition) SetDimensionTransformations(v []DimensionTransformation) {
	o.DimensionTransformations = v
}

// GetTransformedMetricsRetention returns the TransformedMetricsRetention field value if set, zero value otherwise.
func (o *TransformationRuleDefinition) GetTransformedMetricsRetention() int64 {
	if o == nil || IsNil(o.TransformedMetricsRetention) {
		var ret int64
		return ret
	}
	return *o.TransformedMetricsRetention
}

// GetTransformedMetricsRetentionOk returns a tuple with the TransformedMetricsRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationRuleDefinition) GetTransformedMetricsRetentionOk() (*int64, bool) {
	if o == nil || IsNil(o.TransformedMetricsRetention) {
		return nil, false
	}
	return o.TransformedMetricsRetention, true
}

// HasTransformedMetricsRetention returns a boolean if a field has been set.
func (o *TransformationRuleDefinition) HasTransformedMetricsRetention() bool {
	if o != nil && !IsNil(o.TransformedMetricsRetention) {
		return true
	}

	return false
}

// SetTransformedMetricsRetention gets a reference to the given int64 and assigns it to the TransformedMetricsRetention field.
func (o *TransformationRuleDefinition) SetTransformedMetricsRetention(v int64) {
	o.TransformedMetricsRetention = &v
}

// GetRetention returns the Retention field value
func (o *TransformationRuleDefinition) GetRetention() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleDefinition) GetRetentionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *TransformationRuleDefinition) SetRetention(v int64) {
	o.Retention = v
}

func (o TransformationRuleDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformationRuleDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["selector"] = o.Selector
	if !IsNil(o.DimensionTransformations) {
		toSerialize["dimensionTransformations"] = o.DimensionTransformations
	}
	if !IsNil(o.TransformedMetricsRetention) {
		toSerialize["transformedMetricsRetention"] = o.TransformedMetricsRetention
	}
	toSerialize["retention"] = o.Retention
	return toSerialize, nil
}

func (o *TransformationRuleDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"selector",
		"retention",
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

	varTransformationRuleDefinition := _TransformationRuleDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransformationRuleDefinition)

	if err != nil {
		return err
	}

	*o = TransformationRuleDefinition(varTransformationRuleDefinition)

	return err
}

type NullableTransformationRuleDefinition struct {
	value *TransformationRuleDefinition
	isSet bool
}

func (v NullableTransformationRuleDefinition) Get() *TransformationRuleDefinition {
	return v.value
}

func (v *NullableTransformationRuleDefinition) Set(val *TransformationRuleDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationRuleDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationRuleDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationRuleDefinition(val *TransformationRuleDefinition) *NullableTransformationRuleDefinition {
	return &NullableTransformationRuleDefinition{value: val, isSet: true}
}

func (v NullableTransformationRuleDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationRuleDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


