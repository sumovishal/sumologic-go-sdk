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

// checks if the MonitorQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorQuery{}

// MonitorQuery A search query.
type MonitorQuery struct {
	// The unique identifier of the row. Defaults to sequential capital letters, `A`, `B`, `C`, etc.
	RowId string `json:"rowId"`
	// The logs or metrics query that defines the stream of data the monitor runs on.
	Query string `json:"query"`
}

type _MonitorQuery MonitorQuery

// NewMonitorQuery instantiates a new MonitorQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorQuery(rowId string, query string) *MonitorQuery {
	this := MonitorQuery{}
	this.RowId = rowId
	this.Query = query
	return &this
}

// NewMonitorQueryWithDefaults instantiates a new MonitorQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorQueryWithDefaults() *MonitorQuery {
	this := MonitorQuery{}
	return &this
}

// GetRowId returns the RowId field value
func (o *MonitorQuery) GetRowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value
// and a boolean to check if the value has been set.
func (o *MonitorQuery) GetRowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowId, true
}

// SetRowId sets field value
func (o *MonitorQuery) SetRowId(v string) {
	o.RowId = v
}

// GetQuery returns the Query field value
func (o *MonitorQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *MonitorQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *MonitorQuery) SetQuery(v string) {
	o.Query = v
}

func (o MonitorQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rowId"] = o.RowId
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

func (o *MonitorQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rowId",
		"query",
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

	varMonitorQuery := _MonitorQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMonitorQuery)

	if err != nil {
		return err
	}

	*o = MonitorQuery(varMonitorQuery)

	return err
}

type NullableMonitorQuery struct {
	value *MonitorQuery
	isSet bool
}

func (v NullableMonitorQuery) Get() *MonitorQuery {
	return v.value
}

func (v *NullableMonitorQuery) Set(val *MonitorQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorQuery(val *MonitorQuery) *NullableMonitorQuery {
	return &NullableMonitorQuery{value: val, isSet: true}
}

func (v NullableMonitorQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


