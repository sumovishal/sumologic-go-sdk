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

// checks if the CseInsightConfidenceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CseInsightConfidenceRequest{}

// CseInsightConfidenceRequest CSE insight JSON object.
type CseInsightConfidenceRequest struct {
	// List of CSE Insight Created logs for which the confidence score should be calculated.
	CseInsight string `json:"cseInsight"`
}

// NewCseInsightConfidenceRequest instantiates a new CseInsightConfidenceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCseInsightConfidenceRequest(cseInsight string) *CseInsightConfidenceRequest {
	this := CseInsightConfidenceRequest{}
	this.CseInsight = cseInsight
	return &this
}

// NewCseInsightConfidenceRequestWithDefaults instantiates a new CseInsightConfidenceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCseInsightConfidenceRequestWithDefaults() *CseInsightConfidenceRequest {
	this := CseInsightConfidenceRequest{}
	return &this
}

// GetCseInsight returns the CseInsight field value
func (o *CseInsightConfidenceRequest) GetCseInsight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CseInsight
}

// GetCseInsightOk returns a tuple with the CseInsight field value
// and a boolean to check if the value has been set.
func (o *CseInsightConfidenceRequest) GetCseInsightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CseInsight, true
}

// SetCseInsight sets field value
func (o *CseInsightConfidenceRequest) SetCseInsight(v string) {
	o.CseInsight = v
}

func (o CseInsightConfidenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CseInsightConfidenceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cseInsight"] = o.CseInsight
	return toSerialize, nil
}

type NullableCseInsightConfidenceRequest struct {
	value *CseInsightConfidenceRequest
	isSet bool
}

func (v NullableCseInsightConfidenceRequest) Get() *CseInsightConfidenceRequest {
	return v.value
}

func (v *NullableCseInsightConfidenceRequest) Set(val *CseInsightConfidenceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCseInsightConfidenceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCseInsightConfidenceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCseInsightConfidenceRequest(val *CseInsightConfidenceRequest) *NullableCseInsightConfidenceRequest {
	return &NullableCseInsightConfidenceRequest{value: val, isSet: true}
}

func (v NullableCseInsightConfidenceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCseInsightConfidenceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


