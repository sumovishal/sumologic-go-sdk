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

// checks if the SpanCalculationMinAggregator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanCalculationMinAggregator{}

// SpanCalculationMinAggregator struct for SpanCalculationMinAggregator
type SpanCalculationMinAggregator struct {
	SpanCalculationAggregator
}

// NewSpanCalculationMinAggregator instantiates a new SpanCalculationMinAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanCalculationMinAggregator(key string) *SpanCalculationMinAggregator {
	this := SpanCalculationMinAggregator{}
	this.Key = key
	return &this
}

// NewSpanCalculationMinAggregatorWithDefaults instantiates a new SpanCalculationMinAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanCalculationMinAggregatorWithDefaults() *SpanCalculationMinAggregator {
	this := SpanCalculationMinAggregator{}
	return &this
}

func (o SpanCalculationMinAggregator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanCalculationMinAggregator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSpanCalculationAggregator, errSpanCalculationAggregator := json.Marshal(o.SpanCalculationAggregator)
	if errSpanCalculationAggregator != nil {
		return map[string]interface{}{}, errSpanCalculationAggregator
	}
	errSpanCalculationAggregator = json.Unmarshal([]byte(serializedSpanCalculationAggregator), &toSerialize)
	if errSpanCalculationAggregator != nil {
		return map[string]interface{}{}, errSpanCalculationAggregator
	}
	return toSerialize, nil
}

type NullableSpanCalculationMinAggregator struct {
	value *SpanCalculationMinAggregator
	isSet bool
}

func (v NullableSpanCalculationMinAggregator) Get() *SpanCalculationMinAggregator {
	return v.value
}

func (v *NullableSpanCalculationMinAggregator) Set(val *SpanCalculationMinAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanCalculationMinAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanCalculationMinAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanCalculationMinAggregator(val *SpanCalculationMinAggregator) *NullableSpanCalculationMinAggregator {
	return &NullableSpanCalculationMinAggregator{value: val, isSet: true}
}

func (v NullableSpanCalculationMinAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanCalculationMinAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


