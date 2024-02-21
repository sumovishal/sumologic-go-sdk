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

// checks if the ReportAutoParsingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportAutoParsingInfo{}

// ReportAutoParsingInfo Auto-parsing information for the panel. This information tells us whether automatic field extraction from JSON log messages is enabled or not
type ReportAutoParsingInfo struct {
	// Can be `intelligent` or `performance`
	Mode *string `json:"mode,omitempty"`
}

// NewReportAutoParsingInfo instantiates a new ReportAutoParsingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportAutoParsingInfo() *ReportAutoParsingInfo {
	this := ReportAutoParsingInfo{}
	var mode string = "performance"
	this.Mode = &mode
	return &this
}

// NewReportAutoParsingInfoWithDefaults instantiates a new ReportAutoParsingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportAutoParsingInfoWithDefaults() *ReportAutoParsingInfo {
	this := ReportAutoParsingInfo{}
	var mode string = "performance"
	this.Mode = &mode
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ReportAutoParsingInfo) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportAutoParsingInfo) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ReportAutoParsingInfo) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ReportAutoParsingInfo) SetMode(v string) {
	o.Mode = &v
}

func (o ReportAutoParsingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportAutoParsingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableReportAutoParsingInfo struct {
	value *ReportAutoParsingInfo
	isSet bool
}

func (v NullableReportAutoParsingInfo) Get() *ReportAutoParsingInfo {
	return v.value
}

func (v *NullableReportAutoParsingInfo) Set(val *ReportAutoParsingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReportAutoParsingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReportAutoParsingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportAutoParsingInfo(val *ReportAutoParsingInfo) *NullableReportAutoParsingInfo {
	return &NullableReportAutoParsingInfo{value: val, isSet: true}
}

func (v NullableReportAutoParsingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportAutoParsingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


