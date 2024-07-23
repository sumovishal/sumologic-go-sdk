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

// checks if the Header type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Header{}

// Header struct for Header
type Header struct {
	// Name of the header field.
	Name string `json:"name"`
	// Value of the header field.
	Value string `json:"value"`
}

type _Header Header

// NewHeader instantiates a new Header object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeader(name string, value string) *Header {
	this := Header{}
	this.Name = name
	this.Value = value
	return &this
}

// NewHeaderWithDefaults instantiates a new Header object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderWithDefaults() *Header {
	this := Header{}
	return &this
}

// GetName returns the Name field value
func (o *Header) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Header) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Header) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *Header) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Header) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Header) SetValue(v string) {
	o.Value = v
}

func (o Header) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Header) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *Header) UnmarshalJSON(data []byte) (err error) {
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

	varHeader := _Header{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHeader)

	if err != nil {
		return err
	}

	*o = Header(varHeader)

	return err
}

type NullableHeader struct {
	value *Header
	isSet bool
}

func (v NullableHeader) Get() *Header {
	return v.value
}

func (v *NullableHeader) Set(val *Header) {
	v.value = val
	v.isSet = true
}

func (v NullableHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeader(val *Header) *NullableHeader {
	return &NullableHeader{value: val, isSet: true}
}

func (v NullableHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


