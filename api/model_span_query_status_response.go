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

// checks if the SpanQueryStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanQueryStatusResponse{}

// SpanQueryStatusResponse struct for SpanQueryStatusResponse
type SpanQueryStatusResponse struct {
	// A list of span analytics queries.
	QueryRows []SpanQueryRowStatus `json:"queryRows"`
	// Status of the query. Possible values: `Processing`, `Finished`, `Error`, `Paused`
	Status string `json:"status"`
}

// NewSpanQueryStatusResponse instantiates a new SpanQueryStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanQueryStatusResponse(queryRows []SpanQueryRowStatus, status string) *SpanQueryStatusResponse {
	this := SpanQueryStatusResponse{}
	this.QueryRows = queryRows
	this.Status = status
	return &this
}

// NewSpanQueryStatusResponseWithDefaults instantiates a new SpanQueryStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanQueryStatusResponseWithDefaults() *SpanQueryStatusResponse {
	this := SpanQueryStatusResponse{}
	return &this
}

// GetQueryRows returns the QueryRows field value
func (o *SpanQueryStatusResponse) GetQueryRows() []SpanQueryRowStatus {
	if o == nil {
		var ret []SpanQueryRowStatus
		return ret
	}

	return o.QueryRows
}

// GetQueryRowsOk returns a tuple with the QueryRows field value
// and a boolean to check if the value has been set.
func (o *SpanQueryStatusResponse) GetQueryRowsOk() ([]SpanQueryRowStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryRows, true
}

// SetQueryRows sets field value
func (o *SpanQueryStatusResponse) SetQueryRows(v []SpanQueryRowStatus) {
	o.QueryRows = v
}

// GetStatus returns the Status field value
func (o *SpanQueryStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SpanQueryStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SpanQueryStatusResponse) SetStatus(v string) {
	o.Status = v
}

func (o SpanQueryStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanQueryStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryRows"] = o.QueryRows
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableSpanQueryStatusResponse struct {
	value *SpanQueryStatusResponse
	isSet bool
}

func (v NullableSpanQueryStatusResponse) Get() *SpanQueryStatusResponse {
	return v.value
}

func (v *NullableSpanQueryStatusResponse) Set(val *SpanQueryStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanQueryStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanQueryStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanQueryStatusResponse(val *SpanQueryStatusResponse) *NullableSpanQueryStatusResponse {
	return &NullableSpanQueryStatusResponse{value: val, isSet: true}
}

func (v NullableSpanQueryStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanQueryStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


