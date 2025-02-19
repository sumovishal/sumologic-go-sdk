# SCIMCreateUserDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | Defines the SCIM schemas for the user | 
**UserName** | **string** | Unique identifier for the user (email) | 
**Name** | [**NameInfo**](NameInfo.md) |  | 
**Emails** | [**[]SCIMCreateUserDefinitionEmailsInner**](SCIMCreateUserDefinitionEmailsInner.md) | Sumo logic accepts only one email address | 
**Roles** | [**[]SCIMCreateUserDefinitionRolesInner**](SCIMCreateUserDefinitionRolesInner.md) | Roles should exactly match with role names within sumologic | 

## Methods

### NewSCIMCreateUserDefinition

`func NewSCIMCreateUserDefinition(schemas []string, userName string, name NameInfo, emails []SCIMCreateUserDefinitionEmailsInner, roles []SCIMCreateUserDefinitionRolesInner, ) *SCIMCreateUserDefinition`

NewSCIMCreateUserDefinition instantiates a new SCIMCreateUserDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMCreateUserDefinitionWithDefaults

`func NewSCIMCreateUserDefinitionWithDefaults() *SCIMCreateUserDefinition`

NewSCIMCreateUserDefinitionWithDefaults instantiates a new SCIMCreateUserDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SCIMCreateUserDefinition) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SCIMCreateUserDefinition) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SCIMCreateUserDefinition) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetUserName

`func (o *SCIMCreateUserDefinition) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SCIMCreateUserDefinition) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SCIMCreateUserDefinition) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *SCIMCreateUserDefinition) GetName() NameInfo`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCIMCreateUserDefinition) GetNameOk() (*NameInfo, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCIMCreateUserDefinition) SetName(v NameInfo)`

SetName sets Name field to given value.


### GetEmails

`func (o *SCIMCreateUserDefinition) GetEmails() []SCIMCreateUserDefinitionEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SCIMCreateUserDefinition) GetEmailsOk() (*[]SCIMCreateUserDefinitionEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SCIMCreateUserDefinition) SetEmails(v []SCIMCreateUserDefinitionEmailsInner)`

SetEmails sets Emails field to given value.


### GetRoles

`func (o *SCIMCreateUserDefinition) GetRoles() []SCIMCreateUserDefinitionRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SCIMCreateUserDefinition) GetRolesOk() (*[]SCIMCreateUserDefinitionRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SCIMCreateUserDefinition) SetRoles(v []SCIMCreateUserDefinitionRolesInner)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


