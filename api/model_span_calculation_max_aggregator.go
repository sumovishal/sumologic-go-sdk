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

// SpanCalculationMaxAggregator struct for SpanCalculationMaxAggregator
type SpanCalculationMaxAggregator struct {
	SpanCalculationAggregator
}

// NewSpanCalculationMaxAggregator instantiates a new SpanCalculationMaxAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanCalculationMaxAggregator(key string) *SpanCalculationMaxAggregator {
	this := SpanCalculationMaxAggregator{}
	this.Key = key
	return &this
}

// NewSpanCalculationMaxAggregatorWithDefaults instantiates a new SpanCalculationMaxAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanCalculationMaxAggregatorWithDefaults() *SpanCalculationMaxAggregator {
	this := SpanCalculationMaxAggregator{}
	return &this
}

func (o SpanCalculationMaxAggregator) MarshalJSON() ([]byte, error) {
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

type NullableSpanCalculationMaxAggregator struct {
	value *SpanCalculationMaxAggregator
	isSet bool
}

func (v NullableSpanCalculationMaxAggregator) Get() *SpanCalculationMaxAggregator {
	return v.value
}

func (v *NullableSpanCalculationMaxAggregator) Set(val *SpanCalculationMaxAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanCalculationMaxAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanCalculationMaxAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanCalculationMaxAggregator(val *SpanCalculationMaxAggregator) *NullableSpanCalculationMaxAggregator {
	return &NullableSpanCalculationMaxAggregator{value: val, isSet: true}
}

func (v NullableSpanCalculationMaxAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanCalculationMaxAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


