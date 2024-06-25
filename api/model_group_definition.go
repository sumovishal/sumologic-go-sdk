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

// checks if the GroupDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupDefinition{}

// GroupDefinition Alert group scope that the schedule applies to
type GroupDefinition struct {
	// Field name of an alert group defined in monitors
	GroupKey string `json:"groupKey"`
	// Values of alert groups generated by monitors
	GroupValues []string `json:"groupValues"`
}

// NewGroupDefinition instantiates a new GroupDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDefinition(groupKey string, groupValues []string) *GroupDefinition {
	this := GroupDefinition{}
	this.GroupKey = groupKey
	this.GroupValues = groupValues
	return &this
}

// NewGroupDefinitionWithDefaults instantiates a new GroupDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDefinitionWithDefaults() *GroupDefinition {
	this := GroupDefinition{}
	return &this
}

// GetGroupKey returns the GroupKey field value
func (o *GroupDefinition) GetGroupKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupKey
}

// GetGroupKeyOk returns a tuple with the GroupKey field value
// and a boolean to check if the value has been set.
func (o *GroupDefinition) GetGroupKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupKey, true
}

// SetGroupKey sets field value
func (o *GroupDefinition) SetGroupKey(v string) {
	o.GroupKey = v
}

// GetGroupValues returns the GroupValues field value
func (o *GroupDefinition) GetGroupValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupValues
}

// GetGroupValuesOk returns a tuple with the GroupValues field value
// and a boolean to check if the value has been set.
func (o *GroupDefinition) GetGroupValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupValues, true
}

// SetGroupValues sets field value
func (o *GroupDefinition) SetGroupValues(v []string) {
	o.GroupValues = v
}

func (o GroupDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupKey"] = o.GroupKey
	toSerialize["groupValues"] = o.GroupValues
	return toSerialize, nil
}

type NullableGroupDefinition struct {
	value *GroupDefinition
	isSet bool
}

func (v NullableGroupDefinition) Get() *GroupDefinition {
	return v.value
}

func (v *NullableGroupDefinition) Set(val *GroupDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDefinition(val *GroupDefinition) *NullableGroupDefinition {
	return &NullableGroupDefinition{value: val, isSet: true}
}

func (v NullableGroupDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


