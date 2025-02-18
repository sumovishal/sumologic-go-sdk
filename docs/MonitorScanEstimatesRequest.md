# MonitorScanEstimatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The logs query that defines the stream of data the monitor runs on. | 
**Triggers** | [**[]TriggerCondition**](TriggerCondition.md) | Defines the conditions of when to send notifications. | 
**Timezone** | **string** | Time zone for the monitor [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 

## Methods

### NewMonitorScanEstimatesRequest

`func NewMonitorScanEstimatesRequest(query string, triggers []TriggerCondition, timezone string, ) *MonitorScanEstimatesRequest`

NewMonitorScanEstimatesRequest instantiates a new MonitorScanEstimatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorScanEstimatesRequestWithDefaults

`func NewMonitorScanEstimatesRequestWithDefaults() *MonitorScanEstimatesRequest`

NewMonitorScanEstimatesRequestWithDefaults instantiates a new MonitorScanEstimatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *MonitorScanEstimatesRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitorScanEstimatesRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MonitorScanEstimatesRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTriggers

`func (o *MonitorScanEstimatesRequest) GetTriggers() []TriggerCondition`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MonitorScanEstimatesRequest) GetTriggersOk() (*[]TriggerCondition, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *MonitorScanEstimatesRequest) SetTriggers(v []TriggerCondition)`

SetTriggers sets Triggers field to given value.


### GetTimezone

`func (o *MonitorScanEstimatesRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MonitorScanEstimatesRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MonitorScanEstimatesRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


