# AsyncUninstallAppJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Whether or not the request is in progress (&#x60;InProgress&#x60;), has completed successfully (&#x60;Success&#x60;), or has completed with an error (&#x60;Failed&#x60;). | 
**Errors** | Pointer to [**[]ErrorDescription**](ErrorDescription.md) | More information about the failure if the status is &#x60;Failed&#x60;. | [optional] 

## Methods

### NewAsyncUninstallAppJobStatus

`func NewAsyncUninstallAppJobStatus(status string, ) *AsyncUninstallAppJobStatus`

NewAsyncUninstallAppJobStatus instantiates a new AsyncUninstallAppJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncUninstallAppJobStatusWithDefaults

`func NewAsyncUninstallAppJobStatusWithDefaults() *AsyncUninstallAppJobStatus`

NewAsyncUninstallAppJobStatusWithDefaults instantiates a new AsyncUninstallAppJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AsyncUninstallAppJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncUninstallAppJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncUninstallAppJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *AsyncUninstallAppJobStatus) GetErrors() []ErrorDescription`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AsyncUninstallAppJobStatus) GetErrorsOk() (*[]ErrorDescription, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AsyncUninstallAppJobStatus) SetErrors(v []ErrorDescription)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AsyncUninstallAppJobStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


