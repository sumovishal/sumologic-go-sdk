/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SearchQueryContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchQueryContext{}

// SearchQueryContext struct for SearchQueryContext
type SearchQueryContext struct {
	EventContext
	// The query id of the log search.
	QueryId string `json:"queryId"`
}

// NewSearchQueryContext instantiates a new SearchQueryContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchQueryContext(queryId string, eventContextType string) *SearchQueryContext {
	this := SearchQueryContext{}
	this.EventContextType = eventContextType
	this.QueryId = queryId
	return &this
}

// NewSearchQueryContextWithDefaults instantiates a new SearchQueryContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchQueryContextWithDefaults() *SearchQueryContext {
	this := SearchQueryContext{}
	return &this
}

// GetQueryId returns the QueryId field value
func (o *SearchQueryContext) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *SearchQueryContext) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *SearchQueryContext) SetQueryId(v string) {
	o.QueryId = v
}

func (o SearchQueryContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchQueryContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEventContext, errEventContext := json.Marshal(o.EventContext)
	if errEventContext != nil {
		return map[string]interface{}{}, errEventContext
	}
	errEventContext = json.Unmarshal([]byte(serializedEventContext), &toSerialize)
	if errEventContext != nil {
		return map[string]interface{}{}, errEventContext
	}
	toSerialize["queryId"] = o.QueryId
	return toSerialize, nil
}

type NullableSearchQueryContext struct {
	value *SearchQueryContext
	isSet bool
}

func (v NullableSearchQueryContext) Get() *SearchQueryContext {
	return v.value
}

func (v *NullableSearchQueryContext) Set(val *SearchQueryContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchQueryContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchQueryContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchQueryContext(val *SearchQueryContext) *NullableSearchQueryContext {
	return &NullableSearchQueryContext{value: val, isSet: true}
}

func (v NullableSearchQueryContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchQueryContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


