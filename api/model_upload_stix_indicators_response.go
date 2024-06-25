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

// checks if the UploadStixIndicatorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadStixIndicatorsResponse{}

// UploadStixIndicatorsResponse struct for UploadStixIndicatorsResponse
type UploadStixIndicatorsResponse struct {
	// A list of invalid indicator IDs that were not ingested
	InvalidIndicators []string `json:"invalidIndicators"`
}

// NewUploadStixIndicatorsResponse instantiates a new UploadStixIndicatorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadStixIndicatorsResponse(invalidIndicators []string) *UploadStixIndicatorsResponse {
	this := UploadStixIndicatorsResponse{}
	this.InvalidIndicators = invalidIndicators
	return &this
}

// NewUploadStixIndicatorsResponseWithDefaults instantiates a new UploadStixIndicatorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadStixIndicatorsResponseWithDefaults() *UploadStixIndicatorsResponse {
	this := UploadStixIndicatorsResponse{}
	return &this
}

// GetInvalidIndicators returns the InvalidIndicators field value
func (o *UploadStixIndicatorsResponse) GetInvalidIndicators() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InvalidIndicators
}

// GetInvalidIndicatorsOk returns a tuple with the InvalidIndicators field value
// and a boolean to check if the value has been set.
func (o *UploadStixIndicatorsResponse) GetInvalidIndicatorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvalidIndicators, true
}

// SetInvalidIndicators sets field value
func (o *UploadStixIndicatorsResponse) SetInvalidIndicators(v []string) {
	o.InvalidIndicators = v
}

func (o UploadStixIndicatorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadStixIndicatorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invalidIndicators"] = o.InvalidIndicators
	return toSerialize, nil
}

type NullableUploadStixIndicatorsResponse struct {
	value *UploadStixIndicatorsResponse
	isSet bool
}

func (v NullableUploadStixIndicatorsResponse) Get() *UploadStixIndicatorsResponse {
	return v.value
}

func (v *NullableUploadStixIndicatorsResponse) Set(val *UploadStixIndicatorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadStixIndicatorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadStixIndicatorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadStixIndicatorsResponse(val *UploadStixIndicatorsResponse) *NullableUploadStixIndicatorsResponse {
	return &NullableUploadStixIndicatorsResponse{value: val, isSet: true}
}

func (v NullableUploadStixIndicatorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadStixIndicatorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


