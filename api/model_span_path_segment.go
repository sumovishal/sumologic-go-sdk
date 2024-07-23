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

// checks if the SpanPathSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanPathSegment{}

// SpanPathSegment struct for SpanPathSegment
type SpanPathSegment struct {
	// Span identifier.
	SpanId string `json:"spanId"`
	// The name of the service this span is part of.
	Service *string `json:"service,omitempty"`
	// Color hex code assigned to the service.
	ServiceColor *string `json:"serviceColor,omitempty"`
	// Number of nanoseconds from the span startedAt the segment started.
	StartOffset int64 `json:"startOffset"`
	// Number of nanoseconds the span segment lasted.
	Duration int64 `json:"duration"`
	// The fraction (value between 0.0 and 1.0) from the trace duration time this segment took.
	Fraction *float64 `json:"fraction,omitempty"`
}

type _SpanPathSegment SpanPathSegment

// NewSpanPathSegment instantiates a new SpanPathSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanPathSegment(spanId string, startOffset int64, duration int64) *SpanPathSegment {
	this := SpanPathSegment{}
	this.SpanId = spanId
	this.StartOffset = startOffset
	this.Duration = duration
	return &this
}

// NewSpanPathSegmentWithDefaults instantiates a new SpanPathSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanPathSegmentWithDefaults() *SpanPathSegment {
	this := SpanPathSegment{}
	return &this
}

// GetSpanId returns the SpanId field value
func (o *SpanPathSegment) GetSpanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value
// and a boolean to check if the value has been set.
func (o *SpanPathSegment) GetSpanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpanId, true
}

// SetSpanId sets field value
func (o *SpanPathSegment) SetSpanId(v string) {
	o.SpanId = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SpanPathSegment) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanPathSegment) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SpanPathSegment) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SpanPathSegment) SetService(v string) {
	o.Service = &v
}

// GetServiceColor returns the ServiceColor field value if set, zero value otherwise.
func (o *SpanPathSegment) GetServiceColor() string {
	if o == nil || IsNil(o.ServiceColor) {
		var ret string
		return ret
	}
	return *o.ServiceColor
}

// GetServiceColorOk returns a tuple with the ServiceColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanPathSegment) GetServiceColorOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceColor) {
		return nil, false
	}
	return o.ServiceColor, true
}

// HasServiceColor returns a boolean if a field has been set.
func (o *SpanPathSegment) HasServiceColor() bool {
	if o != nil && !IsNil(o.ServiceColor) {
		return true
	}

	return false
}

// SetServiceColor gets a reference to the given string and assigns it to the ServiceColor field.
func (o *SpanPathSegment) SetServiceColor(v string) {
	o.ServiceColor = &v
}

// GetStartOffset returns the StartOffset field value
func (o *SpanPathSegment) GetStartOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartOffset
}

// GetStartOffsetOk returns a tuple with the StartOffset field value
// and a boolean to check if the value has been set.
func (o *SpanPathSegment) GetStartOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartOffset, true
}

// SetStartOffset sets field value
func (o *SpanPathSegment) SetStartOffset(v int64) {
	o.StartOffset = v
}

// GetDuration returns the Duration field value
func (o *SpanPathSegment) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SpanPathSegment) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SpanPathSegment) SetDuration(v int64) {
	o.Duration = v
}

// GetFraction returns the Fraction field value if set, zero value otherwise.
func (o *SpanPathSegment) GetFraction() float64 {
	if o == nil || IsNil(o.Fraction) {
		var ret float64
		return ret
	}
	return *o.Fraction
}

// GetFractionOk returns a tuple with the Fraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanPathSegment) GetFractionOk() (*float64, bool) {
	if o == nil || IsNil(o.Fraction) {
		return nil, false
	}
	return o.Fraction, true
}

// HasFraction returns a boolean if a field has been set.
func (o *SpanPathSegment) HasFraction() bool {
	if o != nil && !IsNil(o.Fraction) {
		return true
	}

	return false
}

// SetFraction gets a reference to the given float64 and assigns it to the Fraction field.
func (o *SpanPathSegment) SetFraction(v float64) {
	o.Fraction = &v
}

func (o SpanPathSegment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanPathSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["spanId"] = o.SpanId
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.ServiceColor) {
		toSerialize["serviceColor"] = o.ServiceColor
	}
	toSerialize["startOffset"] = o.StartOffset
	toSerialize["duration"] = o.Duration
	if !IsNil(o.Fraction) {
		toSerialize["fraction"] = o.Fraction
	}
	return toSerialize, nil
}

func (o *SpanPathSegment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"spanId",
		"startOffset",
		"duration",
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

	varSpanPathSegment := _SpanPathSegment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpanPathSegment)

	if err != nil {
		return err
	}

	*o = SpanPathSegment(varSpanPathSegment)

	return err
}

type NullableSpanPathSegment struct {
	value *SpanPathSegment
	isSet bool
}

func (v NullableSpanPathSegment) Get() *SpanPathSegment {
	return v.value
}

func (v *NullableSpanPathSegment) Set(val *SpanPathSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanPathSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanPathSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanPathSegment(val *SpanPathSegment) *NullableSpanPathSegment {
	return &NullableSpanPathSegment{value: val, isSet: true}
}

func (v NullableSpanPathSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanPathSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


