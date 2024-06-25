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

// checks if the Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Request{}

// Request Evaluate SLI using occurrences of successful events over compliance period.
type Request struct {
	// Type of Raw Data Queries for SLI (Logs/Metrics).
	QueryType string `json:"queryType"`
	// Queries for defining SLI.
	Queries []SliQueryGroup `json:"queries"`
	// Compared against threshold query's raw data points to determine success.
	Threshold *float32 `json:"threshold,omitempty"`
	// Comparison function with threshold (LessThan/GreaterThan/LessThanOrEqual/GreaterThanOrEqual).
	Op *string `json:"op,omitempty"`
}

// NewRequest instantiates a new Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequest(queryType string, queries []SliQueryGroup, evaluationType string) *Request {
	this := Request{}
	this.QueryType = queryType
	this.Queries = queries
	this.EvaluationType = evaluationType
	return &this
}

// NewRequestWithDefaults instantiates a new Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestWithDefaults() *Request {
	this := Request{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *Request) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *Request) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *Request) SetQueryType(v string) {
	o.QueryType = v
}

// GetQueries returns the Queries field value
func (o *Request) GetQueries() []SliQueryGroup {
	if o == nil {
		var ret []SliQueryGroup
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *Request) GetQueriesOk() ([]SliQueryGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *Request) SetQueries(v []SliQueryGroup) {
	o.Queries = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *Request) GetThreshold() float32 {
	if o == nil || IsNil(o.Threshold) {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetThresholdOk() (*float32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *Request) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *Request) SetThreshold(v float32) {
	o.Threshold = &v
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *Request) GetOp() string {
	if o == nil || IsNil(o.Op) {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetOpOk() (*string, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *Request) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *Request) SetOp(v string) {
	o.Op = &v
}

func (o Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryType"] = o.QueryType
	toSerialize["queries"] = o.Queries
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	return toSerialize, nil
}

type NullableRequest struct {
	value *Request
	isSet bool
}

func (v NullableRequest) Get() *Request {
	return v.value
}

func (v *NullableRequest) Set(val *Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequest(val *Request) *NullableRequest {
	return &NullableRequest{value: val, isSet: true}
}

func (v NullableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


