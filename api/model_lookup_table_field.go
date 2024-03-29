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

// checks if the LookupTableField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupTableField{}

// LookupTableField The definition of the field.
type LookupTableField struct {
	// The name of the field.
	FieldName string `json:"fieldName"`
	// The data type of the field. Supported types:   - `boolean`   - `int`   - `long`   - `double`   - `string`
	FieldType string `json:"fieldType"`
}

// NewLookupTableField instantiates a new LookupTableField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupTableField(fieldName string, fieldType string) *LookupTableField {
	this := LookupTableField{}
	this.FieldName = fieldName
	this.FieldType = fieldType
	return &this
}

// NewLookupTableFieldWithDefaults instantiates a new LookupTableField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupTableFieldWithDefaults() *LookupTableField {
	this := LookupTableField{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *LookupTableField) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *LookupTableField) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *LookupTableField) SetFieldName(v string) {
	o.FieldName = v
}

// GetFieldType returns the FieldType field value
func (o *LookupTableField) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *LookupTableField) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *LookupTableField) SetFieldType(v string) {
	o.FieldType = v
}

func (o LookupTableField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupTableField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	toSerialize["fieldType"] = o.FieldType
	return toSerialize, nil
}

type NullableLookupTableField struct {
	value *LookupTableField
	isSet bool
}

func (v NullableLookupTableField) Get() *LookupTableField {
	return v.value
}

func (v *NullableLookupTableField) Set(val *LookupTableField) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupTableField) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupTableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupTableField(val *LookupTableField) *NullableLookupTableField {
	return &NullableLookupTableField{value: val, isSet: true}
}

func (v NullableLookupTableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupTableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


