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

// checks if the UsageForecastResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageForecastResponse{}

// UsageForecastResponse Usage forecast for the organization.
type UsageForecastResponse struct {
	// Average credit usage per day till now.
	AverageUsage *float64 `json:"averageUsage,omitempty"`
	// Percentage of total credits used till date.
	UsagePercentage *float64 `json:"usagePercentage,omitempty"`
	// Total expected usage by the end of contract period.
	ForecastedUsage *float64 `json:"forecastedUsage,omitempty"`
	// Percentage of allocated credits that will be used in the contract period.
	ForecastedUsagePercentage *float64 `json:"forecastedUsagePercentage,omitempty"`
	// Days remaining till all the credits are consumed.
	RemainingDays *float64 `json:"remainingDays,omitempty"`
}

// NewUsageForecastResponse instantiates a new UsageForecastResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageForecastResponse() *UsageForecastResponse {
	this := UsageForecastResponse{}
	return &this
}

// NewUsageForecastResponseWithDefaults instantiates a new UsageForecastResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageForecastResponseWithDefaults() *UsageForecastResponse {
	this := UsageForecastResponse{}
	return &this
}

// GetAverageUsage returns the AverageUsage field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetAverageUsage() float64 {
	if o == nil || IsNil(o.AverageUsage) {
		var ret float64
		return ret
	}
	return *o.AverageUsage
}

// GetAverageUsageOk returns a tuple with the AverageUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetAverageUsageOk() (*float64, bool) {
	if o == nil || IsNil(o.AverageUsage) {
		return nil, false
	}
	return o.AverageUsage, true
}

// HasAverageUsage returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasAverageUsage() bool {
	if o != nil && !IsNil(o.AverageUsage) {
		return true
	}

	return false
}

// SetAverageUsage gets a reference to the given float64 and assigns it to the AverageUsage field.
func (o *UsageForecastResponse) SetAverageUsage(v float64) {
	o.AverageUsage = &v
}

// GetUsagePercentage returns the UsagePercentage field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetUsagePercentage() float64 {
	if o == nil || IsNil(o.UsagePercentage) {
		var ret float64
		return ret
	}
	return *o.UsagePercentage
}

// GetUsagePercentageOk returns a tuple with the UsagePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetUsagePercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.UsagePercentage) {
		return nil, false
	}
	return o.UsagePercentage, true
}

// HasUsagePercentage returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasUsagePercentage() bool {
	if o != nil && !IsNil(o.UsagePercentage) {
		return true
	}

	return false
}

// SetUsagePercentage gets a reference to the given float64 and assigns it to the UsagePercentage field.
func (o *UsageForecastResponse) SetUsagePercentage(v float64) {
	o.UsagePercentage = &v
}

// GetForecastedUsage returns the ForecastedUsage field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetForecastedUsage() float64 {
	if o == nil || IsNil(o.ForecastedUsage) {
		var ret float64
		return ret
	}
	return *o.ForecastedUsage
}

// GetForecastedUsageOk returns a tuple with the ForecastedUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetForecastedUsageOk() (*float64, bool) {
	if o == nil || IsNil(o.ForecastedUsage) {
		return nil, false
	}
	return o.ForecastedUsage, true
}

// HasForecastedUsage returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasForecastedUsage() bool {
	if o != nil && !IsNil(o.ForecastedUsage) {
		return true
	}

	return false
}

// SetForecastedUsage gets a reference to the given float64 and assigns it to the ForecastedUsage field.
func (o *UsageForecastResponse) SetForecastedUsage(v float64) {
	o.ForecastedUsage = &v
}

// GetForecastedUsagePercentage returns the ForecastedUsagePercentage field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetForecastedUsagePercentage() float64 {
	if o == nil || IsNil(o.ForecastedUsagePercentage) {
		var ret float64
		return ret
	}
	return *o.ForecastedUsagePercentage
}

// GetForecastedUsagePercentageOk returns a tuple with the ForecastedUsagePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetForecastedUsagePercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.ForecastedUsagePercentage) {
		return nil, false
	}
	return o.ForecastedUsagePercentage, true
}

// HasForecastedUsagePercentage returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasForecastedUsagePercentage() bool {
	if o != nil && !IsNil(o.ForecastedUsagePercentage) {
		return true
	}

	return false
}

// SetForecastedUsagePercentage gets a reference to the given float64 and assigns it to the ForecastedUsagePercentage field.
func (o *UsageForecastResponse) SetForecastedUsagePercentage(v float64) {
	o.ForecastedUsagePercentage = &v
}

// GetRemainingDays returns the RemainingDays field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetRemainingDays() float64 {
	if o == nil || IsNil(o.RemainingDays) {
		var ret float64
		return ret
	}
	return *o.RemainingDays
}

// GetRemainingDaysOk returns a tuple with the RemainingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetRemainingDaysOk() (*float64, bool) {
	if o == nil || IsNil(o.RemainingDays) {
		return nil, false
	}
	return o.RemainingDays, true
}

// HasRemainingDays returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasRemainingDays() bool {
	if o != nil && !IsNil(o.RemainingDays) {
		return true
	}

	return false
}

// SetRemainingDays gets a reference to the given float64 and assigns it to the RemainingDays field.
func (o *UsageForecastResponse) SetRemainingDays(v float64) {
	o.RemainingDays = &v
}

func (o UsageForecastResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageForecastResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AverageUsage) {
		toSerialize["averageUsage"] = o.AverageUsage
	}
	if !IsNil(o.UsagePercentage) {
		toSerialize["usagePercentage"] = o.UsagePercentage
	}
	if !IsNil(o.ForecastedUsage) {
		toSerialize["forecastedUsage"] = o.ForecastedUsage
	}
	if !IsNil(o.ForecastedUsagePercentage) {
		toSerialize["forecastedUsagePercentage"] = o.ForecastedUsagePercentage
	}
	if !IsNil(o.RemainingDays) {
		toSerialize["remainingDays"] = o.RemainingDays
	}
	return toSerialize, nil
}

type NullableUsageForecastResponse struct {
	value *UsageForecastResponse
	isSet bool
}

func (v NullableUsageForecastResponse) Get() *UsageForecastResponse {
	return v.value
}

func (v *NullableUsageForecastResponse) Set(val *UsageForecastResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageForecastResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageForecastResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageForecastResponse(val *UsageForecastResponse) *NullableUsageForecastResponse {
	return &NullableUsageForecastResponse{value: val, isSet: true}
}

func (v NullableUsageForecastResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageForecastResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


