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

// checks if the PathItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PathItem{}

// PathItem struct for PathItem
type PathItem struct {
	// Identifier of the path element.
	Id string `json:"id"`
	// Name of the path element.
	Name string `json:"name"`
}

// NewPathItem instantiates a new PathItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathItem(id string, name string) *PathItem {
	this := PathItem{}
	this.Id = id
	this.Name = name
	return &this
}

// NewPathItemWithDefaults instantiates a new PathItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathItemWithDefaults() *PathItem {
	this := PathItem{}
	return &this
}

// GetId returns the Id field value
func (o *PathItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PathItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PathItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PathItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PathItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PathItem) SetName(v string) {
	o.Name = v
}

func (o PathItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePathItem struct {
	value *PathItem
	isSet bool
}

func (v NullablePathItem) Get() *PathItem {
	return v.value
}

func (v *NullablePathItem) Set(val *PathItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePathItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePathItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathItem(val *PathItem) *NullablePathItem {
	return &NullablePathItem{value: val, isSet: true}
}

func (v NullablePathItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


