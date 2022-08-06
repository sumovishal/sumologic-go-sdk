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

// ServiceMapPanelAllOf A panel for service map.
type ServiceMapPanelAllOf struct {
	// Filter services by the application custom tag.
	Application *string `json:"application,omitempty"`
	// Show only the specific service and its connections to other services.
	Service *string `json:"service,omitempty"`
	// Show remote services, like databases or external calls, automatically detected in client traffic.
	ShowRemoteServices *bool `json:"showRemoteServices,omitempty"`
}

// NewServiceMapPanelAllOf instantiates a new ServiceMapPanelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceMapPanelAllOf() *ServiceMapPanelAllOf {
	this := ServiceMapPanelAllOf{}
	return &this
}

// NewServiceMapPanelAllOfWithDefaults instantiates a new ServiceMapPanelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceMapPanelAllOfWithDefaults() *ServiceMapPanelAllOf {
	this := ServiceMapPanelAllOf{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ServiceMapPanelAllOf) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapPanelAllOf) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ServiceMapPanelAllOf) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ServiceMapPanelAllOf) SetApplication(v string) {
	o.Application = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceMapPanelAllOf) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapPanelAllOf) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceMapPanelAllOf) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ServiceMapPanelAllOf) SetService(v string) {
	o.Service = &v
}

// GetShowRemoteServices returns the ShowRemoteServices field value if set, zero value otherwise.
func (o *ServiceMapPanelAllOf) GetShowRemoteServices() bool {
	if o == nil || o.ShowRemoteServices == nil {
		var ret bool
		return ret
	}
	return *o.ShowRemoteServices
}

// GetShowRemoteServicesOk returns a tuple with the ShowRemoteServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceMapPanelAllOf) GetShowRemoteServicesOk() (*bool, bool) {
	if o == nil || o.ShowRemoteServices == nil {
		return nil, false
	}
	return o.ShowRemoteServices, true
}

// HasShowRemoteServices returns a boolean if a field has been set.
func (o *ServiceMapPanelAllOf) HasShowRemoteServices() bool {
	if o != nil && o.ShowRemoteServices != nil {
		return true
	}

	return false
}

// SetShowRemoteServices gets a reference to the given bool and assigns it to the ShowRemoteServices field.
func (o *ServiceMapPanelAllOf) SetShowRemoteServices(v bool) {
	o.ShowRemoteServices = &v
}

func (o ServiceMapPanelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ShowRemoteServices != nil {
		toSerialize["showRemoteServices"] = o.ShowRemoteServices
	}
	return json.Marshal(toSerialize)
}

type NullableServiceMapPanelAllOf struct {
	value *ServiceMapPanelAllOf
	isSet bool
}

func (v NullableServiceMapPanelAllOf) Get() *ServiceMapPanelAllOf {
	return v.value
}

func (v *NullableServiceMapPanelAllOf) Set(val *ServiceMapPanelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceMapPanelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceMapPanelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceMapPanelAllOf(val *ServiceMapPanelAllOf) *NullableServiceMapPanelAllOf {
	return &NullableServiceMapPanelAllOf{value: val, isSet: true}
}

func (v NullableServiceMapPanelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceMapPanelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


