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

// checks if the TransformationRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformationRulesResponse{}

// TransformationRulesResponse A generic response for transformation rule.
type TransformationRulesResponse struct {
	// List of transformation rules.
	Data []TransformationRuleResponse `json:"data"`
	// Next continuation token.
	Next *string `json:"next,omitempty"`
}

// NewTransformationRulesResponse instantiates a new TransformationRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformationRulesResponse(data []TransformationRuleResponse) *TransformationRulesResponse {
	this := TransformationRulesResponse{}
	this.Data = data
	return &this
}

// NewTransformationRulesResponseWithDefaults instantiates a new TransformationRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformationRulesResponseWithDefaults() *TransformationRulesResponse {
	this := TransformationRulesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TransformationRulesResponse) GetData() []TransformationRuleResponse {
	if o == nil {
		var ret []TransformationRuleResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransformationRulesResponse) GetDataOk() ([]TransformationRuleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *TransformationRulesResponse) SetData(v []TransformationRuleResponse) {
	o.Data = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *TransformationRulesResponse) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationRulesResponse) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *TransformationRulesResponse) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *TransformationRulesResponse) SetNext(v string) {
	o.Next = &v
}

func (o TransformationRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformationRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableTransformationRulesResponse struct {
	value *TransformationRulesResponse
	isSet bool
}

func (v NullableTransformationRulesResponse) Get() *TransformationRulesResponse {
	return v.value
}

func (v *NullableTransformationRulesResponse) Set(val *TransformationRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationRulesResponse(val *TransformationRulesResponse) *NullableTransformationRulesResponse {
	return &NullableTransformationRulesResponse{value: val, isSet: true}
}

func (v NullableTransformationRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


