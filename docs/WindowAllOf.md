# WindowAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | **float32** | Threshold for classifying window as successful or unsuccessful. | 
**Op** | **string** | Comparison function with window threshold (LessThan/GreaterThan/LessThanOrEqual/GreaterThanOrEqual). | 
**Aggregation** | Pointer to **string** | Aggregation function applied over each window to arrive at SLI. Must be &#x60;Avg&#x60;, &#x60;Min&#x60;, &#x60;Max&#x60;, &#x60;Sum&#x60;, or percentile of the form &#x60;pX&#x60; where &#x60;X&#x60; is an integer between 0 and 100. | [optional] 
**Size** | **string** | Size of the aggregation window (minimum of 1m and maximum of 1h). | 

## Methods

### NewWindowAllOf

`func NewWindowAllOf(threshold float32, op string, size string, ) *WindowAllOf`

NewWindowAllOf instantiates a new WindowAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowAllOfWithDefaults

`func NewWindowAllOfWithDefaults() *WindowAllOf`

NewWindowAllOfWithDefaults instantiates a new WindowAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *WindowAllOf) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *WindowAllOf) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *WindowAllOf) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.


### GetOp

`func (o *WindowAllOf) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *WindowAllOf) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *WindowAllOf) SetOp(v string)`

SetOp sets Op field to given value.


### GetAggregation

`func (o *WindowAllOf) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *WindowAllOf) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *WindowAllOf) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *WindowAllOf) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetSize

`func (o *WindowAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *WindowAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *WindowAllOf) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


