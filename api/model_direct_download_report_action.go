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

// checks if the DirectDownloadReportAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectDownloadReportAction{}

// DirectDownloadReportAction struct for DirectDownloadReportAction
type DirectDownloadReportAction struct {
	ReportAction
}

// NewDirectDownloadReportAction instantiates a new DirectDownloadReportAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDownloadReportAction(actionType string) *DirectDownloadReportAction {
	this := DirectDownloadReportAction{}
	this.ActionType = actionType
	return &this
}

// NewDirectDownloadReportActionWithDefaults instantiates a new DirectDownloadReportAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDownloadReportActionWithDefaults() *DirectDownloadReportAction {
	this := DirectDownloadReportAction{}
	return &this
}

func (o DirectDownloadReportAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDownloadReportAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedReportAction, errReportAction := json.Marshal(o.ReportAction)
	if errReportAction != nil {
		return map[string]interface{}{}, errReportAction
	}
	errReportAction = json.Unmarshal([]byte(serializedReportAction), &toSerialize)
	if errReportAction != nil {
		return map[string]interface{}{}, errReportAction
	}
	return toSerialize, nil
}

type NullableDirectDownloadReportAction struct {
	value *DirectDownloadReportAction
	isSet bool
}

func (v NullableDirectDownloadReportAction) Get() *DirectDownloadReportAction {
	return v.value
}

func (v *NullableDirectDownloadReportAction) Set(val *DirectDownloadReportAction) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDownloadReportAction) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDownloadReportAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDownloadReportAction(val *DirectDownloadReportAction) *NullableDirectDownloadReportAction {
	return &NullableDirectDownloadReportAction{value: val, isSet: true}
}

func (v NullableDirectDownloadReportAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDownloadReportAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


