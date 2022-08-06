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

// SlosLibrarySloUpdate struct for SlosLibrarySloUpdate
type SlosLibrarySloUpdate struct {
	SlosLibraryBaseUpdate
	// Type of SLI Signal (latency, error, throughput, availability or other).
	SignalType string `json:"signalType"`
	Compliance Compliance `json:"compliance"`
	Indicator Sli `json:"indicator"`
	// Name of the service.
	Service *string `json:"service,omitempty"`
	// Name of the application.
	Application *string `json:"application,omitempty"`
}

// NewSlosLibrarySloUpdate instantiates a new SlosLibrarySloUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibrarySloUpdate(signalType string, compliance Compliance, indicator Sli, name string, version int64, type_ string) *SlosLibrarySloUpdate {
	this := SlosLibrarySloUpdate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Version = version
	this.Type = type_
	this.SignalType = signalType
	this.Compliance = compliance
	this.Indicator = indicator
	return &this
}

// NewSlosLibrarySloUpdateWithDefaults instantiates a new SlosLibrarySloUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibrarySloUpdateWithDefaults() *SlosLibrarySloUpdate {
	this := SlosLibrarySloUpdate{}
	return &this
}

// GetSignalType returns the SignalType field value
func (o *SlosLibrarySloUpdate) GetSignalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloUpdate) GetSignalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalType, true
}

// SetSignalType sets field value
func (o *SlosLibrarySloUpdate) SetSignalType(v string) {
	o.SignalType = v
}

// GetCompliance returns the Compliance field value
func (o *SlosLibrarySloUpdate) GetCompliance() Compliance {
	if o == nil {
		var ret Compliance
		return ret
	}

	return o.Compliance
}

// GetComplianceOk returns a tuple with the Compliance field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloUpdate) GetComplianceOk() (*Compliance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compliance, true
}

// SetCompliance sets field value
func (o *SlosLibrarySloUpdate) SetCompliance(v Compliance) {
	o.Compliance = v
}

// GetIndicator returns the Indicator field value
func (o *SlosLibrarySloUpdate) GetIndicator() Sli {
	if o == nil {
		var ret Sli
		return ret
	}

	return o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloUpdate) GetIndicatorOk() (*Sli, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indicator, true
}

// SetIndicator sets field value
func (o *SlosLibrarySloUpdate) SetIndicator(v Sli) {
	o.Indicator = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SlosLibrarySloUpdate) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloUpdate) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SlosLibrarySloUpdate) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SlosLibrarySloUpdate) SetService(v string) {
	o.Service = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *SlosLibrarySloUpdate) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloUpdate) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *SlosLibrarySloUpdate) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *SlosLibrarySloUpdate) SetApplication(v string) {
	o.Application = &v
}

func (o SlosLibrarySloUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSlosLibraryBaseUpdate, errSlosLibraryBaseUpdate := json.Marshal(o.SlosLibraryBaseUpdate)
	if errSlosLibraryBaseUpdate != nil {
		return []byte{}, errSlosLibraryBaseUpdate
	}
	errSlosLibraryBaseUpdate = json.Unmarshal([]byte(serializedSlosLibraryBaseUpdate), &toSerialize)
	if errSlosLibraryBaseUpdate != nil {
		return []byte{}, errSlosLibraryBaseUpdate
	}
	if true {
		toSerialize["signalType"] = o.SignalType
	}
	if true {
		toSerialize["compliance"] = o.Compliance
	}
	if true {
		toSerialize["indicator"] = o.Indicator
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	return json.Marshal(toSerialize)
}

type NullableSlosLibrarySloUpdate struct {
	value *SlosLibrarySloUpdate
	isSet bool
}

func (v NullableSlosLibrarySloUpdate) Get() *SlosLibrarySloUpdate {
	return v.value
}

func (v *NullableSlosLibrarySloUpdate) Set(val *SlosLibrarySloUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibrarySloUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibrarySloUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibrarySloUpdate(val *SlosLibrarySloUpdate) *NullableSlosLibrarySloUpdate {
	return &NullableSlosLibrarySloUpdate{value: val, isSet: true}
}

func (v NullableSlosLibrarySloUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibrarySloUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


