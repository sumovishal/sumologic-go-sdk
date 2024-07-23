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

// checks if the CustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomField{}

// CustomField struct for CustomField
type CustomField struct {
	// Field name.
	FieldName string `json:"fieldName"`
	// Identifier of the field.
	FieldId string `json:"fieldId"`
	// Field type. Possible values are `String`, `Long`, `Int`, `Double`, and `Boolean`.
	DataType string `json:"dataType" validate:"regexp=^(String|Long|Int|Double|Boolean)$"`
	// Indicates whether the field is enabled and its values are being accepted. Possible values are `Enabled` and `Disabled`.
	State string `json:"state" validate:"regexp=^(Enabled|Disabled)$"`
}

type _CustomField CustomField

// NewCustomField instantiates a new CustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomField(fieldName string, fieldId string, dataType string, state string) *CustomField {
	this := CustomField{}
	this.FieldName = fieldName
	this.FieldId = fieldId
	this.DataType = dataType
	this.State = state
	return &this
}

// NewCustomFieldWithDefaults instantiates a new CustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldWithDefaults() *CustomField {
	this := CustomField{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *CustomField) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *CustomField) SetFieldName(v string) {
	o.FieldName = v
}

// GetFieldId returns the FieldId field value
func (o *CustomField) GetFieldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetFieldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldId, true
}

// SetFieldId sets field value
func (o *CustomField) SetFieldId(v string) {
	o.FieldId = v
}

// GetDataType returns the DataType field value
func (o *CustomField) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *CustomField) SetDataType(v string) {
	o.DataType = v
}

// GetState returns the State field value
func (o *CustomField) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CustomField) SetState(v string) {
	o.State = v
}

func (o CustomField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	toSerialize["fieldId"] = o.FieldId
	toSerialize["dataType"] = o.DataType
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *CustomField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fieldName",
		"fieldId",
		"dataType",
		"state",
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

	varCustomField := _CustomField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomField)

	if err != nil {
		return err
	}

	*o = CustomField(varCustomField)

	return err
}

type NullableCustomField struct {
	value *CustomField
	isSet bool
}

func (v NullableCustomField) Get() *CustomField {
	return v.value
}

func (v *NullableCustomField) Set(val *CustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomField(val *CustomField) *NullableCustomField {
	return &NullableCustomField{value: val, isSet: true}
}

func (v NullableCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


