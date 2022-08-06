# DashboardTemplateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the dashboard. | 
**PanelToSessionIdMap** | Pointer to **map[string]string** | A map of panel to session id. The session id will be used to fetch data of the panel for the report. If not specified, a new session id will be created for the panel.  | [optional] 
**TimeRange** | Pointer to [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] 
**VariableValues** | Pointer to [**VariablesValuesData**](VariablesValuesData.md) |  | [optional] 

## Methods

### NewDashboardTemplateAllOf

`func NewDashboardTemplateAllOf(id string, ) *DashboardTemplateAllOf`

NewDashboardTemplateAllOf instantiates a new DashboardTemplateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardTemplateAllOfWithDefaults

`func NewDashboardTemplateAllOfWithDefaults() *DashboardTemplateAllOf`

NewDashboardTemplateAllOfWithDefaults instantiates a new DashboardTemplateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardTemplateAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardTemplateAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardTemplateAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetPanelToSessionIdMap

`func (o *DashboardTemplateAllOf) GetPanelToSessionIdMap() map[string]string`

GetPanelToSessionIdMap returns the PanelToSessionIdMap field if non-nil, zero value otherwise.

### GetPanelToSessionIdMapOk

`func (o *DashboardTemplateAllOf) GetPanelToSessionIdMapOk() (*map[string]string, bool)`

GetPanelToSessionIdMapOk returns a tuple with the PanelToSessionIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelToSessionIdMap

`func (o *DashboardTemplateAllOf) SetPanelToSessionIdMap(v map[string]string)`

SetPanelToSessionIdMap sets PanelToSessionIdMap field to given value.

### HasPanelToSessionIdMap

`func (o *DashboardTemplateAllOf) HasPanelToSessionIdMap() bool`

HasPanelToSessionIdMap returns a boolean if a field has been set.

### GetTimeRange

`func (o *DashboardTemplateAllOf) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DashboardTemplateAllOf) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DashboardTemplateAllOf) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *DashboardTemplateAllOf) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetVariableValues

`func (o *DashboardTemplateAllOf) GetVariableValues() VariablesValuesData`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *DashboardTemplateAllOf) GetVariableValuesOk() (*VariablesValuesData, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *DashboardTemplateAllOf) SetVariableValues(v VariablesValuesData)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *DashboardTemplateAllOf) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


