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

// checks if the CseSignalNotificationSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CseSignalNotificationSyncDefinition{}

// CseSignalNotificationSyncDefinition struct for CseSignalNotificationSyncDefinition
type CseSignalNotificationSyncDefinition struct {
	ScheduleNotificationSyncDefinition
	// Name of the Cloud SIEM Enterprise Record to be created.
	RecordType string `json:"recordType"`
}

// NewCseSignalNotificationSyncDefinition instantiates a new CseSignalNotificationSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCseSignalNotificationSyncDefinition(recordType string, taskType string) *CseSignalNotificationSyncDefinition {
	this := CseSignalNotificationSyncDefinition{}
	this.TaskType = taskType
	this.RecordType = recordType
	return &this
}

// NewCseSignalNotificationSyncDefinitionWithDefaults instantiates a new CseSignalNotificationSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCseSignalNotificationSyncDefinitionWithDefaults() *CseSignalNotificationSyncDefinition {
	this := CseSignalNotificationSyncDefinition{}
	return &this
}

// GetRecordType returns the RecordType field value
func (o *CseSignalNotificationSyncDefinition) GetRecordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value
// and a boolean to check if the value has been set.
func (o *CseSignalNotificationSyncDefinition) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordType, true
}

// SetRecordType sets field value
func (o *CseSignalNotificationSyncDefinition) SetRecordType(v string) {
	o.RecordType = v
}

func (o CseSignalNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CseSignalNotificationSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedScheduleNotificationSyncDefinition, errScheduleNotificationSyncDefinition := json.Marshal(o.ScheduleNotificationSyncDefinition)
	if errScheduleNotificationSyncDefinition != nil {
		return map[string]interface{}{}, errScheduleNotificationSyncDefinition
	}
	errScheduleNotificationSyncDefinition = json.Unmarshal([]byte(serializedScheduleNotificationSyncDefinition), &toSerialize)
	if errScheduleNotificationSyncDefinition != nil {
		return map[string]interface{}{}, errScheduleNotificationSyncDefinition
	}
	toSerialize["recordType"] = o.RecordType
	return toSerialize, nil
}

type NullableCseSignalNotificationSyncDefinition struct {
	value *CseSignalNotificationSyncDefinition
	isSet bool
}

func (v NullableCseSignalNotificationSyncDefinition) Get() *CseSignalNotificationSyncDefinition {
	return v.value
}

func (v *NullableCseSignalNotificationSyncDefinition) Set(val *CseSignalNotificationSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCseSignalNotificationSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCseSignalNotificationSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCseSignalNotificationSyncDefinition(val *CseSignalNotificationSyncDefinition) *NullableCseSignalNotificationSyncDefinition {
	return &NullableCseSignalNotificationSyncDefinition{value: val, isSet: true}
}

func (v NullableCseSignalNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCseSignalNotificationSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


