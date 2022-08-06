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

// SliQueriesValidationResult Validation result for the SLI queries.
type SliQueriesValidationResult struct {
	// Whether or not if queries are valid.
	IsValid *bool `json:"isValid,omitempty"`
	// Message from validation.
	Message *string `json:"message,omitempty"`
}

// NewSliQueriesValidationResult instantiates a new SliQueriesValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliQueriesValidationResult() *SliQueriesValidationResult {
	this := SliQueriesValidationResult{}
	var message string = ""
	this.Message = &message
	return &this
}

// NewSliQueriesValidationResultWithDefaults instantiates a new SliQueriesValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliQueriesValidationResultWithDefaults() *SliQueriesValidationResult {
	this := SliQueriesValidationResult{}
	var message string = ""
	this.Message = &message
	return &this
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *SliQueriesValidationResult) GetIsValid() bool {
	if o == nil || o.IsValid == nil {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliQueriesValidationResult) GetIsValidOk() (*bool, bool) {
	if o == nil || o.IsValid == nil {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *SliQueriesValidationResult) HasIsValid() bool {
	if o != nil && o.IsValid != nil {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *SliQueriesValidationResult) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SliQueriesValidationResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliQueriesValidationResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SliQueriesValidationResult) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SliQueriesValidationResult) SetMessage(v string) {
	o.Message = &v
}

func (o SliQueriesValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsValid != nil {
		toSerialize["isValid"] = o.IsValid
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableSliQueriesValidationResult struct {
	value *SliQueriesValidationResult
	isSet bool
}

func (v NullableSliQueriesValidationResult) Get() *SliQueriesValidationResult {
	return v.value
}

func (v *NullableSliQueriesValidationResult) Set(val *SliQueriesValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSliQueriesValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSliQueriesValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliQueriesValidationResult(val *SliQueriesValidationResult) *NullableSliQueriesValidationResult {
	return &NullableSliQueriesValidationResult{value: val, isSet: true}
}

func (v NullableSliQueriesValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliQueriesValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


