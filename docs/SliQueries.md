# SliQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryGroup** | [**SliQueryGroup**](SliQueryGroup.md) |  | 
**QueryType** | **string** | Type of queries for SLI (Logs/Metrics). | 

## Methods

### NewSliQueries

`func NewSliQueries(queryGroup SliQueryGroup, queryType string, ) *SliQueries`

NewSliQueries instantiates a new SliQueries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliQueriesWithDefaults

`func NewSliQueriesWithDefaults() *SliQueries`

NewSliQueriesWithDefaults instantiates a new SliQueries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryGroup

`func (o *SliQueries) GetQueryGroup() SliQueryGroup`

GetQueryGroup returns the QueryGroup field if non-nil, zero value otherwise.

### GetQueryGroupOk

`func (o *SliQueries) GetQueryGroupOk() (*SliQueryGroup, bool)`

GetQueryGroupOk returns a tuple with the QueryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryGroup

`func (o *SliQueries) SetQueryGroup(v SliQueryGroup)`

SetQueryGroup sets QueryGroup field to given value.


### GetQueryType

`func (o *SliQueries) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *SliQueries) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *SliQueries) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


