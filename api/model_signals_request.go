/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// SignalsRequest Signal Request object.
type SignalsRequest struct {
	// A list of signal types to compute. Can be `LogFluctuation`, `DimensionalityExplanation`, `GisBenchmark` or `Anomalies` 
	SignalTypes []string `json:"signalTypes"`
	SignalContext SignalContext `json:"signalContext"`
}

// NewSignalsRequest instantiates a new SignalsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalsRequest(signalTypes []string, signalContext SignalContext) *SignalsRequest {
	this := SignalsRequest{}
	this.SignalTypes = signalTypes
	this.SignalContext = signalContext
	return &this
}

// NewSignalsRequestWithDefaults instantiates a new SignalsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalsRequestWithDefaults() *SignalsRequest {
	this := SignalsRequest{}
	return &this
}

// GetSignalTypes returns the SignalTypes field value
func (o *SignalsRequest) GetSignalTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SignalTypes
}

// GetSignalTypesOk returns a tuple with the SignalTypes field value
// and a boolean to check if the value has been set.
func (o *SignalsRequest) GetSignalTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignalTypes, true
}

// SetSignalTypes sets field value
func (o *SignalsRequest) SetSignalTypes(v []string) {
	o.SignalTypes = v
}

// GetSignalContext returns the SignalContext field value
func (o *SignalsRequest) GetSignalContext() SignalContext {
	if o == nil {
		var ret SignalContext
		return ret
	}

	return o.SignalContext
}

// GetSignalContextOk returns a tuple with the SignalContext field value
// and a boolean to check if the value has been set.
func (o *SignalsRequest) GetSignalContextOk() (*SignalContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalContext, true
}

// SetSignalContext sets field value
func (o *SignalsRequest) SetSignalContext(v SignalContext) {
	o.SignalContext = v
}

func (o SignalsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["signalTypes"] = o.SignalTypes
	}
	if true {
		toSerialize["signalContext"] = o.SignalContext
	}
	return json.Marshal(toSerialize)
}

type NullableSignalsRequest struct {
	value *SignalsRequest
	isSet bool
}

func (v NullableSignalsRequest) Get() *SignalsRequest {
	return v.value
}

func (v *NullableSignalsRequest) Set(val *SignalsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalsRequest(val *SignalsRequest) *NullableSignalsRequest {
	return &NullableSignalsRequest{value: val, isSet: true}
}

func (v NullableSignalsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


