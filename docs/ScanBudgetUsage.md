# ScanBudgetUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetId** | **string** | Budget id. | 
**Usage** | **int64** | Budget usage (in bytes). | 
**UsagePercentage** | **int64** | Budget usage percentage. | 

## Methods

### NewScanBudgetUsage

`func NewScanBudgetUsage(budgetId string, usage int64, usagePercentage int64, ) *ScanBudgetUsage`

NewScanBudgetUsage instantiates a new ScanBudgetUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanBudgetUsageWithDefaults

`func NewScanBudgetUsageWithDefaults() *ScanBudgetUsage`

NewScanBudgetUsageWithDefaults instantiates a new ScanBudgetUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetId

`func (o *ScanBudgetUsage) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *ScanBudgetUsage) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *ScanBudgetUsage) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetUsage

`func (o *ScanBudgetUsage) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ScanBudgetUsage) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ScanBudgetUsage) SetUsage(v int64)`

SetUsage sets Usage field to given value.


### GetUsagePercentage

`func (o *ScanBudgetUsage) GetUsagePercentage() int64`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *ScanBudgetUsage) GetUsagePercentageOk() (*int64, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *ScanBudgetUsage) SetUsagePercentage(v int64)`

SetUsagePercentage sets UsagePercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


