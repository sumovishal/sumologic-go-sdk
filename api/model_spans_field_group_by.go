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

// checks if the SpansFieldGroupBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpansFieldGroupBy{}

// SpansFieldGroupBy struct for SpansFieldGroupBy
type SpansFieldGroupBy struct {
	SpansGroupBy
	// A name of the field to group by.
	FieldName string `json:"fieldName"`
}

type _SpansFieldGroupBy SpansFieldGroupBy

// NewSpansFieldGroupBy instantiates a new SpansFieldGroupBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpansFieldGroupBy(fieldName string, type_ string) *SpansFieldGroupBy {
	this := SpansFieldGroupBy{}
	this.Type = type_
	this.FieldName = fieldName
	return &this
}

// NewSpansFieldGroupByWithDefaults instantiates a new SpansFieldGroupBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpansFieldGroupByWithDefaults() *SpansFieldGroupBy {
	this := SpansFieldGroupBy{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *SpansFieldGroupBy) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *SpansFieldGroupBy) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *SpansFieldGroupBy) SetFieldName(v string) {
	o.FieldName = v
}

func (o SpansFieldGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpansFieldGroupBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSpansGroupBy, errSpansGroupBy := json.Marshal(o.SpansGroupBy)
	if errSpansGroupBy != nil {
		return map[string]interface{}{}, errSpansGroupBy
	}
	errSpansGroupBy = json.Unmarshal([]byte(serializedSpansGroupBy), &toSerialize)
	if errSpansGroupBy != nil {
		return map[string]interface{}{}, errSpansGroupBy
	}
	toSerialize["fieldName"] = o.FieldName
	return toSerialize, nil
}

func (o *SpansFieldGroupBy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fieldName",
		"type",
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

	varSpansFieldGroupBy := _SpansFieldGroupBy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpansFieldGroupBy)

	if err != nil {
		return err
	}

	*o = SpansFieldGroupBy(varSpansFieldGroupBy)

	return err
}

type NullableSpansFieldGroupBy struct {
	value *SpansFieldGroupBy
	isSet bool
}

func (v NullableSpansFieldGroupBy) Get() *SpansFieldGroupBy {
	return v.value
}

func (v *NullableSpansFieldGroupBy) Set(val *SpansFieldGroupBy) {
	v.value = val
	v.isSet = true
}

func (v NullableSpansFieldGroupBy) IsSet() bool {
	return v.isSet
}

func (v *NullableSpansFieldGroupBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpansFieldGroupBy(val *SpansFieldGroupBy) *NullableSpansFieldGroupBy {
	return &NullableSpansFieldGroupBy{value: val, isSet: true}
}

func (v NullableSpansFieldGroupBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpansFieldGroupBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


