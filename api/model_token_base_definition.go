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

// checks if the TokenBaseDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenBaseDefinition{}

// TokenBaseDefinition struct for TokenBaseDefinition
type TokenBaseDefinition struct {
	// Name of the token.
	Name string `json:"name"`
	// Description of the token.
	Description *string `json:"description,omitempty"`
	// Status of the token. Can be `Active`, or `Inactive`.
	Status string `json:"status"`
	// Type of the token. Valid values: 1) CollectorRegistration
	Type string `json:"type"`
}

// NewTokenBaseDefinition instantiates a new TokenBaseDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenBaseDefinition(name string, status string, type_ string) *TokenBaseDefinition {
	this := TokenBaseDefinition{}
	this.Name = name
	this.Status = status
	this.Type = type_
	return &this
}

// NewTokenBaseDefinitionWithDefaults instantiates a new TokenBaseDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenBaseDefinitionWithDefaults() *TokenBaseDefinition {
	this := TokenBaseDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *TokenBaseDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenBaseDefinition) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TokenBaseDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TokenBaseDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TokenBaseDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *TokenBaseDefinition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TokenBaseDefinition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *TokenBaseDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenBaseDefinition) SetType(v string) {
	o.Type = v
}

func (o TokenBaseDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenBaseDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTokenBaseDefinition struct {
	value *TokenBaseDefinition
	isSet bool
}

func (v NullableTokenBaseDefinition) Get() *TokenBaseDefinition {
	return v.value
}

func (v *NullableTokenBaseDefinition) Set(val *TokenBaseDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenBaseDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenBaseDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenBaseDefinition(val *TokenBaseDefinition) *NullableTokenBaseDefinition {
	return &NullableTokenBaseDefinition{value: val, isSet: true}
}

func (v NullableTokenBaseDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenBaseDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


