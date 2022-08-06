# StaticConditionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | **string** | The relative time range of the monitor. Valid values of time ranges are &#x60;5m&#x60;, &#x60;10m&#x60;, &#x60;15m&#x60;, &#x60;30m&#x60;, &#x60;1h&#x60;, &#x60;3h&#x60;, &#x60;6h&#x60;, &#x60;12h&#x60;, or &#x60;24h&#x60;. | 
**Threshold** | Pointer to **float64** | The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [optional] [default to 0.0]
**ThresholdType** | Pointer to **string** | The comparison type for the &#x60;threshold&#x60; evaluation. This defines how you want the data value compared. Valid values:   1. &#x60;LessThan&#x60;: Less than than the configured threshold.   2. &#x60;GreaterThan&#x60;: Greater than the configured threshold.   3. &#x60;LessThanOrEqual&#x60;: Less than or equal to the configured threshold.   4. &#x60;GreaterThanOrEqual&#x60;: Greater than or equal to the configured threshold. ThresholdType value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [optional] [default to "GreaterThanOrEqual"]
**Field** | Pointer to **string** | The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If &#x60;field&#x60; is not specified, monitor would default to result count instead. | [optional] 
**OccurrenceType** | **string** | The criteria to evaluate the threshold and thresholdType in the given time range. Valid values:   1. &#x60;AtLeastOnce&#x60;: Trigger if the threshold is met at least once. (NOTE: This is the only valid value if monitorType is &#x60;Metrics&#x60;.)   2. &#x60;Always&#x60;: Trigger if the threshold is met continuously. (NOTE: This is the only valid value if monitorType is &#x60;Metrics&#x60;.)   3. &#x60;ResultCount&#x60;: Trigger if the threshold is met against the count of results. (NOTE: This is the only valid value if monitorType is &#x60;Logs&#x60;.)   4. &#x60;MissingData&#x60;: Trigger if the data is missing. (NOTE: This is valid for both &#x60;Logs&#x60; and &#x60;Metrics&#x60; monitorTypes) | 
**TriggerSource** | **string** | Determines which time series from queries to use for Metrics MissingData and ResolvedMissingData triggers Valid values:   1. &#x60;AllTimeSeries&#x60;: Evaluate the condition against all time series. (NOTE: This option is only valid if monitorType is &#x60;Metrics&#x60;)   2. &#x60;AnyTimeSeries&#x60;: Evaluate the condition against any time series. (NOTE: This option is only valid if monitorType is &#x60;Metrics&#x60;)   3. &#x60;AllResults&#x60;: Evaluate the condition against results from all queries. (NOTE: This option is only valid if monitorType is &#x60;Logs&#x60;) | 

## Methods

### NewStaticConditionAllOf

`func NewStaticConditionAllOf(timeRange string, occurrenceType string, triggerSource string, ) *StaticConditionAllOf`

NewStaticConditionAllOf instantiates a new StaticConditionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticConditionAllOfWithDefaults

`func NewStaticConditionAllOfWithDefaults() *StaticConditionAllOf`

NewStaticConditionAllOfWithDefaults instantiates a new StaticConditionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *StaticConditionAllOf) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *StaticConditionAllOf) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *StaticConditionAllOf) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetThreshold

`func (o *StaticConditionAllOf) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *StaticConditionAllOf) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *StaticConditionAllOf) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *StaticConditionAllOf) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetThresholdType

`func (o *StaticConditionAllOf) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *StaticConditionAllOf) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *StaticConditionAllOf) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.

### HasThresholdType

`func (o *StaticConditionAllOf) HasThresholdType() bool`

HasThresholdType returns a boolean if a field has been set.

### GetField

`func (o *StaticConditionAllOf) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *StaticConditionAllOf) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *StaticConditionAllOf) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *StaticConditionAllOf) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOccurrenceType

`func (o *StaticConditionAllOf) GetOccurrenceType() string`

GetOccurrenceType returns the OccurrenceType field if non-nil, zero value otherwise.

### GetOccurrenceTypeOk

`func (o *StaticConditionAllOf) GetOccurrenceTypeOk() (*string, bool)`

GetOccurrenceTypeOk returns a tuple with the OccurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceType

`func (o *StaticConditionAllOf) SetOccurrenceType(v string)`

SetOccurrenceType sets OccurrenceType field to given value.


### GetTriggerSource

`func (o *StaticConditionAllOf) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *StaticConditionAllOf) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *StaticConditionAllOf) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


