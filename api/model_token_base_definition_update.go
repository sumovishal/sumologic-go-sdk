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

// checks if the TokenBaseDefinitionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenBaseDefinitionUpdate{}

// TokenBaseDefinitionUpdate struct for TokenBaseDefinitionUpdate
type TokenBaseDefinitionUpdate struct {
	// Name of the token.
	Name string `json:"name"`
	// Description of the token.
	Description *string `json:"description,omitempty"`
	// Status of the token. Can be `Active`, or `Inactive`.
	Status string `json:"status" validate:"regexp=^(Active|Inactive)$"`
	// Type of the token. Valid values: 1) CollectorRegistration
	Type string `json:"type" validate:"regexp=^(CollectorRegistration)$"`
	// Version of the token.
	Version int64 `json:"version"`
}

type _TokenBaseDefinitionUpdate TokenBaseDefinitionUpdate

// NewTokenBaseDefinitionUpdate instantiates a new TokenBaseDefinitionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenBaseDefinitionUpdate(name string, status string, type_ string, version int64) *TokenBaseDefinitionUpdate {
	this := TokenBaseDefinitionUpdate{}
	this.Name = name
	this.Status = status
	this.Type = type_
	this.Version = version
	return &this
}

// NewTokenBaseDefinitionUpdateWithDefaults instantiates a new TokenBaseDefinitionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenBaseDefinitionUpdateWithDefaults() *TokenBaseDefinitionUpdate {
	this := TokenBaseDefinitionUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *TokenBaseDefinitionUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinitionUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenBaseDefinitionUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TokenBaseDefinitionUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinitionUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TokenBaseDefinitionUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TokenBaseDefinitionUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *TokenBaseDefinitionUpdate) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinitionUpdate) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TokenBaseDefinitionUpdate) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *TokenBaseDefinitionUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinitionUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenBaseDefinitionUpdate) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *TokenBaseDefinitionUpdate) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TokenBaseDefinitionUpdate) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *TokenBaseDefinitionUpdate) SetVersion(v int64) {
	o.Version = v
}

func (o TokenBaseDefinitionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenBaseDefinitionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *TokenBaseDefinitionUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"status",
		"type",
		"version",
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

	varTokenBaseDefinitionUpdate := _TokenBaseDefinitionUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenBaseDefinitionUpdate)

	if err != nil {
		return err
	}

	*o = TokenBaseDefinitionUpdate(varTokenBaseDefinitionUpdate)

	return err
}

type NullableTokenBaseDefinitionUpdate struct {
	value *TokenBaseDefinitionUpdate
	isSet bool
}

func (v NullableTokenBaseDefinitionUpdate) Get() *TokenBaseDefinitionUpdate {
	return v.value
}

func (v *NullableTokenBaseDefinitionUpdate) Set(val *TokenBaseDefinitionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenBaseDefinitionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenBaseDefinitionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenBaseDefinitionUpdate(val *TokenBaseDefinitionUpdate) *NullableTokenBaseDefinitionUpdate {
	return &NullableTokenBaseDefinitionUpdate{value: val, isSet: true}
}

func (v NullableTokenBaseDefinitionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenBaseDefinitionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


