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

// checks if the DashboardSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardSyncDefinition{}

// DashboardSyncDefinition struct for DashboardSyncDefinition
type DashboardSyncDefinition struct {
	ContentSyncDefinition
	// A description of the dashboard.
	Description string `json:"description"`
	// Supported values are:   - `1` for small   - `2` for medium   - `3` for large
	DetailLevel int32 `json:"detailLevel"`
	// Visual settings for the panel.
	Properties string `json:"properties"`
	// The panels of the dashboard. _Dashboard links are not supported._
	Panels []ReportPanelSyncDefinition `json:"panels"`
	// The filters for the dashboard. Filters allow you to control the amount of information displayed in your dashboards.
	Filters []ReportFilterSyncDefinition `json:"filters"`
}

// NewDashboardSyncDefinition instantiates a new DashboardSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSyncDefinition(description string, detailLevel int32, properties string, panels []ReportPanelSyncDefinition, filters []ReportFilterSyncDefinition, type_ string, name string) *DashboardSyncDefinition {
	this := DashboardSyncDefinition{}
	this.Type = type_
	this.Name = name
	this.Description = description
	this.DetailLevel = detailLevel
	this.Properties = properties
	this.Panels = panels
	this.Filters = filters
	return &this
}

// NewDashboardSyncDefinitionWithDefaults instantiates a new DashboardSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSyncDefinitionWithDefaults() *DashboardSyncDefinition {
	this := DashboardSyncDefinition{}
	return &this
}

// GetDescription returns the Description field value
func (o *DashboardSyncDefinition) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DashboardSyncDefinition) SetDescription(v string) {
	o.Description = v
}

// GetDetailLevel returns the DetailLevel field value
func (o *DashboardSyncDefinition) GetDetailLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DetailLevel
}

// GetDetailLevelOk returns a tuple with the DetailLevel field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinition) GetDetailLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetailLevel, true
}

// SetDetailLevel sets field value
func (o *DashboardSyncDefinition) SetDetailLevel(v int32) {
	o.DetailLevel = v
}

// GetProperties returns the Properties field value
func (o *DashboardSyncDefinition) GetProperties() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinition) GetPropertiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *DashboardSyncDefinition) SetProperties(v string) {
	o.Properties = v
}

// GetPanels returns the Panels field value
func (o *DashboardSyncDefinition) GetPanels() []ReportPanelSyncDefinition {
	if o == nil {
		var ret []ReportPanelSyncDefinition
		return ret
	}

	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinition) GetPanelsOk() ([]ReportPanelSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Panels, true
}

// SetPanels sets field value
func (o *DashboardSyncDefinition) SetPanels(v []ReportPanelSyncDefinition) {
	o.Panels = v
}

// GetFilters returns the Filters field value
func (o *DashboardSyncDefinition) GetFilters() []ReportFilterSyncDefinition {
	if o == nil {
		var ret []ReportFilterSyncDefinition
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinition) GetFiltersOk() ([]ReportFilterSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *DashboardSyncDefinition) SetFilters(v []ReportFilterSyncDefinition) {
	o.Filters = v
}

func (o DashboardSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContentSyncDefinition, errContentSyncDefinition := json.Marshal(o.ContentSyncDefinition)
	if errContentSyncDefinition != nil {
		return map[string]interface{}{}, errContentSyncDefinition
	}
	errContentSyncDefinition = json.Unmarshal([]byte(serializedContentSyncDefinition), &toSerialize)
	if errContentSyncDefinition != nil {
		return map[string]interface{}{}, errContentSyncDefinition
	}
	toSerialize["description"] = o.Description
	toSerialize["detailLevel"] = o.DetailLevel
	toSerialize["properties"] = o.Properties
	toSerialize["panels"] = o.Panels
	toSerialize["filters"] = o.Filters
	return toSerialize, nil
}

type NullableDashboardSyncDefinition struct {
	value *DashboardSyncDefinition
	isSet bool
}

func (v NullableDashboardSyncDefinition) Get() *DashboardSyncDefinition {
	return v.value
}

func (v *NullableDashboardSyncDefinition) Set(val *DashboardSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSyncDefinition(val *DashboardSyncDefinition) *NullableDashboardSyncDefinition {
	return &NullableDashboardSyncDefinition{value: val, isSet: true}
}

func (v NullableDashboardSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


