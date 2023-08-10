/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// checks if the UploadStixIndicatorsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadStixIndicatorsRequest{}

// UploadStixIndicatorsRequest struct for UploadStixIndicatorsRequest
type UploadStixIndicatorsRequest struct {
	// User-provided text to identify the source of the indicator
	Source string `json:"source"`
	// When this indicator was first loaded into Sumo. Timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	Imported time.Time `json:"imported"`
	// The list of stix threat intel indicators to upload.
	Indicators []StixIndicator `json:"indicators"`
}

// NewUploadStixIndicatorsRequest instantiates a new UploadStixIndicatorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadStixIndicatorsRequest(source string, imported time.Time, indicators []StixIndicator) *UploadStixIndicatorsRequest {
	this := UploadStixIndicatorsRequest{}
	this.Source = source
	this.Imported = imported
	this.Indicators = indicators
	return &this
}

// NewUploadStixIndicatorsRequestWithDefaults instantiates a new UploadStixIndicatorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadStixIndicatorsRequestWithDefaults() *UploadStixIndicatorsRequest {
	this := UploadStixIndicatorsRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *UploadStixIndicatorsRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UploadStixIndicatorsRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UploadStixIndicatorsRequest) SetSource(v string) {
	o.Source = v
}

// GetImported returns the Imported field value
func (o *UploadStixIndicatorsRequest) GetImported() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Imported
}

// GetImportedOk returns a tuple with the Imported field value
// and a boolean to check if the value has been set.
func (o *UploadStixIndicatorsRequest) GetImportedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Imported, true
}

// SetImported sets field value
func (o *UploadStixIndicatorsRequest) SetImported(v time.Time) {
	o.Imported = v
}

// GetIndicators returns the Indicators field value
func (o *UploadStixIndicatorsRequest) GetIndicators() []StixIndicator {
	if o == nil {
		var ret []StixIndicator
		return ret
	}

	return o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value
// and a boolean to check if the value has been set.
func (o *UploadStixIndicatorsRequest) GetIndicatorsOk() ([]StixIndicator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indicators, true
}

// SetIndicators sets field value
func (o *UploadStixIndicatorsRequest) SetIndicators(v []StixIndicator) {
	o.Indicators = v
}

func (o UploadStixIndicatorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadStixIndicatorsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["imported"] = o.Imported
	toSerialize["indicators"] = o.Indicators
	return toSerialize, nil
}

type NullableUploadStixIndicatorsRequest struct {
	value *UploadStixIndicatorsRequest
	isSet bool
}

func (v NullableUploadStixIndicatorsRequest) Get() *UploadStixIndicatorsRequest {
	return v.value
}

func (v *NullableUploadStixIndicatorsRequest) Set(val *UploadStixIndicatorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadStixIndicatorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadStixIndicatorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadStixIndicatorsRequest(val *UploadStixIndicatorsRequest) *NullableUploadStixIndicatorsRequest {
	return &NullableUploadStixIndicatorsRequest{value: val, isSet: true}
}

func (v NullableUploadStixIndicatorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadStixIndicatorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


