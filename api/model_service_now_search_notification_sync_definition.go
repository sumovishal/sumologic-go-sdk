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

// ServiceNowSearchNotificationSyncDefinition struct for ServiceNowSearchNotificationSyncDefinition
type ServiceNowSearchNotificationSyncDefinition struct {
	ScheduleNotificationSyncDefinition
	// ServiceNow identifier.
	ExternalId string `json:"externalId"`
	Fields *ServiceNowFieldsSyncDefinition `json:"fields,omitempty"`
}

// NewServiceNowSearchNotificationSyncDefinition instantiates a new ServiceNowSearchNotificationSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNowSearchNotificationSyncDefinition(externalId string, taskType string) *ServiceNowSearchNotificationSyncDefinition {
	this := ServiceNowSearchNotificationSyncDefinition{}
	this.TaskType = taskType
	this.ExternalId = externalId
	return &this
}

// NewServiceNowSearchNotificationSyncDefinitionWithDefaults instantiates a new ServiceNowSearchNotificationSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNowSearchNotificationSyncDefinitionWithDefaults() *ServiceNowSearchNotificationSyncDefinition {
	this := ServiceNowSearchNotificationSyncDefinition{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *ServiceNowSearchNotificationSyncDefinition) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowSearchNotificationSyncDefinition) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *ServiceNowSearchNotificationSyncDefinition) SetExternalId(v string) {
	o.ExternalId = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ServiceNowSearchNotificationSyncDefinition) GetFields() ServiceNowFieldsSyncDefinition {
	if o == nil || o.Fields == nil {
		var ret ServiceNowFieldsSyncDefinition
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowSearchNotificationSyncDefinition) GetFieldsOk() (*ServiceNowFieldsSyncDefinition, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ServiceNowSearchNotificationSyncDefinition) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given ServiceNowFieldsSyncDefinition and assigns it to the Fields field.
func (o *ServiceNowSearchNotificationSyncDefinition) SetFields(v ServiceNowFieldsSyncDefinition) {
	o.Fields = &v
}

func (o ServiceNowSearchNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
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
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableServiceNowSearchNotificationSyncDefinition struct {
	value *ServiceNowSearchNotificationSyncDefinition
	isSet bool
}

func (v NullableServiceNowSearchNotificationSyncDefinition) Get() *ServiceNowSearchNotificationSyncDefinition {
	return v.value
}

func (v *NullableServiceNowSearchNotificationSyncDefinition) Set(val *ServiceNowSearchNotificationSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNowSearchNotificationSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNowSearchNotificationSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNowSearchNotificationSyncDefinition(val *ServiceNowSearchNotificationSyncDefinition) *NullableServiceNowSearchNotificationSyncDefinition {
	return &NullableServiceNowSearchNotificationSyncDefinition{value: val, isSet: true}
}

func (v NullableServiceNowSearchNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNowSearchNotificationSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


