# MetricsSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the metrics search page. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**Description** | Pointer to **string** | Description of the metrics search page. | [optional] 
**Queries** | [**[]Query**](Query.md) | Queries of the metrics search page. | 
**VisualSettings** | Pointer to **string** | Visual settings of the metrics search page. | [optional] 

## Methods

### NewMetricsSearch

`func NewMetricsSearch(title string, timeRange ResolvableTimeRange, queries []Query, ) *MetricsSearch`

NewMetricsSearch instantiates a new MetricsSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsSearchWithDefaults

`func NewMetricsSearchWithDefaults() *MetricsSearch`

NewMetricsSearchWithDefaults instantiates a new MetricsSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *MetricsSearch) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MetricsSearch) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MetricsSearch) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTimeRange

`func (o *MetricsSearch) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsSearch) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsSearch) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetDescription

`func (o *MetricsSearch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsSearch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsSearch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsSearch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQueries

`func (o *MetricsSearch) GetQueries() []Query`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MetricsSearch) GetQueriesOk() (*[]Query, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MetricsSearch) SetQueries(v []Query)`

SetQueries sets Queries field to given value.


### GetVisualSettings

`func (o *MetricsSearch) GetVisualSettings() string`

GetVisualSettings returns the VisualSettings field if non-nil, zero value otherwise.

### GetVisualSettingsOk

`func (o *MetricsSearch) GetVisualSettingsOk() (*string, bool)`

GetVisualSettingsOk returns a tuple with the VisualSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualSettings

`func (o *MetricsSearch) SetVisualSettings(v string)`

SetVisualSettings sets VisualSettings field to given value.

### HasVisualSettings

`func (o *MetricsSearch) HasVisualSettings() bool`

HasVisualSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


