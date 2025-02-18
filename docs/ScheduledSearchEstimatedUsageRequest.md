# ScheduledSearchEstimatedUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | The text of a logs search query. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**CronSchedule** | Pointer to **string** | Cron-like expression specifying the search&#39;s schedule. Field scheduleType must be set to \&quot;Custom\&quot;, otherwise, scheduleType takes precedence over cronSchedule. | [optional] 
**ScheduleType** | **string** | Run schedule of the scheduled search. Set to \&quot;Custom\&quot; to specify the schedule with a CRON expression. Please note that  with Custom, 1Day and 1Week schedule types you need to provide the corresponding cron expression to determine when to  actually run the search. e.g. Sample Valid Cron for 1Day is \&quot;0 0 16 ? * 2-6 *\&quot;. Possible schedule types are:   - &#x60;RealTime&#x60;   - &#x60;15Minutes&#x60;   - &#x60;1Hour&#x60;   - &#x60;2Hours&#x60;   - &#x60;4Hours&#x60;   - &#x60;6Hours&#x60;   - &#x60;8Hours&#x60;   - &#x60;12Hours&#x60;   - &#x60;1Day&#x60;   - &#x60;1Week&#x60;   - &#x60;Custom&#x60; | 
**ByReceiptTime** | Pointer to **bool** | Set it to true to run the search using receipt time. By default, searches do not run by receipt time. | [optional] [default to false]
**QueryParameters** | Pointer to [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | An array of search query parameter objects. | [optional] 
**TimeZone** | **string** | Time zone identifier for the estimates. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 

## Methods

### NewScheduledSearchEstimatedUsageRequest

`func NewScheduledSearchEstimatedUsageRequest(queryString string, timeRange ResolvableTimeRange, scheduleType string, timeZone string, ) *ScheduledSearchEstimatedUsageRequest`

NewScheduledSearchEstimatedUsageRequest instantiates a new ScheduledSearchEstimatedUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledSearchEstimatedUsageRequestWithDefaults

`func NewScheduledSearchEstimatedUsageRequestWithDefaults() *ScheduledSearchEstimatedUsageRequest`

NewScheduledSearchEstimatedUsageRequestWithDefaults instantiates a new ScheduledSearchEstimatedUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *ScheduledSearchEstimatedUsageRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *ScheduledSearchEstimatedUsageRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *ScheduledSearchEstimatedUsageRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetTimeRange

`func (o *ScheduledSearchEstimatedUsageRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *ScheduledSearchEstimatedUsageRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *ScheduledSearchEstimatedUsageRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetCronSchedule

`func (o *ScheduledSearchEstimatedUsageRequest) GetCronSchedule() string`

GetCronSchedule returns the CronSchedule field if non-nil, zero value otherwise.

### GetCronScheduleOk

`func (o *ScheduledSearchEstimatedUsageRequest) GetCronScheduleOk() (*string, bool)`

GetCronScheduleOk returns a tuple with the CronSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronSchedule

`func (o *ScheduledSearchEstimatedUsageRequest) SetCronSchedule(v string)`

SetCronSchedule sets CronSchedule field to given value.

### HasCronSchedule

`func (o *ScheduledSearchEstimatedUsageRequest) HasCronSchedule() bool`

HasCronSchedule returns a boolean if a field has been set.

### GetScheduleType

`func (o *ScheduledSearchEstimatedUsageRequest) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ScheduledSearchEstimatedUsageRequest) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ScheduledSearchEstimatedUsageRequest) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.


### GetByReceiptTime

`func (o *ScheduledSearchEstimatedUsageRequest) GetByReceiptTime() bool`

GetByReceiptTime returns the ByReceiptTime field if non-nil, zero value otherwise.

### GetByReceiptTimeOk

`func (o *ScheduledSearchEstimatedUsageRequest) GetByReceiptTimeOk() (*bool, bool)`

GetByReceiptTimeOk returns a tuple with the ByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByReceiptTime

`func (o *ScheduledSearchEstimatedUsageRequest) SetByReceiptTime(v bool)`

SetByReceiptTime sets ByReceiptTime field to given value.

### HasByReceiptTime

`func (o *ScheduledSearchEstimatedUsageRequest) HasByReceiptTime() bool`

HasByReceiptTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *ScheduledSearchEstimatedUsageRequest) GetQueryParameters() []QueryParameterSyncDefinition`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *ScheduledSearchEstimatedUsageRequest) GetQueryParametersOk() (*[]QueryParameterSyncDefinition, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *ScheduledSearchEstimatedUsageRequest) SetQueryParameters(v []QueryParameterSyncDefinition)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *ScheduledSearchEstimatedUsageRequest) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### GetTimeZone

`func (o *ScheduledSearchEstimatedUsageRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ScheduledSearchEstimatedUsageRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ScheduledSearchEstimatedUsageRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


