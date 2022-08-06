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

// TraceLightEventsResponse struct for TraceLightEventsResponse
type TraceLightEventsResponse struct {
	// Map of span ids to lists of their events, without their attributes.
	SpanEvents *map[string][]LightSpanEvent `json:"spanEvents,omitempty"`
	// Next continuation token.
	Next *string `json:"next,omitempty"`
}

// NewTraceLightEventsResponse instantiates a new TraceLightEventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceLightEventsResponse() *TraceLightEventsResponse {
	this := TraceLightEventsResponse{}
	return &this
}

// NewTraceLightEventsResponseWithDefaults instantiates a new TraceLightEventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceLightEventsResponseWithDefaults() *TraceLightEventsResponse {
	this := TraceLightEventsResponse{}
	return &this
}

// GetSpanEvents returns the SpanEvents field value if set, zero value otherwise.
func (o *TraceLightEventsResponse) GetSpanEvents() map[string][]LightSpanEvent {
	if o == nil || o.SpanEvents == nil {
		var ret map[string][]LightSpanEvent
		return ret
	}
	return *o.SpanEvents
}

// GetSpanEventsOk returns a tuple with the SpanEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceLightEventsResponse) GetSpanEventsOk() (*map[string][]LightSpanEvent, bool) {
	if o == nil || o.SpanEvents == nil {
		return nil, false
	}
	return o.SpanEvents, true
}

// HasSpanEvents returns a boolean if a field has been set.
func (o *TraceLightEventsResponse) HasSpanEvents() bool {
	if o != nil && o.SpanEvents != nil {
		return true
	}

	return false
}

// SetSpanEvents gets a reference to the given map[string][]LightSpanEvent and assigns it to the SpanEvents field.
func (o *TraceLightEventsResponse) SetSpanEvents(v map[string][]LightSpanEvent) {
	o.SpanEvents = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *TraceLightEventsResponse) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceLightEventsResponse) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *TraceLightEventsResponse) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *TraceLightEventsResponse) SetNext(v string) {
	o.Next = &v
}

func (o TraceLightEventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpanEvents != nil {
		toSerialize["spanEvents"] = o.SpanEvents
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	return json.Marshal(toSerialize)
}

type NullableTraceLightEventsResponse struct {
	value *TraceLightEventsResponse
	isSet bool
}

func (v NullableTraceLightEventsResponse) Get() *TraceLightEventsResponse {
	return v.value
}

func (v *NullableTraceLightEventsResponse) Set(val *TraceLightEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceLightEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceLightEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceLightEventsResponse(val *TraceLightEventsResponse) *NullableTraceLightEventsResponse {
	return &NullableTraceLightEventsResponse{value: val, isSet: true}
}

func (v NullableTraceLightEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceLightEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


