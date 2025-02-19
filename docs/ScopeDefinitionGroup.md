# ScopeDefinitionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The name of the scope group | 
**Label** | **string** | The label for the scope group | 
**ParentId** | Pointer to **string** | The ID of the parent scope group | [optional] 

## Methods

### NewScopeDefinitionGroup

`func NewScopeDefinitionGroup(id string, label string, ) *ScopeDefinitionGroup`

NewScopeDefinitionGroup instantiates a new ScopeDefinitionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeDefinitionGroupWithDefaults

`func NewScopeDefinitionGroupWithDefaults() *ScopeDefinitionGroup`

NewScopeDefinitionGroupWithDefaults instantiates a new ScopeDefinitionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScopeDefinitionGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopeDefinitionGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopeDefinitionGroup) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *ScopeDefinitionGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ScopeDefinitionGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ScopeDefinitionGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetParentId

`func (o *ScopeDefinitionGroup) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ScopeDefinitionGroup) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ScopeDefinitionGroup) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ScopeDefinitionGroup) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


