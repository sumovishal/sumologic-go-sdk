/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContentAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentAllOf{}

// ContentAllOf struct for ContentAllOf
type ContentAllOf struct {
	// Identifier of the content item.
	Id string `json:"id"`
	// The name of the content item.
	Name string `json:"name"`
	// Type of the content item. Supported values are:   1. Folder   2. Search   3. Report (for old dashboards)   4. Dashboard (for new dashboards)   5. Lookups
	ItemType string `json:"itemType"`
	// Identifier of the parent content item.
	ParentId string `json:"parentId"`
	// List of permissions the user has on the content item.
	Permissions []string `json:"permissions"`
}

// NewContentAllOf instantiates a new ContentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentAllOf(id string, name string, itemType string, parentId string, permissions []string) *ContentAllOf {
	this := ContentAllOf{}
	this.Id = id
	this.Name = name
	this.ItemType = itemType
	this.ParentId = parentId
	this.Permissions = permissions
	return &this
}

// NewContentAllOfWithDefaults instantiates a new ContentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentAllOfWithDefaults() *ContentAllOf {
	this := ContentAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *ContentAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContentAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContentAllOf) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ContentAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContentAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContentAllOf) SetName(v string) {
	o.Name = v
}

// GetItemType returns the ItemType field value
func (o *ContentAllOf) GetItemType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
func (o *ContentAllOf) GetItemTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemType, true
}

// SetItemType sets field value
func (o *ContentAllOf) SetItemType(v string) {
	o.ItemType = v
}

// GetParentId returns the ParentId field value
func (o *ContentAllOf) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ContentAllOf) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *ContentAllOf) SetParentId(v string) {
	o.ParentId = v
}

// GetPermissions returns the Permissions field value
func (o *ContentAllOf) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ContentAllOf) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ContentAllOf) SetPermissions(v []string) {
	o.Permissions = v
}

func (o ContentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["itemType"] = o.ItemType
	toSerialize["parentId"] = o.ParentId
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

type NullableContentAllOf struct {
	value *ContentAllOf
	isSet bool
}

func (v NullableContentAllOf) Get() *ContentAllOf {
	return v.value
}

func (v *NullableContentAllOf) Set(val *ContentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentAllOf(val *ContentAllOf) *NullableContentAllOf {
	return &NullableContentAllOf{value: val, isSet: true}
}

func (v NullableContentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


