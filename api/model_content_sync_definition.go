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

// checks if the ContentSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSyncDefinition{}

// ContentSyncDefinition struct for ContentSyncDefinition
type ContentSyncDefinition struct {
	// The content item type. **Note:**  - `MewboardSyncDefinition` _is depreciated, and will soon be removed. Please use_ `DashboardV2SyncDefinition`    _instead_.  - Dashboard links are not supported for dashboards.
	Type string `json:"type"`
	// The name of the item.
	Name string `json:"name"`
}

// NewContentSyncDefinition instantiates a new ContentSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSyncDefinition(type_ string, name string) *ContentSyncDefinition {
	this := ContentSyncDefinition{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewContentSyncDefinitionWithDefaults instantiates a new ContentSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSyncDefinitionWithDefaults() *ContentSyncDefinition {
	this := ContentSyncDefinition{}
	return &this
}

// GetType returns the Type field value
func (o *ContentSyncDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentSyncDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentSyncDefinition) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ContentSyncDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContentSyncDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContentSyncDefinition) SetName(v string) {
	o.Name = v
}

func (o ContentSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableContentSyncDefinition struct {
	value *ContentSyncDefinition
	isSet bool
}

func (v NullableContentSyncDefinition) Get() *ContentSyncDefinition {
	return v.value
}

func (v *NullableContentSyncDefinition) Set(val *ContentSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSyncDefinition(val *ContentSyncDefinition) *NullableContentSyncDefinition {
	return &NullableContentSyncDefinition{value: val, isSet: true}
}

func (v NullableContentSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


