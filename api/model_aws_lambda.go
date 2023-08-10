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

// checks if the AWSLambda type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSLambda{}

// AWSLambda struct for AWSLambda
type AWSLambda struct {
	Action
	// The identifier of the connection.
	ConnectionId string `json:"connectionId"`
	// The override of the default JSON payload of the connection. Should be in JSON format.
	PayloadOverride *string `json:"payloadOverride,omitempty"`
	// The override of the resolution JSON payload of the connection. Should be in JSON format.
	ResolutionPayloadOverride *string `json:"resolutionPayloadOverride,omitempty"`
}

// NewAWSLambda instantiates a new AWSLambda object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSLambda(connectionId string, connectionType string) *AWSLambda {
	this := AWSLambda{}
	this.ConnectionType = connectionType
	this.ConnectionId = connectionId
	return &this
}

// NewAWSLambdaWithDefaults instantiates a new AWSLambda object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSLambdaWithDefaults() *AWSLambda {
	this := AWSLambda{}
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *AWSLambda) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *AWSLambda) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *AWSLambda) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetPayloadOverride returns the PayloadOverride field value if set, zero value otherwise.
func (o *AWSLambda) GetPayloadOverride() string {
	if o == nil || IsNil(o.PayloadOverride) {
		var ret string
		return ret
	}
	return *o.PayloadOverride
}

// GetPayloadOverrideOk returns a tuple with the PayloadOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLambda) GetPayloadOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadOverride) {
		return nil, false
	}
	return o.PayloadOverride, true
}

// HasPayloadOverride returns a boolean if a field has been set.
func (o *AWSLambda) HasPayloadOverride() bool {
	if o != nil && !IsNil(o.PayloadOverride) {
		return true
	}

	return false
}

// SetPayloadOverride gets a reference to the given string and assigns it to the PayloadOverride field.
func (o *AWSLambda) SetPayloadOverride(v string) {
	o.PayloadOverride = &v
}

// GetResolutionPayloadOverride returns the ResolutionPayloadOverride field value if set, zero value otherwise.
func (o *AWSLambda) GetResolutionPayloadOverride() string {
	if o == nil || IsNil(o.ResolutionPayloadOverride) {
		var ret string
		return ret
	}
	return *o.ResolutionPayloadOverride
}

// GetResolutionPayloadOverrideOk returns a tuple with the ResolutionPayloadOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLambda) GetResolutionPayloadOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ResolutionPayloadOverride) {
		return nil, false
	}
	return o.ResolutionPayloadOverride, true
}

// HasResolutionPayloadOverride returns a boolean if a field has been set.
func (o *AWSLambda) HasResolutionPayloadOverride() bool {
	if o != nil && !IsNil(o.ResolutionPayloadOverride) {
		return true
	}

	return false
}

// SetResolutionPayloadOverride gets a reference to the given string and assigns it to the ResolutionPayloadOverride field.
func (o *AWSLambda) SetResolutionPayloadOverride(v string) {
	o.ResolutionPayloadOverride = &v
}

func (o AWSLambda) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSLambda) ToMap() (map[string]interface{}, error) {
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

type NullableAWSLambda struct {
	value *AWSLambda
	isSet bool
}

func (v NullableAWSLambda) Get() *AWSLambda {
	return v.value
}

func (v *NullableAWSLambda) Set(val *AWSLambda) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSLambda) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSLambda) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSLambda(val *AWSLambda) *NullableAWSLambda {
	return &NullableAWSLambda{value: val, isSet: true}
}

func (v NullableAWSLambda) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSLambda) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


