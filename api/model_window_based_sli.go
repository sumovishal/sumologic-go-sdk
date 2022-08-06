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

// WindowBasedSli Evaluate SLI using successful or unsuccessful windows over compliance period.
type WindowBasedSli struct {
	// Threshold for classifying window as successful or unsuccessful.
	Threshold float32 `json:"threshold"`
	// Comparison function with window threshold (LessThan/GreaterThan/LessThanOrEqual/GreaterThanOrEqual).
	Op string `json:"op"`
	// Aggregation function applied over each window to arrive at SLI. Must be `Avg`, `Min`, `Max`, `Sum`, or percentile of the form `pX` where `X` is an integer between 1 and 99.
	Aggregation *string `json:"aggregation,omitempty"`
	// Size of the aggregation window (minimum of 1m and maximum of 1h).
	Size string `json:"size"`
}

// NewWindowBasedSli instantiates a new WindowBasedSli object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowBasedSli(threshold float32, op string, size string, evaluationType string, queryType string, queries []SliQueryGroup) *WindowBasedSli {
	this := WindowBasedSli{}
	this.EvaluationType = evaluationType
	this.QueryType = queryType
	this.Queries = queries
	this.Threshold = threshold
	this.Op = op
	this.Size = size
	return &this
}

// NewWindowBasedSliWithDefaults instantiates a new WindowBasedSli object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowBasedSliWithDefaults() *WindowBasedSli {
	this := WindowBasedSli{}
	return &this
}

// GetThreshold returns the Threshold field value
func (o *WindowBasedSli) GetThreshold() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *WindowBasedSli) GetThresholdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *WindowBasedSli) SetThreshold(v float32) {
	o.Threshold = v
}

// GetOp returns the Op field value
func (o *WindowBasedSli) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *WindowBasedSli) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *WindowBasedSli) SetOp(v string) {
	o.Op = v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *WindowBasedSli) GetAggregation() string {
	if o == nil || o.Aggregation == nil {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowBasedSli) GetAggregationOk() (*string, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *WindowBasedSli) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *WindowBasedSli) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetSize returns the Size field value
func (o *WindowBasedSli) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *WindowBasedSli) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *WindowBasedSli) SetSize(v string) {
	o.Size = v
}

func (o WindowBasedSli) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["op"] = o.Op
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableWindowBasedSli struct {
	value *WindowBasedSli
	isSet bool
}

func (v NullableWindowBasedSli) Get() *WindowBasedSli {
	return v.value
}

func (v *NullableWindowBasedSli) Set(val *WindowBasedSli) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowBasedSli) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowBasedSli) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowBasedSli(val *WindowBasedSli) *NullableWindowBasedSli {
	return &NullableWindowBasedSli{value: val, isSet: true}
}

func (v NullableWindowBasedSli) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowBasedSli) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


