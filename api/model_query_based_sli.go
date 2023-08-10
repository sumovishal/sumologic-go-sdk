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

// checks if the QueryBasedSli type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryBasedSli{}

// QueryBasedSli Common properties for query based SLIs
type QueryBasedSli struct {
	// Type of Raw Data Queries for SLI (Logs/Metrics).
	QueryType string `json:"queryType"`
	// Queries for defining SLI.
	Queries []SliQueryGroup `json:"queries"`
}

// NewQueryBasedSli instantiates a new QueryBasedSli object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryBasedSli(queryType string, queries []SliQueryGroup) *QueryBasedSli {
	this := QueryBasedSli{}
	this.QueryType = queryType
	this.Queries = queries
	return &this
}

// NewQueryBasedSliWithDefaults instantiates a new QueryBasedSli object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryBasedSliWithDefaults() *QueryBasedSli {
	this := QueryBasedSli{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *QueryBasedSli) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *QueryBasedSli) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *QueryBasedSli) SetQueryType(v string) {
	o.QueryType = v
}

// GetQueries returns the Queries field value
func (o *QueryBasedSli) GetQueries() []SliQueryGroup {
	if o == nil {
		var ret []SliQueryGroup
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *QueryBasedSli) GetQueriesOk() ([]SliQueryGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *QueryBasedSli) SetQueries(v []SliQueryGroup) {
	o.Queries = v
}

func (o QueryBasedSli) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryBasedSli) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryType"] = o.QueryType
	toSerialize["queries"] = o.Queries
	return toSerialize, nil
}

type NullableQueryBasedSli struct {
	value *QueryBasedSli
	isSet bool
}

func (v NullableQueryBasedSli) Get() *QueryBasedSli {
	return v.value
}

func (v *NullableQueryBasedSli) Set(val *QueryBasedSli) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryBasedSli) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryBasedSli) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryBasedSli(val *QueryBasedSli) *NullableQueryBasedSli {
	return &NullableQueryBasedSli{value: val, isSet: true}
}

func (v NullableQueryBasedSli) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryBasedSli) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


