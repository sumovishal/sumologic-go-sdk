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

// checks if the SaveToViewNotificationSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveToViewNotificationSyncDefinition{}

// SaveToViewNotificationSyncDefinition struct for SaveToViewNotificationSyncDefinition
type SaveToViewNotificationSyncDefinition struct {
	ScheduleNotificationSyncDefinition
	// Name of the View to save the notification to.
	ViewName string `json:"viewName"`
}

// NewSaveToViewNotificationSyncDefinition instantiates a new SaveToViewNotificationSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveToViewNotificationSyncDefinition(viewName string, taskType string) *SaveToViewNotificationSyncDefinition {
	this := SaveToViewNotificationSyncDefinition{}
	this.TaskType = taskType
	this.ViewName = viewName
	return &this
}

// NewSaveToViewNotificationSyncDefinitionWithDefaults instantiates a new SaveToViewNotificationSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveToViewNotificationSyncDefinitionWithDefaults() *SaveToViewNotificationSyncDefinition {
	this := SaveToViewNotificationSyncDefinition{}
	return &this
}

// GetViewName returns the ViewName field value
func (o *SaveToViewNotificationSyncDefinition) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *SaveToViewNotificationSyncDefinition) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value
func (o *SaveToViewNotificationSyncDefinition) SetViewName(v string) {
	o.ViewName = v
}

func (o SaveToViewNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveToViewNotificationSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedScheduleNotificationSyncDefinition, errScheduleNotificationSyncDefinition := json.Marshal(o.ScheduleNotificationSyncDefinition)
	if errScheduleNotificationSyncDefinition != nil {
		return map[string]interface{}{}, errScheduleNotificationSyncDefinition
	}
	errScheduleNotificationSyncDefinition = json.Unmarshal([]byte(serializedScheduleNotificationSyncDefinition), &toSerialize)
	if errScheduleNotificationSyncDefinition != nil {
		return map[string]interface{}{}, errScheduleNotificationSyncDefinition
	}
	toSerialize["viewName"] = o.ViewName
	return toSerialize, nil
}

type NullableSaveToViewNotificationSyncDefinition struct {
	value *SaveToViewNotificationSyncDefinition
	isSet bool
}

func (v NullableSaveToViewNotificationSyncDefinition) Get() *SaveToViewNotificationSyncDefinition {
	return v.value
}

func (v *NullableSaveToViewNotificationSyncDefinition) Set(val *SaveToViewNotificationSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveToViewNotificationSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveToViewNotificationSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveToViewNotificationSyncDefinition(val *SaveToViewNotificationSyncDefinition) *NullableSaveToViewNotificationSyncDefinition {
	return &NullableSaveToViewNotificationSyncDefinition{value: val, isSet: true}
}

func (v NullableSaveToViewNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveToViewNotificationSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


