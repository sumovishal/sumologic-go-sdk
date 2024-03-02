# PaginatedReportSchedules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportSchedules** | [**[]ReportSchedule**](ReportSchedule.md) | List of dashboard report schedules. | 
**Next** | Pointer to **string** | Next continuation token. &#x60;token&#x60; is set to null when no more pages are left. | [optional] 

## Methods

### NewPaginatedReportSchedules

`func NewPaginatedReportSchedules(reportSchedules []ReportSchedule, ) *PaginatedReportSchedules`

NewPaginatedReportSchedules instantiates a new PaginatedReportSchedules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedReportSchedulesWithDefaults

`func NewPaginatedReportSchedulesWithDefaults() *PaginatedReportSchedules`

NewPaginatedReportSchedulesWithDefaults instantiates a new PaginatedReportSchedules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportSchedules

`func (o *PaginatedReportSchedules) GetReportSchedules() []ReportSchedule`

GetReportSchedules returns the ReportSchedules field if non-nil, zero value otherwise.

### GetReportSchedulesOk

`func (o *PaginatedReportSchedules) GetReportSchedulesOk() (*[]ReportSchedule, bool)`

GetReportSchedulesOk returns a tuple with the ReportSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportSchedules

`func (o *PaginatedReportSchedules) SetReportSchedules(v []ReportSchedule)`

SetReportSchedules sets ReportSchedules field to given value.


### GetNext

`func (o *PaginatedReportSchedules) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedReportSchedules) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedReportSchedules) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedReportSchedules) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


