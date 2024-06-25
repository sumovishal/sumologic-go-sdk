# MutingSchedulesLibraryMutingScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**ScheduleDefinition**](ScheduleDefinition.md) |  | 
**Monitor** | Pointer to [**MonitorScope**](MonitorScope.md) |  | [optional] 
**NotificationGroups** | Pointer to [**[]GroupDefinition**](GroupDefinition.md) |  | [optional] [default to []]

## Methods

### NewMutingSchedulesLibraryMutingScheduleAllOf

`func NewMutingSchedulesLibraryMutingScheduleAllOf(schedule ScheduleDefinition, ) *MutingSchedulesLibraryMutingScheduleAllOf`

NewMutingSchedulesLibraryMutingScheduleAllOf instantiates a new MutingSchedulesLibraryMutingScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutingSchedulesLibraryMutingScheduleAllOfWithDefaults

`func NewMutingSchedulesLibraryMutingScheduleAllOfWithDefaults() *MutingSchedulesLibraryMutingScheduleAllOf`

NewMutingSchedulesLibraryMutingScheduleAllOfWithDefaults instantiates a new MutingSchedulesLibraryMutingScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetSchedule() ScheduleDefinition`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetScheduleOk() (*ScheduleDefinition, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) SetSchedule(v ScheduleDefinition)`

SetSchedule sets Schedule field to given value.


### GetMonitor

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetMonitor() MonitorScope`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetMonitorOk() (*MonitorScope, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) SetMonitor(v MonitorScope)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetNotificationGroups

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetNotificationGroups() []GroupDefinition`

GetNotificationGroups returns the NotificationGroups field if non-nil, zero value otherwise.

### GetNotificationGroupsOk

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) GetNotificationGroupsOk() (*[]GroupDefinition, bool)`

GetNotificationGroupsOk returns a tuple with the NotificationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationGroups

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) SetNotificationGroups(v []GroupDefinition)`

SetNotificationGroups sets NotificationGroups field to given value.

### HasNotificationGroups

`func (o *MutingSchedulesLibraryMutingScheduleAllOf) HasNotificationGroups() bool`

HasNotificationGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


