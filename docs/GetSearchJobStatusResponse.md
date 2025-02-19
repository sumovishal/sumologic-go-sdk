# GetSearchJobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | Search job state. | [optional] 
**MessageCount** | Pointer to **int32** | Number of messages found or produced so far. | [optional] 
**RecordCount** | Pointer to **int32** | Number of records found or produced so far. | [optional] 
**HistogramBuckets** | Pointer to [**[]HistogramBucket**](HistogramBucket.md) | Histogram buckets for the query. | [optional] 
**PendingErrors** | Pointer to **[]string** | Pending errors that have accumulated since the last time the status was requested. | [optional] 
**PendingWarnings** | Pointer to **[]string** | Pending warnings that have accumulated since the last time the status was requested. | [optional] 

## Methods

### NewGetSearchJobStatusResponse

`func NewGetSearchJobStatusResponse() *GetSearchJobStatusResponse`

NewGetSearchJobStatusResponse instantiates a new GetSearchJobStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSearchJobStatusResponseWithDefaults

`func NewGetSearchJobStatusResponseWithDefaults() *GetSearchJobStatusResponse`

NewGetSearchJobStatusResponseWithDefaults instantiates a new GetSearchJobStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *GetSearchJobStatusResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSearchJobStatusResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSearchJobStatusResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSearchJobStatusResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageCount

`func (o *GetSearchJobStatusResponse) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *GetSearchJobStatusResponse) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *GetSearchJobStatusResponse) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.

### HasMessageCount

`func (o *GetSearchJobStatusResponse) HasMessageCount() bool`

HasMessageCount returns a boolean if a field has been set.

### GetRecordCount

`func (o *GetSearchJobStatusResponse) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *GetSearchJobStatusResponse) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *GetSearchJobStatusResponse) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *GetSearchJobStatusResponse) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.

### GetHistogramBuckets

`func (o *GetSearchJobStatusResponse) GetHistogramBuckets() []HistogramBucket`

GetHistogramBuckets returns the HistogramBuckets field if non-nil, zero value otherwise.

### GetHistogramBucketsOk

`func (o *GetSearchJobStatusResponse) GetHistogramBucketsOk() (*[]HistogramBucket, bool)`

GetHistogramBucketsOk returns a tuple with the HistogramBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramBuckets

`func (o *GetSearchJobStatusResponse) SetHistogramBuckets(v []HistogramBucket)`

SetHistogramBuckets sets HistogramBuckets field to given value.

### HasHistogramBuckets

`func (o *GetSearchJobStatusResponse) HasHistogramBuckets() bool`

HasHistogramBuckets returns a boolean if a field has been set.

### GetPendingErrors

`func (o *GetSearchJobStatusResponse) GetPendingErrors() []string`

GetPendingErrors returns the PendingErrors field if non-nil, zero value otherwise.

### GetPendingErrorsOk

`func (o *GetSearchJobStatusResponse) GetPendingErrorsOk() (*[]string, bool)`

GetPendingErrorsOk returns a tuple with the PendingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingErrors

`func (o *GetSearchJobStatusResponse) SetPendingErrors(v []string)`

SetPendingErrors sets PendingErrors field to given value.

### HasPendingErrors

`func (o *GetSearchJobStatusResponse) HasPendingErrors() bool`

HasPendingErrors returns a boolean if a field has been set.

### GetPendingWarnings

`func (o *GetSearchJobStatusResponse) GetPendingWarnings() []string`

GetPendingWarnings returns the PendingWarnings field if non-nil, zero value otherwise.

### GetPendingWarningsOk

`func (o *GetSearchJobStatusResponse) GetPendingWarningsOk() (*[]string, bool)`

GetPendingWarningsOk returns a tuple with the PendingWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingWarnings

`func (o *GetSearchJobStatusResponse) SetPendingWarnings(v []string)`

SetPendingWarnings sets PendingWarnings field to given value.

### HasPendingWarnings

`func (o *GetSearchJobStatusResponse) HasPendingWarnings() bool`

HasPendingWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


