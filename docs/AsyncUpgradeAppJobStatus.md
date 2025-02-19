# AsyncUpgradeAppJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Whether or not the request is in progress (&#x60;InProgress&#x60;), has completed successfully (&#x60;Success&#x60;), or has completed with an error (&#x60;Failed&#x60;). | 
**InstanceId** | Pointer to **string** | Instance identifier of the upgraded app. This field is not set yet but is a placeholder for future use. | [optional] 
**Path** | Pointer to **string** | Path of the folder in which the app was upgraded. | [optional] 
**FolderId** | Pointer to **string** | Identifier of the folder in which the app was upgraded. | [optional] 
**Error** | Pointer to [**ErrorDescription**](ErrorDescription.md) |  | [optional] 

## Methods

### NewAsyncUpgradeAppJobStatus

`func NewAsyncUpgradeAppJobStatus(status string, ) *AsyncUpgradeAppJobStatus`

NewAsyncUpgradeAppJobStatus instantiates a new AsyncUpgradeAppJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncUpgradeAppJobStatusWithDefaults

`func NewAsyncUpgradeAppJobStatusWithDefaults() *AsyncUpgradeAppJobStatus`

NewAsyncUpgradeAppJobStatusWithDefaults instantiates a new AsyncUpgradeAppJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AsyncUpgradeAppJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncUpgradeAppJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncUpgradeAppJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInstanceId

`func (o *AsyncUpgradeAppJobStatus) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AsyncUpgradeAppJobStatus) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AsyncUpgradeAppJobStatus) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AsyncUpgradeAppJobStatus) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetPath

`func (o *AsyncUpgradeAppJobStatus) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AsyncUpgradeAppJobStatus) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AsyncUpgradeAppJobStatus) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AsyncUpgradeAppJobStatus) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetFolderId

`func (o *AsyncUpgradeAppJobStatus) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AsyncUpgradeAppJobStatus) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AsyncUpgradeAppJobStatus) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AsyncUpgradeAppJobStatus) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetError

`func (o *AsyncUpgradeAppJobStatus) GetError() ErrorDescription`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AsyncUpgradeAppJobStatus) GetErrorOk() (*ErrorDescription, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AsyncUpgradeAppJobStatus) SetError(v ErrorDescription)`

SetError sets Error field to given value.

### HasError

`func (o *AsyncUpgradeAppJobStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


