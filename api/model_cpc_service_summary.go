/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CpcServiceSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpcServiceSummary{}

// CpcServiceSummary struct for CpcServiceSummary
type CpcServiceSummary struct {
	// The name of the service.
	Service string `json:"service"`
	// The color hex code assigned to the service.
	Color string `json:"color"`
	CpcSummary CpcSummary `json:"cpcSummary"`
}

// NewCpcServiceSummary instantiates a new CpcServiceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpcServiceSummary(service string, color string, cpcSummary CpcSummary) *CpcServiceSummary {
	this := CpcServiceSummary{}
	this.Service = service
	this.Color = color
	this.CpcSummary = cpcSummary
	return &this
}

// NewCpcServiceSummaryWithDefaults instantiates a new CpcServiceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpcServiceSummaryWithDefaults() *CpcServiceSummary {
	this := CpcServiceSummary{}
	return &this
}

// GetService returns the Service field value
func (o *CpcServiceSummary) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *CpcServiceSummary) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *CpcServiceSummary) SetService(v string) {
	o.Service = v
}

// GetColor returns the Color field value
func (o *CpcServiceSummary) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *CpcServiceSummary) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *CpcServiceSummary) SetColor(v string) {
	o.Color = v
}

// GetCpcSummary returns the CpcSummary field value
func (o *CpcServiceSummary) GetCpcSummary() CpcSummary {
	if o == nil {
		var ret CpcSummary
		return ret
	}

	return o.CpcSummary
}

// GetCpcSummaryOk returns a tuple with the CpcSummary field value
// and a boolean to check if the value has been set.
func (o *CpcServiceSummary) GetCpcSummaryOk() (*CpcSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpcSummary, true
}

// SetCpcSummary sets field value
func (o *CpcServiceSummary) SetCpcSummary(v CpcSummary) {
	o.CpcSummary = v
}

func (o CpcServiceSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpcServiceSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service"] = o.Service
	toSerialize["color"] = o.Color
	toSerialize["cpcSummary"] = o.CpcSummary
	return toSerialize, nil
}

type NullableCpcServiceSummary struct {
	value *CpcServiceSummary
	isSet bool
}

func (v NullableCpcServiceSummary) Get() *CpcServiceSummary {
	return v.value
}

func (v *NullableCpcServiceSummary) Set(val *CpcServiceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCpcServiceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCpcServiceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpcServiceSummary(val *CpcServiceSummary) *NullableCpcServiceSummary {
	return &NullableCpcServiceSummary{value: val, isSet: true}
}

func (v NullableCpcServiceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpcServiceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


