# MonitorsLibraryMonitorExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorType** | **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor.   3. &#x60;Slo&#x60;: A SLO based monitor. Currently SLO based monitor is available in closed beta (Notify your Sumo Logic representative in order to get the early access). | 
**EvaluationDelay** | Pointer to **string** | The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time. | [optional] [default to "0m"]
**AlertName** | Pointer to **string** | The name of the alert(s) triggered from this monitor. Monitor name will be used if not specified. All template variables can be used here except {{AlertName}}, {{AlertResponseURL}}, {{ResultsJson}}, and {{Playbook}}. | [optional] 
**RunAs** | Pointer to **map[string]interface{}** |  | [optional] 
**NotificationGroupFields** | Pointer to **[]string** | The set of fields to be used to group alert notifications for a monitor. The value of this field will be considered only when &#39;groupNotifications&#39; is true. The fields with very high cardinality such as &#x60;_blockid&#x60;, &#x60;_raw&#x60;, &#x60;_messagetime&#x60;, &#x60;_receipttime&#x60;, and &#x60;_messageid&#x60; are not allowed for Alert Grouping. | [optional] 
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | 
**Triggers** | [**[]TriggerCondition**](TriggerCondition.md) | Defines the conditions of when to send notifications. | 
**TimeZone** | Pointer to **string** | Time zone identifier for monitor notifications. Follow the format in [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | [optional] 
**Notifications** | Pointer to [**[]MonitorNotification**](MonitorNotification.md) | The notifications the monitor will send when the respective trigger condition is met. | [optional] [default to []]
**IsDisabled** | Pointer to **bool** | Whether or not the monitor is disabled. Disabled monitors will not run, and will not generate or send notifications. | [optional] [default to false]
**GroupNotifications** | Pointer to **bool** | Whether or not to group notifications for individual items that meet the trigger condition. | [optional] [default to true]
**Playbook** | Pointer to **string** | Notes such as links and instruction to help you resolve alerts triggered by this monitor. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more. | [optional] [default to ""]
**SloId** | Pointer to **string** | Identifier of the SLO definition for the monitor. This is only applicable for SLO type monitors. | [optional] 
**AutomatedPlaybookIds** | Pointer to **[]string** | The set of automated playbook ids for a monitor. | [optional] [default to []]

## Methods

### NewMonitorsLibraryMonitorExport

`func NewMonitorsLibraryMonitorExport(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, ) *MonitorsLibraryMonitorExport`

NewMonitorsLibraryMonitorExport instantiates a new MonitorsLibraryMonitorExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorsLibraryMonitorExportWithDefaults

`func NewMonitorsLibraryMonitorExportWithDefaults() *MonitorsLibraryMonitorExport`

NewMonitorsLibraryMonitorExportWithDefaults instantiates a new MonitorsLibraryMonitorExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *MonitorsLibraryMonitorExport) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *MonitorsLibraryMonitorExport) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *MonitorsLibraryMonitorExport) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.


### GetEvaluationDelay

`func (o *MonitorsLibraryMonitorExport) GetEvaluationDelay() string`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *MonitorsLibraryMonitorExport) GetEvaluationDelayOk() (*string, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDelay

`func (o *MonitorsLibraryMonitorExport) SetEvaluationDelay(v string)`

SetEvaluationDelay sets EvaluationDelay field to given value.

### HasEvaluationDelay

`func (o *MonitorsLibraryMonitorExport) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### GetAlertName

`func (o *MonitorsLibraryMonitorExport) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *MonitorsLibraryMonitorExport) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *MonitorsLibraryMonitorExport) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *MonitorsLibraryMonitorExport) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetRunAs

`func (o *MonitorsLibraryMonitorExport) GetRunAs() map[string]interface{}`

GetRunAs returns the RunAs field if non-nil, zero value otherwise.

### GetRunAsOk

`func (o *MonitorsLibraryMonitorExport) GetRunAsOk() (*map[string]interface{}, bool)`

GetRunAsOk returns a tuple with the RunAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAs

`func (o *MonitorsLibraryMonitorExport) SetRunAs(v map[string]interface{})`

SetRunAs sets RunAs field to given value.

### HasRunAs

`func (o *MonitorsLibraryMonitorExport) HasRunAs() bool`

HasRunAs returns a boolean if a field has been set.

### GetNotificationGroupFields

`func (o *MonitorsLibraryMonitorExport) GetNotificationGroupFields() []string`

GetNotificationGroupFields returns the NotificationGroupFields field if non-nil, zero value otherwise.

### GetNotificationGroupFieldsOk

`func (o *MonitorsLibraryMonitorExport) GetNotificationGroupFieldsOk() (*[]string, bool)`

GetNotificationGroupFieldsOk returns a tuple with the NotificationGroupFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationGroupFields

`func (o *MonitorsLibraryMonitorExport) SetNotificationGroupFields(v []string)`

SetNotificationGroupFields sets NotificationGroupFields field to given value.

### HasNotificationGroupFields

`func (o *MonitorsLibraryMonitorExport) HasNotificationGroupFields() bool`

HasNotificationGroupFields returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorsLibraryMonitorExport) GetQueries() []MonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorsLibraryMonitorExport) GetQueriesOk() (*[]MonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorsLibraryMonitorExport) SetQueries(v []MonitorQuery)`

SetQueries sets Queries field to given value.


### GetTriggers

`func (o *MonitorsLibraryMonitorExport) GetTriggers() []TriggerCondition`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MonitorsLibraryMonitorExport) GetTriggersOk() (*[]TriggerCondition, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *MonitorsLibraryMonitorExport) SetTriggers(v []TriggerCondition)`

SetTriggers sets Triggers field to given value.


### GetTimeZone

`func (o *MonitorsLibraryMonitorExport) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MonitorsLibraryMonitorExport) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MonitorsLibraryMonitorExport) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MonitorsLibraryMonitorExport) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetNotifications

`func (o *MonitorsLibraryMonitorExport) GetNotifications() []MonitorNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *MonitorsLibraryMonitorExport) GetNotificationsOk() (*[]MonitorNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *MonitorsLibraryMonitorExport) SetNotifications(v []MonitorNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *MonitorsLibraryMonitorExport) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetIsDisabled

`func (o *MonitorsLibraryMonitorExport) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *MonitorsLibraryMonitorExport) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *MonitorsLibraryMonitorExport) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *MonitorsLibraryMonitorExport) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetGroupNotifications

`func (o *MonitorsLibraryMonitorExport) GetGroupNotifications() bool`

GetGroupNotifications returns the GroupNotifications field if non-nil, zero value otherwise.

### GetGroupNotificationsOk

`func (o *MonitorsLibraryMonitorExport) GetGroupNotificationsOk() (*bool, bool)`

GetGroupNotificationsOk returns a tuple with the GroupNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNotifications

`func (o *MonitorsLibraryMonitorExport) SetGroupNotifications(v bool)`

SetGroupNotifications sets GroupNotifications field to given value.

### HasGroupNotifications

`func (o *MonitorsLibraryMonitorExport) HasGroupNotifications() bool`

HasGroupNotifications returns a boolean if a field has been set.

### GetPlaybook

`func (o *MonitorsLibraryMonitorExport) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *MonitorsLibraryMonitorExport) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *MonitorsLibraryMonitorExport) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *MonitorsLibraryMonitorExport) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetSloId

`func (o *MonitorsLibraryMonitorExport) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *MonitorsLibraryMonitorExport) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *MonitorsLibraryMonitorExport) SetSloId(v string)`

SetSloId sets SloId field to given value.

### HasSloId

`func (o *MonitorsLibraryMonitorExport) HasSloId() bool`

HasSloId returns a boolean if a field has been set.

### GetAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorExport) GetAutomatedPlaybookIds() []string`

GetAutomatedPlaybookIds returns the AutomatedPlaybookIds field if non-nil, zero value otherwise.

### GetAutomatedPlaybookIdsOk

`func (o *MonitorsLibraryMonitorExport) GetAutomatedPlaybookIdsOk() (*[]string, bool)`

GetAutomatedPlaybookIdsOk returns a tuple with the AutomatedPlaybookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorExport) SetAutomatedPlaybookIds(v []string)`

SetAutomatedPlaybookIds sets AutomatedPlaybookIds field to given value.

### HasAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorExport) HasAutomatedPlaybookIds() bool`

HasAutomatedPlaybookIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


