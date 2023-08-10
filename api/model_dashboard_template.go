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

// checks if the DashboardTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardTemplate{}

// DashboardTemplate struct for DashboardTemplate
type DashboardTemplate struct {
	Template
	// Id of the dashboard.
	Id string `json:"id"`
	// A map of panel to session id. The session id will be used to fetch data of the panel for the report. If not specified, a new session id will be created for the panel. 
	PanelToSessionIdMap *map[string]string `json:"panelToSessionIdMap,omitempty"`
	TimeRange *ResolvableTimeRange `json:"timeRange,omitempty"`
	VariableValues *VariablesValuesData `json:"variableValues,omitempty"`
}

// NewDashboardTemplate instantiates a new DashboardTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardTemplate(id string, templateType string) *DashboardTemplate {
	this := DashboardTemplate{}
	this.TemplateType = templateType
	this.Id = id
	return &this
}

// NewDashboardTemplateWithDefaults instantiates a new DashboardTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardTemplateWithDefaults() *DashboardTemplate {
	this := DashboardTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *DashboardTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DashboardTemplate) SetId(v string) {
	o.Id = v
}

// GetPanelToSessionIdMap returns the PanelToSessionIdMap field value if set, zero value otherwise.
func (o *DashboardTemplate) GetPanelToSessionIdMap() map[string]string {
	if o == nil || IsNil(o.PanelToSessionIdMap) {
		var ret map[string]string
		return ret
	}
	return *o.PanelToSessionIdMap
}

// GetPanelToSessionIdMapOk returns a tuple with the PanelToSessionIdMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardTemplate) GetPanelToSessionIdMapOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PanelToSessionIdMap) {
		return nil, false
	}
	return o.PanelToSessionIdMap, true
}

// HasPanelToSessionIdMap returns a boolean if a field has been set.
func (o *DashboardTemplate) HasPanelToSessionIdMap() bool {
	if o != nil && !IsNil(o.PanelToSessionIdMap) {
		return true
	}

	return false
}

// SetPanelToSessionIdMap gets a reference to the given map[string]string and assigns it to the PanelToSessionIdMap field.
func (o *DashboardTemplate) SetPanelToSessionIdMap(v map[string]string) {
	o.PanelToSessionIdMap = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *DashboardTemplate) GetTimeRange() ResolvableTimeRange {
	if o == nil || IsNil(o.TimeRange) {
		var ret ResolvableTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardTemplate) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *DashboardTemplate) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given ResolvableTimeRange and assigns it to the TimeRange field.
func (o *DashboardTemplate) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = &v
}

// GetVariableValues returns the VariableValues field value if set, zero value otherwise.
func (o *DashboardTemplate) GetVariableValues() VariablesValuesData {
	if o == nil || IsNil(o.VariableValues) {
		var ret VariablesValuesData
		return ret
	}
	return *o.VariableValues
}

// GetVariableValuesOk returns a tuple with the VariableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardTemplate) GetVariableValuesOk() (*VariablesValuesData, bool) {
	if o == nil || IsNil(o.VariableValues) {
		return nil, false
	}
	return o.VariableValues, true
}

// HasVariableValues returns a boolean if a field has been set.
func (o *DashboardTemplate) HasVariableValues() bool {
	if o != nil && !IsNil(o.VariableValues) {
		return true
	}

	return false
}

// SetVariableValues gets a reference to the given VariablesValuesData and assigns it to the VariableValues field.
func (o *DashboardTemplate) SetVariableValues(v VariablesValuesData) {
	o.VariableValues = &v
}

func (o DashboardTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTemplate, errTemplate := json.Marshal(o.Template)
	if errTemplate != nil {
		return map[string]interface{}{}, errTemplate
	}
	errTemplate = json.Unmarshal([]byte(serializedTemplate), &toSerialize)
	if errTemplate != nil {
		return map[string]interface{}{}, errTemplate
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.PanelToSessionIdMap) {
		toSerialize["panelToSessionIdMap"] = o.PanelToSessionIdMap
	}
	if !IsNil(o.TimeRange) {
		toSerialize["timeRange"] = o.TimeRange
	}
	if !IsNil(o.VariableValues) {
		toSerialize["variableValues"] = o.VariableValues
	}
	return toSerialize, nil
}

type NullableDashboardTemplate struct {
	value *DashboardTemplate
	isSet bool
}

func (v NullableDashboardTemplate) Get() *DashboardTemplate {
	return v.value
}

func (v *NullableDashboardTemplate) Set(val *DashboardTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardTemplate(val *DashboardTemplate) *NullableDashboardTemplate {
	return &NullableDashboardTemplate{value: val, isSet: true}
}

func (v NullableDashboardTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


