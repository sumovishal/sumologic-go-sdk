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

// LogSearchQueryTimeRangeBaseExceptParsingMode Definition of a log search with query and timerange.
type LogSearchQueryTimeRangeBaseExceptParsingMode struct {
	// Query to perform.
	QueryString string `json:"queryString"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// This has the value `true` if the search is to be run by receipt time and `false` if it is to be run by message time.
	RunByReceiptTime *bool `json:"runByReceiptTime,omitempty"`
	// Definition of the query parameters.
	QueryParameters []LogSearchQueryParameterSyncDefinition `json:"queryParameters,omitempty"`
}

// NewLogSearchQueryTimeRangeBaseExceptParsingMode instantiates a new LogSearchQueryTimeRangeBaseExceptParsingMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchQueryTimeRangeBaseExceptParsingMode(queryString string, timeRange ResolvableTimeRange) *LogSearchQueryTimeRangeBaseExceptParsingMode {
	this := LogSearchQueryTimeRangeBaseExceptParsingMode{}
	this.QueryString = queryString
	this.TimeRange = timeRange
	var runByReceiptTime bool = false
	this.RunByReceiptTime = &runByReceiptTime
	return &this
}

// NewLogSearchQueryTimeRangeBaseExceptParsingModeWithDefaults instantiates a new LogSearchQueryTimeRangeBaseExceptParsingMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchQueryTimeRangeBaseExceptParsingModeWithDefaults() *LogSearchQueryTimeRangeBaseExceptParsingMode {
	this := LogSearchQueryTimeRangeBaseExceptParsingMode{}
	var runByReceiptTime bool = false
	this.RunByReceiptTime = &runByReceiptTime
	return &this
}

// GetQueryString returns the QueryString field value
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) SetQueryString(v string) {
	o.QueryString = v
}

// GetTimeRange returns the TimeRange field value
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetRunByReceiptTime returns the RunByReceiptTime field value if set, zero value otherwise.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetRunByReceiptTime() bool {
	if o == nil || o.RunByReceiptTime == nil {
		var ret bool
		return ret
	}
	return *o.RunByReceiptTime
}

// GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetRunByReceiptTimeOk() (*bool, bool) {
	if o == nil || o.RunByReceiptTime == nil {
		return nil, false
	}
	return o.RunByReceiptTime, true
}

// HasRunByReceiptTime returns a boolean if a field has been set.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) HasRunByReceiptTime() bool {
	if o != nil && o.RunByReceiptTime != nil {
		return true
	}

	return false
}

// SetRunByReceiptTime gets a reference to the given bool and assigns it to the RunByReceiptTime field.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) SetRunByReceiptTime(v bool) {
	o.RunByReceiptTime = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetQueryParameters() []LogSearchQueryParameterSyncDefinition {
	if o == nil || o.QueryParameters == nil {
		var ret []LogSearchQueryParameterSyncDefinition
		return ret
	}
	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetQueryParametersOk() ([]LogSearchQueryParameterSyncDefinition, bool) {
	if o == nil || o.QueryParameters == nil {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) HasQueryParameters() bool {
	if o != nil && o.QueryParameters != nil {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given []LogSearchQueryParameterSyncDefinition and assigns it to the QueryParameters field.
func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) SetQueryParameters(v []LogSearchQueryParameterSyncDefinition) {
	o.QueryParameters = v
}

func (o LogSearchQueryTimeRangeBaseExceptParsingMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryString"] = o.QueryString
	}
	if true {
		toSerialize["timeRange"] = o.TimeRange
	}
	if o.RunByReceiptTime != nil {
		toSerialize["runByReceiptTime"] = o.RunByReceiptTime
	}
	if o.QueryParameters != nil {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	return json.Marshal(toSerialize)
}

type NullableLogSearchQueryTimeRangeBaseExceptParsingMode struct {
	value *LogSearchQueryTimeRangeBaseExceptParsingMode
	isSet bool
}

func (v NullableLogSearchQueryTimeRangeBaseExceptParsingMode) Get() *LogSearchQueryTimeRangeBaseExceptParsingMode {
	return v.value
}

func (v *NullableLogSearchQueryTimeRangeBaseExceptParsingMode) Set(val *LogSearchQueryTimeRangeBaseExceptParsingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchQueryTimeRangeBaseExceptParsingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchQueryTimeRangeBaseExceptParsingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchQueryTimeRangeBaseExceptParsingMode(val *LogSearchQueryTimeRangeBaseExceptParsingMode) *NullableLogSearchQueryTimeRangeBaseExceptParsingMode {
	return &NullableLogSearchQueryTimeRangeBaseExceptParsingMode{value: val, isSet: true}
}

func (v NullableLogSearchQueryTimeRangeBaseExceptParsingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchQueryTimeRangeBaseExceptParsingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


