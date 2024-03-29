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

// checks if the UpdateScheduledViewDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateScheduledViewDefinition{}

// UpdateScheduledViewDefinition struct for UpdateScheduledViewDefinition
type UpdateScheduledViewDefinition struct {
	// An optional ID of a data forwarding configuration to be used by the scheduled view.
	DataForwardingId *string `json:"dataForwardingId,omitempty"`
	// The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled.
	RetentionPeriod *int32 `json:"retentionPeriod,omitempty"`
	// This is required if the newly specified `retentionPeriod` is less than the existing retention period.  In such a situation, a value of `true` says that data between the existing retention period and the new retention period should be deleted immediately; if `false`, such data will be deleted after seven days. This property is optional and ignored if the specified `retentionPeriod` is greater than or equal to the current retention period.
	ReduceRetentionPeriodImmediately *bool `json:"reduceRetentionPeriodImmediately,omitempty"`
}

// NewUpdateScheduledViewDefinition instantiates a new UpdateScheduledViewDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateScheduledViewDefinition() *UpdateScheduledViewDefinition {
	this := UpdateScheduledViewDefinition{}
	var retentionPeriod int32 = -1
	this.RetentionPeriod = &retentionPeriod
	var reduceRetentionPeriodImmediately bool = false
	this.ReduceRetentionPeriodImmediately = &reduceRetentionPeriodImmediately
	return &this
}

// NewUpdateScheduledViewDefinitionWithDefaults instantiates a new UpdateScheduledViewDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateScheduledViewDefinitionWithDefaults() *UpdateScheduledViewDefinition {
	this := UpdateScheduledViewDefinition{}
	var retentionPeriod int32 = -1
	this.RetentionPeriod = &retentionPeriod
	var reduceRetentionPeriodImmediately bool = false
	this.ReduceRetentionPeriodImmediately = &reduceRetentionPeriodImmediately
	return &this
}

// GetDataForwardingId returns the DataForwardingId field value if set, zero value otherwise.
func (o *UpdateScheduledViewDefinition) GetDataForwardingId() string {
	if o == nil || IsNil(o.DataForwardingId) {
		var ret string
		return ret
	}
	return *o.DataForwardingId
}

// GetDataForwardingIdOk returns a tuple with the DataForwardingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledViewDefinition) GetDataForwardingIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataForwardingId) {
		return nil, false
	}
	return o.DataForwardingId, true
}

// HasDataForwardingId returns a boolean if a field has been set.
func (o *UpdateScheduledViewDefinition) HasDataForwardingId() bool {
	if o != nil && !IsNil(o.DataForwardingId) {
		return true
	}

	return false
}

// SetDataForwardingId gets a reference to the given string and assigns it to the DataForwardingId field.
func (o *UpdateScheduledViewDefinition) SetDataForwardingId(v string) {
	o.DataForwardingId = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *UpdateScheduledViewDefinition) GetRetentionPeriod() int32 {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret int32
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledViewDefinition) GetRetentionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RetentionPeriod) {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *UpdateScheduledViewDefinition) HasRetentionPeriod() bool {
	if o != nil && !IsNil(o.RetentionPeriod) {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int32 and assigns it to the RetentionPeriod field.
func (o *UpdateScheduledViewDefinition) SetRetentionPeriod(v int32) {
	o.RetentionPeriod = &v
}

// GetReduceRetentionPeriodImmediately returns the ReduceRetentionPeriodImmediately field value if set, zero value otherwise.
func (o *UpdateScheduledViewDefinition) GetReduceRetentionPeriodImmediately() bool {
	if o == nil || IsNil(o.ReduceRetentionPeriodImmediately) {
		var ret bool
		return ret
	}
	return *o.ReduceRetentionPeriodImmediately
}

// GetReduceRetentionPeriodImmediatelyOk returns a tuple with the ReduceRetentionPeriodImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledViewDefinition) GetReduceRetentionPeriodImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceRetentionPeriodImmediately) {
		return nil, false
	}
	return o.ReduceRetentionPeriodImmediately, true
}

// HasReduceRetentionPeriodImmediately returns a boolean if a field has been set.
func (o *UpdateScheduledViewDefinition) HasReduceRetentionPeriodImmediately() bool {
	if o != nil && !IsNil(o.ReduceRetentionPeriodImmediately) {
		return true
	}

	return false
}

// SetReduceRetentionPeriodImmediately gets a reference to the given bool and assigns it to the ReduceRetentionPeriodImmediately field.
func (o *UpdateScheduledViewDefinition) SetReduceRetentionPeriodImmediately(v bool) {
	o.ReduceRetentionPeriodImmediately = &v
}

func (o UpdateScheduledViewDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateScheduledViewDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataForwardingId) {
		toSerialize["dataForwardingId"] = o.DataForwardingId
	}
	if !IsNil(o.RetentionPeriod) {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if !IsNil(o.ReduceRetentionPeriodImmediately) {
		toSerialize["reduceRetentionPeriodImmediately"] = o.ReduceRetentionPeriodImmediately
	}
	return toSerialize, nil
}

type NullableUpdateScheduledViewDefinition struct {
	value *UpdateScheduledViewDefinition
	isSet bool
}

func (v NullableUpdateScheduledViewDefinition) Get() *UpdateScheduledViewDefinition {
	return v.value
}

func (v *NullableUpdateScheduledViewDefinition) Set(val *UpdateScheduledViewDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateScheduledViewDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateScheduledViewDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateScheduledViewDefinition(val *UpdateScheduledViewDefinition) *NullableUpdateScheduledViewDefinition {
	return &NullableUpdateScheduledViewDefinition{value: val, isSet: true}
}

func (v NullableUpdateScheduledViewDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateScheduledViewDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


