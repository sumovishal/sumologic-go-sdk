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

// checks if the ContentCopyParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentCopyParams{}

// ContentCopyParams struct for ContentCopyParams
type ContentCopyParams struct {
	// Identifier of the parent folder to copy to.
	ParentId string `json:"parentId"`
	// Optionally provide a new name.
	Name *string `json:"name,omitempty"`
	// Optionally provide a new description.
	Description *string `json:"description,omitempty"`
}

// NewContentCopyParams instantiates a new ContentCopyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentCopyParams(parentId string) *ContentCopyParams {
	this := ContentCopyParams{}
	this.ParentId = parentId
	return &this
}

// NewContentCopyParamsWithDefaults instantiates a new ContentCopyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentCopyParamsWithDefaults() *ContentCopyParams {
	this := ContentCopyParams{}
	return &this
}

// GetParentId returns the ParentId field value
func (o *ContentCopyParams) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ContentCopyParams) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *ContentCopyParams) SetParentId(v string) {
	o.ParentId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContentCopyParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentCopyParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContentCopyParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContentCopyParams) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContentCopyParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentCopyParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContentCopyParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContentCopyParams) SetDescription(v string) {
	o.Description = &v
}

func (o ContentCopyParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentCopyParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parentId"] = o.ParentId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableContentCopyParams struct {
	value *ContentCopyParams
	isSet bool
}

func (v NullableContentCopyParams) Get() *ContentCopyParams {
	return v.value
}

func (v *NullableContentCopyParams) Set(val *ContentCopyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableContentCopyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableContentCopyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentCopyParams(val *ContentCopyParams) *NullableContentCopyParams {
	return &NullableContentCopyParams{value: val, isSet: true}
}

func (v NullableContentCopyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentCopyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


