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

// checks if the SlosLibrarySloResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlosLibrarySloResponse{}

// SlosLibrarySloResponse struct for SlosLibrarySloResponse
type SlosLibrarySloResponse struct {
	SlosLibraryBaseResponse
	// Type of SLI Signal (latency, error, throughput, availability or other).
	SignalType string `json:"signalType" validate:"regexp=^(Latency|Error|Throughput|Availability|Other)$"`
	Compliance Compliance `json:"compliance"`
	Indicator Sli `json:"indicator"`
	// Name of the service.
	Service *string `json:"service,omitempty"`
	// Name of the application.
	Application *string `json:"application,omitempty"`
	// Tags to be associated with the SLO.
	Tags *map[string]string `json:"tags,omitempty"`
	// Current SLO Version. This is incremented on every change of a critical field of the SLO (i.e, SLI or Compliance period timezone), that requires a recompute of the SLI values over the compliance period.
	SloVersion *int64 `json:"sloVersion,omitempty"`
}

type _SlosLibrarySloResponse SlosLibrarySloResponse

// NewSlosLibrarySloResponse instantiates a new SlosLibrarySloResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlosLibrarySloResponse(signalType string, compliance Compliance, indicator Sli, id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool) *SlosLibrarySloResponse {
	this := SlosLibrarySloResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Version = version
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.ParentId = parentId
	this.ContentType = contentType
	this.Type = type_
	this.IsSystem = isSystem
	this.IsMutable = isMutable
	this.SignalType = signalType
	this.Compliance = compliance
	this.Indicator = indicator
	return &this
}

// NewSlosLibrarySloResponseWithDefaults instantiates a new SlosLibrarySloResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlosLibrarySloResponseWithDefaults() *SlosLibrarySloResponse {
	this := SlosLibrarySloResponse{}
	return &this
}

// GetSignalType returns the SignalType field value
func (o *SlosLibrarySloResponse) GetSignalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloResponse) GetSignalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalType, true
}

// SetSignalType sets field value
func (o *SlosLibrarySloResponse) SetSignalType(v string) {
	o.SignalType = v
}

// GetCompliance returns the Compliance field value
func (o *SlosLibrarySloResponse) GetCompliance() Compliance {
	if o == nil {
		var ret Compliance
		return ret
	}

	return o.Compliance
}

// GetComplianceOk returns a tuple with the Compliance field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloResponse) GetComplianceOk() (*Compliance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compliance, true
}

// SetCompliance sets field value
func (o *SlosLibrarySloResponse) SetCompliance(v Compliance) {
	o.Compliance = v
}

// GetIndicator returns the Indicator field value
func (o *SlosLibrarySloResponse) GetIndicator() Sli {
	if o == nil {
		var ret Sli
		return ret
	}

	return o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloResponse) GetIndicatorOk() (*Sli, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indicator, true
}

// SetIndicator sets field value
func (o *SlosLibrarySloResponse) SetIndicator(v Sli) {
	o.Indicator = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SlosLibrarySloResponse) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloResponse) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SlosLibrarySloResponse) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SlosLibrarySloResponse) SetService(v string) {
	o.Service = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *SlosLibrarySloResponse) GetApplication() string {
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloResponse) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *SlosLibrarySloResponse) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *SlosLibrarySloResponse) SetApplication(v string) {
	o.Application = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SlosLibrarySloResponse) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloResponse) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SlosLibrarySloResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *SlosLibrarySloResponse) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetSloVersion returns the SloVersion field value if set, zero value otherwise.
func (o *SlosLibrarySloResponse) GetSloVersion() int64 {
	if o == nil || IsNil(o.SloVersion) {
		var ret int64
		return ret
	}
	return *o.SloVersion
}

// GetSloVersionOk returns a tuple with the SloVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlosLibrarySloResponse) GetSloVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.SloVersion) {
		return nil, false
	}
	return o.SloVersion, true
}

// HasSloVersion returns a boolean if a field has been set.
func (o *SlosLibrarySloResponse) HasSloVersion() bool {
	if o != nil && !IsNil(o.SloVersion) {
		return true
	}

	return false
}

// SetSloVersion gets a reference to the given int64 and assigns it to the SloVersion field.
func (o *SlosLibrarySloResponse) SetSloVersion(v int64) {
	o.SloVersion = &v
}

func (o SlosLibrarySloResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlosLibrarySloResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSlosLibraryBaseResponse, errSlosLibraryBaseResponse := json.Marshal(o.SlosLibraryBaseResponse)
	if errSlosLibraryBaseResponse != nil {
		return map[string]interface{}{}, errSlosLibraryBaseResponse
	}
	errSlosLibraryBaseResponse = json.Unmarshal([]byte(serializedSlosLibraryBaseResponse), &toSerialize)
	if errSlosLibraryBaseResponse != nil {
		return map[string]interface{}{}, errSlosLibraryBaseResponse
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
	if !IsNil(o.SloVersion) {
		toSerialize["sloVersion"] = o.SloVersion
	}
	return toSerialize, nil
}

func (o *SlosLibrarySloResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signalType",
		"compliance",
		"indicator",
		"id",
		"name",
		"description",
		"version",
		"createdAt",
		"createdBy",
		"modifiedAt",
		"modifiedBy",
		"parentId",
		"contentType",
		"type",
		"isSystem",
		"isMutable",
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

	varSlosLibrarySloResponse := _SlosLibrarySloResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSlosLibrarySloResponse)

	if err != nil {
		return err
	}

	*o = SlosLibrarySloResponse(varSlosLibrarySloResponse)

	return err
}

type NullableSlosLibrarySloResponse struct {
	value *SlosLibrarySloResponse
	isSet bool
}

func (v NullableSlosLibrarySloResponse) Get() *SlosLibrarySloResponse {
	return v.value
}

func (v *NullableSlosLibrarySloResponse) Set(val *SlosLibrarySloResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSlosLibrarySloResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSlosLibrarySloResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlosLibrarySloResponse(val *SlosLibrarySloResponse) *NullableSlosLibrarySloResponse {
	return &NullableSlosLibrarySloResponse{value: val, isSet: true}
}

func (v NullableSlosLibrarySloResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlosLibrarySloResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


