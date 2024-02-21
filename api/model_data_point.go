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

// checks if the DataPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPoint{}

// DataPoint Data for visualizing monitor chart.
type DataPoint struct {
	// Type of the data point.
	DataPointType *string `json:"dataPointType,omitempty"`
}

// NewDataPoint instantiates a new DataPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPoint() *DataPoint {
	this := DataPoint{}
	return &this
}

// NewDataPointWithDefaults instantiates a new DataPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPointWithDefaults() *DataPoint {
	this := DataPoint{}
	return &this
}

// GetDataPointType returns the DataPointType field value if set, zero value otherwise.
func (o *DataPoint) GetDataPointType() string {
	if o == nil || IsNil(o.DataPointType) {
		var ret string
		return ret
	}
	return *o.DataPointType
}

// GetDataPointTypeOk returns a tuple with the DataPointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPoint) GetDataPointTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataPointType) {
		return nil, false
	}
	return o.DataPointType, true
}

// HasDataPointType returns a boolean if a field has been set.
func (o *DataPoint) HasDataPointType() bool {
	if o != nil && !IsNil(o.DataPointType) {
		return true
	}

	return false
}

// SetDataPointType gets a reference to the given string and assigns it to the DataPointType field.
func (o *DataPoint) SetDataPointType(v string) {
	o.DataPointType = &v
}

func (o DataPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataPointType) {
		toSerialize["dataPointType"] = o.DataPointType
	}
	return toSerialize, nil
}

type NullableDataPoint struct {
	value *DataPoint
	isSet bool
}

func (v NullableDataPoint) Get() *DataPoint {
	return v.value
}

func (v *NullableDataPoint) Set(val *DataPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPoint(val *DataPoint) *NullableDataPoint {
	return &NullableDataPoint{value: val, isSet: true}
}

func (v NullableDataPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


