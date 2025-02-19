# AnomalyCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | Pointer to **string** | The relative time range of the monitor. Valid values of time ranges are &#x60;-5m&#x60;, &#x60;-10m&#x60;, &#x60;-15m&#x60;, &#x60;-30m&#x60;, &#x60;-1h&#x60;, &#x60;-3h&#x60;, &#x60;-6h&#x60;, &#x60;-12h&#x60;, &#x60;-24h&#x60; or &#x60;-1d&#x60;. | [optional] 
**Sensitivity** | Pointer to **float64** | The triggering sensitivity of the anomaly model used for this monitor. | [optional] [default to 0.5]
**AnomalyDetectorType** | Pointer to **string** | The type of anomaly model that will be used for evaluating this monitor. Only &#x60;Cluster&#x60; option is supported currently. | [optional] 
**Field** | Pointer to **string** | The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If &#x60;field&#x60; is not specified, monitor would default to result count instead. | [optional] 
**MinAnomalyCount** | Pointer to **int32** | The minimum number of anomalies required to exist in the current time range for the condition to trigger. | [optional] [default to 1]
**Direction** | Pointer to **string** | Specifies which direction should trigger violations. | [optional] [default to "Both"]

## Methods

### NewAnomalyCondition

`func NewAnomalyCondition() *AnomalyCondition`

NewAnomalyCondition instantiates a new AnomalyCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyConditionWithDefaults

`func NewAnomalyConditionWithDefaults() *AnomalyCondition`

NewAnomalyConditionWithDefaults instantiates a new AnomalyCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *AnomalyCondition) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *AnomalyCondition) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *AnomalyCondition) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *AnomalyCondition) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetSensitivity

`func (o *AnomalyCondition) GetSensitivity() float64`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *AnomalyCondition) GetSensitivityOk() (*float64, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *AnomalyCondition) SetSensitivity(v float64)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *AnomalyCondition) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### GetAnomalyDetectorType

`func (o *AnomalyCondition) GetAnomalyDetectorType() string`

GetAnomalyDetectorType returns the AnomalyDetectorType field if non-nil, zero value otherwise.

### GetAnomalyDetectorTypeOk

`func (o *AnomalyCondition) GetAnomalyDetectorTypeOk() (*string, bool)`

GetAnomalyDetectorTypeOk returns a tuple with the AnomalyDetectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyDetectorType

`func (o *AnomalyCondition) SetAnomalyDetectorType(v string)`

SetAnomalyDetectorType sets AnomalyDetectorType field to given value.

### HasAnomalyDetectorType

`func (o *AnomalyCondition) HasAnomalyDetectorType() bool`

HasAnomalyDetectorType returns a boolean if a field has been set.

### GetField

`func (o *AnomalyCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AnomalyCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AnomalyCondition) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *AnomalyCondition) HasField() bool`

HasField returns a boolean if a field has been set.

### GetMinAnomalyCount

`func (o *AnomalyCondition) GetMinAnomalyCount() int32`

GetMinAnomalyCount returns the MinAnomalyCount field if non-nil, zero value otherwise.

### GetMinAnomalyCountOk

`func (o *AnomalyCondition) GetMinAnomalyCountOk() (*int32, bool)`

GetMinAnomalyCountOk returns a tuple with the MinAnomalyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAnomalyCount

`func (o *AnomalyCondition) SetMinAnomalyCount(v int32)`

SetMinAnomalyCount sets MinAnomalyCount field to given value.

### HasMinAnomalyCount

`func (o *AnomalyCondition) HasMinAnomalyCount() bool`

HasMinAnomalyCount returns a boolean if a field has been set.

### GetDirection

`func (o *AnomalyCondition) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AnomalyCondition) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AnomalyCondition) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AnomalyCondition) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


