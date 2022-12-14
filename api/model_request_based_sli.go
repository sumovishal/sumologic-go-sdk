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

// RequestBasedSli struct for RequestBasedSli
type RequestBasedSli struct {
	Sli
	// Compared against threshold query's raw data points to determine success.
	Threshold *float32 `json:"threshold,omitempty"`
	// Comparison function with threshold (LessThan/GreaterThan/LessThanOrEqual/GreaterThanOrEqual).
	Op *string `json:"op,omitempty"`
}

// NewRequestBasedSli instantiates a new RequestBasedSli object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestBasedSli(evaluationType string, queryType string, queries []SliQueryGroup) *RequestBasedSli {
	this := RequestBasedSli{}
	this.EvaluationType = evaluationType
	this.QueryType = queryType
	this.Queries = queries
	return &this
}

// NewRequestBasedSliWithDefaults instantiates a new RequestBasedSli object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestBasedSliWithDefaults() *RequestBasedSli {
	this := RequestBasedSli{}
	return &this
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RequestBasedSli) GetThreshold() float32 {
	if o == nil || o.Threshold == nil {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBasedSli) GetThresholdOk() (*float32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RequestBasedSli) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *RequestBasedSli) SetThreshold(v float32) {
	o.Threshold = &v
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *RequestBasedSli) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBasedSli) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *RequestBasedSli) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *RequestBasedSli) SetOp(v string) {
	o.Op = &v
}

func (o RequestBasedSli) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSli, errSli := json.Marshal(o.Sli)
	if errSli != nil {
		return []byte{}, errSli
	}
	errSli = json.Unmarshal([]byte(serializedSli), &toSerialize)
	if errSli != nil {
		return []byte{}, errSli
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	return json.Marshal(toSerialize)
}

type NullableRequestBasedSli struct {
	value *RequestBasedSli
	isSet bool
}

func (v NullableRequestBasedSli) Get() *RequestBasedSli {
	return v.value
}

func (v *NullableRequestBasedSli) Set(val *RequestBasedSli) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestBasedSli) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestBasedSli) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestBasedSli(val *RequestBasedSli) *NullableRequestBasedSli {
	return &NullableRequestBasedSli{value: val, isSet: true}
}

func (v NullableRequestBasedSli) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestBasedSli) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


