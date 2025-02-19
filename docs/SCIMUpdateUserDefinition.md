# SCIMUpdateUserDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | Defines the SCIM schemas for the user | 
**Name** | [**NameInfo**](NameInfo.md) |  | 
**Active** | **bool** | Indicates whether the user is active. | 
**Roles** | [**[]SCIMCreateUserDefinitionRolesInner**](SCIMCreateUserDefinitionRolesInner.md) | Roles should exactly match with role names within sumologic | 

## Methods

### NewSCIMUpdateUserDefinition

`func NewSCIMUpdateUserDefinition(schemas []string, name NameInfo, active bool, roles []SCIMCreateUserDefinitionRolesInner, ) *SCIMUpdateUserDefinition`

NewSCIMUpdateUserDefinition instantiates a new SCIMUpdateUserDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMUpdateUserDefinitionWithDefaults

`func NewSCIMUpdateUserDefinitionWithDefaults() *SCIMUpdateUserDefinition`

NewSCIMUpdateUserDefinitionWithDefaults instantiates a new SCIMUpdateUserDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SCIMUpdateUserDefinition) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SCIMUpdateUserDefinition) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SCIMUpdateUserDefinition) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetName

`func (o *SCIMUpdateUserDefinition) GetName() NameInfo`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCIMUpdateUserDefinition) GetNameOk() (*NameInfo, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCIMUpdateUserDefinition) SetName(v NameInfo)`

SetName sets Name field to given value.


### GetActive

`func (o *SCIMUpdateUserDefinition) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SCIMUpdateUserDefinition) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SCIMUpdateUserDefinition) SetActive(v bool)`

SetActive sets Active field to given value.


### GetRoles

`func (o *SCIMUpdateUserDefinition) GetRoles() []SCIMCreateUserDefinitionRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SCIMUpdateUserDefinition) GetRolesOk() (*[]SCIMCreateUserDefinitionRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SCIMUpdateUserDefinition) SetRoles(v []SCIMCreateUserDefinitionRolesInner)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


