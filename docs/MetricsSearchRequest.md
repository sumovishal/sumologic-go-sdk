# MetricsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the metrics search page. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**Description** | Pointer to **string** | Description of the metrics search page. | [optional] 
**Queries** | [**[]Query**](Query.md) | Queries of the metrics search page. | 
**VisualSettings** | Pointer to **string** | Visual settings of the metrics search page. | [optional] 
**FolderId** | Pointer to **string** | The identifier of the folder to save the metrics search in. By default it is saved in your personal folder.  | [optional] 

## Methods

### NewMetricsSearchRequest

`func NewMetricsSearchRequest(title string, timeRange ResolvableTimeRange, queries []Query, ) *MetricsSearchRequest`

NewMetricsSearchRequest instantiates a new MetricsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsSearchRequestWithDefaults

`func NewMetricsSearchRequestWithDefaults() *MetricsSearchRequest`

NewMetricsSearchRequestWithDefaults instantiates a new MetricsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *MetricsSearchRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MetricsSearchRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MetricsSearchRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTimeRange

`func (o *MetricsSearchRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsSearchRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsSearchRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetDescription

`func (o *MetricsSearchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsSearchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsSearchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsSearchRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQueries

`func (o *MetricsSearchRequest) GetQueries() []Query`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MetricsSearchRequest) GetQueriesOk() (*[]Query, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MetricsSearchRequest) SetQueries(v []Query)`

SetQueries sets Queries field to given value.


### GetVisualSettings

`func (o *MetricsSearchRequest) GetVisualSettings() string`

GetVisualSettings returns the VisualSettings field if non-nil, zero value otherwise.

### GetVisualSettingsOk

`func (o *MetricsSearchRequest) GetVisualSettingsOk() (*string, bool)`

GetVisualSettingsOk returns a tuple with the VisualSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualSettings

`func (o *MetricsSearchRequest) SetVisualSettings(v string)`

SetVisualSettings sets VisualSettings field to given value.

### HasVisualSettings

`func (o *MetricsSearchRequest) HasVisualSettings() bool`

HasVisualSettings returns a boolean if a field has been set.

### GetFolderId

`func (o *MetricsSearchRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MetricsSearchRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MetricsSearchRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MetricsSearchRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


