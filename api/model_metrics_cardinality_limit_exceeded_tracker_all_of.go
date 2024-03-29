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

// checks if the MetricsCardinalityLimitExceededTrackerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsCardinalityLimitExceededTrackerAllOf{}

// MetricsCardinalityLimitExceededTrackerAllOf struct for MetricsCardinalityLimitExceededTrackerAllOf
type MetricsCardinalityLimitExceededTrackerAllOf struct {
	// The retention of metrics that exceeded the limit.
	Retention *string `json:"retention,omitempty"`
}

// NewMetricsCardinalityLimitExceededTrackerAllOf instantiates a new MetricsCardinalityLimitExceededTrackerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsCardinalityLimitExceededTrackerAllOf() *MetricsCardinalityLimitExceededTrackerAllOf {
	this := MetricsCardinalityLimitExceededTrackerAllOf{}
	return &this
}

// NewMetricsCardinalityLimitExceededTrackerAllOfWithDefaults instantiates a new MetricsCardinalityLimitExceededTrackerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsCardinalityLimitExceededTrackerAllOfWithDefaults() *MetricsCardinalityLimitExceededTrackerAllOf {
	this := MetricsCardinalityLimitExceededTrackerAllOf{}
	return &this
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *MetricsCardinalityLimitExceededTrackerAllOf) GetRetention() string {
	if o == nil || IsNil(o.Retention) {
		var ret string
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsCardinalityLimitExceededTrackerAllOf) GetRetentionOk() (*string, bool) {
	if o == nil || IsNil(o.Retention) {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *MetricsCardinalityLimitExceededTrackerAllOf) HasRetention() bool {
	if o != nil && !IsNil(o.Retention) {
		return true
	}

	return false
}

// SetRetention gets a reference to the given string and assigns it to the Retention field.
func (o *MetricsCardinalityLimitExceededTrackerAllOf) SetRetention(v string) {
	o.Retention = &v
}

func (o MetricsCardinalityLimitExceededTrackerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsCardinalityLimitExceededTrackerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Retention) {
		toSerialize["retention"] = o.Retention
	}
	return toSerialize, nil
}

type NullableMetricsCardinalityLimitExceededTrackerAllOf struct {
	value *MetricsCardinalityLimitExceededTrackerAllOf
	isSet bool
}

func (v NullableMetricsCardinalityLimitExceededTrackerAllOf) Get() *MetricsCardinalityLimitExceededTrackerAllOf {
	return v.value
}

func (v *NullableMetricsCardinalityLimitExceededTrackerAllOf) Set(val *MetricsCardinalityLimitExceededTrackerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsCardinalityLimitExceededTrackerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsCardinalityLimitExceededTrackerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsCardinalityLimitExceededTrackerAllOf(val *MetricsCardinalityLimitExceededTrackerAllOf) *NullableMetricsCardinalityLimitExceededTrackerAllOf {
	return &NullableMetricsCardinalityLimitExceededTrackerAllOf{value: val, isSet: true}
}

func (v NullableMetricsCardinalityLimitExceededTrackerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsCardinalityLimitExceededTrackerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


