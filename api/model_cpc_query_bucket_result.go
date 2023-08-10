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

// checks if the CpcQueryBucketResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpcQueryBucketResult{}

// CpcQueryBucketResult struct for CpcQueryBucketResult
type CpcQueryBucketResult struct {
	// A unique identifier of a time bucket.
	BucketId string `json:"bucketId"`
	// A start of the time bucket in the [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartTimestamp time.Time `json:"startTimestamp"`
	// The length of a time bucket expressed in milliseconds.
	Length int64 `json:"length"`
	// The total number of traces matching the search criteria based on which the CPC data is aggregated.
	TotalNumOfTraces int64 `json:"totalNumOfTraces"`
	// The average duration in nanoseconds of the traces matching the search criteria  based on which the CPC data is aggregated.
	AvgTraceDuration float64 `json:"avgTraceDuration"`
	// The summary of aggregated Critical Path Contribution data on a per service basis.  Each element of the array corresponds to a summary for a specific service.
	PerServiceCpcSummaries []CpcServiceSummary `json:"perServiceCpcSummaries"`
	OtherServicesCpcSummary CpcSummary `json:"otherServicesCpcSummary"`
	IdleTimeCpcSummary CpcSummary `json:"idleTimeCpcSummary"`
}

// NewCpcQueryBucketResult instantiates a new CpcQueryBucketResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpcQueryBucketResult(bucketId string, startTimestamp time.Time, length int64, totalNumOfTraces int64, avgTraceDuration float64, perServiceCpcSummaries []CpcServiceSummary, otherServicesCpcSummary CpcSummary, idleTimeCpcSummary CpcSummary) *CpcQueryBucketResult {
	this := CpcQueryBucketResult{}
	this.BucketId = bucketId
	this.StartTimestamp = startTimestamp
	this.Length = length
	this.TotalNumOfTraces = totalNumOfTraces
	this.AvgTraceDuration = avgTraceDuration
	this.PerServiceCpcSummaries = perServiceCpcSummaries
	this.OtherServicesCpcSummary = otherServicesCpcSummary
	this.IdleTimeCpcSummary = idleTimeCpcSummary
	return &this
}

// NewCpcQueryBucketResultWithDefaults instantiates a new CpcQueryBucketResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpcQueryBucketResultWithDefaults() *CpcQueryBucketResult {
	this := CpcQueryBucketResult{}
	return &this
}

// GetBucketId returns the BucketId field value
func (o *CpcQueryBucketResult) GetBucketId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketId
}

// GetBucketIdOk returns a tuple with the BucketId field value
// and a boolean to check if the value has been set.
func (o *CpcQueryBucketResult) GetBucketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketId, true
}

// SetBucketId sets field value
func (o *CpcQueryBucketResult) SetBucketId(v string) {
	o.BucketId = v
}

// GetStartTimestamp returns the StartTimestamp field value
func (o *CpcQueryBucketResult) GetStartTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value
// and a boolean to check if the value has been set.
func (o *CpcQueryBucketResult) GetStartTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimestamp, true
}

// SetStartTimestamp sets field value
func (o *CpcQueryBucketResult) SetStartTimestamp(v time.Time) {
	o.StartTimestamp = v
}

// GetLength returns the Length field value
func (o *CpcQueryBucketResult) GetLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *CpcQueryBucketResult) GetLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *CpcQueryBucketResult) SetLength(v int64) {
	o.Length = v
}

// GetTotalNumOfTraces returns the TotalNumOfTraces field value
func (o *CpcQueryBucketResult) GetTotalNumOfTraces() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalNumOfTraces
}

// GetTotalNumOfTracesOk returns a tuple with the TotalNumOfTraces field value
// and a boolean to check if the value has been set.
func (o *CpcQueryBucketResult) GetTotalNumOfTracesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNumOfTraces, true
}

// SetTotalNumOfTraces sets field value
func (o *CpcQueryBucketResult) SetTotalNumOfTraces(v int64) {
	o.TotalNumOfTraces = v
}

// GetAvgTraceDuration returns the AvgTraceDuration field value
func (o *CpcQueryBucketResult) GetAvgTraceDuration() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AvgTraceDuration
}

// GetAvgTraceDurationOk returns a tuple with the AvgTraceDuration field value
// and a boolean to check if the value has been set.
func (o *CpcQueryBucketResult) GetAvgTraceDurationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgTraceDuration, true
}

// SetAvgTraceDuration sets field value
func (o *CpcQueryBucketResult) SetAvgTraceDuration(v float64) {
	o.AvgTraceDuration = v
}

// GetPerServiceCpcSummaries returns the PerServiceCpcSummaries field value
func (o *CpcQueryBucketResult) GetPerServiceCpcSummaries() []CpcServiceSummary {
	if o == nil {
		var ret []CpcServiceSummary
		return ret
	}

	return o.PerServiceCpcSummaries
}

// GetPerServiceCpcSummariesOk returns a tuple with the PerServiceCpcSummaries field value
// and a boolean to check if the value has been set.
func (o *CpcQueryBucketResult) GetPerServiceCpcSummariesOk() ([]CpcServiceSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerServiceCpcSummaries, true
}

// SetPerServiceCpcSummaries sets field value
func (o *CpcQueryBucketResult) SetPerServiceCpcSummaries(v []CpcServiceSummary) {
	o.PerServiceCpcSummaries = v
}

// GetOtherServicesCpcSummary returns the OtherServicesCpcSummary field value
func (o *CpcQueryBucketResult) GetOtherServicesCpcSummary() CpcSummary {
	if o == nil {
		var ret CpcSummary
		return ret
	}

	return o.OtherServicesCpcSummary
}

// GetOtherServicesCpcSummaryOk returns a tuple with the OtherServicesCpcSummary field value
// and a boolean to check if the value has been set.
func (o *CpcQueryBucketResult) GetOtherServicesCpcSummaryOk() (*CpcSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtherServicesCpcSummary, true
}

// SetOtherServicesCpcSummary sets field value
func (o *CpcQueryBucketResult) SetOtherServicesCpcSummary(v CpcSummary) {
	o.OtherServicesCpcSummary = v
}

// GetIdleTimeCpcSummary returns the IdleTimeCpcSummary field value
func (o *CpcQueryBucketResult) GetIdleTimeCpcSummary() CpcSummary {
	if o == nil {
		var ret CpcSummary
		return ret
	}

	return o.IdleTimeCpcSummary
}

// GetIdleTimeCpcSummaryOk returns a tuple with the IdleTimeCpcSummary field value
// and a boolean to check if the value has been set.
func (o *CpcQueryBucketResult) GetIdleTimeCpcSummaryOk() (*CpcSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdleTimeCpcSummary, true
}

// SetIdleTimeCpcSummary sets field value
func (o *CpcQueryBucketResult) SetIdleTimeCpcSummary(v CpcSummary) {
	o.IdleTimeCpcSummary = v
}

func (o CpcQueryBucketResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpcQueryBucketResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucketId"] = o.BucketId
	toSerialize["startTimestamp"] = o.StartTimestamp
	toSerialize["length"] = o.Length
	toSerialize["totalNumOfTraces"] = o.TotalNumOfTraces
	toSerialize["avgTraceDuration"] = o.AvgTraceDuration
	toSerialize["perServiceCpcSummaries"] = o.PerServiceCpcSummaries
	toSerialize["otherServicesCpcSummary"] = o.OtherServicesCpcSummary
	toSerialize["idleTimeCpcSummary"] = o.IdleTimeCpcSummary
	return toSerialize, nil
}

type NullableCpcQueryBucketResult struct {
	value *CpcQueryBucketResult
	isSet bool
}

func (v NullableCpcQueryBucketResult) Get() *CpcQueryBucketResult {
	return v.value
}

func (v *NullableCpcQueryBucketResult) Set(val *CpcQueryBucketResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCpcQueryBucketResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCpcQueryBucketResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpcQueryBucketResult(val *CpcQueryBucketResult) *NullableCpcQueryBucketResult {
	return &NullableCpcQueryBucketResult{value: val, isSet: true}
}

func (v NullableCpcQueryBucketResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpcQueryBucketResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


