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

// checks if the VisualDataAxes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualDataAxes{}

// VisualDataAxes struct for VisualDataAxes
type VisualDataAxes struct {
	// The data of the primary x axis.
	X []VisualAxisData `json:"x"`
	// The data of the primary y axis.
	Y []VisualAxisData `json:"y"`
	// The data of the secondary x axis.
	X2 []VisualAxisData `json:"x2,omitempty"`
	// The data of the secondary y axis.
	Y2 []VisualAxisData `json:"y2,omitempty"`
}

type _VisualDataAxes VisualDataAxes

// NewVisualDataAxes instantiates a new VisualDataAxes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualDataAxes(x []VisualAxisData, y []VisualAxisData) *VisualDataAxes {
	this := VisualDataAxes{}
	this.X = x
	this.Y = y
	return &this
}

// NewVisualDataAxesWithDefaults instantiates a new VisualDataAxes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualDataAxesWithDefaults() *VisualDataAxes {
	this := VisualDataAxes{}
	return &this
}

// GetX returns the X field value
func (o *VisualDataAxes) GetX() []VisualAxisData {
	if o == nil {
		var ret []VisualAxisData
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *VisualDataAxes) GetXOk() ([]VisualAxisData, bool) {
	if o == nil {
		return nil, false
	}
	return o.X, true
}

// SetX sets field value
func (o *VisualDataAxes) SetX(v []VisualAxisData) {
	o.X = v
}

// GetY returns the Y field value
func (o *VisualDataAxes) GetY() []VisualAxisData {
	if o == nil {
		var ret []VisualAxisData
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *VisualDataAxes) GetYOk() ([]VisualAxisData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Y, true
}

// SetY sets field value
func (o *VisualDataAxes) SetY(v []VisualAxisData) {
	o.Y = v
}

// GetX2 returns the X2 field value if set, zero value otherwise.
func (o *VisualDataAxes) GetX2() []VisualAxisData {
	if o == nil || IsNil(o.X2) {
		var ret []VisualAxisData
		return ret
	}
	return o.X2
}

// GetX2Ok returns a tuple with the X2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataAxes) GetX2Ok() ([]VisualAxisData, bool) {
	if o == nil || IsNil(o.X2) {
		return nil, false
	}
	return o.X2, true
}

// HasX2 returns a boolean if a field has been set.
func (o *VisualDataAxes) HasX2() bool {
	if o != nil && !IsNil(o.X2) {
		return true
	}

	return false
}

// SetX2 gets a reference to the given []VisualAxisData and assigns it to the X2 field.
func (o *VisualDataAxes) SetX2(v []VisualAxisData) {
	o.X2 = v
}

// GetY2 returns the Y2 field value if set, zero value otherwise.
func (o *VisualDataAxes) GetY2() []VisualAxisData {
	if o == nil || IsNil(o.Y2) {
		var ret []VisualAxisData
		return ret
	}
	return o.Y2
}

// GetY2Ok returns a tuple with the Y2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataAxes) GetY2Ok() ([]VisualAxisData, bool) {
	if o == nil || IsNil(o.Y2) {
		return nil, false
	}
	return o.Y2, true
}

// HasY2 returns a boolean if a field has been set.
func (o *VisualDataAxes) HasY2() bool {
	if o != nil && !IsNil(o.Y2) {
		return true
	}

	return false
}

// SetY2 gets a reference to the given []VisualAxisData and assigns it to the Y2 field.
func (o *VisualDataAxes) SetY2(v []VisualAxisData) {
	o.Y2 = v
}

func (o VisualDataAxes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualDataAxes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	if !IsNil(o.X2) {
		toSerialize["x2"] = o.X2
	}
	if !IsNil(o.Y2) {
		toSerialize["y2"] = o.Y2
	}
	return toSerialize, nil
}

func (o *VisualDataAxes) UnmarshalJSON(data []byte) (err error) {
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

	varVisualDataAxes := _VisualDataAxes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVisualDataAxes)

	if err != nil {
		return err
	}

	*o = VisualDataAxes(varVisualDataAxes)

	return err
}

type NullableVisualDataAxes struct {
	value *VisualDataAxes
	isSet bool
}

func (v NullableVisualDataAxes) Get() *VisualDataAxes {
	return v.value
}

func (v *NullableVisualDataAxes) Set(val *VisualDataAxes) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualDataAxes) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualDataAxes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualDataAxes(val *VisualDataAxes) *NullableVisualDataAxes {
	return &NullableVisualDataAxes{value: val, isSet: true}
}

func (v NullableVisualDataAxes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualDataAxes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


