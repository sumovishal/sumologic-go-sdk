# Selector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[][]OtTag**]([]OtTag.md) | tags filter for agents | [optional] 
**Names** | Pointer to **[]string** | names to select custom agents | [optional] 

## Methods

### NewSelector

`func NewSelector() *Selector`

NewSelector instantiates a new Selector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectorWithDefaults

`func NewSelectorWithDefaults() *Selector`

NewSelectorWithDefaults instantiates a new Selector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *Selector) GetTags() [][]OtTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Selector) GetTagsOk() (*[][]OtTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Selector) SetTags(v [][]OtTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Selector) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNames

`func (o *Selector) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Selector) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Selector) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *Selector) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


