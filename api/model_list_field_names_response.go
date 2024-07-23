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

// checks if the ListFieldNamesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFieldNamesResponse{}

// ListFieldNamesResponse struct for ListFieldNamesResponse
type ListFieldNamesResponse struct {
	// List of all built-in and custom field names.
	Data []FieldName `json:"data"`
}

type _ListFieldNamesResponse ListFieldNamesResponse

// NewListFieldNamesResponse instantiates a new ListFieldNamesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFieldNamesResponse(data []FieldName) *ListFieldNamesResponse {
	this := ListFieldNamesResponse{}
	this.Data = data
	return &this
}

// NewListFieldNamesResponseWithDefaults instantiates a new ListFieldNamesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFieldNamesResponseWithDefaults() *ListFieldNamesResponse {
	this := ListFieldNamesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListFieldNamesResponse) GetData() []FieldName {
	if o == nil {
		var ret []FieldName
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListFieldNamesResponse) GetDataOk() ([]FieldName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListFieldNamesResponse) SetData(v []FieldName) {
	o.Data = v
}

func (o ListFieldNamesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFieldNamesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ListFieldNamesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListFieldNamesResponse := _ListFieldNamesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListFieldNamesResponse)

	if err != nil {
		return err
	}

	*o = ListFieldNamesResponse(varListFieldNamesResponse)

	return err
}

type NullableListFieldNamesResponse struct {
	value *ListFieldNamesResponse
	isSet bool
}

func (v NullableListFieldNamesResponse) Get() *ListFieldNamesResponse {
	return v.value
}

func (v *NullableListFieldNamesResponse) Set(val *ListFieldNamesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFieldNamesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFieldNamesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFieldNamesResponse(val *ListFieldNamesResponse) *NullableListFieldNamesResponse {
	return &NullableListFieldNamesResponse{value: val, isSet: true}
}

func (v NullableListFieldNamesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFieldNamesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


