/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// MonitorsLibraryMonitorAllOf struct for MonitorsLibraryMonitorAllOf
type MonitorsLibraryMonitorAllOf struct {
	// The type of monitor. Valid values:   1. `Logs`: A logs query monitor.   2. `Metrics`: A metrics query monitor.   3. `Slo`: A SLO based monitor. Currently SLO based monitor is available in closed beta (Notify your Sumo Logic representative in order to get the early access).
	MonitorType string `json:"monitorType"`
	// The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time.
	EvaluationDelay *string `json:"evaluationDelay,omitempty"`
	// The name of the alert(s) triggered from this monitor. Monitor name will be used if not specified.
	AlertName *string `json:"alertName,omitempty"`
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
}

// NewMonitorsLibraryMonitorAllOf instantiates a new MonitorsLibraryMonitorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorsLibraryMonitorAllOf(monitorType string, queries []MonitorQuery, triggers []TriggerCondition) *MonitorsLibraryMonitorAllOf {
	this := MonitorsLibraryMonitorAllOf{}
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

// NewMonitorsLibraryMonitorAllOfWithDefaults instantiates a new MonitorsLibraryMonitorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorsLibraryMonitorAllOfWithDefaults() *MonitorsLibraryMonitorAllOf {
	this := MonitorsLibraryMonitorAllOf{}
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
func (o *MonitorsLibraryMonitorAllOf) GetMonitorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetMonitorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorType, true
}

// SetMonitorType sets field value
func (o *MonitorsLibraryMonitorAllOf) SetMonitorType(v string) {
	o.MonitorType = v
}

// GetEvaluationDelay returns the EvaluationDelay field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorAllOf) GetEvaluationDelay() string {
	if o == nil || o.EvaluationDelay == nil {
		var ret string
		return ret
	}
	return *o.EvaluationDelay
}

// GetEvaluationDelayOk returns a tuple with the EvaluationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetEvaluationDelayOk() (*string, bool) {
	if o == nil || o.EvaluationDelay == nil {
		return nil, false
	}
	return o.EvaluationDelay, true
}

// HasEvaluationDelay returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorAllOf) HasEvaluationDelay() bool {
	if o != nil && o.EvaluationDelay != nil {
		return true
	}

	return false
}

// SetEvaluationDelay gets a reference to the given string and assigns it to the EvaluationDelay field.
func (o *MonitorsLibraryMonitorAllOf) SetEvaluationDelay(v string) {
	o.EvaluationDelay = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorAllOf) GetAlertName() string {
	if o == nil || o.AlertName == nil {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetAlertNameOk() (*string, bool) {
	if o == nil || o.AlertName == nil {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorAllOf) HasAlertName() bool {
	if o != nil && o.AlertName != nil {
		return true
	}

	return false
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *MonitorsLibraryMonitorAllOf) SetAlertName(v string) {
	o.AlertName = &v
}

// GetQueries returns the Queries field value
func (o *MonitorsLibraryMonitorAllOf) GetQueries() []MonitorQuery {
	if o == nil {
		var ret []MonitorQuery
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetQueriesOk() ([]MonitorQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *MonitorsLibraryMonitorAllOf) SetQueries(v []MonitorQuery) {
	o.Queries = v
}

// GetTriggers returns the Triggers field value
func (o *MonitorsLibraryMonitorAllOf) GetTriggers() []TriggerCondition {
	if o == nil {
		var ret []TriggerCondition
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetTriggersOk() ([]TriggerCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers, true
}

// SetTriggers sets field value
func (o *MonitorsLibraryMonitorAllOf) SetTriggers(v []TriggerCondition) {
	o.Triggers = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorAllOf) GetNotifications() []MonitorNotification {
	if o == nil || o.Notifications == nil {
		var ret []MonitorNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetNotificationsOk() ([]MonitorNotification, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorAllOf) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []MonitorNotification and assigns it to the Notifications field.
func (o *MonitorsLibraryMonitorAllOf) SetNotifications(v []MonitorNotification) {
	o.Notifications = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorAllOf) GetIsDisabled() bool {
	if o == nil || o.IsDisabled == nil {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetIsDisabledOk() (*bool, bool) {
	if o == nil || o.IsDisabled == nil {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorAllOf) HasIsDisabled() bool {
	if o != nil && o.IsDisabled != nil {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *MonitorsLibraryMonitorAllOf) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetGroupNotifications returns the GroupNotifications field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorAllOf) GetGroupNotifications() bool {
	if o == nil || o.GroupNotifications == nil {
		var ret bool
		return ret
	}
	return *o.GroupNotifications
}

// GetGroupNotificationsOk returns a tuple with the GroupNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetGroupNotificationsOk() (*bool, bool) {
	if o == nil || o.GroupNotifications == nil {
		return nil, false
	}
	return o.GroupNotifications, true
}

// HasGroupNotifications returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorAllOf) HasGroupNotifications() bool {
	if o != nil && o.GroupNotifications != nil {
		return true
	}

	return false
}

// SetGroupNotifications gets a reference to the given bool and assigns it to the GroupNotifications field.
func (o *MonitorsLibraryMonitorAllOf) SetGroupNotifications(v bool) {
	o.GroupNotifications = &v
}

// GetPlaybook returns the Playbook field value if set, zero value otherwise.
func (o *MonitorsLibraryMonitorAllOf) GetPlaybook() string {
	if o == nil || o.Playbook == nil {
		var ret string
		return ret
	}
	return *o.Playbook
}

// GetPlaybookOk returns a tuple with the Playbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorsLibraryMonitorAllOf) GetPlaybookOk() (*string, bool) {
	if o == nil || o.Playbook == nil {
		return nil, false
	}
	return o.Playbook, true
}

// HasPlaybook returns a boolean if a field has been set.
func (o *MonitorsLibraryMonitorAllOf) HasPlaybook() bool {
	if o != nil && o.Playbook != nil {
		return true
	}

	return false
}

// SetPlaybook gets a reference to the given string and assigns it to the Playbook field.
func (o *MonitorsLibraryMonitorAllOf) SetPlaybook(v string) {
	o.Playbook = &v
}

func (o MonitorsLibraryMonitorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["monitorType"] = o.MonitorType
	}
	if o.EvaluationDelay != nil {
		toSerialize["evaluationDelay"] = o.EvaluationDelay
	}
	if o.AlertName != nil {
		toSerialize["alertName"] = o.AlertName
	}
	if true {
		toSerialize["queries"] = o.Queries
	}
	if true {
		toSerialize["triggers"] = o.Triggers
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	if o.IsDisabled != nil {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if o.GroupNotifications != nil {
		toSerialize["groupNotifications"] = o.GroupNotifications
	}
	if o.Playbook != nil {
		toSerialize["playbook"] = o.Playbook
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorsLibraryMonitorAllOf struct {
	value *MonitorsLibraryMonitorAllOf
	isSet bool
}

func (v NullableMonitorsLibraryMonitorAllOf) Get() *MonitorsLibraryMonitorAllOf {
	return v.value
}

func (v *NullableMonitorsLibraryMonitorAllOf) Set(val *MonitorsLibraryMonitorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorsLibraryMonitorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorsLibraryMonitorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorsLibraryMonitorAllOf(val *MonitorsLibraryMonitorAllOf) *NullableMonitorsLibraryMonitorAllOf {
	return &NullableMonitorsLibraryMonitorAllOf{value: val, isSet: true}
}

func (v NullableMonitorsLibraryMonitorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorsLibraryMonitorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


