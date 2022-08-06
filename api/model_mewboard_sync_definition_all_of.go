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

// MewboardSyncDefinitionAllOf struct for MewboardSyncDefinitionAllOf
type MewboardSyncDefinitionAllOf struct {
	// A description of the dashboard.
	Description *string `json:"description,omitempty"`
	// The title of the dashboard.
	Title string `json:"title"`
	// Theme for the dashboard. Must be `light` or `dark`.
	Theme *string `json:"theme,omitempty"`
	TopologyLabelMap *TopologyLabelMap `json:"topologyLabelMap,omitempty"`
	// Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard.
	RefreshInterval *int32 `json:"refreshInterval,omitempty"`
	TimeRange *ResolvableTimeRange `json:"timeRange,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
	// Children panels that the container panel contains.
	Panels []Panel `json:"panels,omitempty"`
	// Variables that could be applied to the panel's children.
	Variables []Variable `json:"variables,omitempty"`
	// Coloring rules to color the panel/data with.
	ColoringRules []ColoringRule `json:"coloringRules,omitempty"`
}

// NewMewboardSyncDefinitionAllOf instantiates a new MewboardSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMewboardSyncDefinitionAllOf(title string) *MewboardSyncDefinitionAllOf {
	this := MewboardSyncDefinitionAllOf{}
	this.Title = title
	var theme string = "light"
	this.Theme = &theme
	return &this
}

// NewMewboardSyncDefinitionAllOfWithDefaults instantiates a new MewboardSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMewboardSyncDefinitionAllOfWithDefaults() *MewboardSyncDefinitionAllOf {
	this := MewboardSyncDefinitionAllOf{}
	var theme string = "light"
	this.Theme = &theme
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MewboardSyncDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value
func (o *MewboardSyncDefinitionAllOf) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *MewboardSyncDefinitionAllOf) SetTitle(v string) {
	o.Title = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetTheme() string {
	if o == nil || o.Theme == nil {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetThemeOk() (*string, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasTheme() bool {
	if o != nil && o.Theme != nil {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *MewboardSyncDefinitionAllOf) SetTheme(v string) {
	o.Theme = &v
}

// GetTopologyLabelMap returns the TopologyLabelMap field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetTopologyLabelMap() TopologyLabelMap {
	if o == nil || o.TopologyLabelMap == nil {
		var ret TopologyLabelMap
		return ret
	}
	return *o.TopologyLabelMap
}

// GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetTopologyLabelMapOk() (*TopologyLabelMap, bool) {
	if o == nil || o.TopologyLabelMap == nil {
		return nil, false
	}
	return o.TopologyLabelMap, true
}

// HasTopologyLabelMap returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasTopologyLabelMap() bool {
	if o != nil && o.TopologyLabelMap != nil {
		return true
	}

	return false
}

// SetTopologyLabelMap gets a reference to the given TopologyLabelMap and assigns it to the TopologyLabelMap field.
func (o *MewboardSyncDefinitionAllOf) SetTopologyLabelMap(v TopologyLabelMap) {
	o.TopologyLabelMap = &v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetRefreshInterval() int32 {
	if o == nil || o.RefreshInterval == nil {
		var ret int32
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetRefreshIntervalOk() (*int32, bool) {
	if o == nil || o.RefreshInterval == nil {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasRefreshInterval() bool {
	if o != nil && o.RefreshInterval != nil {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given int32 and assigns it to the RefreshInterval field.
func (o *MewboardSyncDefinitionAllOf) SetRefreshInterval(v int32) {
	o.RefreshInterval = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetTimeRange() ResolvableTimeRange {
	if o == nil || o.TimeRange == nil {
		var ret ResolvableTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil || o.TimeRange == nil {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasTimeRange() bool {
	if o != nil && o.TimeRange != nil {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given ResolvableTimeRange and assigns it to the TimeRange field.
func (o *MewboardSyncDefinitionAllOf) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetLayout() Layout {
	if o == nil || o.Layout == nil {
		var ret Layout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetLayoutOk() (*Layout, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given Layout and assigns it to the Layout field.
func (o *MewboardSyncDefinitionAllOf) SetLayout(v Layout) {
	o.Layout = &v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetPanels() []Panel {
	if o == nil || o.Panels == nil {
		var ret []Panel
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetPanelsOk() ([]Panel, bool) {
	if o == nil || o.Panels == nil {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasPanels() bool {
	if o != nil && o.Panels != nil {
		return true
	}

	return false
}

// SetPanels gets a reference to the given []Panel and assigns it to the Panels field.
func (o *MewboardSyncDefinitionAllOf) SetPanels(v []Panel) {
	o.Panels = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetVariables() []Variable {
	if o == nil || o.Variables == nil {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetVariablesOk() ([]Variable, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *MewboardSyncDefinitionAllOf) SetVariables(v []Variable) {
	o.Variables = v
}

// GetColoringRules returns the ColoringRules field value if set, zero value otherwise.
func (o *MewboardSyncDefinitionAllOf) GetColoringRules() []ColoringRule {
	if o == nil || o.ColoringRules == nil {
		var ret []ColoringRule
		return ret
	}
	return o.ColoringRules
}

// GetColoringRulesOk returns a tuple with the ColoringRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinitionAllOf) GetColoringRulesOk() ([]ColoringRule, bool) {
	if o == nil || o.ColoringRules == nil {
		return nil, false
	}
	return o.ColoringRules, true
}

// HasColoringRules returns a boolean if a field has been set.
func (o *MewboardSyncDefinitionAllOf) HasColoringRules() bool {
	if o != nil && o.ColoringRules != nil {
		return true
	}

	return false
}

// SetColoringRules gets a reference to the given []ColoringRule and assigns it to the ColoringRules field.
func (o *MewboardSyncDefinitionAllOf) SetColoringRules(v []ColoringRule) {
	o.ColoringRules = v
}

func (o MewboardSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Theme != nil {
		toSerialize["theme"] = o.Theme
	}
	if o.TopologyLabelMap != nil {
		toSerialize["topologyLabelMap"] = o.TopologyLabelMap
	}
	if o.RefreshInterval != nil {
		toSerialize["refreshInterval"] = o.RefreshInterval
	}
	if o.TimeRange != nil {
		toSerialize["timeRange"] = o.TimeRange
	}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.Panels != nil {
		toSerialize["panels"] = o.Panels
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.ColoringRules != nil {
		toSerialize["coloringRules"] = o.ColoringRules
	}
	return json.Marshal(toSerialize)
}

type NullableMewboardSyncDefinitionAllOf struct {
	value *MewboardSyncDefinitionAllOf
	isSet bool
}

func (v NullableMewboardSyncDefinitionAllOf) Get() *MewboardSyncDefinitionAllOf {
	return v.value
}

func (v *NullableMewboardSyncDefinitionAllOf) Set(val *MewboardSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMewboardSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMewboardSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMewboardSyncDefinitionAllOf(val *MewboardSyncDefinitionAllOf) *NullableMewboardSyncDefinitionAllOf {
	return &NullableMewboardSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableMewboardSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMewboardSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


