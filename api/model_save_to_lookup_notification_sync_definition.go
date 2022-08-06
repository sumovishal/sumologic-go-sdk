/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// SaveToLookupNotificationSyncDefinition struct for SaveToLookupNotificationSyncDefinition
type SaveToLookupNotificationSyncDefinition struct {
	ScheduleNotificationSyncDefinition
	// The path of the lookup table that will store the results of the scheduled search.
	LookupFilePath string `json:"lookupFilePath"`
	// This indicates whether the file contents will be merged with existing data in the lookup table or not. If this is true then data with the same primary keys will be updated while the rest of the rows will be appended.
	IsLookupMergeOperation bool `json:"isLookupMergeOperation"`
}

// NewSaveToLookupNotificationSyncDefinition instantiates a new SaveToLookupNotificationSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveToLookupNotificationSyncDefinition(lookupFilePath string, isLookupMergeOperation bool, taskType string) *SaveToLookupNotificationSyncDefinition {
	this := SaveToLookupNotificationSyncDefinition{}
	this.TaskType = taskType
	this.LookupFilePath = lookupFilePath
	this.IsLookupMergeOperation = isLookupMergeOperation
	return &this
}

// NewSaveToLookupNotificationSyncDefinitionWithDefaults instantiates a new SaveToLookupNotificationSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveToLookupNotificationSyncDefinitionWithDefaults() *SaveToLookupNotificationSyncDefinition {
	this := SaveToLookupNotificationSyncDefinition{}
	return &this
}

// GetLookupFilePath returns the LookupFilePath field value
func (o *SaveToLookupNotificationSyncDefinition) GetLookupFilePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupFilePath
}

// GetLookupFilePathOk returns a tuple with the LookupFilePath field value
// and a boolean to check if the value has been set.
func (o *SaveToLookupNotificationSyncDefinition) GetLookupFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupFilePath, true
}

// SetLookupFilePath sets field value
func (o *SaveToLookupNotificationSyncDefinition) SetLookupFilePath(v string) {
	o.LookupFilePath = v
}

// GetIsLookupMergeOperation returns the IsLookupMergeOperation field value
func (o *SaveToLookupNotificationSyncDefinition) GetIsLookupMergeOperation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLookupMergeOperation
}

// GetIsLookupMergeOperationOk returns a tuple with the IsLookupMergeOperation field value
// and a boolean to check if the value has been set.
func (o *SaveToLookupNotificationSyncDefinition) GetIsLookupMergeOperationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLookupMergeOperation, true
}

// SetIsLookupMergeOperation sets field value
func (o *SaveToLookupNotificationSyncDefinition) SetIsLookupMergeOperation(v bool) {
	o.IsLookupMergeOperation = v
}

func (o SaveToLookupNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedScheduleNotificationSyncDefinition, errScheduleNotificationSyncDefinition := json.Marshal(o.ScheduleNotificationSyncDefinition)
	if errScheduleNotificationSyncDefinition != nil {
		return []byte{}, errScheduleNotificationSyncDefinition
	}
	errScheduleNotificationSyncDefinition = json.Unmarshal([]byte(serializedScheduleNotificationSyncDefinition), &toSerialize)
	if errScheduleNotificationSyncDefinition != nil {
		return []byte{}, errScheduleNotificationSyncDefinition
	}
	if true {
		toSerialize["lookupFilePath"] = o.LookupFilePath
	}
	if true {
		toSerialize["isLookupMergeOperation"] = o.IsLookupMergeOperation
	}
	return json.Marshal(toSerialize)
}

type NullableSaveToLookupNotificationSyncDefinition struct {
	value *SaveToLookupNotificationSyncDefinition
	isSet bool
}

func (v NullableSaveToLookupNotificationSyncDefinition) Get() *SaveToLookupNotificationSyncDefinition {
	return v.value
}

func (v *NullableSaveToLookupNotificationSyncDefinition) Set(val *SaveToLookupNotificationSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveToLookupNotificationSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveToLookupNotificationSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveToLookupNotificationSyncDefinition(val *SaveToLookupNotificationSyncDefinition) *NullableSaveToLookupNotificationSyncDefinition {
	return &NullableSaveToLookupNotificationSyncDefinition{value: val, isSet: true}
}

func (v NullableSaveToLookupNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveToLookupNotificationSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


