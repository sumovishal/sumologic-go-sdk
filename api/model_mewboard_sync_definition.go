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

// checks if the MewboardSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MewboardSyncDefinition{}

// MewboardSyncDefinition struct for MewboardSyncDefinition
type MewboardSyncDefinition struct {
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

// NewMewboardSyncDefinition instantiates a new MewboardSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMewboardSyncDefinition(title string, type_ string, name string) *MewboardSyncDefinition {
	this := MewboardSyncDefinition{}
	this.Type = type_
	this.Name = name
	this.Title = title
	var theme string = "light"
	this.Theme = &theme
	return &this
}

// NewMewboardSyncDefinitionWithDefaults instantiates a new MewboardSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMewboardSyncDefinitionWithDefaults() *MewboardSyncDefinition {
	this := MewboardSyncDefinition{}
	var theme string = "light"
	this.Theme = &theme
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MewboardSyncDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value
func (o *MewboardSyncDefinition) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *MewboardSyncDefinition) SetTitle(v string) {
	o.Title = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *MewboardSyncDefinition) SetTheme(v string) {
	o.Theme = &v
}

// GetTopologyLabelMap returns the TopologyLabelMap field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetTopologyLabelMap() TopologyLabelMap {
	if o == nil || IsNil(o.TopologyLabelMap) {
		var ret TopologyLabelMap
		return ret
	}
	return *o.TopologyLabelMap
}

// GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetTopologyLabelMapOk() (*TopologyLabelMap, bool) {
	if o == nil || IsNil(o.TopologyLabelMap) {
		return nil, false
	}
	return o.TopologyLabelMap, true
}

// HasTopologyLabelMap returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasTopologyLabelMap() bool {
	if o != nil && !IsNil(o.TopologyLabelMap) {
		return true
	}

	return false
}

// SetTopologyLabelMap gets a reference to the given TopologyLabelMap and assigns it to the TopologyLabelMap field.
func (o *MewboardSyncDefinition) SetTopologyLabelMap(v TopologyLabelMap) {
	o.TopologyLabelMap = &v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetRefreshInterval() int32 {
	if o == nil || IsNil(o.RefreshInterval) {
		var ret int32
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetRefreshIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshInterval) {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasRefreshInterval() bool {
	if o != nil && !IsNil(o.RefreshInterval) {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given int32 and assigns it to the RefreshInterval field.
func (o *MewboardSyncDefinition) SetRefreshInterval(v int32) {
	o.RefreshInterval = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetTimeRange() ResolvableTimeRange {
	if o == nil || IsNil(o.TimeRange) {
		var ret ResolvableTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given ResolvableTimeRange and assigns it to the TimeRange field.
func (o *MewboardSyncDefinition) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetLayout() Layout {
	if o == nil || IsNil(o.Layout) {
		var ret Layout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetLayoutOk() (*Layout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given Layout and assigns it to the Layout field.
func (o *MewboardSyncDefinition) SetLayout(v Layout) {
	o.Layout = &v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetPanels() []Panel {
	if o == nil || IsNil(o.Panels) {
		var ret []Panel
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetPanelsOk() ([]Panel, bool) {
	if o == nil || IsNil(o.Panels) {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasPanels() bool {
	if o != nil && !IsNil(o.Panels) {
		return true
	}

	return false
}

// SetPanels gets a reference to the given []Panel and assigns it to the Panels field.
func (o *MewboardSyncDefinition) SetPanels(v []Panel) {
	o.Panels = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetVariables() []Variable {
	if o == nil || IsNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetVariablesOk() ([]Variable, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *MewboardSyncDefinition) SetVariables(v []Variable) {
	o.Variables = v
}

// GetColoringRules returns the ColoringRules field value if set, zero value otherwise.
func (o *MewboardSyncDefinition) GetColoringRules() []ColoringRule {
	if o == nil || IsNil(o.ColoringRules) {
		var ret []ColoringRule
		return ret
	}
	return o.ColoringRules
}

// GetColoringRulesOk returns a tuple with the ColoringRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MewboardSyncDefinition) GetColoringRulesOk() ([]ColoringRule, bool) {
	if o == nil || IsNil(o.ColoringRules) {
		return nil, false
	}
	return o.ColoringRules, true
}

// HasColoringRules returns a boolean if a field has been set.
func (o *MewboardSyncDefinition) HasColoringRules() bool {
	if o != nil && !IsNil(o.ColoringRules) {
		return true
	}

	return false
}

// SetColoringRules gets a reference to the given []ColoringRule and assigns it to the ColoringRules field.
func (o *MewboardSyncDefinition) SetColoringRules(v []ColoringRule) {
	o.ColoringRules = v
}

func (o MewboardSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MewboardSyncDefinition) ToMap() (map[string]interface{}, error) {
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

type NullableMewboardSyncDefinition struct {
	value *MewboardSyncDefinition
	isSet bool
}

func (v NullableMewboardSyncDefinition) Get() *MewboardSyncDefinition {
	return v.value
}

func (v *NullableMewboardSyncDefinition) Set(val *MewboardSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMewboardSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMewboardSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMewboardSyncDefinition(val *MewboardSyncDefinition) *NullableMewboardSyncDefinition {
	return &NullableMewboardSyncDefinition{value: val, isSet: true}
}

func (v NullableMewboardSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMewboardSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


