/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// checks if the SignalsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignalsResponse{}

// SignalsResponse Signal response object.
type SignalsResponse struct {
	// The type of the signal to compute. Can be `LogFluctuation`, `DimensionalityExplanation`, `GisBenchmark` or `Anomalies` 
	SignalType string `json:"signalType"`
	// The id for the signal result in hex format.
	SignalId string `json:"signalId"`
	// Start time of the signal.
	StartTime time.Time `json:"startTime"`
	// End time of the signal.
	EndTime time.Time `json:"endTime"`
	// Description of the payload.
	Summary string `json:"summary"`
	// Json string for computed signal.
	Payload string `json:"payload"`
	// Raw data queries for the computed signal.
	OpenInQueries []OpenInQuery `json:"openInQueries"`
}

// NewSignalsResponse instantiates a new SignalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalsResponse(signalType string, signalId string, startTime time.Time, endTime time.Time, summary string, payload string, openInQueries []OpenInQuery) *SignalsResponse {
	this := SignalsResponse{}
	this.SignalType = signalType
	this.SignalId = signalId
	this.StartTime = startTime
	this.EndTime = endTime
	this.Summary = summary
	this.Payload = payload
	this.OpenInQueries = openInQueries
	return &this
}

// NewSignalsResponseWithDefaults instantiates a new SignalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalsResponseWithDefaults() *SignalsResponse {
	this := SignalsResponse{}
	return &this
}

// GetSignalType returns the SignalType field value
func (o *SignalsResponse) GetSignalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value
// and a boolean to check if the value has been set.
func (o *SignalsResponse) GetSignalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalType, true
}

// SetSignalType sets field value
func (o *SignalsResponse) SetSignalType(v string) {
	o.SignalType = v
}

// GetSignalId returns the SignalId field value
func (o *SignalsResponse) GetSignalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value
// and a boolean to check if the value has been set.
func (o *SignalsResponse) GetSignalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalId, true
}

// SetSignalId sets field value
func (o *SignalsResponse) SetSignalId(v string) {
	o.SignalId = v
}

// GetStartTime returns the StartTime field value
func (o *SignalsResponse) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *SignalsResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *SignalsResponse) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *SignalsResponse) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SignalsResponse) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *SignalsResponse) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetSummary returns the Summary field value
func (o *SignalsResponse) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *SignalsResponse) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *SignalsResponse) SetSummary(v string) {
	o.Summary = v
}

// GetPayload returns the Payload field value
func (o *SignalsResponse) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *SignalsResponse) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *SignalsResponse) SetPayload(v string) {
	o.Payload = v
}

// GetOpenInQueries returns the OpenInQueries field value
func (o *SignalsResponse) GetOpenInQueries() []OpenInQuery {
	if o == nil {
		var ret []OpenInQuery
		return ret
	}

	return o.OpenInQueries
}

// GetOpenInQueriesOk returns a tuple with the OpenInQueries field value
// and a boolean to check if the value has been set.
func (o *SignalsResponse) GetOpenInQueriesOk() ([]OpenInQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenInQueries, true
}

// SetOpenInQueries sets field value
func (o *SignalsResponse) SetOpenInQueries(v []OpenInQuery) {
	o.OpenInQueries = v
}

func (o SignalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignalsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signalType"] = o.SignalType
	toSerialize["signalId"] = o.SignalId
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	toSerialize["summary"] = o.Summary
	toSerialize["payload"] = o.Payload
	toSerialize["openInQueries"] = o.OpenInQueries
	return toSerialize, nil
}

type NullableSignalsResponse struct {
	value *SignalsResponse
	isSet bool
}

func (v NullableSignalsResponse) Get() *SignalsResponse {
	return v.value
}

func (v *NullableSignalsResponse) Set(val *SignalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalsResponse(val *SignalsResponse) *NullableSignalsResponse {
	return &NullableSignalsResponse{value: val, isSet: true}
}

func (v NullableSignalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


