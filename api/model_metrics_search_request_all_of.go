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

// checks if the MetricsSearchRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsSearchRequestAllOf{}

// MetricsSearchRequestAllOf struct for MetricsSearchRequestAllOf
type MetricsSearchRequestAllOf struct {
	// The identifier of the folder to save the metrics search in. By default it is saved in your personal folder. 
	FolderId *string `json:"folderId,omitempty"`
}

// NewMetricsSearchRequestAllOf instantiates a new MetricsSearchRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsSearchRequestAllOf() *MetricsSearchRequestAllOf {
	this := MetricsSearchRequestAllOf{}
	return &this
}

// NewMetricsSearchRequestAllOfWithDefaults instantiates a new MetricsSearchRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsSearchRequestAllOfWithDefaults() *MetricsSearchRequestAllOf {
	this := MetricsSearchRequestAllOf{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *MetricsSearchRequestAllOf) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchRequestAllOf) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *MetricsSearchRequestAllOf) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *MetricsSearchRequestAllOf) SetFolderId(v string) {
	o.FolderId = &v
}

func (o MetricsSearchRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsSearchRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	return toSerialize, nil
}

type NullableMetricsSearchRequestAllOf struct {
	value *MetricsSearchRequestAllOf
	isSet bool
}

func (v NullableMetricsSearchRequestAllOf) Get() *MetricsSearchRequestAllOf {
	return v.value
}

func (v *NullableMetricsSearchRequestAllOf) Set(val *MetricsSearchRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsSearchRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsSearchRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsSearchRequestAllOf(val *MetricsSearchRequestAllOf) *NullableMetricsSearchRequestAllOf {
	return &NullableMetricsSearchRequestAllOf{value: val, isSet: true}
}

func (v NullableMetricsSearchRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsSearchRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


