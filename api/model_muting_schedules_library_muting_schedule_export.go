/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MutingSchedulesLibraryMutingScheduleExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutingSchedulesLibraryMutingScheduleExport{}

// MutingSchedulesLibraryMutingScheduleExport struct for MutingSchedulesLibraryMutingScheduleExport
type MutingSchedulesLibraryMutingScheduleExport struct {
	MutingSchedulesLibraryBaseExport
	Schedule ScheduleDefinition `json:"schedule"`
	Monitor *MonitorScope `json:"monitor,omitempty"`
	NotificationGroups []GroupDefinition `json:"notificationGroups,omitempty"`
}

type _MutingSchedulesLibraryMutingScheduleExport MutingSchedulesLibraryMutingScheduleExport

// NewMutingSchedulesLibraryMutingScheduleExport instantiates a new MutingSchedulesLibraryMutingScheduleExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutingSchedulesLibraryMutingScheduleExport(schedule ScheduleDefinition, name string, type_ string) *MutingSchedulesLibraryMutingScheduleExport {
	this := MutingSchedulesLibraryMutingScheduleExport{}
	this.Name = name
	this.Type = type_
	this.Schedule = schedule
	return &this
}

// NewMutingSchedulesLibraryMutingScheduleExportWithDefaults instantiates a new MutingSchedulesLibraryMutingScheduleExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutingSchedulesLibraryMutingScheduleExportWithDefaults() *MutingSchedulesLibraryMutingScheduleExport {
	this := MutingSchedulesLibraryMutingScheduleExport{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *MutingSchedulesLibraryMutingScheduleExport) GetSchedule() ScheduleDefinition {
	if o == nil {
		var ret ScheduleDefinition
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryMutingScheduleExport) GetScheduleOk() (*ScheduleDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *MutingSchedulesLibraryMutingScheduleExport) SetSchedule(v ScheduleDefinition) {
	o.Schedule = v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *MutingSchedulesLibraryMutingScheduleExport) GetMonitor() MonitorScope {
	if o == nil || IsNil(o.Monitor) {
		var ret MonitorScope
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryMutingScheduleExport) GetMonitorOk() (*MonitorScope, bool) {
	if o == nil || IsNil(o.Monitor) {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *MutingSchedulesLibraryMutingScheduleExport) HasMonitor() bool {
	if o != nil && !IsNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given MonitorScope and assigns it to the Monitor field.
func (o *MutingSchedulesLibraryMutingScheduleExport) SetMonitor(v MonitorScope) {
	o.Monitor = &v
}

// GetNotificationGroups returns the NotificationGroups field value if set, zero value otherwise.
func (o *MutingSchedulesLibraryMutingScheduleExport) GetNotificationGroups() []GroupDefinition {
	if o == nil || IsNil(o.NotificationGroups) {
		var ret []GroupDefinition
		return ret
	}
	return o.NotificationGroups
}

// GetNotificationGroupsOk returns a tuple with the NotificationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingSchedulesLibraryMutingScheduleExport) GetNotificationGroupsOk() ([]GroupDefinition, bool) {
	if o == nil || IsNil(o.NotificationGroups) {
		return nil, false
	}
	return o.NotificationGroups, true
}

// HasNotificationGroups returns a boolean if a field has been set.
func (o *MutingSchedulesLibraryMutingScheduleExport) HasNotificationGroups() bool {
	if o != nil && !IsNil(o.NotificationGroups) {
		return true
	}

	return false
}

// SetNotificationGroups gets a reference to the given []GroupDefinition and assigns it to the NotificationGroups field.
func (o *MutingSchedulesLibraryMutingScheduleExport) SetNotificationGroups(v []GroupDefinition) {
	o.NotificationGroups = v
}

func (o MutingSchedulesLibraryMutingScheduleExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutingSchedulesLibraryMutingScheduleExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMutingSchedulesLibraryBaseExport, errMutingSchedulesLibraryBaseExport := json.Marshal(o.MutingSchedulesLibraryBaseExport)
	if errMutingSchedulesLibraryBaseExport != nil {
		return map[string]interface{}{}, errMutingSchedulesLibraryBaseExport
	}
	errMutingSchedulesLibraryBaseExport = json.Unmarshal([]byte(serializedMutingSchedulesLibraryBaseExport), &toSerialize)
	if errMutingSchedulesLibraryBaseExport != nil {
		return map[string]interface{}{}, errMutingSchedulesLibraryBaseExport
	}
	toSerialize["schedule"] = o.Schedule
	if !IsNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	if !IsNil(o.NotificationGroups) {
		toSerialize["notificationGroups"] = o.NotificationGroups
	}
	return toSerialize, nil
}

func (o *MutingSchedulesLibraryMutingScheduleExport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schedule",
		"name",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMutingSchedulesLibraryMutingScheduleExport := _MutingSchedulesLibraryMutingScheduleExport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMutingSchedulesLibraryMutingScheduleExport)

	if err != nil {
		return err
	}

	*o = MutingSchedulesLibraryMutingScheduleExport(varMutingSchedulesLibraryMutingScheduleExport)

	return err
}

type NullableMutingSchedulesLibraryMutingScheduleExport struct {
	value *MutingSchedulesLibraryMutingScheduleExport
	isSet bool
}

func (v NullableMutingSchedulesLibraryMutingScheduleExport) Get() *MutingSchedulesLibraryMutingScheduleExport {
	return v.value
}

func (v *NullableMutingSchedulesLibraryMutingScheduleExport) Set(val *MutingSchedulesLibraryMutingScheduleExport) {
	v.value = val
	v.isSet = true
}

func (v NullableMutingSchedulesLibraryMutingScheduleExport) IsSet() bool {
	return v.isSet
}

func (v *NullableMutingSchedulesLibraryMutingScheduleExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutingSchedulesLibraryMutingScheduleExport(val *MutingSchedulesLibraryMutingScheduleExport) *NullableMutingSchedulesLibraryMutingScheduleExport {
	return &NullableMutingSchedulesLibraryMutingScheduleExport{value: val, isSet: true}
}

func (v NullableMutingSchedulesLibraryMutingScheduleExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutingSchedulesLibraryMutingScheduleExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


