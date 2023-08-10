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

// checks if the ServiceNowDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceNowDefinitionAllOf{}

// ServiceNowDefinitionAllOf struct for ServiceNowDefinitionAllOf
type ServiceNowDefinitionAllOf struct {
	// URL for the ServiceNow connection.
	Url string `json:"url"`
	// User name for the ServiceNow connection.
	Username string `json:"username"`
	// User password for the ServiceNow connection.
	Password string `json:"password"`
}

// NewServiceNowDefinitionAllOf instantiates a new ServiceNowDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNowDefinitionAllOf(url string, username string, password string) *ServiceNowDefinitionAllOf {
	this := ServiceNowDefinitionAllOf{}
	this.Url = url
	this.Username = username
	this.Password = password
	return &this
}

// NewServiceNowDefinitionAllOfWithDefaults instantiates a new ServiceNowDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNowDefinitionAllOfWithDefaults() *ServiceNowDefinitionAllOf {
	this := ServiceNowDefinitionAllOf{}
	return &this
}

// GetUrl returns the Url field value
func (o *ServiceNowDefinitionAllOf) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceNowDefinitionAllOf) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ServiceNowDefinitionAllOf) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *ServiceNowDefinitionAllOf) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ServiceNowDefinitionAllOf) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ServiceNowDefinitionAllOf) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *ServiceNowDefinitionAllOf) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ServiceNowDefinitionAllOf) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ServiceNowDefinitionAllOf) SetPassword(v string) {
	o.Password = v
}

func (o ServiceNowDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceNowDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableServiceNowDefinitionAllOf struct {
	value *ServiceNowDefinitionAllOf
	isSet bool
}

func (v NullableServiceNowDefinitionAllOf) Get() *ServiceNowDefinitionAllOf {
	return v.value
}

func (v *NullableServiceNowDefinitionAllOf) Set(val *ServiceNowDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNowDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNowDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNowDefinitionAllOf(val *ServiceNowDefinitionAllOf) *NullableServiceNowDefinitionAllOf {
	return &NullableServiceNowDefinitionAllOf{value: val, isSet: true}
}

func (v NullableServiceNowDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNowDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


