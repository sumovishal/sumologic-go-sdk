# ReportScheduleRequest

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

## Methods

### NewReportScheduleRequest

`func NewReportScheduleRequest(dashboardId string, reportFormat string, scheduleType string, timeZone string, emailNotification ReportScheduleRequestEmailNotification, ) *ReportScheduleRequest`

NewReportScheduleRequest instantiates a new ReportScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportScheduleRequestWithDefaults

`func NewReportScheduleRequestWithDefaults() *ReportScheduleRequest`

NewReportScheduleRequestWithDefaults instantiates a new ReportScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardId

`func (o *ReportScheduleRequest) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *ReportScheduleRequest) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *ReportScheduleRequest) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.


### GetTimeRange

`func (o *ReportScheduleRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *ReportScheduleRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *ReportScheduleRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *ReportScheduleRequest) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetVariableValues

`func (o *ReportScheduleRequest) GetVariableValues() VariablesValuesData`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *ReportScheduleRequest) GetVariableValuesOk() (*VariablesValuesData, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *ReportScheduleRequest) SetVariableValues(v VariablesValuesData)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *ReportScheduleRequest) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.

### GetReportFormat

`func (o *ReportScheduleRequest) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *ReportScheduleRequest) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *ReportScheduleRequest) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.


### GetScheduleType

`func (o *ReportScheduleRequest) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ReportScheduleRequest) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ReportScheduleRequest) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.


### GetCronExpression

`func (o *ReportScheduleRequest) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ReportScheduleRequest) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ReportScheduleRequest) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *ReportScheduleRequest) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetTimeZone

`func (o *ReportScheduleRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ReportScheduleRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ReportScheduleRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetEmailNotification

`func (o *ReportScheduleRequest) GetEmailNotification() ReportScheduleRequestEmailNotification`

GetEmailNotification returns the EmailNotification field if non-nil, zero value otherwise.

### GetEmailNotificationOk

`func (o *ReportScheduleRequest) GetEmailNotificationOk() (*ReportScheduleRequestEmailNotification, bool)`

GetEmailNotificationOk returns a tuple with the EmailNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotification

`func (o *ReportScheduleRequest) SetEmailNotification(v ReportScheduleRequestEmailNotification)`

SetEmailNotification sets EmailNotification field to given value.


### GetIsActive

`func (o *ReportScheduleRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ReportScheduleRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ReportScheduleRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ReportScheduleRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


