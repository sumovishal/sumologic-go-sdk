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

// checks if the SlosLibrarySloExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlosLibrarySloExport{}

// SlosLibrarySloExport struct for SlosLibrarySloExport
type SlosLibrarySloExport struct {
	SlosLibraryBaseExport
	// Type of SLI Signal (latency, error, throughput, availability or other).
	SignalType string `json:"signalType"`
	Compliance Compliance `json:"compliance"`
	Indicator Sli `json:"indicator"`
	// Name of the service.
	Service *string `json:"service,omitempty"`
	// Name of the application.
	Application *string `json:"application,omitempty"`
	// Tags to be associated with the SLO.
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewSlosLibrarySloExport instantiates a new SlosLibrarySloExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibrarySloExport(signalType string, compliance Compliance, indicator Sli, name string, type_ string) *SlosLibrarySloExport {
	this := SlosLibrarySloExport{}
	this.Name = name
	this.Type = type_
	this.SignalType = signalType
	this.Compliance = compliance
	this.Indicator = indicator
	return &this
}

// NewSlosLibrarySloExportWithDefaults instantiates a new SlosLibrarySloExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibrarySloExportWithDefaults() *SlosLibrarySloExport {
	this := SlosLibrarySloExport{}
	return &this
}

// GetSignalType returns the SignalType field value
func (o *SlosLibrarySloExport) GetSignalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloExport) GetSignalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalType, true
}

// SetSignalType sets field value
func (o *SlosLibrarySloExport) SetSignalType(v string) {
	o.SignalType = v
}

// GetCompliance returns the Compliance field value
func (o *SlosLibrarySloExport) GetCompliance() Compliance {
	if o == nil {
		var ret Compliance
		return ret
	}

	return o.Compliance
}

// GetComplianceOk returns a tuple with the Compliance field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloExport) GetComplianceOk() (*Compliance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compliance, true
}

// SetCompliance sets field value
func (o *SlosLibrarySloExport) SetCompliance(v Compliance) {
	o.Compliance = v
}

// GetIndicator returns the Indicator field value
func (o *SlosLibrarySloExport) GetIndicator() Sli {
	if o == nil {
		var ret Sli
		return ret
	}

	return o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloExport) GetIndicatorOk() (*Sli, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indicator, true
}

// SetIndicator sets field value
func (o *SlosLibrarySloExport) SetIndicator(v Sli) {
	o.Indicator = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SlosLibrarySloExport) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloExport) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SlosLibrarySloExport) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SlosLibrarySloExport) SetService(v string) {
	o.Service = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *SlosLibrarySloExport) GetApplication() string {
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloExport) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *SlosLibrarySloExport) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *SlosLibrarySloExport) SetApplication(v string) {
	o.Application = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SlosLibrarySloExport) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloExport) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SlosLibrarySloExport) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SlosLibrarySloExport) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o SlosLibrarySloExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlosLibrarySloExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSlosLibraryBaseExport, errSlosLibraryBaseExport := json.Marshal(o.SlosLibraryBaseExport)
	if errSlosLibraryBaseExport != nil {
		return map[string]interface{}{}, errSlosLibraryBaseExport
	}
	errSlosLibraryBaseExport = json.Unmarshal([]byte(serializedSlosLibraryBaseExport), &toSerialize)
	if errSlosLibraryBaseExport != nil {
		return map[string]interface{}{}, errSlosLibraryBaseExport
	}
	toSerialize["signalType"] = o.SignalType
	toSerialize["compliance"] = o.Compliance
	toSerialize["indicator"] = o.Indicator
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableSlosLibrarySloExport struct {
	value *SlosLibrarySloExport
	isSet bool
}

func (v NullableSlosLibrarySloExport) Get() *SlosLibrarySloExport {
	return v.value
}

func (v *NullableSlosLibrarySloExport) Set(val *SlosLibrarySloExport) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibrarySloExport) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibrarySloExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibrarySloExport(val *SlosLibrarySloExport) *NullableSlosLibrarySloExport {
	return &NullableSlosLibrarySloExport{value: val, isSet: true}
}

func (v NullableSlosLibrarySloExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibrarySloExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


