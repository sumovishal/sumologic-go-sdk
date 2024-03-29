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

// checks if the SpanQueryResultSpansResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanQueryResultSpansResponse{}

// SpanQueryResultSpansResponse struct for SpanQueryResultSpansResponse
type SpanQueryResultSpansResponse struct {
	// List of trace spans.
	SpanPage []SpanQuerySpanData `json:"spanPage"`
	// Next continuation token.
	Next *string `json:"next,omitempty"`
}

// NewSpanQueryResultSpansResponse instantiates a new SpanQueryResultSpansResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanQueryResultSpansResponse(spanPage []SpanQuerySpanData) *SpanQueryResultSpansResponse {
	this := SpanQueryResultSpansResponse{}
	this.SpanPage = spanPage
	return &this
}

// NewSpanQueryResultSpansResponseWithDefaults instantiates a new SpanQueryResultSpansResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanQueryResultSpansResponseWithDefaults() *SpanQueryResultSpansResponse {
	this := SpanQueryResultSpansResponse{}
	return &this
}

// GetSpanPage returns the SpanPage field value
func (o *SpanQueryResultSpansResponse) GetSpanPage() []SpanQuerySpanData {
	if o == nil {
		var ret []SpanQuerySpanData
		return ret
	}

	return o.SpanPage
}

// GetSpanPageOk returns a tuple with the SpanPage field value
// and a boolean to check if the value has been set.
func (o *SpanQueryResultSpansResponse) GetSpanPageOk() ([]SpanQuerySpanData, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpanPage, true
}

// SetSpanPage sets field value
func (o *SpanQueryResultSpansResponse) SetSpanPage(v []SpanQuerySpanData) {
	o.SpanPage = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SpanQueryResultSpansResponse) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQueryResultSpansResponse) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SpanQueryResultSpansResponse) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *SpanQueryResultSpansResponse) SetNext(v string) {
	o.Next = &v
}

func (o SpanQueryResultSpansResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanQueryResultSpansResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["spanPage"] = o.SpanPage
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableSpanQueryResultSpansResponse struct {
	value *SpanQueryResultSpansResponse
	isSet bool
}

func (v NullableSpanQueryResultSpansResponse) Get() *SpanQueryResultSpansResponse {
	return v.value
}

func (v *NullableSpanQueryResultSpansResponse) Set(val *SpanQueryResultSpansResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanQueryResultSpansResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanQueryResultSpansResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanQueryResultSpansResponse(val *SpanQueryResultSpansResponse) *NullableSpanQueryResultSpansResponse {
	return &NullableSpanQueryResultSpansResponse{value: val, isSet: true}
}

func (v NullableSpanQueryResultSpansResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanQueryResultSpansResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


