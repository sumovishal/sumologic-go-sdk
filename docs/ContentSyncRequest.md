# ContentSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceChildOrgInfo** | [**ChildOrgInfo**](ChildOrgInfo.md) |  | 
**DestinationChildOrgInfo** | [**DestinationChildOrgInfo**](DestinationChildOrgInfo.md) |  | 
**ContentList** | [**[]Content1**](Content1.md) | List of Content and Configuration Information. | 

## Methods

### NewContentSyncRequest

`func NewContentSyncRequest(sourceChildOrgInfo ChildOrgInfo, destinationChildOrgInfo DestinationChildOrgInfo, contentList []Content1, ) *ContentSyncRequest`

NewContentSyncRequest instantiates a new ContentSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSyncRequestWithDefaults

`func NewContentSyncRequestWithDefaults() *ContentSyncRequest`

NewContentSyncRequestWithDefaults instantiates a new ContentSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceChildOrgInfo

`func (o *ContentSyncRequest) GetSourceChildOrgInfo() ChildOrgInfo`

GetSourceChildOrgInfo returns the SourceChildOrgInfo field if non-nil, zero value otherwise.

### GetSourceChildOrgInfoOk

`func (o *ContentSyncRequest) GetSourceChildOrgInfoOk() (*ChildOrgInfo, bool)`

GetSourceChildOrgInfoOk returns a tuple with the SourceChildOrgInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChildOrgInfo

`func (o *ContentSyncRequest) SetSourceChildOrgInfo(v ChildOrgInfo)`

SetSourceChildOrgInfo sets SourceChildOrgInfo field to given value.


### GetDestinationChildOrgInfo

`func (o *ContentSyncRequest) GetDestinationChildOrgInfo() DestinationChildOrgInfo`

GetDestinationChildOrgInfo returns the DestinationChildOrgInfo field if non-nil, zero value otherwise.

### GetDestinationChildOrgInfoOk

`func (o *ContentSyncRequest) GetDestinationChildOrgInfoOk() (*DestinationChildOrgInfo, bool)`

GetDestinationChildOrgInfoOk returns a tuple with the DestinationChildOrgInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationChildOrgInfo

`func (o *ContentSyncRequest) SetDestinationChildOrgInfo(v DestinationChildOrgInfo)`

SetDestinationChildOrgInfo sets DestinationChildOrgInfo field to given value.


### GetContentList

`func (o *ContentSyncRequest) GetContentList() []Content1`

GetContentList returns the ContentList field if non-nil, zero value otherwise.

### GetContentListOk

`func (o *ContentSyncRequest) GetContentListOk() (*[]Content1, bool)`

GetContentListOk returns a tuple with the ContentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentList

`func (o *ContentSyncRequest) SetContentList(v []Content1)`

SetContentList sets ContentList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


