# ServiceAccountModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the service account. | 
**IsActive** | Pointer to **bool** | True if the service account is active. | [optional] 

## Methods

### NewServiceAccountModelAllOf

`func NewServiceAccountModelAllOf(id string, ) *ServiceAccountModelAllOf`

NewServiceAccountModelAllOf instantiates a new ServiceAccountModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountModelAllOfWithDefaults

`func NewServiceAccountModelAllOfWithDefaults() *ServiceAccountModelAllOf`

NewServiceAccountModelAllOfWithDefaults instantiates a new ServiceAccountModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAccountModelAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountModelAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountModelAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *ServiceAccountModelAllOf) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ServiceAccountModelAllOf) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ServiceAccountModelAllOf) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ServiceAccountModelAllOf) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


