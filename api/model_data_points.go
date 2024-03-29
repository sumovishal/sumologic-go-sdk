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

// checks if the DataPoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPoints{}

// DataPoints Denotes the data points as a result of the groupBy function performed on the usage data.
type DataPoints struct {
	TimeRange BeginBoundedTimeRange `json:"timeRange"`
	// An array of objects denoting the value and unit of the data points.
	Values []DataValue `json:"values,omitempty"`
}

// NewDataPoints instantiates a new DataPoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPoints(timeRange BeginBoundedTimeRange) *DataPoints {
	this := DataPoints{}
	this.TimeRange = timeRange
	return &this
}

// NewDataPointsWithDefaults instantiates a new DataPoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPointsWithDefaults() *DataPoints {
	this := DataPoints{}
	return &this
}

// GetTimeRange returns the TimeRange field value
func (o *DataPoints) GetTimeRange() BeginBoundedTimeRange {
	if o == nil {
		var ret BeginBoundedTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *DataPoints) GetTimeRangeOk() (*BeginBoundedTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *DataPoints) SetTimeRange(v BeginBoundedTimeRange) {
	o.TimeRange = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *DataPoints) GetValues() []DataValue {
	if o == nil || IsNil(o.Values) {
		var ret []DataValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPoints) GetValuesOk() ([]DataValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *DataPoints) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []DataValue and assigns it to the Values field.
func (o *DataPoints) SetValues(v []DataValue) {
	o.Values = v
}

func (o DataPoints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataPoints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableDataPoints struct {
	value *DataPoints
	isSet bool
}

func (v NullableDataPoints) Get() *DataPoints {
	return v.value
}

func (v *NullableDataPoints) Set(val *DataPoints) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPoints) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPoints(val *DataPoints) *NullableDataPoints {
	return &NullableDataPoints{value: val, isSet: true}
}

func (v NullableDataPoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


