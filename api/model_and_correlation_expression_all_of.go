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

// checks if the AndCorrelationExpressionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AndCorrelationExpressionAllOf{}

// AndCorrelationExpressionAllOf struct for AndCorrelationExpressionAllOf
type AndCorrelationExpressionAllOf struct {
	// List of correlation expressions to be evaluated with AND boolean operator.
	CorrelationExpressions []CorrelationExpression `json:"correlationExpressions"`
}

// NewAndCorrelationExpressionAllOf instantiates a new AndCorrelationExpressionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAndCorrelationExpressionAllOf(correlationExpressions []CorrelationExpression) *AndCorrelationExpressionAllOf {
	this := AndCorrelationExpressionAllOf{}
	this.CorrelationExpressions = correlationExpressions
	return &this
}

// NewAndCorrelationExpressionAllOfWithDefaults instantiates a new AndCorrelationExpressionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAndCorrelationExpressionAllOfWithDefaults() *AndCorrelationExpressionAllOf {
	this := AndCorrelationExpressionAllOf{}
	return &this
}

// GetCorrelationExpressions returns the CorrelationExpressions field value
func (o *AndCorrelationExpressionAllOf) GetCorrelationExpressions() []CorrelationExpression {
	if o == nil {
		var ret []CorrelationExpression
		return ret
	}

	return o.CorrelationExpressions
}

// GetCorrelationExpressionsOk returns a tuple with the CorrelationExpressions field value
// and a boolean to check if the value has been set.
func (o *AndCorrelationExpressionAllOf) GetCorrelationExpressionsOk() ([]CorrelationExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.CorrelationExpressions, true
}

// SetCorrelationExpressions sets field value
func (o *AndCorrelationExpressionAllOf) SetCorrelationExpressions(v []CorrelationExpression) {
	o.CorrelationExpressions = v
}

func (o AndCorrelationExpressionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AndCorrelationExpressionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["correlationExpressions"] = o.CorrelationExpressions
	return toSerialize, nil
}

type NullableAndCorrelationExpressionAllOf struct {
	value *AndCorrelationExpressionAllOf
	isSet bool
}

func (v NullableAndCorrelationExpressionAllOf) Get() *AndCorrelationExpressionAllOf {
	return v.value
}

func (v *NullableAndCorrelationExpressionAllOf) Set(val *AndCorrelationExpressionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAndCorrelationExpressionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAndCorrelationExpressionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAndCorrelationExpressionAllOf(val *AndCorrelationExpressionAllOf) *NullableAndCorrelationExpressionAllOf {
	return &NullableAndCorrelationExpressionAllOf{value: val, isSet: true}
}

func (v NullableAndCorrelationExpressionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAndCorrelationExpressionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


