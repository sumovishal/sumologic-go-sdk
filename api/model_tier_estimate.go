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

// checks if the TierEstimate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TierEstimate{}

// TierEstimate estimate for a tier
type TierEstimate struct {
	// Name of the data tier
	Tier *string `json:"tier,omitempty"`
	// estimate data scanned per monitor scan in bytes
	PerScanInBytes *int64 `json:"perScanInBytes,omitempty"`
	// estimate data scanned per day in bytes
	PerDayInBytes *int64 `json:"perDayInBytes,omitempty"`
	// estimate data scanned per year in bytes
	PerYearInBytes *int64 `json:"perYearInBytes,omitempty"`
	// one-time scan for log anomaly monitor
	TrainingScanInBytes *int64 `json:"trainingScanInBytes,omitempty"`
}

// NewTierEstimate instantiates a new TierEstimate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTierEstimate() *TierEstimate {
	this := TierEstimate{}
	return &this
}

// NewTierEstimateWithDefaults instantiates a new TierEstimate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTierEstimateWithDefaults() *TierEstimate {
	this := TierEstimate{}
	return &this
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *TierEstimate) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TierEstimate) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *TierEstimate) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *TierEstimate) SetTier(v string) {
	o.Tier = &v
}

// GetPerScanInBytes returns the PerScanInBytes field value if set, zero value otherwise.
func (o *TierEstimate) GetPerScanInBytes() int64 {
	if o == nil || IsNil(o.PerScanInBytes) {
		var ret int64
		return ret
	}
	return *o.PerScanInBytes
}

// GetPerScanInBytesOk returns a tuple with the PerScanInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TierEstimate) GetPerScanInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.PerScanInBytes) {
		return nil, false
	}
	return o.PerScanInBytes, true
}

// HasPerScanInBytes returns a boolean if a field has been set.
func (o *TierEstimate) HasPerScanInBytes() bool {
	if o != nil && !IsNil(o.PerScanInBytes) {
		return true
	}

	return false
}

// SetPerScanInBytes gets a reference to the given int64 and assigns it to the PerScanInBytes field.
func (o *TierEstimate) SetPerScanInBytes(v int64) {
	o.PerScanInBytes = &v
}

// GetPerDayInBytes returns the PerDayInBytes field value if set, zero value otherwise.
func (o *TierEstimate) GetPerDayInBytes() int64 {
	if o == nil || IsNil(o.PerDayInBytes) {
		var ret int64
		return ret
	}
	return *o.PerDayInBytes
}

// GetPerDayInBytesOk returns a tuple with the PerDayInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TierEstimate) GetPerDayInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.PerDayInBytes) {
		return nil, false
	}
	return o.PerDayInBytes, true
}

// HasPerDayInBytes returns a boolean if a field has been set.
func (o *TierEstimate) HasPerDayInBytes() bool {
	if o != nil && !IsNil(o.PerDayInBytes) {
		return true
	}

	return false
}

// SetPerDayInBytes gets a reference to the given int64 and assigns it to the PerDayInBytes field.
func (o *TierEstimate) SetPerDayInBytes(v int64) {
	o.PerDayInBytes = &v
}

// GetPerYearInBytes returns the PerYearInBytes field value if set, zero value otherwise.
func (o *TierEstimate) GetPerYearInBytes() int64 {
	if o == nil || IsNil(o.PerYearInBytes) {
		var ret int64
		return ret
	}
	return *o.PerYearInBytes
}

// GetPerYearInBytesOk returns a tuple with the PerYearInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TierEstimate) GetPerYearInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.PerYearInBytes) {
		return nil, false
	}
	return o.PerYearInBytes, true
}

// HasPerYearInBytes returns a boolean if a field has been set.
func (o *TierEstimate) HasPerYearInBytes() bool {
	if o != nil && !IsNil(o.PerYearInBytes) {
		return true
	}

	return false
}

// SetPerYearInBytes gets a reference to the given int64 and assigns it to the PerYearInBytes field.
func (o *TierEstimate) SetPerYearInBytes(v int64) {
	o.PerYearInBytes = &v
}

// GetTrainingScanInBytes returns the TrainingScanInBytes field value if set, zero value otherwise.
func (o *TierEstimate) GetTrainingScanInBytes() int64 {
	if o == nil || IsNil(o.TrainingScanInBytes) {
		var ret int64
		return ret
	}
	return *o.TrainingScanInBytes
}

// GetTrainingScanInBytesOk returns a tuple with the TrainingScanInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TierEstimate) GetTrainingScanInBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.TrainingScanInBytes) {
		return nil, false
	}
	return o.TrainingScanInBytes, true
}

// HasTrainingScanInBytes returns a boolean if a field has been set.
func (o *TierEstimate) HasTrainingScanInBytes() bool {
	if o != nil && !IsNil(o.TrainingScanInBytes) {
		return true
	}

	return false
}

// SetTrainingScanInBytes gets a reference to the given int64 and assigns it to the TrainingScanInBytes field.
func (o *TierEstimate) SetTrainingScanInBytes(v int64) {
	o.TrainingScanInBytes = &v
}

func (o TierEstimate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TierEstimate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !IsNil(o.PerScanInBytes) {
		toSerialize["perScanInBytes"] = o.PerScanInBytes
	}
	if !IsNil(o.PerDayInBytes) {
		toSerialize["perDayInBytes"] = o.PerDayInBytes
	}
	if !IsNil(o.PerYearInBytes) {
		toSerialize["perYearInBytes"] = o.PerYearInBytes
	}
	if !IsNil(o.TrainingScanInBytes) {
		toSerialize["trainingScanInBytes"] = o.TrainingScanInBytes
	}
	return toSerialize, nil
}

type NullableTierEstimate struct {
	value *TierEstimate
	isSet bool
}

func (v NullableTierEstimate) Get() *TierEstimate {
	return v.value
}

func (v *NullableTierEstimate) Set(val *TierEstimate) {
	v.value = val
	v.isSet = true
}

func (v NullableTierEstimate) IsSet() bool {
	return v.isSet
}

func (v *NullableTierEstimate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTierEstimate(val *TierEstimate) *NullableTierEstimate {
	return &NullableTierEstimate{value: val, isSet: true}
}

func (v NullableTierEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTierEstimate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


