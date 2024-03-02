# MetricsSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the metrics search page. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**Description** | Pointer to **string** | Description of the metrics search page. | [optional] 
**Queries** | [**[]Query**](Query.md) | Queries of the metrics search page. | 
**VisualSettings** | Pointer to **string** | Visual settings of the metrics search page. | [optional] 
**FolderId** | Pointer to **string** | The identifier of the folder to save the metrics search in. By default it is saved in your personal folder.  | [optional] 
**Id** | Pointer to **string** | Unique identifier for the metrics search page. | [optional] 

## Methods

### NewMetricsSearchResponse

`func NewMetricsSearchResponse(title string, timeRange ResolvableTimeRange, queries []Query, ) *MetricsSearchResponse`

NewMetricsSearchResponse instantiates a new MetricsSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsSearchResponseWithDefaults

`func NewMetricsSearchResponseWithDefaults() *MetricsSearchResponse`

NewMetricsSearchResponseWithDefaults instantiates a new MetricsSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *MetricsSearchResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MetricsSearchResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MetricsSearchResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTimeRange

`func (o *MetricsSearchResponse) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsSearchResponse) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsSearchResponse) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetDescription

`func (o *MetricsSearchResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsSearchResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsSearchResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsSearchResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQueries

`func (o *MetricsSearchResponse) GetQueries() []Query`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MetricsSearchResponse) GetQueriesOk() (*[]Query, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MetricsSearchResponse) SetQueries(v []Query)`

SetQueries sets Queries field to given value.


### GetVisualSettings

`func (o *MetricsSearchResponse) GetVisualSettings() string`

GetVisualSettings returns the VisualSettings field if non-nil, zero value otherwise.

### GetVisualSettingsOk

`func (o *MetricsSearchResponse) GetVisualSettingsOk() (*string, bool)`

GetVisualSettingsOk returns a tuple with the VisualSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualSettings

`func (o *MetricsSearchResponse) SetVisualSettings(v string)`

SetVisualSettings sets VisualSettings field to given value.

### HasVisualSettings

`func (o *MetricsSearchResponse) HasVisualSettings() bool`

HasVisualSettings returns a boolean if a field has been set.

### GetFolderId

`func (o *MetricsSearchResponse) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MetricsSearchResponse) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MetricsSearchResponse) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MetricsSearchResponse) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetId

`func (o *MetricsSearchResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsSearchResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsSearchResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricsSearchResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


