# LogSearchDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the item in the content library. | 
**Description** | Pointer to **string** | Item description in the content library. | [optional] 
**Schedule** | Pointer to [**LogSearchScheduleSyncDefinition**](LogSearchScheduleSyncDefinition.md) |  | [optional] 
**Properties** | Pointer to **string** | Aggregate Results Settings and View configurations, Legends settings, and different visualisation settings overrides. Leave this field empty to use the defaults. This property contains JSON object encoded as a string.  | [optional] 

## Methods

### NewLogSearchDefinitionAllOf

`func NewLogSearchDefinitionAllOf(name string, ) *LogSearchDefinitionAllOf`

NewLogSearchDefinitionAllOf instantiates a new LogSearchDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchDefinitionAllOfWithDefaults

`func NewLogSearchDefinitionAllOfWithDefaults() *LogSearchDefinitionAllOf`

NewLogSearchDefinitionAllOfWithDefaults instantiates a new LogSearchDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LogSearchDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogSearchDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogSearchDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LogSearchDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogSearchDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogSearchDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogSearchDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchedule

`func (o *LogSearchDefinitionAllOf) GetSchedule() LogSearchScheduleSyncDefinition`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *LogSearchDefinitionAllOf) GetScheduleOk() (*LogSearchScheduleSyncDefinition, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *LogSearchDefinitionAllOf) SetSchedule(v LogSearchScheduleSyncDefinition)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *LogSearchDefinitionAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetProperties

`func (o *LogSearchDefinitionAllOf) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *LogSearchDefinitionAllOf) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *LogSearchDefinitionAllOf) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *LogSearchDefinitionAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


