/*
Sumo Logic API

Go client for Sumo Logic API.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// checks if the MonitorTemplatesLibraryMonitorTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorTemplatesLibraryMonitorTemplateResponse{}

// MonitorTemplatesLibraryMonitorTemplateResponse struct for MonitorTemplatesLibraryMonitorTemplateResponse
type MonitorTemplatesLibraryMonitorTemplateResponse struct {
	MonitorTemplatesLibraryBaseResponse
	// The type of monitor template. Valid values:   1. `Logs`: A logs query monitor template.   2. `Metrics`: A metrics query monitor template.   3. `Slo`: A SLO based monitor template.
	MonitorType string `json:"monitorType"`
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

// NewMonitorTemplatesLibraryMonitorTemplateResponse instantiates a new MonitorTemplatesLibraryMonitorTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorTemplatesLibraryMonitorTemplateResponse(monitorType string, appId string, queries []MonitorQuery, triggers []TriggerCondition, id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool) *MonitorTemplatesLibraryMonitorTemplateResponse {
	this := MonitorTemplatesLibraryMonitorTemplateResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Version = version
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.ParentId = parentId
	this.ContentType = contentType
	this.Type = type_
	this.IsSystem = isSystem
	this.IsMutable = isMutable
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

// NewMonitorTemplatesLibraryMonitorTemplateResponseWithDefaults instantiates a new MonitorTemplatesLibraryMonitorTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorTemplatesLibraryMonitorTemplateResponseWithDefaults() *MonitorTemplatesLibraryMonitorTemplateResponse {
	this := MonitorTemplatesLibraryMonitorTemplateResponse{}
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
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetMonitorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetMonitorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorType, true
}

// SetMonitorType sets field value
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetMonitorType(v string) {
	o.MonitorType = v
}

// GetAppId returns the AppId field value
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetAppId(v string) {
	o.AppId = v
}

// GetEvaluationDelay returns the EvaluationDelay field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetEvaluationDelay() string {
	if o == nil || IsNil(o.EvaluationDelay) {
		var ret string
		return ret
	}
	return *o.EvaluationDelay
}

// GetEvaluationDelayOk returns a tuple with the EvaluationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetEvaluationDelayOk() (*string, bool) {
	if o == nil || IsNil(o.EvaluationDelay) {
		return nil, false
	}
	return o.EvaluationDelay, true
}

// HasEvaluationDelay returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasEvaluationDelay() bool {
	if o != nil && !IsNil(o.EvaluationDelay) {
		return true
	}

	return false
}

// SetEvaluationDelay gets a reference to the given string and assigns it to the EvaluationDelay field.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetEvaluationDelay(v string) {
	o.EvaluationDelay = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetAlertName() string {
	if o == nil || IsNil(o.AlertName) {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetAlertNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlertName) {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasAlertName() bool {
	if o != nil && !IsNil(o.AlertName) {
		return true
	}

	return false
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetAlertName(v string) {
	o.AlertName = &v
}

// GetQueries returns the Queries field value
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetQueries() []MonitorQuery {
	if o == nil {
		var ret []MonitorQuery
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetQueriesOk() ([]MonitorQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetQueries(v []MonitorQuery) {
	o.Queries = v
}

// GetTriggers returns the Triggers field value
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetTriggers() []TriggerCondition {
	if o == nil {
		var ret []TriggerCondition
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetTriggersOk() ([]TriggerCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers, true
}

// SetTriggers sets field value
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetTriggers(v []TriggerCondition) {
	o.Triggers = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetGroupNotifications returns the GroupNotifications field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetGroupNotifications() bool {
	if o == nil || IsNil(o.GroupNotifications) {
		var ret bool
		return ret
	}
	return *o.GroupNotifications
}

// GetGroupNotificationsOk returns a tuple with the GroupNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetGroupNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupNotifications) {
		return nil, false
	}
	return o.GroupNotifications, true
}

// HasGroupNotifications returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasGroupNotifications() bool {
	if o != nil && !IsNil(o.GroupNotifications) {
		return true
	}

	return false
}

// SetGroupNotifications gets a reference to the given bool and assigns it to the GroupNotifications field.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetGroupNotifications(v bool) {
	o.GroupNotifications = &v
}

// GetPlaybook returns the Playbook field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetPlaybook() string {
	if o == nil || IsNil(o.Playbook) {
		var ret string
		return ret
	}
	return *o.Playbook
}

// GetPlaybookOk returns a tuple with the Playbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetPlaybookOk() (*string, bool) {
	if o == nil || IsNil(o.Playbook) {
		return nil, false
	}
	return o.Playbook, true
}

// HasPlaybook returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasPlaybook() bool {
	if o != nil && !IsNil(o.Playbook) {
		return true
	}

	return false
}

// SetPlaybook gets a reference to the given string and assigns it to the Playbook field.
func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetPlaybook(v string) {
	o.Playbook = &v
}

func (o MonitorTemplatesLibraryMonitorTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorTemplatesLibraryMonitorTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMonitorTemplatesLibraryBaseResponse, errMonitorTemplatesLibraryBaseResponse := json.Marshal(o.MonitorTemplatesLibraryBaseResponse)
	if errMonitorTemplatesLibraryBaseResponse != nil {
		return map[string]interface{}{}, errMonitorTemplatesLibraryBaseResponse
	}
	errMonitorTemplatesLibraryBaseResponse = json.Unmarshal([]byte(serializedMonitorTemplatesLibraryBaseResponse), &toSerialize)
	if errMonitorTemplatesLibraryBaseResponse != nil {
		return map[string]interface{}{}, errMonitorTemplatesLibraryBaseResponse
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

type NullableMonitorTemplatesLibraryMonitorTemplateResponse struct {
	value *MonitorTemplatesLibraryMonitorTemplateResponse
	isSet bool
}

func (v NullableMonitorTemplatesLibraryMonitorTemplateResponse) Get() *MonitorTemplatesLibraryMonitorTemplateResponse {
	return v.value
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplateResponse) Set(val *MonitorTemplatesLibraryMonitorTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorTemplatesLibraryMonitorTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorTemplatesLibraryMonitorTemplateResponse(val *MonitorTemplatesLibraryMonitorTemplateResponse) *NullableMonitorTemplatesLibraryMonitorTemplateResponse {
	return &NullableMonitorTemplatesLibraryMonitorTemplateResponse{value: val, isSet: true}
}

func (v NullableMonitorTemplatesLibraryMonitorTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


