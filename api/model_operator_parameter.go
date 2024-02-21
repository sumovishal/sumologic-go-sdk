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

// checks if the OperatorParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatorParameter{}

// OperatorParameter The operator parameter for operator data.
type OperatorParameter struct {
	// The key of the operator parameter.
	Key string `json:"key"`
	// The value of the operator parameter.
	Value string `json:"value"`
}

// NewOperatorParameter instantiates a new OperatorParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorParameter(key string, value string) *OperatorParameter {
	this := OperatorParameter{}
	this.Key = key
	this.Value = value
	return &this
}

// NewOperatorParameterWithDefaults instantiates a new OperatorParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorParameterWithDefaults() *OperatorParameter {
	this := OperatorParameter{}
	return &this
}

// GetKey returns the Key field value
func (o *OperatorParameter) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *OperatorParameter) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *OperatorParameter) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *OperatorParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OperatorParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OperatorParameter) SetValue(v string) {
	o.Value = v
}

func (o OperatorParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatorParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableOperatorParameter struct {
	value *OperatorParameter
	isSet bool
}

func (v NullableOperatorParameter) Get() *OperatorParameter {
	return v.value
}

func (v *NullableOperatorParameter) Set(val *OperatorParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorParameter(val *OperatorParameter) *NullableOperatorParameter {
	return &NullableOperatorParameter{value: val, isSet: true}
}

func (v NullableOperatorParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


