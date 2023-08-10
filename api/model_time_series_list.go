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

// checks if the TimeSeriesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSeriesList{}

// TimeSeriesList struct for TimeSeriesList
type TimeSeriesList struct {
	// A list of timeseries returned by corresponding query.
	TimeSeries []TimeSeries `json:"timeSeries"`
	// Unit of the query.
	Unit *string `json:"unit,omitempty"`
	// Time shift value if specified in request in human readable format.
	TimeShiftLabel *string `json:"timeShiftLabel,omitempty"`
	ResultContext *MetricsQueryResultContext `json:"resultContext,omitempty"`
}

// NewTimeSeriesList instantiates a new TimeSeriesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSeriesList(timeSeries []TimeSeries) *TimeSeriesList {
	this := TimeSeriesList{}
	this.TimeSeries = timeSeries
	return &this
}

// NewTimeSeriesListWithDefaults instantiates a new TimeSeriesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSeriesListWithDefaults() *TimeSeriesList {
	this := TimeSeriesList{}
	return &this
}

// GetTimeSeries returns the TimeSeries field value
func (o *TimeSeriesList) GetTimeSeries() []TimeSeries {
	if o == nil {
		var ret []TimeSeries
		return ret
	}

	return o.TimeSeries
}

// GetTimeSeriesOk returns a tuple with the TimeSeries field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesList) GetTimeSeriesOk() ([]TimeSeries, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeSeries, true
}

// SetTimeSeries sets field value
func (o *TimeSeriesList) SetTimeSeries(v []TimeSeries) {
	o.TimeSeries = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *TimeSeriesList) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesList) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *TimeSeriesList) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *TimeSeriesList) SetUnit(v string) {
	o.Unit = &v
}

// GetTimeShiftLabel returns the TimeShiftLabel field value if set, zero value otherwise.
func (o *TimeSeriesList) GetTimeShiftLabel() string {
	if o == nil || IsNil(o.TimeShiftLabel) {
		var ret string
		return ret
	}
	return *o.TimeShiftLabel
}

// GetTimeShiftLabelOk returns a tuple with the TimeShiftLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesList) GetTimeShiftLabelOk() (*string, bool) {
	if o == nil || IsNil(o.TimeShiftLabel) {
		return nil, false
	}
	return o.TimeShiftLabel, true
}

// HasTimeShiftLabel returns a boolean if a field has been set.
func (o *TimeSeriesList) HasTimeShiftLabel() bool {
	if o != nil && !IsNil(o.TimeShiftLabel) {
		return true
	}

	return false
}

// SetTimeShiftLabel gets a reference to the given string and assigns it to the TimeShiftLabel field.
func (o *TimeSeriesList) SetTimeShiftLabel(v string) {
	o.TimeShiftLabel = &v
}

// GetResultContext returns the ResultContext field value if set, zero value otherwise.
func (o *TimeSeriesList) GetResultContext() MetricsQueryResultContext {
	if o == nil || IsNil(o.ResultContext) {
		var ret MetricsQueryResultContext
		return ret
	}
	return *o.ResultContext
}

// GetResultContextOk returns a tuple with the ResultContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesList) GetResultContextOk() (*MetricsQueryResultContext, bool) {
	if o == nil || IsNil(o.ResultContext) {
		return nil, false
	}
	return o.ResultContext, true
}

// HasResultContext returns a boolean if a field has been set.
func (o *TimeSeriesList) HasResultContext() bool {
	if o != nil && !IsNil(o.ResultContext) {
		return true
	}

	return false
}

// SetResultContext gets a reference to the given MetricsQueryResultContext and assigns it to the ResultContext field.
func (o *TimeSeriesList) SetResultContext(v MetricsQueryResultContext) {
	o.ResultContext = &v
}

func (o TimeSeriesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSeriesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeSeries"] = o.TimeSeries
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.TimeShiftLabel) {
		toSerialize["timeShiftLabel"] = o.TimeShiftLabel
	}
	if !IsNil(o.ResultContext) {
		toSerialize["resultContext"] = o.ResultContext
	}
	return toSerialize, nil
}

type NullableTimeSeriesList struct {
	value *TimeSeriesList
	isSet bool
}

func (v NullableTimeSeriesList) Get() *TimeSeriesList {
	return v.value
}

func (v *NullableTimeSeriesList) Set(val *TimeSeriesList) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSeriesList) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSeriesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSeriesList(val *TimeSeriesList) *NullableTimeSeriesList {
	return &NullableTimeSeriesList{value: val, isSet: true}
}

func (v NullableTimeSeriesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSeriesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


