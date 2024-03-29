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

// checks if the LogSearchEstimatedUsageRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogSearchEstimatedUsageRequestAllOf{}

// LogSearchEstimatedUsageRequestAllOf struct for LogSearchEstimatedUsageRequestAllOf
type LogSearchEstimatedUsageRequestAllOf struct {
	// Time zone to get the estimated usage details. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). 
	Timezone string `json:"timezone"`
}

// NewLogSearchEstimatedUsageRequestAllOf instantiates a new LogSearchEstimatedUsageRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchEstimatedUsageRequestAllOf(timezone string) *LogSearchEstimatedUsageRequestAllOf {
	this := LogSearchEstimatedUsageRequestAllOf{}
	this.Timezone = timezone
	return &this
}

// NewLogSearchEstimatedUsageRequestAllOfWithDefaults instantiates a new LogSearchEstimatedUsageRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchEstimatedUsageRequestAllOfWithDefaults() *LogSearchEstimatedUsageRequestAllOf {
	this := LogSearchEstimatedUsageRequestAllOf{}
	return &this
}

// GetTimezone returns the Timezone field value
func (o *LogSearchEstimatedUsageRequestAllOf) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *LogSearchEstimatedUsageRequestAllOf) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *LogSearchEstimatedUsageRequestAllOf) SetTimezone(v string) {
	o.Timezone = v
}

func (o LogSearchEstimatedUsageRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogSearchEstimatedUsageRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timezone"] = o.Timezone
	return toSerialize, nil
}

type NullableLogSearchEstimatedUsageRequestAllOf struct {
	value *LogSearchEstimatedUsageRequestAllOf
	isSet bool
}

func (v NullableLogSearchEstimatedUsageRequestAllOf) Get() *LogSearchEstimatedUsageRequestAllOf {
	return v.value
}

func (v *NullableLogSearchEstimatedUsageRequestAllOf) Set(val *LogSearchEstimatedUsageRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchEstimatedUsageRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchEstimatedUsageRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchEstimatedUsageRequestAllOf(val *LogSearchEstimatedUsageRequestAllOf) *NullableLogSearchEstimatedUsageRequestAllOf {
	return &NullableLogSearchEstimatedUsageRequestAllOf{value: val, isSet: true}
}

func (v NullableLogSearchEstimatedUsageRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchEstimatedUsageRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


