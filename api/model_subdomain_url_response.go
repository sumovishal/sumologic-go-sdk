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

// checks if the SubdomainUrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubdomainUrlResponse{}

// SubdomainUrlResponse struct for SubdomainUrlResponse
type SubdomainUrlResponse struct {
	// Login URL corresponding to the subdomain.
	Url string `json:"url"`
}

// NewSubdomainUrlResponse instantiates a new SubdomainUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubdomainUrlResponse(url string) *SubdomainUrlResponse {
	this := SubdomainUrlResponse{}
	this.Url = url
	return &this
}

// NewSubdomainUrlResponseWithDefaults instantiates a new SubdomainUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubdomainUrlResponseWithDefaults() *SubdomainUrlResponse {
	this := SubdomainUrlResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *SubdomainUrlResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SubdomainUrlResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SubdomainUrlResponse) SetUrl(v string) {
	o.Url = v
}

func (o SubdomainUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubdomainUrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableSubdomainUrlResponse struct {
	value *SubdomainUrlResponse
	isSet bool
}

func (v NullableSubdomainUrlResponse) Get() *SubdomainUrlResponse {
	return v.value
}

func (v *NullableSubdomainUrlResponse) Set(val *SubdomainUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubdomainUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubdomainUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubdomainUrlResponse(val *SubdomainUrlResponse) *NullableSubdomainUrlResponse {
	return &NullableSubdomainUrlResponse{value: val, isSet: true}
}

func (v NullableSubdomainUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubdomainUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


