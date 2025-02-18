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

// checks if the LogSearchQueryEstimationQueryDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogSearchQueryEstimationQueryDefinition{}

// LogSearchQueryEstimationQueryDefinition Definition of the log search with query and timerange.
type LogSearchQueryEstimationQueryDefinition struct {
	// Log search Query to compute the estimated volume of data scanned.
	QueryString string `json:"queryString"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// This has the value `true` if the search is to be run by receipt time and `false` if it is to be run by message time.
	RunByReceiptTime *bool `json:"runByReceiptTime,omitempty"`
	// Values for search template used in the search query. Learn more about the search templates here : https://help.sumologic.com/docs/search/get-started-with-search/build-search/search-templates/
	QueryParameters []LogSearchQueryParameterSyncDefinitionBase `json:"queryParameters,omitempty"`
}

type _LogSearchQueryEstimationQueryDefinition LogSearchQueryEstimationQueryDefinition

// NewLogSearchQueryEstimationQueryDefinition instantiates a new LogSearchQueryEstimationQueryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchQueryEstimationQueryDefinition(queryString string, timeRange ResolvableTimeRange) *LogSearchQueryEstimationQueryDefinition {
	this := LogSearchQueryEstimationQueryDefinition{}
	this.QueryString = queryString
	this.TimeRange = timeRange
	var runByReceiptTime bool = false
	this.RunByReceiptTime = &runByReceiptTime
	return &this
}

// NewLogSearchQueryEstimationQueryDefinitionWithDefaults instantiates a new LogSearchQueryEstimationQueryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchQueryEstimationQueryDefinitionWithDefaults() *LogSearchQueryEstimationQueryDefinition {
	this := LogSearchQueryEstimationQueryDefinition{}
	var runByReceiptTime bool = false
	this.RunByReceiptTime = &runByReceiptTime
	return &this
}

// GetQueryString returns the QueryString field value
func (o *LogSearchQueryEstimationQueryDefinition) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *LogSearchQueryEstimationQueryDefinition) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *LogSearchQueryEstimationQueryDefinition) SetQueryString(v string) {
	o.QueryString = v
}

// GetTimeRange returns the TimeRange field value
func (o *LogSearchQueryEstimationQueryDefinition) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *LogSearchQueryEstimationQueryDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *LogSearchQueryEstimationQueryDefinition) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetRunByReceiptTime returns the RunByReceiptTime field value if set, zero value otherwise.
func (o *LogSearchQueryEstimationQueryDefinition) GetRunByReceiptTime() bool {
	if o == nil || IsNil(o.RunByReceiptTime) {
		var ret bool
		return ret
	}
	return *o.RunByReceiptTime
}

// GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchQueryEstimationQueryDefinition) GetRunByReceiptTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.RunByReceiptTime) {
		return nil, false
	}
	return o.RunByReceiptTime, true
}

// HasRunByReceiptTime returns a boolean if a field has been set.
func (o *LogSearchQueryEstimationQueryDefinition) HasRunByReceiptTime() bool {
	if o != nil && !IsNil(o.RunByReceiptTime) {
		return true
	}

	return false
}

// SetRunByReceiptTime gets a reference to the given bool and assigns it to the RunByReceiptTime field.
func (o *LogSearchQueryEstimationQueryDefinition) SetRunByReceiptTime(v bool) {
	o.RunByReceiptTime = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *LogSearchQueryEstimationQueryDefinition) GetQueryParameters() []LogSearchQueryParameterSyncDefinitionBase {
	if o == nil || IsNil(o.QueryParameters) {
		var ret []LogSearchQueryParameterSyncDefinitionBase
		return ret
	}
	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchQueryEstimationQueryDefinition) GetQueryParametersOk() ([]LogSearchQueryParameterSyncDefinitionBase, bool) {
	if o == nil || IsNil(o.QueryParameters) {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *LogSearchQueryEstimationQueryDefinition) HasQueryParameters() bool {
	if o != nil && !IsNil(o.QueryParameters) {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given []LogSearchQueryParameterSyncDefinitionBase and assigns it to the QueryParameters field.
func (o *LogSearchQueryEstimationQueryDefinition) SetQueryParameters(v []LogSearchQueryParameterSyncDefinitionBase) {
	o.QueryParameters = v
}

func (o LogSearchQueryEstimationQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogSearchQueryEstimationQueryDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryString"] = o.QueryString
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.RunByReceiptTime) {
		toSerialize["runByReceiptTime"] = o.RunByReceiptTime
	}
	if !IsNil(o.QueryParameters) {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	return toSerialize, nil
}

func (o *LogSearchQueryEstimationQueryDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryString",
		"timeRange",
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

	varLogSearchQueryEstimationQueryDefinition := _LogSearchQueryEstimationQueryDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogSearchQueryEstimationQueryDefinition)

	if err != nil {
		return err
	}

	*o = LogSearchQueryEstimationQueryDefinition(varLogSearchQueryEstimationQueryDefinition)

	return err
}

type NullableLogSearchQueryEstimationQueryDefinition struct {
	value *LogSearchQueryEstimationQueryDefinition
	isSet bool
}

func (v NullableLogSearchQueryEstimationQueryDefinition) Get() *LogSearchQueryEstimationQueryDefinition {
	return v.value
}

func (v *NullableLogSearchQueryEstimationQueryDefinition) Set(val *LogSearchQueryEstimationQueryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchQueryEstimationQueryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchQueryEstimationQueryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchQueryEstimationQueryDefinition(val *LogSearchQueryEstimationQueryDefinition) *NullableLogSearchQueryEstimationQueryDefinition {
	return &NullableLogSearchQueryEstimationQueryDefinition{value: val, isSet: true}
}

func (v NullableLogSearchQueryEstimationQueryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchQueryEstimationQueryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


