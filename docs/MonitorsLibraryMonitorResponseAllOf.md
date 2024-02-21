# MonitorsLibraryMonitorResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorType** | **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor.   3. &#x60;Slo&#x60;: A SLO based monitor. Currently SLO based monitor is available in closed beta (Notify your Sumo Logic representative in order to get the early access). | 
**EvaluationDelay** | Pointer to **string** | The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time. | [optional] [default to "0m"]
**AlertName** | Pointer to **string** | The name of the alert(s) triggered from this monitor. Monitor name will be used if not specified. All template variables can be used here except {{AlertName}}, {{AlertResponseURL}}, {{ResultsJson}}, and {{Playbook}}. | [optional] 
**RunAs** | Pointer to [**MonitorsLibraryMonitorAllOfRunAs**](MonitorsLibraryMonitorAllOfRunAs.md) |  | [optional] 
**NotificationGroupFields** | Pointer to **[]string** | The set of fields to be used to group alert notifications for a monitor. The value of this field will be considered only when &#39;groupNotifications&#39; is true. The fields with very high cardinality such as &#x60;_blockid&#x60;, &#x60;_raw&#x60;, &#x60;_messagetime&#x60;, &#x60;_receipttime&#x60;, and &#x60;_messageid&#x60; are not allowed for Alert Grouping. | [optional] 
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | 
**Triggers** | [**[]TriggerCondition**](TriggerCondition.md) | Defines the conditions of when to send notifications. | 
**Notifications** | Pointer to [**[]MonitorNotification**](MonitorNotification.md) | The notifications the monitor will send when the respective trigger condition is met. | [optional] [default to []]
**IsDisabled** | Pointer to **bool** | Whether or not the monitor is disabled. Disabled monitors will not run, and will not generate or send notifications. | [optional] [default to false]
**Status** | Pointer to **[]string** | The current status of the monitor. Each monitor can have one or more status values. Valid values:   1. &#x60;Normal&#x60;: The monitor is running normally and does not have any currently triggered conditions.   2. &#x60;Critical&#x60;: The Critical trigger condition has been met.   3. &#x60;Warning&#x60;: The Warning trigger condition has been met.   4. &#x60;MissingData&#x60;: The MissingData trigger condition has been met.   5. &#x60;Disabled&#x60;: The monitor has been disabled and is not currently running. | [optional] 
**GroupNotifications** | Pointer to **bool** | Whether or not to group notifications for individual items that meet the trigger condition. | [optional] [default to true]
**Warnings** | Pointer to **map[string]string** | Monitor manager warnings | [optional] 
**Playbook** | Pointer to **string** | Notes such as links and instruction to help you resolve alerts triggered by this monitor. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more. | [optional] [default to ""]
**SloId** | Pointer to **string** | Identifier of the SLO definition for the monitor. This is only applicable for SLO type monitors. | [optional] 
**AutomatedPlaybookIds** | Pointer to **[]string** | The set of automated playbook ids for a monitor. | [optional] [default to []]

## Methods

### NewMonitorsLibraryMonitorResponseAllOf

`func NewMonitorsLibraryMonitorResponseAllOf(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, ) *MonitorsLibraryMonitorResponseAllOf`

NewMonitorsLibraryMonitorResponseAllOf instantiates a new MonitorsLibraryMonitorResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorsLibraryMonitorResponseAllOfWithDefaults

`func NewMonitorsLibraryMonitorResponseAllOfWithDefaults() *MonitorsLibraryMonitorResponseAllOf`

NewMonitorsLibraryMonitorResponseAllOfWithDefaults instantiates a new MonitorsLibraryMonitorResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *MonitorsLibraryMonitorResponseAllOf) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *MonitorsLibraryMonitorResponseAllOf) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.


### GetEvaluationDelay

`func (o *MonitorsLibraryMonitorResponseAllOf) GetEvaluationDelay() string`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetEvaluationDelayOk() (*string, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDelay

`func (o *MonitorsLibraryMonitorResponseAllOf) SetEvaluationDelay(v string)`

SetEvaluationDelay sets EvaluationDelay field to given value.

### HasEvaluationDelay

`func (o *MonitorsLibraryMonitorResponseAllOf) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### GetAlertName

`func (o *MonitorsLibraryMonitorResponseAllOf) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *MonitorsLibraryMonitorResponseAllOf) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *MonitorsLibraryMonitorResponseAllOf) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetRunAs

`func (o *MonitorsLibraryMonitorResponseAllOf) GetRunAs() MonitorsLibraryMonitorAllOfRunAs`

GetRunAs returns the RunAs field if non-nil, zero value otherwise.

### GetRunAsOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetRunAsOk() (*MonitorsLibraryMonitorAllOfRunAs, bool)`

GetRunAsOk returns a tuple with the RunAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAs

`func (o *MonitorsLibraryMonitorResponseAllOf) SetRunAs(v MonitorsLibraryMonitorAllOfRunAs)`

SetRunAs sets RunAs field to given value.

### HasRunAs

`func (o *MonitorsLibraryMonitorResponseAllOf) HasRunAs() bool`

HasRunAs returns a boolean if a field has been set.

### GetNotificationGroupFields

`func (o *MonitorsLibraryMonitorResponseAllOf) GetNotificationGroupFields() []string`

GetNotificationGroupFields returns the NotificationGroupFields field if non-nil, zero value otherwise.

### GetNotificationGroupFieldsOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetNotificationGroupFieldsOk() (*[]string, bool)`

GetNotificationGroupFieldsOk returns a tuple with the NotificationGroupFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationGroupFields

`func (o *MonitorsLibraryMonitorResponseAllOf) SetNotificationGroupFields(v []string)`

SetNotificationGroupFields sets NotificationGroupFields field to given value.

### HasNotificationGroupFields

`func (o *MonitorsLibraryMonitorResponseAllOf) HasNotificationGroupFields() bool`

HasNotificationGroupFields returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorsLibraryMonitorResponseAllOf) GetQueries() []MonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetQueriesOk() (*[]MonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorsLibraryMonitorResponseAllOf) SetQueries(v []MonitorQuery)`

SetQueries sets Queries field to given value.


### GetTriggers

`func (o *MonitorsLibraryMonitorResponseAllOf) GetTriggers() []TriggerCondition`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetTriggersOk() (*[]TriggerCondition, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *MonitorsLibraryMonitorResponseAllOf) SetTriggers(v []TriggerCondition)`

SetTriggers sets Triggers field to given value.


### GetNotifications

`func (o *MonitorsLibraryMonitorResponseAllOf) GetNotifications() []MonitorNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetNotificationsOk() (*[]MonitorNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *MonitorsLibraryMonitorResponseAllOf) SetNotifications(v []MonitorNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *MonitorsLibraryMonitorResponseAllOf) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetIsDisabled

`func (o *MonitorsLibraryMonitorResponseAllOf) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *MonitorsLibraryMonitorResponseAllOf) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *MonitorsLibraryMonitorResponseAllOf) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetStatus

`func (o *MonitorsLibraryMonitorResponseAllOf) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitorsLibraryMonitorResponseAllOf) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitorsLibraryMonitorResponseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGroupNotifications

`func (o *MonitorsLibraryMonitorResponseAllOf) GetGroupNotifications() bool`

GetGroupNotifications returns the GroupNotifications field if non-nil, zero value otherwise.

### GetGroupNotificationsOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetGroupNotificationsOk() (*bool, bool)`

GetGroupNotificationsOk returns a tuple with the GroupNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNotifications

`func (o *MonitorsLibraryMonitorResponseAllOf) SetGroupNotifications(v bool)`

SetGroupNotifications sets GroupNotifications field to given value.

### HasGroupNotifications

`func (o *MonitorsLibraryMonitorResponseAllOf) HasGroupNotifications() bool`

HasGroupNotifications returns a boolean if a field has been set.

### GetWarnings

`func (o *MonitorsLibraryMonitorResponseAllOf) GetWarnings() map[string]string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetWarningsOk() (*map[string]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MonitorsLibraryMonitorResponseAllOf) SetWarnings(v map[string]string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MonitorsLibraryMonitorResponseAllOf) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetPlaybook

`func (o *MonitorsLibraryMonitorResponseAllOf) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *MonitorsLibraryMonitorResponseAllOf) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *MonitorsLibraryMonitorResponseAllOf) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetSloId

`func (o *MonitorsLibraryMonitorResponseAllOf) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *MonitorsLibraryMonitorResponseAllOf) SetSloId(v string)`

SetSloId sets SloId field to given value.

### HasSloId

`func (o *MonitorsLibraryMonitorResponseAllOf) HasSloId() bool`

HasSloId returns a boolean if a field has been set.

### GetAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorResponseAllOf) GetAutomatedPlaybookIds() []string`

GetAutomatedPlaybookIds returns the AutomatedPlaybookIds field if non-nil, zero value otherwise.

### GetAutomatedPlaybookIdsOk

`func (o *MonitorsLibraryMonitorResponseAllOf) GetAutomatedPlaybookIdsOk() (*[]string, bool)`

GetAutomatedPlaybookIdsOk returns a tuple with the AutomatedPlaybookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorResponseAllOf) SetAutomatedPlaybookIds(v []string)`

SetAutomatedPlaybookIds sets AutomatedPlaybookIds field to given value.

### HasAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorResponseAllOf) HasAutomatedPlaybookIds() bool`

HasAutomatedPlaybookIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


