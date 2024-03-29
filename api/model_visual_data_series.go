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

// checks if the VisualDataSeries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualDataSeries{}

// VisualDataSeries struct for VisualDataSeries
type VisualDataSeries struct {
	// The id of the query.
	QueryId string `json:"queryId"`
	// The meaning of 'name' depends on the series type.   - For results of type 'timeseries', it is the value of the 'metric' key.   - For results of type 'nontimeseries', it is the name of one of the fields that is not part of 'xAxisKeys'.   - For results of type 'table', it is the comma-separated string of names of all fields. 
	Name string `json:"name"`
	// A list of data points in the visual series.
	DataPoints []VisualPointData `json:"dataPoints"`
	AggregateInfo *VisualAggregateData `json:"aggregateInfo,omitempty"`
	MetaData *VisualMetaData `json:"metaData,omitempty"`
	// Type of the visual series.
	SeriesType *string `json:"seriesType,omitempty"`
	// Keys that will be plotted as a point on the x axis.
	XAxisKeys []string `json:"xAxisKeys,omitempty"`
	// Value that represents if the series values are String or Double
	ValueType *string `json:"valueType,omitempty"`
	// Source of the visual series.
	Source *string `json:"source,omitempty"`
	// Keys that will be plotted as a point on the x axis and their data type
	XAxisKeyTypes *map[string]string `json:"xAxisKeyTypes,omitempty"`
	QueryInfo *MetricsQueryResultInfo `json:"queryInfo,omitempty"`
}

// NewVisualDataSeries instantiates a new VisualDataSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualDataSeries(queryId string, name string, dataPoints []VisualPointData) *VisualDataSeries {
	this := VisualDataSeries{}
	this.QueryId = queryId
	this.Name = name
	this.DataPoints = dataPoints
	return &this
}

// NewVisualDataSeriesWithDefaults instantiates a new VisualDataSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualDataSeriesWithDefaults() *VisualDataSeries {
	this := VisualDataSeries{}
	return &this
}

// GetQueryId returns the QueryId field value
func (o *VisualDataSeries) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *VisualDataSeries) SetQueryId(v string) {
	o.QueryId = v
}

// GetName returns the Name field value
func (o *VisualDataSeries) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VisualDataSeries) SetName(v string) {
	o.Name = v
}

// GetDataPoints returns the DataPoints field value
func (o *VisualDataSeries) GetDataPoints() []VisualPointData {
	if o == nil {
		var ret []VisualPointData
		return ret
	}

	return o.DataPoints
}

// GetDataPointsOk returns a tuple with the DataPoints field value
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetDataPointsOk() ([]VisualPointData, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataPoints, true
}

// SetDataPoints sets field value
func (o *VisualDataSeries) SetDataPoints(v []VisualPointData) {
	o.DataPoints = v
}

// GetAggregateInfo returns the AggregateInfo field value if set, zero value otherwise.
func (o *VisualDataSeries) GetAggregateInfo() VisualAggregateData {
	if o == nil || IsNil(o.AggregateInfo) {
		var ret VisualAggregateData
		return ret
	}
	return *o.AggregateInfo
}

// GetAggregateInfoOk returns a tuple with the AggregateInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetAggregateInfoOk() (*VisualAggregateData, bool) {
	if o == nil || IsNil(o.AggregateInfo) {
		return nil, false
	}
	return o.AggregateInfo, true
}

// HasAggregateInfo returns a boolean if a field has been set.
func (o *VisualDataSeries) HasAggregateInfo() bool {
	if o != nil && !IsNil(o.AggregateInfo) {
		return true
	}

	return false
}

// SetAggregateInfo gets a reference to the given VisualAggregateData and assigns it to the AggregateInfo field.
func (o *VisualDataSeries) SetAggregateInfo(v VisualAggregateData) {
	o.AggregateInfo = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *VisualDataSeries) GetMetaData() VisualMetaData {
	if o == nil || IsNil(o.MetaData) {
		var ret VisualMetaData
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetMetaDataOk() (*VisualMetaData, bool) {
	if o == nil || IsNil(o.MetaData) {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *VisualDataSeries) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given VisualMetaData and assigns it to the MetaData field.
func (o *VisualDataSeries) SetMetaData(v VisualMetaData) {
	o.MetaData = &v
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *VisualDataSeries) GetSeriesType() string {
	if o == nil || IsNil(o.SeriesType) {
		var ret string
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetSeriesTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesType) {
		return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *VisualDataSeries) HasSeriesType() bool {
	if o != nil && !IsNil(o.SeriesType) {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given string and assigns it to the SeriesType field.
func (o *VisualDataSeries) SetSeriesType(v string) {
	o.SeriesType = &v
}

// GetXAxisKeys returns the XAxisKeys field value if set, zero value otherwise.
func (o *VisualDataSeries) GetXAxisKeys() []string {
	if o == nil || IsNil(o.XAxisKeys) {
		var ret []string
		return ret
	}
	return o.XAxisKeys
}

// GetXAxisKeysOk returns a tuple with the XAxisKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetXAxisKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.XAxisKeys) {
		return nil, false
	}
	return o.XAxisKeys, true
}

// HasXAxisKeys returns a boolean if a field has been set.
func (o *VisualDataSeries) HasXAxisKeys() bool {
	if o != nil && !IsNil(o.XAxisKeys) {
		return true
	}

	return false
}

// SetXAxisKeys gets a reference to the given []string and assigns it to the XAxisKeys field.
func (o *VisualDataSeries) SetXAxisKeys(v []string) {
	o.XAxisKeys = v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *VisualDataSeries) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *VisualDataSeries) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *VisualDataSeries) SetValueType(v string) {
	o.ValueType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *VisualDataSeries) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *VisualDataSeries) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *VisualDataSeries) SetSource(v string) {
	o.Source = &v
}

// GetXAxisKeyTypes returns the XAxisKeyTypes field value if set, zero value otherwise.
func (o *VisualDataSeries) GetXAxisKeyTypes() map[string]string {
	if o == nil || IsNil(o.XAxisKeyTypes) {
		var ret map[string]string
		return ret
	}
	return *o.XAxisKeyTypes
}

// GetXAxisKeyTypesOk returns a tuple with the XAxisKeyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetXAxisKeyTypesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.XAxisKeyTypes) {
		return nil, false
	}
	return o.XAxisKeyTypes, true
}

// HasXAxisKeyTypes returns a boolean if a field has been set.
func (o *VisualDataSeries) HasXAxisKeyTypes() bool {
	if o != nil && !IsNil(o.XAxisKeyTypes) {
		return true
	}

	return false
}

// SetXAxisKeyTypes gets a reference to the given map[string]string and assigns it to the XAxisKeyTypes field.
func (o *VisualDataSeries) SetXAxisKeyTypes(v map[string]string) {
	o.XAxisKeyTypes = &v
}

// GetQueryInfo returns the QueryInfo field value if set, zero value otherwise.
func (o *VisualDataSeries) GetQueryInfo() MetricsQueryResultInfo {
	if o == nil || IsNil(o.QueryInfo) {
		var ret MetricsQueryResultInfo
		return ret
	}
	return *o.QueryInfo
}

// GetQueryInfoOk returns a tuple with the QueryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualDataSeries) GetQueryInfoOk() (*MetricsQueryResultInfo, bool) {
	if o == nil || IsNil(o.QueryInfo) {
		return nil, false
	}
	return o.QueryInfo, true
}

// HasQueryInfo returns a boolean if a field has been set.
func (o *VisualDataSeries) HasQueryInfo() bool {
	if o != nil && !IsNil(o.QueryInfo) {
		return true
	}

	return false
}

// SetQueryInfo gets a reference to the given MetricsQueryResultInfo and assigns it to the QueryInfo field.
func (o *VisualDataSeries) SetQueryInfo(v MetricsQueryResultInfo) {
	o.QueryInfo = &v
}

func (o VisualDataSeries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualDataSeries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryId"] = o.QueryId
	toSerialize["name"] = o.Name
	toSerialize["dataPoints"] = o.DataPoints
	if !IsNil(o.AggregateInfo) {
		toSerialize["aggregateInfo"] = o.AggregateInfo
	}
	if !IsNil(o.MetaData) {
		toSerialize["metaData"] = o.MetaData
	}
	if !IsNil(o.SeriesType) {
		toSerialize["seriesType"] = o.SeriesType
	}
	if !IsNil(o.XAxisKeys) {
		toSerialize["xAxisKeys"] = o.XAxisKeys
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.XAxisKeyTypes) {
		toSerialize["xAxisKeyTypes"] = o.XAxisKeyTypes
	}
	if !IsNil(o.QueryInfo) {
		toSerialize["queryInfo"] = o.QueryInfo
	}
	return toSerialize, nil
}

type NullableVisualDataSeries struct {
	value *VisualDataSeries
	isSet bool
}

func (v NullableVisualDataSeries) Get() *VisualDataSeries {
	return v.value
}

func (v *NullableVisualDataSeries) Set(val *VisualDataSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualDataSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualDataSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualDataSeries(val *VisualDataSeries) *NullableVisualDataSeries {
	return &NullableVisualDataSeries{value: val, isSet: true}
}

func (v NullableVisualDataSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualDataSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


