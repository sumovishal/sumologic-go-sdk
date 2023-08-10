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

// checks if the SpansTimeGroupBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpansTimeGroupBy{}

// SpansTimeGroupBy struct for SpansTimeGroupBy
type SpansTimeGroupBy struct {
	SpansGroupBy
	// A fixed interval grouping in the following format <#><time_period>,  supported <time_period> values are weeks (w), days (d), hours (h), minutes (m), and seconds (s). 
	FieldValue string `json:"fieldValue"`
}

// NewSpansTimeGroupBy instantiates a new SpansTimeGroupBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpansTimeGroupBy(fieldValue string, type_ string) *SpansTimeGroupBy {
	this := SpansTimeGroupBy{}
	this.Type = type_
	this.FieldValue = fieldValue
	return &this
}

// NewSpansTimeGroupByWithDefaults instantiates a new SpansTimeGroupBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpansTimeGroupByWithDefaults() *SpansTimeGroupBy {
	this := SpansTimeGroupBy{}
	return &this
}

// GetFieldValue returns the FieldValue field value
func (o *SpansTimeGroupBy) GetFieldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value
// and a boolean to check if the value has been set.
func (o *SpansTimeGroupBy) GetFieldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldValue, true
}

// SetFieldValue sets field value
func (o *SpansTimeGroupBy) SetFieldValue(v string) {
	o.FieldValue = v
}

func (o SpansTimeGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpansTimeGroupBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSpansGroupBy, errSpansGroupBy := json.Marshal(o.SpansGroupBy)
	if errSpansGroupBy != nil {
		return map[string]interface{}{}, errSpansGroupBy
	}
	errSpansGroupBy = json.Unmarshal([]byte(serializedSpansGroupBy), &toSerialize)
	if errSpansGroupBy != nil {
		return map[string]interface{}{}, errSpansGroupBy
	}
	toSerialize["fieldValue"] = o.FieldValue
	return toSerialize, nil
}

type NullableSpansTimeGroupBy struct {
	value *SpansTimeGroupBy
	isSet bool
}

func (v NullableSpansTimeGroupBy) Get() *SpansTimeGroupBy {
	return v.value
}

func (v *NullableSpansTimeGroupBy) Set(val *SpansTimeGroupBy) {
	v.value = val
	v.isSet = true
}

func (v NullableSpansTimeGroupBy) IsSet() bool {
	return v.isSet
}

func (v *NullableSpansTimeGroupBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpansTimeGroupBy(val *SpansTimeGroupBy) *NullableSpansTimeGroupBy {
	return &NullableSpansTimeGroupBy{value: val, isSet: true}
}

func (v NullableSpansTimeGroupBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpansTimeGroupBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


