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

// checks if the BeginAsyncJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeginAsyncJobResponse{}

// BeginAsyncJobResponse struct for BeginAsyncJobResponse
type BeginAsyncJobResponse struct {
	// Identifier to get the status of an asynchronous job.
	Id string `json:"id"`
}

// NewBeginAsyncJobResponse instantiates a new BeginAsyncJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeginAsyncJobResponse(id string) *BeginAsyncJobResponse {
	this := BeginAsyncJobResponse{}
	this.Id = id
	return &this
}

// NewBeginAsyncJobResponseWithDefaults instantiates a new BeginAsyncJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeginAsyncJobResponseWithDefaults() *BeginAsyncJobResponse {
	this := BeginAsyncJobResponse{}
	return &this
}

// GetId returns the Id field value
func (o *BeginAsyncJobResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BeginAsyncJobResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BeginAsyncJobResponse) SetId(v string) {
	o.Id = v
}

func (o BeginAsyncJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeginAsyncJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableBeginAsyncJobResponse struct {
	value *BeginAsyncJobResponse
	isSet bool
}

func (v NullableBeginAsyncJobResponse) Get() *BeginAsyncJobResponse {
	return v.value
}

func (v *NullableBeginAsyncJobResponse) Set(val *BeginAsyncJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBeginAsyncJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBeginAsyncJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeginAsyncJobResponse(val *BeginAsyncJobResponse) *NullableBeginAsyncJobResponse {
	return &NullableBeginAsyncJobResponse{value: val, isSet: true}
}

func (v NullableBeginAsyncJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeginAsyncJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


