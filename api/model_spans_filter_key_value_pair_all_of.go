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

// checks if the SpansFilterKeyValuePairAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpansFilterKeyValuePairAllOf{}

// SpansFilterKeyValuePairAllOf A representation of a span filter where both the field name and field value are provided, e.g. http.status_code > 500. 
type SpansFilterKeyValuePairAllOf struct {
	// A symbol that indicates an operation to be performed between a `fieldName` and `fieldValue`.
	Operator string `json:"operator"`
	// The second argument of the operation applied to a `fieldName`.
	FieldValue string `json:"fieldValue"`
}

// NewSpansFilterKeyValuePairAllOf instantiates a new SpansFilterKeyValuePairAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpansFilterKeyValuePairAllOf(operator string, fieldValue string) *SpansFilterKeyValuePairAllOf {
	this := SpansFilterKeyValuePairAllOf{}
	this.Operator = operator
	this.FieldValue = fieldValue
	return &this
}

// NewSpansFilterKeyValuePairAllOfWithDefaults instantiates a new SpansFilterKeyValuePairAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpansFilterKeyValuePairAllOfWithDefaults() *SpansFilterKeyValuePairAllOf {
	this := SpansFilterKeyValuePairAllOf{}
	return &this
}

// GetOperator returns the Operator field value
func (o *SpansFilterKeyValuePairAllOf) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SpansFilterKeyValuePairAllOf) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *SpansFilterKeyValuePairAllOf) SetOperator(v string) {
	o.Operator = v
}

// GetFieldValue returns the FieldValue field value
func (o *SpansFilterKeyValuePairAllOf) GetFieldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value
// and a boolean to check if the value has been set.
func (o *SpansFilterKeyValuePairAllOf) GetFieldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldValue, true
}

// SetFieldValue sets field value
func (o *SpansFilterKeyValuePairAllOf) SetFieldValue(v string) {
	o.FieldValue = v
}

func (o SpansFilterKeyValuePairAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpansFilterKeyValuePairAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operator"] = o.Operator
	toSerialize["fieldValue"] = o.FieldValue
	return toSerialize, nil
}

type NullableSpansFilterKeyValuePairAllOf struct {
	value *SpansFilterKeyValuePairAllOf
	isSet bool
}

func (v NullableSpansFilterKeyValuePairAllOf) Get() *SpansFilterKeyValuePairAllOf {
	return v.value
}

func (v *NullableSpansFilterKeyValuePairAllOf) Set(val *SpansFilterKeyValuePairAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpansFilterKeyValuePairAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpansFilterKeyValuePairAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpansFilterKeyValuePairAllOf(val *SpansFilterKeyValuePairAllOf) *NullableSpansFilterKeyValuePairAllOf {
	return &NullableSpansFilterKeyValuePairAllOf{value: val, isSet: true}
}

func (v NullableSpansFilterKeyValuePairAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpansFilterKeyValuePairAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


