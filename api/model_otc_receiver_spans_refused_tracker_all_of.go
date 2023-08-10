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

// checks if the OTCReceiverSpansRefusedTrackerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCReceiverSpansRefusedTrackerAllOf{}

// OTCReceiverSpansRefusedTrackerAllOf struct for OTCReceiverSpansRefusedTrackerAllOf
type OTCReceiverSpansRefusedTrackerAllOf struct {
	// The collector instance ID, e.g. `974b444b-4b45-4f32-aa03-1dbf2a16826d`.
	InstanceId *string `json:"instanceId,omitempty"`
	// The collector instance address, e.g. `172.16.1.14`.
	InstanceAddress *string `json:"instanceAddress,omitempty"`
	// The collector receiver ID, e.g. `otlphttp/2`.
	ReceiverId *string `json:"receiverId,omitempty"`
	// The count of refused spans.
	Count *string `json:"count,omitempty"`
}

// NewOTCReceiverSpansRefusedTrackerAllOf instantiates a new OTCReceiverSpansRefusedTrackerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCReceiverSpansRefusedTrackerAllOf() *OTCReceiverSpansRefusedTrackerAllOf {
	this := OTCReceiverSpansRefusedTrackerAllOf{}
	return &this
}

// NewOTCReceiverSpansRefusedTrackerAllOfWithDefaults instantiates a new OTCReceiverSpansRefusedTrackerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCReceiverSpansRefusedTrackerAllOfWithDefaults() *OTCReceiverSpansRefusedTrackerAllOf {
	this := OTCReceiverSpansRefusedTrackerAllOf{}
	return &this
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *OTCReceiverSpansRefusedTrackerAllOf) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverSpansRefusedTrackerAllOf) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *OTCReceiverSpansRefusedTrackerAllOf) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *OTCReceiverSpansRefusedTrackerAllOf) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceAddress returns the InstanceAddress field value if set, zero value otherwise.
func (o *OTCReceiverSpansRefusedTrackerAllOf) GetInstanceAddress() string {
	if o == nil || IsNil(o.InstanceAddress) {
		var ret string
		return ret
	}
	return *o.InstanceAddress
}

// GetInstanceAddressOk returns a tuple with the InstanceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverSpansRefusedTrackerAllOf) GetInstanceAddressOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceAddress) {
		return nil, false
	}
	return o.InstanceAddress, true
}

// HasInstanceAddress returns a boolean if a field has been set.
func (o *OTCReceiverSpansRefusedTrackerAllOf) HasInstanceAddress() bool {
	if o != nil && !IsNil(o.InstanceAddress) {
		return true
	}

	return false
}

// SetInstanceAddress gets a reference to the given string and assigns it to the InstanceAddress field.
func (o *OTCReceiverSpansRefusedTrackerAllOf) SetInstanceAddress(v string) {
	o.InstanceAddress = &v
}

// GetReceiverId returns the ReceiverId field value if set, zero value otherwise.
func (o *OTCReceiverSpansRefusedTrackerAllOf) GetReceiverId() string {
	if o == nil || IsNil(o.ReceiverId) {
		var ret string
		return ret
	}
	return *o.ReceiverId
}

// GetReceiverIdOk returns a tuple with the ReceiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverSpansRefusedTrackerAllOf) GetReceiverIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiverId) {
		return nil, false
	}
	return o.ReceiverId, true
}

// HasReceiverId returns a boolean if a field has been set.
func (o *OTCReceiverSpansRefusedTrackerAllOf) HasReceiverId() bool {
	if o != nil && !IsNil(o.ReceiverId) {
		return true
	}

	return false
}

// SetReceiverId gets a reference to the given string and assigns it to the ReceiverId field.
func (o *OTCReceiverSpansRefusedTrackerAllOf) SetReceiverId(v string) {
	o.ReceiverId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OTCReceiverSpansRefusedTrackerAllOf) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCReceiverSpansRefusedTrackerAllOf) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OTCReceiverSpansRefusedTrackerAllOf) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *OTCReceiverSpansRefusedTrackerAllOf) SetCount(v string) {
	o.Count = &v
}

func (o OTCReceiverSpansRefusedTrackerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCReceiverSpansRefusedTrackerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.InstanceAddress) {
		toSerialize["instanceAddress"] = o.InstanceAddress
	}
	if !IsNil(o.ReceiverId) {
		toSerialize["receiverId"] = o.ReceiverId
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableOTCReceiverSpansRefusedTrackerAllOf struct {
	value *OTCReceiverSpansRefusedTrackerAllOf
	isSet bool
}

func (v NullableOTCReceiverSpansRefusedTrackerAllOf) Get() *OTCReceiverSpansRefusedTrackerAllOf {
	return v.value
}

func (v *NullableOTCReceiverSpansRefusedTrackerAllOf) Set(val *OTCReceiverSpansRefusedTrackerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCReceiverSpansRefusedTrackerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCReceiverSpansRefusedTrackerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCReceiverSpansRefusedTrackerAllOf(val *OTCReceiverSpansRefusedTrackerAllOf) *NullableOTCReceiverSpansRefusedTrackerAllOf {
	return &NullableOTCReceiverSpansRefusedTrackerAllOf{value: val, isSet: true}
}

func (v NullableOTCReceiverSpansRefusedTrackerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCReceiverSpansRefusedTrackerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


