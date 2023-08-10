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

// checks if the MonitorSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorSubscription{}

// MonitorSubscription The monitor subscription. Alerts can be filtered by a monitor subscription status.
type MonitorSubscription struct {
	// The id of the subscription target. It can be either a monitor or a folder id.
	TargetId *string `json:"targetId,omitempty"`
}

// NewMonitorSubscription instantiates a new MonitorSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorSubscription() *MonitorSubscription {
	this := MonitorSubscription{}
	return &this
}

// NewMonitorSubscriptionWithDefaults instantiates a new MonitorSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorSubscriptionWithDefaults() *MonitorSubscription {
	this := MonitorSubscription{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *MonitorSubscription) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSubscription) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *MonitorSubscription) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *MonitorSubscription) SetTargetId(v string) {
	o.TargetId = &v
}

func (o MonitorSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	return toSerialize, nil
}

type NullableMonitorSubscription struct {
	value *MonitorSubscription
	isSet bool
}

func (v NullableMonitorSubscription) Get() *MonitorSubscription {
	return v.value
}

func (v *NullableMonitorSubscription) Set(val *MonitorSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorSubscription(val *MonitorSubscription) *NullableMonitorSubscription {
	return &NullableMonitorSubscription{value: val, isSet: true}
}

func (v NullableMonitorSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


