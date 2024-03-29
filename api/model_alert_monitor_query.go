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

// checks if the AlertMonitorQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertMonitorQuery{}

// AlertMonitorQuery struct for AlertMonitorQuery
type AlertMonitorQuery struct {
	// The unique identifier of the row. Defaults to sequential capital letters, `A`, `B`, `C`, etc.
	RowId string `json:"rowId"`
	// The logs or metrics query that defines the stream of data the monitor runs on.
	Query string `json:"query"`
	// Indicates whether the current row is the trigger (final) row.
	IsTriggerRow bool `json:"isTriggerRow"`
}

// NewAlertMonitorQuery instantiates a new AlertMonitorQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertMonitorQuery(rowId string, query string, isTriggerRow bool) *AlertMonitorQuery {
	this := AlertMonitorQuery{}
	this.RowId = rowId
	this.Query = query
	this.IsTriggerRow = isTriggerRow
	return &this
}

// NewAlertMonitorQueryWithDefaults instantiates a new AlertMonitorQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertMonitorQueryWithDefaults() *AlertMonitorQuery {
	this := AlertMonitorQuery{}
	return &this
}

// GetRowId returns the RowId field value
func (o *AlertMonitorQuery) GetRowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value
// and a boolean to check if the value has been set.
func (o *AlertMonitorQuery) GetRowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowId, true
}

// SetRowId sets field value
func (o *AlertMonitorQuery) SetRowId(v string) {
	o.RowId = v
}

// GetQuery returns the Query field value
func (o *AlertMonitorQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *AlertMonitorQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *AlertMonitorQuery) SetQuery(v string) {
	o.Query = v
}

// GetIsTriggerRow returns the IsTriggerRow field value
func (o *AlertMonitorQuery) GetIsTriggerRow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTriggerRow
}

// GetIsTriggerRowOk returns a tuple with the IsTriggerRow field value
// and a boolean to check if the value has been set.
func (o *AlertMonitorQuery) GetIsTriggerRowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTriggerRow, true
}

// SetIsTriggerRow sets field value
func (o *AlertMonitorQuery) SetIsTriggerRow(v bool) {
	o.IsTriggerRow = v
}

func (o AlertMonitorQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertMonitorQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rowId"] = o.RowId
	toSerialize["query"] = o.Query
	toSerialize["isTriggerRow"] = o.IsTriggerRow
	return toSerialize, nil
}

type NullableAlertMonitorQuery struct {
	value *AlertMonitorQuery
	isSet bool
}

func (v NullableAlertMonitorQuery) Get() *AlertMonitorQuery {
	return v.value
}

func (v *NullableAlertMonitorQuery) Set(val *AlertMonitorQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertMonitorQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertMonitorQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertMonitorQuery(val *AlertMonitorQuery) *NullableAlertMonitorQuery {
	return &NullableAlertMonitorQuery{value: val, isSet: true}
}

func (v NullableAlertMonitorQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertMonitorQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


