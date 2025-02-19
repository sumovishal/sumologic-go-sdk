# SCIMPatchUserDefinitionOperationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | We will only support &#x60;replace&#x60; operation. | [optional] 
**Path** | Pointer to **string** | Attribute path to modify | [optional] 
**Value** | Pointer to [**SCIMPatchUserDefinitionOperationsInnerValue**](SCIMPatchUserDefinitionOperationsInnerValue.md) |  | [optional] 

## Methods

### NewSCIMPatchUserDefinitionOperationsInner

`func NewSCIMPatchUserDefinitionOperationsInner() *SCIMPatchUserDefinitionOperationsInner`

NewSCIMPatchUserDefinitionOperationsInner instantiates a new SCIMPatchUserDefinitionOperationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMPatchUserDefinitionOperationsInnerWithDefaults

`func NewSCIMPatchUserDefinitionOperationsInnerWithDefaults() *SCIMPatchUserDefinitionOperationsInner`

NewSCIMPatchUserDefinitionOperationsInnerWithDefaults instantiates a new SCIMPatchUserDefinitionOperationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *SCIMPatchUserDefinitionOperationsInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SCIMPatchUserDefinitionOperationsInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SCIMPatchUserDefinitionOperationsInner) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SCIMPatchUserDefinitionOperationsInner) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *SCIMPatchUserDefinitionOperationsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SCIMPatchUserDefinitionOperationsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SCIMPatchUserDefinitionOperationsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SCIMPatchUserDefinitionOperationsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *SCIMPatchUserDefinitionOperationsInner) GetValue() SCIMPatchUserDefinitionOperationsInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SCIMPatchUserDefinitionOperationsInner) GetValueOk() (*SCIMPatchUserDefinitionOperationsInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SCIMPatchUserDefinitionOperationsInner) SetValue(v SCIMPatchUserDefinitionOperationsInnerValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *SCIMPatchUserDefinitionOperationsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


