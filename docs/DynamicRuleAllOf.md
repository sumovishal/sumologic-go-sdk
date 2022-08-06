# DynamicRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the dynamic parsing rule. | 
**IsSystemRule** | **bool** | Whether the rule has been defined by the system, rather than by a user. | 

## Methods

### NewDynamicRuleAllOf

`func NewDynamicRuleAllOf(id string, isSystemRule bool, ) *DynamicRuleAllOf`

NewDynamicRuleAllOf instantiates a new DynamicRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicRuleAllOfWithDefaults

`func NewDynamicRuleAllOfWithDefaults() *DynamicRuleAllOf`

NewDynamicRuleAllOfWithDefaults instantiates a new DynamicRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DynamicRuleAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicRuleAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicRuleAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetIsSystemRule

`func (o *DynamicRuleAllOf) GetIsSystemRule() bool`

GetIsSystemRule returns the IsSystemRule field if non-nil, zero value otherwise.

### GetIsSystemRuleOk

`func (o *DynamicRuleAllOf) GetIsSystemRuleOk() (*bool, bool)`

GetIsSystemRuleOk returns a tuple with the IsSystemRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemRule

`func (o *DynamicRuleAllOf) SetIsSystemRule(v bool)`

SetIsSystemRule sets IsSystemRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


