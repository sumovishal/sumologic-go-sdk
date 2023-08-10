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

// checks if the DashboardSyncDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardSyncDefinitionAllOf{}

// DashboardSyncDefinitionAllOf struct for DashboardSyncDefinitionAllOf
type DashboardSyncDefinitionAllOf struct {
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

// NewDashboardSyncDefinitionAllOf instantiates a new DashboardSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSyncDefinitionAllOf(description string, detailLevel int32, properties string, panels []ReportPanelSyncDefinition, filters []ReportFilterSyncDefinition) *DashboardSyncDefinitionAllOf {
	this := DashboardSyncDefinitionAllOf{}
	this.Description = description
	this.DetailLevel = detailLevel
	this.Properties = properties
	this.Panels = panels
	this.Filters = filters
	return &this
}

// NewDashboardSyncDefinitionAllOfWithDefaults instantiates a new DashboardSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSyncDefinitionAllOfWithDefaults() *DashboardSyncDefinitionAllOf {
	this := DashboardSyncDefinitionAllOf{}
	return &this
}

// GetDescription returns the Description field value
func (o *DashboardSyncDefinitionAllOf) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DashboardSyncDefinitionAllOf) SetDescription(v string) {
	o.Description = v
}

// GetDetailLevel returns the DetailLevel field value
func (o *DashboardSyncDefinitionAllOf) GetDetailLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DetailLevel
}

// GetDetailLevelOk returns a tuple with the DetailLevel field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinitionAllOf) GetDetailLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetailLevel, true
}

// SetDetailLevel sets field value
func (o *DashboardSyncDefinitionAllOf) SetDetailLevel(v int32) {
	o.DetailLevel = v
}

// GetProperties returns the Properties field value
func (o *DashboardSyncDefinitionAllOf) GetProperties() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinitionAllOf) GetPropertiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *DashboardSyncDefinitionAllOf) SetProperties(v string) {
	o.Properties = v
}

// GetPanels returns the Panels field value
func (o *DashboardSyncDefinitionAllOf) GetPanels() []ReportPanelSyncDefinition {
	if o == nil {
		var ret []ReportPanelSyncDefinition
		return ret
	}

	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinitionAllOf) GetPanelsOk() ([]ReportPanelSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Panels, true
}

// SetPanels sets field value
func (o *DashboardSyncDefinitionAllOf) SetPanels(v []ReportPanelSyncDefinition) {
	o.Panels = v
}

// GetFilters returns the Filters field value
func (o *DashboardSyncDefinitionAllOf) GetFilters() []ReportFilterSyncDefinition {
	if o == nil {
		var ret []ReportFilterSyncDefinition
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *DashboardSyncDefinitionAllOf) GetFiltersOk() ([]ReportFilterSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *DashboardSyncDefinitionAllOf) SetFilters(v []ReportFilterSyncDefinition) {
	o.Filters = v
}

func (o DashboardSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardSyncDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["detailLevel"] = o.DetailLevel
	toSerialize["properties"] = o.Properties
	toSerialize["panels"] = o.Panels
	toSerialize["filters"] = o.Filters
	return toSerialize, nil
}

type NullableDashboardSyncDefinitionAllOf struct {
	value *DashboardSyncDefinitionAllOf
	isSet bool
}

func (v NullableDashboardSyncDefinitionAllOf) Get() *DashboardSyncDefinitionAllOf {
	return v.value
}

func (v *NullableDashboardSyncDefinitionAllOf) Set(val *DashboardSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSyncDefinitionAllOf(val *DashboardSyncDefinitionAllOf) *NullableDashboardSyncDefinitionAllOf {
	return &NullableDashboardSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableDashboardSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


