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

// checks if the MonitorTemplatesLibraryBaseExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorTemplatesLibraryBaseExport{}

// MonitorTemplatesLibraryBaseExport struct for MonitorTemplatesLibraryBaseExport
type MonitorTemplatesLibraryBaseExport struct {
	// Name of the monitortemplate or folder.
	Name string `json:"name"`
	// Description of the monitortemplate or folder.
	Description *string `json:"description,omitempty"`
	// Type of the object model.
	Type string `json:"type"`
}

// NewMonitorTemplatesLibraryBaseExport instantiates a new MonitorTemplatesLibraryBaseExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorTemplatesLibraryBaseExport(name string, type_ string) *MonitorTemplatesLibraryBaseExport {
	this := MonitorTemplatesLibraryBaseExport{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewMonitorTemplatesLibraryBaseExportWithDefaults instantiates a new MonitorTemplatesLibraryBaseExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorTemplatesLibraryBaseExportWithDefaults() *MonitorTemplatesLibraryBaseExport {
	this := MonitorTemplatesLibraryBaseExport{}
	return &this
}

// GetName returns the Name field value
func (o *MonitorTemplatesLibraryBaseExport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryBaseExport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MonitorTemplatesLibraryBaseExport) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryBaseExport) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryBaseExport) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryBaseExport) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MonitorTemplatesLibraryBaseExport) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *MonitorTemplatesLibraryBaseExport) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryBaseExport) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MonitorTemplatesLibraryBaseExport) SetType(v string) {
	o.Type = v
}

func (o MonitorTemplatesLibraryBaseExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorTemplatesLibraryBaseExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMonitorTemplatesLibraryBaseExport struct {
	value *MonitorTemplatesLibraryBaseExport
	isSet bool
}

func (v NullableMonitorTemplatesLibraryBaseExport) Get() *MonitorTemplatesLibraryBaseExport {
	return v.value
}

func (v *NullableMonitorTemplatesLibraryBaseExport) Set(val *MonitorTemplatesLibraryBaseExport) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorTemplatesLibraryBaseExport) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorTemplatesLibraryBaseExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorTemplatesLibraryBaseExport(val *MonitorTemplatesLibraryBaseExport) *NullableMonitorTemplatesLibraryBaseExport {
	return &NullableMonitorTemplatesLibraryBaseExport{value: val, isSet: true}
}

func (v NullableMonitorTemplatesLibraryBaseExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorTemplatesLibraryBaseExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


