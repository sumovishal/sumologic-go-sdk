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

// checks if the IncidentTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncidentTemplate{}

// IncidentTemplate struct for IncidentTemplate
type IncidentTemplate struct {
	// Unique identifier of the incident template.
	Id int32 `json:"id"`
	// Name of the incident template.
	Name string `json:"name"`
}

// NewIncidentTemplate instantiates a new IncidentTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTemplate(id int32, name string) *IncidentTemplate {
	this := IncidentTemplate{}
	this.Id = id
	this.Name = name
	return &this
}

// NewIncidentTemplateWithDefaults instantiates a new IncidentTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTemplateWithDefaults() *IncidentTemplate {
	this := IncidentTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *IncidentTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IncidentTemplate) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *IncidentTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IncidentTemplate) SetName(v string) {
	o.Name = v
}

func (o IncidentTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncidentTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableIncidentTemplate struct {
	value *IncidentTemplate
	isSet bool
}

func (v NullableIncidentTemplate) Get() *IncidentTemplate {
	return v.value
}

func (v *NullableIncidentTemplate) Set(val *IncidentTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTemplate(val *IncidentTemplate) *NullableIncidentTemplate {
	return &NullableIncidentTemplate{value: val, isSet: true}
}

func (v NullableIncidentTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


