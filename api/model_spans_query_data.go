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

// checks if the SpansQueryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpansQueryData{}

// SpansQueryData The data format describing a basic spans query.
type SpansQueryData struct {
	// A list of filters for the spans query.
	Filters []SpansFilter `json:"filters"`
	// A list of used visualization methods for the spans query.
	Visualizations []SpansVisualization `json:"visualizations"`
	// A list of group-by clauses for the spans query.
	GroupBy []SpansGroupBy `json:"groupBy"`
	// A list of limits that will be applied to the spans query.
	Limit []SpansLimitItem `json:"limit"`
}

type _SpansQueryData SpansQueryData

// NewSpansQueryData instantiates a new SpansQueryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpansQueryData(filters []SpansFilter, visualizations []SpansVisualization, groupBy []SpansGroupBy, limit []SpansLimitItem) *SpansQueryData {
	this := SpansQueryData{}
	this.Filters = filters
	this.Visualizations = visualizations
	this.GroupBy = groupBy
	this.Limit = limit
	return &this
}

// NewSpansQueryDataWithDefaults instantiates a new SpansQueryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpansQueryDataWithDefaults() *SpansQueryData {
	this := SpansQueryData{}
	return &this
}

// GetFilters returns the Filters field value
func (o *SpansQueryData) GetFilters() []SpansFilter {
	if o == nil {
		var ret []SpansFilter
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *SpansQueryData) GetFiltersOk() ([]SpansFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *SpansQueryData) SetFilters(v []SpansFilter) {
	o.Filters = v
}

// GetVisualizations returns the Visualizations field value
func (o *SpansQueryData) GetVisualizations() []SpansVisualization {
	if o == nil {
		var ret []SpansVisualization
		return ret
	}

	return o.Visualizations
}

// GetVisualizationsOk returns a tuple with the Visualizations field value
// and a boolean to check if the value has been set.
func (o *SpansQueryData) GetVisualizationsOk() ([]SpansVisualization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Visualizations, true
}

// SetVisualizations sets field value
func (o *SpansQueryData) SetVisualizations(v []SpansVisualization) {
	o.Visualizations = v
}

// GetGroupBy returns the GroupBy field value
func (o *SpansQueryData) GetGroupBy() []SpansGroupBy {
	if o == nil {
		var ret []SpansGroupBy
		return ret
	}

	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *SpansQueryData) GetGroupByOk() ([]SpansGroupBy, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// SetGroupBy sets field value
func (o *SpansQueryData) SetGroupBy(v []SpansGroupBy) {
	o.GroupBy = v
}

// GetLimit returns the Limit field value
func (o *SpansQueryData) GetLimit() []SpansLimitItem {
	if o == nil {
		var ret []SpansLimitItem
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *SpansQueryData) GetLimitOk() ([]SpansLimitItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit, true
}

// SetLimit sets field value
func (o *SpansQueryData) SetLimit(v []SpansLimitItem) {
	o.Limit = v
}

func (o SpansQueryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpansQueryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filters"] = o.Filters
	toSerialize["visualizations"] = o.Visualizations
	toSerialize["groupBy"] = o.GroupBy
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

func (o *SpansQueryData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filters",
		"visualizations",
		"groupBy",
		"limit",
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

	varSpansQueryData := _SpansQueryData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpansQueryData)

	if err != nil {
		return err
	}

	*o = SpansQueryData(varSpansQueryData)

	return err
}

type NullableSpansQueryData struct {
	value *SpansQueryData
	isSet bool
}

func (v NullableSpansQueryData) Get() *SpansQueryData {
	return v.value
}

func (v *NullableSpansQueryData) Set(val *SpansQueryData) {
	v.value = val
	v.isSet = true
}

func (v NullableSpansQueryData) IsSet() bool {
	return v.isSet
}

func (v *NullableSpansQueryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpansQueryData(val *SpansQueryData) *NullableSpansQueryData {
	return &NullableSpansQueryData{value: val, isSet: true}
}

func (v NullableSpansQueryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpansQueryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


