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

// checks if the OutlierSeriesDataPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutlierSeriesDataPoint{}

// OutlierSeriesDataPoint struct for OutlierSeriesDataPoint
type OutlierSeriesDataPoint struct {
	DataPoint
	// Epoch unix time stamp.
	X int64 `json:"x"`
	Y OutlierDataValue `json:"y"`
}

type _OutlierSeriesDataPoint OutlierSeriesDataPoint

// NewOutlierSeriesDataPoint instantiates a new OutlierSeriesDataPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutlierSeriesDataPoint(x int64, y OutlierDataValue) *OutlierSeriesDataPoint {
	this := OutlierSeriesDataPoint{}
	this.X = x
	this.Y = y
	return &this
}

// NewOutlierSeriesDataPointWithDefaults instantiates a new OutlierSeriesDataPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutlierSeriesDataPointWithDefaults() *OutlierSeriesDataPoint {
	this := OutlierSeriesDataPoint{}
	return &this
}

// GetX returns the X field value
func (o *OutlierSeriesDataPoint) GetX() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *OutlierSeriesDataPoint) GetXOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *OutlierSeriesDataPoint) SetX(v int64) {
	o.X = v
}

// GetY returns the Y field value
func (o *OutlierSeriesDataPoint) GetY() OutlierDataValue {
	if o == nil {
		var ret OutlierDataValue
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *OutlierSeriesDataPoint) GetYOk() (*OutlierDataValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *OutlierSeriesDataPoint) SetY(v OutlierDataValue) {
	o.Y = v
}

func (o OutlierSeriesDataPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutlierSeriesDataPoint) ToMap() (map[string]interface{}, error) {
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

func (o *OutlierSeriesDataPoint) UnmarshalJSON(data []byte) (err error) {
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

	varOutlierSeriesDataPoint := _OutlierSeriesDataPoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutlierSeriesDataPoint)

	if err != nil {
		return err
	}

	*o = OutlierSeriesDataPoint(varOutlierSeriesDataPoint)

	return err
}

type NullableOutlierSeriesDataPoint struct {
	value *OutlierSeriesDataPoint
	isSet bool
}

func (v NullableOutlierSeriesDataPoint) Get() *OutlierSeriesDataPoint {
	return v.value
}

func (v *NullableOutlierSeriesDataPoint) Set(val *OutlierSeriesDataPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableOutlierSeriesDataPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableOutlierSeriesDataPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutlierSeriesDataPoint(val *OutlierSeriesDataPoint) *NullableOutlierSeriesDataPoint {
	return &NullableOutlierSeriesDataPoint{value: val, isSet: true}
}

func (v NullableOutlierSeriesDataPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutlierSeriesDataPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


