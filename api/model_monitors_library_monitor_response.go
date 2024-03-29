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

// checks if the MonitorsLibraryMonitorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorsLibraryMonitorResponse{}

// MonitorsLibraryMonitorResponse struct for MonitorsLibraryMonitorResponse
type MonitorsLibraryMonitorResponse struct {
	MonitorsLibraryBaseResponse
	// The type of monitor. Valid values:   1. `Logs`: A logs query monitor.   2. `Metrics`: A metrics query monitor.   3. `Slo`: A SLO based monitor. Currently SLO based monitor is available in closed beta (Notify your Sumo Logic representative in order to get the early access).
	MonitorType string `json:"monitorType"`
	// The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time.
	EvaluationDelay *string `json:"evaluationDelay,omitempty"`
	// The name of the alert(s) triggered from this monitor. Monitor name will be used if not specified.
	AlertName *string `json:"alertName,omitempty"`
	RunAs *MonitorsLibraryMonitorAllOfRunAs `json:"runAs,omitempty"`
	// The set of fields to be used to group alert notifications for a monitor. The value of this field will be considered only when 'groupNotifications' is true. The fields with very high cardinality such as `_raw`, `_messagetime`, `_receipttime`, and `_messageid` are not allowed for Alert Grouping.
	NotificationGroupFields []string `json:"notificationGroupFields,omitempty"`
	// All queries from the monitor.
	Queries []MonitorQuery `json:"queries"`
	// Defines the conditions of when to send notifications.
	Triggers []TriggerCondition `json:"triggers"`
	// The notifications the monitor will send when the respective trigger condition is met.
	Notifications []MonitorNotification `json:"notifications,omitempty"`
	// Whether or not the monitor is disabled. Disabled monitors will not run, and will not generate or send notifications.
	IsDisabled *bool `json:"isDisabled,omitempty"`
	// The current status of the monitor. Each monitor can have one or more status values. Valid values:   1. `Normal`: The monitor is running normally and does not have any currently triggered conditions.   2. `Critical`: The Critical trigger condition has been met.   3. `Warning`: The Warning trigger condition has been met.   4. `MissingData`: The MissingData trigger condition has been met.   5. `Disabled`: The monitor has been disabled and is not currently running.
	Status []string `json:"status,omitempty"`
	// Whether or not to group notifications for individual items that meet the trigger condition.
	GroupNotifications *bool `json:"groupNotifications,omitempty"`
	// Monitor manager warnings
	Warnings *map[string]string `json:"warnings,omitempty"`
	// Notes such as links and instruction to help you resolve alerts triggered by this monitor. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more.
	Playbook *string `json:"playbook,omitempty"`
	// Identifier of the SLO definition for the monitor. This is only applicable for SLO type monitors.
	SloId *string `json:"sloId,omitempty"`
	// The set of automated playbook ids for a monitor.
	AutomatedPlaybookIds []string `json:"automatedPlaybookIds,omitempty"`
}

// NewMonitorsLibraryMonitorResponse instantiates a new MonitorsLibraryMonitorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorsLibraryMonitorResponse(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool) *MonitorsLibraryMonitorResponse {
	this := MonitorsLibraryMonitorResponse{}
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

// NewMonitorsLibraryMonitorResponseWithDefaults instantiates a new MonitorsLibraryMonitorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorsLibraryMonitorResponseWithDefaults() *MonitorsLibraryMonitorResponse {
	this := MonitorsLibraryMonitorResponse{}
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
func (o *MonitorsLibraryMonitorResponse) GetMonitorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetMonitorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorType, true
}

// SetMonitorType sets field value
func (o *MonitorsLibraryMonitorResponse) SetMonitorType(v string) {
	o.MonitorType = v
}

// GetEvaluationDelay returns the EvaluationDelay field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetEvaluationDelay() string {
	if o == nil || IsNil(o.EvaluationDelay) {
		var ret string
		return ret
	}
	return *o.EvaluationDelay
}

// GetEvaluationDelayOk returns a tuple with the EvaluationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetEvaluationDelayOk() (*string, bool) {
	if o == nil || IsNil(o.EvaluationDelay) {
		return nil, false
	}
	return o.EvaluationDelay, true
}

// HasEvaluationDelay returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasEvaluationDelay() bool {
	if o != nil && !IsNil(o.EvaluationDelay) {
		return true
	}

	return false
}

// SetEvaluationDelay gets a reference to the given string and assigns it to the EvaluationDelay field.
func (o *MonitorsLibraryMonitorResponse) SetEvaluationDelay(v string) {
	o.EvaluationDelay = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetAlertName() string {
	if o == nil || IsNil(o.AlertName) {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetAlertNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlertName) {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasAlertName() bool {
	if o != nil && !IsNil(o.AlertName) {
		return true
	}

	return false
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *MonitorsLibraryMonitorResponse) SetAlertName(v string) {
	o.AlertName = &v
}

// GetRunAs returns the RunAs field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetRunAs() MonitorsLibraryMonitorAllOfRunAs {
	if o == nil || IsNil(o.RunAs) {
		var ret MonitorsLibraryMonitorAllOfRunAs
		return ret
	}
	return *o.RunAs
}

// GetRunAsOk returns a tuple with the RunAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetRunAsOk() (*MonitorsLibraryMonitorAllOfRunAs, bool) {
	if o == nil || IsNil(o.RunAs) {
		return nil, false
	}
	return o.RunAs, true
}

// HasRunAs returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasRunAs() bool {
	if o != nil && !IsNil(o.RunAs) {
		return true
	}

	return false
}

// SetRunAs gets a reference to the given MonitorsLibraryMonitorAllOfRunAs and assigns it to the RunAs field.
func (o *MonitorsLibraryMonitorResponse) SetRunAs(v MonitorsLibraryMonitorAllOfRunAs) {
	o.RunAs = &v
}

// GetNotificationGroupFields returns the NotificationGroupFields field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetNotificationGroupFields() []string {
	if o == nil || IsNil(o.NotificationGroupFields) {
		var ret []string
		return ret
	}
	return o.NotificationGroupFields
}

// GetNotificationGroupFieldsOk returns a tuple with the NotificationGroupFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetNotificationGroupFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.NotificationGroupFields) {
		return nil, false
	}
	return o.NotificationGroupFields, true
}

// HasNotificationGroupFields returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasNotificationGroupFields() bool {
	if o != nil && !IsNil(o.NotificationGroupFields) {
		return true
	}

	return false
}

// SetNotificationGroupFields gets a reference to the given []string and assigns it to the NotificationGroupFields field.
func (o *MonitorsLibraryMonitorResponse) SetNotificationGroupFields(v []string) {
	o.NotificationGroupFields = v
}

// GetQueries returns the Queries field value
func (o *MonitorsLibraryMonitorResponse) GetQueries() []MonitorQuery {
	if o == nil {
		var ret []MonitorQuery
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetQueriesOk() ([]MonitorQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *MonitorsLibraryMonitorResponse) SetQueries(v []MonitorQuery) {
	o.Queries = v
}

// GetTriggers returns the Triggers field value
func (o *MonitorsLibraryMonitorResponse) GetTriggers() []TriggerCondition {
	if o == nil {
		var ret []TriggerCondition
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetTriggersOk() ([]TriggerCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers, true
}

// SetTriggers sets field value
func (o *MonitorsLibraryMonitorResponse) SetTriggers(v []TriggerCondition) {
	o.Triggers = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetNotifications() []MonitorNotification {
	if o == nil || IsNil(o.Notifications) {
		var ret []MonitorNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetNotificationsOk() ([]MonitorNotification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []MonitorNotification and assigns it to the Notifications field.
func (o *MonitorsLibraryMonitorResponse) SetNotifications(v []MonitorNotification) {
	o.Notifications = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *MonitorsLibraryMonitorResponse) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetStatus() []string {
	if o == nil || IsNil(o.Status) {
		var ret []string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetStatusOk() ([]string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []string and assigns it to the Status field.
func (o *MonitorsLibraryMonitorResponse) SetStatus(v []string) {
	o.Status = v
}

// GetGroupNotifications returns the GroupNotifications field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetGroupNotifications() bool {
	if o == nil || IsNil(o.GroupNotifications) {
		var ret bool
		return ret
	}
	return *o.GroupNotifications
}

// GetGroupNotificationsOk returns a tuple with the GroupNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetGroupNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupNotifications) {
		return nil, false
	}
	return o.GroupNotifications, true
}

// HasGroupNotifications returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasGroupNotifications() bool {
	if o != nil && !IsNil(o.GroupNotifications) {
		return true
	}

	return false
}

// SetGroupNotifications gets a reference to the given bool and assigns it to the GroupNotifications field.
func (o *MonitorsLibraryMonitorResponse) SetGroupNotifications(v bool) {
	o.GroupNotifications = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetWarnings() map[string]string {
	if o == nil || IsNil(o.Warnings) {
		var ret map[string]string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetWarningsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given map[string]string and assigns it to the Warnings field.
func (o *MonitorsLibraryMonitorResponse) SetWarnings(v map[string]string) {
	o.Warnings = &v
}

// GetPlaybook returns the Playbook field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetPlaybook() string {
	if o == nil || IsNil(o.Playbook) {
		var ret string
		return ret
	}
	return *o.Playbook
}

// GetPlaybookOk returns a tuple with the Playbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetPlaybookOk() (*string, bool) {
	if o == nil || IsNil(o.Playbook) {
		return nil, false
	}
	return o.Playbook, true
}

// HasPlaybook returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasPlaybook() bool {
	if o != nil && !IsNil(o.Playbook) {
		return true
	}

	return false
}

// SetPlaybook gets a reference to the given string and assigns it to the Playbook field.
func (o *MonitorsLibraryMonitorResponse) SetPlaybook(v string) {
	o.Playbook = &v
}

// GetSloId returns the SloId field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetSloId() string {
	if o == nil || IsNil(o.SloId) {
		var ret string
		return ret
	}
	return *o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetSloIdOk() (*string, bool) {
	if o == nil || IsNil(o.SloId) {
		return nil, false
	}
	return o.SloId, true
}

// HasSloId returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasSloId() bool {
	if o != nil && !IsNil(o.SloId) {
		return true
	}

	return false
}

// SetSloId gets a reference to the given string and assigns it to the SloId field.
func (o *MonitorsLibraryMonitorResponse) SetSloId(v string) {
	o.SloId = &v
}

// GetAutomatedPlaybookIds returns the AutomatedPlaybookIds field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorResponse) GetAutomatedPlaybookIds() []string {
	if o == nil || IsNil(o.AutomatedPlaybookIds) {
		var ret []string
		return ret
	}
	return o.AutomatedPlaybookIds
}

// GetAutomatedPlaybookIdsOk returns a tuple with the AutomatedPlaybookIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorResponse) GetAutomatedPlaybookIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutomatedPlaybookIds) {
		return nil, false
	}
	return o.AutomatedPlaybookIds, true
}

// HasAutomatedPlaybookIds returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorResponse) HasAutomatedPlaybookIds() bool {
	if o != nil && !IsNil(o.AutomatedPlaybookIds) {
		return true
	}

	return false
}

// SetAutomatedPlaybookIds gets a reference to the given []string and assigns it to the AutomatedPlaybookIds field.
func (o *MonitorsLibraryMonitorResponse) SetAutomatedPlaybookIds(v []string) {
	o.AutomatedPlaybookIds = v
}

func (o MonitorsLibraryMonitorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorsLibraryMonitorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMonitorsLibraryBaseResponse, errMonitorsLibraryBaseResponse := json.Marshal(o.MonitorsLibraryBaseResponse)
	if errMonitorsLibraryBaseResponse != nil {
		return map[string]interface{}{}, errMonitorsLibraryBaseResponse
	}
	errMonitorsLibraryBaseResponse = json.Unmarshal([]byte(serializedMonitorsLibraryBaseResponse), &toSerialize)
	if errMonitorsLibraryBaseResponse != nil {
		return map[string]interface{}{}, errMonitorsLibraryBaseResponse
	}
	toSerialize["monitorType"] = o.MonitorType
	if !IsNil(o.EvaluationDelay) {
		toSerialize["evaluationDelay"] = o.EvaluationDelay
	}
	if !IsNil(o.AlertName) {
		toSerialize["alertName"] = o.AlertName
	}
	if !IsNil(o.RunAs) {
		toSerialize["runAs"] = o.RunAs
	}
	if !IsNil(o.NotificationGroupFields) {
		toSerialize["notificationGroupFields"] = o.NotificationGroupFields
	}
	toSerialize["queries"] = o.Queries
	toSerialize["triggers"] = o.Triggers
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.GroupNotifications) {
		toSerialize["groupNotifications"] = o.GroupNotifications
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Playbook) {
		toSerialize["playbook"] = o.Playbook
	}
	if !IsNil(o.SloId) {
		toSerialize["sloId"] = o.SloId
	}
	if !IsNil(o.AutomatedPlaybookIds) {
		toSerialize["automatedPlaybookIds"] = o.AutomatedPlaybookIds
	}
	return toSerialize, nil
}

type NullableMonitorsLibraryMonitorResponse struct {
	value *MonitorsLibraryMonitorResponse
	isSet bool
}

func (v NullableMonitorsLibraryMonitorResponse) Get() *MonitorsLibraryMonitorResponse {
	return v.value
}

func (v *NullableMonitorsLibraryMonitorResponse) Set(val *MonitorsLibraryMonitorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorsLibraryMonitorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorsLibraryMonitorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorsLibraryMonitorResponse(val *MonitorsLibraryMonitorResponse) *NullableMonitorsLibraryMonitorResponse {
	return &NullableMonitorsLibraryMonitorResponse{value: val, isSet: true}
}

func (v NullableMonitorsLibraryMonitorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorsLibraryMonitorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


