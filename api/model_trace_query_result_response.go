/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TraceQueryResultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceQueryResultResponse{}

// TraceQueryResultResponse struct for TraceQueryResultResponse
type TraceQueryResultResponse struct {
	// List of traces matching the query.
	Results []TraceDetail `json:"results"`
	// Next continuation token.
	Next *string `json:"next,omitempty"`
}

type _TraceQueryResultResponse TraceQueryResultResponse

// NewTraceQueryResultResponse instantiates a new TraceQueryResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceQueryResultResponse(results []TraceDetail) *TraceQueryResultResponse {
	this := TraceQueryResultResponse{}
	this.Results = results
	return &this
}

// NewTraceQueryResultResponseWithDefaults instantiates a new TraceQueryResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceQueryResultResponseWithDefaults() *TraceQueryResultResponse {
	this := TraceQueryResultResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *TraceQueryResultResponse) GetResults() []TraceDetail {
	if o == nil {
		var ret []TraceDetail
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *TraceQueryResultResponse) GetResultsOk() ([]TraceDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *TraceQueryResultResponse) SetResults(v []TraceDetail) {
	o.Results = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *TraceQueryResultResponse) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceQueryResultResponse) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *TraceQueryResultResponse) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *TraceQueryResultResponse) SetNext(v string) {
	o.Next = &v
}

func (o TraceQueryResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceQueryResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

func (o *TraceQueryResultResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTraceQueryResultResponse := _TraceQueryResultResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTraceQueryResultResponse)

	if err != nil {
		return err
	}

	*o = TraceQueryResultResponse(varTraceQueryResultResponse)

	return err
}

type NullableTraceQueryResultResponse struct {
	value *TraceQueryResultResponse
	isSet bool
}

func (v NullableTraceQueryResultResponse) Get() *TraceQueryResultResponse {
	return v.value
}

func (v *NullableTraceQueryResultResponse) Set(val *TraceQueryResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceQueryResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceQueryResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceQueryResultResponse(val *TraceQueryResultResponse) *NullableTraceQueryResultResponse {
	return &NullableTraceQueryResultResponse{value: val, isSet: true}
}

func (v NullableTraceQueryResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceQueryResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


