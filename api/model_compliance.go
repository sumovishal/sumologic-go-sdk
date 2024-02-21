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

// checks if the Compliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Compliance{}

// Compliance struct for Compliance
type Compliance struct {
	// Compliance Type (rolling or calendar)
	ComplianceType string `json:"complianceType"`
	// Target percentage for the SLI over the compliance period.
	Target float32 `json:"target"`
	// Time zone for the SLO compliance. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
}

// NewCompliance instantiates a new Compliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompliance(complianceType string, target float32, timezone string) *Compliance {
	this := Compliance{}
	this.ComplianceType = complianceType
	this.Target = target
	this.Timezone = timezone
	return &this
}

// NewComplianceWithDefaults instantiates a new Compliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplianceWithDefaults() *Compliance {
	this := Compliance{}
	return &this
}

// GetComplianceType returns the ComplianceType field value
func (o *Compliance) GetComplianceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceType
}

// GetComplianceTypeOk returns a tuple with the ComplianceType field value
// and a boolean to check if the value has been set.
func (o *Compliance) GetComplianceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceType, true
}

// SetComplianceType sets field value
func (o *Compliance) SetComplianceType(v string) {
	o.ComplianceType = v
}

// GetTarget returns the Target field value
func (o *Compliance) GetTarget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Compliance) GetTargetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Compliance) SetTarget(v float32) {
	o.Target = v
}

// GetTimezone returns the Timezone field value
func (o *Compliance) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *Compliance) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *Compliance) SetTimezone(v string) {
	o.Timezone = v
}

func (o Compliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Compliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["complianceType"] = o.ComplianceType
	toSerialize["target"] = o.Target
	toSerialize["timezone"] = o.Timezone
	return toSerialize, nil
}

type NullableCompliance struct {
	value *Compliance
	isSet bool
}

func (v NullableCompliance) Get() *Compliance {
	return v.value
}

func (v *NullableCompliance) Set(val *Compliance) {
	v.value = val
	v.isSet = true
}

func (v NullableCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompliance(val *Compliance) *NullableCompliance {
	return &NullableCompliance{value: val, isSet: true}
}

func (v NullableCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


