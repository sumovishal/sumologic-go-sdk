# ScanBudgetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ScanBudget**](ScanBudget.md) | List of scan budgets. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewScanBudgetList

`func NewScanBudgetList(data []ScanBudget, ) *ScanBudgetList`

NewScanBudgetList instantiates a new ScanBudgetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanBudgetListWithDefaults

`func NewScanBudgetListWithDefaults() *ScanBudgetList`

NewScanBudgetListWithDefaults instantiates a new ScanBudgetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ScanBudgetList) GetData() []ScanBudget`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ScanBudgetList) GetDataOk() (*[]ScanBudget, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ScanBudgetList) SetData(v []ScanBudget)`

SetData sets Data field to given value.


### GetNext

`func (o *ScanBudgetList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ScanBudgetList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ScanBudgetList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ScanBudgetList) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


