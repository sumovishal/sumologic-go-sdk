/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// LogQueryVariableSourceDefinitionAllOf Variable with values that are powered by a log query.
type LogQueryVariableSourceDefinitionAllOf struct {
	// A log query.
	Query string `json:"query"`
	// A field in log query to populate the variable values.
	Field string `json:"field"`
}

// NewLogQueryVariableSourceDefinitionAllOf instantiates a new LogQueryVariableSourceDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogQueryVariableSourceDefinitionAllOf(query string, field string) *LogQueryVariableSourceDefinitionAllOf {
	this := LogQueryVariableSourceDefinitionAllOf{}
	this.Query = query
	this.Field = field
	return &this
}

// NewLogQueryVariableSourceDefinitionAllOfWithDefaults instantiates a new LogQueryVariableSourceDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogQueryVariableSourceDefinitionAllOfWithDefaults() *LogQueryVariableSourceDefinitionAllOf {
	this := LogQueryVariableSourceDefinitionAllOf{}
	return &this
}

// GetQuery returns the Query field value
func (o *LogQueryVariableSourceDefinitionAllOf) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LogQueryVariableSourceDefinitionAllOf) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *LogQueryVariableSourceDefinitionAllOf) SetQuery(v string) {
	o.Query = v
}

// GetField returns the Field field value
func (o *LogQueryVariableSourceDefinitionAllOf) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *LogQueryVariableSourceDefinitionAllOf) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *LogQueryVariableSourceDefinitionAllOf) SetField(v string) {
	o.Field = v
}

func (o LogQueryVariableSourceDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["field"] = o.Field
	}
	return json.Marshal(toSerialize)
}

type NullableLogQueryVariableSourceDefinitionAllOf struct {
	value *LogQueryVariableSourceDefinitionAllOf
	isSet bool
}

func (v NullableLogQueryVariableSourceDefinitionAllOf) Get() *LogQueryVariableSourceDefinitionAllOf {
	return v.value
}

func (v *NullableLogQueryVariableSourceDefinitionAllOf) Set(val *LogQueryVariableSourceDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLogQueryVariableSourceDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLogQueryVariableSourceDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogQueryVariableSourceDefinitionAllOf(val *LogQueryVariableSourceDefinitionAllOf) *NullableLogQueryVariableSourceDefinitionAllOf {
	return &NullableLogQueryVariableSourceDefinitionAllOf{value: val, isSet: true}
}

func (v NullableLogQueryVariableSourceDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogQueryVariableSourceDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


