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

// checks if the AlertsLibraryBaseUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsLibraryBaseUpdate{}

// AlertsLibraryBaseUpdate struct for AlertsLibraryBaseUpdate
type AlertsLibraryBaseUpdate struct {
	// The name of the alert or folder.
	Name string `json:"name"`
	// The description of the alert or folder.
	Description *string `json:"description,omitempty"`
	// The version of the alert or folder.
	Version int64 `json:"version"`
	// Type of the object model.
	Type string `json:"type"`
}

// NewAlertsLibraryBaseUpdate instantiates a new AlertsLibraryBaseUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryBaseUpdate(name string, version int64, type_ string) *AlertsLibraryBaseUpdate {
	this := AlertsLibraryBaseUpdate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Version = version
	this.Type = type_
	return &this
}

// NewAlertsLibraryBaseUpdateWithDefaults instantiates a new AlertsLibraryBaseUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryBaseUpdateWithDefaults() *AlertsLibraryBaseUpdate {
	this := AlertsLibraryBaseUpdate{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetName returns the Name field value
func (o *AlertsLibraryBaseUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlertsLibraryBaseUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertsLibraryBaseUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertsLibraryBaseUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertsLibraryBaseUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value
func (o *AlertsLibraryBaseUpdate) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseUpdate) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AlertsLibraryBaseUpdate) SetVersion(v int64) {
	o.Version = v
}

// GetType returns the Type field value
func (o *AlertsLibraryBaseUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlertsLibraryBaseUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AlertsLibraryBaseUpdate) SetType(v string) {
	o.Type = v
}

func (o AlertsLibraryBaseUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsLibraryBaseUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["version"] = o.Version
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAlertsLibraryBaseUpdate struct {
	value *AlertsLibraryBaseUpdate
	isSet bool
}

func (v NullableAlertsLibraryBaseUpdate) Get() *AlertsLibraryBaseUpdate {
	return v.value
}

func (v *NullableAlertsLibraryBaseUpdate) Set(val *AlertsLibraryBaseUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryBaseUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryBaseUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryBaseUpdate(val *AlertsLibraryBaseUpdate) *NullableAlertsLibraryBaseUpdate {
	return &NullableAlertsLibraryBaseUpdate{value: val, isSet: true}
}

func (v NullableAlertsLibraryBaseUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryBaseUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


