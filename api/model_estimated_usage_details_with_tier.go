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

// checks if the EstimatedUsageDetailsWithTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimatedUsageDetailsWithTier{}

// EstimatedUsageDetailsWithTier struct for EstimatedUsageDetailsWithTier
type EstimatedUsageDetailsWithTier struct {
	// Name of the data tier. Supported Values are Continuous, Frequent, Infrequent
	Tier *string `json:"tier,omitempty"`
	// Amount of data scanned in bytes, to run the query.
	DataScannedInBytes *int64 `json:"dataScannedInBytes,omitempty"`
}

// NewEstimatedUsageDetailsWithTier instantiates a new EstimatedUsageDetailsWithTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimatedUsageDetailsWithTier() *EstimatedUsageDetailsWithTier {
	this := EstimatedUsageDetailsWithTier{}
	return &this
}

// NewEstimatedUsageDetailsWithTierWithDefaults instantiates a new EstimatedUsageDetailsWithTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimatedUsageDetailsWithTierWithDefaults() *EstimatedUsageDetailsWithTier {
	this := EstimatedUsageDetailsWithTier{}
	return &this
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *EstimatedUsageDetailsWithTier) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedUsageDetailsWithTier) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *EstimatedUsageDetailsWithTier) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *EstimatedUsageDetailsWithTier) SetTier(v string) {
	o.Tier = &v
}

// GetDataScannedInBytes returns the DataScannedInBytes field value if set, zero value otherwise.
func (o *EstimatedUsageDetailsWithTier) GetDataScannedInBytes() int64 {
	if o == nil || IsNil(o.DataScannedInBytes) {
		var ret int64
		return ret
	}
	return *o.DataScannedInBytes
}

// GetDataScannedInBytesOk returns a tuple with the DataScannedInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedUsageDetailsWithTier) GetDataScannedInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DataScannedInBytes) {
		return nil, false
	}
	return o.DataScannedInBytes, true
}

// HasDataScannedInBytes returns a boolean if a field has been set.
func (o *EstimatedUsageDetailsWithTier) HasDataScannedInBytes() bool {
	if o != nil && !IsNil(o.DataScannedInBytes) {
		return true
	}

	return false
}

// SetDataScannedInBytes gets a reference to the given int64 and assigns it to the DataScannedInBytes field.
func (o *EstimatedUsageDetailsWithTier) SetDataScannedInBytes(v int64) {
	o.DataScannedInBytes = &v
}

func (o EstimatedUsageDetailsWithTier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimatedUsageDetailsWithTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !IsNil(o.DataScannedInBytes) {
		toSerialize["dataScannedInBytes"] = o.DataScannedInBytes
	}
	return toSerialize, nil
}

type NullableEstimatedUsageDetailsWithTier struct {
	value *EstimatedUsageDetailsWithTier
	isSet bool
}

func (v NullableEstimatedUsageDetailsWithTier) Get() *EstimatedUsageDetailsWithTier {
	return v.value
}

func (v *NullableEstimatedUsageDetailsWithTier) Set(val *EstimatedUsageDetailsWithTier) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedUsageDetailsWithTier) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedUsageDetailsWithTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedUsageDetailsWithTier(val *EstimatedUsageDetailsWithTier) *NullableEstimatedUsageDetailsWithTier {
	return &NullableEstimatedUsageDetailsWithTier{value: val, isSet: true}
}

func (v NullableEstimatedUsageDetailsWithTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedUsageDetailsWithTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


