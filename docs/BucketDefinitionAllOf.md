# BucketDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the data forwarding destination. | 
**InvalidatedBySystem** | Pointer to **bool** | True if invalidated by the system. | [optional] 

## Methods

### NewBucketDefinitionAllOf

`func NewBucketDefinitionAllOf(id string, ) *BucketDefinitionAllOf`

NewBucketDefinitionAllOf instantiates a new BucketDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketDefinitionAllOfWithDefaults

`func NewBucketDefinitionAllOfWithDefaults() *BucketDefinitionAllOf`

NewBucketDefinitionAllOfWithDefaults instantiates a new BucketDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BucketDefinitionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketDefinitionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketDefinitionAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetInvalidatedBySystem

`func (o *BucketDefinitionAllOf) GetInvalidatedBySystem() bool`

GetInvalidatedBySystem returns the InvalidatedBySystem field if non-nil, zero value otherwise.

### GetInvalidatedBySystemOk

`func (o *BucketDefinitionAllOf) GetInvalidatedBySystemOk() (*bool, bool)`

GetInvalidatedBySystemOk returns a tuple with the InvalidatedBySystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidatedBySystem

`func (o *BucketDefinitionAllOf) SetInvalidatedBySystem(v bool)`

SetInvalidatedBySystem sets InvalidatedBySystem field to given value.

### HasInvalidatedBySystem

`func (o *BucketDefinitionAllOf) HasInvalidatedBySystem() bool`

HasInvalidatedBySystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


