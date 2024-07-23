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

// checks if the DroppedField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DroppedField{}

// DroppedField struct for DroppedField
type DroppedField struct {
	// Field name.
	FieldName string `json:"fieldName"`
}

type _DroppedField DroppedField

// NewDroppedField instantiates a new DroppedField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDroppedField(fieldName string) *DroppedField {
	this := DroppedField{}
	this.FieldName = fieldName
	return &this
}

// NewDroppedFieldWithDefaults instantiates a new DroppedField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDroppedFieldWithDefaults() *DroppedField {
	this := DroppedField{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *DroppedField) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *DroppedField) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *DroppedField) SetFieldName(v string) {
	o.FieldName = v
}

func (o DroppedField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DroppedField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	return toSerialize, nil
}

func (o *DroppedField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fieldName",
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

	varDroppedField := _DroppedField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDroppedField)

	if err != nil {
		return err
	}

	*o = DroppedField(varDroppedField)

	return err
}

type NullableDroppedField struct {
	value *DroppedField
	isSet bool
}

func (v NullableDroppedField) Get() *DroppedField {
	return v.value
}

func (v *NullableDroppedField) Set(val *DroppedField) {
	v.value = val
	v.isSet = true
}

func (v NullableDroppedField) IsSet() bool {
	return v.isSet
}

func (v *NullableDroppedField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDroppedField(val *DroppedField) *NullableDroppedField {
	return &NullableDroppedField{value: val, isSet: true}
}

func (v NullableDroppedField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDroppedField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


