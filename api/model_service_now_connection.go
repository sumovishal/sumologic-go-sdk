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

// checks if the ServiceNowConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceNowConnection{}

// ServiceNowConnection struct for ServiceNowConnection
type ServiceNowConnection struct {
	Connection
	// URL for the ServiceNow connection.
	Url string `json:"url"`
	// User name for the ServiceNow connection.
	Username string `json:"username"`
}

type _ServiceNowConnection ServiceNowConnection

// NewServiceNowConnection instantiates a new ServiceNowConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNowConnection(url string, username string, type_ string, id string, name string, description string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *ServiceNowConnection {
	this := ServiceNowConnection{}
	this.Type = type_
	this.Id = id
	this.Name = name
	this.Description = description
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Url = url
	this.Username = username
	return &this
}

// NewServiceNowConnectionWithDefaults instantiates a new ServiceNowConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNowConnectionWithDefaults() *ServiceNowConnection {
	this := ServiceNowConnection{}
	return &this
}

// GetUrl returns the Url field value
func (o *ServiceNowConnection) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceNowConnection) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ServiceNowConnection) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *ServiceNowConnection) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ServiceNowConnection) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ServiceNowConnection) SetUsername(v string) {
	o.Username = v
}

func (o ServiceNowConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceNowConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedConnection, errConnection := json.Marshal(o.Connection)
	if errConnection != nil {
		return map[string]interface{}{}, errConnection
	}
	errConnection = json.Unmarshal([]byte(serializedConnection), &toSerialize)
	if errConnection != nil {
		return map[string]interface{}{}, errConnection
	}
	toSerialize["url"] = o.Url
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *ServiceNowConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"username",
		"type",
		"id",
		"name",
		"description",
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

	varServiceNowConnection := _ServiceNowConnection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceNowConnection)

	if err != nil {
		return err
	}

	*o = ServiceNowConnection(varServiceNowConnection)

	return err
}

type NullableServiceNowConnection struct {
	value *ServiceNowConnection
	isSet bool
}

func (v NullableServiceNowConnection) Get() *ServiceNowConnection {
	return v.value
}

func (v *NullableServiceNowConnection) Set(val *ServiceNowConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNowConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNowConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNowConnection(val *ServiceNowConnection) *NullableServiceNowConnection {
	return &NullableServiceNowConnection{value: val, isSet: true}
}

func (v NullableServiceNowConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNowConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


