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

// checks if the OnDemandProvisioningInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnDemandProvisioningInfo{}

// OnDemandProvisioningInfo struct for OnDemandProvisioningInfo
type OnDemandProvisioningInfo struct {
	// First name attribute of the new user account.
	FirstNameAttribute *string `json:"firstNameAttribute,omitempty"`
	// Last name attribute of the new user account.
	LastNameAttribute *string `json:"lastNameAttribute,omitempty"`
	// Sumo Logic RBAC roles to be assigned when user accounts are provisioned.
	OnDemandProvisioningRoles []string `json:"onDemandProvisioningRoles"`
}

type _OnDemandProvisioningInfo OnDemandProvisioningInfo

// NewOnDemandProvisioningInfo instantiates a new OnDemandProvisioningInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnDemandProvisioningInfo(onDemandProvisioningRoles []string) *OnDemandProvisioningInfo {
	this := OnDemandProvisioningInfo{}
	var firstNameAttribute string = ""
	this.FirstNameAttribute = &firstNameAttribute
	var lastNameAttribute string = ""
	this.LastNameAttribute = &lastNameAttribute
	this.OnDemandProvisioningRoles = onDemandProvisioningRoles
	return &this
}

// NewOnDemandProvisioningInfoWithDefaults instantiates a new OnDemandProvisioningInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnDemandProvisioningInfoWithDefaults() *OnDemandProvisioningInfo {
	this := OnDemandProvisioningInfo{}
	var firstNameAttribute string = ""
	this.FirstNameAttribute = &firstNameAttribute
	var lastNameAttribute string = ""
	this.LastNameAttribute = &lastNameAttribute
	return &this
}

// GetFirstNameAttribute returns the FirstNameAttribute field value if set, zero value otherwise.
func (o *OnDemandProvisioningInfo) GetFirstNameAttribute() string {
	if o == nil || IsNil(o.FirstNameAttribute) {
		var ret string
		return ret
	}
	return *o.FirstNameAttribute
}

// GetFirstNameAttributeOk returns a tuple with the FirstNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandProvisioningInfo) GetFirstNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.FirstNameAttribute) {
		return nil, false
	}
	return o.FirstNameAttribute, true
}

// HasFirstNameAttribute returns a boolean if a field has been set.
func (o *OnDemandProvisioningInfo) HasFirstNameAttribute() bool {
	if o != nil && !IsNil(o.FirstNameAttribute) {
		return true
	}

	return false
}

// SetFirstNameAttribute gets a reference to the given string and assigns it to the FirstNameAttribute field.
func (o *OnDemandProvisioningInfo) SetFirstNameAttribute(v string) {
	o.FirstNameAttribute = &v
}

// GetLastNameAttribute returns the LastNameAttribute field value if set, zero value otherwise.
func (o *OnDemandProvisioningInfo) GetLastNameAttribute() string {
	if o == nil || IsNil(o.LastNameAttribute) {
		var ret string
		return ret
	}
	return *o.LastNameAttribute
}

// GetLastNameAttributeOk returns a tuple with the LastNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandProvisioningInfo) GetLastNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.LastNameAttribute) {
		return nil, false
	}
	return o.LastNameAttribute, true
}

// HasLastNameAttribute returns a boolean if a field has been set.
func (o *OnDemandProvisioningInfo) HasLastNameAttribute() bool {
	if o != nil && !IsNil(o.LastNameAttribute) {
		return true
	}

	return false
}

// SetLastNameAttribute gets a reference to the given string and assigns it to the LastNameAttribute field.
func (o *OnDemandProvisioningInfo) SetLastNameAttribute(v string) {
	o.LastNameAttribute = &v
}

// GetOnDemandProvisioningRoles returns the OnDemandProvisioningRoles field value
func (o *OnDemandProvisioningInfo) GetOnDemandProvisioningRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OnDemandProvisioningRoles
}

// GetOnDemandProvisioningRolesOk returns a tuple with the OnDemandProvisioningRoles field value
// and a boolean to check if the value has been set.
func (o *OnDemandProvisioningInfo) GetOnDemandProvisioningRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnDemandProvisioningRoles, true
}

// SetOnDemandProvisioningRoles sets field value
func (o *OnDemandProvisioningInfo) SetOnDemandProvisioningRoles(v []string) {
	o.OnDemandProvisioningRoles = v
}

func (o OnDemandProvisioningInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnDemandProvisioningInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstNameAttribute) {
		toSerialize["firstNameAttribute"] = o.FirstNameAttribute
	}
	if !IsNil(o.LastNameAttribute) {
		toSerialize["lastNameAttribute"] = o.LastNameAttribute
	}
	toSerialize["onDemandProvisioningRoles"] = o.OnDemandProvisioningRoles
	return toSerialize, nil
}

func (o *OnDemandProvisioningInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"onDemandProvisioningRoles",
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

	varOnDemandProvisioningInfo := _OnDemandProvisioningInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnDemandProvisioningInfo)

	if err != nil {
		return err
	}

	*o = OnDemandProvisioningInfo(varOnDemandProvisioningInfo)

	return err
}

type NullableOnDemandProvisioningInfo struct {
	value *OnDemandProvisioningInfo
	isSet bool
}

func (v NullableOnDemandProvisioningInfo) Get() *OnDemandProvisioningInfo {
	return v.value
}

func (v *NullableOnDemandProvisioningInfo) Set(val *OnDemandProvisioningInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOnDemandProvisioningInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOnDemandProvisioningInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnDemandProvisioningInfo(val *OnDemandProvisioningInfo) *NullableOnDemandProvisioningInfo {
	return &NullableOnDemandProvisioningInfo{value: val, isSet: true}
}

func (v NullableOnDemandProvisioningInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnDemandProvisioningInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


