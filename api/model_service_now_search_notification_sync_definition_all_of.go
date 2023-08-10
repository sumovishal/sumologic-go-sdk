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

// checks if the ServiceNowSearchNotificationSyncDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceNowSearchNotificationSyncDefinitionAllOf{}

// ServiceNowSearchNotificationSyncDefinitionAllOf struct for ServiceNowSearchNotificationSyncDefinitionAllOf
type ServiceNowSearchNotificationSyncDefinitionAllOf struct {
	// ServiceNow identifier.
	ExternalId string `json:"externalId"`
	Fields *ServiceNowFieldsSyncDefinition `json:"fields,omitempty"`
}

// NewServiceNowSearchNotificationSyncDefinitionAllOf instantiates a new ServiceNowSearchNotificationSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNowSearchNotificationSyncDefinitionAllOf(externalId string) *ServiceNowSearchNotificationSyncDefinitionAllOf {
	this := ServiceNowSearchNotificationSyncDefinitionAllOf{}
	this.ExternalId = externalId
	return &this
}

// NewServiceNowSearchNotificationSyncDefinitionAllOfWithDefaults instantiates a new ServiceNowSearchNotificationSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNowSearchNotificationSyncDefinitionAllOfWithDefaults() *ServiceNowSearchNotificationSyncDefinitionAllOf {
	this := ServiceNowSearchNotificationSyncDefinitionAllOf{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) SetExternalId(v string) {
	o.ExternalId = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) GetFields() ServiceNowFieldsSyncDefinition {
	if o == nil || IsNil(o.Fields) {
		var ret ServiceNowFieldsSyncDefinition
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) GetFieldsOk() (*ServiceNowFieldsSyncDefinition, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ServiceNowFieldsSyncDefinition and assigns it to the Fields field.
func (o *ServiceNowSearchNotificationSyncDefinitionAllOf) SetFields(v ServiceNowFieldsSyncDefinition) {
	o.Fields = &v
}

func (o ServiceNowSearchNotificationSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceNowSearchNotificationSyncDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableServiceNowSearchNotificationSyncDefinitionAllOf struct {
	value *ServiceNowSearchNotificationSyncDefinitionAllOf
	isSet bool
}

func (v NullableServiceNowSearchNotificationSyncDefinitionAllOf) Get() *ServiceNowSearchNotificationSyncDefinitionAllOf {
	return v.value
}

func (v *NullableServiceNowSearchNotificationSyncDefinitionAllOf) Set(val *ServiceNowSearchNotificationSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNowSearchNotificationSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNowSearchNotificationSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNowSearchNotificationSyncDefinitionAllOf(val *ServiceNowSearchNotificationSyncDefinitionAllOf) *NullableServiceNowSearchNotificationSyncDefinitionAllOf {
	return &NullableServiceNowSearchNotificationSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableServiceNowSearchNotificationSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNowSearchNotificationSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


