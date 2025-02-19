# LinkedSourceTemplatesUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectorId** | **string** | otCollector id for which tags are edited. | 
**AddedSourceTemplates** | Pointer to [**[]LinkingUpdatedSourceTemplateDetails**](LinkingUpdatedSourceTemplateDetails.md) | list of sourceTemplates which are linked to otCollector. | [optional] 
**RemovedSourceTemplates** | Pointer to [**[]LinkingUpdatedSourceTemplateDetails**](LinkingUpdatedSourceTemplateDetails.md) | list of sourceTemplates which are removed from otCollector linking. | [optional] 

## Methods

### NewLinkedSourceTemplatesUpdateResponse

`func NewLinkedSourceTemplatesUpdateResponse(collectorId string, ) *LinkedSourceTemplatesUpdateResponse`

NewLinkedSourceTemplatesUpdateResponse instantiates a new LinkedSourceTemplatesUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedSourceTemplatesUpdateResponseWithDefaults

`func NewLinkedSourceTemplatesUpdateResponseWithDefaults() *LinkedSourceTemplatesUpdateResponse`

NewLinkedSourceTemplatesUpdateResponseWithDefaults instantiates a new LinkedSourceTemplatesUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectorId

`func (o *LinkedSourceTemplatesUpdateResponse) GetCollectorId() string`

GetCollectorId returns the CollectorId field if non-nil, zero value otherwise.

### GetCollectorIdOk

`func (o *LinkedSourceTemplatesUpdateResponse) GetCollectorIdOk() (*string, bool)`

GetCollectorIdOk returns a tuple with the CollectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorId

`func (o *LinkedSourceTemplatesUpdateResponse) SetCollectorId(v string)`

SetCollectorId sets CollectorId field to given value.


### GetAddedSourceTemplates

`func (o *LinkedSourceTemplatesUpdateResponse) GetAddedSourceTemplates() []LinkingUpdatedSourceTemplateDetails`

GetAddedSourceTemplates returns the AddedSourceTemplates field if non-nil, zero value otherwise.

### GetAddedSourceTemplatesOk

`func (o *LinkedSourceTemplatesUpdateResponse) GetAddedSourceTemplatesOk() (*[]LinkingUpdatedSourceTemplateDetails, bool)`

GetAddedSourceTemplatesOk returns a tuple with the AddedSourceTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedSourceTemplates

`func (o *LinkedSourceTemplatesUpdateResponse) SetAddedSourceTemplates(v []LinkingUpdatedSourceTemplateDetails)`

SetAddedSourceTemplates sets AddedSourceTemplates field to given value.

### HasAddedSourceTemplates

`func (o *LinkedSourceTemplatesUpdateResponse) HasAddedSourceTemplates() bool`

HasAddedSourceTemplates returns a boolean if a field has been set.

### GetRemovedSourceTemplates

`func (o *LinkedSourceTemplatesUpdateResponse) GetRemovedSourceTemplates() []LinkingUpdatedSourceTemplateDetails`

GetRemovedSourceTemplates returns the RemovedSourceTemplates field if non-nil, zero value otherwise.

### GetRemovedSourceTemplatesOk

`func (o *LinkedSourceTemplatesUpdateResponse) GetRemovedSourceTemplatesOk() (*[]LinkingUpdatedSourceTemplateDetails, bool)`

GetRemovedSourceTemplatesOk returns a tuple with the RemovedSourceTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedSourceTemplates

`func (o *LinkedSourceTemplatesUpdateResponse) SetRemovedSourceTemplates(v []LinkingUpdatedSourceTemplateDetails)`

SetRemovedSourceTemplates sets RemovedSourceTemplates field to given value.

### HasRemovedSourceTemplates

`func (o *LinkedSourceTemplatesUpdateResponse) HasRemovedSourceTemplates() bool`

HasRemovedSourceTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


