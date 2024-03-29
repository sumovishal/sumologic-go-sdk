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

// checks if the CpcSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpcSummary{}

// CpcSummary struct for CpcSummary
type CpcSummary struct {
	// The total number of traces matching the search criteria for a given service  based on which the CPC data is aggregated.
	NumOfTraces int64 `json:"numOfTraces"`
	// The total fraction (value between 0.0 and 1.0) of the trace duration time consumed  by a given service (or a group of services) in the critical path of analyzed traces.
	AvgPercentageInTrace float64 `json:"avgPercentageInTrace"`
	// The average time in nanoseconds spent by a given service (or a group of services) in the critical path of analyzed traces.
	AvgTimeInTrace float64 `json:"avgTimeInTrace"`
	// The total time in nanoseconds spent by a given service (or a group of services) in the critical path  of analyzed traces.
	TotalTimeTaken int64 `json:"totalTimeTaken"`
}

// NewCpcSummary instantiates a new CpcSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpcSummary(numOfTraces int64, avgPercentageInTrace float64, avgTimeInTrace float64, totalTimeTaken int64) *CpcSummary {
	this := CpcSummary{}
	this.NumOfTraces = numOfTraces
	this.AvgPercentageInTrace = avgPercentageInTrace
	this.AvgTimeInTrace = avgTimeInTrace
	this.TotalTimeTaken = totalTimeTaken
	return &this
}

// NewCpcSummaryWithDefaults instantiates a new CpcSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpcSummaryWithDefaults() *CpcSummary {
	this := CpcSummary{}
	return &this
}

// GetNumOfTraces returns the NumOfTraces field value
func (o *CpcSummary) GetNumOfTraces() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumOfTraces
}

// GetNumOfTracesOk returns a tuple with the NumOfTraces field value
// and a boolean to check if the value has been set.
func (o *CpcSummary) GetNumOfTracesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumOfTraces, true
}

// SetNumOfTraces sets field value
func (o *CpcSummary) SetNumOfTraces(v int64) {
	o.NumOfTraces = v
}

// GetAvgPercentageInTrace returns the AvgPercentageInTrace field value
func (o *CpcSummary) GetAvgPercentageInTrace() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AvgPercentageInTrace
}

// GetAvgPercentageInTraceOk returns a tuple with the AvgPercentageInTrace field value
// and a boolean to check if the value has been set.
func (o *CpcSummary) GetAvgPercentageInTraceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgPercentageInTrace, true
}

// SetAvgPercentageInTrace sets field value
func (o *CpcSummary) SetAvgPercentageInTrace(v float64) {
	o.AvgPercentageInTrace = v
}

// GetAvgTimeInTrace returns the AvgTimeInTrace field value
func (o *CpcSummary) GetAvgTimeInTrace() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AvgTimeInTrace
}

// GetAvgTimeInTraceOk returns a tuple with the AvgTimeInTrace field value
// and a boolean to check if the value has been set.
func (o *CpcSummary) GetAvgTimeInTraceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgTimeInTrace, true
}

// SetAvgTimeInTrace sets field value
func (o *CpcSummary) SetAvgTimeInTrace(v float64) {
	o.AvgTimeInTrace = v
}

// GetTotalTimeTaken returns the TotalTimeTaken field value
func (o *CpcSummary) GetTotalTimeTaken() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalTimeTaken
}

// GetTotalTimeTakenOk returns a tuple with the TotalTimeTaken field value
// and a boolean to check if the value has been set.
func (o *CpcSummary) GetTotalTimeTakenOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTimeTaken, true
}

// SetTotalTimeTaken sets field value
func (o *CpcSummary) SetTotalTimeTaken(v int64) {
	o.TotalTimeTaken = v
}

func (o CpcSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpcSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numOfTraces"] = o.NumOfTraces
	toSerialize["avgPercentageInTrace"] = o.AvgPercentageInTrace
	toSerialize["avgTimeInTrace"] = o.AvgTimeInTrace
	toSerialize["totalTimeTaken"] = o.TotalTimeTaken
	return toSerialize, nil
}

type NullableCpcSummary struct {
	value *CpcSummary
	isSet bool
}

func (v NullableCpcSummary) Get() *CpcSummary {
	return v.value
}

func (v *NullableCpcSummary) Set(val *CpcSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCpcSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCpcSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpcSummary(val *CpcSummary) *NullableCpcSummary {
	return &NullableCpcSummary{value: val, isSet: true}
}

func (v NullableCpcSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpcSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


