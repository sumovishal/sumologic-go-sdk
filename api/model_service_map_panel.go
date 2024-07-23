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

// checks if the ServiceMapPanel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceMapPanel{}

// ServiceMapPanel struct for ServiceMapPanel
type ServiceMapPanel struct {
	Panel
	// Filter services by the application custom tag.
	Application *string `json:"application,omitempty"`
	// Show only the specific service and its connections to other services.
	Service *string `json:"service,omitempty"`
	// Show remote services, like databases or external calls, automatically detected in client traffic.
	ShowRemoteServices *bool `json:"showRemoteServices,omitempty"`
	// Show only service map data specific to the provided environment.
	Environment *string `json:"environment,omitempty"`
}

type _ServiceMapPanel ServiceMapPanel

// NewServiceMapPanel instantiates a new ServiceMapPanel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceMapPanel(key string, panelType string) *ServiceMapPanel {
	this := ServiceMapPanel{}
	this.Key = key
	var keepVisualSettingsConsistentWithParent bool = true
	this.KeepVisualSettingsConsistentWithParent = &keepVisualSettingsConsistentWithParent
	this.PanelType = panelType
	return &this
}

// NewServiceMapPanelWithDefaults instantiates a new ServiceMapPanel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceMapPanelWithDefaults() *ServiceMapPanel {
	this := ServiceMapPanel{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ServiceMapPanel) GetApplication() string {
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapPanel) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ServiceMapPanel) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ServiceMapPanel) SetApplication(v string) {
	o.Application = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceMapPanel) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapPanel) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceMapPanel) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ServiceMapPanel) SetService(v string) {
	o.Service = &v
}

// GetShowRemoteServices returns the ShowRemoteServices field value if set, zero value otherwise.
func (o *ServiceMapPanel) GetShowRemoteServices() bool {
	if o == nil || IsNil(o.ShowRemoteServices) {
		var ret bool
		return ret
	}
	return *o.ShowRemoteServices
}

// GetShowRemoteServicesOk returns a tuple with the ShowRemoteServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapPanel) GetShowRemoteServicesOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowRemoteServices) {
		return nil, false
	}
	return o.ShowRemoteServices, true
}

// HasShowRemoteServices returns a boolean if a field has been set.
func (o *ServiceMapPanel) HasShowRemoteServices() bool {
	if o != nil && !IsNil(o.ShowRemoteServices) {
		return true
	}

	return false
}

// SetShowRemoteServices gets a reference to the given bool and assigns it to the ShowRemoteServices field.
func (o *ServiceMapPanel) SetShowRemoteServices(v bool) {
	o.ShowRemoteServices = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ServiceMapPanel) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapPanel) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ServiceMapPanel) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ServiceMapPanel) SetEnvironment(v string) {
	o.Environment = &v
}

func (o ServiceMapPanel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceMapPanel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPanel, errPanel := json.Marshal(o.Panel)
	if errPanel != nil {
		return map[string]interface{}{}, errPanel
	}
	errPanel = json.Unmarshal([]byte(serializedPanel), &toSerialize)
	if errPanel != nil {
		return map[string]interface{}{}, errPanel
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.ShowRemoteServices) {
		toSerialize["showRemoteServices"] = o.ShowRemoteServices
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

func (o *ServiceMapPanel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"panelType",
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

	varServiceMapPanel := _ServiceMapPanel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceMapPanel)

	if err != nil {
		return err
	}

	*o = ServiceMapPanel(varServiceMapPanel)

	return err
}

type NullableServiceMapPanel struct {
	value *ServiceMapPanel
	isSet bool
}

func (v NullableServiceMapPanel) Get() *ServiceMapPanel {
	return v.value
}

func (v *NullableServiceMapPanel) Set(val *ServiceMapPanel) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceMapPanel) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceMapPanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceMapPanel(val *ServiceMapPanel) *NullableServiceMapPanel {
	return &NullableServiceMapPanel{value: val, isSet: true}
}

func (v NullableServiceMapPanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceMapPanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


