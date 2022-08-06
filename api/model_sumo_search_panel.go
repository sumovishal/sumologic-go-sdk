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

// SumoSearchPanel struct for SumoSearchPanel
type SumoSearchPanel struct {
	Panel
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

// NewSumoSearchPanel instantiates a new SumoSearchPanel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSumoSearchPanel(queries []Query, key string, panelType string) *SumoSearchPanel {
	this := SumoSearchPanel{}
	this.Key = key
	var keepVisualSettingsConsistentWithParent bool = true
	this.KeepVisualSettingsConsistentWithParent = &keepVisualSettingsConsistentWithParent
	this.PanelType = panelType
	this.Queries = queries
	return &this
}

// NewSumoSearchPanelWithDefaults instantiates a new SumoSearchPanel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSumoSearchPanelWithDefaults() *SumoSearchPanel {
	this := SumoSearchPanel{}
	return &this
}

// GetQueries returns the Queries field value
func (o *SumoSearchPanel) GetQueries() []Query {
	if o == nil {
		var ret []Query
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *SumoSearchPanel) GetQueriesOk() ([]Query, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *SumoSearchPanel) SetQueries(v []Query) {
	o.Queries = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SumoSearchPanel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoSearchPanel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SumoSearchPanel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SumoSearchPanel) SetDescription(v string) {
	o.Description = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *SumoSearchPanel) GetTimeRange() ResolvableTimeRange {
	if o == nil || o.TimeRange == nil {
		var ret ResolvableTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoSearchPanel) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil || o.TimeRange == nil {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *SumoSearchPanel) HasTimeRange() bool {
	if o != nil && o.TimeRange != nil {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given ResolvableTimeRange and assigns it to the TimeRange field.
func (o *SumoSearchPanel) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = &v
}

// GetColoringRules returns the ColoringRules field value if set, zero value otherwise.
func (o *SumoSearchPanel) GetColoringRules() []ColoringRule {
	if o == nil || o.ColoringRules == nil {
		var ret []ColoringRule
		return ret
	}
	return o.ColoringRules
}

// GetColoringRulesOk returns a tuple with the ColoringRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoSearchPanel) GetColoringRulesOk() ([]ColoringRule, bool) {
	if o == nil || o.ColoringRules == nil {
		return nil, false
	}
	return o.ColoringRules, true
}

// HasColoringRules returns a boolean if a field has been set.
func (o *SumoSearchPanel) HasColoringRules() bool {
	if o != nil && o.ColoringRules != nil {
		return true
	}

	return false
}

// SetColoringRules gets a reference to the given []ColoringRule and assigns it to the ColoringRules field.
func (o *SumoSearchPanel) SetColoringRules(v []ColoringRule) {
	o.ColoringRules = v
}

// GetLinkedDashboards returns the LinkedDashboards field value if set, zero value otherwise.
func (o *SumoSearchPanel) GetLinkedDashboards() []LinkedDashboard {
	if o == nil || o.LinkedDashboards == nil {
		var ret []LinkedDashboard
		return ret
	}
	return o.LinkedDashboards
}

// GetLinkedDashboardsOk returns a tuple with the LinkedDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SumoSearchPanel) GetLinkedDashboardsOk() ([]LinkedDashboard, bool) {
	if o == nil || o.LinkedDashboards == nil {
		return nil, false
	}
	return o.LinkedDashboards, true
}

// HasLinkedDashboards returns a boolean if a field has been set.
func (o *SumoSearchPanel) HasLinkedDashboards() bool {
	if o != nil && o.LinkedDashboards != nil {
		return true
	}

	return false
}

// SetLinkedDashboards gets a reference to the given []LinkedDashboard and assigns it to the LinkedDashboards field.
func (o *SumoSearchPanel) SetLinkedDashboards(v []LinkedDashboard) {
	o.LinkedDashboards = v
}

func (o SumoSearchPanel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPanel, errPanel := json.Marshal(o.Panel)
	if errPanel != nil {
		return []byte{}, errPanel
	}
	errPanel = json.Unmarshal([]byte(serializedPanel), &toSerialize)
	if errPanel != nil {
		return []byte{}, errPanel
	}
	if true {
		toSerialize["queries"] = o.Queries
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TimeRange != nil {
		toSerialize["timeRange"] = o.TimeRange
	}
	if o.ColoringRules != nil {
		toSerialize["coloringRules"] = o.ColoringRules
	}
	if o.LinkedDashboards != nil {
		toSerialize["linkedDashboards"] = o.LinkedDashboards
	}
	return json.Marshal(toSerialize)
}

type NullableSumoSearchPanel struct {
	value *SumoSearchPanel
	isSet bool
}

func (v NullableSumoSearchPanel) Get() *SumoSearchPanel {
	return v.value
}

func (v *NullableSumoSearchPanel) Set(val *SumoSearchPanel) {
	v.value = val
	v.isSet = true
}

func (v NullableSumoSearchPanel) IsSet() bool {
	return v.isSet
}

func (v *NullableSumoSearchPanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSumoSearchPanel(val *SumoSearchPanel) *NullableSumoSearchPanel {
	return &NullableSumoSearchPanel{value: val, isSet: true}
}

func (v NullableSumoSearchPanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSumoSearchPanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


