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

// LogSearchEstimatedUsageDefinitionAllOf struct for LogSearchEstimatedUsageDefinitionAllOf
type LogSearchEstimatedUsageDefinitionAllOf struct {
	EstimatedUsageDetails EstimatedUsageDetails `json:"estimatedUsageDetails"`
}

// NewLogSearchEstimatedUsageDefinitionAllOf instantiates a new LogSearchEstimatedUsageDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchEstimatedUsageDefinitionAllOf(estimatedUsageDetails EstimatedUsageDetails) *LogSearchEstimatedUsageDefinitionAllOf {
	this := LogSearchEstimatedUsageDefinitionAllOf{}
	this.EstimatedUsageDetails = estimatedUsageDetails
	return &this
}

// NewLogSearchEstimatedUsageDefinitionAllOfWithDefaults instantiates a new LogSearchEstimatedUsageDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchEstimatedUsageDefinitionAllOfWithDefaults() *LogSearchEstimatedUsageDefinitionAllOf {
	this := LogSearchEstimatedUsageDefinitionAllOf{}
	return &this
}

// GetEstimatedUsageDetails returns the EstimatedUsageDetails field value
func (o *LogSearchEstimatedUsageDefinitionAllOf) GetEstimatedUsageDetails() EstimatedUsageDetails {
	if o == nil {
		var ret EstimatedUsageDetails
		return ret
	}

	return o.EstimatedUsageDetails
}

// GetEstimatedUsageDetailsOk returns a tuple with the EstimatedUsageDetails field value
// and a boolean to check if the value has been set.
func (o *LogSearchEstimatedUsageDefinitionAllOf) GetEstimatedUsageDetailsOk() (*EstimatedUsageDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedUsageDetails, true
}

// SetEstimatedUsageDetails sets field value
func (o *LogSearchEstimatedUsageDefinitionAllOf) SetEstimatedUsageDetails(v EstimatedUsageDetails) {
	o.EstimatedUsageDetails = v
}

func (o LogSearchEstimatedUsageDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["estimatedUsageDetails"] = o.EstimatedUsageDetails
	}
	return json.Marshal(toSerialize)
}

type NullableLogSearchEstimatedUsageDefinitionAllOf struct {
	value *LogSearchEstimatedUsageDefinitionAllOf
	isSet bool
}

func (v NullableLogSearchEstimatedUsageDefinitionAllOf) Get() *LogSearchEstimatedUsageDefinitionAllOf {
	return v.value
}

func (v *NullableLogSearchEstimatedUsageDefinitionAllOf) Set(val *LogSearchEstimatedUsageDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchEstimatedUsageDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchEstimatedUsageDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchEstimatedUsageDefinitionAllOf(val *LogSearchEstimatedUsageDefinitionAllOf) *NullableLogSearchEstimatedUsageDefinitionAllOf {
	return &NullableLogSearchEstimatedUsageDefinitionAllOf{value: val, isSet: true}
}

func (v NullableLogSearchEstimatedUsageDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchEstimatedUsageDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


