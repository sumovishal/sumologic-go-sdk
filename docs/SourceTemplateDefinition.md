# SourceTemplateDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaRef** | Pointer to [**SchemaRef**](SchemaRef.md) |  | [optional] 
**Id** | Pointer to **string** | id of source template. | [optional] 
**InputJson** | Pointer to **map[string]interface{}** | inputJson of source template | [optional] 
**Config** | Pointer to **string** | configuration of source template | [optional] 
**Selector** | Pointer to [**Selector**](Selector.md) |  | [optional] 
**TotalCollectorLinked** | Pointer to **int32** | count of total collector linked with this source template. | [optional] [default to 0]
**CreatedAt** | Pointer to **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Modification timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] 
**CreatedBy** | Pointer to **string** | Id of the user who created source template | [optional] 
**ModifiedBy** | Pointer to **string** | Id of the user who last modified the source template | [optional] 
**Status** | Pointer to **string** | Status of Source template | [optional] 
**IsEnabled** | Pointer to **bool** | A boolean parameter to get if the source template is enabled. | [optional] [default to true]

## Methods

### NewSourceTemplateDefinition

`func NewSourceTemplateDefinition() *SourceTemplateDefinition`

NewSourceTemplateDefinition instantiates a new SourceTemplateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceTemplateDefinitionWithDefaults

`func NewSourceTemplateDefinitionWithDefaults() *SourceTemplateDefinition`

NewSourceTemplateDefinitionWithDefaults instantiates a new SourceTemplateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaRef

`func (o *SourceTemplateDefinition) GetSchemaRef() SchemaRef`

GetSchemaRef returns the SchemaRef field if non-nil, zero value otherwise.

### GetSchemaRefOk

`func (o *SourceTemplateDefinition) GetSchemaRefOk() (*SchemaRef, bool)`

GetSchemaRefOk returns a tuple with the SchemaRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRef

`func (o *SourceTemplateDefinition) SetSchemaRef(v SchemaRef)`

SetSchemaRef sets SchemaRef field to given value.

### HasSchemaRef

`func (o *SourceTemplateDefinition) HasSchemaRef() bool`

HasSchemaRef returns a boolean if a field has been set.

### GetId

`func (o *SourceTemplateDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceTemplateDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceTemplateDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SourceTemplateDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputJson

`func (o *SourceTemplateDefinition) GetInputJson() map[string]interface{}`

GetInputJson returns the InputJson field if non-nil, zero value otherwise.

### GetInputJsonOk

`func (o *SourceTemplateDefinition) GetInputJsonOk() (*map[string]interface{}, bool)`

GetInputJsonOk returns a tuple with the InputJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputJson

`func (o *SourceTemplateDefinition) SetInputJson(v map[string]interface{})`

SetInputJson sets InputJson field to given value.

### HasInputJson

`func (o *SourceTemplateDefinition) HasInputJson() bool`

HasInputJson returns a boolean if a field has been set.

### GetConfig

`func (o *SourceTemplateDefinition) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SourceTemplateDefinition) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SourceTemplateDefinition) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SourceTemplateDefinition) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSelector

`func (o *SourceTemplateDefinition) GetSelector() Selector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *SourceTemplateDefinition) GetSelectorOk() (*Selector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *SourceTemplateDefinition) SetSelector(v Selector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *SourceTemplateDefinition) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetTotalCollectorLinked

`func (o *SourceTemplateDefinition) GetTotalCollectorLinked() int32`

GetTotalCollectorLinked returns the TotalCollectorLinked field if non-nil, zero value otherwise.

### GetTotalCollectorLinkedOk

`func (o *SourceTemplateDefinition) GetTotalCollectorLinkedOk() (*int32, bool)`

GetTotalCollectorLinkedOk returns a tuple with the TotalCollectorLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollectorLinked

`func (o *SourceTemplateDefinition) SetTotalCollectorLinked(v int32)`

SetTotalCollectorLinked sets TotalCollectorLinked field to given value.

### HasTotalCollectorLinked

`func (o *SourceTemplateDefinition) HasTotalCollectorLinked() bool`

HasTotalCollectorLinked returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SourceTemplateDefinition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SourceTemplateDefinition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SourceTemplateDefinition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SourceTemplateDefinition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SourceTemplateDefinition) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SourceTemplateDefinition) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SourceTemplateDefinition) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SourceTemplateDefinition) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SourceTemplateDefinition) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SourceTemplateDefinition) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SourceTemplateDefinition) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SourceTemplateDefinition) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *SourceTemplateDefinition) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SourceTemplateDefinition) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SourceTemplateDefinition) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *SourceTemplateDefinition) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetStatus

`func (o *SourceTemplateDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceTemplateDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceTemplateDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceTemplateDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsEnabled

`func (o *SourceTemplateDefinition) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SourceTemplateDefinition) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SourceTemplateDefinition) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SourceTemplateDefinition) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


