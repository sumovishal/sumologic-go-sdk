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

// checks if the ResolvableTimeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolvableTimeRange{}

// ResolvableTimeRange struct for ResolvableTimeRange
type ResolvableTimeRange struct {
	// Type of the time range. Value must be either `CompleteLiteralTimeRange` or `BeginBoundedTimeRange`.
	Type string `json:"type"`
}

type _ResolvableTimeRange ResolvableTimeRange

// NewResolvableTimeRange instantiates a new ResolvableTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolvableTimeRange(type_ string) *ResolvableTimeRange {
	this := ResolvableTimeRange{}
	this.Type = type_
	return &this
}

// NewResolvableTimeRangeWithDefaults instantiates a new ResolvableTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolvableTimeRangeWithDefaults() *ResolvableTimeRange {
	this := ResolvableTimeRange{}
	return &this
}

// GetType returns the Type field value
func (o *ResolvableTimeRange) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResolvableTimeRange) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResolvableTimeRange) SetType(v string) {
	o.Type = v
}

func (o ResolvableTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolvableTimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ResolvableTimeRange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varResolvableTimeRange := _ResolvableTimeRange{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResolvableTimeRange)

	if err != nil {
		return err
	}

	*o = ResolvableTimeRange(varResolvableTimeRange)

	return err
}

type NullableResolvableTimeRange struct {
	value *ResolvableTimeRange
	isSet bool
}

func (v NullableResolvableTimeRange) Get() *ResolvableTimeRange {
	return v.value
}

func (v *NullableResolvableTimeRange) Set(val *ResolvableTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableResolvableTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableResolvableTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolvableTimeRange(val *ResolvableTimeRange) *NullableResolvableTimeRange {
	return &NullableResolvableTimeRange{value: val, isSet: true}
}

func (v NullableResolvableTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolvableTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


