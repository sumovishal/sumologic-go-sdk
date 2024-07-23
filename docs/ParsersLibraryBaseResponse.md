# ParsersLibraryBaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the folder or parser. | 
**Name** | **string** | Name of the folder or parser. | 
**Description** | **string** | Description of the folder or parser. | 
**Version** | **int64** | Version of the folder or parser. | 
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.  | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**ParentId** | **string** | Identifier of the parent folder. | 
**ContentType** | **string** | Type of the content. Valid values:   1) Folder   2) Parser  | 
**Type** | **string** | Type of the object model. | 
**IsLocked** | **bool** | Whether the object is locked. | 
**IsSystem** | **bool** | System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can&#39;t be updated. | 
**IsMutable** | **bool** | Immutable objects are \&quot;READ-ONLY\&quot;. | 

## Methods

### NewParsersLibraryBaseResponse

`func NewParsersLibraryBaseResponse(id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isLocked bool, isSystem bool, isMutable bool, ) *ParsersLibraryBaseResponse`

NewParsersLibraryBaseResponse instantiates a new ParsersLibraryBaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsersLibraryBaseResponseWithDefaults

`func NewParsersLibraryBaseResponseWithDefaults() *ParsersLibraryBaseResponse`

NewParsersLibraryBaseResponseWithDefaults instantiates a new ParsersLibraryBaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParsersLibraryBaseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParsersLibraryBaseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParsersLibraryBaseResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ParsersLibraryBaseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParsersLibraryBaseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParsersLibraryBaseResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ParsersLibraryBaseResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParsersLibraryBaseResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParsersLibraryBaseResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *ParsersLibraryBaseResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ParsersLibraryBaseResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ParsersLibraryBaseResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *ParsersLibraryBaseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParsersLibraryBaseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParsersLibraryBaseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ParsersLibraryBaseResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ParsersLibraryBaseResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ParsersLibraryBaseResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *ParsersLibraryBaseResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParsersLibraryBaseResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParsersLibraryBaseResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ParsersLibraryBaseResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ParsersLibraryBaseResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ParsersLibraryBaseResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetParentId

`func (o *ParsersLibraryBaseResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ParsersLibraryBaseResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ParsersLibraryBaseResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetContentType

`func (o *ParsersLibraryBaseResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ParsersLibraryBaseResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ParsersLibraryBaseResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetType

`func (o *ParsersLibraryBaseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParsersLibraryBaseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParsersLibraryBaseResponse) SetType(v string)`

SetType sets Type field to given value.


### GetIsLocked

`func (o *ParsersLibraryBaseResponse) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ParsersLibraryBaseResponse) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ParsersLibraryBaseResponse) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsSystem

`func (o *ParsersLibraryBaseResponse) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *ParsersLibraryBaseResponse) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *ParsersLibraryBaseResponse) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetIsMutable

`func (o *ParsersLibraryBaseResponse) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *ParsersLibraryBaseResponse) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *ParsersLibraryBaseResponse) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


