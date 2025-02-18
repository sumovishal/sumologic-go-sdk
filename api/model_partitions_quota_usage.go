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

// checks if the PartitionsQuotaUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartitionsQuotaUsage{}

// PartitionsQuotaUsage struct for PartitionsQuotaUsage
type PartitionsQuotaUsage struct {
	// Maximum number of Partitions allowed.
	Quota int32 `json:"quota"`
	// Remaining number of Partitions allowed.
	Remaining int32 `json:"remaining"`
}

type _PartitionsQuotaUsage PartitionsQuotaUsage

// NewPartitionsQuotaUsage instantiates a new PartitionsQuotaUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartitionsQuotaUsage(quota int32, remaining int32) *PartitionsQuotaUsage {
	this := PartitionsQuotaUsage{}
	this.Quota = quota
	this.Remaining = remaining
	return &this
}

// NewPartitionsQuotaUsageWithDefaults instantiates a new PartitionsQuotaUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionsQuotaUsageWithDefaults() *PartitionsQuotaUsage {
	this := PartitionsQuotaUsage{}
	return &this
}

// GetQuota returns the Quota field value
func (o *PartitionsQuotaUsage) GetQuota() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value
// and a boolean to check if the value has been set.
func (o *PartitionsQuotaUsage) GetQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quota, true
}

// SetQuota sets field value
func (o *PartitionsQuotaUsage) SetQuota(v int32) {
	o.Quota = v
}

// GetRemaining returns the Remaining field value
func (o *PartitionsQuotaUsage) GetRemaining() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *PartitionsQuotaUsage) GetRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *PartitionsQuotaUsage) SetRemaining(v int32) {
	o.Remaining = v
}

func (o PartitionsQuotaUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartitionsQuotaUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quota"] = o.Quota
	toSerialize["remaining"] = o.Remaining
	return toSerialize, nil
}

func (o *PartitionsQuotaUsage) UnmarshalJSON(data []byte) (err error) {
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

	varPartitionsQuotaUsage := _PartitionsQuotaUsage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPartitionsQuotaUsage)

	if err != nil {
		return err
	}

	*o = PartitionsQuotaUsage(varPartitionsQuotaUsage)

	return err
}

type NullablePartitionsQuotaUsage struct {
	value *PartitionsQuotaUsage
	isSet bool
}

func (v NullablePartitionsQuotaUsage) Get() *PartitionsQuotaUsage {
	return v.value
}

func (v *NullablePartitionsQuotaUsage) Set(val *PartitionsQuotaUsage) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitionsQuotaUsage) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitionsQuotaUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitionsQuotaUsage(val *PartitionsQuotaUsage) *NullablePartitionsQuotaUsage {
	return &NullablePartitionsQuotaUsage{value: val, isSet: true}
}

func (v NullablePartitionsQuotaUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitionsQuotaUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


