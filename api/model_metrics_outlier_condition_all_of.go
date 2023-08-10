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

// checks if the MetricsOutlierConditionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsOutlierConditionAllOf{}

// MetricsOutlierConditionAllOf A rule that defines how metrics monitor should evaluate outlier data and trigger notifications.
type MetricsOutlierConditionAllOf struct {
	// The time range used to compute the baseline.
	BaselineWindow *string `json:"baselineWindow,omitempty"`
	// Specifies which direction should trigger violations.
	Direction *string `json:"direction,omitempty"`
	// How much should the indicator be different from the baseline for each datapoint.
	Threshold *float64 `json:"threshold,omitempty"`
}

// NewMetricsOutlierConditionAllOf instantiates a new MetricsOutlierConditionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsOutlierConditionAllOf() *MetricsOutlierConditionAllOf {
	this := MetricsOutlierConditionAllOf{}
	var baselineWindow string = "1d"
	this.BaselineWindow = &baselineWindow
	var direction string = "Both"
	this.Direction = &direction
	var threshold float64 = 3.0
	this.Threshold = &threshold
	return &this
}

// NewMetricsOutlierConditionAllOfWithDefaults instantiates a new MetricsOutlierConditionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsOutlierConditionAllOfWithDefaults() *MetricsOutlierConditionAllOf {
	this := MetricsOutlierConditionAllOf{}
	var baselineWindow string = "1d"
	this.BaselineWindow = &baselineWindow
	var direction string = "Both"
	this.Direction = &direction
	var threshold float64 = 3.0
	this.Threshold = &threshold
	return &this
}

// GetBaselineWindow returns the BaselineWindow field value if set, zero value otherwise.
func (o *MetricsOutlierConditionAllOf) GetBaselineWindow() string {
	if o == nil || IsNil(o.BaselineWindow) {
		var ret string
		return ret
	}
	return *o.BaselineWindow
}

// GetBaselineWindowOk returns a tuple with the BaselineWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlierConditionAllOf) GetBaselineWindowOk() (*string, bool) {
	if o == nil || IsNil(o.BaselineWindow) {
		return nil, false
	}
	return o.BaselineWindow, true
}

// HasBaselineWindow returns a boolean if a field has been set.
func (o *MetricsOutlierConditionAllOf) HasBaselineWindow() bool {
	if o != nil && !IsNil(o.BaselineWindow) {
		return true
	}

	return false
}

// SetBaselineWindow gets a reference to the given string and assigns it to the BaselineWindow field.
func (o *MetricsOutlierConditionAllOf) SetBaselineWindow(v string) {
	o.BaselineWindow = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *MetricsOutlierConditionAllOf) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlierConditionAllOf) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MetricsOutlierConditionAllOf) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *MetricsOutlierConditionAllOf) SetDirection(v string) {
	o.Direction = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *MetricsOutlierConditionAllOf) GetThreshold() float64 {
	if o == nil || IsNil(o.Threshold) {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlierConditionAllOf) GetThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *MetricsOutlierConditionAllOf) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *MetricsOutlierConditionAllOf) SetThreshold(v float64) {
	o.Threshold = &v
}

func (o MetricsOutlierConditionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsOutlierConditionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaselineWindow) {
		toSerialize["baselineWindow"] = o.BaselineWindow
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	return toSerialize, nil
}

type NullableMetricsOutlierConditionAllOf struct {
	value *MetricsOutlierConditionAllOf
	isSet bool
}

func (v NullableMetricsOutlierConditionAllOf) Get() *MetricsOutlierConditionAllOf {
	return v.value
}

func (v *NullableMetricsOutlierConditionAllOf) Set(val *MetricsOutlierConditionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsOutlierConditionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsOutlierConditionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsOutlierConditionAllOf(val *MetricsOutlierConditionAllOf) *NullableMetricsOutlierConditionAllOf {
	return &NullableMetricsOutlierConditionAllOf{value: val, isSet: true}
}

func (v NullableMetricsOutlierConditionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsOutlierConditionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


