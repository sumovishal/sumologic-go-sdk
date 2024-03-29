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

// checks if the ServiceNowConnectionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceNowConnectionAllOf{}

// ServiceNowConnectionAllOf struct for ServiceNowConnectionAllOf
type ServiceNowConnectionAllOf struct {
	// URL for the ServiceNow connection.
	Url string `json:"url"`
	// User name for the ServiceNow connection.
	Username string `json:"username"`
}

// NewServiceNowConnectionAllOf instantiates a new ServiceNowConnectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNowConnectionAllOf(url string, username string) *ServiceNowConnectionAllOf {
	this := ServiceNowConnectionAllOf{}
	this.Url = url
	this.Username = username
	return &this
}

// NewServiceNowConnectionAllOfWithDefaults instantiates a new ServiceNowConnectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNowConnectionAllOfWithDefaults() *ServiceNowConnectionAllOf {
	this := ServiceNowConnectionAllOf{}
	return &this
}

// GetUrl returns the Url field value
func (o *ServiceNowConnectionAllOf) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceNowConnectionAllOf) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ServiceNowConnectionAllOf) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *ServiceNowConnectionAllOf) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ServiceNowConnectionAllOf) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ServiceNowConnectionAllOf) SetUsername(v string) {
	o.Username = v
}

func (o ServiceNowConnectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceNowConnectionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableServiceNowConnectionAllOf struct {
	value *ServiceNowConnectionAllOf
	isSet bool
}

func (v NullableServiceNowConnectionAllOf) Get() *ServiceNowConnectionAllOf {
	return v.value
}

func (v *NullableServiceNowConnectionAllOf) Set(val *ServiceNowConnectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNowConnectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNowConnectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNowConnectionAllOf(val *ServiceNowConnectionAllOf) *NullableServiceNowConnectionAllOf {
	return &NullableServiceNowConnectionAllOf{value: val, isSet: true}
}

func (v NullableServiceNowConnectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNowConnectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


