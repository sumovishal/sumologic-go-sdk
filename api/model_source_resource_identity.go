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

// SourceResourceIdentity struct for SourceResourceIdentity
type SourceResourceIdentity struct {
	ResourceIdentity
	// The unique identifier of the Collector this Source belongs to.
	CollectorId *string `json:"collectorId,omitempty"`
	// The name of the Collector this Source belongs to.
	CollectorName *string `json:"collectorName,omitempty"`
}

// NewSourceResourceIdentity instantiates a new SourceResourceIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceResourceIdentity(id string, type_ string) *SourceResourceIdentity {
	this := SourceResourceIdentity{}
	this.Id = id
	var name string = "Unknown"
	this.Name = &name
	this.Type = type_
	var collectorId string = "Unknown"
	this.CollectorId = &collectorId
	var collectorName string = "Unknown"
	this.CollectorName = &collectorName
	return &this
}

// NewSourceResourceIdentityWithDefaults instantiates a new SourceResourceIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceResourceIdentityWithDefaults() *SourceResourceIdentity {
	this := SourceResourceIdentity{}
	var collectorId string = "Unknown"
	this.CollectorId = &collectorId
	var collectorName string = "Unknown"
	this.CollectorName = &collectorName
	return &this
}

// GetCollectorId returns the CollectorId field value if set, zero value otherwise.
func (o *SourceResourceIdentity) GetCollectorId() string {
	if o == nil || o.CollectorId == nil {
		var ret string
		return ret
	}
	return *o.CollectorId
}

// GetCollectorIdOk returns a tuple with the CollectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceResourceIdentity) GetCollectorIdOk() (*string, bool) {
	if o == nil || o.CollectorId == nil {
		return nil, false
	}
	return o.CollectorId, true
}

// HasCollectorId returns a boolean if a field has been set.
func (o *SourceResourceIdentity) HasCollectorId() bool {
	if o != nil && o.CollectorId != nil {
		return true
	}

	return false
}

// SetCollectorId gets a reference to the given string and assigns it to the CollectorId field.
func (o *SourceResourceIdentity) SetCollectorId(v string) {
	o.CollectorId = &v
}

// GetCollectorName returns the CollectorName field value if set, zero value otherwise.
func (o *SourceResourceIdentity) GetCollectorName() string {
	if o == nil || o.CollectorName == nil {
		var ret string
		return ret
	}
	return *o.CollectorName
}

// GetCollectorNameOk returns a tuple with the CollectorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceResourceIdentity) GetCollectorNameOk() (*string, bool) {
	if o == nil || o.CollectorName == nil {
		return nil, false
	}
	return o.CollectorName, true
}

// HasCollectorName returns a boolean if a field has been set.
func (o *SourceResourceIdentity) HasCollectorName() bool {
	if o != nil && o.CollectorName != nil {
		return true
	}

	return false
}

// SetCollectorName gets a reference to the given string and assigns it to the CollectorName field.
func (o *SourceResourceIdentity) SetCollectorName(v string) {
	o.CollectorName = &v
}

func (o SourceResourceIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedResourceIdentity, errResourceIdentity := json.Marshal(o.ResourceIdentity)
	if errResourceIdentity != nil {
		return []byte{}, errResourceIdentity
	}
	errResourceIdentity = json.Unmarshal([]byte(serializedResourceIdentity), &toSerialize)
	if errResourceIdentity != nil {
		return []byte{}, errResourceIdentity
	}
	if o.CollectorId != nil {
		toSerialize["collectorId"] = o.CollectorId
	}
	if o.CollectorName != nil {
		toSerialize["collectorName"] = o.CollectorName
	}
	return json.Marshal(toSerialize)
}

type NullableSourceResourceIdentity struct {
	value *SourceResourceIdentity
	isSet bool
}

func (v NullableSourceResourceIdentity) Get() *SourceResourceIdentity {
	return v.value
}

func (v *NullableSourceResourceIdentity) Set(val *SourceResourceIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceResourceIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceResourceIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceResourceIdentity(val *SourceResourceIdentity) *NullableSourceResourceIdentity {
	return &NullableSourceResourceIdentity{value: val, isSet: true}
}

func (v NullableSourceResourceIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceResourceIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


