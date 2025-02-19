# SCIMPatchUserDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | Defines the SCIM schemas for the patch operation | 
**Operations** | [**[]SCIMPatchUserDefinitionOperationsInner**](SCIMPatchUserDefinitionOperationsInner.md) | Updates one or more attributes of a SCIM resource using a sequence of operations | 

## Methods

### NewSCIMPatchUserDefinition

`func NewSCIMPatchUserDefinition(schemas []string, operations []SCIMPatchUserDefinitionOperationsInner, ) *SCIMPatchUserDefinition`

NewSCIMPatchUserDefinition instantiates a new SCIMPatchUserDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMPatchUserDefinitionWithDefaults

`func NewSCIMPatchUserDefinitionWithDefaults() *SCIMPatchUserDefinition`

NewSCIMPatchUserDefinitionWithDefaults instantiates a new SCIMPatchUserDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SCIMPatchUserDefinition) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SCIMPatchUserDefinition) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SCIMPatchUserDefinition) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetOperations

`func (o *SCIMPatchUserDefinition) GetOperations() []SCIMPatchUserDefinitionOperationsInner`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *SCIMPatchUserDefinition) GetOperationsOk() (*[]SCIMPatchUserDefinitionOperationsInner, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *SCIMPatchUserDefinition) SetOperations(v []SCIMPatchUserDefinitionOperationsInner)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


