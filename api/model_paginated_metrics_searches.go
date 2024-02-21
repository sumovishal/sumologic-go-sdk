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

// checks if the PaginatedMetricsSearches type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedMetricsSearches{}

// PaginatedMetricsSearches struct for PaginatedMetricsSearches
type PaginatedMetricsSearches struct {
	// List of metrics search pages.
	MetricsSearches []MetricsSearch `json:"metricsSearches"`
	// Next continuation token. `token` is set to null when no more pages are left.
	Next *string `json:"next,omitempty"`
}

// NewPaginatedMetricsSearches instantiates a new PaginatedMetricsSearches object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedMetricsSearches(metricsSearches []MetricsSearch) *PaginatedMetricsSearches {
	this := PaginatedMetricsSearches{}
	this.MetricsSearches = metricsSearches
	return &this
}

// NewPaginatedMetricsSearchesWithDefaults instantiates a new PaginatedMetricsSearches object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedMetricsSearchesWithDefaults() *PaginatedMetricsSearches {
	this := PaginatedMetricsSearches{}
	return &this
}

// GetMetricsSearches returns the MetricsSearches field value
func (o *PaginatedMetricsSearches) GetMetricsSearches() []MetricsSearch {
	if o == nil {
		var ret []MetricsSearch
		return ret
	}

	return o.MetricsSearches
}

// GetMetricsSearchesOk returns a tuple with the MetricsSearches field value
// and a boolean to check if the value has been set.
func (o *PaginatedMetricsSearches) GetMetricsSearchesOk() ([]MetricsSearch, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricsSearches, true
}

// SetMetricsSearches sets field value
func (o *PaginatedMetricsSearches) SetMetricsSearches(v []MetricsSearch) {
	o.MetricsSearches = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedMetricsSearches) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedMetricsSearches) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedMetricsSearches) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedMetricsSearches) SetNext(v string) {
	o.Next = &v
}

func (o PaginatedMetricsSearches) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedMetricsSearches) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricsSearches"] = o.MetricsSearches
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullablePaginatedMetricsSearches struct {
	value *PaginatedMetricsSearches
	isSet bool
}

func (v NullablePaginatedMetricsSearches) Get() *PaginatedMetricsSearches {
	return v.value
}

func (v *NullablePaginatedMetricsSearches) Set(val *PaginatedMetricsSearches) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedMetricsSearches) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedMetricsSearches) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedMetricsSearches(val *PaginatedMetricsSearches) *NullablePaginatedMetricsSearches {
	return &NullablePaginatedMetricsSearches{value: val, isSet: true}
}

func (v NullablePaginatedMetricsSearches) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedMetricsSearches) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


