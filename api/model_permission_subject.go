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

// checks if the PermissionSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionSubject{}

// PermissionSubject Identifier for the entity (subject) that is granted the permission on resource(s).
type PermissionSubject struct {
	// Type of subject for the permission. Valid values are: `user` or `role` or `org`.
	SubjectType string `json:"subjectType" validate:"regexp=^(user|role|org)$"`
	// The identifier that belongs to the subject type chosen above. For e.g. if the subjectType is set to `user`, subjectId should be the identifier of a user (same goes for `role` or `org` subjectType).
	SubjectId string `json:"subjectId"`
}

type _PermissionSubject PermissionSubject

// NewPermissionSubject instantiates a new PermissionSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionSubject(subjectType string, subjectId string) *PermissionSubject {
	this := PermissionSubject{}
	this.SubjectType = subjectType
	this.SubjectId = subjectId
	return &this
}

// NewPermissionSubjectWithDefaults instantiates a new PermissionSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionSubjectWithDefaults() *PermissionSubject {
	this := PermissionSubject{}
	return &this
}

// GetSubjectType returns the SubjectType field value
func (o *PermissionSubject) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *PermissionSubject) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value
func (o *PermissionSubject) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetSubjectId returns the SubjectId field value
func (o *PermissionSubject) GetSubjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value
// and a boolean to check if the value has been set.
func (o *PermissionSubject) GetSubjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectId, true
}

// SetSubjectId sets field value
func (o *PermissionSubject) SetSubjectId(v string) {
	o.SubjectId = v
}

func (o PermissionSubject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionSubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subjectType"] = o.SubjectType
	toSerialize["subjectId"] = o.SubjectId
	return toSerialize, nil
}

func (o *PermissionSubject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subjectType",
		"subjectId",
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

	varPermissionSubject := _PermissionSubject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionSubject)

	if err != nil {
		return err
	}

	*o = PermissionSubject(varPermissionSubject)

	return err
}

type NullablePermissionSubject struct {
	value *PermissionSubject
	isSet bool
}

func (v NullablePermissionSubject) Get() *PermissionSubject {
	return v.value
}

func (v *NullablePermissionSubject) Set(val *PermissionSubject) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionSubject) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionSubject(val *PermissionSubject) *NullablePermissionSubject {
	return &NullablePermissionSubject{value: val, isSet: true}
}

func (v NullablePermissionSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


