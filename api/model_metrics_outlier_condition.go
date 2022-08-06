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

// MetricsOutlierCondition struct for MetricsOutlierCondition
type MetricsOutlierCondition struct {
	TriggerCondition
	// The time range used to compute the baseline.
	BaselineWindow *string `json:"baselineWindow,omitempty"`
	// Specifies which direction should trigger violations.
	Direction *string `json:"direction,omitempty"`
	// How much should the indicator be different from the baseline for each datapoint.
	Threshold *float64 `json:"threshold,omitempty"`
}

// NewMetricsOutlierCondition instantiates a new MetricsOutlierCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsOutlierCondition(triggerType string) *MetricsOutlierCondition {
	this := MetricsOutlierCondition{}
	var detectionMethod string = "StaticCondition"
	this.DetectionMethod = &detectionMethod
	this.TriggerType = triggerType
	var baselineWindow string = "1d"
	this.BaselineWindow = &baselineWindow
	var direction string = "Both"
	this.Direction = &direction
	var threshold float64 = 3.0
	this.Threshold = &threshold
	return &this
}

// NewMetricsOutlierConditionWithDefaults instantiates a new MetricsOutlierCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsOutlierConditionWithDefaults() *MetricsOutlierCondition {
	this := MetricsOutlierCondition{}
	var baselineWindow string = "1d"
	this.BaselineWindow = &baselineWindow
	var direction string = "Both"
	this.Direction = &direction
	var threshold float64 = 3.0
	this.Threshold = &threshold
	return &this
}

// GetBaselineWindow returns the BaselineWindow field value if set, zero value otherwise.
func (o *MetricsOutlierCondition) GetBaselineWindow() string {
	if o == nil || o.BaselineWindow == nil {
		var ret string
		return ret
	}
	return *o.BaselineWindow
}

// GetBaselineWindowOk returns a tuple with the BaselineWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlierCondition) GetBaselineWindowOk() (*string, bool) {
	if o == nil || o.BaselineWindow == nil {
		return nil, false
	}
	return o.BaselineWindow, true
}

// HasBaselineWindow returns a boolean if a field has been set.
func (o *MetricsOutlierCondition) HasBaselineWindow() bool {
	if o != nil && o.BaselineWindow != nil {
		return true
	}

	return false
}

// SetBaselineWindow gets a reference to the given string and assigns it to the BaselineWindow field.
func (o *MetricsOutlierCondition) SetBaselineWindow(v string) {
	o.BaselineWindow = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *MetricsOutlierCondition) GetDirection() string {
	if o == nil || o.Direction == nil {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlierCondition) GetDirectionOk() (*string, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MetricsOutlierCondition) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *MetricsOutlierCondition) SetDirection(v string) {
	o.Direction = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *MetricsOutlierCondition) GetThreshold() float64 {
	if o == nil || o.Threshold == nil {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlierCondition) GetThresholdOk() (*float64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *MetricsOutlierCondition) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *MetricsOutlierCondition) SetThreshold(v float64) {
	o.Threshold = &v
}

func (o MetricsOutlierCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTriggerCondition, errTriggerCondition := json.Marshal(o.TriggerCondition)
	if errTriggerCondition != nil {
		return []byte{}, errTriggerCondition
	}
	errTriggerCondition = json.Unmarshal([]byte(serializedTriggerCondition), &toSerialize)
	if errTriggerCondition != nil {
		return []byte{}, errTriggerCondition
	}
	if o.BaselineWindow != nil {
		toSerialize["baselineWindow"] = o.BaselineWindow
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsOutlierCondition struct {
	value *MetricsOutlierCondition
	isSet bool
}

func (v NullableMetricsOutlierCondition) Get() *MetricsOutlierCondition {
	return v.value
}

func (v *NullableMetricsOutlierCondition) Set(val *MetricsOutlierCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsOutlierCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsOutlierCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsOutlierCondition(val *MetricsOutlierCondition) *NullableMetricsOutlierCondition {
	return &NullableMetricsOutlierCondition{value: val, isSet: true}
}

func (v NullableMetricsOutlierCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsOutlierCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


