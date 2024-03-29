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

// checks if the MutingSchedulesLibraryMutingSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutingSchedulesLibraryMutingSchedule{}

// MutingSchedulesLibraryMutingSchedule struct for MutingSchedulesLibraryMutingSchedule
type MutingSchedulesLibraryMutingSchedule struct {
	MutingSchedulesLibraryBase
	Schedule ScheduleDefinition `json:"schedule"`
	Monitor *MonitorScope `json:"monitor,omitempty"`
}

// NewMutingSchedulesLibraryMutingSchedule instantiates a new MutingSchedulesLibraryMutingSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutingSchedulesLibraryMutingSchedule(schedule ScheduleDefinition, name string, type_ string) *MutingSchedulesLibraryMutingSchedule {
	this := MutingSchedulesLibraryMutingSchedule{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Type = type_
	this.Schedule = schedule
	return &this
}

// NewMutingSchedulesLibraryMutingScheduleWithDefaults instantiates a new MutingSchedulesLibraryMutingSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutingSchedulesLibraryMutingScheduleWithDefaults() *MutingSchedulesLibraryMutingSchedule {
	this := MutingSchedulesLibraryMutingSchedule{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *MutingSchedulesLibraryMutingSchedule) GetSchedule() ScheduleDefinition {
	if o == nil {
		var ret ScheduleDefinition
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryMutingSchedule) GetScheduleOk() (*ScheduleDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *MutingSchedulesLibraryMutingSchedule) SetSchedule(v ScheduleDefinition) {
	o.Schedule = v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *MutingSchedulesLibraryMutingSchedule) GetMonitor() MonitorScope {
	if o == nil || IsNil(o.Monitor) {
		var ret MonitorScope
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryMutingSchedule) GetMonitorOk() (*MonitorScope, bool) {
	if o == nil || IsNil(o.Monitor) {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *MutingSchedulesLibraryMutingSchedule) HasMonitor() bool {
	if o != nil && !IsNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given MonitorScope and assigns it to the Monitor field.
func (o *MutingSchedulesLibraryMutingSchedule) SetMonitor(v MonitorScope) {
	o.Monitor = &v
}

func (o MutingSchedulesLibraryMutingSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutingSchedulesLibraryMutingSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMutingSchedulesLibraryBase, errMutingSchedulesLibraryBase := json.Marshal(o.MutingSchedulesLibraryBase)
	if errMutingSchedulesLibraryBase != nil {
		return map[string]interface{}{}, errMutingSchedulesLibraryBase
	}
	errMutingSchedulesLibraryBase = json.Unmarshal([]byte(serializedMutingSchedulesLibraryBase), &toSerialize)
	if errMutingSchedulesLibraryBase != nil {
		return map[string]interface{}{}, errMutingSchedulesLibraryBase
	}
	toSerialize["schedule"] = o.Schedule
	if !IsNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	return toSerialize, nil
}

type NullableMutingSchedulesLibraryMutingSchedule struct {
	value *MutingSchedulesLibraryMutingSchedule
	isSet bool
}

func (v NullableMutingSchedulesLibraryMutingSchedule) Get() *MutingSchedulesLibraryMutingSchedule {
	return v.value
}

func (v *NullableMutingSchedulesLibraryMutingSchedule) Set(val *MutingSchedulesLibraryMutingSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableMutingSchedulesLibraryMutingSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableMutingSchedulesLibraryMutingSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutingSchedulesLibraryMutingSchedule(val *MutingSchedulesLibraryMutingSchedule) *NullableMutingSchedulesLibraryMutingSchedule {
	return &NullableMutingSchedulesLibraryMutingSchedule{value: val, isSet: true}
}

func (v NullableMutingSchedulesLibraryMutingSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutingSchedulesLibraryMutingSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


