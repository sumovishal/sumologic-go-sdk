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

// checks if the BooleanEventAttributeValueAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BooleanEventAttributeValueAllOf{}

// BooleanEventAttributeValueAllOf struct for BooleanEventAttributeValueAllOf
type BooleanEventAttributeValueAllOf struct {
	Value bool `json:"value"`
}

// NewBooleanEventAttributeValueAllOf instantiates a new BooleanEventAttributeValueAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBooleanEventAttributeValueAllOf(value bool) *BooleanEventAttributeValueAllOf {
	this := BooleanEventAttributeValueAllOf{}
	this.Value = value
	return &this
}

// NewBooleanEventAttributeValueAllOfWithDefaults instantiates a new BooleanEventAttributeValueAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBooleanEventAttributeValueAllOfWithDefaults() *BooleanEventAttributeValueAllOf {
	this := BooleanEventAttributeValueAllOf{}
	return &this
}

// GetValue returns the Value field value
func (o *BooleanEventAttributeValueAllOf) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BooleanEventAttributeValueAllOf) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BooleanEventAttributeValueAllOf) SetValue(v bool) {
	o.Value = v
}

func (o BooleanEventAttributeValueAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BooleanEventAttributeValueAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableBooleanEventAttributeValueAllOf struct {
	value *BooleanEventAttributeValueAllOf
	isSet bool
}

func (v NullableBooleanEventAttributeValueAllOf) Get() *BooleanEventAttributeValueAllOf {
	return v.value
}

func (v *NullableBooleanEventAttributeValueAllOf) Set(val *BooleanEventAttributeValueAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBooleanEventAttributeValueAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBooleanEventAttributeValueAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooleanEventAttributeValueAllOf(val *BooleanEventAttributeValueAllOf) *NullableBooleanEventAttributeValueAllOf {
	return &NullableBooleanEventAttributeValueAllOf{value: val, isSet: true}
}

func (v NullableBooleanEventAttributeValueAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooleanEventAttributeValueAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


