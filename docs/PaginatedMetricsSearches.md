# PaginatedMetricsSearches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricsSearches** | [**[]MetricsSearchResponse**](MetricsSearchResponse.md) | List of metrics search pages. | 
**Next** | Pointer to **string** | Next continuation token. &#x60;token&#x60; is set to null when no more pages are left. | [optional] 

## Methods

### NewPaginatedMetricsSearches

`func NewPaginatedMetricsSearches(metricsSearches []MetricsSearchResponse, ) *PaginatedMetricsSearches`

NewPaginatedMetricsSearches instantiates a new PaginatedMetricsSearches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedMetricsSearchesWithDefaults

`func NewPaginatedMetricsSearchesWithDefaults() *PaginatedMetricsSearches`

NewPaginatedMetricsSearchesWithDefaults instantiates a new PaginatedMetricsSearches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricsSearches

`func (o *PaginatedMetricsSearches) GetMetricsSearches() []MetricsSearchResponse`

GetMetricsSearches returns the MetricsSearches field if non-nil, zero value otherwise.

### GetMetricsSearchesOk

`func (o *PaginatedMetricsSearches) GetMetricsSearchesOk() (*[]MetricsSearchResponse, bool)`

GetMetricsSearchesOk returns a tuple with the MetricsSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsSearches

`func (o *PaginatedMetricsSearches) SetMetricsSearches(v []MetricsSearchResponse)`

SetMetricsSearches sets MetricsSearches field to given value.


### GetNext

`func (o *PaginatedMetricsSearches) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedMetricsSearches) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedMetricsSearches) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedMetricsSearches) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


