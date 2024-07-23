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

// checks if the MonitorTemplatesLibraryMonitorTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorTemplatesLibraryMonitorTemplate{}

// MonitorTemplatesLibraryMonitorTemplate struct for MonitorTemplatesLibraryMonitorTemplate
type MonitorTemplatesLibraryMonitorTemplate struct {
	MonitorTemplatesLibraryBase
	// The type of monitor template. Valid values:   1. `Logs`: A logs query monitor template.   2. `Metrics`: A metrics query monitor template.   3. `Slo`: A SLO based monitor template.
	MonitorType string `json:"monitorType" validate:"regexp=^(Logs|Metrics|Slo)$"`
	// The app id which related to the monitor template. It will be UUID string
	AppId string `json:"appId"`
	// The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time.
	EvaluationDelay *string `json:"evaluationDelay,omitempty"`
	// The name of the alert(s) triggered from the monitor created based on the template. Monitor name will be used if not specified.
	AlertName *string `json:"alertName,omitempty"`
	// All queries from the monitor.
	Queries []MonitorQuery `json:"queries"`
	// Defines the conditions of when to send notifications.
	Triggers []TriggerCondition `json:"triggers"`
	// Whether or not the monitor template is disabled.
	IsDisabled *bool `json:"isDisabled,omitempty"`
	// Whether or not to group notifications for individual items that meet the trigger condition.
	GroupNotifications *bool `json:"groupNotifications,omitempty"`
	// Notes such as links and instruction to help you resolve alerts triggered by this monitor template. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more.
	Playbook *string `json:"playbook,omitempty"`
}

type _MonitorTemplatesLibraryMonitorTemplate MonitorTemplatesLibraryMonitorTemplate

// NewMonitorTemplatesLibraryMonitorTemplate instantiates a new MonitorTemplatesLibraryMonitorTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorTemplatesLibraryMonitorTemplate(monitorType string, appId string, queries []MonitorQuery, triggers []TriggerCondition, name string, type_ string) *MonitorTemplatesLibraryMonitorTemplate {
	this := MonitorTemplatesLibraryMonitorTemplate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Type = type_
	this.MonitorType = monitorType
	this.AppId = appId
	var evaluationDelay string = "0m"
	this.EvaluationDelay = &evaluationDelay
	this.Queries = queries
	this.Triggers = triggers
	var isDisabled bool = false
	this.IsDisabled = &isDisabled
	var groupNotifications bool = true
	this.GroupNotifications = &groupNotifications
	var playbook string = ""
	this.Playbook = &playbook
	return &this
}

// NewMonitorTemplatesLibraryMonitorTemplateWithDefaults instantiates a new MonitorTemplatesLibraryMonitorTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorTemplatesLibraryMonitorTemplateWithDefaults() *MonitorTemplatesLibraryMonitorTemplate {
	this := MonitorTemplatesLibraryMonitorTemplate{}
	var evaluationDelay string = "0m"
	this.EvaluationDelay = &evaluationDelay
	var isDisabled bool = false
	this.IsDisabled = &isDisabled
	var groupNotifications bool = true
	this.GroupNotifications = &groupNotifications
	var playbook string = ""
	this.Playbook = &playbook
	return &this
}

// GetMonitorType returns the MonitorType field value
func (o *MonitorTemplatesLibraryMonitorTemplate) GetMonitorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetMonitorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorType, true
}

// SetMonitorType sets field value
func (o *MonitorTemplatesLibraryMonitorTemplate) SetMonitorType(v string) {
	o.MonitorType = v
}

// GetAppId returns the AppId field value
func (o *MonitorTemplatesLibraryMonitorTemplate) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MonitorTemplatesLibraryMonitorTemplate) SetAppId(v string) {
	o.AppId = v
}

// GetEvaluationDelay returns the EvaluationDelay field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetEvaluationDelay() string {
	if o == nil || IsNil(o.EvaluationDelay) {
		var ret string
		return ret
	}
	return *o.EvaluationDelay
}

// GetEvaluationDelayOk returns a tuple with the EvaluationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetEvaluationDelayOk() (*string, bool) {
	if o == nil || IsNil(o.EvaluationDelay) {
		return nil, false
	}
	return o.EvaluationDelay, true
}

// HasEvaluationDelay returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) HasEvaluationDelay() bool {
	if o != nil && !IsNil(o.EvaluationDelay) {
		return true
	}

	return false
}

// SetEvaluationDelay gets a reference to the given string and assigns it to the EvaluationDelay field.
func (o *MonitorTemplatesLibraryMonitorTemplate) SetEvaluationDelay(v string) {
	o.EvaluationDelay = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetAlertName() string {
	if o == nil || IsNil(o.AlertName) {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetAlertNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlertName) {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) HasAlertName() bool {
	if o != nil && !IsNil(o.AlertName) {
		return true
	}

	return false
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *MonitorTemplatesLibraryMonitorTemplate) SetAlertName(v string) {
	o.AlertName = &v
}

// GetQueries returns the Queries field value
func (o *MonitorTemplatesLibraryMonitorTemplate) GetQueries() []MonitorQuery {
	if o == nil {
		var ret []MonitorQuery
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetQueriesOk() ([]MonitorQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *MonitorTemplatesLibraryMonitorTemplate) SetQueries(v []MonitorQuery) {
	o.Queries = v
}

// GetTriggers returns the Triggers field value
func (o *MonitorTemplatesLibraryMonitorTemplate) GetTriggers() []TriggerCondition {
	if o == nil {
		var ret []TriggerCondition
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetTriggersOk() ([]TriggerCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers, true
}

// SetTriggers sets field value
func (o *MonitorTemplatesLibraryMonitorTemplate) SetTriggers(v []TriggerCondition) {
	o.Triggers = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *MonitorTemplatesLibraryMonitorTemplate) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetGroupNotifications returns the GroupNotifications field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetGroupNotifications() bool {
	if o == nil || IsNil(o.GroupNotifications) {
		var ret bool
		return ret
	}
	return *o.GroupNotifications
}

// GetGroupNotificationsOk returns a tuple with the GroupNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetGroupNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupNotifications) {
		return nil, false
	}
	return o.GroupNotifications, true
}

// HasGroupNotifications returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) HasGroupNotifications() bool {
	if o != nil && !IsNil(o.GroupNotifications) {
		return true
	}

	return false
}

// SetGroupNotifications gets a reference to the given bool and assigns it to the GroupNotifications field.
func (o *MonitorTemplatesLibraryMonitorTemplate) SetGroupNotifications(v bool) {
	o.GroupNotifications = &v
}

// GetPlaybook returns the Playbook field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetPlaybook() string {
	if o == nil || IsNil(o.Playbook) {
		var ret string
		return ret
	}
	return *o.Playbook
}

// GetPlaybookOk returns a tuple with the Playbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) GetPlaybookOk() (*string, bool) {
	if o == nil || IsNil(o.Playbook) {
		return nil, false
	}
	return o.Playbook, true
}

// HasPlaybook returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplate) HasPlaybook() bool {
	if o != nil && !IsNil(o.Playbook) {
		return true
	}

	return false
}

// SetPlaybook gets a reference to the given string and assigns it to the Playbook field.
func (o *MonitorTemplatesLibraryMonitorTemplate) SetPlaybook(v string) {
	o.Playbook = &v
}

func (o MonitorTemplatesLibraryMonitorTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorTemplatesLibraryMonitorTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMonitorTemplatesLibraryBase, errMonitorTemplatesLibraryBase := json.Marshal(o.MonitorTemplatesLibraryBase)
	if errMonitorTemplatesLibraryBase != nil {
		return map[string]interface{}{}, errMonitorTemplatesLibraryBase
	}
	errMonitorTemplatesLibraryBase = json.Unmarshal([]byte(serializedMonitorTemplatesLibraryBase), &toSerialize)
	if errMonitorTemplatesLibraryBase != nil {
		return map[string]interface{}{}, errMonitorTemplatesLibraryBase
	}
	toSerialize["monitorType"] = o.MonitorType
	toSerialize["appId"] = o.AppId
	if !IsNil(o.EvaluationDelay) {
		toSerialize["evaluationDelay"] = o.EvaluationDelay
	}
	if !IsNil(o.AlertName) {
		toSerialize["alertName"] = o.AlertName
	}
	toSerialize["queries"] = o.Queries
	toSerialize["triggers"] = o.Triggers
	if !IsNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if !IsNil(o.GroupNotifications) {
		toSerialize["groupNotifications"] = o.GroupNotifications
	}
	if !IsNil(o.Playbook) {
		toSerialize["playbook"] = o.Playbook
	}
	return toSerialize, nil
}

func (o *MonitorTemplatesLibraryMonitorTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"monitorType",
		"appId",
		"queries",
		"triggers",
		"name",
		"type",
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

	varMonitorTemplatesLibraryMonitorTemplate := _MonitorTemplatesLibraryMonitorTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMonitorTemplatesLibraryMonitorTemplate)

	if err != nil {
		return err
	}

	*o = MonitorTemplatesLibraryMonitorTemplate(varMonitorTemplatesLibraryMonitorTemplate)

	return err
}

type NullableMonitorTemplatesLibraryMonitorTemplate struct {
	value *MonitorTemplatesLibraryMonitorTemplate
	isSet bool
}

func (v NullableMonitorTemplatesLibraryMonitorTemplate) Get() *MonitorTemplatesLibraryMonitorTemplate {
	return v.value
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplate) Set(val *MonitorTemplatesLibraryMonitorTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorTemplatesLibraryMonitorTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorTemplatesLibraryMonitorTemplate(val *MonitorTemplatesLibraryMonitorTemplate) *NullableMonitorTemplatesLibraryMonitorTemplate {
	return &NullableMonitorTemplatesLibraryMonitorTemplate{value: val, isSet: true}
}

func (v NullableMonitorTemplatesLibraryMonitorTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


