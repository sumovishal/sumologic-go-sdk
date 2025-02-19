# ContentSyncStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Content Sync Job status. | 
**Progress** | **int32** | Content Sync Job progress percentage. | 

## Methods

### NewContentSyncStatusResponse

`func NewContentSyncStatusResponse(status string, progress int32, ) *ContentSyncStatusResponse`

NewContentSyncStatusResponse instantiates a new ContentSyncStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSyncStatusResponseWithDefaults

`func NewContentSyncStatusResponseWithDefaults() *ContentSyncStatusResponse`

NewContentSyncStatusResponseWithDefaults instantiates a new ContentSyncStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ContentSyncStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContentSyncStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContentSyncStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *ContentSyncStatusResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ContentSyncStatusResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ContentSyncStatusResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


