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

// checks if the MonitorScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorScope{}

// MonitorScope Monitor scope that the schedule applies to
type MonitorScope struct {
	// List of monitor Ids in hex. Must be empty if `all` is true.
	Ids []string `json:"ids,omitempty"`
	// true if the schedule applies to all monitors
	All *bool `json:"all,omitempty"`
}

// NewMonitorScope instantiates a new MonitorScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorScope() *MonitorScope {
	this := MonitorScope{}
	var all bool = false
	this.All = &all
	return &this
}

// NewMonitorScopeWithDefaults instantiates a new MonitorScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorScopeWithDefaults() *MonitorScope {
	this := MonitorScope{}
	var all bool = false
	this.All = &all
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *MonitorScope) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorScope) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *MonitorScope) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *MonitorScope) SetIds(v []string) {
	o.Ids = v
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *MonitorScope) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorScope) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *MonitorScope) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *MonitorScope) SetAll(v bool) {
	o.All = &v
}

func (o MonitorScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	return toSerialize, nil
}

type NullableMonitorScope struct {
	value *MonitorScope
	isSet bool
}

func (v NullableMonitorScope) Get() *MonitorScope {
	return v.value
}

func (v *NullableMonitorScope) Set(val *MonitorScope) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorScope) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorScope(val *MonitorScope) *NullableMonitorScope {
	return &NullableMonitorScope{value: val, isSet: true}
}

func (v NullableMonitorScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


