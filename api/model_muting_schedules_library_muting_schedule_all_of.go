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

// checks if the MutingSchedulesLibraryMutingScheduleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutingSchedulesLibraryMutingScheduleAllOf{}

// MutingSchedulesLibraryMutingScheduleAllOf struct for MutingSchedulesLibraryMutingScheduleAllOf
type MutingSchedulesLibraryMutingScheduleAllOf struct {
	Schedule ScheduleDefinition `json:"schedule"`
	Monitor *MonitorScope `json:"monitor,omitempty"`
	NotificationGroups []GroupDefinition `json:"notificationGroups,omitempty"`
}

// NewMutingSchedulesLibraryMutingScheduleAllOf instantiates a new MutingSchedulesLibraryMutingScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutingSchedulesLibraryMutingScheduleAllOf(schedule ScheduleDefinition) *MutingSchedulesLibraryMutingScheduleAllOf {
	this := MutingSchedulesLibraryMutingScheduleAllOf{}
	this.Schedule = schedule
	return &this
}

// NewMutingSchedulesLibraryMutingScheduleAllOfWithDefaults instantiates a new MutingSchedulesLibraryMutingScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutingSchedulesLibraryMutingScheduleAllOfWithDefaults() *MutingSchedulesLibraryMutingScheduleAllOf {
	this := MutingSchedulesLibraryMutingScheduleAllOf{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetSchedule() ScheduleDefinition {
	if o == nil {
		var ret ScheduleDefinition
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetScheduleOk() (*ScheduleDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *MutingSchedulesLibraryMutingScheduleAllOf) SetSchedule(v ScheduleDefinition) {
	o.Schedule = v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetMonitor() MonitorScope {
	if o == nil || IsNil(o.Monitor) {
		var ret MonitorScope
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetMonitorOk() (*MonitorScope, bool) {
	if o == nil || IsNil(o.Monitor) {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) HasMonitor() bool {
	if o != nil && !IsNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given MonitorScope and assigns it to the Monitor field.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) SetMonitor(v MonitorScope) {
	o.Monitor = &v
}

// GetNotificationGroups returns the NotificationGroups field value if set, zero value otherwise.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetNotificationGroups() []GroupDefinition {
	if o == nil || IsNil(o.NotificationGroups) {
		var ret []GroupDefinition
		return ret
	}
	return o.NotificationGroups
}

// GetNotificationGroupsOk returns a tuple with the NotificationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetNotificationGroupsOk() ([]GroupDefinition, bool) {
	if o == nil || IsNil(o.NotificationGroups) {
		return nil, false
	}
	return o.NotificationGroups, true
}

// HasNotificationGroups returns a boolean if a field has been set.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) HasNotificationGroups() bool {
	if o != nil && !IsNil(o.NotificationGroups) {
		return true
	}

	return false
}

// SetNotificationGroups gets a reference to the given []GroupDefinition and assigns it to the NotificationGroups field.
func (o *MutingSchedulesLibraryMutingScheduleAllOf) SetNotificationGroups(v []GroupDefinition) {
	o.NotificationGroups = v
}

func (o MutingSchedulesLibraryMutingScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutingSchedulesLibraryMutingScheduleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schedule"] = o.Schedule
	if !IsNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	if !IsNil(o.NotificationGroups) {
		toSerialize["notificationGroups"] = o.NotificationGroups
	}
	return toSerialize, nil
}

type NullableMutingSchedulesLibraryMutingScheduleAllOf struct {
	value *MutingSchedulesLibraryMutingScheduleAllOf
	isSet bool
}

func (v NullableMutingSchedulesLibraryMutingScheduleAllOf) Get() *MutingSchedulesLibraryMutingScheduleAllOf {
	return v.value
}

func (v *NullableMutingSchedulesLibraryMutingScheduleAllOf) Set(val *MutingSchedulesLibraryMutingScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMutingSchedulesLibraryMutingScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMutingSchedulesLibraryMutingScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutingSchedulesLibraryMutingScheduleAllOf(val *MutingSchedulesLibraryMutingScheduleAllOf) *NullableMutingSchedulesLibraryMutingScheduleAllOf {
	return &NullableMutingSchedulesLibraryMutingScheduleAllOf{value: val, isSet: true}
}

func (v NullableMutingSchedulesLibraryMutingScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutingSchedulesLibraryMutingScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


