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

// ConnectionDefinition struct for ConnectionDefinition
type ConnectionDefinition struct {
	// Type of connection. Valid values are `WebhookDefinition`, `ServiceNowDefinition`.
	Type string `json:"type"`
	// Name of connection. Name should be a valid alphanumeric value.
	Name string `json:"name"`
	// Description of the connection.
	Description *string `json:"description,omitempty"`
}

// NewConnectionDefinition instantiates a new ConnectionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionDefinition(type_ string, name string) *ConnectionDefinition {
	this := ConnectionDefinition{}
	this.Type = type_
	this.Name = name
	var description string = ""
	this.Description = &description
	return &this
}

// NewConnectionDefinitionWithDefaults instantiates a new ConnectionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionDefinitionWithDefaults() *ConnectionDefinition {
	this := ConnectionDefinition{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetType returns the Type field value
func (o *ConnectionDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectionDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectionDefinition) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ConnectionDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectionDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectionDefinition) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectionDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectionDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectionDefinition) SetDescription(v string) {
	o.Description = &v
}

func (o ConnectionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionDefinition struct {
	value *ConnectionDefinition
	isSet bool
}

func (v NullableConnectionDefinition) Get() *ConnectionDefinition {
	return v.value
}

func (v *NullableConnectionDefinition) Set(val *ConnectionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionDefinition(val *ConnectionDefinition) *NullableConnectionDefinition {
	return &NullableConnectionDefinition{value: val, isSet: true}
}

func (v NullableConnectionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


