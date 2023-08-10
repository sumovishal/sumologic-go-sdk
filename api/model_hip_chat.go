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

// checks if the HipChat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HipChat{}

// HipChat struct for HipChat
type HipChat struct {
	Action
	// The identifier of the connection.
	ConnectionId string `json:"connectionId"`
	// The override of the default JSON payload of the connection. Should be in JSON format.
	PayloadOverride *string `json:"payloadOverride,omitempty"`
}

// NewHipChat instantiates a new HipChat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHipChat(connectionId string, connectionType string) *HipChat {
	this := HipChat{}
	this.ConnectionType = connectionType
	this.ConnectionId = connectionId
	return &this
}

// NewHipChatWithDefaults instantiates a new HipChat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHipChatWithDefaults() *HipChat {
	this := HipChat{}
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *HipChat) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *HipChat) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *HipChat) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetPayloadOverride returns the PayloadOverride field value if set, zero value otherwise.
func (o *HipChat) GetPayloadOverride() string {
	if o == nil || IsNil(o.PayloadOverride) {
		var ret string
		return ret
	}
	return *o.PayloadOverride
}

// GetPayloadOverrideOk returns a tuple with the PayloadOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HipChat) GetPayloadOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadOverride) {
		return nil, false
	}
	return o.PayloadOverride, true
}

// HasPayloadOverride returns a boolean if a field has been set.
func (o *HipChat) HasPayloadOverride() bool {
	if o != nil && !IsNil(o.PayloadOverride) {
		return true
	}

	return false
}

// SetPayloadOverride gets a reference to the given string and assigns it to the PayloadOverride field.
func (o *HipChat) SetPayloadOverride(v string) {
	o.PayloadOverride = &v
}

func (o HipChat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HipChat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAction, errAction := json.Marshal(o.Action)
	if errAction != nil {
		return map[string]interface{}{}, errAction
	}
	errAction = json.Unmarshal([]byte(serializedAction), &toSerialize)
	if errAction != nil {
		return map[string]interface{}{}, errAction
	}
	toSerialize["connectionId"] = o.ConnectionId
	if !IsNil(o.PayloadOverride) {
		toSerialize["payloadOverride"] = o.PayloadOverride
	}
	return toSerialize, nil
}

type NullableHipChat struct {
	value *HipChat
	isSet bool
}

func (v NullableHipChat) Get() *HipChat {
	return v.value
}

func (v *NullableHipChat) Set(val *HipChat) {
	v.value = val
	v.isSet = true
}

func (v NullableHipChat) IsSet() bool {
	return v.isSet
}

func (v *NullableHipChat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHipChat(val *HipChat) *NullableHipChat {
	return &NullableHipChat{value: val, isSet: true}
}

func (v NullableHipChat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHipChat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


