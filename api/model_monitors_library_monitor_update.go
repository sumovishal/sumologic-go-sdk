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

// checks if the MonitorsLibraryMonitorUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorsLibraryMonitorUpdate{}

// MonitorsLibraryMonitorUpdate struct for MonitorsLibraryMonitorUpdate
type MonitorsLibraryMonitorUpdate struct {
	MonitorsLibraryBaseUpdate
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
	// Whether or not to group notifications for individual items that meet the trigger condition.
	GroupNotifications *bool `json:"groupNotifications,omitempty"`
	// Notes such as links and instruction to help you resolve alerts triggered by this monitor. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more.
	Playbook *string `json:"playbook,omitempty"`
	// Identifier of the SLO definition for the monitor. This is only applicable for SLO type monitors.
	SloId *string `json:"sloId,omitempty"`
	// The set of automated playbook ids for a monitor.
	AutomatedPlaybookIds []string `json:"automatedPlaybookIds,omitempty"`
}

// NewMonitorsLibraryMonitorUpdate instantiates a new MonitorsLibraryMonitorUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorsLibraryMonitorUpdate(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, name string, version int64, type_ string) *MonitorsLibraryMonitorUpdate {
	this := MonitorsLibraryMonitorUpdate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Version = version
	this.Type = type_
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

// NewMonitorsLibraryMonitorUpdateWithDefaults instantiates a new MonitorsLibraryMonitorUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorsLibraryMonitorUpdateWithDefaults() *MonitorsLibraryMonitorUpdate {
	this := MonitorsLibraryMonitorUpdate{}
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
func (o *MonitorsLibraryMonitorUpdate) GetMonitorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetMonitorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorType, true
}

// SetMonitorType sets field value
func (o *MonitorsLibraryMonitorUpdate) SetMonitorType(v string) {
	o.MonitorType = v
}

// GetEvaluationDelay returns the EvaluationDelay field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetEvaluationDelay() string {
	if o == nil || IsNil(o.EvaluationDelay) {
		var ret string
		return ret
	}
	return *o.EvaluationDelay
}

// GetEvaluationDelayOk returns a tuple with the EvaluationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetEvaluationDelayOk() (*string, bool) {
	if o == nil || IsNil(o.EvaluationDelay) {
		return nil, false
	}
	return o.EvaluationDelay, true
}

// HasEvaluationDelay returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasEvaluationDelay() bool {
	if o != nil && !IsNil(o.EvaluationDelay) {
		return true
	}

	return false
}

// SetEvaluationDelay gets a reference to the given string and assigns it to the EvaluationDelay field.
func (o *MonitorsLibraryMonitorUpdate) SetEvaluationDelay(v string) {
	o.EvaluationDelay = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetAlertName() string {
	if o == nil || IsNil(o.AlertName) {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetAlertNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlertName) {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasAlertName() bool {
	if o != nil && !IsNil(o.AlertName) {
		return true
	}

	return false
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *MonitorsLibraryMonitorUpdate) SetAlertName(v string) {
	o.AlertName = &v
}

// GetRunAs returns the RunAs field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetRunAs() MonitorsLibraryMonitorAllOfRunAs {
	if o == nil || IsNil(o.RunAs) {
		var ret MonitorsLibraryMonitorAllOfRunAs
		return ret
	}
	return *o.RunAs
}

// GetRunAsOk returns a tuple with the RunAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetRunAsOk() (*MonitorsLibraryMonitorAllOfRunAs, bool) {
	if o == nil || IsNil(o.RunAs) {
		return nil, false
	}
	return o.RunAs, true
}

// HasRunAs returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasRunAs() bool {
	if o != nil && !IsNil(o.RunAs) {
		return true
	}

	return false
}

// SetRunAs gets a reference to the given MonitorsLibraryMonitorAllOfRunAs and assigns it to the RunAs field.
func (o *MonitorsLibraryMonitorUpdate) SetRunAs(v MonitorsLibraryMonitorAllOfRunAs) {
	o.RunAs = &v
}

// GetNotificationGroupFields returns the NotificationGroupFields field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetNotificationGroupFields() []string {
	if o == nil || IsNil(o.NotificationGroupFields) {
		var ret []string
		return ret
	}
	return o.NotificationGroupFields
}

// GetNotificationGroupFieldsOk returns a tuple with the NotificationGroupFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetNotificationGroupFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.NotificationGroupFields) {
		return nil, false
	}
	return o.NotificationGroupFields, true
}

// HasNotificationGroupFields returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasNotificationGroupFields() bool {
	if o != nil && !IsNil(o.NotificationGroupFields) {
		return true
	}

	return false
}

// SetNotificationGroupFields gets a reference to the given []string and assigns it to the NotificationGroupFields field.
func (o *MonitorsLibraryMonitorUpdate) SetNotificationGroupFields(v []string) {
	o.NotificationGroupFields = v
}

// GetQueries returns the Queries field value
func (o *MonitorsLibraryMonitorUpdate) GetQueries() []MonitorQuery {
	if o == nil {
		var ret []MonitorQuery
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetQueriesOk() ([]MonitorQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *MonitorsLibraryMonitorUpdate) SetQueries(v []MonitorQuery) {
	o.Queries = v
}

// GetTriggers returns the Triggers field value
func (o *MonitorsLibraryMonitorUpdate) GetTriggers() []TriggerCondition {
	if o == nil {
		var ret []TriggerCondition
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetTriggersOk() ([]TriggerCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers, true
}

// SetTriggers sets field value
func (o *MonitorsLibraryMonitorUpdate) SetTriggers(v []TriggerCondition) {
	o.Triggers = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetNotifications() []MonitorNotification {
	if o == nil || IsNil(o.Notifications) {
		var ret []MonitorNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetNotificationsOk() ([]MonitorNotification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []MonitorNotification and assigns it to the Notifications field.
func (o *MonitorsLibraryMonitorUpdate) SetNotifications(v []MonitorNotification) {
	o.Notifications = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *MonitorsLibraryMonitorUpdate) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetGroupNotifications returns the GroupNotifications field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetGroupNotifications() bool {
	if o == nil || IsNil(o.GroupNotifications) {
		var ret bool
		return ret
	}
	return *o.GroupNotifications
}

// GetGroupNotificationsOk returns a tuple with the GroupNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetGroupNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupNotifications) {
		return nil, false
	}
	return o.GroupNotifications, true
}

// HasGroupNotifications returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasGroupNotifications() bool {
	if o != nil && !IsNil(o.GroupNotifications) {
		return true
	}

	return false
}

// SetGroupNotifications gets a reference to the given bool and assigns it to the GroupNotifications field.
func (o *MonitorsLibraryMonitorUpdate) SetGroupNotifications(v bool) {
	o.GroupNotifications = &v
}

// GetPlaybook returns the Playbook field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetPlaybook() string {
	if o == nil || IsNil(o.Playbook) {
		var ret string
		return ret
	}
	return *o.Playbook
}

// GetPlaybookOk returns a tuple with the Playbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetPlaybookOk() (*string, bool) {
	if o == nil || IsNil(o.Playbook) {
		return nil, false
	}
	return o.Playbook, true
}

// HasPlaybook returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasPlaybook() bool {
	if o != nil && !IsNil(o.Playbook) {
		return true
	}

	return false
}

// SetPlaybook gets a reference to the given string and assigns it to the Playbook field.
func (o *MonitorsLibraryMonitorUpdate) SetPlaybook(v string) {
	o.Playbook = &v
}

// GetSloId returns the SloId field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetSloId() string {
	if o == nil || IsNil(o.SloId) {
		var ret string
		return ret
	}
	return *o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetSloIdOk() (*string, bool) {
	if o == nil || IsNil(o.SloId) {
		return nil, false
	}
	return o.SloId, true
}

// HasSloId returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasSloId() bool {
	if o != nil && !IsNil(o.SloId) {
		return true
	}

	return false
}

// SetSloId gets a reference to the given string and assigns it to the SloId field.
func (o *MonitorsLibraryMonitorUpdate) SetSloId(v string) {
	o.SloId = &v
}

// GetAutomatedPlaybookIds returns the AutomatedPlaybookIds field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorUpdate) GetAutomatedPlaybookIds() []string {
	if o == nil || IsNil(o.AutomatedPlaybookIds) {
		var ret []string
		return ret
	}
	return o.AutomatedPlaybookIds
}

// GetAutomatedPlaybookIdsOk returns a tuple with the AutomatedPlaybookIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorUpdate) GetAutomatedPlaybookIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutomatedPlaybookIds) {
		return nil, false
	}
	return o.AutomatedPlaybookIds, true
}

// HasAutomatedPlaybookIds returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorUpdate) HasAutomatedPlaybookIds() bool {
	if o != nil && !IsNil(o.AutomatedPlaybookIds) {
		return true
	}

	return false
}

// SetAutomatedPlaybookIds gets a reference to the given []string and assigns it to the AutomatedPlaybookIds field.
func (o *MonitorsLibraryMonitorUpdate) SetAutomatedPlaybookIds(v []string) {
	o.AutomatedPlaybookIds = v
}

func (o MonitorsLibraryMonitorUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorsLibraryMonitorUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMonitorsLibraryBaseUpdate, errMonitorsLibraryBaseUpdate := json.Marshal(o.MonitorsLibraryBaseUpdate)
	if errMonitorsLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errMonitorsLibraryBaseUpdate
	}
	errMonitorsLibraryBaseUpdate = json.Unmarshal([]byte(serializedMonitorsLibraryBaseUpdate), &toSerialize)
	if errMonitorsLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errMonitorsLibraryBaseUpdate
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
	if !IsNil(o.GroupNotifications) {
		toSerialize["groupNotifications"] = o.GroupNotifications
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

type NullableMonitorsLibraryMonitorUpdate struct {
	value *MonitorsLibraryMonitorUpdate
	isSet bool
}

func (v NullableMonitorsLibraryMonitorUpdate) Get() *MonitorsLibraryMonitorUpdate {
	return v.value
}

func (v *NullableMonitorsLibraryMonitorUpdate) Set(val *MonitorsLibraryMonitorUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorsLibraryMonitorUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorsLibraryMonitorUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorsLibraryMonitorUpdate(val *MonitorsLibraryMonitorUpdate) *NullableMonitorsLibraryMonitorUpdate {
	return &NullableMonitorsLibraryMonitorUpdate{value: val, isSet: true}
}

func (v NullableMonitorsLibraryMonitorUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorsLibraryMonitorUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


