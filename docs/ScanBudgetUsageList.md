# ScanBudgetUsageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ScanBudgetUsage**](ScanBudgetUsage.md) | List of budget usages | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewScanBudgetUsageList

`func NewScanBudgetUsageList(data []ScanBudgetUsage, ) *ScanBudgetUsageList`

NewScanBudgetUsageList instantiates a new ScanBudgetUsageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanBudgetUsageListWithDefaults

`func NewScanBudgetUsageListWithDefaults() *ScanBudgetUsageList`

NewScanBudgetUsageListWithDefaults instantiates a new ScanBudgetUsageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ScanBudgetUsageList) GetData() []ScanBudgetUsage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ScanBudgetUsageList) GetDataOk() (*[]ScanBudgetUsage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ScanBudgetUsageList) SetData(v []ScanBudgetUsage)`

SetData sets Data field to given value.


### GetNext

`func (o *ScanBudgetUsageList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ScanBudgetUsageList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ScanBudgetUsageList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ScanBudgetUsageList) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


