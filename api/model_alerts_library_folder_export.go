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

// checks if the AlertsLibraryFolderExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsLibraryFolderExport{}

// AlertsLibraryFolderExport struct for AlertsLibraryFolderExport
type AlertsLibraryFolderExport struct {
	AlertsLibraryBaseExport
	// The items in the folder. A multi-type list of types alert or folder.
	Children []AlertsLibraryBaseExport `json:"children,omitempty"`
}

type _AlertsLibraryFolderExport AlertsLibraryFolderExport

// NewAlertsLibraryFolderExport instantiates a new AlertsLibraryFolderExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryFolderExport(name string, type_ string) *AlertsLibraryFolderExport {
	this := AlertsLibraryFolderExport{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewAlertsLibraryFolderExportWithDefaults instantiates a new AlertsLibraryFolderExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryFolderExportWithDefaults() *AlertsLibraryFolderExport {
	this := AlertsLibraryFolderExport{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *AlertsLibraryFolderExport) GetChildren() []AlertsLibraryBaseExport {
	if o == nil || IsNil(o.Children) {
		var ret []AlertsLibraryBaseExport
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryFolderExport) GetChildrenOk() ([]AlertsLibraryBaseExport, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *AlertsLibraryFolderExport) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []AlertsLibraryBaseExport and assigns it to the Children field.
func (o *AlertsLibraryFolderExport) SetChildren(v []AlertsLibraryBaseExport) {
	o.Children = v
}

func (o AlertsLibraryFolderExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsLibraryFolderExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAlertsLibraryBaseExport, errAlertsLibraryBaseExport := json.Marshal(o.AlertsLibraryBaseExport)
	if errAlertsLibraryBaseExport != nil {
		return map[string]interface{}{}, errAlertsLibraryBaseExport
	}
	errAlertsLibraryBaseExport = json.Unmarshal([]byte(serializedAlertsLibraryBaseExport), &toSerialize)
	if errAlertsLibraryBaseExport != nil {
		return map[string]interface{}{}, errAlertsLibraryBaseExport
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

func (o *AlertsLibraryFolderExport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varAlertsLibraryFolderExport := _AlertsLibraryFolderExport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertsLibraryFolderExport)

	if err != nil {
		return err
	}

	*o = AlertsLibraryFolderExport(varAlertsLibraryFolderExport)

	return err
}

type NullableAlertsLibraryFolderExport struct {
	value *AlertsLibraryFolderExport
	isSet bool
}

func (v NullableAlertsLibraryFolderExport) Get() *AlertsLibraryFolderExport {
	return v.value
}

func (v *NullableAlertsLibraryFolderExport) Set(val *AlertsLibraryFolderExport) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryFolderExport) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryFolderExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryFolderExport(val *AlertsLibraryFolderExport) *NullableAlertsLibraryFolderExport {
	return &NullableAlertsLibraryFolderExport{value: val, isSet: true}
}

func (v NullableAlertsLibraryFolderExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryFolderExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


