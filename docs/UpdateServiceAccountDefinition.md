# UpdateServiceAccountDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the service account. | [optional] 
**IsActive** | Pointer to **bool** | This has the value &#x60;true&#x60; if the service account is active and &#x60;false&#x60; if it has been deactivated. | [optional] 
**RoleIds** | Pointer to **[]string** | List of role identifiers associated with the service account. | [optional] 
**Email** | Pointer to **string** | New email address of the service account. | [optional] 

## Methods

### NewUpdateServiceAccountDefinition

`func NewUpdateServiceAccountDefinition() *UpdateServiceAccountDefinition`

NewUpdateServiceAccountDefinition instantiates a new UpdateServiceAccountDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceAccountDefinitionWithDefaults

`func NewUpdateServiceAccountDefinitionWithDefaults() *UpdateServiceAccountDefinition`

NewUpdateServiceAccountDefinitionWithDefaults instantiates a new UpdateServiceAccountDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateServiceAccountDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServiceAccountDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServiceAccountDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServiceAccountDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateServiceAccountDefinition) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateServiceAccountDefinition) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateServiceAccountDefinition) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateServiceAccountDefinition) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRoleIds

`func (o *UpdateServiceAccountDefinition) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UpdateServiceAccountDefinition) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UpdateServiceAccountDefinition) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *UpdateServiceAccountDefinition) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateServiceAccountDefinition) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateServiceAccountDefinition) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateServiceAccountDefinition) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateServiceAccountDefinition) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


