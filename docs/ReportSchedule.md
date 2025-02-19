# ReportSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardId** | **string** | Identifier of dashboard the schedule will generate report for. | 
**TimeRange** | Pointer to [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] 
**VariableValues** | Pointer to [**VariablesValuesData**](VariablesValuesData.md) |  | [optional] 
**ReportFormat** | **string** | File format of the report. Can be &#x60;Pdf&#x60; or &#x60;Png&#x60;. &#x60;Pdf&#x60; is portable document format. &#x60;Png&#x60; is portable graphics image format. | 
**ScheduleType** | **string** | Run schedule of the scheduled report. Set to \&quot;Custom\&quot; to specify the schedule with a CRON expression. Possible schedule types are:   - &#x60;RealTime&#x60;   - &#x60;15Minutes&#x60;   - &#x60;1Hour&#x60;   - &#x60;2Hours&#x60;   - &#x60;4Hours&#x60;   - &#x60;6Hours&#x60;   - &#x60;8Hours&#x60;   - &#x60;12Hours&#x60;   - &#x60;1Day&#x60;   - &#x60;1Week&#x60;   - &#x60;Custom&#x60; | 
**CronExpression** | Pointer to **string** | Cron-like expression specifying the report&#39;s schedule. Field scheduleType must be set to \&quot;Custom\&quot;, otherwise, scheduleType takes precedence over cronExpression. | [optional] 
**TimeZone** | **string** | Time zone identifier for time specification. Either an abbreviation such as \&quot;PST\&quot;, a full name such as \&quot;America/Los_Angeles\&quot;, or a custom ID such as \&quot;GMT-8:00\&quot;. Note that the support of abbreviations is for JDK 1.1.x compatibility only and full names should be used. | 
**EmailNotification** | [**ReportScheduleRequestEmailNotification**](ReportScheduleRequestEmailNotification.md) |  | 
**IsActive** | Pointer to **bool** | Is the dashboard report schedule active | [optional] [default to true]
**ScheduleId** | Pointer to **string** | Identifier of the dashboard report schedule. | [optional] 

## Methods

### NewReportSchedule

`func NewReportSchedule(dashboardId string, reportFormat string, scheduleType string, timeZone string, emailNotification ReportScheduleRequestEmailNotification, ) *ReportSchedule`

NewReportSchedule instantiates a new ReportSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportScheduleWithDefaults

`func NewReportScheduleWithDefaults() *ReportSchedule`

NewReportScheduleWithDefaults instantiates a new ReportSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardId

`func (o *ReportSchedule) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *ReportSchedule) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *ReportSchedule) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.


### GetTimeRange

`func (o *ReportSchedule) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *ReportSchedule) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *ReportSchedule) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *ReportSchedule) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetVariableValues

`func (o *ReportSchedule) GetVariableValues() VariablesValuesData`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *ReportSchedule) GetVariableValuesOk() (*VariablesValuesData, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *ReportSchedule) SetVariableValues(v VariablesValuesData)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *ReportSchedule) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.

### GetReportFormat

`func (o *ReportSchedule) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *ReportSchedule) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *ReportSchedule) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.


### GetScheduleType

`func (o *ReportSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ReportSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ReportSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.


### GetCronExpression

`func (o *ReportSchedule) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ReportSchedule) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ReportSchedule) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *ReportSchedule) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetTimeZone

`func (o *ReportSchedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ReportSchedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ReportSchedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetEmailNotification

`func (o *ReportSchedule) GetEmailNotification() ReportScheduleRequestEmailNotification`

GetEmailNotification returns the EmailNotification field if non-nil, zero value otherwise.

### GetEmailNotificationOk

`func (o *ReportSchedule) GetEmailNotificationOk() (*ReportScheduleRequestEmailNotification, bool)`

GetEmailNotificationOk returns a tuple with the EmailNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotification

`func (o *ReportSchedule) SetEmailNotification(v ReportScheduleRequestEmailNotification)`

SetEmailNotification sets EmailNotification field to given value.


### GetIsActive

`func (o *ReportSchedule) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ReportSchedule) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ReportSchedule) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ReportSchedule) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetScheduleId

`func (o *ReportSchedule) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ReportSchedule) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ReportSchedule) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *ReportSchedule) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


