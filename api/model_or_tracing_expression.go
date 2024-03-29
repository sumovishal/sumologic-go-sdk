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

// checks if the OrTracingExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrTracingExpression{}

// OrTracingExpression Evaluates to true, if at least one expression evaluates to true, otherwise evaluates to false.
type OrTracingExpression struct {
	TraceQueryExpression
	Expressions []TraceQueryExpression `json:"expressions"`
}

// NewOrTracingExpression instantiates a new OrTracingExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrTracingExpression(expressions []TraceQueryExpression, type_ string) *OrTracingExpression {
	this := OrTracingExpression{}
	this.Type = type_
	this.Expressions = expressions
	return &this
}

// NewOrTracingExpressionWithDefaults instantiates a new OrTracingExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrTracingExpressionWithDefaults() *OrTracingExpression {
	this := OrTracingExpression{}
	return &this
}

// GetExpressions returns the Expressions field value
func (o *OrTracingExpression) GetExpressions() []TraceQueryExpression {
	if o == nil {
		var ret []TraceQueryExpression
		return ret
	}

	return o.Expressions
}

// GetExpressionsOk returns a tuple with the Expressions field value
// and a boolean to check if the value has been set.
func (o *OrTracingExpression) GetExpressionsOk() ([]TraceQueryExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expressions, true
}

// SetExpressions sets field value
func (o *OrTracingExpression) SetExpressions(v []TraceQueryExpression) {
	o.Expressions = v
}

func (o OrTracingExpression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrTracingExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTraceQueryExpression, errTraceQueryExpression := json.Marshal(o.TraceQueryExpression)
	if errTraceQueryExpression != nil {
		return map[string]interface{}{}, errTraceQueryExpression
	}
	errTraceQueryExpression = json.Unmarshal([]byte(serializedTraceQueryExpression), &toSerialize)
	if errTraceQueryExpression != nil {
		return map[string]interface{}{}, errTraceQueryExpression
	}
	toSerialize["expressions"] = o.Expressions
	return toSerialize, nil
}

type NullableOrTracingExpression struct {
	value *OrTracingExpression
	isSet bool
}

func (v NullableOrTracingExpression) Get() *OrTracingExpression {
	return v.value
}

func (v *NullableOrTracingExpression) Set(val *OrTracingExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableOrTracingExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableOrTracingExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrTracingExpression(val *OrTracingExpression) *NullableOrTracingExpression {
	return &NullableOrTracingExpression{value: val, isSet: true}
}

func (v NullableOrTracingExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrTracingExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


