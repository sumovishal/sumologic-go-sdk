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

// checks if the BeginAsyncJobResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeginAsyncJobResponseV2{}

// BeginAsyncJobResponseV2 struct for BeginAsyncJobResponseV2
type BeginAsyncJobResponseV2 struct {
	// Identifier of the asynchronous job. Use it to get status of the job.
	JobId string `json:"jobId"`
}

type _BeginAsyncJobResponseV2 BeginAsyncJobResponseV2

// NewBeginAsyncJobResponseV2 instantiates a new BeginAsyncJobResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeginAsyncJobResponseV2(jobId string) *BeginAsyncJobResponseV2 {
	this := BeginAsyncJobResponseV2{}
	this.JobId = jobId
	return &this
}

// NewBeginAsyncJobResponseV2WithDefaults instantiates a new BeginAsyncJobResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeginAsyncJobResponseV2WithDefaults() *BeginAsyncJobResponseV2 {
	this := BeginAsyncJobResponseV2{}
	return &this
}

// GetJobId returns the JobId field value
func (o *BeginAsyncJobResponseV2) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *BeginAsyncJobResponseV2) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *BeginAsyncJobResponseV2) SetJobId(v string) {
	o.JobId = v
}

func (o BeginAsyncJobResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeginAsyncJobResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobId"] = o.JobId
	return toSerialize, nil
}

func (o *BeginAsyncJobResponseV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobId",
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

	varBeginAsyncJobResponseV2 := _BeginAsyncJobResponseV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBeginAsyncJobResponseV2)

	if err != nil {
		return err
	}

	*o = BeginAsyncJobResponseV2(varBeginAsyncJobResponseV2)

	return err
}

type NullableBeginAsyncJobResponseV2 struct {
	value *BeginAsyncJobResponseV2
	isSet bool
}

func (v NullableBeginAsyncJobResponseV2) Get() *BeginAsyncJobResponseV2 {
	return v.value
}

func (v *NullableBeginAsyncJobResponseV2) Set(val *BeginAsyncJobResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableBeginAsyncJobResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableBeginAsyncJobResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeginAsyncJobResponseV2(val *BeginAsyncJobResponseV2) *NullableBeginAsyncJobResponseV2 {
	return &NullableBeginAsyncJobResponseV2{value: val, isSet: true}
}

func (v NullableBeginAsyncJobResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeginAsyncJobResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


