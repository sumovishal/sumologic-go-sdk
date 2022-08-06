# TracesListPanelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | Pointer to [**[]Query**](Query.md) | Traces queries of the panel. | [optional] 
**TimeRange** | Pointer to [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] 

## Methods

### NewTracesListPanelAllOf

`func NewTracesListPanelAllOf() *TracesListPanelAllOf`

NewTracesListPanelAllOf instantiates a new TracesListPanelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracesListPanelAllOfWithDefaults

`func NewTracesListPanelAllOfWithDefaults() *TracesListPanelAllOf`

NewTracesListPanelAllOfWithDefaults instantiates a new TracesListPanelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *TracesListPanelAllOf) GetQueries() []Query`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *TracesListPanelAllOf) GetQueriesOk() (*[]Query, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *TracesListPanelAllOf) SetQueries(v []Query)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *TracesListPanelAllOf) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetTimeRange

`func (o *TracesListPanelAllOf) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *TracesListPanelAllOf) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *TracesListPanelAllOf) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *TracesListPanelAllOf) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


