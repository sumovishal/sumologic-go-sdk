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

// checks if the BucketDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketDefinitionAllOf{}

// BucketDefinitionAllOf struct for BucketDefinitionAllOf
type BucketDefinitionAllOf struct {
	// The unique identifier of the data forwarding destination.
	Id string `json:"id"`
	// True if invalidated by the system.
	InvalidatedBySystem *bool `json:"invalidatedBySystem,omitempty"`
}

// NewBucketDefinitionAllOf instantiates a new BucketDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketDefinitionAllOf(id string) *BucketDefinitionAllOf {
	this := BucketDefinitionAllOf{}
	this.Id = id
	return &this
}

// NewBucketDefinitionAllOfWithDefaults instantiates a new BucketDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketDefinitionAllOfWithDefaults() *BucketDefinitionAllOf {
	this := BucketDefinitionAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *BucketDefinitionAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BucketDefinitionAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BucketDefinitionAllOf) SetId(v string) {
	o.Id = v
}

// GetInvalidatedBySystem returns the InvalidatedBySystem field value if set, zero value otherwise.
func (o *BucketDefinitionAllOf) GetInvalidatedBySystem() bool {
	if o == nil || IsNil(o.InvalidatedBySystem) {
		var ret bool
		return ret
	}
	return *o.InvalidatedBySystem
}

// GetInvalidatedBySystemOk returns a tuple with the InvalidatedBySystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinitionAllOf) GetInvalidatedBySystemOk() (*bool, bool) {
	if o == nil || IsNil(o.InvalidatedBySystem) {
		return nil, false
	}
	return o.InvalidatedBySystem, true
}

// HasInvalidatedBySystem returns a boolean if a field has been set.
func (o *BucketDefinitionAllOf) HasInvalidatedBySystem() bool {
	if o != nil && !IsNil(o.InvalidatedBySystem) {
		return true
	}

	return false
}

// SetInvalidatedBySystem gets a reference to the given bool and assigns it to the InvalidatedBySystem field.
func (o *BucketDefinitionAllOf) SetInvalidatedBySystem(v bool) {
	o.InvalidatedBySystem = &v
}

func (o BucketDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.InvalidatedBySystem) {
		toSerialize["invalidatedBySystem"] = o.InvalidatedBySystem
	}
	return toSerialize, nil
}

type NullableBucketDefinitionAllOf struct {
	value *BucketDefinitionAllOf
	isSet bool
}

func (v NullableBucketDefinitionAllOf) Get() *BucketDefinitionAllOf {
	return v.value
}

func (v *NullableBucketDefinitionAllOf) Set(val *BucketDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketDefinitionAllOf(val *BucketDefinitionAllOf) *NullableBucketDefinitionAllOf {
	return &NullableBucketDefinitionAllOf{value: val, isSet: true}
}

func (v NullableBucketDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


