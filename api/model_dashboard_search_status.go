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

// checks if the DashboardSearchStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardSearchStatus{}

// DashboardSearchStatus struct for DashboardSearchStatus
type DashboardSearchStatus struct {
	// Current state of the search.
	State string `json:"state"`
	// Percentage of search completed.
	PercentCompleted *int32 `json:"percentCompleted,omitempty"`
}

// NewDashboardSearchStatus instantiates a new DashboardSearchStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSearchStatus(state string) *DashboardSearchStatus {
	this := DashboardSearchStatus{}
	this.State = state
	return &this
}

// NewDashboardSearchStatusWithDefaults instantiates a new DashboardSearchStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSearchStatusWithDefaults() *DashboardSearchStatus {
	this := DashboardSearchStatus{}
	return &this
}

// GetState returns the State field value
func (o *DashboardSearchStatus) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchStatus) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DashboardSearchStatus) SetState(v string) {
	o.State = v
}

// GetPercentCompleted returns the PercentCompleted field value if set, zero value otherwise.
func (o *DashboardSearchStatus) GetPercentCompleted() int32 {
	if o == nil || IsNil(o.PercentCompleted) {
		var ret int32
		return ret
	}
	return *o.PercentCompleted
}

// GetPercentCompletedOk returns a tuple with the PercentCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchStatus) GetPercentCompletedOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentCompleted) {
		return nil, false
	}
	return o.PercentCompleted, true
}

// HasPercentCompleted returns a boolean if a field has been set.
func (o *DashboardSearchStatus) HasPercentCompleted() bool {
	if o != nil && !IsNil(o.PercentCompleted) {
		return true
	}

	return false
}

// SetPercentCompleted gets a reference to the given int32 and assigns it to the PercentCompleted field.
func (o *DashboardSearchStatus) SetPercentCompleted(v int32) {
	o.PercentCompleted = &v
}

func (o DashboardSearchStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardSearchStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	if !IsNil(o.PercentCompleted) {
		toSerialize["percentCompleted"] = o.PercentCompleted
	}
	return toSerialize, nil
}

type NullableDashboardSearchStatus struct {
	value *DashboardSearchStatus
	isSet bool
}

func (v NullableDashboardSearchStatus) Get() *DashboardSearchStatus {
	return v.value
}

func (v *NullableDashboardSearchStatus) Set(val *DashboardSearchStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSearchStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSearchStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSearchStatus(val *DashboardSearchStatus) *NullableDashboardSearchStatus {
	return &NullableDashboardSearchStatus{value: val, isSet: true}
}

func (v NullableDashboardSearchStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSearchStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


