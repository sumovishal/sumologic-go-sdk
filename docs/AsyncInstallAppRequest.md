# AsyncInstallAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the app to install. You can either specify a specific version of the app or use &#x60;latest&#x60; to install the latest version of the app. _If version is not specified, the latest version of the app will be installed_.  | [optional] [default to "latest"]
**Parameters** | Pointer to **map[string]string** | Map of additional parameters for the app installation. | [optional] 

## Methods

### NewAsyncInstallAppRequest

`func NewAsyncInstallAppRequest() *AsyncInstallAppRequest`

NewAsyncInstallAppRequest instantiates a new AsyncInstallAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncInstallAppRequestWithDefaults

`func NewAsyncInstallAppRequestWithDefaults() *AsyncInstallAppRequest`

NewAsyncInstallAppRequestWithDefaults instantiates a new AsyncInstallAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *AsyncInstallAppRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AsyncInstallAppRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AsyncInstallAppRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AsyncInstallAppRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetParameters

`func (o *AsyncInstallAppRequest) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AsyncInstallAppRequest) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AsyncInstallAppRequest) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AsyncInstallAppRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


