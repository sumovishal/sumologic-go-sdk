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

// checks if the SpansCalculationVisualization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpansCalculationVisualization{}

// SpansCalculationVisualization struct for SpansCalculationVisualization
type SpansCalculationVisualization struct {
	SpansVisualization
	// A field by which the spans are aggregated.
	Field string `json:"field"`
	Aggregator SpanCalculationAggregator `json:"aggregator"`
}

type _SpansCalculationVisualization SpansCalculationVisualization

// NewSpansCalculationVisualization instantiates a new SpansCalculationVisualization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpansCalculationVisualization(field string, aggregator SpanCalculationAggregator, type_ string, name string) *SpansCalculationVisualization {
	this := SpansCalculationVisualization{}
	this.Type = type_
	this.Name = name
	this.Field = field
	this.Aggregator = aggregator
	return &this
}

// NewSpansCalculationVisualizationWithDefaults instantiates a new SpansCalculationVisualization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpansCalculationVisualizationWithDefaults() *SpansCalculationVisualization {
	this := SpansCalculationVisualization{}
	return &this
}

// GetField returns the Field field value
func (o *SpansCalculationVisualization) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *SpansCalculationVisualization) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *SpansCalculationVisualization) SetField(v string) {
	o.Field = v
}

// GetAggregator returns the Aggregator field value
func (o *SpansCalculationVisualization) GetAggregator() SpanCalculationAggregator {
	if o == nil {
		var ret SpanCalculationAggregator
		return ret
	}

	return o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value
// and a boolean to check if the value has been set.
func (o *SpansCalculationVisualization) GetAggregatorOk() (*SpanCalculationAggregator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregator, true
}

// SetAggregator sets field value
func (o *SpansCalculationVisualization) SetAggregator(v SpanCalculationAggregator) {
	o.Aggregator = v
}

func (o SpansCalculationVisualization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpansCalculationVisualization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSpansVisualization, errSpansVisualization := json.Marshal(o.SpansVisualization)
	if errSpansVisualization != nil {
		return map[string]interface{}{}, errSpansVisualization
	}
	errSpansVisualization = json.Unmarshal([]byte(serializedSpansVisualization), &toSerialize)
	if errSpansVisualization != nil {
		return map[string]interface{}{}, errSpansVisualization
	}
	toSerialize["field"] = o.Field
	toSerialize["aggregator"] = o.Aggregator
	return toSerialize, nil
}

func (o *SpansCalculationVisualization) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
		"aggregator",
		"type",
		"name",
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

	varSpansCalculationVisualization := _SpansCalculationVisualization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpansCalculationVisualization)

	if err != nil {
		return err
	}

	*o = SpansCalculationVisualization(varSpansCalculationVisualization)

	return err
}

type NullableSpansCalculationVisualization struct {
	value *SpansCalculationVisualization
	isSet bool
}

func (v NullableSpansCalculationVisualization) Get() *SpansCalculationVisualization {
	return v.value
}

func (v *NullableSpansCalculationVisualization) Set(val *SpansCalculationVisualization) {
	v.value = val
	v.isSet = true
}

func (v NullableSpansCalculationVisualization) IsSet() bool {
	return v.isSet
}

func (v *NullableSpansCalculationVisualization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpansCalculationVisualization(val *SpansCalculationVisualization) *NullableSpansCalculationVisualization {
	return &NullableSpansCalculationVisualization{value: val, isSet: true}
}

func (v NullableSpansCalculationVisualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpansCalculationVisualization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


