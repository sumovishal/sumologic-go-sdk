# LinkingUpdatedSourceTemplateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceTemplateDefinition** | [**SourceTemplateDefinition**](SourceTemplateDefinition.md) |  | 
**ReasonTags** | [**[][]CollectorTag**]([]CollectorTag.md) | tags which are responsible for source template and collector linking impact. | 

## Methods

### NewLinkingUpdatedSourceTemplateDetails

`func NewLinkingUpdatedSourceTemplateDetails(sourceTemplateDefinition SourceTemplateDefinition, reasonTags [][]CollectorTag, ) *LinkingUpdatedSourceTemplateDetails`

NewLinkingUpdatedSourceTemplateDetails instantiates a new LinkingUpdatedSourceTemplateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkingUpdatedSourceTemplateDetailsWithDefaults

`func NewLinkingUpdatedSourceTemplateDetailsWithDefaults() *LinkingUpdatedSourceTemplateDetails`

NewLinkingUpdatedSourceTemplateDetailsWithDefaults instantiates a new LinkingUpdatedSourceTemplateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceTemplateDefinition

`func (o *LinkingUpdatedSourceTemplateDetails) GetSourceTemplateDefinition() SourceTemplateDefinition`

GetSourceTemplateDefinition returns the SourceTemplateDefinition field if non-nil, zero value otherwise.

### GetSourceTemplateDefinitionOk

`func (o *LinkingUpdatedSourceTemplateDetails) GetSourceTemplateDefinitionOk() (*SourceTemplateDefinition, bool)`

GetSourceTemplateDefinitionOk returns a tuple with the SourceTemplateDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTemplateDefinition

`func (o *LinkingUpdatedSourceTemplateDetails) SetSourceTemplateDefinition(v SourceTemplateDefinition)`

SetSourceTemplateDefinition sets SourceTemplateDefinition field to given value.


### GetReasonTags

`func (o *LinkingUpdatedSourceTemplateDetails) GetReasonTags() [][]CollectorTag`

GetReasonTags returns the ReasonTags field if non-nil, zero value otherwise.

### GetReasonTagsOk

`func (o *LinkingUpdatedSourceTemplateDetails) GetReasonTagsOk() (*[][]CollectorTag, bool)`

GetReasonTagsOk returns a tuple with the ReasonTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonTags

`func (o *LinkingUpdatedSourceTemplateDetails) SetReasonTags(v [][]CollectorTag)`

SetReasonTags sets ReasonTags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


