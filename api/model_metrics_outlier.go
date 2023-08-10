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

// checks if the MetricsOutlier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsOutlier{}

// MetricsOutlier The parameters extracted from the metrics outlier query.
type MetricsOutlier struct {
	// The query string after trimming out the outlier clause.
	TrimmedQuery *string `json:"trimmedQuery,omitempty"`
	// The time range used to compute the baseline.
	BaselineWindow *string `json:"baselineWindow,omitempty"`
	BaselineTimeRangeWindow *ResolvableTimeRange `json:"baselineTimeRangeWindow,omitempty"`
	// Specifies which direction should trigger violations. Valid values:   1. `Both`: Both positive and negative deviations   2. `Up`: Positive deviations only   3. `Down`: Negative deviations only example: \"Up\" pattern: \"^(Both|Up|Down)$\" default: \"Both\" x-pattern-message: \"should be one of the following: 'Both', 'Up', 'Down'\"
	Direction *string `json:"direction,omitempty"`
	// How much should the indicator be different from the baseline for each datapoint.
	Threshold *float64 `json:"threshold,omitempty"`
}

// NewMetricsOutlier instantiates a new MetricsOutlier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsOutlier() *MetricsOutlier {
	this := MetricsOutlier{}
	var baselineWindow string = "5m"
	this.BaselineWindow = &baselineWindow
	var threshold float64 = 3.0
	this.Threshold = &threshold
	return &this
}

// NewMetricsOutlierWithDefaults instantiates a new MetricsOutlier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsOutlierWithDefaults() *MetricsOutlier {
	this := MetricsOutlier{}
	var baselineWindow string = "5m"
	this.BaselineWindow = &baselineWindow
	var threshold float64 = 3.0
	this.Threshold = &threshold
	return &this
}

// GetTrimmedQuery returns the TrimmedQuery field value if set, zero value otherwise.
func (o *MetricsOutlier) GetTrimmedQuery() string {
	if o == nil || IsNil(o.TrimmedQuery) {
		var ret string
		return ret
	}
	return *o.TrimmedQuery
}

// GetTrimmedQueryOk returns a tuple with the TrimmedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlier) GetTrimmedQueryOk() (*string, bool) {
	if o == nil || IsNil(o.TrimmedQuery) {
		return nil, false
	}
	return o.TrimmedQuery, true
}

// HasTrimmedQuery returns a boolean if a field has been set.
func (o *MetricsOutlier) HasTrimmedQuery() bool {
	if o != nil && !IsNil(o.TrimmedQuery) {
		return true
	}

	return false
}

// SetTrimmedQuery gets a reference to the given string and assigns it to the TrimmedQuery field.
func (o *MetricsOutlier) SetTrimmedQuery(v string) {
	o.TrimmedQuery = &v
}

// GetBaselineWindow returns the BaselineWindow field value if set, zero value otherwise.
func (o *MetricsOutlier) GetBaselineWindow() string {
	if o == nil || IsNil(o.BaselineWindow) {
		var ret string
		return ret
	}
	return *o.BaselineWindow
}

// GetBaselineWindowOk returns a tuple with the BaselineWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlier) GetBaselineWindowOk() (*string, bool) {
	if o == nil || IsNil(o.BaselineWindow) {
		return nil, false
	}
	return o.BaselineWindow, true
}

// HasBaselineWindow returns a boolean if a field has been set.
func (o *MetricsOutlier) HasBaselineWindow() bool {
	if o != nil && !IsNil(o.BaselineWindow) {
		return true
	}

	return false
}

// SetBaselineWindow gets a reference to the given string and assigns it to the BaselineWindow field.
func (o *MetricsOutlier) SetBaselineWindow(v string) {
	o.BaselineWindow = &v
}

// GetBaselineTimeRangeWindow returns the BaselineTimeRangeWindow field value if set, zero value otherwise.
func (o *MetricsOutlier) GetBaselineTimeRangeWindow() ResolvableTimeRange {
	if o == nil || IsNil(o.BaselineTimeRangeWindow) {
		var ret ResolvableTimeRange
		return ret
	}
	return *o.BaselineTimeRangeWindow
}

// GetBaselineTimeRangeWindowOk returns a tuple with the BaselineTimeRangeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlier) GetBaselineTimeRangeWindowOk() (*ResolvableTimeRange, bool) {
	if o == nil || IsNil(o.BaselineTimeRangeWindow) {
		return nil, false
	}
	return o.BaselineTimeRangeWindow, true
}

// HasBaselineTimeRangeWindow returns a boolean if a field has been set.
func (o *MetricsOutlier) HasBaselineTimeRangeWindow() bool {
	if o != nil && !IsNil(o.BaselineTimeRangeWindow) {
		return true
	}

	return false
}

// SetBaselineTimeRangeWindow gets a reference to the given ResolvableTimeRange and assigns it to the BaselineTimeRangeWindow field.
func (o *MetricsOutlier) SetBaselineTimeRangeWindow(v ResolvableTimeRange) {
	o.BaselineTimeRangeWindow = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *MetricsOutlier) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlier) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MetricsOutlier) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *MetricsOutlier) SetDirection(v string) {
	o.Direction = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *MetricsOutlier) GetThreshold() float64 {
	if o == nil || IsNil(o.Threshold) {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsOutlier) GetThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *MetricsOutlier) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *MetricsOutlier) SetThreshold(v float64) {
	o.Threshold = &v
}

func (o MetricsOutlier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsOutlier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrimmedQuery) {
		toSerialize["trimmedQuery"] = o.TrimmedQuery
	}
	if !IsNil(o.BaselineWindow) {
		toSerialize["baselineWindow"] = o.BaselineWindow
	}
	if !IsNil(o.BaselineTimeRangeWindow) {
		toSerialize["baselineTimeRangeWindow"] = o.BaselineTimeRangeWindow
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	return toSerialize, nil
}

type NullableMetricsOutlier struct {
	value *MetricsOutlier
	isSet bool
}

func (v NullableMetricsOutlier) Get() *MetricsOutlier {
	return v.value
}

func (v *NullableMetricsOutlier) Set(val *MetricsOutlier) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsOutlier) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsOutlier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsOutlier(val *MetricsOutlier) *NullableMetricsOutlier {
	return &NullableMetricsOutlier{value: val, isSet: true}
}

func (v NullableMetricsOutlier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsOutlier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


