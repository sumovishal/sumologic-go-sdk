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

// checks if the LogQueryVariableSourceDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogQueryVariableSourceDefinition{}

// LogQueryVariableSourceDefinition struct for LogQueryVariableSourceDefinition
type LogQueryVariableSourceDefinition struct {
	VariableSourceDefinition
	// A log query.
	Query string `json:"query"`
	// A field in log query to populate the variable values.
	Field string `json:"field"`
}

// NewLogQueryVariableSourceDefinition instantiates a new LogQueryVariableSourceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogQueryVariableSourceDefinition(query string, field string, variableSourceType string) *LogQueryVariableSourceDefinition {
	this := LogQueryVariableSourceDefinition{}
	this.VariableSourceType = variableSourceType
	this.Query = query
	this.Field = field
	return &this
}

// NewLogQueryVariableSourceDefinitionWithDefaults instantiates a new LogQueryVariableSourceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogQueryVariableSourceDefinitionWithDefaults() *LogQueryVariableSourceDefinition {
	this := LogQueryVariableSourceDefinition{}
	return &this
}

// GetQuery returns the Query field value
func (o *LogQueryVariableSourceDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LogQueryVariableSourceDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *LogQueryVariableSourceDefinition) SetQuery(v string) {
	o.Query = v
}

// GetField returns the Field field value
func (o *LogQueryVariableSourceDefinition) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *LogQueryVariableSourceDefinition) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *LogQueryVariableSourceDefinition) SetField(v string) {
	o.Field = v
}

func (o LogQueryVariableSourceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogQueryVariableSourceDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVariableSourceDefinition, errVariableSourceDefinition := json.Marshal(o.VariableSourceDefinition)
	if errVariableSourceDefinition != nil {
		return map[string]interface{}{}, errVariableSourceDefinition
	}
	errVariableSourceDefinition = json.Unmarshal([]byte(serializedVariableSourceDefinition), &toSerialize)
	if errVariableSourceDefinition != nil {
		return map[string]interface{}{}, errVariableSourceDefinition
	}
	toSerialize["query"] = o.Query
	toSerialize["field"] = o.Field
	return toSerialize, nil
}

type NullableLogQueryVariableSourceDefinition struct {
	value *LogQueryVariableSourceDefinition
	isSet bool
}

func (v NullableLogQueryVariableSourceDefinition) Get() *LogQueryVariableSourceDefinition {
	return v.value
}

func (v *NullableLogQueryVariableSourceDefinition) Set(val *LogQueryVariableSourceDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogQueryVariableSourceDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogQueryVariableSourceDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogQueryVariableSourceDefinition(val *LogQueryVariableSourceDefinition) *NullableLogQueryVariableSourceDefinition {
	return &NullableLogQueryVariableSourceDefinition{value: val, isSet: true}
}

func (v NullableLogQueryVariableSourceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogQueryVariableSourceDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


