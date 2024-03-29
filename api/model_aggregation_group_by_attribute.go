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

// checks if the AggregationGroupByAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregationGroupByAttribute{}

// AggregationGroupByAttribute Base group by attribute object.
type AggregationGroupByAttribute struct {
	// Attribute type of the object model.
	AttributeType string `json:"attributeType"`
}

// NewAggregationGroupByAttribute instantiates a new AggregationGroupByAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationGroupByAttribute(attributeType string) *AggregationGroupByAttribute {
	this := AggregationGroupByAttribute{}
	this.AttributeType = attributeType
	return &this
}

// NewAggregationGroupByAttributeWithDefaults instantiates a new AggregationGroupByAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationGroupByAttributeWithDefaults() *AggregationGroupByAttribute {
	this := AggregationGroupByAttribute{}
	return &this
}

// GetAttributeType returns the AttributeType field value
func (o *AggregationGroupByAttribute) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *AggregationGroupByAttribute) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *AggregationGroupByAttribute) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o AggregationGroupByAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregationGroupByAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeType"] = o.AttributeType
	return toSerialize, nil
}

type NullableAggregationGroupByAttribute struct {
	value *AggregationGroupByAttribute
	isSet bool
}

func (v NullableAggregationGroupByAttribute) Get() *AggregationGroupByAttribute {
	return v.value
}

func (v *NullableAggregationGroupByAttribute) Set(val *AggregationGroupByAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationGroupByAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationGroupByAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationGroupByAttribute(val *AggregationGroupByAttribute) *NullableAggregationGroupByAttribute {
	return &NullableAggregationGroupByAttribute{value: val, isSet: true}
}

func (v NullableAggregationGroupByAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationGroupByAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


