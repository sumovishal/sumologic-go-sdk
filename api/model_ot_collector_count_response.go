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

// checks if the OTCollectorCountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCollectorCountResponse{}

// OTCollectorCountResponse response for total count of otCollectors.
type OTCollectorCountResponse struct {
	// Total number of OT Collector for a customer.
	TotalCount int32 `json:"totalCount"`
}

// NewOTCollectorCountResponse instantiates a new OTCollectorCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCollectorCountResponse(totalCount int32) *OTCollectorCountResponse {
	this := OTCollectorCountResponse{}
	this.TotalCount = totalCount
	return &this
}

// NewOTCollectorCountResponseWithDefaults instantiates a new OTCollectorCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCollectorCountResponseWithDefaults() *OTCollectorCountResponse {
	this := OTCollectorCountResponse{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *OTCollectorCountResponse) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *OTCollectorCountResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *OTCollectorCountResponse) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o OTCollectorCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCollectorCountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

type NullableOTCollectorCountResponse struct {
	value *OTCollectorCountResponse
	isSet bool
}

func (v NullableOTCollectorCountResponse) Get() *OTCollectorCountResponse {
	return v.value
}

func (v *NullableOTCollectorCountResponse) Set(val *OTCollectorCountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCollectorCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCollectorCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCollectorCountResponse(val *OTCollectorCountResponse) *NullableOTCollectorCountResponse {
	return &NullableOTCollectorCountResponse{value: val, isSet: true}
}

func (v NullableOTCollectorCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCollectorCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


