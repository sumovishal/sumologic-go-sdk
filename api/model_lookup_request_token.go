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

// checks if the LookupRequestToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupRequestToken{}

// LookupRequestToken Allows you to track the status of an upload or export request.
type LookupRequestToken struct {
	// The identifier used to track the request.
	Id string `json:"id"`
}

// NewLookupRequestToken instantiates a new LookupRequestToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupRequestToken(id string) *LookupRequestToken {
	this := LookupRequestToken{}
	this.Id = id
	return &this
}

// NewLookupRequestTokenWithDefaults instantiates a new LookupRequestToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupRequestTokenWithDefaults() *LookupRequestToken {
	this := LookupRequestToken{}
	return &this
}

// GetId returns the Id field value
func (o *LookupRequestToken) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LookupRequestToken) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LookupRequestToken) SetId(v string) {
	o.Id = v
}

func (o LookupRequestToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupRequestToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableLookupRequestToken struct {
	value *LookupRequestToken
	isSet bool
}

func (v NullableLookupRequestToken) Get() *LookupRequestToken {
	return v.value
}

func (v *NullableLookupRequestToken) Set(val *LookupRequestToken) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupRequestToken) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupRequestToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupRequestToken(val *LookupRequestToken) *NullableLookupRequestToken {
	return &NullableLookupRequestToken{value: val, isSet: true}
}

func (v NullableLookupRequestToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupRequestToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


