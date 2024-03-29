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

// checks if the SharedBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedBucket{}

// SharedBucket A shared bucket contains capacities which can be used my multiple entitlements which are linked to the bucket. There will be a 1:many mapping between SharedBucket:Entitlement.
type SharedBucket struct {
	// Name of the bucket.
	Name string `json:"name"`
	// The text to be displayed on UI for this bucket.
	Label string `json:"label"`
	// List of entitlement types which can consume from this bucket.
	LinkedEntitlementTypes []string `json:"linkedEntitlementTypes"`
	// List of capacities alloted.
	Capacitites []Capacity `json:"capacitites,omitempty"`
}

// NewSharedBucket instantiates a new SharedBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedBucket(name string, label string, linkedEntitlementTypes []string) *SharedBucket {
	this := SharedBucket{}
	this.Name = name
	this.Label = label
	this.LinkedEntitlementTypes = linkedEntitlementTypes
	return &this
}

// NewSharedBucketWithDefaults instantiates a new SharedBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedBucketWithDefaults() *SharedBucket {
	this := SharedBucket{}
	return &this
}

// GetName returns the Name field value
func (o *SharedBucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SharedBucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SharedBucket) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *SharedBucket) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SharedBucket) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SharedBucket) SetLabel(v string) {
	o.Label = v
}

// GetLinkedEntitlementTypes returns the LinkedEntitlementTypes field value
func (o *SharedBucket) GetLinkedEntitlementTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LinkedEntitlementTypes
}

// GetLinkedEntitlementTypesOk returns a tuple with the LinkedEntitlementTypes field value
// and a boolean to check if the value has been set.
func (o *SharedBucket) GetLinkedEntitlementTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedEntitlementTypes, true
}

// SetLinkedEntitlementTypes sets field value
func (o *SharedBucket) SetLinkedEntitlementTypes(v []string) {
	o.LinkedEntitlementTypes = v
}

// GetCapacitites returns the Capacitites field value if set, zero value otherwise.
func (o *SharedBucket) GetCapacitites() []Capacity {
	if o == nil || IsNil(o.Capacitites) {
		var ret []Capacity
		return ret
	}
	return o.Capacitites
}

// GetCapacititesOk returns a tuple with the Capacitites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedBucket) GetCapacititesOk() ([]Capacity, bool) {
	if o == nil || IsNil(o.Capacitites) {
		return nil, false
	}
	return o.Capacitites, true
}

// HasCapacitites returns a boolean if a field has been set.
func (o *SharedBucket) HasCapacitites() bool {
	if o != nil && !IsNil(o.Capacitites) {
		return true
	}

	return false
}

// SetCapacitites gets a reference to the given []Capacity and assigns it to the Capacitites field.
func (o *SharedBucket) SetCapacitites(v []Capacity) {
	o.Capacitites = v
}

func (o SharedBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["label"] = o.Label
	toSerialize["linkedEntitlementTypes"] = o.LinkedEntitlementTypes
	if !IsNil(o.Capacitites) {
		toSerialize["capacitites"] = o.Capacitites
	}
	return toSerialize, nil
}

type NullableSharedBucket struct {
	value *SharedBucket
	isSet bool
}

func (v NullableSharedBucket) Get() *SharedBucket {
	return v.value
}

func (v *NullableSharedBucket) Set(val *SharedBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedBucket(val *SharedBucket) *NullableSharedBucket {
	return &NullableSharedBucket{value: val, isSet: true}
}

func (v NullableSharedBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


