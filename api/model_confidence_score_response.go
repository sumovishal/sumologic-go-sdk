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

// ConfidenceScoreResponse CSE insight confidence score.
type ConfidenceScoreResponse struct {
	// List of confidence scores to the CSE Insights.
	ConfidenceScore string `json:"confidenceScore"`
}

// NewConfidenceScoreResponse instantiates a new ConfidenceScoreResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfidenceScoreResponse(confidenceScore string) *ConfidenceScoreResponse {
	this := ConfidenceScoreResponse{}
	this.ConfidenceScore = confidenceScore
	return &this
}

// NewConfidenceScoreResponseWithDefaults instantiates a new ConfidenceScoreResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfidenceScoreResponseWithDefaults() *ConfidenceScoreResponse {
	this := ConfidenceScoreResponse{}
	return &this
}

// GetConfidenceScore returns the ConfidenceScore field value
func (o *ConfidenceScoreResponse) GetConfidenceScore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfidenceScore
}

// GetConfidenceScoreOk returns a tuple with the ConfidenceScore field value
// and a boolean to check if the value has been set.
func (o *ConfidenceScoreResponse) GetConfidenceScoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfidenceScore, true
}

// SetConfidenceScore sets field value
func (o *ConfidenceScoreResponse) SetConfidenceScore(v string) {
	o.ConfidenceScore = v
}

func (o ConfidenceScoreResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["confidenceScore"] = o.ConfidenceScore
	}
	return json.Marshal(toSerialize)
}

type NullableConfidenceScoreResponse struct {
	value *ConfidenceScoreResponse
	isSet bool
}

func (v NullableConfidenceScoreResponse) Get() *ConfidenceScoreResponse {
	return v.value
}

func (v *NullableConfidenceScoreResponse) Set(val *ConfidenceScoreResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfidenceScoreResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfidenceScoreResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfidenceScoreResponse(val *ConfidenceScoreResponse) *NullableConfidenceScoreResponse {
	return &NullableConfidenceScoreResponse{value: val, isSet: true}
}

func (v NullableConfidenceScoreResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfidenceScoreResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


