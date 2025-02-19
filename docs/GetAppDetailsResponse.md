# GetAppDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | UUID of the app. | 
**Version** | **string** | Version of the app. | 
**BaseUrl** | **string** | URL prefix for where the app is stored. | 
**Manifest** | **string** | Content of the manifest YAML file, as Base64-encoded string. | 
**Config** | Pointer to **string** | Content of the config YAML file, as Base64-encoded string. | [optional] 
**Readme** | Pointer to **string** | Content of the README markdown file, as Base64-encoded string. | [optional] 
**Files** | Pointer to **map[string]string** | Content of various files part of app package, as Base64-encoded string. | [optional] 

## Methods

### NewGetAppDetailsResponse

`func NewGetAppDetailsResponse(uuid string, version string, baseUrl string, manifest string, ) *GetAppDetailsResponse`

NewGetAppDetailsResponse instantiates a new GetAppDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppDetailsResponseWithDefaults

`func NewGetAppDetailsResponseWithDefaults() *GetAppDetailsResponse`

NewGetAppDetailsResponseWithDefaults instantiates a new GetAppDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *GetAppDetailsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetAppDetailsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetAppDetailsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetVersion

`func (o *GetAppDetailsResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetAppDetailsResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetAppDetailsResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBaseUrl

`func (o *GetAppDetailsResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *GetAppDetailsResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *GetAppDetailsResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetManifest

`func (o *GetAppDetailsResponse) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *GetAppDetailsResponse) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *GetAppDetailsResponse) SetManifest(v string)`

SetManifest sets Manifest field to given value.


### GetConfig

`func (o *GetAppDetailsResponse) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetAppDetailsResponse) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetAppDetailsResponse) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetAppDetailsResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetReadme

`func (o *GetAppDetailsResponse) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *GetAppDetailsResponse) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *GetAppDetailsResponse) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *GetAppDetailsResponse) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### GetFiles

`func (o *GetAppDetailsResponse) GetFiles() map[string]string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetAppDetailsResponse) GetFilesOk() (*map[string]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetAppDetailsResponse) SetFiles(v map[string]string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GetAppDetailsResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


