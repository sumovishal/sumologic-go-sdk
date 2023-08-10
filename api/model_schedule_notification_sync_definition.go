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

// checks if the ScheduleNotificationSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleNotificationSyncDefinition{}

// ScheduleNotificationSyncDefinition struct for ScheduleNotificationSyncDefinition
type ScheduleNotificationSyncDefinition struct {
	// Delivery channel for notifications.
	TaskType string `json:"taskType"`
}

// NewScheduleNotificationSyncDefinition instantiates a new ScheduleNotificationSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleNotificationSyncDefinition(taskType string) *ScheduleNotificationSyncDefinition {
	this := ScheduleNotificationSyncDefinition{}
	this.TaskType = taskType
	return &this
}

// NewScheduleNotificationSyncDefinitionWithDefaults instantiates a new ScheduleNotificationSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleNotificationSyncDefinitionWithDefaults() *ScheduleNotificationSyncDefinition {
	this := ScheduleNotificationSyncDefinition{}
	return &this
}

// GetTaskType returns the TaskType field value
func (o *ScheduleNotificationSyncDefinition) GetTaskType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *ScheduleNotificationSyncDefinition) GetTaskTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value
func (o *ScheduleNotificationSyncDefinition) SetTaskType(v string) {
	o.TaskType = v
}

func (o ScheduleNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleNotificationSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["taskType"] = o.TaskType
	return toSerialize, nil
}

type NullableScheduleNotificationSyncDefinition struct {
	value *ScheduleNotificationSyncDefinition
	isSet bool
}

func (v NullableScheduleNotificationSyncDefinition) Get() *ScheduleNotificationSyncDefinition {
	return v.value
}

func (v *NullableScheduleNotificationSyncDefinition) Set(val *ScheduleNotificationSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleNotificationSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleNotificationSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleNotificationSyncDefinition(val *ScheduleNotificationSyncDefinition) *NullableScheduleNotificationSyncDefinition {
	return &NullableScheduleNotificationSyncDefinition{value: val, isSet: true}
}

func (v NullableScheduleNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleNotificationSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


