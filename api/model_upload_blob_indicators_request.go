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

// checks if the UploadBlobIndicatorsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadBlobIndicatorsRequest{}

// UploadBlobIndicatorsRequest struct for UploadBlobIndicatorsRequest
type UploadBlobIndicatorsRequest struct {
	// User-provided text to identify the source of the indicator
	Source string `json:"source"`
	// When this indicator was first loaded into Sumo. Timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	Imported time.Time `json:"imported"`
	// The input format.
	Format string `json:"format"`
	// The blob containing indicators.
	Blob string `json:"blob"`
}

// NewUploadBlobIndicatorsRequest instantiates a new UploadBlobIndicatorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadBlobIndicatorsRequest(source string, imported time.Time, format string, blob string) *UploadBlobIndicatorsRequest {
	this := UploadBlobIndicatorsRequest{}
	this.Source = source
	this.Imported = imported
	this.Format = format
	this.Blob = blob
	return &this
}

// NewUploadBlobIndicatorsRequestWithDefaults instantiates a new UploadBlobIndicatorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadBlobIndicatorsRequestWithDefaults() *UploadBlobIndicatorsRequest {
	this := UploadBlobIndicatorsRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *UploadBlobIndicatorsRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UploadBlobIndicatorsRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UploadBlobIndicatorsRequest) SetSource(v string) {
	o.Source = v
}

// GetImported returns the Imported field value
func (o *UploadBlobIndicatorsRequest) GetImported() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Imported
}

// GetImportedOk returns a tuple with the Imported field value
// and a boolean to check if the value has been set.
func (o *UploadBlobIndicatorsRequest) GetImportedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Imported, true
}

// SetImported sets field value
func (o *UploadBlobIndicatorsRequest) SetImported(v time.Time) {
	o.Imported = v
}

// GetFormat returns the Format field value
func (o *UploadBlobIndicatorsRequest) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *UploadBlobIndicatorsRequest) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *UploadBlobIndicatorsRequest) SetFormat(v string) {
	o.Format = v
}

// GetBlob returns the Blob field value
func (o *UploadBlobIndicatorsRequest) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *UploadBlobIndicatorsRequest) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *UploadBlobIndicatorsRequest) SetBlob(v string) {
	o.Blob = v
}

func (o UploadBlobIndicatorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadBlobIndicatorsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["imported"] = o.Imported
	toSerialize["format"] = o.Format
	toSerialize["blob"] = o.Blob
	return toSerialize, nil
}

type NullableUploadBlobIndicatorsRequest struct {
	value *UploadBlobIndicatorsRequest
	isSet bool
}

func (v NullableUploadBlobIndicatorsRequest) Get() *UploadBlobIndicatorsRequest {
	return v.value
}

func (v *NullableUploadBlobIndicatorsRequest) Set(val *UploadBlobIndicatorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadBlobIndicatorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadBlobIndicatorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadBlobIndicatorsRequest(val *UploadBlobIndicatorsRequest) *NullableUploadBlobIndicatorsRequest {
	return &NullableUploadBlobIndicatorsRequest{value: val, isSet: true}
}

func (v NullableUploadBlobIndicatorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadBlobIndicatorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


