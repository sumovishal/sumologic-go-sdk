# SavedSearchSyncDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTimeRange** | **string** | Default time range for the search. Possible types of time ranges are:   - relative time range: e.g. \&quot;-1d -12h\&quot; represents a time range from one day ago to 12 hours ago.   - absolute time range: e.g. \&quot;01-04-2017 20:32:00 to 01-04-2017 20:35:00\&quot; represents a time range   from April 1st, 2017 at 8:32 PM until April 1st, 2017 at 8:35 PM. | 

## Methods

### NewSavedSearchSyncDefinitionAllOf

`func NewSavedSearchSyncDefinitionAllOf(defaultTimeRange string, ) *SavedSearchSyncDefinitionAllOf`

NewSavedSearchSyncDefinitionAllOf instantiates a new SavedSearchSyncDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedSearchSyncDefinitionAllOfWithDefaults

`func NewSavedSearchSyncDefinitionAllOfWithDefaults() *SavedSearchSyncDefinitionAllOf`

NewSavedSearchSyncDefinitionAllOfWithDefaults instantiates a new SavedSearchSyncDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTimeRange

`func (o *SavedSearchSyncDefinitionAllOf) GetDefaultTimeRange() string`

GetDefaultTimeRange returns the DefaultTimeRange field if non-nil, zero value otherwise.

### GetDefaultTimeRangeOk

`func (o *SavedSearchSyncDefinitionAllOf) GetDefaultTimeRangeOk() (*string, bool)`

GetDefaultTimeRangeOk returns a tuple with the DefaultTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeRange

`func (o *SavedSearchSyncDefinitionAllOf) SetDefaultTimeRange(v string)`

SetDefaultTimeRange sets DefaultTimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


