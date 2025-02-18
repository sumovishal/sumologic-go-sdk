# AsyncInstallAppJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Whether or not the request is in progress (&#x60;InProgress&#x60;), has completed successfully (&#x60;Success&#x60;), or has completed with an error (&#x60;Failed&#x60;). | 
**InstanceId** | Pointer to **string** | Instance identifier of the installed app. This field is not set yet but is a placeholder for future use. | [optional] 
**Path** | Pointer to **string** | Path of the folder in which the app was installed. | [optional] 
**FolderId** | Pointer to **string** | Identifier of the folder in which the app was installed. | [optional] 
**Error** | Pointer to [**ErrorDescription**](ErrorDescription.md) |  | [optional] 

## Methods

### NewAsyncInstallAppJobStatus

`func NewAsyncInstallAppJobStatus(status string, ) *AsyncInstallAppJobStatus`

NewAsyncInstallAppJobStatus instantiates a new AsyncInstallAppJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncInstallAppJobStatusWithDefaults

`func NewAsyncInstallAppJobStatusWithDefaults() *AsyncInstallAppJobStatus`

NewAsyncInstallAppJobStatusWithDefaults instantiates a new AsyncInstallAppJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AsyncInstallAppJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncInstallAppJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncInstallAppJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInstanceId

`func (o *AsyncInstallAppJobStatus) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AsyncInstallAppJobStatus) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AsyncInstallAppJobStatus) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AsyncInstallAppJobStatus) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetPath

`func (o *AsyncInstallAppJobStatus) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AsyncInstallAppJobStatus) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AsyncInstallAppJobStatus) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AsyncInstallAppJobStatus) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetFolderId

`func (o *AsyncInstallAppJobStatus) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AsyncInstallAppJobStatus) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AsyncInstallAppJobStatus) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AsyncInstallAppJobStatus) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetError

`func (o *AsyncInstallAppJobStatus) GetError() ErrorDescription`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AsyncInstallAppJobStatus) GetErrorOk() (*ErrorDescription, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AsyncInstallAppJobStatus) SetError(v ErrorDescription)`

SetError sets Error field to given value.

### HasError

`func (o *AsyncInstallAppJobStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


