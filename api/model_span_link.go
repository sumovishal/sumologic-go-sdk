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

// SpanLink Details of the linked span.
type SpanLink struct {
	// Trace identifier of the linked span.
	TraceId string `json:"traceId"`
	// Span identifier of the linked span.
	SpanId string `json:"spanId"`
}

// NewSpanLink instantiates a new SpanLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanLink(traceId string, spanId string) *SpanLink {
	this := SpanLink{}
	this.TraceId = traceId
	this.SpanId = spanId
	return &this
}

// NewSpanLinkWithDefaults instantiates a new SpanLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanLinkWithDefaults() *SpanLink {
	this := SpanLink{}
	return &this
}

// GetTraceId returns the TraceId field value
func (o *SpanLink) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *SpanLink) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *SpanLink) SetTraceId(v string) {
	o.TraceId = v
}

// GetSpanId returns the SpanId field value
func (o *SpanLink) GetSpanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *SpanLink) GetSpanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value
func (o *SpanLink) SetSpanId(v string) {
	o.SpanId = v
}

func (o SpanLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["traceId"] = o.TraceId
	}
	if true {
		toSerialize["spanId"] = o.SpanId
	}
	return json.Marshal(toSerialize)
}

type NullableSpanLink struct {
	value *SpanLink
	isSet bool
}

func (v NullableSpanLink) Get() *SpanLink {
	return v.value
}

func (v *NullableSpanLink) Set(val *SpanLink) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanLink) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanLink(val *SpanLink) *NullableSpanLink {
	return &NullableSpanLink{value: val, isSet: true}
}

func (v NullableSpanLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


