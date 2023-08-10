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

// checks if the CompliancePeriodRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompliancePeriodRef{}

// CompliancePeriodRef Reference to the compliance period of the SLO.
type CompliancePeriodRef struct {
	// Type of reference to the compliance period. Must be `Relative`.
	ComplianceRefType string `json:"complianceRefType"`
	// Relative shift of compliance period from the latest/current compliance period.
	RelativeShift *int32 `json:"relativeShift,omitempty"`
}

// NewCompliancePeriodRef instantiates a new CompliancePeriodRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompliancePeriodRef(complianceRefType string) *CompliancePeriodRef {
	this := CompliancePeriodRef{}
	this.ComplianceRefType = complianceRefType
	return &this
}

// NewCompliancePeriodRefWithDefaults instantiates a new CompliancePeriodRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompliancePeriodRefWithDefaults() *CompliancePeriodRef {
	this := CompliancePeriodRef{}
	return &this
}

// GetComplianceRefType returns the ComplianceRefType field value
func (o *CompliancePeriodRef) GetComplianceRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceRefType
}

// GetComplianceRefTypeOk returns a tuple with the ComplianceRefType field value
// and a boolean to check if the value has been set.
func (o *CompliancePeriodRef) GetComplianceRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceRefType, true
}

// SetComplianceRefType sets field value
func (o *CompliancePeriodRef) SetComplianceRefType(v string) {
	o.ComplianceRefType = v
}

// GetRelativeShift returns the RelativeShift field value if set, zero value otherwise.
func (o *CompliancePeriodRef) GetRelativeShift() int32 {
	if o == nil || IsNil(o.RelativeShift) {
		var ret int32
		return ret
	}
	return *o.RelativeShift
}

// GetRelativeShiftOk returns a tuple with the RelativeShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompliancePeriodRef) GetRelativeShiftOk() (*int32, bool) {
	if o == nil || IsNil(o.RelativeShift) {
		return nil, false
	}
	return o.RelativeShift, true
}

// HasRelativeShift returns a boolean if a field has been set.
func (o *CompliancePeriodRef) HasRelativeShift() bool {
	if o != nil && !IsNil(o.RelativeShift) {
		return true
	}

	return false
}

// SetRelativeShift gets a reference to the given int32 and assigns it to the RelativeShift field.
func (o *CompliancePeriodRef) SetRelativeShift(v int32) {
	o.RelativeShift = &v
}

func (o CompliancePeriodRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompliancePeriodRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["complianceRefType"] = o.ComplianceRefType
	if !IsNil(o.RelativeShift) {
		toSerialize["relativeShift"] = o.RelativeShift
	}
	return toSerialize, nil
}

type NullableCompliancePeriodRef struct {
	value *CompliancePeriodRef
	isSet bool
}

func (v NullableCompliancePeriodRef) Get() *CompliancePeriodRef {
	return v.value
}

func (v *NullableCompliancePeriodRef) Set(val *CompliancePeriodRef) {
	v.value = val
	v.isSet = true
}

func (v NullableCompliancePeriodRef) IsSet() bool {
	return v.isSet
}

func (v *NullableCompliancePeriodRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompliancePeriodRef(val *CompliancePeriodRef) *NullableCompliancePeriodRef {
	return &NullableCompliancePeriodRef{value: val, isSet: true}
}

func (v NullableCompliancePeriodRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompliancePeriodRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


