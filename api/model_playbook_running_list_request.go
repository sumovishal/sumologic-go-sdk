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

// checks if the PlaybookRunningListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaybookRunningListRequest{}

// PlaybookRunningListRequest The request parameters for getting all running playbooks' status.
type PlaybookRunningListRequest struct {
	// The alert id.
	AlertId string `json:"alertId"`
}

// NewPlaybookRunningListRequest instantiates a new PlaybookRunningListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaybookRunningListRequest(alertId string) *PlaybookRunningListRequest {
	this := PlaybookRunningListRequest{}
	this.AlertId = alertId
	return &this
}

// NewPlaybookRunningListRequestWithDefaults instantiates a new PlaybookRunningListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaybookRunningListRequestWithDefaults() *PlaybookRunningListRequest {
	this := PlaybookRunningListRequest{}
	return &this
}

// GetAlertId returns the AlertId field value
func (o *PlaybookRunningListRequest) GetAlertId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value
// and a boolean to check if the value has been set.
func (o *PlaybookRunningListRequest) GetAlertIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertId, true
}

// SetAlertId sets field value
func (o *PlaybookRunningListRequest) SetAlertId(v string) {
	o.AlertId = v
}

func (o PlaybookRunningListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaybookRunningListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alertId"] = o.AlertId
	return toSerialize, nil
}

type NullablePlaybookRunningListRequest struct {
	value *PlaybookRunningListRequest
	isSet bool
}

func (v NullablePlaybookRunningListRequest) Get() *PlaybookRunningListRequest {
	return v.value
}

func (v *NullablePlaybookRunningListRequest) Set(val *PlaybookRunningListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybookRunningListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybookRunningListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybookRunningListRequest(val *PlaybookRunningListRequest) *NullablePlaybookRunningListRequest {
	return &NullablePlaybookRunningListRequest{value: val, isSet: true}
}

func (v NullablePlaybookRunningListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybookRunningListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


