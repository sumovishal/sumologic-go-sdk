# SliStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status showing whether the SLI related metrics for the SLO were successfully computed or had an irrecoverable error during calculation or computation is in progress. | 
**SliPercentage** | Pointer to **float64** | SLI percentage for the compliance period. | [optional] 
**ErrorBudgetRemainingPercentage** | Pointer to **float64** | A string representing the error budget remaining in percentage. | [optional] 
**AbsoluteErrorBudgetRemaining** | Pointer to **string** | A string representing the absolute error budget remaining in terms of time or number of requests. | [optional] 
**ProgressPercentage** | Pointer to **float64** | The percentage progress of SLI metrics calculation in the system if SLI metrics are currently being computed. | [optional] 

## Methods

### NewSliStatus

`func NewSliStatus(status string, ) *SliStatus`

NewSliStatus instantiates a new SliStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliStatusWithDefaults

`func NewSliStatusWithDefaults() *SliStatus`

NewSliStatusWithDefaults instantiates a new SliStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SliStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SliStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SliStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSliPercentage

`func (o *SliStatus) GetSliPercentage() float64`

GetSliPercentage returns the SliPercentage field if non-nil, zero value otherwise.

### GetSliPercentageOk

`func (o *SliStatus) GetSliPercentageOk() (*float64, bool)`

GetSliPercentageOk returns a tuple with the SliPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliPercentage

`func (o *SliStatus) SetSliPercentage(v float64)`

SetSliPercentage sets SliPercentage field to given value.

### HasSliPercentage

`func (o *SliStatus) HasSliPercentage() bool`

HasSliPercentage returns a boolean if a field has been set.

### GetErrorBudgetRemainingPercentage

`func (o *SliStatus) GetErrorBudgetRemainingPercentage() float64`

GetErrorBudgetRemainingPercentage returns the ErrorBudgetRemainingPercentage field if non-nil, zero value otherwise.

### GetErrorBudgetRemainingPercentageOk

`func (o *SliStatus) GetErrorBudgetRemainingPercentageOk() (*float64, bool)`

GetErrorBudgetRemainingPercentageOk returns a tuple with the ErrorBudgetRemainingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBudgetRemainingPercentage

`func (o *SliStatus) SetErrorBudgetRemainingPercentage(v float64)`

SetErrorBudgetRemainingPercentage sets ErrorBudgetRemainingPercentage field to given value.

### HasErrorBudgetRemainingPercentage

`func (o *SliStatus) HasErrorBudgetRemainingPercentage() bool`

HasErrorBudgetRemainingPercentage returns a boolean if a field has been set.

### GetAbsoluteErrorBudgetRemaining

`func (o *SliStatus) GetAbsoluteErrorBudgetRemaining() string`

GetAbsoluteErrorBudgetRemaining returns the AbsoluteErrorBudgetRemaining field if non-nil, zero value otherwise.

### GetAbsoluteErrorBudgetRemainingOk

`func (o *SliStatus) GetAbsoluteErrorBudgetRemainingOk() (*string, bool)`

GetAbsoluteErrorBudgetRemainingOk returns a tuple with the AbsoluteErrorBudgetRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteErrorBudgetRemaining

`func (o *SliStatus) SetAbsoluteErrorBudgetRemaining(v string)`

SetAbsoluteErrorBudgetRemaining sets AbsoluteErrorBudgetRemaining field to given value.

### HasAbsoluteErrorBudgetRemaining

`func (o *SliStatus) HasAbsoluteErrorBudgetRemaining() bool`

HasAbsoluteErrorBudgetRemaining returns a boolean if a field has been set.

### GetProgressPercentage

`func (o *SliStatus) GetProgressPercentage() float64`

GetProgressPercentage returns the ProgressPercentage field if non-nil, zero value otherwise.

### GetProgressPercentageOk

`func (o *SliStatus) GetProgressPercentageOk() (*float64, bool)`

GetProgressPercentageOk returns a tuple with the ProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercentage

`func (o *SliStatus) SetProgressPercentage(v float64)`

SetProgressPercentage sets ProgressPercentage field to given value.

### HasProgressPercentage

`func (o *SliStatus) HasProgressPercentage() bool`

HasProgressPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


