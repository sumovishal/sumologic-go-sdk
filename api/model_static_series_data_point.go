/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the StaticSeriesDataPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticSeriesDataPoint{}

// StaticSeriesDataPoint struct for StaticSeriesDataPoint
type StaticSeriesDataPoint struct {
	DataPoint
	// Epoch unix time stamp.
	X int64 `json:"x"`
	// The value of the data point.
	Y float64 `json:"y"`
}

type _StaticSeriesDataPoint StaticSeriesDataPoint

// NewStaticSeriesDataPoint instantiates a new StaticSeriesDataPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticSeriesDataPoint(x int64, y float64) *StaticSeriesDataPoint {
	this := StaticSeriesDataPoint{}
	this.X = x
	this.Y = y
	return &this
}

// NewStaticSeriesDataPointWithDefaults instantiates a new StaticSeriesDataPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticSeriesDataPointWithDefaults() *StaticSeriesDataPoint {
	this := StaticSeriesDataPoint{}
	return &this
}

// GetX returns the X field value
func (o *StaticSeriesDataPoint) GetX() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *StaticSeriesDataPoint) GetXOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *StaticSeriesDataPoint) SetX(v int64) {
	o.X = v
}

// GetY returns the Y field value
func (o *StaticSeriesDataPoint) GetY() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *StaticSeriesDataPoint) GetYOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *StaticSeriesDataPoint) SetY(v float64) {
	o.Y = v
}

func (o StaticSeriesDataPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticSeriesDataPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDataPoint, errDataPoint := json.Marshal(o.DataPoint)
	if errDataPoint != nil {
		return map[string]interface{}{}, errDataPoint
	}
	errDataPoint = json.Unmarshal([]byte(serializedDataPoint), &toSerialize)
	if errDataPoint != nil {
		return map[string]interface{}{}, errDataPoint
	}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return toSerialize, nil
}

func (o *StaticSeriesDataPoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"x",
		"y",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStaticSeriesDataPoint := _StaticSeriesDataPoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticSeriesDataPoint)

	if err != nil {
		return err
	}

	*o = StaticSeriesDataPoint(varStaticSeriesDataPoint)

	return err
}

type NullableStaticSeriesDataPoint struct {
	value *StaticSeriesDataPoint
	isSet bool
}

func (v NullableStaticSeriesDataPoint) Get() *StaticSeriesDataPoint {
	return v.value
}

func (v *NullableStaticSeriesDataPoint) Set(val *StaticSeriesDataPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticSeriesDataPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticSeriesDataPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticSeriesDataPoint(val *StaticSeriesDataPoint) *NullableStaticSeriesDataPoint {
	return &NullableStaticSeriesDataPoint{value: val, isSet: true}
}

func (v NullableStaticSeriesDataPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticSeriesDataPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


