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

// checks if the MetricsQueryRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsQueryRow{}

// MetricsQueryRow struct for MetricsQueryRow
type MetricsQueryRow struct {
	// Row id for the query row, A to Z letter.
	RowId string `json:"rowId" validate:"regexp=[A-Z]"`
	// A metric query consists of a metric, one or more filters and optionally, one or more [Metrics Operators](https://help.sumologic.com/?cid=10144). Strictly speaking, both filters and operators are optional.  Most of the [Metrics Operators](https://help.sumologic.com/?cid=10144) are allowed in the query string except `fillmissing`, `outlier`, `quantize` and `timeshift`.    * `fillmissing`: Not supported in API.   * `outlier`: Not supported in API.   * `quantize`: Only supported through `quantization` param.   * `timeshift`: Only supported through `timeshift` param.   In practice, your metric queries will almost always contain filters that narrow the scope of your query. For more information about the query language see [Metrics Queries](https://help.sumologic.com/?cid=1079).
	Query string `json:"query"`
	// Segregates time series data by time period. This allows you to create aggregated results in buckets of fixed intervals (for example, 5-minute intervals). The value is in milliseconds.
	Quantization *int64 `json:"quantization,omitempty"`
	// We use the term rollup to refer to the aggregation function Sumo Logic uses when quantizing metrics. Can be `Avg`, `Sum`, `Min`, `Max`, `Count` or `None`.
	Rollup *string `json:"rollup,omitempty" validate:"regexp=^(Count|Min|Max|Sum|Avg|None)$|^$"`
	// Shifts the time series from your metrics query by the specified amount of time. This can help when comparing a time series across multiple time periods. Specified as a signed duration in milliseconds.
	Timeshift *int64 `json:"timeshift,omitempty"`
}

type _MetricsQueryRow MetricsQueryRow

// NewMetricsQueryRow instantiates a new MetricsQueryRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsQueryRow(rowId string, query string) *MetricsQueryRow {
	this := MetricsQueryRow{}
	this.RowId = rowId
	this.Query = query
	return &this
}

// NewMetricsQueryRowWithDefaults instantiates a new MetricsQueryRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsQueryRowWithDefaults() *MetricsQueryRow {
	this := MetricsQueryRow{}
	return &this
}

// GetRowId returns the RowId field value
func (o *MetricsQueryRow) GetRowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value
// and a boolean to check if the value has been set.
func (o *MetricsQueryRow) GetRowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowId, true
}

// SetRowId sets field value
func (o *MetricsQueryRow) SetRowId(v string) {
	o.RowId = v
}

// GetQuery returns the Query field value
func (o *MetricsQueryRow) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *MetricsQueryRow) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *MetricsQueryRow) SetQuery(v string) {
	o.Query = v
}

// GetQuantization returns the Quantization field value if set, zero value otherwise.
func (o *MetricsQueryRow) GetQuantization() int64 {
	if o == nil || IsNil(o.Quantization) {
		var ret int64
		return ret
	}
	return *o.Quantization
}

// GetQuantizationOk returns a tuple with the Quantization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryRow) GetQuantizationOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantization) {
		return nil, false
	}
	return o.Quantization, true
}

// HasQuantization returns a boolean if a field has been set.
func (o *MetricsQueryRow) HasQuantization() bool {
	if o != nil && !IsNil(o.Quantization) {
		return true
	}

	return false
}

// SetQuantization gets a reference to the given int64 and assigns it to the Quantization field.
func (o *MetricsQueryRow) SetQuantization(v int64) {
	o.Quantization = &v
}

// GetRollup returns the Rollup field value if set, zero value otherwise.
func (o *MetricsQueryRow) GetRollup() string {
	if o == nil || IsNil(o.Rollup) {
		var ret string
		return ret
	}
	return *o.Rollup
}

// GetRollupOk returns a tuple with the Rollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryRow) GetRollupOk() (*string, bool) {
	if o == nil || IsNil(o.Rollup) {
		return nil, false
	}
	return o.Rollup, true
}

// HasRollup returns a boolean if a field has been set.
func (o *MetricsQueryRow) HasRollup() bool {
	if o != nil && !IsNil(o.Rollup) {
		return true
	}

	return false
}

// SetRollup gets a reference to the given string and assigns it to the Rollup field.
func (o *MetricsQueryRow) SetRollup(v string) {
	o.Rollup = &v
}

// GetTimeshift returns the Timeshift field value if set, zero value otherwise.
func (o *MetricsQueryRow) GetTimeshift() int64 {
	if o == nil || IsNil(o.Timeshift) {
		var ret int64
		return ret
	}
	return *o.Timeshift
}

// GetTimeshiftOk returns a tuple with the Timeshift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryRow) GetTimeshiftOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeshift) {
		return nil, false
	}
	return o.Timeshift, true
}

// HasTimeshift returns a boolean if a field has been set.
func (o *MetricsQueryRow) HasTimeshift() bool {
	if o != nil && !IsNil(o.Timeshift) {
		return true
	}

	return false
}

// SetTimeshift gets a reference to the given int64 and assigns it to the Timeshift field.
func (o *MetricsQueryRow) SetTimeshift(v int64) {
	o.Timeshift = &v
}

func (o MetricsQueryRow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsQueryRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rowId"] = o.RowId
	toSerialize["query"] = o.Query
	if !IsNil(o.Quantization) {
		toSerialize["quantization"] = o.Quantization
	}
	if !IsNil(o.Rollup) {
		toSerialize["rollup"] = o.Rollup
	}
	if !IsNil(o.Timeshift) {
		toSerialize["timeshift"] = o.Timeshift
	}
	return toSerialize, nil
}

func (o *MetricsQueryRow) UnmarshalJSON(data []byte) (err error) {
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

	varMetricsQueryRow := _MetricsQueryRow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetricsQueryRow)

	if err != nil {
		return err
	}

	*o = MetricsQueryRow(varMetricsQueryRow)

	return err
}

type NullableMetricsQueryRow struct {
	value *MetricsQueryRow
	isSet bool
}

func (v NullableMetricsQueryRow) Get() *MetricsQueryRow {
	return v.value
}

func (v *NullableMetricsQueryRow) Set(val *MetricsQueryRow) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsQueryRow) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsQueryRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsQueryRow(val *MetricsQueryRow) *NullableMetricsQueryRow {
	return &NullableMetricsQueryRow{value: val, isSet: true}
}

func (v NullableMetricsQueryRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsQueryRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


