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

// DimensionKeyValue The key and value pair for each metric dimension.
type DimensionKeyValue struct {
	// The key of the metric dimension.
	Key *string `json:"key,omitempty"`
	// The value of the metric dimension.
	Value *string `json:"value,omitempty"`
}

// NewDimensionKeyValue instantiates a new DimensionKeyValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionKeyValue() *DimensionKeyValue {
	this := DimensionKeyValue{}
	return &this
}

// NewDimensionKeyValueWithDefaults instantiates a new DimensionKeyValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionKeyValueWithDefaults() *DimensionKeyValue {
	this := DimensionKeyValue{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DimensionKeyValue) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionKeyValue) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DimensionKeyValue) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DimensionKeyValue) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DimensionKeyValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionKeyValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DimensionKeyValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DimensionKeyValue) SetValue(v string) {
	o.Value = &v
}

func (o DimensionKeyValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDimensionKeyValue struct {
	value *DimensionKeyValue
	isSet bool
}

func (v NullableDimensionKeyValue) Get() *DimensionKeyValue {
	return v.value
}

func (v *NullableDimensionKeyValue) Set(val *DimensionKeyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionKeyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionKeyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionKeyValue(val *DimensionKeyValue) *NullableDimensionKeyValue {
	return &NullableDimensionKeyValue{value: val, isSet: true}
}

func (v NullableDimensionKeyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimensionKeyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


