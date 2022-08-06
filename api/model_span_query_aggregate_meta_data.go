/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// SpanQueryAggregateMetaData struct for SpanQueryAggregateMetaData
type SpanQueryAggregateMetaData struct {
	// The value of the metadata.
	Data map[string]string `json:"data"`
}

// NewSpanQueryAggregateMetaData instantiates a new SpanQueryAggregateMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanQueryAggregateMetaData(data map[string]string) *SpanQueryAggregateMetaData {
	this := SpanQueryAggregateMetaData{}
	this.Data = data
	return &this
}

// NewSpanQueryAggregateMetaDataWithDefaults instantiates a new SpanQueryAggregateMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanQueryAggregateMetaDataWithDefaults() *SpanQueryAggregateMetaData {
	this := SpanQueryAggregateMetaData{}
	return &this
}

// GetData returns the Data field value
func (o *SpanQueryAggregateMetaData) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SpanQueryAggregateMetaData) GetDataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SpanQueryAggregateMetaData) SetData(v map[string]string) {
	o.Data = v
}

func (o SpanQueryAggregateMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSpanQueryAggregateMetaData struct {
	value *SpanQueryAggregateMetaData
	isSet bool
}

func (v NullableSpanQueryAggregateMetaData) Get() *SpanQueryAggregateMetaData {
	return v.value
}

func (v *NullableSpanQueryAggregateMetaData) Set(val *SpanQueryAggregateMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanQueryAggregateMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanQueryAggregateMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanQueryAggregateMetaData(val *SpanQueryAggregateMetaData) *NullableSpanQueryAggregateMetaData {
	return &NullableSpanQueryAggregateMetaData{value: val, isSet: true}
}

func (v NullableSpanQueryAggregateMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanQueryAggregateMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


