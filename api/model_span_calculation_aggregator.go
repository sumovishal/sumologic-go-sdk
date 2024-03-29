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

// checks if the SpanCalculationAggregator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanCalculationAggregator{}

// SpanCalculationAggregator struct for SpanCalculationAggregator
type SpanCalculationAggregator struct {
	// A specific aggregation type applied to spans.
	Key string `json:"key"`
}

// NewSpanCalculationAggregator instantiates a new SpanCalculationAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanCalculationAggregator(key string) *SpanCalculationAggregator {
	this := SpanCalculationAggregator{}
	this.Key = key
	return &this
}

// NewSpanCalculationAggregatorWithDefaults instantiates a new SpanCalculationAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanCalculationAggregatorWithDefaults() *SpanCalculationAggregator {
	this := SpanCalculationAggregator{}
	return &this
}

// GetKey returns the Key field value
func (o *SpanCalculationAggregator) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SpanCalculationAggregator) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SpanCalculationAggregator) SetKey(v string) {
	o.Key = v
}

func (o SpanCalculationAggregator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanCalculationAggregator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableSpanCalculationAggregator struct {
	value *SpanCalculationAggregator
	isSet bool
}

func (v NullableSpanCalculationAggregator) Get() *SpanCalculationAggregator {
	return v.value
}

func (v *NullableSpanCalculationAggregator) Set(val *SpanCalculationAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanCalculationAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanCalculationAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanCalculationAggregator(val *SpanCalculationAggregator) *NullableSpanCalculationAggregator {
	return &NullableSpanCalculationAggregator{value: val, isSet: true}
}

func (v NullableSpanCalculationAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanCalculationAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


