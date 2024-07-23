# ParsersLibraryBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the folder or parser. | 
**Description** | **string** | Description of the folder or parser. | 
**Type** | **string** | Type of the object model. | 
**IsLocked** | Pointer to **bool** | Locking/Unlocking requires the &#x60;LockParsers&#x60; capability. Locked objects can only be &#x60;Localized&#x60;. Updating or moving requires unlocking the object. Locking/Unlocking recursively locks all of the objects children. All children of a locked object must be locked. | [optional] [default to false]

## Methods

### NewParsersLibraryBase

`func NewParsersLibraryBase(name string, description string, type_ string, ) *ParsersLibraryBase`

NewParsersLibraryBase instantiates a new ParsersLibraryBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsersLibraryBaseWithDefaults

`func NewParsersLibraryBaseWithDefaults() *ParsersLibraryBase`

NewParsersLibraryBaseWithDefaults instantiates a new ParsersLibraryBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParsersLibraryBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParsersLibraryBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParsersLibraryBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ParsersLibraryBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParsersLibraryBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParsersLibraryBase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ParsersLibraryBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParsersLibraryBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParsersLibraryBase) SetType(v string)`

SetType sets Type field to given value.


### GetIsLocked

`func (o *ParsersLibraryBase) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ParsersLibraryBase) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ParsersLibraryBase) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ParsersLibraryBase) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


