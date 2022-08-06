/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// SpanCalculationSumAggregator struct for SpanCalculationSumAggregator
type SpanCalculationSumAggregator struct {
	SpanCalculationAggregator
}

// NewSpanCalculationSumAggregator instantiates a new SpanCalculationSumAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanCalculationSumAggregator(key string) *SpanCalculationSumAggregator {
	this := SpanCalculationSumAggregator{}
	this.Key = key
	return &this
}

// NewSpanCalculationSumAggregatorWithDefaults instantiates a new SpanCalculationSumAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanCalculationSumAggregatorWithDefaults() *SpanCalculationSumAggregator {
	this := SpanCalculationSumAggregator{}
	return &this
}

func (o SpanCalculationSumAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSpanCalculationAggregator, errSpanCalculationAggregator := json.Marshal(o.SpanCalculationAggregator)
	if errSpanCalculationAggregator != nil {
		return []byte{}, errSpanCalculationAggregator
	}
	errSpanCalculationAggregator = json.Unmarshal([]byte(serializedSpanCalculationAggregator), &toSerialize)
	if errSpanCalculationAggregator != nil {
		return []byte{}, errSpanCalculationAggregator
	}
	return json.Marshal(toSerialize)
}

type NullableSpanCalculationSumAggregator struct {
	value *SpanCalculationSumAggregator
	isSet bool
}

func (v NullableSpanCalculationSumAggregator) Get() *SpanCalculationSumAggregator {
	return v.value
}

func (v *NullableSpanCalculationSumAggregator) Set(val *SpanCalculationSumAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanCalculationSumAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanCalculationSumAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanCalculationSumAggregator(val *SpanCalculationSumAggregator) *NullableSpanCalculationSumAggregator {
	return &NullableSpanCalculationSumAggregator{value: val, isSet: true}
}

func (v NullableSpanCalculationSumAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanCalculationSumAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


