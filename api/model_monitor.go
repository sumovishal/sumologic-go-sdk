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

// checks if the Monitor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Monitor{}

// Monitor SLI definition based on monitors.
type Monitor struct {
	Sli
	// Monitors over which the SLO is defined.
	MonitorTriggers []MonitorTrigger `json:"monitorTriggers"`
}

type _Monitor Monitor

// NewMonitor instantiates a new Monitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitor(monitorTriggers []MonitorTrigger, evaluationType string) *Monitor {
	this := Monitor{}
	this.EvaluationType = evaluationType
	this.MonitorTriggers = monitorTriggers
	return &this
}

// NewMonitorWithDefaults instantiates a new Monitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorWithDefaults() *Monitor {
	this := Monitor{}
	return &this
}

// GetMonitorTriggers returns the MonitorTriggers field value
func (o *Monitor) GetMonitorTriggers() []MonitorTrigger {
	if o == nil {
		var ret []MonitorTrigger
		return ret
	}

	return o.MonitorTriggers
}

// GetMonitorTriggersOk returns a tuple with the MonitorTriggers field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetMonitorTriggersOk() ([]MonitorTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitorTriggers, true
}

// SetMonitorTriggers sets field value
func (o *Monitor) SetMonitorTriggers(v []MonitorTrigger) {
	o.MonitorTriggers = v
}

func (o Monitor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Monitor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["monitorTriggers"] = o.MonitorTriggers
	return toSerialize, nil
}

func (o *Monitor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"monitorTriggers",
		"evaluationType",
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

	varMonitor := _Monitor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMonitor)

	if err != nil {
		return err
	}

	*o = Monitor(varMonitor)

	return err
}

type NullableMonitor struct {
	value *Monitor
	isSet bool
}

func (v NullableMonitor) Get() *Monitor {
	return v.value
}

func (v *NullableMonitor) Set(val *Monitor) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitor(val *Monitor) *NullableMonitor {
	return &NullableMonitor{value: val, isSet: true}
}

func (v NullableMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


