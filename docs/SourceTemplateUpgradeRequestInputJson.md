# SourceTemplateUpgradeRequestInputJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of source template. | 
**Receivers** | **map[string]interface{}** | receiver information of source template | 
**Description** | Pointer to **string** | description of source template | [optional] 
**Processors** | Pointer to **map[string]interface{}** | processors for source template | [optional] 

## Methods

### NewSourceTemplateUpgradeRequestInputJson

`func NewSourceTemplateUpgradeRequestInputJson(name string, receivers map[string]interface{}, ) *SourceTemplateUpgradeRequestInputJson`

NewSourceTemplateUpgradeRequestInputJson instantiates a new SourceTemplateUpgradeRequestInputJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceTemplateUpgradeRequestInputJsonWithDefaults

`func NewSourceTemplateUpgradeRequestInputJsonWithDefaults() *SourceTemplateUpgradeRequestInputJson`

NewSourceTemplateUpgradeRequestInputJsonWithDefaults instantiates a new SourceTemplateUpgradeRequestInputJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourceTemplateUpgradeRequestInputJson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceTemplateUpgradeRequestInputJson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceTemplateUpgradeRequestInputJson) SetName(v string)`

SetName sets Name field to given value.


### GetReceivers

`func (o *SourceTemplateUpgradeRequestInputJson) GetReceivers() map[string]interface{}`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *SourceTemplateUpgradeRequestInputJson) GetReceiversOk() (*map[string]interface{}, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *SourceTemplateUpgradeRequestInputJson) SetReceivers(v map[string]interface{})`

SetReceivers sets Receivers field to given value.


### GetDescription

`func (o *SourceTemplateUpgradeRequestInputJson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SourceTemplateUpgradeRequestInputJson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SourceTemplateUpgradeRequestInputJson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SourceTemplateUpgradeRequestInputJson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProcessors

`func (o *SourceTemplateUpgradeRequestInputJson) GetProcessors() map[string]interface{}`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *SourceTemplateUpgradeRequestInputJson) GetProcessorsOk() (*map[string]interface{}, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *SourceTemplateUpgradeRequestInputJson) SetProcessors(v map[string]interface{})`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *SourceTemplateUpgradeRequestInputJson) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


