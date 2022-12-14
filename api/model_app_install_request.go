/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// AppInstallRequest JSON object containing name, description, destinationFolderId, and dataSourceType.
type AppInstallRequest struct {
	// Preferred name of the app to be installed. This will be the name of the app in the selected installation folder.
	Name string `json:"name"`
	// Preferred description of the app to be installed. This will be displayed as the app description in the selected installation folder.
	Description string `json:"description"`
	// Identifier of the folder in which the app will be installed in hexadecimal format.
	DestinationFolderId string `json:"destinationFolderId"`
	// Dictionary of properties specifying log-source name and value.
	DataSourceValues *map[string]string `json:"dataSourceValues,omitempty"`
}

// NewAppInstallRequest instantiates a new AppInstallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInstallRequest(name string, description string, destinationFolderId string) *AppInstallRequest {
	this := AppInstallRequest{}
	this.Name = name
	this.Description = description
	this.DestinationFolderId = destinationFolderId
	return &this
}

// NewAppInstallRequestWithDefaults instantiates a new AppInstallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInstallRequestWithDefaults() *AppInstallRequest {
	this := AppInstallRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AppInstallRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppInstallRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppInstallRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AppInstallRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AppInstallRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AppInstallRequest) SetDescription(v string) {
	o.Description = v
}

// GetDestinationFolderId returns the DestinationFolderId field value
func (o *AppInstallRequest) GetDestinationFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationFolderId
}

// GetDestinationFolderIdOk returns a tuple with the DestinationFolderId field value
// and a boolean to check if the value has been set.
func (o *AppInstallRequest) GetDestinationFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationFolderId, true
}

// SetDestinationFolderId sets field value
func (o *AppInstallRequest) SetDestinationFolderId(v string) {
	o.DestinationFolderId = v
}

// GetDataSourceValues returns the DataSourceValues field value if set, zero value otherwise.
func (o *AppInstallRequest) GetDataSourceValues() map[string]string {
	if o == nil || o.DataSourceValues == nil {
		var ret map[string]string
		return ret
	}
	return *o.DataSourceValues
}

// GetDataSourceValuesOk returns a tuple with the DataSourceValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInstallRequest) GetDataSourceValuesOk() (*map[string]string, bool) {
	if o == nil || o.DataSourceValues == nil {
		return nil, false
	}
	return o.DataSourceValues, true
}

// HasDataSourceValues returns a boolean if a field has been set.
func (o *AppInstallRequest) HasDataSourceValues() bool {
	if o != nil && o.DataSourceValues != nil {
		return true
	}

	return false
}

// SetDataSourceValues gets a reference to the given map[string]string and assigns it to the DataSourceValues field.
func (o *AppInstallRequest) SetDataSourceValues(v map[string]string) {
	o.DataSourceValues = &v
}

func (o AppInstallRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["destinationFolderId"] = o.DestinationFolderId
	}
	if o.DataSourceValues != nil {
		toSerialize["dataSourceValues"] = o.DataSourceValues
	}
	return json.Marshal(toSerialize)
}

type NullableAppInstallRequest struct {
	value *AppInstallRequest
	isSet bool
}

func (v NullableAppInstallRequest) Get() *AppInstallRequest {
	return v.value
}

func (v *NullableAppInstallRequest) Set(val *AppInstallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInstallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInstallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInstallRequest(val *AppInstallRequest) *NullableAppInstallRequest {
	return &NullableAppInstallRequest{value: val, isSet: true}
}

func (v NullableAppInstallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInstallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


