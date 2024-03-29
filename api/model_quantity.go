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

// checks if the Quantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Quantity{}

// Quantity Details of unit of consumption and its value.
type Quantity struct {
	// The value of the consumable in units.
	Value int64 `json:"value"`
	// The unit of the consumable. Units are provided in: 1. `GB` 2. `DPM`(Data Points Per Minute) 3. `Credits` 4. `Days` 
	Unit string `json:"unit"`
}

// NewQuantity instantiates a new Quantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuantity(value int64, unit string) *Quantity {
	this := Quantity{}
	this.Value = value
	this.Unit = unit
	return &this
}

// NewQuantityWithDefaults instantiates a new Quantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuantityWithDefaults() *Quantity {
	this := Quantity{}
	return &this
}

// GetValue returns the Value field value
func (o *Quantity) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Quantity) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Quantity) SetValue(v int64) {
	o.Value = v
}

// GetUnit returns the Unit field value
func (o *Quantity) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Quantity) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Quantity) SetUnit(v string) {
	o.Unit = v
}

func (o Quantity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Quantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["unit"] = o.Unit
	return toSerialize, nil
}

type NullableQuantity struct {
	value *Quantity
	isSet bool
}

func (v NullableQuantity) Get() *Quantity {
	return v.value
}

func (v *NullableQuantity) Set(val *Quantity) {
	v.value = val
	v.isSet = true
}

func (v NullableQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullableQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuantity(val *Quantity) *NullableQuantity {
	return &NullableQuantity{value: val, isSet: true}
}

func (v NullableQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


