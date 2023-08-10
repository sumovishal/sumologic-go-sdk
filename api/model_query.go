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

// checks if the Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Query{}

// Query struct for Query
type Query struct {
	// The metrics, traces or logs query.
	QueryString string `json:"queryString"`
	// The type of the query, either `Metrics`, `Traces`, `Spans` or `Logs`.
	QueryType string `json:"queryType"`
	// The key for metric, traces or log queries. Used as an identifier for queries. It is displayed on the panel builder and used for display overrides and query toggling. 
	QueryKey string `json:"queryKey"`
	// The mode of the metrics query that the user was editing. Can be `Basic` or `Advanced`. Will ONLY be specified for metrics queries. 
	MetricsQueryMode *string `json:"metricsQueryMode,omitempty"`
	MetricsQueryData *MetricsQueryData `json:"metricsQueryData,omitempty"`
	TracesQueryData *TracesQueryData `json:"tracesQueryData,omitempty"`
	SpansQueryData *SpansQueryData `json:"spansQueryData,omitempty"`
	// This field only applies for queryType of `Logs` but other query types may be supported in the future. Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `Auto`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParseMode *string `json:"parseMode,omitempty"`
	// This field only applies for queryType of `Logs` but other query types may be supported in the future. Define the time source of this query. Possible values are `Message` and `Receipt`. `Message` will use the timeStamp on the message, while `Receipt` will use the timestamp it was received by Sumo.
	TimeSource *string `json:"timeSource,omitempty"`
	// This field only applies for queryType of `Metrics` but other query types may be supported in the future. Determines if the row should be returned in the response. Can be used in conjunction with a join, if only the result of the join is needed, and not the intermediate rows. Setting `transient` to `true`  wherever the intermediate results aren't required speeds up the computation and reduces the amount of data  transferred over the network.
	Transient *bool `json:"transient,omitempty"`
	// This field only applies for queryType of `Metrics` but other query types may be supported in the future. Specifies the output cardinality limitations for the query, which is the maximum number of timeseries returned in the result.
	OutputCardinalityLimit *int32 `json:"outputCardinalityLimit,omitempty"`
}

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery(queryString string, queryType string, queryKey string) *Query {
	this := Query{}
	this.QueryString = queryString
	this.QueryType = queryType
	this.QueryKey = queryKey
	var parseMode string = "Auto"
	this.ParseMode = &parseMode
	var timeSource string = "Message"
	this.TimeSource = &timeSource
	var transient bool = false
	this.Transient = &transient
	var outputCardinalityLimit int32 = 1000
	this.OutputCardinalityLimit = &outputCardinalityLimit
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	var parseMode string = "Auto"
	this.ParseMode = &parseMode
	var timeSource string = "Message"
	this.TimeSource = &timeSource
	var transient bool = false
	this.Transient = &transient
	var outputCardinalityLimit int32 = 1000
	this.OutputCardinalityLimit = &outputCardinalityLimit
	return &this
}

// GetQueryString returns the QueryString field value
func (o *Query) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *Query) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *Query) SetQueryString(v string) {
	o.QueryString = v
}

// GetQueryType returns the QueryType field value
func (o *Query) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *Query) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *Query) SetQueryType(v string) {
	o.QueryType = v
}

// GetQueryKey returns the QueryKey field value
func (o *Query) GetQueryKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryKey
}

// GetQueryKeyOk returns a tuple with the QueryKey field value
// and a boolean to check if the value has been set.
func (o *Query) GetQueryKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryKey, true
}

// SetQueryKey sets field value
func (o *Query) SetQueryKey(v string) {
	o.QueryKey = v
}

// GetMetricsQueryMode returns the MetricsQueryMode field value if set, zero value otherwise.
func (o *Query) GetMetricsQueryMode() string {
	if o == nil || IsNil(o.MetricsQueryMode) {
		var ret string
		return ret
	}
	return *o.MetricsQueryMode
}

// GetMetricsQueryModeOk returns a tuple with the MetricsQueryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetMetricsQueryModeOk() (*string, bool) {
	if o == nil || IsNil(o.MetricsQueryMode) {
		return nil, false
	}
	return o.MetricsQueryMode, true
}

// HasMetricsQueryMode returns a boolean if a field has been set.
func (o *Query) HasMetricsQueryMode() bool {
	if o != nil && !IsNil(o.MetricsQueryMode) {
		return true
	}

	return false
}

// SetMetricsQueryMode gets a reference to the given string and assigns it to the MetricsQueryMode field.
func (o *Query) SetMetricsQueryMode(v string) {
	o.MetricsQueryMode = &v
}

// GetMetricsQueryData returns the MetricsQueryData field value if set, zero value otherwise.
func (o *Query) GetMetricsQueryData() MetricsQueryData {
	if o == nil || IsNil(o.MetricsQueryData) {
		var ret MetricsQueryData
		return ret
	}
	return *o.MetricsQueryData
}

// GetMetricsQueryDataOk returns a tuple with the MetricsQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetMetricsQueryDataOk() (*MetricsQueryData, bool) {
	if o == nil || IsNil(o.MetricsQueryData) {
		return nil, false
	}
	return o.MetricsQueryData, true
}

// HasMetricsQueryData returns a boolean if a field has been set.
func (o *Query) HasMetricsQueryData() bool {
	if o != nil && !IsNil(o.MetricsQueryData) {
		return true
	}

	return false
}

// SetMetricsQueryData gets a reference to the given MetricsQueryData and assigns it to the MetricsQueryData field.
func (o *Query) SetMetricsQueryData(v MetricsQueryData) {
	o.MetricsQueryData = &v
}

// GetTracesQueryData returns the TracesQueryData field value if set, zero value otherwise.
func (o *Query) GetTracesQueryData() TracesQueryData {
	if o == nil || IsNil(o.TracesQueryData) {
		var ret TracesQueryData
		return ret
	}
	return *o.TracesQueryData
}

// GetTracesQueryDataOk returns a tuple with the TracesQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTracesQueryDataOk() (*TracesQueryData, bool) {
	if o == nil || IsNil(o.TracesQueryData) {
		return nil, false
	}
	return o.TracesQueryData, true
}

// HasTracesQueryData returns a boolean if a field has been set.
func (o *Query) HasTracesQueryData() bool {
	if o != nil && !IsNil(o.TracesQueryData) {
		return true
	}

	return false
}

// SetTracesQueryData gets a reference to the given TracesQueryData and assigns it to the TracesQueryData field.
func (o *Query) SetTracesQueryData(v TracesQueryData) {
	o.TracesQueryData = &v
}

// GetSpansQueryData returns the SpansQueryData field value if set, zero value otherwise.
func (o *Query) GetSpansQueryData() SpansQueryData {
	if o == nil || IsNil(o.SpansQueryData) {
		var ret SpansQueryData
		return ret
	}
	return *o.SpansQueryData
}

// GetSpansQueryDataOk returns a tuple with the SpansQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetSpansQueryDataOk() (*SpansQueryData, bool) {
	if o == nil || IsNil(o.SpansQueryData) {
		return nil, false
	}
	return o.SpansQueryData, true
}

// HasSpansQueryData returns a boolean if a field has been set.
func (o *Query) HasSpansQueryData() bool {
	if o != nil && !IsNil(o.SpansQueryData) {
		return true
	}

	return false
}

// SetSpansQueryData gets a reference to the given SpansQueryData and assigns it to the SpansQueryData field.
func (o *Query) SetSpansQueryData(v SpansQueryData) {
	o.SpansQueryData = &v
}

// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *Query) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *Query) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *Query) SetParseMode(v string) {
	o.ParseMode = &v
}

// GetTimeSource returns the TimeSource field value if set, zero value otherwise.
func (o *Query) GetTimeSource() string {
	if o == nil || IsNil(o.TimeSource) {
		var ret string
		return ret
	}
	return *o.TimeSource
}

// GetTimeSourceOk returns a tuple with the TimeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTimeSourceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeSource) {
		return nil, false
	}
	return o.TimeSource, true
}

// HasTimeSource returns a boolean if a field has been set.
func (o *Query) HasTimeSource() bool {
	if o != nil && !IsNil(o.TimeSource) {
		return true
	}

	return false
}

// SetTimeSource gets a reference to the given string and assigns it to the TimeSource field.
func (o *Query) SetTimeSource(v string) {
	o.TimeSource = &v
}

// GetTransient returns the Transient field value if set, zero value otherwise.
func (o *Query) GetTransient() bool {
	if o == nil || IsNil(o.Transient) {
		var ret bool
		return ret
	}
	return *o.Transient
}

// GetTransientOk returns a tuple with the Transient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTransientOk() (*bool, bool) {
	if o == nil || IsNil(o.Transient) {
		return nil, false
	}
	return o.Transient, true
}

// HasTransient returns a boolean if a field has been set.
func (o *Query) HasTransient() bool {
	if o != nil && !IsNil(o.Transient) {
		return true
	}

	return false
}

// SetTransient gets a reference to the given bool and assigns it to the Transient field.
func (o *Query) SetTransient(v bool) {
	o.Transient = &v
}

// GetOutputCardinalityLimit returns the OutputCardinalityLimit field value if set, zero value otherwise.
func (o *Query) GetOutputCardinalityLimit() int32 {
	if o == nil || IsNil(o.OutputCardinalityLimit) {
		var ret int32
		return ret
	}
	return *o.OutputCardinalityLimit
}

// GetOutputCardinalityLimitOk returns a tuple with the OutputCardinalityLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetOutputCardinalityLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.OutputCardinalityLimit) {
		return nil, false
	}
	return o.OutputCardinalityLimit, true
}

// HasOutputCardinalityLimit returns a boolean if a field has been set.
func (o *Query) HasOutputCardinalityLimit() bool {
	if o != nil && !IsNil(o.OutputCardinalityLimit) {
		return true
	}

	return false
}

// SetOutputCardinalityLimit gets a reference to the given int32 and assigns it to the OutputCardinalityLimit field.
func (o *Query) SetOutputCardinalityLimit(v int32) {
	o.OutputCardinalityLimit = &v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryString"] = o.QueryString
	toSerialize["queryType"] = o.QueryType
	toSerialize["queryKey"] = o.QueryKey
	if !IsNil(o.MetricsQueryMode) {
		toSerialize["metricsQueryMode"] = o.MetricsQueryMode
	}
	if !IsNil(o.MetricsQueryData) {
		toSerialize["metricsQueryData"] = o.MetricsQueryData
	}
	if !IsNil(o.TracesQueryData) {
		toSerialize["tracesQueryData"] = o.TracesQueryData
	}
	if !IsNil(o.SpansQueryData) {
		toSerialize["spansQueryData"] = o.SpansQueryData
	}
	if !IsNil(o.ParseMode) {
		toSerialize["parseMode"] = o.ParseMode
	}
	if !IsNil(o.TimeSource) {
		toSerialize["timeSource"] = o.TimeSource
	}
	if !IsNil(o.Transient) {
		toSerialize["transient"] = o.Transient
	}
	if !IsNil(o.OutputCardinalityLimit) {
		toSerialize["outputCardinalityLimit"] = o.OutputCardinalityLimit
	}
	return toSerialize, nil
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


