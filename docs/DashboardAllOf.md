# DashboardAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the dashboard. This id is used to get detailed information about the dashboard, such as panels, variables and the layout.  | [optional] 
**ContentId** | Pointer to **string** | Content identifier for the dashboard. This id is used to connect to the Sumo Content Library and get general metadata about the dashboard. Use this id if you want to search for dashboards in Sumo folders.  | [optional] 
**ScheduleId** | Pointer to **string** | Scheduled report identifier for the dashboard. Only most recently modified report schedule is rerun per dashboard. This id is used to manage the schedule details through the scheduled report API.  | [optional] 
**ScheduleCount** | Pointer to **int32** | Count of report schedules for the dashboard. | [optional] 

## Methods

### NewDashboardAllOf

`func NewDashboardAllOf() *DashboardAllOf`

NewDashboardAllOf instantiates a new DashboardAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardAllOfWithDefaults

`func NewDashboardAllOfWithDefaults() *DashboardAllOf`

NewDashboardAllOfWithDefaults instantiates a new DashboardAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentId

`func (o *DashboardAllOf) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *DashboardAllOf) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *DashboardAllOf) SetContentId(v string)`

SetContentId sets ContentId field to given value.

### HasContentId

`func (o *DashboardAllOf) HasContentId() bool`

HasContentId returns a boolean if a field has been set.

### GetScheduleId

`func (o *DashboardAllOf) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DashboardAllOf) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DashboardAllOf) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *DashboardAllOf) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetScheduleCount

`func (o *DashboardAllOf) GetScheduleCount() int32`

GetScheduleCount returns the ScheduleCount field if non-nil, zero value otherwise.

### GetScheduleCountOk

`func (o *DashboardAllOf) GetScheduleCountOk() (*int32, bool)`

GetScheduleCountOk returns a tuple with the ScheduleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCount

`func (o *DashboardAllOf) SetScheduleCount(v int32)`

SetScheduleCount sets ScheduleCount field to given value.

### HasScheduleCount

`func (o *DashboardAllOf) HasScheduleCount() bool`

HasScheduleCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


