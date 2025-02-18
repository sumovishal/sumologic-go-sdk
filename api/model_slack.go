/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Slack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Slack{}

// Slack struct for Slack
type Slack struct {
	Action
	// The identifier of the connection.
	ConnectionId string `json:"connectionId"`
	// The override of the default JSON payload of the connection. Should be in JSON format.
	PayloadOverride *string `json:"payloadOverride,omitempty"`
	// The override of the resolution JSON payload of the connection. Should be in JSON format.
	ResolutionPayloadOverride *string `json:"resolutionPayloadOverride,omitempty"`
}

type _Slack Slack

// NewSlack instantiates a new Slack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlack(connectionId string, connectionType string) *Slack {
	this := Slack{}
	this.ConnectionType = connectionType
	this.ConnectionId = connectionId
	return &this
}

// NewSlackWithDefaults instantiates a new Slack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackWithDefaults() *Slack {
	this := Slack{}
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *Slack) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *Slack) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *Slack) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetPayloadOverride returns the PayloadOverride field value if set, zero value otherwise.
func (o *Slack) GetPayloadOverride() string {
	if o == nil || IsNil(o.PayloadOverride) {
		var ret string
		return ret
	}
	return *o.PayloadOverride
}

// GetPayloadOverrideOk returns a tuple with the PayloadOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slack) GetPayloadOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadOverride) {
		return nil, false
	}
	return o.PayloadOverride, true
}

// HasPayloadOverride returns a boolean if a field has been set.
func (o *Slack) HasPayloadOverride() bool {
	if o != nil && !IsNil(o.PayloadOverride) {
		return true
	}

	return false
}

// SetPayloadOverride gets a reference to the given string and assigns it to the PayloadOverride field.
func (o *Slack) SetPayloadOverride(v string) {
	o.PayloadOverride = &v
}

// GetResolutionPayloadOverride returns the ResolutionPayloadOverride field value if set, zero value otherwise.
func (o *Slack) GetResolutionPayloadOverride() string {
	if o == nil || IsNil(o.ResolutionPayloadOverride) {
		var ret string
		return ret
	}
	return *o.ResolutionPayloadOverride
}

// GetResolutionPayloadOverrideOk returns a tuple with the ResolutionPayloadOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slack) GetResolutionPayloadOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ResolutionPayloadOverride) {
		return nil, false
	}
	return o.ResolutionPayloadOverride, true
}

// HasResolutionPayloadOverride returns a boolean if a field has been set.
func (o *Slack) HasResolutionPayloadOverride() bool {
	if o != nil && !IsNil(o.ResolutionPayloadOverride) {
		return true
	}

	return false
}

// SetResolutionPayloadOverride gets a reference to the given string and assigns it to the ResolutionPayloadOverride field.
func (o *Slack) SetResolutionPayloadOverride(v string) {
	o.ResolutionPayloadOverride = &v
}

func (o Slack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Slack) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ResolutionPayloadOverride) {
		toSerialize["resolutionPayloadOverride"] = o.ResolutionPayloadOverride
	}
	return toSerialize, nil
}

func (o *Slack) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionId",
		"connectionType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSlack := _Slack{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSlack)

	if err != nil {
		return err
	}

	*o = Slack(varSlack)

	return err
}

type NullableSlack struct {
	value *Slack
	isSet bool
}

func (v NullableSlack) Get() *Slack {
	return v.value
}

func (v *NullableSlack) Set(val *Slack) {
	v.value = val
	v.isSet = true
}

func (v NullableSlack) IsSet() bool {
	return v.isSet
}

func (v *NullableSlack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlack(val *Slack) *NullableSlack {
	return &NullableSlack{value: val, isSet: true}
}

func (v NullableSlack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


