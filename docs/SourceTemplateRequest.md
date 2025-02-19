# SourceTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaRef** | [**SchemaRef**](SchemaRef.md) |  | 
**InputJson** | [**SourceTemplateRequestInputJson**](SourceTemplateRequestInputJson.md) |  | 
**Selector** | Pointer to [**Selector**](Selector.md) |  | [optional] 

## Methods

### NewSourceTemplateRequest

`func NewSourceTemplateRequest(schemaRef SchemaRef, inputJson SourceTemplateRequestInputJson, ) *SourceTemplateRequest`

NewSourceTemplateRequest instantiates a new SourceTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceTemplateRequestWithDefaults

`func NewSourceTemplateRequestWithDefaults() *SourceTemplateRequest`

NewSourceTemplateRequestWithDefaults instantiates a new SourceTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaRef

`func (o *SourceTemplateRequest) GetSchemaRef() SchemaRef`

GetSchemaRef returns the SchemaRef field if non-nil, zero value otherwise.

### GetSchemaRefOk

`func (o *SourceTemplateRequest) GetSchemaRefOk() (*SchemaRef, bool)`

GetSchemaRefOk returns a tuple with the SchemaRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRef

`func (o *SourceTemplateRequest) SetSchemaRef(v SchemaRef)`

SetSchemaRef sets SchemaRef field to given value.


### GetInputJson

`func (o *SourceTemplateRequest) GetInputJson() SourceTemplateRequestInputJson`

GetInputJson returns the InputJson field if non-nil, zero value otherwise.

### GetInputJsonOk

`func (o *SourceTemplateRequest) GetInputJsonOk() (*SourceTemplateRequestInputJson, bool)`

GetInputJsonOk returns a tuple with the InputJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputJson

`func (o *SourceTemplateRequest) SetInputJson(v SourceTemplateRequestInputJson)`

SetInputJson sets InputJson field to given value.


### GetSelector

`func (o *SourceTemplateRequest) GetSelector() Selector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *SourceTemplateRequest) GetSelectorOk() (*Selector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *SourceTemplateRequest) SetSelector(v Selector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *SourceTemplateRequest) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


