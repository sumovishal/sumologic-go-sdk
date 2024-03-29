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

// checks if the WarningDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarningDetails{}

// WarningDetails Warning while computing signals.
type WarningDetails struct {
	// Warning code.
	Code string `json:"code"`
	// Warning message.
	Message string `json:"message"`
	// Details related to warning.
	Detail string `json:"detail"`
}

// NewWarningDetails instantiates a new WarningDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarningDetails(code string, message string, detail string) *WarningDetails {
	this := WarningDetails{}
	this.Code = code
	this.Message = message
	this.Detail = detail
	return &this
}

// NewWarningDetailsWithDefaults instantiates a new WarningDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningDetailsWithDefaults() *WarningDetails {
	this := WarningDetails{}
	return &this
}

// GetCode returns the Code field value
func (o *WarningDetails) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *WarningDetails) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *WarningDetails) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *WarningDetails) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *WarningDetails) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *WarningDetails) SetMessage(v string) {
	o.Message = v
}

// GetDetail returns the Detail field value
func (o *WarningDetails) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *WarningDetails) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *WarningDetails) SetDetail(v string) {
	o.Detail = v
}

func (o WarningDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarningDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["detail"] = o.Detail
	return toSerialize, nil
}

type NullableWarningDetails struct {
	value *WarningDetails
	isSet bool
}

func (v NullableWarningDetails) Get() *WarningDetails {
	return v.value
}

func (v *NullableWarningDetails) Set(val *WarningDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableWarningDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableWarningDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarningDetails(val *WarningDetails) *NullableWarningDetails {
	return &NullableWarningDetails{value: val, isSet: true}
}

func (v NullableWarningDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarningDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


