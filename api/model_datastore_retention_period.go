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

// checks if the DatastoreRetentionPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatastoreRetentionPeriod{}

// DatastoreRetentionPeriod struct for DatastoreRetentionPeriod
type DatastoreRetentionPeriod struct {
	// Retention period in days.
	RetentionPeriod int64 `json:"retentionPeriod"`
}

// NewDatastoreRetentionPeriod instantiates a new DatastoreRetentionPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatastoreRetentionPeriod(retentionPeriod int64) *DatastoreRetentionPeriod {
	this := DatastoreRetentionPeriod{}
	this.RetentionPeriod = retentionPeriod
	return &this
}

// NewDatastoreRetentionPeriodWithDefaults instantiates a new DatastoreRetentionPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatastoreRetentionPeriodWithDefaults() *DatastoreRetentionPeriod {
	this := DatastoreRetentionPeriod{}
	return &this
}

// GetRetentionPeriod returns the RetentionPeriod field value
func (o *DatastoreRetentionPeriod) GetRetentionPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value
// and a boolean to check if the value has been set.
func (o *DatastoreRetentionPeriod) GetRetentionPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionPeriod, true
}

// SetRetentionPeriod sets field value
func (o *DatastoreRetentionPeriod) SetRetentionPeriod(v int64) {
	o.RetentionPeriod = v
}

func (o DatastoreRetentionPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatastoreRetentionPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["retentionPeriod"] = o.RetentionPeriod
	return toSerialize, nil
}

type NullableDatastoreRetentionPeriod struct {
	value *DatastoreRetentionPeriod
	isSet bool
}

func (v NullableDatastoreRetentionPeriod) Get() *DatastoreRetentionPeriod {
	return v.value
}

func (v *NullableDatastoreRetentionPeriod) Set(val *DatastoreRetentionPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreRetentionPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreRetentionPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreRetentionPeriod(val *DatastoreRetentionPeriod) *NullableDatastoreRetentionPeriod {
	return &NullableDatastoreRetentionPeriod{value: val, isSet: true}
}

func (v NullableDatastoreRetentionPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreRetentionPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


