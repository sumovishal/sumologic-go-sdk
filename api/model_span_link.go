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

// checks if the SpanLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanLink{}

// SpanLink Details of the linked span.
type SpanLink struct {
	// Trace identifier of the linked span.
	TraceId string `json:"traceId"`
	// Span identifier of the linked span.
	SpanId string `json:"spanId"`
}

type _SpanLink SpanLink

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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["traceId"] = o.TraceId
	toSerialize["spanId"] = o.SpanId
	return toSerialize, nil
}

func (o *SpanLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"traceId",
		"spanId",
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

	varSpanLink := _SpanLink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpanLink)

	if err != nil {
		return err
	}

	*o = SpanLink(varSpanLink)

	return err
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


