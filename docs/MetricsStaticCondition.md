# MetricsStaticCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | **string** | The relative time range of the monitor. Valid values of time ranges are &#x60;-5m&#x60;, &#x60;-10m&#x60;, &#x60;-15m&#x60;, &#x60;-30m&#x60;, &#x60;-1h&#x60;, &#x60;-3h&#x60;, &#x60;-6h&#x60;, &#x60;-12h&#x60;, or &#x60;-24h&#x60;. | 
**Threshold** | **float64** | The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [default to 0.0]
**ThresholdType** | **string** | The comparison type for the &#x60;threshold&#x60; evaluation. This defines how you want the data value compared. Valid values:   1. &#x60;LessThan&#x60;: Less than than the configured threshold.   2. &#x60;GreaterThan&#x60;: Greater than the configured threshold.   3. &#x60;LessThanOrEqual&#x60;: Less than or equal to the configured threshold.   4. &#x60;GreaterThanOrEqual&#x60;: Greater than or equal to the configured threshold. ThresholdType value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [default to "GreaterThanOrEqual"]
**OccurrenceType** | **string** | The criteria to evaluate the threshold and thresholdType in the given time range. Valid values:   1. &#x60;AtLeastOnce&#x60;: Trigger if the threshold is met at least once. (NOTE: This is the only valid value if monitorType is &#x60;Metrics&#x60;.)   2. &#x60;Always&#x60;: Trigger if the threshold is met continuously. (NOTE: This is the only valid value if monitorType is &#x60;Metrics&#x60;.)   3. &#x60;ResultCount&#x60;: Trigger if the threshold is met against the count of results. (NOTE: This is the only valid value if monitorType is &#x60;Logs&#x60;.)   4. &#x60;MissingData&#x60;: Trigger if the data is missing. (NOTE: This is valid for both &#x60;Logs&#x60; and &#x60;Metrics&#x60; monitorTypes) | 
**MinDataPoints** | Pointer to **int32** | The minimum number of data points required for the monitor to alert or resolve within the time range specified. This field will always be set to 1 for &#x60;AtleastOnce&#x60; occurrence type and for &#x60;Always&#x60;, if not specified by user it will default to 2. | [optional] 

## Methods

### NewMetricsStaticCondition

`func NewMetricsStaticCondition(timeRange string, threshold float64, thresholdType string, occurrenceType string, ) *MetricsStaticCondition`

NewMetricsStaticCondition instantiates a new MetricsStaticCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsStaticConditionWithDefaults

`func NewMetricsStaticConditionWithDefaults() *MetricsStaticCondition`

NewMetricsStaticConditionWithDefaults instantiates a new MetricsStaticCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *MetricsStaticCondition) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsStaticCondition) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsStaticCondition) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetThreshold

`func (o *MetricsStaticCondition) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricsStaticCondition) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricsStaticCondition) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetThresholdType

`func (o *MetricsStaticCondition) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *MetricsStaticCondition) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *MetricsStaticCondition) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.


### GetOccurrenceType

`func (o *MetricsStaticCondition) GetOccurrenceType() string`

GetOccurrenceType returns the OccurrenceType field if non-nil, zero value otherwise.

### GetOccurrenceTypeOk

`func (o *MetricsStaticCondition) GetOccurrenceTypeOk() (*string, bool)`

GetOccurrenceTypeOk returns a tuple with the OccurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceType

`func (o *MetricsStaticCondition) SetOccurrenceType(v string)`

SetOccurrenceType sets OccurrenceType field to given value.


### GetMinDataPoints

`func (o *MetricsStaticCondition) GetMinDataPoints() int32`

GetMinDataPoints returns the MinDataPoints field if non-nil, zero value otherwise.

### GetMinDataPointsOk

`func (o *MetricsStaticCondition) GetMinDataPointsOk() (*int32, bool)`

GetMinDataPointsOk returns a tuple with the MinDataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDataPoints

`func (o *MetricsStaticCondition) SetMinDataPoints(v int32)`

SetMinDataPoints sets MinDataPoints field to given value.

### HasMinDataPoints

`func (o *MetricsStaticCondition) HasMinDataPoints() bool`

HasMinDataPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


