/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DataForwardingRuleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataForwardingRuleAllOf{}

// DataForwardingRuleAllOf struct for DataForwardingRuleAllOf
type DataForwardingRuleAllOf struct {
	// The unique identifier of the data forwarding rule.
	Id *string `json:"id,omitempty"`
}

// NewDataForwardingRuleAllOf instantiates a new DataForwardingRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataForwardingRuleAllOf() *DataForwardingRuleAllOf {
	this := DataForwardingRuleAllOf{}
	return &this
}

// NewDataForwardingRuleAllOfWithDefaults instantiates a new DataForwardingRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataForwardingRuleAllOfWithDefaults() *DataForwardingRuleAllOf {
	this := DataForwardingRuleAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataForwardingRuleAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataForwardingRuleAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataForwardingRuleAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataForwardingRuleAllOf) SetId(v string) {
	o.Id = &v
}

func (o DataForwardingRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataForwardingRuleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDataForwardingRuleAllOf struct {
	value *DataForwardingRuleAllOf
	isSet bool
}

func (v NullableDataForwardingRuleAllOf) Get() *DataForwardingRuleAllOf {
	return v.value
}

func (v *NullableDataForwardingRuleAllOf) Set(val *DataForwardingRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataForwardingRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataForwardingRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataForwardingRuleAllOf(val *DataForwardingRuleAllOf) *NullableDataForwardingRuleAllOf {
	return &NullableDataForwardingRuleAllOf{value: val, isSet: true}
}

func (v NullableDataForwardingRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataForwardingRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


