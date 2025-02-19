# SCIMUserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | Defines the SCIM schemas for the user | 
**UserName** | **string** | Unique identifier for the user (email) | 
**Name** | [**NameInfo**](NameInfo.md) |  | 
**Emails** | [**[]SCIMCreateUserDefinitionEmailsInner**](SCIMCreateUserDefinitionEmailsInner.md) | Sumo logic accepts only one email address | 
**Roles** | [**[]SCIMCreateUserDefinitionRolesInner**](SCIMCreateUserDefinitionRolesInner.md) | Roles should exactly match with role names within sumologic | 
**Id** | **string** | Unique SCIM identifier for the user | 
**Active** | Pointer to **bool** | True if the user is active | [optional] 
**Meta** | Pointer to [**ResourceData**](ResourceData.md) |  | [optional] 

## Methods

### NewSCIMUserModel

`func NewSCIMUserModel(schemas []string, userName string, name NameInfo, emails []SCIMCreateUserDefinitionEmailsInner, roles []SCIMCreateUserDefinitionRolesInner, id string, ) *SCIMUserModel`

NewSCIMUserModel instantiates a new SCIMUserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMUserModelWithDefaults

`func NewSCIMUserModelWithDefaults() *SCIMUserModel`

NewSCIMUserModelWithDefaults instantiates a new SCIMUserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SCIMUserModel) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SCIMUserModel) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SCIMUserModel) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetUserName

`func (o *SCIMUserModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SCIMUserModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SCIMUserModel) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *SCIMUserModel) GetName() NameInfo`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCIMUserModel) GetNameOk() (*NameInfo, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCIMUserModel) SetName(v NameInfo)`

SetName sets Name field to given value.


### GetEmails

`func (o *SCIMUserModel) GetEmails() []SCIMCreateUserDefinitionEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SCIMUserModel) GetEmailsOk() (*[]SCIMCreateUserDefinitionEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SCIMUserModel) SetEmails(v []SCIMCreateUserDefinitionEmailsInner)`

SetEmails sets Emails field to given value.


### GetRoles

`func (o *SCIMUserModel) GetRoles() []SCIMCreateUserDefinitionRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SCIMUserModel) GetRolesOk() (*[]SCIMCreateUserDefinitionRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SCIMUserModel) SetRoles(v []SCIMCreateUserDefinitionRolesInner)`

SetRoles sets Roles field to given value.


### GetId

`func (o *SCIMUserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SCIMUserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SCIMUserModel) SetId(v string)`

SetId sets Id field to given value.


### GetActive

`func (o *SCIMUserModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SCIMUserModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SCIMUserModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SCIMUserModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMeta

`func (o *SCIMUserModel) GetMeta() ResourceData`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SCIMUserModel) GetMetaOk() (*ResourceData, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SCIMUserModel) SetMeta(v ResourceData)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SCIMUserModel) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


