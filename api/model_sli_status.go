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

// checks if the SliStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SliStatus{}

// SliStatus Status of the SLI computation. If the status is successful, also contains the SLI value and error budget remaining for the current compliance period.
type SliStatus struct {
	// Whether the SLI computation is complete / had an error / is in progress.
	Status string `json:"status"`
	// SLI percentage for the compliance period. Available if `status` is `Success`.
	SliPercentage *float64 `json:"sliPercentage,omitempty"`
	// Percentage of error budget remaining for the compliance period. Available if `status` is `Success`.
	ErrorBudgetRemainingPercentage *float64 `json:"errorBudgetRemainingPercentage,omitempty"`
	// Formatted string for the absolute error budget remaining (time duration for window-based SLIs, request count for request-based SLIs). Available if `status` is `Success`.
	AbsoluteErrorBudgetRemaining *string `json:"absoluteErrorBudgetRemaining,omitempty"`
	// SLI computation progress.
	ProgressPercentage *float64 `json:"progressPercentage,omitempty"`
}

// NewSliStatus instantiates a new SliStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliStatus(status string) *SliStatus {
	this := SliStatus{}
	this.Status = status
	return &this
}

// NewSliStatusWithDefaults instantiates a new SliStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliStatusWithDefaults() *SliStatus {
	this := SliStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *SliStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SliStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SliStatus) SetStatus(v string) {
	o.Status = v
}

// GetSliPercentage returns the SliPercentage field value if set, zero value otherwise.
func (o *SliStatus) GetSliPercentage() float64 {
	if o == nil || IsNil(o.SliPercentage) {
		var ret float64
		return ret
	}
	return *o.SliPercentage
}

// GetSliPercentageOk returns a tuple with the SliPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliStatus) GetSliPercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.SliPercentage) {
		return nil, false
	}
	return o.SliPercentage, true
}

// HasSliPercentage returns a boolean if a field has been set.
func (o *SliStatus) HasSliPercentage() bool {
	if o != nil && !IsNil(o.SliPercentage) {
		return true
	}

	return false
}

// SetSliPercentage gets a reference to the given float64 and assigns it to the SliPercentage field.
func (o *SliStatus) SetSliPercentage(v float64) {
	o.SliPercentage = &v
}

// GetErrorBudgetRemainingPercentage returns the ErrorBudgetRemainingPercentage field value if set, zero value otherwise.
func (o *SliStatus) GetErrorBudgetRemainingPercentage() float64 {
	if o == nil || IsNil(o.ErrorBudgetRemainingPercentage) {
		var ret float64
		return ret
	}
	return *o.ErrorBudgetRemainingPercentage
}

// GetErrorBudgetRemainingPercentageOk returns a tuple with the ErrorBudgetRemainingPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliStatus) GetErrorBudgetRemainingPercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.ErrorBudgetRemainingPercentage) {
		return nil, false
	}
	return o.ErrorBudgetRemainingPercentage, true
}

// HasErrorBudgetRemainingPercentage returns a boolean if a field has been set.
func (o *SliStatus) HasErrorBudgetRemainingPercentage() bool {
	if o != nil && !IsNil(o.ErrorBudgetRemainingPercentage) {
		return true
	}

	return false
}

// SetErrorBudgetRemainingPercentage gets a reference to the given float64 and assigns it to the ErrorBudgetRemainingPercentage field.
func (o *SliStatus) SetErrorBudgetRemainingPercentage(v float64) {
	o.ErrorBudgetRemainingPercentage = &v
}

// GetAbsoluteErrorBudgetRemaining returns the AbsoluteErrorBudgetRemaining field value if set, zero value otherwise.
func (o *SliStatus) GetAbsoluteErrorBudgetRemaining() string {
	if o == nil || IsNil(o.AbsoluteErrorBudgetRemaining) {
		var ret string
		return ret
	}
	return *o.AbsoluteErrorBudgetRemaining
}

// GetAbsoluteErrorBudgetRemainingOk returns a tuple with the AbsoluteErrorBudgetRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliStatus) GetAbsoluteErrorBudgetRemainingOk() (*string, bool) {
	if o == nil || IsNil(o.AbsoluteErrorBudgetRemaining) {
		return nil, false
	}
	return o.AbsoluteErrorBudgetRemaining, true
}

// HasAbsoluteErrorBudgetRemaining returns a boolean if a field has been set.
func (o *SliStatus) HasAbsoluteErrorBudgetRemaining() bool {
	if o != nil && !IsNil(o.AbsoluteErrorBudgetRemaining) {
		return true
	}

	return false
}

// SetAbsoluteErrorBudgetRemaining gets a reference to the given string and assigns it to the AbsoluteErrorBudgetRemaining field.
func (o *SliStatus) SetAbsoluteErrorBudgetRemaining(v string) {
	o.AbsoluteErrorBudgetRemaining = &v
}

// GetProgressPercentage returns the ProgressPercentage field value if set, zero value otherwise.
func (o *SliStatus) GetProgressPercentage() float64 {
	if o == nil || IsNil(o.ProgressPercentage) {
		var ret float64
		return ret
	}
	return *o.ProgressPercentage
}

// GetProgressPercentageOk returns a tuple with the ProgressPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliStatus) GetProgressPercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.ProgressPercentage) {
		return nil, false
	}
	return o.ProgressPercentage, true
}

// HasProgressPercentage returns a boolean if a field has been set.
func (o *SliStatus) HasProgressPercentage() bool {
	if o != nil && !IsNil(o.ProgressPercentage) {
		return true
	}

	return false
}

// SetProgressPercentage gets a reference to the given float64 and assigns it to the ProgressPercentage field.
func (o *SliStatus) SetProgressPercentage(v float64) {
	o.ProgressPercentage = &v
}

func (o SliStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SliStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.SliPercentage) {
		toSerialize["sliPercentage"] = o.SliPercentage
	}
	if !IsNil(o.ErrorBudgetRemainingPercentage) {
		toSerialize["errorBudgetRemainingPercentage"] = o.ErrorBudgetRemainingPercentage
	}
	if !IsNil(o.AbsoluteErrorBudgetRemaining) {
		toSerialize["absoluteErrorBudgetRemaining"] = o.AbsoluteErrorBudgetRemaining
	}
	if !IsNil(o.ProgressPercentage) {
		toSerialize["progressPercentage"] = o.ProgressPercentage
	}
	return toSerialize, nil
}

type NullableSliStatus struct {
	value *SliStatus
	isSet bool
}

func (v NullableSliStatus) Get() *SliStatus {
	return v.value
}

func (v *NullableSliStatus) Set(val *SliStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSliStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSliStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliStatus(val *SliStatus) *NullableSliStatus {
	return &NullableSliStatus{value: val, isSet: true}
}

func (v NullableSliStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


