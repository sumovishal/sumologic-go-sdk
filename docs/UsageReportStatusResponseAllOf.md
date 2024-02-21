# UsageReportStatusResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status export | [optional] 
**StatusMessage** | Pointer to **string** | Status message export | [optional] 
**ReportDownloadURL** | Pointer to **string** | S3 presigned download URL for the report. It is valid for 10 minutes. | [optional] 

## Methods

### NewUsageReportStatusResponseAllOf

`func NewUsageReportStatusResponseAllOf() *UsageReportStatusResponseAllOf`

NewUsageReportStatusResponseAllOf instantiates a new UsageReportStatusResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageReportStatusResponseAllOfWithDefaults

`func NewUsageReportStatusResponseAllOfWithDefaults() *UsageReportStatusResponseAllOf`

NewUsageReportStatusResponseAllOfWithDefaults instantiates a new UsageReportStatusResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UsageReportStatusResponseAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsageReportStatusResponseAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsageReportStatusResponseAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsageReportStatusResponseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UsageReportStatusResponseAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UsageReportStatusResponseAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UsageReportStatusResponseAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UsageReportStatusResponseAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetReportDownloadURL

`func (o *UsageReportStatusResponseAllOf) GetReportDownloadURL() string`

GetReportDownloadURL returns the ReportDownloadURL field if non-nil, zero value otherwise.

### GetReportDownloadURLOk

`func (o *UsageReportStatusResponseAllOf) GetReportDownloadURLOk() (*string, bool)`

GetReportDownloadURLOk returns a tuple with the ReportDownloadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDownloadURL

`func (o *UsageReportStatusResponseAllOf) SetReportDownloadURL(v string)`

SetReportDownloadURL sets ReportDownloadURL field to given value.

### HasReportDownloadURL

`func (o *UsageReportStatusResponseAllOf) HasReportDownloadURL() bool`

HasReportDownloadURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


