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

// checks if the CompliancePeriods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompliancePeriods{}

// CompliancePeriods Compliance periods along with SLO data availability progress.
type CompliancePeriods struct {
	// Time zone for the compliance periods as per the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
	// List of CompliancePeriodProgress.
	Periods []CompliancePeriodProgress `json:"periods"`
}

// NewCompliancePeriods instantiates a new CompliancePeriods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompliancePeriods(timezone string, periods []CompliancePeriodProgress) *CompliancePeriods {
	this := CompliancePeriods{}
	this.Timezone = timezone
	this.Periods = periods
	return &this
}

// NewCompliancePeriodsWithDefaults instantiates a new CompliancePeriods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompliancePeriodsWithDefaults() *CompliancePeriods {
	this := CompliancePeriods{}
	return &this
}

// GetTimezone returns the Timezone field value
func (o *CompliancePeriods) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *CompliancePeriods) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *CompliancePeriods) SetTimezone(v string) {
	o.Timezone = v
}

// GetPeriods returns the Periods field value
func (o *CompliancePeriods) GetPeriods() []CompliancePeriodProgress {
	if o == nil {
		var ret []CompliancePeriodProgress
		return ret
	}

	return o.Periods
}

// GetPeriodsOk returns a tuple with the Periods field value
// and a boolean to check if the value has been set.
func (o *CompliancePeriods) GetPeriodsOk() ([]CompliancePeriodProgress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Periods, true
}

// SetPeriods sets field value
func (o *CompliancePeriods) SetPeriods(v []CompliancePeriodProgress) {
	o.Periods = v
}

func (o CompliancePeriods) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompliancePeriods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timezone"] = o.Timezone
	toSerialize["periods"] = o.Periods
	return toSerialize, nil
}

type NullableCompliancePeriods struct {
	value *CompliancePeriods
	isSet bool
}

func (v NullableCompliancePeriods) Get() *CompliancePeriods {
	return v.value
}

func (v *NullableCompliancePeriods) Set(val *CompliancePeriods) {
	v.value = val
	v.isSet = true
}

func (v NullableCompliancePeriods) IsSet() bool {
	return v.isSet
}

func (v *NullableCompliancePeriods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompliancePeriods(val *CompliancePeriods) *NullableCompliancePeriods {
	return &NullableCompliancePeriods{value: val, isSet: true}
}

func (v NullableCompliancePeriods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompliancePeriods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


