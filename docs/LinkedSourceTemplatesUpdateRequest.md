# LinkedSourceTemplatesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectorId** | **string** | otCollector id for which tags are edited. | 
**Tags** | Pointer to **map[string]string** | JSON map of key-value metadata to apply to the otCollector. | [optional] [default to {}]

## Methods

### NewLinkedSourceTemplatesUpdateRequest

`func NewLinkedSourceTemplatesUpdateRequest(collectorId string, ) *LinkedSourceTemplatesUpdateRequest`

NewLinkedSourceTemplatesUpdateRequest instantiates a new LinkedSourceTemplatesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedSourceTemplatesUpdateRequestWithDefaults

`func NewLinkedSourceTemplatesUpdateRequestWithDefaults() *LinkedSourceTemplatesUpdateRequest`

NewLinkedSourceTemplatesUpdateRequestWithDefaults instantiates a new LinkedSourceTemplatesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectorId

`func (o *LinkedSourceTemplatesUpdateRequest) GetCollectorId() string`

GetCollectorId returns the CollectorId field if non-nil, zero value otherwise.

### GetCollectorIdOk

`func (o *LinkedSourceTemplatesUpdateRequest) GetCollectorIdOk() (*string, bool)`

GetCollectorIdOk returns a tuple with the CollectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorId

`func (o *LinkedSourceTemplatesUpdateRequest) SetCollectorId(v string)`

SetCollectorId sets CollectorId field to given value.


### GetTags

`func (o *LinkedSourceTemplatesUpdateRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LinkedSourceTemplatesUpdateRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LinkedSourceTemplatesUpdateRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LinkedSourceTemplatesUpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


