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

// checks if the DashboardV2SyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardV2SyncDefinition{}

// DashboardV2SyncDefinition struct for DashboardV2SyncDefinition
type DashboardV2SyncDefinition struct {
	ContentSyncDefinition
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

// NewDashboardV2SyncDefinition instantiates a new DashboardV2SyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardV2SyncDefinition(title string, type_ string, name string) *DashboardV2SyncDefinition {
	this := DashboardV2SyncDefinition{}
	this.Type = type_
	this.Name = name
	this.Title = title
	var theme string = "light"
	this.Theme = &theme
	return &this
}

// NewDashboardV2SyncDefinitionWithDefaults instantiates a new DashboardV2SyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardV2SyncDefinitionWithDefaults() *DashboardV2SyncDefinition {
	this := DashboardV2SyncDefinition{}
	var theme string = "light"
	this.Theme = &theme
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DashboardV2SyncDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value
func (o *DashboardV2SyncDefinition) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DashboardV2SyncDefinition) SetTitle(v string) {
	o.Title = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *DashboardV2SyncDefinition) SetTheme(v string) {
	o.Theme = &v
}

// GetTopologyLabelMap returns the TopologyLabelMap field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetTopologyLabelMap() TopologyLabelMap {
	if o == nil || IsNil(o.TopologyLabelMap) {
		var ret TopologyLabelMap
		return ret
	}
	return *o.TopologyLabelMap
}

// GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetTopologyLabelMapOk() (*TopologyLabelMap, bool) {
	if o == nil || IsNil(o.TopologyLabelMap) {
		return nil, false
	}
	return o.TopologyLabelMap, true
}

// HasTopologyLabelMap returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasTopologyLabelMap() bool {
	if o != nil && !IsNil(o.TopologyLabelMap) {
		return true
	}

	return false
}

// SetTopologyLabelMap gets a reference to the given TopologyLabelMap and assigns it to the TopologyLabelMap field.
func (o *DashboardV2SyncDefinition) SetTopologyLabelMap(v TopologyLabelMap) {
	o.TopologyLabelMap = &v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetRefreshInterval() int32 {
	if o == nil || IsNil(o.RefreshInterval) {
		var ret int32
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetRefreshIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshInterval) {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasRefreshInterval() bool {
	if o != nil && !IsNil(o.RefreshInterval) {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given int32 and assigns it to the RefreshInterval field.
func (o *DashboardV2SyncDefinition) SetRefreshInterval(v int32) {
	o.RefreshInterval = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetTimeRange() ResolvableTimeRange {
	if o == nil || IsNil(o.TimeRange) {
		var ret ResolvableTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given ResolvableTimeRange and assigns it to the TimeRange field.
func (o *DashboardV2SyncDefinition) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetLayout() Layout {
	if o == nil || IsNil(o.Layout) {
		var ret Layout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetLayoutOk() (*Layout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given Layout and assigns it to the Layout field.
func (o *DashboardV2SyncDefinition) SetLayout(v Layout) {
	o.Layout = &v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetPanels() []Panel {
	if o == nil || IsNil(o.Panels) {
		var ret []Panel
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetPanelsOk() ([]Panel, bool) {
	if o == nil || IsNil(o.Panels) {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasPanels() bool {
	if o != nil && !IsNil(o.Panels) {
		return true
	}

	return false
}

// SetPanels gets a reference to the given []Panel and assigns it to the Panels field.
func (o *DashboardV2SyncDefinition) SetPanels(v []Panel) {
	o.Panels = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetVariables() []Variable {
	if o == nil || IsNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetVariablesOk() ([]Variable, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *DashboardV2SyncDefinition) SetVariables(v []Variable) {
	o.Variables = v
}

// GetColoringRules returns the ColoringRules field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetColoringRules() []ColoringRule {
	if o == nil || IsNil(o.ColoringRules) {
		var ret []ColoringRule
		return ret
	}
	return o.ColoringRules
}

// GetColoringRulesOk returns a tuple with the ColoringRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetColoringRulesOk() ([]ColoringRule, bool) {
	if o == nil || IsNil(o.ColoringRules) {
		return nil, false
	}
	return o.ColoringRules, true
}

// HasColoringRules returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasColoringRules() bool {
	if o != nil && !IsNil(o.ColoringRules) {
		return true
	}

	return false
}

// SetColoringRules gets a reference to the given []ColoringRule and assigns it to the ColoringRules field.
func (o *DashboardV2SyncDefinition) SetColoringRules(v []ColoringRule) {
	o.ColoringRules = v
}

func (o DashboardV2SyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardV2SyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContentSyncDefinition, errContentSyncDefinition := json.Marshal(o.ContentSyncDefinition)
	if errContentSyncDefinition != nil {
		return map[string]interface{}{}, errContentSyncDefinition
	}
	errContentSyncDefinition = json.Unmarshal([]byte(serializedContentSyncDefinition), &toSerialize)
	if errContentSyncDefinition != nil {
		return map[string]interface{}{}, errContentSyncDefinition
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["title"] = o.Title
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.TopologyLabelMap) {
		toSerialize["topologyLabelMap"] = o.TopologyLabelMap
	}
	if !IsNil(o.RefreshInterval) {
		toSerialize["refreshInterval"] = o.RefreshInterval
	}
	if !IsNil(o.TimeRange) {
		toSerialize["timeRange"] = o.TimeRange
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.Panels) {
		toSerialize["panels"] = o.Panels
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !IsNil(o.ColoringRules) {
		toSerialize["coloringRules"] = o.ColoringRules
	}
	return toSerialize, nil
}

type NullableDashboardV2SyncDefinition struct {
	value *DashboardV2SyncDefinition
	isSet bool
}

func (v NullableDashboardV2SyncDefinition) Get() *DashboardV2SyncDefinition {
	return v.value
}

func (v *NullableDashboardV2SyncDefinition) Set(val *DashboardV2SyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardV2SyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardV2SyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardV2SyncDefinition(val *DashboardV2SyncDefinition) *NullableDashboardV2SyncDefinition {
	return &NullableDashboardV2SyncDefinition{value: val, isSet: true}
}

func (v NullableDashboardV2SyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardV2SyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


