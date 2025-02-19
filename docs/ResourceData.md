# ResourceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | The name of the resource type of the resource. | [optional] 
**Created** | Pointer to **time.Time** | Creation timestamp in date-time format. | [optional] 
**LastModified** | Pointer to **time.Time** | Last modification timestamp in date-time format. | [optional] 

## Methods

### NewResourceData

`func NewResourceData() *ResourceData`

NewResourceData instantiates a new ResourceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDataWithDefaults

`func NewResourceDataWithDefaults() *ResourceData`

NewResourceDataWithDefaults instantiates a new ResourceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *ResourceData) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceData) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceData) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceData) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetCreated

`func (o *ResourceData) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResourceData) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResourceData) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResourceData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *ResourceData) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResourceData) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResourceData) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResourceData) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


