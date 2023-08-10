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

// checks if the SlosLibrarySloUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlosLibrarySloUpdate{}

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
	// Tags to be associated with the SLO.
	Tags *map[string]string `json:"tags,omitempty"`
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
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloUpdate) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SlosLibrarySloUpdate) HasService() bool {
	if o != nil && !IsNil(o.Service) {
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
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloUpdate) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *SlosLibrarySloUpdate) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *SlosLibrarySloUpdate) SetApplication(v string) {
	o.Application = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SlosLibrarySloUpdate) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloUpdate) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SlosLibrarySloUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SlosLibrarySloUpdate) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o SlosLibrarySloUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlosLibrarySloUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSlosLibraryBaseUpdate, errSlosLibraryBaseUpdate := json.Marshal(o.SlosLibraryBaseUpdate)
	if errSlosLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errSlosLibraryBaseUpdate
	}
	errSlosLibraryBaseUpdate = json.Unmarshal([]byte(serializedSlosLibraryBaseUpdate), &toSerialize)
	if errSlosLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errSlosLibraryBaseUpdate
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


