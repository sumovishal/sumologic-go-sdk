/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SumoSearchPanelAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SumoSearchPanelAllOf{}

// SumoSearchPanelAllOf A panel that has logs and metrics search queries.
type SumoSearchPanelAllOf struct {
	// Metrics and log queries of the panel.
	Queries []Query `json:"queries"`
	// Description of the panel.
	Description *string `json:"description,omitempty"`
	TimeRange *ResolvableTimeRange `json:"timeRange,omitempty"`
	// Rules to set the color of data.
	ColoringRules []ColoringRule `json:"coloringRules,omitempty"`
	// List of linked dashboards.
	LinkedDashboards []LinkedDashboard `json:"linkedDashboards,omitempty"`
}

// NewSumoSearchPanelAllOf instantiates a new SumoSearchPanelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSumoSearchPanelAllOf(queries []Query) *SumoSearchPanelAllOf {
	this := SumoSearchPanelAllOf{}
	this.Queries = queries
	return &this
}

// NewSumoSearchPanelAllOfWithDefaults instantiates a new SumoSearchPanelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSumoSearchPanelAllOfWithDefaults() *SumoSearchPanelAllOf {
	this := SumoSearchPanelAllOf{}
	return &this
}

// GetQueries returns the Queries field value
func (o *SumoSearchPanelAllOf) GetQueries() []Query {
	if o == nil {
		var ret []Query
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *SumoSearchPanelAllOf) GetQueriesOk() ([]Query, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *SumoSearchPanelAllOf) SetQueries(v []Query) {
	o.Queries = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SumoSearchPanelAllOf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoSearchPanelAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SumoSearchPanelAllOf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SumoSearchPanelAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *SumoSearchPanelAllOf) GetTimeRange() ResolvableTimeRange {
	if o == nil || IsNil(o.TimeRange) {
		var ret ResolvableTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoSearchPanelAllOf) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *SumoSearchPanelAllOf) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given ResolvableTimeRange and assigns it to the TimeRange field.
func (o *SumoSearchPanelAllOf) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = &v
}

// GetColoringRules returns the ColoringRules field value if set, zero value otherwise.
func (o *SumoSearchPanelAllOf) GetColoringRules() []ColoringRule {
	if o == nil || IsNil(o.ColoringRules) {
		var ret []ColoringRule
		return ret
	}
	return o.ColoringRules
}

// GetColoringRulesOk returns a tuple with the ColoringRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoSearchPanelAllOf) GetColoringRulesOk() ([]ColoringRule, bool) {
	if o == nil || IsNil(o.ColoringRules) {
		return nil, false
	}
	return o.ColoringRules, true
}

// HasColoringRules returns a boolean if a field has been set.
func (o *SumoSearchPanelAllOf) HasColoringRules() bool {
	if o != nil && !IsNil(o.ColoringRules) {
		return true
	}

	return false
}

// SetColoringRules gets a reference to the given []ColoringRule and assigns it to the ColoringRules field.
func (o *SumoSearchPanelAllOf) SetColoringRules(v []ColoringRule) {
	o.ColoringRules = v
}

// GetLinkedDashboards returns the LinkedDashboards field value if set, zero value otherwise.
func (o *SumoSearchPanelAllOf) GetLinkedDashboards() []LinkedDashboard {
	if o == nil || IsNil(o.LinkedDashboards) {
		var ret []LinkedDashboard
		return ret
	}
	return o.LinkedDashboards
}

// GetLinkedDashboardsOk returns a tuple with the LinkedDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoSearchPanelAllOf) GetLinkedDashboardsOk() ([]LinkedDashboard, bool) {
	if o == nil || IsNil(o.LinkedDashboards) {
		return nil, false
	}
	return o.LinkedDashboards, true
}

// HasLinkedDashboards returns a boolean if a field has been set.
func (o *SumoSearchPanelAllOf) HasLinkedDashboards() bool {
	if o != nil && !IsNil(o.LinkedDashboards) {
		return true
	}

	return false
}

// SetLinkedDashboards gets a reference to the given []LinkedDashboard and assigns it to the LinkedDashboards field.
func (o *SumoSearchPanelAllOf) SetLinkedDashboards(v []LinkedDashboard) {
	o.LinkedDashboards = v
}

func (o SumoSearchPanelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SumoSearchPanelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queries"] = o.Queries
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TimeRange) {
		toSerialize["timeRange"] = o.TimeRange
	}
	if !IsNil(o.ColoringRules) {
		toSerialize["coloringRules"] = o.ColoringRules
	}
	if !IsNil(o.LinkedDashboards) {
		toSerialize["linkedDashboards"] = o.LinkedDashboards
	}
	return toSerialize, nil
}

type NullableSumoSearchPanelAllOf struct {
	value *SumoSearchPanelAllOf
	isSet bool
}

func (v NullableSumoSearchPanelAllOf) Get() *SumoSearchPanelAllOf {
	return v.value
}

func (v *NullableSumoSearchPanelAllOf) Set(val *SumoSearchPanelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSumoSearchPanelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSumoSearchPanelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSumoSearchPanelAllOf(val *SumoSearchPanelAllOf) *NullableSumoSearchPanelAllOf {
	return &NullableSumoSearchPanelAllOf{value: val, isSet: true}
}

func (v NullableSumoSearchPanelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSumoSearchPanelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


