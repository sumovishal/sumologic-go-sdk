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

// checks if the NextInstancesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NextInstancesRequest{}

// NextInstancesRequest struct for NextInstancesRequest
type NextInstancesRequest struct {
	// Time zone for the schedule per [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
	// Schedule start date in the format of `yyyy-mm-dd`
	StartDate string `json:"startDate"`
	// Schedule start time in the format of `hh:mm`
	StartTime string `json:"startTime"`
	// RRule (Recurrence Rule)
	Rrule string `json:"rrule"`
}

// NewNextInstancesRequest instantiates a new NextInstancesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextInstancesRequest(timezone string, startDate string, startTime string, rrule string) *NextInstancesRequest {
	this := NextInstancesRequest{}
	this.Timezone = timezone
	this.StartDate = startDate
	this.StartTime = startTime
	this.Rrule = rrule
	return &this
}

// NewNextInstancesRequestWithDefaults instantiates a new NextInstancesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextInstancesRequestWithDefaults() *NextInstancesRequest {
	this := NextInstancesRequest{}
	return &this
}

// GetTimezone returns the Timezone field value
func (o *NextInstancesRequest) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *NextInstancesRequest) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *NextInstancesRequest) SetTimezone(v string) {
	o.Timezone = v
}

// GetStartDate returns the StartDate field value
func (o *NextInstancesRequest) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *NextInstancesRequest) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *NextInstancesRequest) SetStartDate(v string) {
	o.StartDate = v
}

// GetStartTime returns the StartTime field value
func (o *NextInstancesRequest) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *NextInstancesRequest) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *NextInstancesRequest) SetStartTime(v string) {
	o.StartTime = v
}

// GetRrule returns the Rrule field value
func (o *NextInstancesRequest) GetRrule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *NextInstancesRequest) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rrule, true
}

// SetRrule sets field value
func (o *NextInstancesRequest) SetRrule(v string) {
	o.Rrule = v
}

func (o NextInstancesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NextInstancesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timezone"] = o.Timezone
	toSerialize["startDate"] = o.StartDate
	toSerialize["startTime"] = o.StartTime
	toSerialize["rrule"] = o.Rrule
	return toSerialize, nil
}

type NullableNextInstancesRequest struct {
	value *NextInstancesRequest
	isSet bool
}

func (v NullableNextInstancesRequest) Get() *NextInstancesRequest {
	return v.value
}

func (v *NullableNextInstancesRequest) Set(val *NextInstancesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNextInstancesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNextInstancesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextInstancesRequest(val *NextInstancesRequest) *NullableNextInstancesRequest {
	return &NullableNextInstancesRequest{value: val, isSet: true}
}

func (v NullableNextInstancesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextInstancesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


