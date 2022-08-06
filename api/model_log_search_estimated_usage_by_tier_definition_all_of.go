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

// LogSearchEstimatedUsageByTierDefinitionAllOf struct for LogSearchEstimatedUsageByTierDefinitionAllOf
type LogSearchEstimatedUsageByTierDefinitionAllOf struct {
	EstimatedUsageDetails []EstimatedUsageDetailsWithTier `json:"estimatedUsageDetails"`
}

// NewLogSearchEstimatedUsageByTierDefinitionAllOf instantiates a new LogSearchEstimatedUsageByTierDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchEstimatedUsageByTierDefinitionAllOf(estimatedUsageDetails []EstimatedUsageDetailsWithTier) *LogSearchEstimatedUsageByTierDefinitionAllOf {
	this := LogSearchEstimatedUsageByTierDefinitionAllOf{}
	this.EstimatedUsageDetails = estimatedUsageDetails
	return &this
}

// NewLogSearchEstimatedUsageByTierDefinitionAllOfWithDefaults instantiates a new LogSearchEstimatedUsageByTierDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchEstimatedUsageByTierDefinitionAllOfWithDefaults() *LogSearchEstimatedUsageByTierDefinitionAllOf {
	this := LogSearchEstimatedUsageByTierDefinitionAllOf{}
	return &this
}

// GetEstimatedUsageDetails returns the EstimatedUsageDetails field value
func (o *LogSearchEstimatedUsageByTierDefinitionAllOf) GetEstimatedUsageDetails() []EstimatedUsageDetailsWithTier {
	if o == nil {
		var ret []EstimatedUsageDetailsWithTier
		return ret
	}

	return o.EstimatedUsageDetails
}

// GetEstimatedUsageDetailsOk returns a tuple with the EstimatedUsageDetails field value
// and a boolean to check if the value has been set.
func (o *LogSearchEstimatedUsageByTierDefinitionAllOf) GetEstimatedUsageDetailsOk() ([]EstimatedUsageDetailsWithTier, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedUsageDetails, true
}

// SetEstimatedUsageDetails sets field value
func (o *LogSearchEstimatedUsageByTierDefinitionAllOf) SetEstimatedUsageDetails(v []EstimatedUsageDetailsWithTier) {
	o.EstimatedUsageDetails = v
}

func (o LogSearchEstimatedUsageByTierDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["estimatedUsageDetails"] = o.EstimatedUsageDetails
	}
	return json.Marshal(toSerialize)
}

type NullableLogSearchEstimatedUsageByTierDefinitionAllOf struct {
	value *LogSearchEstimatedUsageByTierDefinitionAllOf
	isSet bool
}

func (v NullableLogSearchEstimatedUsageByTierDefinitionAllOf) Get() *LogSearchEstimatedUsageByTierDefinitionAllOf {
	return v.value
}

func (v *NullableLogSearchEstimatedUsageByTierDefinitionAllOf) Set(val *LogSearchEstimatedUsageByTierDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchEstimatedUsageByTierDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchEstimatedUsageByTierDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchEstimatedUsageByTierDefinitionAllOf(val *LogSearchEstimatedUsageByTierDefinitionAllOf) *NullableLogSearchEstimatedUsageByTierDefinitionAllOf {
	return &NullableLogSearchEstimatedUsageByTierDefinitionAllOf{value: val, isSet: true}
}

func (v NullableLogSearchEstimatedUsageByTierDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchEstimatedUsageByTierDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


