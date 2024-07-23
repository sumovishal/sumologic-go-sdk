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

// checks if the VisualMetaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualMetaData{}

// VisualMetaData struct for VisualMetaData
type VisualMetaData struct {
	// The value of the metadata.
	Data map[string]string `json:"data"`
}

type _VisualMetaData VisualMetaData

// NewVisualMetaData instantiates a new VisualMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualMetaData(data map[string]string) *VisualMetaData {
	this := VisualMetaData{}
	this.Data = data
	return &this
}

// NewVisualMetaDataWithDefaults instantiates a new VisualMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualMetaDataWithDefaults() *VisualMetaData {
	this := VisualMetaData{}
	return &this
}

// GetData returns the Data field value
func (o *VisualMetaData) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *VisualMetaData) GetDataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *VisualMetaData) SetData(v map[string]string) {
	o.Data = v
}

func (o VisualMetaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualMetaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *VisualMetaData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varVisualMetaData := _VisualMetaData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVisualMetaData)

	if err != nil {
		return err
	}

	*o = VisualMetaData(varVisualMetaData)

	return err
}

type NullableVisualMetaData struct {
	value *VisualMetaData
	isSet bool
}

func (v NullableVisualMetaData) Get() *VisualMetaData {
	return v.value
}

func (v *NullableVisualMetaData) Set(val *VisualMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualMetaData(val *VisualMetaData) *NullableVisualMetaData {
	return &NullableVisualMetaData{value: val, isSet: true}
}

func (v NullableVisualMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


