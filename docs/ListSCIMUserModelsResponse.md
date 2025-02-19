# ListSCIMUserModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResults** | Pointer to **int32** | Total number of users that match the filter criteria | [optional] 
**StartIndex** | Pointer to **int32** | The index of the first returned result | [optional] 
**ItemsPerPage** | Pointer to **int32** | The number of results returned in this page | [optional] 
**Resources** | Pointer to [**[]SCIMUserModel**](SCIMUserModel.md) | List of SCIM user resources | [optional] 

## Methods

### NewListSCIMUserModelsResponse

`func NewListSCIMUserModelsResponse() *ListSCIMUserModelsResponse`

NewListSCIMUserModelsResponse instantiates a new ListSCIMUserModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSCIMUserModelsResponseWithDefaults

`func NewListSCIMUserModelsResponseWithDefaults() *ListSCIMUserModelsResponse`

NewListSCIMUserModelsResponseWithDefaults instantiates a new ListSCIMUserModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResults

`func (o *ListSCIMUserModelsResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ListSCIMUserModelsResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ListSCIMUserModelsResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ListSCIMUserModelsResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetStartIndex

`func (o *ListSCIMUserModelsResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ListSCIMUserModelsResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ListSCIMUserModelsResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ListSCIMUserModelsResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *ListSCIMUserModelsResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ListSCIMUserModelsResponse) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ListSCIMUserModelsResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ListSCIMUserModelsResponse) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetResources

`func (o *ListSCIMUserModelsResponse) GetResources() []SCIMUserModel`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ListSCIMUserModelsResponse) GetResourcesOk() (*[]SCIMUserModel, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ListSCIMUserModelsResponse) SetResources(v []SCIMUserModel)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ListSCIMUserModelsResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


