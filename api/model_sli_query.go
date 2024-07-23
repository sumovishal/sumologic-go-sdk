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

// checks if the SliQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SliQuery{}

// SliQuery Group of queries to allow for query arithmetic.
type SliQuery struct {
	// Unique id of the row. Used for query arithmetic, only for metric queries.
	RowId string `json:"rowId"`
	// Query String.
	Query string `json:"query"`
	// Determines whether to use count of rows (for logs) or data points (for metrics) in query result or specific field.
	UseRowCount bool `json:"useRowCount"`
	// Field of log query output to compare against. To be used only for logs based data type when `useRowCount` is false.
	Field *string `json:"field,omitempty"`
}

type _SliQuery SliQuery

// NewSliQuery instantiates a new SliQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliQuery(rowId string, query string, useRowCount bool) *SliQuery {
	this := SliQuery{}
	this.RowId = rowId
	this.Query = query
	this.UseRowCount = useRowCount
	return &this
}

// NewSliQueryWithDefaults instantiates a new SliQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliQueryWithDefaults() *SliQuery {
	this := SliQuery{}
	return &this
}

// GetRowId returns the RowId field value
func (o *SliQuery) GetRowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value
// and a boolean to check if the value has been set.
func (o *SliQuery) GetRowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowId, true
}

// SetRowId sets field value
func (o *SliQuery) SetRowId(v string) {
	o.RowId = v
}

// GetQuery returns the Query field value
func (o *SliQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SliQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SliQuery) SetQuery(v string) {
	o.Query = v
}

// GetUseRowCount returns the UseRowCount field value
func (o *SliQuery) GetUseRowCount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseRowCount
}

// GetUseRowCountOk returns a tuple with the UseRowCount field value
// and a boolean to check if the value has been set.
func (o *SliQuery) GetUseRowCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseRowCount, true
}

// SetUseRowCount sets field value
func (o *SliQuery) SetUseRowCount(v bool) {
	o.UseRowCount = v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *SliQuery) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliQuery) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *SliQuery) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *SliQuery) SetField(v string) {
	o.Field = &v
}

func (o SliQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SliQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rowId"] = o.RowId
	toSerialize["query"] = o.Query
	toSerialize["useRowCount"] = o.UseRowCount
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	return toSerialize, nil
}

func (o *SliQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rowId",
		"query",
		"useRowCount",
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

	varSliQuery := _SliQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSliQuery)

	if err != nil {
		return err
	}

	*o = SliQuery(varSliQuery)

	return err
}

type NullableSliQuery struct {
	value *SliQuery
	isSet bool
}

func (v NullableSliQuery) Get() *SliQuery {
	return v.value
}

func (v *NullableSliQuery) Set(val *SliQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSliQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSliQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliQuery(val *SliQuery) *NullableSliQuery {
	return &NullableSliQuery{value: val, isSet: true}
}

func (v NullableSliQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


