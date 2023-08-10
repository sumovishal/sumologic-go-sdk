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

// checks if the TraceMessageBusSpanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceMessageBusSpanInfo{}

// TraceMessageBusSpanInfo struct for TraceMessageBusSpanInfo
type TraceMessageBusSpanInfo struct {
	TraceSpanInfo
	// An address at which messages can be exchanged e.g. a Kafka record has an associated \"topic name\" that can be stored using this tag.
	Destination *string `json:"destination,omitempty"`
}

// NewTraceMessageBusSpanInfo instantiates a new TraceMessageBusSpanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceMessageBusSpanInfo(type_ string) *TraceMessageBusSpanInfo {
	this := TraceMessageBusSpanInfo{}
	this.Type = type_
	return &this
}

// NewTraceMessageBusSpanInfoWithDefaults instantiates a new TraceMessageBusSpanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceMessageBusSpanInfoWithDefaults() *TraceMessageBusSpanInfo {
	this := TraceMessageBusSpanInfo{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *TraceMessageBusSpanInfo) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceMessageBusSpanInfo) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *TraceMessageBusSpanInfo) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *TraceMessageBusSpanInfo) SetDestination(v string) {
	o.Destination = &v
}

func (o TraceMessageBusSpanInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceMessageBusSpanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTraceSpanInfo, errTraceSpanInfo := json.Marshal(o.TraceSpanInfo)
	if errTraceSpanInfo != nil {
		return map[string]interface{}{}, errTraceSpanInfo
	}
	errTraceSpanInfo = json.Unmarshal([]byte(serializedTraceSpanInfo), &toSerialize)
	if errTraceSpanInfo != nil {
		return map[string]interface{}{}, errTraceSpanInfo
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return toSerialize, nil
}

type NullableTraceMessageBusSpanInfo struct {
	value *TraceMessageBusSpanInfo
	isSet bool
}

func (v NullableTraceMessageBusSpanInfo) Get() *TraceMessageBusSpanInfo {
	return v.value
}

func (v *NullableTraceMessageBusSpanInfo) Set(val *TraceMessageBusSpanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceMessageBusSpanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceMessageBusSpanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceMessageBusSpanInfo(val *TraceMessageBusSpanInfo) *NullableTraceMessageBusSpanInfo {
	return &NullableTraceMessageBusSpanInfo{value: val, isSet: true}
}

func (v NullableTraceMessageBusSpanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceMessageBusSpanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


