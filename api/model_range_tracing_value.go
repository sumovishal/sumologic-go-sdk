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

// checks if the RangeTracingValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeTracingValue{}

// RangeTracingValue struct for RangeTracingValue
type RangeTracingValue struct {
	TracingValue
	From TracingValue `json:"from"`
	To TracingValue `json:"to"`
}

// NewRangeTracingValue instantiates a new RangeTracingValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeTracingValue(from TracingValue, to TracingValue, type_ string) *RangeTracingValue {
	this := RangeTracingValue{}
	this.Type = type_
	this.From = from
	this.To = to
	return &this
}

// NewRangeTracingValueWithDefaults instantiates a new RangeTracingValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeTracingValueWithDefaults() *RangeTracingValue {
	this := RangeTracingValue{}
	return &this
}

// GetFrom returns the From field value
func (o *RangeTracingValue) GetFrom() TracingValue {
	if o == nil {
		var ret TracingValue
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *RangeTracingValue) GetFromOk() (*TracingValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *RangeTracingValue) SetFrom(v TracingValue) {
	o.From = v
}

// GetTo returns the To field value
func (o *RangeTracingValue) GetTo() TracingValue {
	if o == nil {
		var ret TracingValue
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *RangeTracingValue) GetToOk() (*TracingValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *RangeTracingValue) SetTo(v TracingValue) {
	o.To = v
}

func (o RangeTracingValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RangeTracingValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTracingValue, errTracingValue := json.Marshal(o.TracingValue)
	if errTracingValue != nil {
		return map[string]interface{}{}, errTracingValue
	}
	errTracingValue = json.Unmarshal([]byte(serializedTracingValue), &toSerialize)
	if errTracingValue != nil {
		return map[string]interface{}{}, errTracingValue
	}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	return toSerialize, nil
}

type NullableRangeTracingValue struct {
	value *RangeTracingValue
	isSet bool
}

func (v NullableRangeTracingValue) Get() *RangeTracingValue {
	return v.value
}

func (v *NullableRangeTracingValue) Set(val *RangeTracingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeTracingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeTracingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeTracingValue(val *RangeTracingValue) *NullableRangeTracingValue {
	return &NullableRangeTracingValue{value: val, isSet: true}
}

func (v NullableRangeTracingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeTracingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


