# CreateServiceAccountDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the service account. | 
**Email** | **string** | Email address of the service account. | 
**RoleIds** | **[]string** | List of roleIds associated with the service account. | 

## Methods

### NewCreateServiceAccountDefinition

`func NewCreateServiceAccountDefinition(name string, email string, roleIds []string, ) *CreateServiceAccountDefinition`

NewCreateServiceAccountDefinition instantiates a new CreateServiceAccountDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountDefinitionWithDefaults

`func NewCreateServiceAccountDefinitionWithDefaults() *CreateServiceAccountDefinition`

NewCreateServiceAccountDefinitionWithDefaults instantiates a new CreateServiceAccountDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateServiceAccountDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceAccountDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceAccountDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateServiceAccountDefinition) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateServiceAccountDefinition) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateServiceAccountDefinition) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoleIds

`func (o *CreateServiceAccountDefinition) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *CreateServiceAccountDefinition) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *CreateServiceAccountDefinition) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


