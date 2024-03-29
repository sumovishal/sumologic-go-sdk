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

// checks if the OTCWarningProcessingSpansTrackerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCWarningProcessingSpansTrackerAllOf{}

// OTCWarningProcessingSpansTrackerAllOf struct for OTCWarningProcessingSpansTrackerAllOf
type OTCWarningProcessingSpansTrackerAllOf struct {
	// The collector instance ID, e.g. `974b444b-4b45-4f32-aa03-1dbf2a16826d`.
	InstanceId *string `json:"instanceId,omitempty"`
	// The collector instance address, e.g. `172.16.1.14`.
	InstanceAddress *string `json:"instanceAddress,omitempty"`
	// The collector processor ID, e.g. `cascading_filter`.
	ProcessorId *string `json:"processorId,omitempty"`
	// The warning message.
	Message *string `json:"message,omitempty"`
}

// NewOTCWarningProcessingSpansTrackerAllOf instantiates a new OTCWarningProcessingSpansTrackerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCWarningProcessingSpansTrackerAllOf() *OTCWarningProcessingSpansTrackerAllOf {
	this := OTCWarningProcessingSpansTrackerAllOf{}
	return &this
}

// NewOTCWarningProcessingSpansTrackerAllOfWithDefaults instantiates a new OTCWarningProcessingSpansTrackerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCWarningProcessingSpansTrackerAllOfWithDefaults() *OTCWarningProcessingSpansTrackerAllOf {
	this := OTCWarningProcessingSpansTrackerAllOf{}
	return &this
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *OTCWarningProcessingSpansTrackerAllOf) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCWarningProcessingSpansTrackerAllOf) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *OTCWarningProcessingSpansTrackerAllOf) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *OTCWarningProcessingSpansTrackerAllOf) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceAddress returns the InstanceAddress field value if set, zero value otherwise.
func (o *OTCWarningProcessingSpansTrackerAllOf) GetInstanceAddress() string {
	if o == nil || IsNil(o.InstanceAddress) {
		var ret string
		return ret
	}
	return *o.InstanceAddress
}

// GetInstanceAddressOk returns a tuple with the InstanceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCWarningProcessingSpansTrackerAllOf) GetInstanceAddressOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceAddress) {
		return nil, false
	}
	return o.InstanceAddress, true
}

// HasInstanceAddress returns a boolean if a field has been set.
func (o *OTCWarningProcessingSpansTrackerAllOf) HasInstanceAddress() bool {
	if o != nil && !IsNil(o.InstanceAddress) {
		return true
	}

	return false
}

// SetInstanceAddress gets a reference to the given string and assigns it to the InstanceAddress field.
func (o *OTCWarningProcessingSpansTrackerAllOf) SetInstanceAddress(v string) {
	o.InstanceAddress = &v
}

// GetProcessorId returns the ProcessorId field value if set, zero value otherwise.
func (o *OTCWarningProcessingSpansTrackerAllOf) GetProcessorId() string {
	if o == nil || IsNil(o.ProcessorId) {
		var ret string
		return ret
	}
	return *o.ProcessorId
}

// GetProcessorIdOk returns a tuple with the ProcessorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCWarningProcessingSpansTrackerAllOf) GetProcessorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorId) {
		return nil, false
	}
	return o.ProcessorId, true
}

// HasProcessorId returns a boolean if a field has been set.
func (o *OTCWarningProcessingSpansTrackerAllOf) HasProcessorId() bool {
	if o != nil && !IsNil(o.ProcessorId) {
		return true
	}

	return false
}

// SetProcessorId gets a reference to the given string and assigns it to the ProcessorId field.
func (o *OTCWarningProcessingSpansTrackerAllOf) SetProcessorId(v string) {
	o.ProcessorId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OTCWarningProcessingSpansTrackerAllOf) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCWarningProcessingSpansTrackerAllOf) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OTCWarningProcessingSpansTrackerAllOf) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OTCWarningProcessingSpansTrackerAllOf) SetMessage(v string) {
	o.Message = &v
}

func (o OTCWarningProcessingSpansTrackerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCWarningProcessingSpansTrackerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.InstanceAddress) {
		toSerialize["instanceAddress"] = o.InstanceAddress
	}
	if !IsNil(o.ProcessorId) {
		toSerialize["processorId"] = o.ProcessorId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableOTCWarningProcessingSpansTrackerAllOf struct {
	value *OTCWarningProcessingSpansTrackerAllOf
	isSet bool
}

func (v NullableOTCWarningProcessingSpansTrackerAllOf) Get() *OTCWarningProcessingSpansTrackerAllOf {
	return v.value
}

func (v *NullableOTCWarningProcessingSpansTrackerAllOf) Set(val *OTCWarningProcessingSpansTrackerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCWarningProcessingSpansTrackerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCWarningProcessingSpansTrackerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCWarningProcessingSpansTrackerAllOf(val *OTCWarningProcessingSpansTrackerAllOf) *NullableOTCWarningProcessingSpansTrackerAllOf {
	return &NullableOTCWarningProcessingSpansTrackerAllOf{value: val, isSet: true}
}

func (v NullableOTCWarningProcessingSpansTrackerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCWarningProcessingSpansTrackerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


