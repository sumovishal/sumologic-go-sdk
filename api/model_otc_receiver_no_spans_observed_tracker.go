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

// checks if the OTCReceiverNoSpansObservedTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCReceiverNoSpansObservedTracker{}

// OTCReceiverNoSpansObservedTracker struct for OTCReceiverNoSpansObservedTracker
type OTCReceiverNoSpansObservedTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The collector instance ID, e.g. `974b444b-4b45-4f32-aa03-1dbf2a16826d`.
	InstanceId *string `json:"instanceId,omitempty"`
	// The collector instance address, e.g. `172.16.1.14`.
	InstanceAddress *string `json:"instanceAddress,omitempty"`
	// The collector receiver ID, e.g. `otlphttp/2`.
	ReceiverId *string `json:"receiverId,omitempty"`
}

// NewOTCReceiverNoSpansObservedTracker instantiates a new OTCReceiverNoSpansObservedTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCReceiverNoSpansObservedTracker(trackerId string, error_ string, description string) *OTCReceiverNoSpansObservedTracker {
	this := OTCReceiverNoSpansObservedTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewOTCReceiverNoSpansObservedTrackerWithDefaults instantiates a new OTCReceiverNoSpansObservedTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCReceiverNoSpansObservedTrackerWithDefaults() *OTCReceiverNoSpansObservedTracker {
	this := OTCReceiverNoSpansObservedTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *OTCReceiverNoSpansObservedTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverNoSpansObservedTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *OTCReceiverNoSpansObservedTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *OTCReceiverNoSpansObservedTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *OTCReceiverNoSpansObservedTracker) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverNoSpansObservedTracker) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *OTCReceiverNoSpansObservedTracker) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *OTCReceiverNoSpansObservedTracker) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceAddress returns the InstanceAddress field value if set, zero value otherwise.
func (o *OTCReceiverNoSpansObservedTracker) GetInstanceAddress() string {
	if o == nil || IsNil(o.InstanceAddress) {
		var ret string
		return ret
	}
	return *o.InstanceAddress
}

// GetInstanceAddressOk returns a tuple with the InstanceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverNoSpansObservedTracker) GetInstanceAddressOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceAddress) {
		return nil, false
	}
	return o.InstanceAddress, true
}

// HasInstanceAddress returns a boolean if a field has been set.
func (o *OTCReceiverNoSpansObservedTracker) HasInstanceAddress() bool {
	if o != nil && !IsNil(o.InstanceAddress) {
		return true
	}

	return false
}

// SetInstanceAddress gets a reference to the given string and assigns it to the InstanceAddress field.
func (o *OTCReceiverNoSpansObservedTracker) SetInstanceAddress(v string) {
	o.InstanceAddress = &v
}

// GetReceiverId returns the ReceiverId field value if set, zero value otherwise.
func (o *OTCReceiverNoSpansObservedTracker) GetReceiverId() string {
	if o == nil || IsNil(o.ReceiverId) {
		var ret string
		return ret
	}
	return *o.ReceiverId
}

// GetReceiverIdOk returns a tuple with the ReceiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverNoSpansObservedTracker) GetReceiverIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiverId) {
		return nil, false
	}
	return o.ReceiverId, true
}

// HasReceiverId returns a boolean if a field has been set.
func (o *OTCReceiverNoSpansObservedTracker) HasReceiverId() bool {
	if o != nil && !IsNil(o.ReceiverId) {
		return true
	}

	return false
}

// SetReceiverId gets a reference to the given string and assigns it to the ReceiverId field.
func (o *OTCReceiverNoSpansObservedTracker) SetReceiverId(v string) {
	o.ReceiverId = &v
}

func (o OTCReceiverNoSpansObservedTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCReceiverNoSpansObservedTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.InstanceAddress) {
		toSerialize["instanceAddress"] = o.InstanceAddress
	}
	if !IsNil(o.ReceiverId) {
		toSerialize["receiverId"] = o.ReceiverId
	}
	return toSerialize, nil
}

type NullableOTCReceiverNoSpansObservedTracker struct {
	value *OTCReceiverNoSpansObservedTracker
	isSet bool
}

func (v NullableOTCReceiverNoSpansObservedTracker) Get() *OTCReceiverNoSpansObservedTracker {
	return v.value
}

func (v *NullableOTCReceiverNoSpansObservedTracker) Set(val *OTCReceiverNoSpansObservedTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCReceiverNoSpansObservedTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCReceiverNoSpansObservedTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCReceiverNoSpansObservedTracker(val *OTCReceiverNoSpansObservedTracker) *NullableOTCReceiverNoSpansObservedTracker {
	return &NullableOTCReceiverNoSpansObservedTracker{value: val, isSet: true}
}

func (v NullableOTCReceiverNoSpansObservedTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCReceiverNoSpansObservedTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


