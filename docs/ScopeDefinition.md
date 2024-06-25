# ScopeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The name of the scope. | 
**Label** | **string** | The UI label for the scope. | 
**Group** | [**ScopeDefinitionGroup**](ScopeDefinitionGroup.md) |  | 

## Methods

### NewScopeDefinition

`func NewScopeDefinition(id string, label string, group ScopeDefinitionGroup, ) *ScopeDefinition`

NewScopeDefinition instantiates a new ScopeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeDefinitionWithDefaults

`func NewScopeDefinitionWithDefaults() *ScopeDefinition`

NewScopeDefinitionWithDefaults instantiates a new ScopeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScopeDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopeDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopeDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *ScopeDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ScopeDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ScopeDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetGroup

`func (o *ScopeDefinition) GetGroup() ScopeDefinitionGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ScopeDefinition) GetGroupOk() (*ScopeDefinitionGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ScopeDefinition) SetGroup(v ScopeDefinitionGroup)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


