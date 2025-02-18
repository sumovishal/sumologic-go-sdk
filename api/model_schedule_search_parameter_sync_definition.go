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

// checks if the ScheduleSearchParameterSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleSearchParameterSyncDefinition{}

// ScheduleSearchParameterSyncDefinition struct for ScheduleSearchParameterSyncDefinition
type ScheduleSearchParameterSyncDefinition struct {
	// Name of scheduled search parameter.
	Name string `json:"name"`
	// Value of scheduled search parameter.
	Value string `json:"value"`
}

type _ScheduleSearchParameterSyncDefinition ScheduleSearchParameterSyncDefinition

// NewScheduleSearchParameterSyncDefinition instantiates a new ScheduleSearchParameterSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleSearchParameterSyncDefinition(name string, value string) *ScheduleSearchParameterSyncDefinition {
	this := ScheduleSearchParameterSyncDefinition{}
	this.Name = name
	this.Value = value
	return &this
}

// NewScheduleSearchParameterSyncDefinitionWithDefaults instantiates a new ScheduleSearchParameterSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleSearchParameterSyncDefinitionWithDefaults() *ScheduleSearchParameterSyncDefinition {
	this := ScheduleSearchParameterSyncDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *ScheduleSearchParameterSyncDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScheduleSearchParameterSyncDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScheduleSearchParameterSyncDefinition) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *ScheduleSearchParameterSyncDefinition) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ScheduleSearchParameterSyncDefinition) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ScheduleSearchParameterSyncDefinition) SetValue(v string) {
	o.Value = v
}

func (o ScheduleSearchParameterSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleSearchParameterSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ScheduleSearchParameterSyncDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varScheduleSearchParameterSyncDefinition := _ScheduleSearchParameterSyncDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScheduleSearchParameterSyncDefinition)

	if err != nil {
		return err
	}

	*o = ScheduleSearchParameterSyncDefinition(varScheduleSearchParameterSyncDefinition)

	return err
}

type NullableScheduleSearchParameterSyncDefinition struct {
	value *ScheduleSearchParameterSyncDefinition
	isSet bool
}

func (v NullableScheduleSearchParameterSyncDefinition) Get() *ScheduleSearchParameterSyncDefinition {
	return v.value
}

func (v *NullableScheduleSearchParameterSyncDefinition) Set(val *ScheduleSearchParameterSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleSearchParameterSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleSearchParameterSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleSearchParameterSyncDefinition(val *ScheduleSearchParameterSyncDefinition) *NullableScheduleSearchParameterSyncDefinition {
	return &NullableScheduleSearchParameterSyncDefinition{value: val, isSet: true}
}

func (v NullableScheduleSearchParameterSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleSearchParameterSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


