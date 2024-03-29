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

// checks if the DashboardMigrationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardMigrationStatus{}

// DashboardMigrationStatus struct for DashboardMigrationStatus
type DashboardMigrationStatus struct {
	// A successful migration to Dashboard(New).
	SuccessCount int32 `json:"successCount"`
	// A failed migration to Dashboard(New).
	FailedCount int32 `json:"failedCount"`
	// The total number of Legacy Dashboards to migrate.
	TotalCount int32 `json:"totalCount"`
}

// NewDashboardMigrationStatus instantiates a new DashboardMigrationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardMigrationStatus(successCount int32, failedCount int32, totalCount int32) *DashboardMigrationStatus {
	this := DashboardMigrationStatus{}
	this.SuccessCount = successCount
	this.FailedCount = failedCount
	this.TotalCount = totalCount
	return &this
}

// NewDashboardMigrationStatusWithDefaults instantiates a new DashboardMigrationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardMigrationStatusWithDefaults() *DashboardMigrationStatus {
	this := DashboardMigrationStatus{}
	return &this
}

// GetSuccessCount returns the SuccessCount field value
func (o *DashboardMigrationStatus) GetSuccessCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value
// and a boolean to check if the value has been set.
func (o *DashboardMigrationStatus) GetSuccessCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessCount, true
}

// SetSuccessCount sets field value
func (o *DashboardMigrationStatus) SetSuccessCount(v int32) {
	o.SuccessCount = v
}

// GetFailedCount returns the FailedCount field value
func (o *DashboardMigrationStatus) GetFailedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailedCount
}

// GetFailedCountOk returns a tuple with the FailedCount field value
// and a boolean to check if the value has been set.
func (o *DashboardMigrationStatus) GetFailedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedCount, true
}

// SetFailedCount sets field value
func (o *DashboardMigrationStatus) SetFailedCount(v int32) {
	o.FailedCount = v
}

// GetTotalCount returns the TotalCount field value
func (o *DashboardMigrationStatus) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *DashboardMigrationStatus) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *DashboardMigrationStatus) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o DashboardMigrationStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardMigrationStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["successCount"] = o.SuccessCount
	toSerialize["failedCount"] = o.FailedCount
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

type NullableDashboardMigrationStatus struct {
	value *DashboardMigrationStatus
	isSet bool
}

func (v NullableDashboardMigrationStatus) Get() *DashboardMigrationStatus {
	return v.value
}

func (v *NullableDashboardMigrationStatus) Set(val *DashboardMigrationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardMigrationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardMigrationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardMigrationStatus(val *DashboardMigrationStatus) *NullableDashboardMigrationStatus {
	return &NullableDashboardMigrationStatus{value: val, isSet: true}
}

func (v NullableDashboardMigrationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardMigrationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


