/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// OpenInQuery Raw data query for the computed signal.
type OpenInQuery struct {
	Query Query `json:"query"`
	// Start time of the query.
	StartTime time.Time `json:"startTime"`
	// End time of the query.
	EndTime time.Time `json:"endTime"`
}

// NewOpenInQuery instantiates a new OpenInQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInQuery(query Query, startTime time.Time, endTime time.Time) *OpenInQuery {
	this := OpenInQuery{}
	this.Query = query
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewOpenInQueryWithDefaults instantiates a new OpenInQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInQueryWithDefaults() *OpenInQuery {
	this := OpenInQuery{}
	return &this
}

// GetQuery returns the Query field value
func (o *OpenInQuery) GetQuery() Query {
	if o == nil {
		var ret Query
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *OpenInQuery) GetQueryOk() (*Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *OpenInQuery) SetQuery(v Query) {
	o.Query = v
}

// GetStartTime returns the StartTime field value
func (o *OpenInQuery) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *OpenInQuery) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *OpenInQuery) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *OpenInQuery) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *OpenInQuery) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *OpenInQuery) SetEndTime(v time.Time) {
	o.EndTime = v
}

func (o OpenInQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableOpenInQuery struct {
	value *OpenInQuery
	isSet bool
}

func (v NullableOpenInQuery) Get() *OpenInQuery {
	return v.value
}

func (v *NullableOpenInQuery) Set(val *OpenInQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInQuery(val *OpenInQuery) *NullableOpenInQuery {
	return &NullableOpenInQuery{value: val, isSet: true}
}

func (v NullableOpenInQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


