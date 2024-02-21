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

// checks if the OTCProcessSpansDroppedTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCProcessSpansDroppedTracker{}

// OTCProcessSpansDroppedTracker struct for OTCProcessSpansDroppedTracker
type OTCProcessSpansDroppedTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The collector instance ID, e.g. `974b444b-4b45-4f32-aa03-1dbf2a16826d`.
	InstanceId *string `json:"instanceId,omitempty"`
	// The collector instance address, e.g. `172.16.1.14`.
	InstanceAddress *string `json:"instanceAddress,omitempty"`
	// The collector processor ID, e.g. `cascading_filter`.
	ProcessorId *string `json:"processorId,omitempty"`
	// The count of dropped spans.
	Count *string `json:"count,omitempty"`
}

// NewOTCProcessSpansDroppedTracker instantiates a new OTCProcessSpansDroppedTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCProcessSpansDroppedTracker(trackerId string, error_ string, description string) *OTCProcessSpansDroppedTracker {
	this := OTCProcessSpansDroppedTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewOTCProcessSpansDroppedTrackerWithDefaults instantiates a new OTCProcessSpansDroppedTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCProcessSpansDroppedTrackerWithDefaults() *OTCProcessSpansDroppedTracker {
	this := OTCProcessSpansDroppedTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *OTCProcessSpansDroppedTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessSpansDroppedTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *OTCProcessSpansDroppedTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *OTCProcessSpansDroppedTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *OTCProcessSpansDroppedTracker) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessSpansDroppedTracker) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *OTCProcessSpansDroppedTracker) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *OTCProcessSpansDroppedTracker) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceAddress returns the InstanceAddress field value if set, zero value otherwise.
func (o *OTCProcessSpansDroppedTracker) GetInstanceAddress() string {
	if o == nil || IsNil(o.InstanceAddress) {
		var ret string
		return ret
	}
	return *o.InstanceAddress
}

// GetInstanceAddressOk returns a tuple with the InstanceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessSpansDroppedTracker) GetInstanceAddressOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceAddress) {
		return nil, false
	}
	return o.InstanceAddress, true
}

// HasInstanceAddress returns a boolean if a field has been set.
func (o *OTCProcessSpansDroppedTracker) HasInstanceAddress() bool {
	if o != nil && !IsNil(o.InstanceAddress) {
		return true
	}

	return false
}

// SetInstanceAddress gets a reference to the given string and assigns it to the InstanceAddress field.
func (o *OTCProcessSpansDroppedTracker) SetInstanceAddress(v string) {
	o.InstanceAddress = &v
}

// GetProcessorId returns the ProcessorId field value if set, zero value otherwise.
func (o *OTCProcessSpansDroppedTracker) GetProcessorId() string {
	if o == nil || IsNil(o.ProcessorId) {
		var ret string
		return ret
	}
	return *o.ProcessorId
}

// GetProcessorIdOk returns a tuple with the ProcessorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessSpansDroppedTracker) GetProcessorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorId) {
		return nil, false
	}
	return o.ProcessorId, true
}

// HasProcessorId returns a boolean if a field has been set.
func (o *OTCProcessSpansDroppedTracker) HasProcessorId() bool {
	if o != nil && !IsNil(o.ProcessorId) {
		return true
	}

	return false
}

// SetProcessorId gets a reference to the given string and assigns it to the ProcessorId field.
func (o *OTCProcessSpansDroppedTracker) SetProcessorId(v string) {
	o.ProcessorId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OTCProcessSpansDroppedTracker) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCProcessSpansDroppedTracker) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OTCProcessSpansDroppedTracker) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *OTCProcessSpansDroppedTracker) SetCount(v string) {
	o.Count = &v
}

func (o OTCProcessSpansDroppedTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCProcessSpansDroppedTracker) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ProcessorId) {
		toSerialize["processorId"] = o.ProcessorId
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableOTCProcessSpansDroppedTracker struct {
	value *OTCProcessSpansDroppedTracker
	isSet bool
}

func (v NullableOTCProcessSpansDroppedTracker) Get() *OTCProcessSpansDroppedTracker {
	return v.value
}

func (v *NullableOTCProcessSpansDroppedTracker) Set(val *OTCProcessSpansDroppedTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCProcessSpansDroppedTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCProcessSpansDroppedTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCProcessSpansDroppedTracker(val *OTCProcessSpansDroppedTracker) *NullableOTCProcessSpansDroppedTracker {
	return &NullableOTCProcessSpansDroppedTracker{value: val, isSet: true}
}

func (v NullableOTCProcessSpansDroppedTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCProcessSpansDroppedTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


