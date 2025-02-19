# AccessKeysLifetimePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeysLifetimeInDays** | **int32** | The number of days it will take for an access key to expire without being rotated/copied. Setting it to 0 (never) means that access keys will never expire. | 

## Methods

### NewAccessKeysLifetimePolicy

`func NewAccessKeysLifetimePolicy(accessKeysLifetimeInDays int32, ) *AccessKeysLifetimePolicy`

NewAccessKeysLifetimePolicy instantiates a new AccessKeysLifetimePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeysLifetimePolicyWithDefaults

`func NewAccessKeysLifetimePolicyWithDefaults() *AccessKeysLifetimePolicy`

NewAccessKeysLifetimePolicyWithDefaults instantiates a new AccessKeysLifetimePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeysLifetimeInDays

`func (o *AccessKeysLifetimePolicy) GetAccessKeysLifetimeInDays() int32`

GetAccessKeysLifetimeInDays returns the AccessKeysLifetimeInDays field if non-nil, zero value otherwise.

### GetAccessKeysLifetimeInDaysOk

`func (o *AccessKeysLifetimePolicy) GetAccessKeysLifetimeInDaysOk() (*int32, bool)`

GetAccessKeysLifetimeInDaysOk returns a tuple with the AccessKeysLifetimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeysLifetimeInDays

`func (o *AccessKeysLifetimePolicy) SetAccessKeysLifetimeInDays(v int32)`

SetAccessKeysLifetimeInDays sets AccessKeysLifetimeInDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


