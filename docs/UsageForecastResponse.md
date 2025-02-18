# UsageForecastResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageUsage** | Pointer to **float64** | Average credit usage per day till now. | [optional] 
**UsagePercentage** | Pointer to **float64** | Percentage of total credits used till date. | [optional] 
**ForecastedUsage** | Pointer to **float64** | Total expected usage by the end of contract period. | [optional] 
**ForecastedUsagePercentage** | Pointer to **float64** | Percentage of allocated credits that will be used in the contract period. | [optional] 
**RemainingDays** | Pointer to **float64** | Days remaining till all the credits are consumed. | [optional] 

## Methods

### NewUsageForecastResponse

`func NewUsageForecastResponse() *UsageForecastResponse`

NewUsageForecastResponse instantiates a new UsageForecastResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageForecastResponseWithDefaults

`func NewUsageForecastResponseWithDefaults() *UsageForecastResponse`

NewUsageForecastResponseWithDefaults instantiates a new UsageForecastResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageUsage

`func (o *UsageForecastResponse) GetAverageUsage() float64`

GetAverageUsage returns the AverageUsage field if non-nil, zero value otherwise.

### GetAverageUsageOk

`func (o *UsageForecastResponse) GetAverageUsageOk() (*float64, bool)`

GetAverageUsageOk returns a tuple with the AverageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageUsage

`func (o *UsageForecastResponse) SetAverageUsage(v float64)`

SetAverageUsage sets AverageUsage field to given value.

### HasAverageUsage

`func (o *UsageForecastResponse) HasAverageUsage() bool`

HasAverageUsage returns a boolean if a field has been set.

### GetUsagePercentage

`func (o *UsageForecastResponse) GetUsagePercentage() float64`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *UsageForecastResponse) GetUsagePercentageOk() (*float64, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *UsageForecastResponse) SetUsagePercentage(v float64)`

SetUsagePercentage sets UsagePercentage field to given value.

### HasUsagePercentage

`func (o *UsageForecastResponse) HasUsagePercentage() bool`

HasUsagePercentage returns a boolean if a field has been set.

### GetForecastedUsage

`func (o *UsageForecastResponse) GetForecastedUsage() float64`

GetForecastedUsage returns the ForecastedUsage field if non-nil, zero value otherwise.

### GetForecastedUsageOk

`func (o *UsageForecastResponse) GetForecastedUsageOk() (*float64, bool)`

GetForecastedUsageOk returns a tuple with the ForecastedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastedUsage

`func (o *UsageForecastResponse) SetForecastedUsage(v float64)`

SetForecastedUsage sets ForecastedUsage field to given value.

### HasForecastedUsage

`func (o *UsageForecastResponse) HasForecastedUsage() bool`

HasForecastedUsage returns a boolean if a field has been set.

### GetForecastedUsagePercentage

`func (o *UsageForecastResponse) GetForecastedUsagePercentage() float64`

GetForecastedUsagePercentage returns the ForecastedUsagePercentage field if non-nil, zero value otherwise.

### GetForecastedUsagePercentageOk

`func (o *UsageForecastResponse) GetForecastedUsagePercentageOk() (*float64, bool)`

GetForecastedUsagePercentageOk returns a tuple with the ForecastedUsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastedUsagePercentage

`func (o *UsageForecastResponse) SetForecastedUsagePercentage(v float64)`

SetForecastedUsagePercentage sets ForecastedUsagePercentage field to given value.

### HasForecastedUsagePercentage

`func (o *UsageForecastResponse) HasForecastedUsagePercentage() bool`

HasForecastedUsagePercentage returns a boolean if a field has been set.

### GetRemainingDays

`func (o *UsageForecastResponse) GetRemainingDays() float64`

GetRemainingDays returns the RemainingDays field if non-nil, zero value otherwise.

### GetRemainingDaysOk

`func (o *UsageForecastResponse) GetRemainingDaysOk() (*float64, bool)`

GetRemainingDaysOk returns a tuple with the RemainingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingDays

`func (o *UsageForecastResponse) SetRemainingDays(v float64)`

SetRemainingDays sets RemainingDays field to given value.

### HasRemainingDays

`func (o *UsageForecastResponse) HasRemainingDays() bool`

HasRemainingDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


