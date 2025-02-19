# SourceTemplateUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaRef** | [**UpgradeSchemaRef**](UpgradeSchemaRef.md) |  | 
**InputJson** | [**SourceTemplateUpgradeRequestInputJson**](SourceTemplateUpgradeRequestInputJson.md) |  | 

## Methods

### NewSourceTemplateUpgradeRequest

`func NewSourceTemplateUpgradeRequest(schemaRef UpgradeSchemaRef, inputJson SourceTemplateUpgradeRequestInputJson, ) *SourceTemplateUpgradeRequest`

NewSourceTemplateUpgradeRequest instantiates a new SourceTemplateUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceTemplateUpgradeRequestWithDefaults

`func NewSourceTemplateUpgradeRequestWithDefaults() *SourceTemplateUpgradeRequest`

NewSourceTemplateUpgradeRequestWithDefaults instantiates a new SourceTemplateUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaRef

`func (o *SourceTemplateUpgradeRequest) GetSchemaRef() UpgradeSchemaRef`

GetSchemaRef returns the SchemaRef field if non-nil, zero value otherwise.

### GetSchemaRefOk

`func (o *SourceTemplateUpgradeRequest) GetSchemaRefOk() (*UpgradeSchemaRef, bool)`

GetSchemaRefOk returns a tuple with the SchemaRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRef

`func (o *SourceTemplateUpgradeRequest) SetSchemaRef(v UpgradeSchemaRef)`

SetSchemaRef sets SchemaRef field to given value.


### GetInputJson

`func (o *SourceTemplateUpgradeRequest) GetInputJson() SourceTemplateUpgradeRequestInputJson`

GetInputJson returns the InputJson field if non-nil, zero value otherwise.

### GetInputJsonOk

`func (o *SourceTemplateUpgradeRequest) GetInputJsonOk() (*SourceTemplateUpgradeRequestInputJson, bool)`

GetInputJsonOk returns a tuple with the InputJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputJson

`func (o *SourceTemplateUpgradeRequest) SetInputJson(v SourceTemplateUpgradeRequestInputJson)`

SetInputJson sets InputJson field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


