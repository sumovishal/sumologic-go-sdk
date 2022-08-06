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

// AlertEntityInfo An entity's name and Id.
type AlertEntityInfo struct {
	// Identifier of the entity.
	EntityId *string `json:"entityId,omitempty"`
	// Name of the entity.
	EntityName *string `json:"entityName,omitempty"`
}

// NewAlertEntityInfo instantiates a new AlertEntityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertEntityInfo() *AlertEntityInfo {
	this := AlertEntityInfo{}
	return &this
}

// NewAlertEntityInfoWithDefaults instantiates a new AlertEntityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertEntityInfoWithDefaults() *AlertEntityInfo {
	this := AlertEntityInfo{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AlertEntityInfo) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEntityInfo) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AlertEntityInfo) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AlertEntityInfo) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *AlertEntityInfo) GetEntityName() string {
	if o == nil || o.EntityName == nil {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEntityInfo) GetEntityNameOk() (*string, bool) {
	if o == nil || o.EntityName == nil {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *AlertEntityInfo) HasEntityName() bool {
	if o != nil && o.EntityName != nil {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *AlertEntityInfo) SetEntityName(v string) {
	o.EntityName = &v
}

func (o AlertEntityInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.EntityName != nil {
		toSerialize["entityName"] = o.EntityName
	}
	return json.Marshal(toSerialize)
}

type NullableAlertEntityInfo struct {
	value *AlertEntityInfo
	isSet bool
}

func (v NullableAlertEntityInfo) Get() *AlertEntityInfo {
	return v.value
}

func (v *NullableAlertEntityInfo) Set(val *AlertEntityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertEntityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertEntityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertEntityInfo(val *AlertEntityInfo) *NullableAlertEntityInfo {
	return &NullableAlertEntityInfo{value: val, isSet: true}
}

func (v NullableAlertEntityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertEntityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


