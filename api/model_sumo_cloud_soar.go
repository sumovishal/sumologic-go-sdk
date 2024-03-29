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

// checks if the SumoCloudSOAR type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SumoCloudSOAR{}

// SumoCloudSOAR struct for SumoCloudSOAR
type SumoCloudSOAR struct {
	Action
	// The identifier of the connection.
	ConnectionId string `json:"connectionId"`
	// The subtype of the connection. Valid values are `Event` or `Incident`.
	ConnectionSubtype *string `json:"connectionSubtype,omitempty"`
	// The override of the default JSON payload of the connection. Should be in JSON format.
	PayloadOverride *string `json:"payloadOverride,omitempty"`
}

// NewSumoCloudSOAR instantiates a new SumoCloudSOAR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSumoCloudSOAR(connectionId string, connectionType string) *SumoCloudSOAR {
	this := SumoCloudSOAR{}
	this.ConnectionType = connectionType
	this.ConnectionId = connectionId
	return &this
}

// NewSumoCloudSOARWithDefaults instantiates a new SumoCloudSOAR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSumoCloudSOARWithDefaults() *SumoCloudSOAR {
	this := SumoCloudSOAR{}
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *SumoCloudSOAR) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *SumoCloudSOAR) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *SumoCloudSOAR) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetConnectionSubtype returns the ConnectionSubtype field value if set, zero value otherwise.
func (o *SumoCloudSOAR) GetConnectionSubtype() string {
	if o == nil || IsNil(o.ConnectionSubtype) {
		var ret string
		return ret
	}
	return *o.ConnectionSubtype
}

// GetConnectionSubtypeOk returns a tuple with the ConnectionSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoCloudSOAR) GetConnectionSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionSubtype) {
		return nil, false
	}
	return o.ConnectionSubtype, true
}

// HasConnectionSubtype returns a boolean if a field has been set.
func (o *SumoCloudSOAR) HasConnectionSubtype() bool {
	if o != nil && !IsNil(o.ConnectionSubtype) {
		return true
	}

	return false
}

// SetConnectionSubtype gets a reference to the given string and assigns it to the ConnectionSubtype field.
func (o *SumoCloudSOAR) SetConnectionSubtype(v string) {
	o.ConnectionSubtype = &v
}

// GetPayloadOverride returns the PayloadOverride field value if set, zero value otherwise.
func (o *SumoCloudSOAR) GetPayloadOverride() string {
	if o == nil || IsNil(o.PayloadOverride) {
		var ret string
		return ret
	}
	return *o.PayloadOverride
}

// GetPayloadOverrideOk returns a tuple with the PayloadOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoCloudSOAR) GetPayloadOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadOverride) {
		return nil, false
	}
	return o.PayloadOverride, true
}

// HasPayloadOverride returns a boolean if a field has been set.
func (o *SumoCloudSOAR) HasPayloadOverride() bool {
	if o != nil && !IsNil(o.PayloadOverride) {
		return true
	}

	return false
}

// SetPayloadOverride gets a reference to the given string and assigns it to the PayloadOverride field.
func (o *SumoCloudSOAR) SetPayloadOverride(v string) {
	o.PayloadOverride = &v
}

func (o SumoCloudSOAR) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SumoCloudSOAR) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConnectionSubtype) {
		toSerialize["connectionSubtype"] = o.ConnectionSubtype
	}
	if !IsNil(o.PayloadOverride) {
		toSerialize["payloadOverride"] = o.PayloadOverride
	}
	return toSerialize, nil
}

type NullableSumoCloudSOAR struct {
	value *SumoCloudSOAR
	isSet bool
}

func (v NullableSumoCloudSOAR) Get() *SumoCloudSOAR {
	return v.value
}

func (v *NullableSumoCloudSOAR) Set(val *SumoCloudSOAR) {
	v.value = val
	v.isSet = true
}

func (v NullableSumoCloudSOAR) IsSet() bool {
	return v.isSet
}

func (v *NullableSumoCloudSOAR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSumoCloudSOAR(val *SumoCloudSOAR) *NullableSumoCloudSOAR {
	return &NullableSumoCloudSOAR{value: val, isSet: true}
}

func (v NullableSumoCloudSOAR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSumoCloudSOAR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


