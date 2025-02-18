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

// checks if the CreateServiceAccountDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceAccountDefinition{}

// CreateServiceAccountDefinition struct for CreateServiceAccountDefinition
type CreateServiceAccountDefinition struct {
	// Name of the service account.
	Name string `json:"name"`
	// Email address of the service account.
	Email string `json:"email"`
	// List of roleIds associated with the service account.
	RoleIds []string `json:"roleIds"`
}

type _CreateServiceAccountDefinition CreateServiceAccountDefinition

// NewCreateServiceAccountDefinition instantiates a new CreateServiceAccountDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceAccountDefinition(name string, email string, roleIds []string) *CreateServiceAccountDefinition {
	this := CreateServiceAccountDefinition{}
	this.Name = name
	this.Email = email
	this.RoleIds = roleIds
	return &this
}

// NewCreateServiceAccountDefinitionWithDefaults instantiates a new CreateServiceAccountDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceAccountDefinitionWithDefaults() *CreateServiceAccountDefinition {
	this := CreateServiceAccountDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *CreateServiceAccountDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateServiceAccountDefinition) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *CreateServiceAccountDefinition) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountDefinition) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateServiceAccountDefinition) SetEmail(v string) {
	o.Email = v
}

// GetRoleIds returns the RoleIds field value
func (o *CreateServiceAccountDefinition) GetRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountDefinition) GetRoleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// SetRoleIds sets field value
func (o *CreateServiceAccountDefinition) SetRoleIds(v []string) {
	o.RoleIds = v
}

func (o CreateServiceAccountDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceAccountDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["roleIds"] = o.RoleIds
	return toSerialize, nil
}

func (o *CreateServiceAccountDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"email",
		"roleIds",
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

	varCreateServiceAccountDefinition := _CreateServiceAccountDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateServiceAccountDefinition)

	if err != nil {
		return err
	}

	*o = CreateServiceAccountDefinition(varCreateServiceAccountDefinition)

	return err
}

type NullableCreateServiceAccountDefinition struct {
	value *CreateServiceAccountDefinition
	isSet bool
}

func (v NullableCreateServiceAccountDefinition) Get() *CreateServiceAccountDefinition {
	return v.value
}

func (v *NullableCreateServiceAccountDefinition) Set(val *CreateServiceAccountDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceAccountDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceAccountDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceAccountDefinition(val *CreateServiceAccountDefinition) *NullableCreateServiceAccountDefinition {
	return &NullableCreateServiceAccountDefinition{value: val, isSet: true}
}

func (v NullableCreateServiceAccountDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceAccountDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


