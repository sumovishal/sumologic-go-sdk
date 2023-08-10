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

// checks if the TraceExistsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceExistsResponse{}

// TraceExistsResponse struct for TraceExistsResponse
type TraceExistsResponse struct {
	// Indicates whether the trace with the given trace id exists.
	Exists bool `json:"exists"`
	// A path to the trace view page in Sumo Logic UI.
	Url *string `json:"url,omitempty"`
}

// NewTraceExistsResponse instantiates a new TraceExistsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceExistsResponse(exists bool) *TraceExistsResponse {
	this := TraceExistsResponse{}
	this.Exists = exists
	return &this
}

// NewTraceExistsResponseWithDefaults instantiates a new TraceExistsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceExistsResponseWithDefaults() *TraceExistsResponse {
	this := TraceExistsResponse{}
	return &this
}

// GetExists returns the Exists field value
func (o *TraceExistsResponse) GetExists() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exists
}

// GetExistsOk returns a tuple with the Exists field value
// and a boolean to check if the value has been set.
func (o *TraceExistsResponse) GetExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exists, true
}

// SetExists sets field value
func (o *TraceExistsResponse) SetExists(v bool) {
	o.Exists = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TraceExistsResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceExistsResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TraceExistsResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TraceExistsResponse) SetUrl(v string) {
	o.Url = &v
}

func (o TraceExistsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceExistsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exists"] = o.Exists
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableTraceExistsResponse struct {
	value *TraceExistsResponse
	isSet bool
}

func (v NullableTraceExistsResponse) Get() *TraceExistsResponse {
	return v.value
}

func (v *NullableTraceExistsResponse) Set(val *TraceExistsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceExistsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceExistsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceExistsResponse(val *TraceExistsResponse) *NullableTraceExistsResponse {
	return &NullableTraceExistsResponse{value: val, isSet: true}
}

func (v NullableTraceExistsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceExistsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


