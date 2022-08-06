/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// AlertsLibraryAlertAllOf struct for AlertsLibraryAlertAllOf
type AlertsLibraryAlertAllOf struct {
	// The Id of the associated monitor.
	MonitorId *string `json:"monitorId,omitempty"`
	// The time at which the alert was resolved.
	ResolvedAt NullableTime `json:"resolvedAt,omitempty"`
	// The time at which the incident started.
	AbnormalityStartTime *time.Time `json:"abnormalityStartTime,omitempty"`
	// The severity of the Alert. Valid values:   1. `Critical`   2. `Warning`   3. `MissingData`
	AlertType *string `json:"alertType,omitempty"`
	// The status of the Alert. Valid values:   1. `Triggered`   2. `Resolved`
	Status *string `json:"status,omitempty"`
	// All queries from the monitor relevant to the alert.
	MonitorQueries []AlertMonitorQuery `json:"monitorQueries,omitempty"`
	// All queries from the monitor relevant to the alert with triggered time series filters.
	TriggerQueries []AlertMonitorQuery `json:"triggerQueries,omitempty"`
	// URL for this monitor's view page
	MonitorUrl *string `json:"monitorUrl,omitempty"`
	// A link to search with the triggering data and time range
	TriggerQueryUrl *string `json:"triggerQueryUrl,omitempty"`
	// Trigger conditions which were breached to create this Alert.
	TriggerConditions []TriggerCondition `json:"triggerConditions,omitempty"`
	// The of the query result which breached the trigger condition.
	TriggerValue *float64 `json:"triggerValue,omitempty"`
	// The type of monitor. Valid values:   1. `Logs`: A logs query monitor.   2. `Metrics`: A metrics query monitor.
	MonitorType *string `json:"monitorType,omitempty"`
	// One or more entity identifiers involved in this Alert.
	EntityIds []string `json:"entityIds,omitempty"`
	// One or more entity involved in this Alert.
	Entities []AlertEntityInfo `json:"entities,omitempty"`
	Notes *string `json:"notes,omitempty"`
	ExtraDetails *ExtraDetails `json:"extraDetails,omitempty"`
	// The condition which triggered this alert.
	AlertCondition NullableString `json:"alertCondition,omitempty"`
}

// NewAlertsLibraryAlertAllOf instantiates a new AlertsLibraryAlertAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryAlertAllOf() *AlertsLibraryAlertAllOf {
	this := AlertsLibraryAlertAllOf{}
	return &this
}

// NewAlertsLibraryAlertAllOfWithDefaults instantiates a new AlertsLibraryAlertAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryAlertAllOfWithDefaults() *AlertsLibraryAlertAllOf {
	this := AlertsLibraryAlertAllOf{}
	return &this
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetMonitorId() string {
	if o == nil || o.MonitorId == nil {
		var ret string
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetMonitorIdOk() (*string, bool) {
	if o == nil || o.MonitorId == nil {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasMonitorId() bool {
	if o != nil && o.MonitorId != nil {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given string and assigns it to the MonitorId field.
func (o *AlertsLibraryAlertAllOf) SetMonitorId(v string) {
	o.MonitorId = &v
}

// GetResolvedAt returns the ResolvedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsLibraryAlertAllOf) GetResolvedAt() time.Time {
	if o == nil || o.ResolvedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ResolvedAt.Get()
}

// GetResolvedAtOk returns a tuple with the ResolvedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertsLibraryAlertAllOf) GetResolvedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedAt.Get(), o.ResolvedAt.IsSet()
}

// HasResolvedAt returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasResolvedAt() bool {
	if o != nil && o.ResolvedAt.IsSet() {
		return true
	}

	return false
}

// SetResolvedAt gets a reference to the given NullableTime and assigns it to the ResolvedAt field.
func (o *AlertsLibraryAlertAllOf) SetResolvedAt(v time.Time) {
	o.ResolvedAt.Set(&v)
}
// SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil
func (o *AlertsLibraryAlertAllOf) SetResolvedAtNil() {
	o.ResolvedAt.Set(nil)
}

// UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
func (o *AlertsLibraryAlertAllOf) UnsetResolvedAt() {
	o.ResolvedAt.Unset()
}

// GetAbnormalityStartTime returns the AbnormalityStartTime field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetAbnormalityStartTime() time.Time {
	if o == nil || o.AbnormalityStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AbnormalityStartTime
}

// GetAbnormalityStartTimeOk returns a tuple with the AbnormalityStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetAbnormalityStartTimeOk() (*time.Time, bool) {
	if o == nil || o.AbnormalityStartTime == nil {
		return nil, false
	}
	return o.AbnormalityStartTime, true
}

// HasAbnormalityStartTime returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasAbnormalityStartTime() bool {
	if o != nil && o.AbnormalityStartTime != nil {
		return true
	}

	return false
}

// SetAbnormalityStartTime gets a reference to the given time.Time and assigns it to the AbnormalityStartTime field.
func (o *AlertsLibraryAlertAllOf) SetAbnormalityStartTime(v time.Time) {
	o.AbnormalityStartTime = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetAlertType() string {
	if o == nil || o.AlertType == nil {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetAlertTypeOk() (*string, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasAlertType() bool {
	if o != nil && o.AlertType != nil {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *AlertsLibraryAlertAllOf) SetAlertType(v string) {
	o.AlertType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlertsLibraryAlertAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetMonitorQueries returns the MonitorQueries field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetMonitorQueries() []AlertMonitorQuery {
	if o == nil || o.MonitorQueries == nil {
		var ret []AlertMonitorQuery
		return ret
	}
	return o.MonitorQueries
}

// GetMonitorQueriesOk returns a tuple with the MonitorQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetMonitorQueriesOk() ([]AlertMonitorQuery, bool) {
	if o == nil || o.MonitorQueries == nil {
		return nil, false
	}
	return o.MonitorQueries, true
}

// HasMonitorQueries returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasMonitorQueries() bool {
	if o != nil && o.MonitorQueries != nil {
		return true
	}

	return false
}

// SetMonitorQueries gets a reference to the given []AlertMonitorQuery and assigns it to the MonitorQueries field.
func (o *AlertsLibraryAlertAllOf) SetMonitorQueries(v []AlertMonitorQuery) {
	o.MonitorQueries = v
}

// GetTriggerQueries returns the TriggerQueries field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetTriggerQueries() []AlertMonitorQuery {
	if o == nil || o.TriggerQueries == nil {
		var ret []AlertMonitorQuery
		return ret
	}
	return o.TriggerQueries
}

// GetTriggerQueriesOk returns a tuple with the TriggerQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetTriggerQueriesOk() ([]AlertMonitorQuery, bool) {
	if o == nil || o.TriggerQueries == nil {
		return nil, false
	}
	return o.TriggerQueries, true
}

// HasTriggerQueries returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasTriggerQueries() bool {
	if o != nil && o.TriggerQueries != nil {
		return true
	}

	return false
}

// SetTriggerQueries gets a reference to the given []AlertMonitorQuery and assigns it to the TriggerQueries field.
func (o *AlertsLibraryAlertAllOf) SetTriggerQueries(v []AlertMonitorQuery) {
	o.TriggerQueries = v
}

// GetMonitorUrl returns the MonitorUrl field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetMonitorUrl() string {
	if o == nil || o.MonitorUrl == nil {
		var ret string
		return ret
	}
	return *o.MonitorUrl
}

// GetMonitorUrlOk returns a tuple with the MonitorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetMonitorUrlOk() (*string, bool) {
	if o == nil || o.MonitorUrl == nil {
		return nil, false
	}
	return o.MonitorUrl, true
}

// HasMonitorUrl returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasMonitorUrl() bool {
	if o != nil && o.MonitorUrl != nil {
		return true
	}

	return false
}

// SetMonitorUrl gets a reference to the given string and assigns it to the MonitorUrl field.
func (o *AlertsLibraryAlertAllOf) SetMonitorUrl(v string) {
	o.MonitorUrl = &v
}

// GetTriggerQueryUrl returns the TriggerQueryUrl field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetTriggerQueryUrl() string {
	if o == nil || o.TriggerQueryUrl == nil {
		var ret string
		return ret
	}
	return *o.TriggerQueryUrl
}

// GetTriggerQueryUrlOk returns a tuple with the TriggerQueryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetTriggerQueryUrlOk() (*string, bool) {
	if o == nil || o.TriggerQueryUrl == nil {
		return nil, false
	}
	return o.TriggerQueryUrl, true
}

// HasTriggerQueryUrl returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasTriggerQueryUrl() bool {
	if o != nil && o.TriggerQueryUrl != nil {
		return true
	}

	return false
}

// SetTriggerQueryUrl gets a reference to the given string and assigns it to the TriggerQueryUrl field.
func (o *AlertsLibraryAlertAllOf) SetTriggerQueryUrl(v string) {
	o.TriggerQueryUrl = &v
}

// GetTriggerConditions returns the TriggerConditions field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetTriggerConditions() []TriggerCondition {
	if o == nil || o.TriggerConditions == nil {
		var ret []TriggerCondition
		return ret
	}
	return o.TriggerConditions
}

// GetTriggerConditionsOk returns a tuple with the TriggerConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetTriggerConditionsOk() ([]TriggerCondition, bool) {
	if o == nil || o.TriggerConditions == nil {
		return nil, false
	}
	return o.TriggerConditions, true
}

// HasTriggerConditions returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasTriggerConditions() bool {
	if o != nil && o.TriggerConditions != nil {
		return true
	}

	return false
}

// SetTriggerConditions gets a reference to the given []TriggerCondition and assigns it to the TriggerConditions field.
func (o *AlertsLibraryAlertAllOf) SetTriggerConditions(v []TriggerCondition) {
	o.TriggerConditions = v
}

// GetTriggerValue returns the TriggerValue field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetTriggerValue() float64 {
	if o == nil || o.TriggerValue == nil {
		var ret float64
		return ret
	}
	return *o.TriggerValue
}

// GetTriggerValueOk returns a tuple with the TriggerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetTriggerValueOk() (*float64, bool) {
	if o == nil || o.TriggerValue == nil {
		return nil, false
	}
	return o.TriggerValue, true
}

// HasTriggerValue returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasTriggerValue() bool {
	if o != nil && o.TriggerValue != nil {
		return true
	}

	return false
}

// SetTriggerValue gets a reference to the given float64 and assigns it to the TriggerValue field.
func (o *AlertsLibraryAlertAllOf) SetTriggerValue(v float64) {
	o.TriggerValue = &v
}

// GetMonitorType returns the MonitorType field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetMonitorType() string {
	if o == nil || o.MonitorType == nil {
		var ret string
		return ret
	}
	return *o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetMonitorTypeOk() (*string, bool) {
	if o == nil || o.MonitorType == nil {
		return nil, false
	}
	return o.MonitorType, true
}

// HasMonitorType returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasMonitorType() bool {
	if o != nil && o.MonitorType != nil {
		return true
	}

	return false
}

// SetMonitorType gets a reference to the given string and assigns it to the MonitorType field.
func (o *AlertsLibraryAlertAllOf) SetMonitorType(v string) {
	o.MonitorType = &v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetEntityIds() []string {
	if o == nil || o.EntityIds == nil {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetEntityIdsOk() ([]string, bool) {
	if o == nil || o.EntityIds == nil {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasEntityIds() bool {
	if o != nil && o.EntityIds != nil {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *AlertsLibraryAlertAllOf) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetEntities() []AlertEntityInfo {
	if o == nil || o.Entities == nil {
		var ret []AlertEntityInfo
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetEntitiesOk() ([]AlertEntityInfo, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []AlertEntityInfo and assigns it to the Entities field.
func (o *AlertsLibraryAlertAllOf) SetEntities(v []AlertEntityInfo) {
	o.Entities = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AlertsLibraryAlertAllOf) SetNotes(v string) {
	o.Notes = &v
}

// GetExtraDetails returns the ExtraDetails field value if set, zero value otherwise.
func (o *AlertsLibraryAlertAllOf) GetExtraDetails() ExtraDetails {
	if o == nil || o.ExtraDetails == nil {
		var ret ExtraDetails
		return ret
	}
	return *o.ExtraDetails
}

// GetExtraDetailsOk returns a tuple with the ExtraDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertAllOf) GetExtraDetailsOk() (*ExtraDetails, bool) {
	if o == nil || o.ExtraDetails == nil {
		return nil, false
	}
	return o.ExtraDetails, true
}

// HasExtraDetails returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasExtraDetails() bool {
	if o != nil && o.ExtraDetails != nil {
		return true
	}

	return false
}

// SetExtraDetails gets a reference to the given ExtraDetails and assigns it to the ExtraDetails field.
func (o *AlertsLibraryAlertAllOf) SetExtraDetails(v ExtraDetails) {
	o.ExtraDetails = &v
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsLibraryAlertAllOf) GetAlertCondition() string {
	if o == nil || o.AlertCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.AlertCondition.Get()
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertsLibraryAlertAllOf) GetAlertConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertCondition.Get(), o.AlertCondition.IsSet()
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *AlertsLibraryAlertAllOf) HasAlertCondition() bool {
	if o != nil && o.AlertCondition.IsSet() {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given NullableString and assigns it to the AlertCondition field.
func (o *AlertsLibraryAlertAllOf) SetAlertCondition(v string) {
	o.AlertCondition.Set(&v)
}
// SetAlertConditionNil sets the value for AlertCondition to be an explicit nil
func (o *AlertsLibraryAlertAllOf) SetAlertConditionNil() {
	o.AlertCondition.Set(nil)
}

// UnsetAlertCondition ensures that no value is present for AlertCondition, not even an explicit nil
func (o *AlertsLibraryAlertAllOf) UnsetAlertCondition() {
	o.AlertCondition.Unset()
}

func (o AlertsLibraryAlertAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MonitorId != nil {
		toSerialize["monitorId"] = o.MonitorId
	}
	if o.ResolvedAt.IsSet() {
		toSerialize["resolvedAt"] = o.ResolvedAt.Get()
	}
	if o.AbnormalityStartTime != nil {
		toSerialize["abnormalityStartTime"] = o.AbnormalityStartTime
	}
	if o.AlertType != nil {
		toSerialize["alertType"] = o.AlertType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.MonitorQueries != nil {
		toSerialize["monitorQueries"] = o.MonitorQueries
	}
	if o.TriggerQueries != nil {
		toSerialize["triggerQueries"] = o.TriggerQueries
	}
	if o.MonitorUrl != nil {
		toSerialize["monitorUrl"] = o.MonitorUrl
	}
	if o.TriggerQueryUrl != nil {
		toSerialize["triggerQueryUrl"] = o.TriggerQueryUrl
	}
	if o.TriggerConditions != nil {
		toSerialize["triggerConditions"] = o.TriggerConditions
	}
	if o.TriggerValue != nil {
		toSerialize["triggerValue"] = o.TriggerValue
	}
	if o.MonitorType != nil {
		toSerialize["monitorType"] = o.MonitorType
	}
	if o.EntityIds != nil {
		toSerialize["entityIds"] = o.EntityIds
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.ExtraDetails != nil {
		toSerialize["extraDetails"] = o.ExtraDetails
	}
	if o.AlertCondition.IsSet() {
		toSerialize["alertCondition"] = o.AlertCondition.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAlertsLibraryAlertAllOf struct {
	value *AlertsLibraryAlertAllOf
	isSet bool
}

func (v NullableAlertsLibraryAlertAllOf) Get() *AlertsLibraryAlertAllOf {
	return v.value
}

func (v *NullableAlertsLibraryAlertAllOf) Set(val *AlertsLibraryAlertAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryAlertAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryAlertAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryAlertAllOf(val *AlertsLibraryAlertAllOf) *NullableAlertsLibraryAlertAllOf {
	return &NullableAlertsLibraryAlertAllOf{value: val, isSet: true}
}

func (v NullableAlertsLibraryAlertAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryAlertAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


