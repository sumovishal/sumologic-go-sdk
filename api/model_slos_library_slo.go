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

// SlosLibrarySlo struct for SlosLibrarySlo
type SlosLibrarySlo struct {
	SlosLibraryBase
	// Type of SLI Signal (latency, error, throughput, availability or other).
	SignalType string `json:"signalType"`
	Compliance Compliance `json:"compliance"`
	Indicator Sli `json:"indicator"`
	// Name of the service.
	Service *string `json:"service,omitempty"`
	// Name of the application.
	Application *string `json:"application,omitempty"`
}

// NewSlosLibrarySlo instantiates a new SlosLibrarySlo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibrarySlo(signalType string, compliance Compliance, indicator Sli, name string, type_ string) *SlosLibrarySlo {
	this := SlosLibrarySlo{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Type = type_
	this.SignalType = signalType
	this.Compliance = compliance
	this.Indicator = indicator
	return &this
}

// NewSlosLibrarySloWithDefaults instantiates a new SlosLibrarySlo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibrarySloWithDefaults() *SlosLibrarySlo {
	this := SlosLibrarySlo{}
	return &this
}

// GetSignalType returns the SignalType field value
func (o *SlosLibrarySlo) GetSignalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySlo) GetSignalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalType, true
}

// SetSignalType sets field value
func (o *SlosLibrarySlo) SetSignalType(v string) {
	o.SignalType = v
}

// GetCompliance returns the Compliance field value
func (o *SlosLibrarySlo) GetCompliance() Compliance {
	if o == nil {
		var ret Compliance
		return ret
	}

	return o.Compliance
}

// GetComplianceOk returns a tuple with the Compliance field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySlo) GetComplianceOk() (*Compliance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compliance, true
}

// SetCompliance sets field value
func (o *SlosLibrarySlo) SetCompliance(v Compliance) {
	o.Compliance = v
}

// GetIndicator returns the Indicator field value
func (o *SlosLibrarySlo) GetIndicator() Sli {
	if o == nil {
		var ret Sli
		return ret
	}

	return o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySlo) GetIndicatorOk() (*Sli, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indicator, true
}

// SetIndicator sets field value
func (o *SlosLibrarySlo) SetIndicator(v Sli) {
	o.Indicator = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SlosLibrarySlo) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySlo) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SlosLibrarySlo) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SlosLibrarySlo) SetService(v string) {
	o.Service = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *SlosLibrarySlo) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySlo) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *SlosLibrarySlo) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *SlosLibrarySlo) SetApplication(v string) {
	o.Application = &v
}

func (o SlosLibrarySlo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSlosLibraryBase, errSlosLibraryBase := json.Marshal(o.SlosLibraryBase)
	if errSlosLibraryBase != nil {
		return []byte{}, errSlosLibraryBase
	}
	errSlosLibraryBase = json.Unmarshal([]byte(serializedSlosLibraryBase), &toSerialize)
	if errSlosLibraryBase != nil {
		return []byte{}, errSlosLibraryBase
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

type NullableSlosLibrarySlo struct {
	value *SlosLibrarySlo
	isSet bool
}

func (v NullableSlosLibrarySlo) Get() *SlosLibrarySlo {
	return v.value
}

func (v *NullableSlosLibrarySlo) Set(val *SlosLibrarySlo) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibrarySlo) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibrarySlo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibrarySlo(val *SlosLibrarySlo) *NullableSlosLibrarySlo {
	return &NullableSlosLibrarySlo{value: val, isSet: true}
}

func (v NullableSlosLibrarySlo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibrarySlo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


