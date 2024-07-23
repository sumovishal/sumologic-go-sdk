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

// checks if the ConfigureSubdomainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigureSubdomainRequest{}

// ConfigureSubdomainRequest struct for ConfigureSubdomainRequest
type ConfigureSubdomainRequest struct {
	// The new subdomain.
	Subdomain string `json:"subdomain" validate:"regexp=^(?!xn--)[a-z0-9]([a-z0-9-]*[a-z0-9])?$"`
}

type _ConfigureSubdomainRequest ConfigureSubdomainRequest

// NewConfigureSubdomainRequest instantiates a new ConfigureSubdomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigureSubdomainRequest(subdomain string) *ConfigureSubdomainRequest {
	this := ConfigureSubdomainRequest{}
	this.Subdomain = subdomain
	return &this
}

// NewConfigureSubdomainRequestWithDefaults instantiates a new ConfigureSubdomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigureSubdomainRequestWithDefaults() *ConfigureSubdomainRequest {
	this := ConfigureSubdomainRequest{}
	return &this
}

// GetSubdomain returns the Subdomain field value
func (o *ConfigureSubdomainRequest) GetSubdomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value
// and a boolean to check if the value has been set.
func (o *ConfigureSubdomainRequest) GetSubdomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subdomain, true
}

// SetSubdomain sets field value
func (o *ConfigureSubdomainRequest) SetSubdomain(v string) {
	o.Subdomain = v
}

func (o ConfigureSubdomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigureSubdomainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subdomain"] = o.Subdomain
	return toSerialize, nil
}

func (o *ConfigureSubdomainRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subdomain",
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

	varConfigureSubdomainRequest := _ConfigureSubdomainRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigureSubdomainRequest)

	if err != nil {
		return err
	}

	*o = ConfigureSubdomainRequest(varConfigureSubdomainRequest)

	return err
}

type NullableConfigureSubdomainRequest struct {
	value *ConfigureSubdomainRequest
	isSet bool
}

func (v NullableConfigureSubdomainRequest) Get() *ConfigureSubdomainRequest {
	return v.value
}

func (v *NullableConfigureSubdomainRequest) Set(val *ConfigureSubdomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigureSubdomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigureSubdomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigureSubdomainRequest(val *ConfigureSubdomainRequest) *NullableConfigureSubdomainRequest {
	return &NullableConfigureSubdomainRequest{value: val, isSet: true}
}

func (v NullableConfigureSubdomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigureSubdomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


