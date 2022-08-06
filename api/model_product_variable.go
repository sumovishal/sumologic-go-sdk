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

// ProductVariable Details of product variable and its quantity.
type ProductVariable struct {
	// Name of a product variable.
	ProductVariableName string `json:"productVariableName"`
	// Unique Identifier of the product variable.
	ProductVariableId string `json:"productVariableId"`
	// Unit of measure for the productvariable.
	Unit string `json:"unit"`
	// Possible values allowed for the productvariable.
	PossibleValues []int64 `json:"possibleValues"`
}

// NewProductVariable instantiates a new ProductVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductVariable(productVariableName string, productVariableId string, unit string, possibleValues []int64) *ProductVariable {
	this := ProductVariable{}
	this.ProductVariableName = productVariableName
	this.ProductVariableId = productVariableId
	this.Unit = unit
	this.PossibleValues = possibleValues
	return &this
}

// NewProductVariableWithDefaults instantiates a new ProductVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductVariableWithDefaults() *ProductVariable {
	this := ProductVariable{}
	return &this
}

// GetProductVariableName returns the ProductVariableName field value
func (o *ProductVariable) GetProductVariableName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVariableName
}

// GetProductVariableNameOk returns a tuple with the ProductVariableName field value
// and a boolean to check if the value has been set.
func (o *ProductVariable) GetProductVariableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVariableName, true
}

// SetProductVariableName sets field value
func (o *ProductVariable) SetProductVariableName(v string) {
	o.ProductVariableName = v
}

// GetProductVariableId returns the ProductVariableId field value
func (o *ProductVariable) GetProductVariableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVariableId
}

// GetProductVariableIdOk returns a tuple with the ProductVariableId field value
// and a boolean to check if the value has been set.
func (o *ProductVariable) GetProductVariableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVariableId, true
}

// SetProductVariableId sets field value
func (o *ProductVariable) SetProductVariableId(v string) {
	o.ProductVariableId = v
}

// GetUnit returns the Unit field value
func (o *ProductVariable) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ProductVariable) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ProductVariable) SetUnit(v string) {
	o.Unit = v
}

// GetPossibleValues returns the PossibleValues field value
func (o *ProductVariable) GetPossibleValues() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.PossibleValues
}

// GetPossibleValuesOk returns a tuple with the PossibleValues field value
// and a boolean to check if the value has been set.
func (o *ProductVariable) GetPossibleValuesOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PossibleValues, true
}

// SetPossibleValues sets field value
func (o *ProductVariable) SetPossibleValues(v []int64) {
	o.PossibleValues = v
}

func (o ProductVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["productVariableName"] = o.ProductVariableName
	}
	if true {
		toSerialize["productVariableId"] = o.ProductVariableId
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["possibleValues"] = o.PossibleValues
	}
	return json.Marshal(toSerialize)
}

type NullableProductVariable struct {
	value *ProductVariable
	isSet bool
}

func (v NullableProductVariable) Get() *ProductVariable {
	return v.value
}

func (v *NullableProductVariable) Set(val *ProductVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableProductVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableProductVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductVariable(val *ProductVariable) *NullableProductVariable {
	return &NullableProductVariable{value: val, isSet: true}
}

func (v NullableProductVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


