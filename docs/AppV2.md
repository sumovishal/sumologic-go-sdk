# AppV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | UUID of the app. | 
**Name** | **string** | Name of the app. | 
**Description** | **string** | Description of the app. | 
**LatestVersion** | **string** | Latest version of the app. | 
**Icon** | **string** | URL of the icon for the app. | 
**Author** | **string** | Author of the app. | 
**AccountTypes** | **[]string** | Account types of which the app is available to. | 
**Beta** | **bool** | Whether the app is in beta. | 
**Installs** | Pointer to **int32** | Number of times the app was installed. | [optional] 
**Attributes** | **map[string][]string** | A map of attributes for this app. Attributes allow to group apps based on different criteria. | 
**Installable** | **bool** | Whether the app is installable or not as not all apps are installable. | 
**ShowOnMarketplace** | **bool** | Whether the app should show up on sumologic.com/applications webpage. | 
**ModifiedAt** | Pointer to **time.Time** | The timestamp in UTC of the most recent modification of the app. | [optional] 

## Methods

### NewAppV2

`func NewAppV2(uuid string, name string, description string, latestVersion string, icon string, author string, accountTypes []string, beta bool, attributes map[string][]string, installable bool, showOnMarketplace bool, ) *AppV2`

NewAppV2 instantiates a new AppV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppV2WithDefaults

`func NewAppV2WithDefaults() *AppV2`

NewAppV2WithDefaults instantiates a new AppV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AppV2) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppV2) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppV2) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AppV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppV2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLatestVersion

`func (o *AppV2) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *AppV2) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *AppV2) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### GetIcon

`func (o *AppV2) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AppV2) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AppV2) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetAuthor

`func (o *AppV2) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AppV2) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AppV2) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetAccountTypes

`func (o *AppV2) GetAccountTypes() []string`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *AppV2) GetAccountTypesOk() (*[]string, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *AppV2) SetAccountTypes(v []string)`

SetAccountTypes sets AccountTypes field to given value.


### GetBeta

`func (o *AppV2) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *AppV2) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *AppV2) SetBeta(v bool)`

SetBeta sets Beta field to given value.


### GetInstalls

`func (o *AppV2) GetInstalls() int32`

GetInstalls returns the Installs field if non-nil, zero value otherwise.

### GetInstallsOk

`func (o *AppV2) GetInstallsOk() (*int32, bool)`

GetInstallsOk returns a tuple with the Installs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalls

`func (o *AppV2) SetInstalls(v int32)`

SetInstalls sets Installs field to given value.

### HasInstalls

`func (o *AppV2) HasInstalls() bool`

HasInstalls returns a boolean if a field has been set.

### GetAttributes

`func (o *AppV2) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppV2) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppV2) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.


### GetInstallable

`func (o *AppV2) GetInstallable() bool`

GetInstallable returns the Installable field if non-nil, zero value otherwise.

### GetInstallableOk

`func (o *AppV2) GetInstallableOk() (*bool, bool)`

GetInstallableOk returns a tuple with the Installable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallable

`func (o *AppV2) SetInstallable(v bool)`

SetInstallable sets Installable field to given value.


### GetShowOnMarketplace

`func (o *AppV2) GetShowOnMarketplace() bool`

GetShowOnMarketplace returns the ShowOnMarketplace field if non-nil, zero value otherwise.

### GetShowOnMarketplaceOk

`func (o *AppV2) GetShowOnMarketplaceOk() (*bool, bool)`

GetShowOnMarketplaceOk returns a tuple with the ShowOnMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnMarketplace

`func (o *AppV2) SetShowOnMarketplace(v bool)`

SetShowOnMarketplace sets ShowOnMarketplace field to given value.


### GetModifiedAt

`func (o *AppV2) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AppV2) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AppV2) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AppV2) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


