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

// checks if the CompleteLiteralTimeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteLiteralTimeRange{}

// CompleteLiteralTimeRange struct for CompleteLiteralTimeRange
type CompleteLiteralTimeRange struct {
	ResolvableTimeRange
	// Name of the complete time range. Possible values are: - `today`, - `yesterday`, - `previous_week`, - `previous_month`.
	RangeName string `json:"rangeName" validate:"regexp=^(today|yesterday|previous_week|previous_month)$"`
}

type _CompleteLiteralTimeRange CompleteLiteralTimeRange

// NewCompleteLiteralTimeRange instantiates a new CompleteLiteralTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteLiteralTimeRange(rangeName string, type_ string) *CompleteLiteralTimeRange {
	this := CompleteLiteralTimeRange{}
	this.Type = type_
	this.RangeName = rangeName
	return &this
}

// NewCompleteLiteralTimeRangeWithDefaults instantiates a new CompleteLiteralTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteLiteralTimeRangeWithDefaults() *CompleteLiteralTimeRange {
	this := CompleteLiteralTimeRange{}
	return &this
}

// GetRangeName returns the RangeName field value
func (o *CompleteLiteralTimeRange) GetRangeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RangeName
}

// GetRangeNameOk returns a tuple with the RangeName field value
// and a boolean to check if the value has been set.
func (o *CompleteLiteralTimeRange) GetRangeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeName, true
}

// SetRangeName sets field value
func (o *CompleteLiteralTimeRange) SetRangeName(v string) {
	o.RangeName = v
}

func (o CompleteLiteralTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteLiteralTimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedResolvableTimeRange, errResolvableTimeRange := json.Marshal(o.ResolvableTimeRange)
	if errResolvableTimeRange != nil {
		return map[string]interface{}{}, errResolvableTimeRange
	}
	errResolvableTimeRange = json.Unmarshal([]byte(serializedResolvableTimeRange), &toSerialize)
	if errResolvableTimeRange != nil {
		return map[string]interface{}{}, errResolvableTimeRange
	}
	toSerialize["rangeName"] = o.RangeName
	return toSerialize, nil
}

func (o *CompleteLiteralTimeRange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rangeName",
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

	varCompleteLiteralTimeRange := _CompleteLiteralTimeRange{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompleteLiteralTimeRange)

	if err != nil {
		return err
	}

	*o = CompleteLiteralTimeRange(varCompleteLiteralTimeRange)

	return err
}

type NullableCompleteLiteralTimeRange struct {
	value *CompleteLiteralTimeRange
	isSet bool
}

func (v NullableCompleteLiteralTimeRange) Get() *CompleteLiteralTimeRange {
	return v.value
}

func (v *NullableCompleteLiteralTimeRange) Set(val *CompleteLiteralTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteLiteralTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteLiteralTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteLiteralTimeRange(val *CompleteLiteralTimeRange) *NullableCompleteLiteralTimeRange {
	return &NullableCompleteLiteralTimeRange{value: val, isSet: true}
}

func (v NullableCompleteLiteralTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteLiteralTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


