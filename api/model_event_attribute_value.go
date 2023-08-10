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

// checks if the EventAttributeValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventAttributeValue{}

// EventAttributeValue struct for EventAttributeValue
type EventAttributeValue struct {
	// Type of the event attribute value.
	Type string `json:"type"`
}

// NewEventAttributeValue instantiates a new EventAttributeValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventAttributeValue(type_ string) *EventAttributeValue {
	this := EventAttributeValue{}
	this.Type = type_
	return &this
}

// NewEventAttributeValueWithDefaults instantiates a new EventAttributeValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventAttributeValueWithDefaults() *EventAttributeValue {
	this := EventAttributeValue{}
	return &this
}

// GetType returns the Type field value
func (o *EventAttributeValue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventAttributeValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventAttributeValue) SetType(v string) {
	o.Type = v
}

func (o EventAttributeValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventAttributeValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableEventAttributeValue struct {
	value *EventAttributeValue
	isSet bool
}

func (v NullableEventAttributeValue) Get() *EventAttributeValue {
	return v.value
}

func (v *NullableEventAttributeValue) Set(val *EventAttributeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEventAttributeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEventAttributeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventAttributeValue(val *EventAttributeValue) *NullableEventAttributeValue {
	return &NullableEventAttributeValue{value: val, isSet: true}
}

func (v NullableEventAttributeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventAttributeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


