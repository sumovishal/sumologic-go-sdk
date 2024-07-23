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

// checks if the FieldQuotaUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldQuotaUsage{}

// FieldQuotaUsage struct for FieldQuotaUsage
type FieldQuotaUsage struct {
	// Maximum number of fields available.
	Quota int32 `json:"quota"`
	// Current number of fields available.
	Remaining int32 `json:"remaining"`
}

type _FieldQuotaUsage FieldQuotaUsage

// NewFieldQuotaUsage instantiates a new FieldQuotaUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldQuotaUsage(quota int32, remaining int32) *FieldQuotaUsage {
	this := FieldQuotaUsage{}
	this.Quota = quota
	this.Remaining = remaining
	return &this
}

// NewFieldQuotaUsageWithDefaults instantiates a new FieldQuotaUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldQuotaUsageWithDefaults() *FieldQuotaUsage {
	this := FieldQuotaUsage{}
	return &this
}

// GetQuota returns the Quota field value
func (o *FieldQuotaUsage) GetQuota() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value
// and a boolean to check if the value has been set.
func (o *FieldQuotaUsage) GetQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quota, true
}

// SetQuota sets field value
func (o *FieldQuotaUsage) SetQuota(v int32) {
	o.Quota = v
}

// GetRemaining returns the Remaining field value
func (o *FieldQuotaUsage) GetRemaining() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *FieldQuotaUsage) GetRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *FieldQuotaUsage) SetRemaining(v int32) {
	o.Remaining = v
}

func (o FieldQuotaUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldQuotaUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quota"] = o.Quota
	toSerialize["remaining"] = o.Remaining
	return toSerialize, nil
}

func (o *FieldQuotaUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quota",
		"remaining",
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

	varFieldQuotaUsage := _FieldQuotaUsage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFieldQuotaUsage)

	if err != nil {
		return err
	}

	*o = FieldQuotaUsage(varFieldQuotaUsage)

	return err
}

type NullableFieldQuotaUsage struct {
	value *FieldQuotaUsage
	isSet bool
}

func (v NullableFieldQuotaUsage) Get() *FieldQuotaUsage {
	return v.value
}

func (v *NullableFieldQuotaUsage) Set(val *FieldQuotaUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldQuotaUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldQuotaUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldQuotaUsage(val *FieldQuotaUsage) *NullableFieldQuotaUsage {
	return &NullableFieldQuotaUsage{value: val, isSet: true}
}

func (v NullableFieldQuotaUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldQuotaUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


