/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// checks if the ViewRetentionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewRetentionProperties{}

// ViewRetentionProperties struct for ViewRetentionProperties
type ViewRetentionProperties struct {
	// If the retention period is scheduled to be updated in the future (i.e., if retention period is previously reduced with value of reduceRetentionPeriodImmediately as false), this property gives the future value of retention period while retentionPeriod gives the current value. retentionPeriod will take up the value of newRetentionPeriod after the scheduled time.
	NewRetentionPeriod *int32 `json:"newRetentionPeriod,omitempty"`
	// When the newRetentionPeriod will become effective in UTC format.
	RetentionEffectiveAt *time.Time `json:"retentionEffectiveAt,omitempty"`
}

// NewViewRetentionProperties instantiates a new ViewRetentionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewRetentionProperties() *ViewRetentionProperties {
	this := ViewRetentionProperties{}
	return &this
}

// NewViewRetentionPropertiesWithDefaults instantiates a new ViewRetentionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewRetentionPropertiesWithDefaults() *ViewRetentionProperties {
	this := ViewRetentionProperties{}
	return &this
}

// GetNewRetentionPeriod returns the NewRetentionPeriod field value if set, zero value otherwise.
func (o *ViewRetentionProperties) GetNewRetentionPeriod() int32 {
	if o == nil || IsNil(o.NewRetentionPeriod) {
		var ret int32
		return ret
	}
	return *o.NewRetentionPeriod
}

// GetNewRetentionPeriodOk returns a tuple with the NewRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewRetentionProperties) GetNewRetentionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.NewRetentionPeriod) {
		return nil, false
	}
	return o.NewRetentionPeriod, true
}

// HasNewRetentionPeriod returns a boolean if a field has been set.
func (o *ViewRetentionProperties) HasNewRetentionPeriod() bool {
	if o != nil && !IsNil(o.NewRetentionPeriod) {
		return true
	}

	return false
}

// SetNewRetentionPeriod gets a reference to the given int32 and assigns it to the NewRetentionPeriod field.
func (o *ViewRetentionProperties) SetNewRetentionPeriod(v int32) {
	o.NewRetentionPeriod = &v
}

// GetRetentionEffectiveAt returns the RetentionEffectiveAt field value if set, zero value otherwise.
func (o *ViewRetentionProperties) GetRetentionEffectiveAt() time.Time {
	if o == nil || IsNil(o.RetentionEffectiveAt) {
		var ret time.Time
		return ret
	}
	return *o.RetentionEffectiveAt
}

// GetRetentionEffectiveAtOk returns a tuple with the RetentionEffectiveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewRetentionProperties) GetRetentionEffectiveAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RetentionEffectiveAt) {
		return nil, false
	}
	return o.RetentionEffectiveAt, true
}

// HasRetentionEffectiveAt returns a boolean if a field has been set.
func (o *ViewRetentionProperties) HasRetentionEffectiveAt() bool {
	if o != nil && !IsNil(o.RetentionEffectiveAt) {
		return true
	}

	return false
}

// SetRetentionEffectiveAt gets a reference to the given time.Time and assigns it to the RetentionEffectiveAt field.
func (o *ViewRetentionProperties) SetRetentionEffectiveAt(v time.Time) {
	o.RetentionEffectiveAt = &v
}

func (o ViewRetentionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewRetentionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewRetentionPeriod) {
		toSerialize["newRetentionPeriod"] = o.NewRetentionPeriod
	}
	if !IsNil(o.RetentionEffectiveAt) {
		toSerialize["retentionEffectiveAt"] = o.RetentionEffectiveAt
	}
	return toSerialize, nil
}

type NullableViewRetentionProperties struct {
	value *ViewRetentionProperties
	isSet bool
}

func (v NullableViewRetentionProperties) Get() *ViewRetentionProperties {
	return v.value
}

func (v *NullableViewRetentionProperties) Set(val *ViewRetentionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableViewRetentionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableViewRetentionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewRetentionProperties(val *ViewRetentionProperties) *NullableViewRetentionProperties {
	return &NullableViewRetentionProperties{value: val, isSet: true}
}

func (v NullableViewRetentionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewRetentionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


