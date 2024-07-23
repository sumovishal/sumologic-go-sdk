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

// checks if the OtTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OtTag{}

// OtTag struct for OtTag
type OtTag struct {
	// key of the given tag.
	Key string `json:"key"`
	// values of the given tag.
	Values []string `json:"values"`
}

type _OtTag OtTag

// NewOtTag instantiates a new OtTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtTag(key string, values []string) *OtTag {
	this := OtTag{}
	this.Key = key
	this.Values = values
	return &this
}

// NewOtTagWithDefaults instantiates a new OtTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtTagWithDefaults() *OtTag {
	this := OtTag{}
	return &this
}

// GetKey returns the Key field value
func (o *OtTag) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *OtTag) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *OtTag) SetKey(v string) {
	o.Key = v
}

// GetValues returns the Values field value
func (o *OtTag) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *OtTag) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *OtTag) SetValues(v []string) {
	o.Values = v
}

func (o OtTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OtTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *OtTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"values",
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

	varOtTag := _OtTag{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOtTag)

	if err != nil {
		return err
	}

	*o = OtTag(varOtTag)

	return err
}

type NullableOtTag struct {
	value *OtTag
	isSet bool
}

func (v NullableOtTag) Get() *OtTag {
	return v.value
}

func (v *NullableOtTag) Set(val *OtTag) {
	v.value = val
	v.isSet = true
}

func (v NullableOtTag) IsSet() bool {
	return v.isSet
}

func (v *NullableOtTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtTag(val *OtTag) *NullableOtTag {
	return &NullableOtTag{value: val, isSet: true}
}

func (v NullableOtTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


