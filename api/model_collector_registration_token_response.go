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
	"time"
)

// checks if the CollectorRegistrationTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectorRegistrationTokenResponse{}

// CollectorRegistrationTokenResponse struct for CollectorRegistrationTokenResponse
type CollectorRegistrationTokenResponse struct {
	TokenBaseResponse
	// The token and URL used to register the Collector as an encoded string.
	EncodedTokenAndUrl string `json:"encodedTokenAndUrl"`
}

type _CollectorRegistrationTokenResponse CollectorRegistrationTokenResponse

// NewCollectorRegistrationTokenResponse instantiates a new CollectorRegistrationTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectorRegistrationTokenResponse(encodedTokenAndUrl string, id string, name string, description string, status string, type_ string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *CollectorRegistrationTokenResponse {
	this := CollectorRegistrationTokenResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Status = status
	this.Type = type_
	this.Version = version
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.EncodedTokenAndUrl = encodedTokenAndUrl
	return &this
}

// NewCollectorRegistrationTokenResponseWithDefaults instantiates a new CollectorRegistrationTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectorRegistrationTokenResponseWithDefaults() *CollectorRegistrationTokenResponse {
	this := CollectorRegistrationTokenResponse{}
	return &this
}

// GetEncodedTokenAndUrl returns the EncodedTokenAndUrl field value
func (o *CollectorRegistrationTokenResponse) GetEncodedTokenAndUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncodedTokenAndUrl
}

// GetEncodedTokenAndUrlOk returns a tuple with the EncodedTokenAndUrl field value
// and a boolean to check if the value has been set.
func (o *CollectorRegistrationTokenResponse) GetEncodedTokenAndUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncodedTokenAndUrl, true
}

// SetEncodedTokenAndUrl sets field value
func (o *CollectorRegistrationTokenResponse) SetEncodedTokenAndUrl(v string) {
	o.EncodedTokenAndUrl = v
}

func (o CollectorRegistrationTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectorRegistrationTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTokenBaseResponse, errTokenBaseResponse := json.Marshal(o.TokenBaseResponse)
	if errTokenBaseResponse != nil {
		return map[string]interface{}{}, errTokenBaseResponse
	}
	errTokenBaseResponse = json.Unmarshal([]byte(serializedTokenBaseResponse), &toSerialize)
	if errTokenBaseResponse != nil {
		return map[string]interface{}{}, errTokenBaseResponse
	}
	toSerialize["encodedTokenAndUrl"] = o.EncodedTokenAndUrl
	return toSerialize, nil
}

func (o *CollectorRegistrationTokenResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"encodedTokenAndUrl",
		"id",
		"name",
		"description",
		"status",
		"type",
		"version",
		"createdAt",
		"createdBy",
		"modifiedAt",
		"modifiedBy",
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

	varCollectorRegistrationTokenResponse := _CollectorRegistrationTokenResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectorRegistrationTokenResponse)

	if err != nil {
		return err
	}

	*o = CollectorRegistrationTokenResponse(varCollectorRegistrationTokenResponse)

	return err
}

type NullableCollectorRegistrationTokenResponse struct {
	value *CollectorRegistrationTokenResponse
	isSet bool
}

func (v NullableCollectorRegistrationTokenResponse) Get() *CollectorRegistrationTokenResponse {
	return v.value
}

func (v *NullableCollectorRegistrationTokenResponse) Set(val *CollectorRegistrationTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectorRegistrationTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectorRegistrationTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectorRegistrationTokenResponse(val *CollectorRegistrationTokenResponse) *NullableCollectorRegistrationTokenResponse {
	return &NullableCollectorRegistrationTokenResponse{value: val, isSet: true}
}

func (v NullableCollectorRegistrationTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectorRegistrationTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


