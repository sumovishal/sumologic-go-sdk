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

// checks if the KillChainPhase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KillChainPhase{}

// KillChainPhase struct for KillChainPhase
type KillChainPhase struct {
	// The name of the kill chain. The value of this property SHOULD be all lowercase and SHOULD use hyphens instead of spaces or underscores as word separators
	KillChainName string `json:"kill_chain_name"`
	// The name of the phase in the kill chain. The value of this property SHOULD be all lowercase and SHOULD use hyphens instead of spaces or underscores as word separators
	PhaseName *string `json:"phase_name,omitempty"`
}

// NewKillChainPhase instantiates a new KillChainPhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKillChainPhase(killChainName string) *KillChainPhase {
	this := KillChainPhase{}
	this.KillChainName = killChainName
	return &this
}

// NewKillChainPhaseWithDefaults instantiates a new KillChainPhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKillChainPhaseWithDefaults() *KillChainPhase {
	this := KillChainPhase{}
	return &this
}

// GetKillChainName returns the KillChainName field value
func (o *KillChainPhase) GetKillChainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KillChainName
}

// GetKillChainNameOk returns a tuple with the KillChainName field value
// and a boolean to check if the value has been set.
func (o *KillChainPhase) GetKillChainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KillChainName, true
}

// SetKillChainName sets field value
func (o *KillChainPhase) SetKillChainName(v string) {
	o.KillChainName = v
}

// GetPhaseName returns the PhaseName field value if set, zero value otherwise.
func (o *KillChainPhase) GetPhaseName() string {
	if o == nil || IsNil(o.PhaseName) {
		var ret string
		return ret
	}
	return *o.PhaseName
}

// GetPhaseNameOk returns a tuple with the PhaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KillChainPhase) GetPhaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.PhaseName) {
		return nil, false
	}
	return o.PhaseName, true
}

// HasPhaseName returns a boolean if a field has been set.
func (o *KillChainPhase) HasPhaseName() bool {
	if o != nil && !IsNil(o.PhaseName) {
		return true
	}

	return false
}

// SetPhaseName gets a reference to the given string and assigns it to the PhaseName field.
func (o *KillChainPhase) SetPhaseName(v string) {
	o.PhaseName = &v
}

func (o KillChainPhase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KillChainPhase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kill_chain_name"] = o.KillChainName
	if !IsNil(o.PhaseName) {
		toSerialize["phase_name"] = o.PhaseName
	}
	return toSerialize, nil
}

type NullableKillChainPhase struct {
	value *KillChainPhase
	isSet bool
}

func (v NullableKillChainPhase) Get() *KillChainPhase {
	return v.value
}

func (v *NullableKillChainPhase) Set(val *KillChainPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableKillChainPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableKillChainPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKillChainPhase(val *KillChainPhase) *NullableKillChainPhase {
	return &NullableKillChainPhase{value: val, isSet: true}
}

func (v NullableKillChainPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKillChainPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


