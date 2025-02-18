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

// checks if the BulkBeginAsyncJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkBeginAsyncJobResponse{}

// BulkBeginAsyncJobResponse struct for BulkBeginAsyncJobResponse
type BulkBeginAsyncJobResponse struct {
	// Map of content identifiers to job identifiers.
	JobIds map[string]string `json:"jobIds"`
	// Map of content identifiers to error messages for all failed job requests
	Errors map[string]BulkErrorResponse `json:"errors"`
}

type _BulkBeginAsyncJobResponse BulkBeginAsyncJobResponse

// NewBulkBeginAsyncJobResponse instantiates a new BulkBeginAsyncJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkBeginAsyncJobResponse(jobIds map[string]string, errors map[string]BulkErrorResponse) *BulkBeginAsyncJobResponse {
	this := BulkBeginAsyncJobResponse{}
	this.JobIds = jobIds
	this.Errors = errors
	return &this
}

// NewBulkBeginAsyncJobResponseWithDefaults instantiates a new BulkBeginAsyncJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkBeginAsyncJobResponseWithDefaults() *BulkBeginAsyncJobResponse {
	this := BulkBeginAsyncJobResponse{}
	return &this
}

// GetJobIds returns the JobIds field value
func (o *BulkBeginAsyncJobResponse) GetJobIds() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.JobIds
}

// GetJobIdsOk returns a tuple with the JobIds field value
// and a boolean to check if the value has been set.
func (o *BulkBeginAsyncJobResponse) GetJobIdsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobIds, true
}

// SetJobIds sets field value
func (o *BulkBeginAsyncJobResponse) SetJobIds(v map[string]string) {
	o.JobIds = v
}

// GetErrors returns the Errors field value
func (o *BulkBeginAsyncJobResponse) GetErrors() map[string]BulkErrorResponse {
	if o == nil {
		var ret map[string]BulkErrorResponse
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *BulkBeginAsyncJobResponse) GetErrorsOk() (*map[string]BulkErrorResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *BulkBeginAsyncJobResponse) SetErrors(v map[string]BulkErrorResponse) {
	o.Errors = v
}

func (o BulkBeginAsyncJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkBeginAsyncJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobIds"] = o.JobIds
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

func (o *BulkBeginAsyncJobResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobIds",
		"errors",
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

	varBulkBeginAsyncJobResponse := _BulkBeginAsyncJobResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkBeginAsyncJobResponse)

	if err != nil {
		return err
	}

	*o = BulkBeginAsyncJobResponse(varBulkBeginAsyncJobResponse)

	return err
}

type NullableBulkBeginAsyncJobResponse struct {
	value *BulkBeginAsyncJobResponse
	isSet bool
}

func (v NullableBulkBeginAsyncJobResponse) Get() *BulkBeginAsyncJobResponse {
	return v.value
}

func (v *NullableBulkBeginAsyncJobResponse) Set(val *BulkBeginAsyncJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkBeginAsyncJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkBeginAsyncJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkBeginAsyncJobResponse(val *BulkBeginAsyncJobResponse) *NullableBulkBeginAsyncJobResponse {
	return &NullableBulkBeginAsyncJobResponse{value: val, isSet: true}
}

func (v NullableBulkBeginAsyncJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkBeginAsyncJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


