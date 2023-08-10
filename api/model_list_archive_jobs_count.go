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

// checks if the ListArchiveJobsCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListArchiveJobsCount{}

// ListArchiveJobsCount struct for ListArchiveJobsCount
type ListArchiveJobsCount struct {
	// List of archive sources with count of jobs having various statuses.
	Data []ArchiveJobsCount `json:"data"`
}

// NewListArchiveJobsCount instantiates a new ListArchiveJobsCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListArchiveJobsCount(data []ArchiveJobsCount) *ListArchiveJobsCount {
	this := ListArchiveJobsCount{}
	this.Data = data
	return &this
}

// NewListArchiveJobsCountWithDefaults instantiates a new ListArchiveJobsCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListArchiveJobsCountWithDefaults() *ListArchiveJobsCount {
	this := ListArchiveJobsCount{}
	return &this
}

// GetData returns the Data field value
func (o *ListArchiveJobsCount) GetData() []ArchiveJobsCount {
	if o == nil {
		var ret []ArchiveJobsCount
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListArchiveJobsCount) GetDataOk() ([]ArchiveJobsCount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListArchiveJobsCount) SetData(v []ArchiveJobsCount) {
	o.Data = v
}

func (o ListArchiveJobsCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListArchiveJobsCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListArchiveJobsCount struct {
	value *ListArchiveJobsCount
	isSet bool
}

func (v NullableListArchiveJobsCount) Get() *ListArchiveJobsCount {
	return v.value
}

func (v *NullableListArchiveJobsCount) Set(val *ListArchiveJobsCount) {
	v.value = val
	v.isSet = true
}

func (v NullableListArchiveJobsCount) IsSet() bool {
	return v.isSet
}

func (v *NullableListArchiveJobsCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListArchiveJobsCount(val *ListArchiveJobsCount) *NullableListArchiveJobsCount {
	return &NullableListArchiveJobsCount{value: val, isSet: true}
}

func (v NullableListArchiveJobsCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListArchiveJobsCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


