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

// checks if the DashboardRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardRequest{}

// DashboardRequest struct for DashboardRequest
type DashboardRequest struct {
	// Title of the dashboard.
	Title string `json:"title"`
	// Description of the dashboard.
	Description *string `json:"description,omitempty"`
	// The identifier of the folder to save the dashboard in. By default it is saved in your personal folder. 
	FolderId *string `json:"folderId,omitempty"`
	TopologyLabelMap *TopologyLabelMap `json:"topologyLabelMap,omitempty"`
	// If set denotes that the dashboard concerns a given domain (e.g. `aws`, `k8s`, `app`).
	Domain *string `json:"domain,omitempty"`
	// If set to non-empty array denotes that the dashboard concerns given hierarchies.
	Hierarchies []string `json:"hierarchies,omitempty"`
	// Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard. Allowed values are `0`, `30`, `60`, `120`, `300`, `900`, `1800`, `3600`, `7200`, `86400`. 
	RefreshInterval *int32 `json:"refreshInterval,omitempty"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// Panels in the dashboard.
	Panels []Panel `json:"panels,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
	// Variables to apply to the panels.
	Variables []Variable `json:"variables,omitempty"`
	// Theme for the dashboard. Either `Light` or `Dark`.
	Theme *string `json:"theme,omitempty" validate:"regexp=^(light|dark|Light|Dark)$"`
	// Is the dashboard public
	IsPublic *bool `json:"isPublic,omitempty"`
	// Whether to highlight threshold violations.
	HighlightViolations *bool `json:"highlightViolations,omitempty"`
}

type _DashboardRequest DashboardRequest

// NewDashboardRequest instantiates a new DashboardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardRequest(title string, timeRange ResolvableTimeRange) *DashboardRequest {
	this := DashboardRequest{}
	this.Title = title
	var domain string = ""
	this.Domain = &domain
	this.TimeRange = timeRange
	var theme string = "Light"
	this.Theme = &theme
	var isPublic bool = false
	this.IsPublic = &isPublic
	var highlightViolations bool = false
	this.HighlightViolations = &highlightViolations
	return &this
}

// NewDashboardRequestWithDefaults instantiates a new DashboardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardRequestWithDefaults() *DashboardRequest {
	this := DashboardRequest{}
	var domain string = ""
	this.Domain = &domain
	var theme string = "Light"
	this.Theme = &theme
	var isPublic bool = false
	this.IsPublic = &isPublic
	var highlightViolations bool = false
	this.HighlightViolations = &highlightViolations
	return &this
}

// GetTitle returns the Title field value
func (o *DashboardRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DashboardRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DashboardRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DashboardRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DashboardRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *DashboardRequest) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *DashboardRequest) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *DashboardRequest) SetFolderId(v string) {
	o.FolderId = &v
}

// GetTopologyLabelMap returns the TopologyLabelMap field value if set, zero value otherwise.
func (o *DashboardRequest) GetTopologyLabelMap() TopologyLabelMap {
	if o == nil || IsNil(o.TopologyLabelMap) {
		var ret TopologyLabelMap
		return ret
	}
	return *o.TopologyLabelMap
}

// GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetTopologyLabelMapOk() (*TopologyLabelMap, bool) {
	if o == nil || IsNil(o.TopologyLabelMap) {
		return nil, false
	}
	return o.TopologyLabelMap, true
}

// HasTopologyLabelMap returns a boolean if a field has been set.
func (o *DashboardRequest) HasTopologyLabelMap() bool {
	if o != nil && !IsNil(o.TopologyLabelMap) {
		return true
	}

	return false
}

// SetTopologyLabelMap gets a reference to the given TopologyLabelMap and assigns it to the TopologyLabelMap field.
func (o *DashboardRequest) SetTopologyLabelMap(v TopologyLabelMap) {
	o.TopologyLabelMap = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DashboardRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DashboardRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DashboardRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetHierarchies returns the Hierarchies field value if set, zero value otherwise.
func (o *DashboardRequest) GetHierarchies() []string {
	if o == nil || IsNil(o.Hierarchies) {
		var ret []string
		return ret
	}
	return o.Hierarchies
}

// GetHierarchiesOk returns a tuple with the Hierarchies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetHierarchiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Hierarchies) {
		return nil, false
	}
	return o.Hierarchies, true
}

// HasHierarchies returns a boolean if a field has been set.
func (o *DashboardRequest) HasHierarchies() bool {
	if o != nil && !IsNil(o.Hierarchies) {
		return true
	}

	return false
}

// SetHierarchies gets a reference to the given []string and assigns it to the Hierarchies field.
func (o *DashboardRequest) SetHierarchies(v []string) {
	o.Hierarchies = v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *DashboardRequest) GetRefreshInterval() int32 {
	if o == nil || IsNil(o.RefreshInterval) {
		var ret int32
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetRefreshIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshInterval) {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *DashboardRequest) HasRefreshInterval() bool {
	if o != nil && !IsNil(o.RefreshInterval) {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given int32 and assigns it to the RefreshInterval field.
func (o *DashboardRequest) SetRefreshInterval(v int32) {
	o.RefreshInterval = &v
}

// GetTimeRange returns the TimeRange field value
func (o *DashboardRequest) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *DashboardRequest) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *DashboardRequest) GetPanels() []Panel {
	if o == nil || IsNil(o.Panels) {
		var ret []Panel
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetPanelsOk() ([]Panel, bool) {
	if o == nil || IsNil(o.Panels) {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *DashboardRequest) HasPanels() bool {
	if o != nil && !IsNil(o.Panels) {
		return true
	}

	return false
}

// SetPanels gets a reference to the given []Panel and assigns it to the Panels field.
func (o *DashboardRequest) SetPanels(v []Panel) {
	o.Panels = v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *DashboardRequest) GetLayout() Layout {
	if o == nil || IsNil(o.Layout) {
		var ret Layout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetLayoutOk() (*Layout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *DashboardRequest) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given Layout and assigns it to the Layout field.
func (o *DashboardRequest) SetLayout(v Layout) {
	o.Layout = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *DashboardRequest) GetVariables() []Variable {
	if o == nil || IsNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetVariablesOk() ([]Variable, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *DashboardRequest) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *DashboardRequest) SetVariables(v []Variable) {
	o.Variables = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *DashboardRequest) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *DashboardRequest) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *DashboardRequest) SetTheme(v string) {
	o.Theme = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *DashboardRequest) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *DashboardRequest) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *DashboardRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetHighlightViolations returns the HighlightViolations field value if set, zero value otherwise.
func (o *DashboardRequest) GetHighlightViolations() bool {
	if o == nil || IsNil(o.HighlightViolations) {
		var ret bool
		return ret
	}
	return *o.HighlightViolations
}

// GetHighlightViolationsOk returns a tuple with the HighlightViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetHighlightViolationsOk() (*bool, bool) {
	if o == nil || IsNil(o.HighlightViolations) {
		return nil, false
	}
	return o.HighlightViolations, true
}

// HasHighlightViolations returns a boolean if a field has been set.
func (o *DashboardRequest) HasHighlightViolations() bool {
	if o != nil && !IsNil(o.HighlightViolations) {
		return true
	}

	return false
}

// SetHighlightViolations gets a reference to the given bool and assigns it to the HighlightViolations field.
func (o *DashboardRequest) SetHighlightViolations(v bool) {
	o.HighlightViolations = &v
}

func (o DashboardRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !IsNil(o.TopologyLabelMap) {
		toSerialize["topologyLabelMap"] = o.TopologyLabelMap
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Hierarchies) {
		toSerialize["hierarchies"] = o.Hierarchies
	}
	if !IsNil(o.RefreshInterval) {
		toSerialize["refreshInterval"] = o.RefreshInterval
	}
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.Panels) {
		toSerialize["panels"] = o.Panels
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	if !IsNil(o.HighlightViolations) {
		toSerialize["highlightViolations"] = o.HighlightViolations
	}
	return toSerialize, nil
}

func (o *DashboardRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varDashboardRequest := _DashboardRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDashboardRequest)

	if err != nil {
		return err
	}

	*o = DashboardRequest(varDashboardRequest)

	return err
}

type NullableDashboardRequest struct {
	value *DashboardRequest
	isSet bool
}

func (v NullableDashboardRequest) Get() *DashboardRequest {
	return v.value
}

func (v *NullableDashboardRequest) Set(val *DashboardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardRequest(val *DashboardRequest) *NullableDashboardRequest {
	return &NullableDashboardRequest{value: val, isSet: true}
}

func (v NullableDashboardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


