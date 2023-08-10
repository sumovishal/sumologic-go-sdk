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

// checks if the OTCProcessHighMemoryUsageTrackerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCProcessHighMemoryUsageTrackerAllOf{}

// OTCProcessHighMemoryUsageTrackerAllOf struct for OTCProcessHighMemoryUsageTrackerAllOf
type OTCProcessHighMemoryUsageTrackerAllOf struct {
	// The collector instance ID, e.g. `974b444b-4b45-4f32-aa03-1dbf2a16826d`.
	InstanceId *string `json:"instanceId,omitempty"`
	// The collector instance address, e.g. `172.16.1.14`.
	InstanceAddress *string `json:"instanceAddress,omitempty"`
	// The collector memory usage in bytes, e.g. `142606592`
	MemoryUsage *string `json:"memoryUsage,omitempty"`
	// The collector memory limit (if set) in bytes, e.g. `4000000000`
	MemoryLimit *string `json:"memoryLimit,omitempty"`
}

// NewOTCProcessHighMemoryUsageTrackerAllOf instantiates a new OTCProcessHighMemoryUsageTrackerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCProcessHighMemoryUsageTrackerAllOf() *OTCProcessHighMemoryUsageTrackerAllOf {
	this := OTCProcessHighMemoryUsageTrackerAllOf{}
	return &this
}

// NewOTCProcessHighMemoryUsageTrackerAllOfWithDefaults instantiates a new OTCProcessHighMemoryUsageTrackerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCProcessHighMemoryUsageTrackerAllOfWithDefaults() *OTCProcessHighMemoryUsageTrackerAllOf {
	this := OTCProcessHighMemoryUsageTrackerAllOf{}
	return &this
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceAddress returns the InstanceAddress field value if set, zero value otherwise.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) GetInstanceAddress() string {
	if o == nil || IsNil(o.InstanceAddress) {
		var ret string
		return ret
	}
	return *o.InstanceAddress
}

// GetInstanceAddressOk returns a tuple with the InstanceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) GetInstanceAddressOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceAddress) {
		return nil, false
	}
	return o.InstanceAddress, true
}

// HasInstanceAddress returns a boolean if a field has been set.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) HasInstanceAddress() bool {
	if o != nil && !IsNil(o.InstanceAddress) {
		return true
	}

	return false
}

// SetInstanceAddress gets a reference to the given string and assigns it to the InstanceAddress field.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) SetInstanceAddress(v string) {
	o.InstanceAddress = &v
}

// GetMemoryUsage returns the MemoryUsage field value if set, zero value otherwise.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) GetMemoryUsage() string {
	if o == nil || IsNil(o.MemoryUsage) {
		var ret string
		return ret
	}
	return *o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) GetMemoryUsageOk() (*string, bool) {
	if o == nil || IsNil(o.MemoryUsage) {
		return nil, false
	}
	return o.MemoryUsage, true
}

// HasMemoryUsage returns a boolean if a field has been set.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) HasMemoryUsage() bool {
	if o != nil && !IsNil(o.MemoryUsage) {
		return true
	}

	return false
}

// SetMemoryUsage gets a reference to the given string and assigns it to the MemoryUsage field.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) SetMemoryUsage(v string) {
	o.MemoryUsage = &v
}

// GetMemoryLimit returns the MemoryLimit field value if set, zero value otherwise.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) GetMemoryLimit() string {
	if o == nil || IsNil(o.MemoryLimit) {
		var ret string
		return ret
	}
	return *o.MemoryLimit
}

// GetMemoryLimitOk returns a tuple with the MemoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) GetMemoryLimitOk() (*string, bool) {
	if o == nil || IsNil(o.MemoryLimit) {
		return nil, false
	}
	return o.MemoryLimit, true
}

// HasMemoryLimit returns a boolean if a field has been set.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) HasMemoryLimit() bool {
	if o != nil && !IsNil(o.MemoryLimit) {
		return true
	}

	return false
}

// SetMemoryLimit gets a reference to the given string and assigns it to the MemoryLimit field.
func (o *OTCProcessHighMemoryUsageTrackerAllOf) SetMemoryLimit(v string) {
	o.MemoryLimit = &v
}

func (o OTCProcessHighMemoryUsageTrackerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCProcessHighMemoryUsageTrackerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.InstanceAddress) {
		toSerialize["instanceAddress"] = o.InstanceAddress
	}
	if !IsNil(o.MemoryUsage) {
		toSerialize["memoryUsage"] = o.MemoryUsage
	}
	if !IsNil(o.MemoryLimit) {
		toSerialize["memoryLimit"] = o.MemoryLimit
	}
	return toSerialize, nil
}

type NullableOTCProcessHighMemoryUsageTrackerAllOf struct {
	value *OTCProcessHighMemoryUsageTrackerAllOf
	isSet bool
}

func (v NullableOTCProcessHighMemoryUsageTrackerAllOf) Get() *OTCProcessHighMemoryUsageTrackerAllOf {
	return v.value
}

func (v *NullableOTCProcessHighMemoryUsageTrackerAllOf) Set(val *OTCProcessHighMemoryUsageTrackerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCProcessHighMemoryUsageTrackerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCProcessHighMemoryUsageTrackerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCProcessHighMemoryUsageTrackerAllOf(val *OTCProcessHighMemoryUsageTrackerAllOf) *NullableOTCProcessHighMemoryUsageTrackerAllOf {
	return &NullableOTCProcessHighMemoryUsageTrackerAllOf{value: val, isSet: true}
}

func (v NullableOTCProcessHighMemoryUsageTrackerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCProcessHighMemoryUsageTrackerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


