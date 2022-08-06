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

// RequestBasedSliAllOf struct for RequestBasedSliAllOf
type RequestBasedSliAllOf struct {
	// Compared against threshold query's raw data points to determine success.
	Threshold *float32 `json:"threshold,omitempty"`
	// Comparison function with threshold (LessThan/GreaterThan/LessThanOrEqual/GreaterThanOrEqual).
	Op *string `json:"op,omitempty"`
}

// NewRequestBasedSliAllOf instantiates a new RequestBasedSliAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestBasedSliAllOf() *RequestBasedSliAllOf {
	this := RequestBasedSliAllOf{}
	return &this
}

// NewRequestBasedSliAllOfWithDefaults instantiates a new RequestBasedSliAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestBasedSliAllOfWithDefaults() *RequestBasedSliAllOf {
	this := RequestBasedSliAllOf{}
	return &this
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RequestBasedSliAllOf) GetThreshold() float32 {
	if o == nil || o.Threshold == nil {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBasedSliAllOf) GetThresholdOk() (*float32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RequestBasedSliAllOf) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *RequestBasedSliAllOf) SetThreshold(v float32) {
	o.Threshold = &v
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *RequestBasedSliAllOf) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBasedSliAllOf) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *RequestBasedSliAllOf) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *RequestBasedSliAllOf) SetOp(v string) {
	o.Op = &v
}

func (o RequestBasedSliAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	return json.Marshal(toSerialize)
}

type NullableRequestBasedSliAllOf struct {
	value *RequestBasedSliAllOf
	isSet bool
}

func (v NullableRequestBasedSliAllOf) Get() *RequestBasedSliAllOf {
	return v.value
}

func (v *NullableRequestBasedSliAllOf) Set(val *RequestBasedSliAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestBasedSliAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestBasedSliAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestBasedSliAllOf(val *RequestBasedSliAllOf) *NullableRequestBasedSliAllOf {
	return &NullableRequestBasedSliAllOf{value: val, isSet: true}
}

func (v NullableRequestBasedSliAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestBasedSliAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


