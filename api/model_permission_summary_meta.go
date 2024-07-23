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

// checks if the PermissionSummaryMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionSummaryMeta{}

// PermissionSummaryMeta Permission Summary with additional information like inheritance, revocation, etc about the permission.
type PermissionSummaryMeta struct {
	// Name of the permission. Example values are: `Read`, `Update`, `Create`, etc.
	Name string `json:"name"`
	// A true value implies that the permission is inherited from some ancestors of the resource. A false value implies that the permission is explicitly assigned to the resource.
	IsInherited bool `json:"isInherited"`
	// A true value implies that the permission is explicitly assigned to the resource. A false value implies that the permission is not explicitly assigned to the resource.
	IsExplicit bool `json:"isExplicit"`
	// A true value implies that the capability required for this permission has been revoked.
	IsRevoked bool `json:"isRevoked"`
	// A true value implies that the permission is recursively cascaded down to all the direct and indirect children of the resource.
	IsRecursive bool `json:"isRecursive"`
	// A true value implies that the permission is defined by the system on the resource and can not be modified by the user. A false value implies that the permission is defined by the user on the resource and can be modified by the user.
	IsSystemDefined bool `json:"isSystemDefined"`
}

type _PermissionSummaryMeta PermissionSummaryMeta

// NewPermissionSummaryMeta instantiates a new PermissionSummaryMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionSummaryMeta(name string, isInherited bool, isExplicit bool, isRevoked bool, isRecursive bool, isSystemDefined bool) *PermissionSummaryMeta {
	this := PermissionSummaryMeta{}
	this.Name = name
	this.IsInherited = isInherited
	this.IsExplicit = isExplicit
	this.IsRevoked = isRevoked
	this.IsRecursive = isRecursive
	this.IsSystemDefined = isSystemDefined
	return &this
}

// NewPermissionSummaryMetaWithDefaults instantiates a new PermissionSummaryMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionSummaryMetaWithDefaults() *PermissionSummaryMeta {
	this := PermissionSummaryMeta{}
	return &this
}

// GetName returns the Name field value
func (o *PermissionSummaryMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PermissionSummaryMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PermissionSummaryMeta) SetName(v string) {
	o.Name = v
}

// GetIsInherited returns the IsInherited field value
func (o *PermissionSummaryMeta) GetIsInherited() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInherited
}

// GetIsInheritedOk returns a tuple with the IsInherited field value
// and a boolean to check if the value has been set.
func (o *PermissionSummaryMeta) GetIsInheritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInherited, true
}

// SetIsInherited sets field value
func (o *PermissionSummaryMeta) SetIsInherited(v bool) {
	o.IsInherited = v
}

// GetIsExplicit returns the IsExplicit field value
func (o *PermissionSummaryMeta) GetIsExplicit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExplicit
}

// GetIsExplicitOk returns a tuple with the IsExplicit field value
// and a boolean to check if the value has been set.
func (o *PermissionSummaryMeta) GetIsExplicitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExplicit, true
}

// SetIsExplicit sets field value
func (o *PermissionSummaryMeta) SetIsExplicit(v bool) {
	o.IsExplicit = v
}

// GetIsRevoked returns the IsRevoked field value
func (o *PermissionSummaryMeta) GetIsRevoked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRevoked
}

// GetIsRevokedOk returns a tuple with the IsRevoked field value
// and a boolean to check if the value has been set.
func (o *PermissionSummaryMeta) GetIsRevokedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRevoked, true
}

// SetIsRevoked sets field value
func (o *PermissionSummaryMeta) SetIsRevoked(v bool) {
	o.IsRevoked = v
}

// GetIsRecursive returns the IsRecursive field value
func (o *PermissionSummaryMeta) GetIsRecursive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRecursive
}

// GetIsRecursiveOk returns a tuple with the IsRecursive field value
// and a boolean to check if the value has been set.
func (o *PermissionSummaryMeta) GetIsRecursiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRecursive, true
}

// SetIsRecursive sets field value
func (o *PermissionSummaryMeta) SetIsRecursive(v bool) {
	o.IsRecursive = v
}

// GetIsSystemDefined returns the IsSystemDefined field value
func (o *PermissionSummaryMeta) GetIsSystemDefined() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystemDefined
}

// GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field value
// and a boolean to check if the value has been set.
func (o *PermissionSummaryMeta) GetIsSystemDefinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystemDefined, true
}

// SetIsSystemDefined sets field value
func (o *PermissionSummaryMeta) SetIsSystemDefined(v bool) {
	o.IsSystemDefined = v
}

func (o PermissionSummaryMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionSummaryMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["isInherited"] = o.IsInherited
	toSerialize["isExplicit"] = o.IsExplicit
	toSerialize["isRevoked"] = o.IsRevoked
	toSerialize["isRecursive"] = o.IsRecursive
	toSerialize["isSystemDefined"] = o.IsSystemDefined
	return toSerialize, nil
}

func (o *PermissionSummaryMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"isInherited",
		"isExplicit",
		"isRevoked",
		"isRecursive",
		"isSystemDefined",
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

	varPermissionSummaryMeta := _PermissionSummaryMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionSummaryMeta)

	if err != nil {
		return err
	}

	*o = PermissionSummaryMeta(varPermissionSummaryMeta)

	return err
}

type NullablePermissionSummaryMeta struct {
	value *PermissionSummaryMeta
	isSet bool
}

func (v NullablePermissionSummaryMeta) Get() *PermissionSummaryMeta {
	return v.value
}

func (v *NullablePermissionSummaryMeta) Set(val *PermissionSummaryMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionSummaryMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionSummaryMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionSummaryMeta(val *PermissionSummaryMeta) *NullablePermissionSummaryMeta {
	return &NullablePermissionSummaryMeta{value: val, isSet: true}
}

func (v NullablePermissionSummaryMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionSummaryMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


