/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderBy{}

// OrderBy struct for OrderBy
type OrderBy struct {
	// Field based on which results should be sorted. When not provided, the default behavior is to sort by timestamp descending. Sortable fields values: `trace_id`, `start_timestamp`, `duration`, `spans_number`, `errors`, `status_code`.
	FieldName string `json:"fieldName"`
	// Type of sorting values - descending or ascending.
	Order string `json:"order" validate:"regexp=^(Asc|Desc)$"`
}

type _OrderBy OrderBy

// NewOrderBy instantiates a new OrderBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBy(fieldName string, order string) *OrderBy {
	this := OrderBy{}
	this.FieldName = fieldName
	this.Order = order
	return &this
}

// NewOrderByWithDefaults instantiates a new OrderBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderByWithDefaults() *OrderBy {
	this := OrderBy{}
	var order string = "Desc"
	this.Order = order
	return &this
}

// GetFieldName returns the FieldName field value
func (o *OrderBy) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *OrderBy) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *OrderBy) SetFieldName(v string) {
	o.FieldName = v
}

// GetOrder returns the Order field value
func (o *OrderBy) GetOrder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *OrderBy) GetOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *OrderBy) SetOrder(v string) {
	o.Order = v
}

func (o OrderBy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

func (o *OrderBy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fieldName",
		"order",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderBy := _OrderBy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderBy)

	if err != nil {
		return err
	}

	*o = OrderBy(varOrderBy)

	return err
}

type NullableOrderBy struct {
	value *OrderBy
	isSet bool
}

func (v NullableOrderBy) Get() *OrderBy {
	return v.value
}

func (v *NullableOrderBy) Set(val *OrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBy(val *OrderBy) *NullableOrderBy {
	return &NullableOrderBy{value: val, isSet: true}
}

func (v NullableOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


