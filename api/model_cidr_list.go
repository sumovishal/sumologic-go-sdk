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

// checks if the CidrList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CidrList{}

// CidrList A list of CIDR notations and/or IP addresses.
type CidrList struct {
	// An array of CIDR notations and/or IP addresses.
	Data []Cidr `json:"data"`
}

// NewCidrList instantiates a new CidrList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCidrList(data []Cidr) *CidrList {
	this := CidrList{}
	this.Data = data
	return &this
}

// NewCidrListWithDefaults instantiates a new CidrList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCidrListWithDefaults() *CidrList {
	this := CidrList{}
	return &this
}

// GetData returns the Data field value
func (o *CidrList) GetData() []Cidr {
	if o == nil {
		var ret []Cidr
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CidrList) GetDataOk() ([]Cidr, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CidrList) SetData(v []Cidr) {
	o.Data = v
}

func (o CidrList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CidrList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCidrList struct {
	value *CidrList
	isSet bool
}

func (v NullableCidrList) Get() *CidrList {
	return v.value
}

func (v *NullableCidrList) Set(val *CidrList) {
	v.value = val
	v.isSet = true
}

func (v NullableCidrList) IsSet() bool {
	return v.isSet
}

func (v *NullableCidrList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCidrList(val *CidrList) *NullableCidrList {
	return &NullableCidrList{value: val, isSet: true}
}

func (v NullableCidrList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCidrList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


