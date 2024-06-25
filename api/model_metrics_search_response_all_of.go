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

// checks if the MetricsSearchResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsSearchResponseAllOf{}

// MetricsSearchResponseAllOf struct for MetricsSearchResponseAllOf
type MetricsSearchResponseAllOf struct {
	// Unique identifier for the metrics search page.
	Id *string `json:"id,omitempty"`
}

// NewMetricsSearchResponseAllOf instantiates a new MetricsSearchResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsSearchResponseAllOf() *MetricsSearchResponseAllOf {
	this := MetricsSearchResponseAllOf{}
	return &this
}

// NewMetricsSearchResponseAllOfWithDefaults instantiates a new MetricsSearchResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsSearchResponseAllOfWithDefaults() *MetricsSearchResponseAllOf {
	this := MetricsSearchResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsSearchResponseAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsSearchResponseAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsSearchResponseAllOf) SetId(v string) {
	o.Id = &v
}

func (o MetricsSearchResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsSearchResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableMetricsSearchResponseAllOf struct {
	value *MetricsSearchResponseAllOf
	isSet bool
}

func (v NullableMetricsSearchResponseAllOf) Get() *MetricsSearchResponseAllOf {
	return v.value
}

func (v *NullableMetricsSearchResponseAllOf) Set(val *MetricsSearchResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsSearchResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsSearchResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsSearchResponseAllOf(val *MetricsSearchResponseAllOf) *NullableMetricsSearchResponseAllOf {
	return &NullableMetricsSearchResponseAllOf{value: val, isSet: true}
}

func (v NullableMetricsSearchResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsSearchResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


