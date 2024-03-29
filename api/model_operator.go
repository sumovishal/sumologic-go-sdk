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

// checks if the Operator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Operator{}

// Operator Result of the aggregations performed on the usages. Operator can be `sum`, `average`, `usagePercentage`, `forecastValue`,`forecastPercentage`, or `forecastRemainingDays`.
type Operator struct {
	// An array of objects denoting the value and unit of the results.
	Values []DataValue `json:"values"`
	// The name of the operator applied to the data.
	Name string `json:"name"`
}

// NewOperator instantiates a new Operator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperator(values []DataValue, name string) *Operator {
	this := Operator{}
	this.Values = values
	this.Name = name
	return &this
}

// NewOperatorWithDefaults instantiates a new Operator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorWithDefaults() *Operator {
	this := Operator{}
	return &this
}

// GetValues returns the Values field value
func (o *Operator) GetValues() []DataValue {
	if o == nil {
		var ret []DataValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Operator) GetValuesOk() ([]DataValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *Operator) SetValues(v []DataValue) {
	o.Values = v
}

// GetName returns the Name field value
func (o *Operator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Operator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Operator) SetName(v string) {
	o.Name = v
}

func (o Operator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Operator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableOperator struct {
	value *Operator
	isSet bool
}

func (v NullableOperator) Get() *Operator {
	return v.value
}

func (v *NullableOperator) Set(val *Operator) {
	v.value = val
	v.isSet = true
}

func (v NullableOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperator(val *Operator) *NullableOperator {
	return &NullableOperator{value: val, isSet: true}
}

func (v NullableOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


