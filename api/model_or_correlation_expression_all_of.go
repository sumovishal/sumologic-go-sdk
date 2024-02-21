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

// checks if the OrCorrelationExpressionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrCorrelationExpressionAllOf{}

// OrCorrelationExpressionAllOf struct for OrCorrelationExpressionAllOf
type OrCorrelationExpressionAllOf struct {
	// List of correlation expressions to be evaluated with OR boolean operator.
	CorrelationExpressions []CorrelationExpression `json:"correlationExpressions"`
}

// NewOrCorrelationExpressionAllOf instantiates a new OrCorrelationExpressionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrCorrelationExpressionAllOf(correlationExpressions []CorrelationExpression) *OrCorrelationExpressionAllOf {
	this := OrCorrelationExpressionAllOf{}
	this.CorrelationExpressions = correlationExpressions
	return &this
}

// NewOrCorrelationExpressionAllOfWithDefaults instantiates a new OrCorrelationExpressionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrCorrelationExpressionAllOfWithDefaults() *OrCorrelationExpressionAllOf {
	this := OrCorrelationExpressionAllOf{}
	return &this
}

// GetCorrelationExpressions returns the CorrelationExpressions field value
func (o *OrCorrelationExpressionAllOf) GetCorrelationExpressions() []CorrelationExpression {
	if o == nil {
		var ret []CorrelationExpression
		return ret
	}

	return o.CorrelationExpressions
}

// GetCorrelationExpressionsOk returns a tuple with the CorrelationExpressions field value
// and a boolean to check if the value has been set.
func (o *OrCorrelationExpressionAllOf) GetCorrelationExpressionsOk() ([]CorrelationExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.CorrelationExpressions, true
}

// SetCorrelationExpressions sets field value
func (o *OrCorrelationExpressionAllOf) SetCorrelationExpressions(v []CorrelationExpression) {
	o.CorrelationExpressions = v
}

func (o OrCorrelationExpressionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrCorrelationExpressionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["correlationExpressions"] = o.CorrelationExpressions
	return toSerialize, nil
}

type NullableOrCorrelationExpressionAllOf struct {
	value *OrCorrelationExpressionAllOf
	isSet bool
}

func (v NullableOrCorrelationExpressionAllOf) Get() *OrCorrelationExpressionAllOf {
	return v.value
}

func (v *NullableOrCorrelationExpressionAllOf) Set(val *OrCorrelationExpressionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrCorrelationExpressionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrCorrelationExpressionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrCorrelationExpressionAllOf(val *OrCorrelationExpressionAllOf) *NullableOrCorrelationExpressionAllOf {
	return &NullableOrCorrelationExpressionAllOf{value: val, isSet: true}
}

func (v NullableOrCorrelationExpressionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrCorrelationExpressionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


