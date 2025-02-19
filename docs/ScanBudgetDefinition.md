# ScanBudgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the budget. | 
**Capacity** | **int64** | Capacity of the budget. | 
**Unit** | **string** | Unit of the budget. | 
**BudgetType** | **string** | Type of the budget. | 
**Scope** | [**ScanBudgetScope**](ScanBudgetScope.md) |  | 
**Window** | **string** | Window of the budget. Use Daily/Weekly/Monthly for creating a time based budget (beta) | 
**ApplicableOn** | **string** | Grouping of the budget. | 
**GroupBy** | **string** | Grouping Entity of the budget. | 
**Action** | **string** | Action to be taken if the budget is breached | 
**Status** | Pointer to **string** | Signifies the state of the budget. (Active/Inactive) | [optional] 

## Methods

### NewScanBudgetDefinition

`func NewScanBudgetDefinition(name string, capacity int64, unit string, budgetType string, scope ScanBudgetScope, window string, applicableOn string, groupBy string, action string, ) *ScanBudgetDefinition`

NewScanBudgetDefinition instantiates a new ScanBudgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanBudgetDefinitionWithDefaults

`func NewScanBudgetDefinitionWithDefaults() *ScanBudgetDefinition`

NewScanBudgetDefinitionWithDefaults instantiates a new ScanBudgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScanBudgetDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScanBudgetDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScanBudgetDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetCapacity

`func (o *ScanBudgetDefinition) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ScanBudgetDefinition) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ScanBudgetDefinition) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.


### GetUnit

`func (o *ScanBudgetDefinition) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ScanBudgetDefinition) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ScanBudgetDefinition) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetBudgetType

`func (o *ScanBudgetDefinition) GetBudgetType() string`

GetBudgetType returns the BudgetType field if non-nil, zero value otherwise.

### GetBudgetTypeOk

`func (o *ScanBudgetDefinition) GetBudgetTypeOk() (*string, bool)`

GetBudgetTypeOk returns a tuple with the BudgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetType

`func (o *ScanBudgetDefinition) SetBudgetType(v string)`

SetBudgetType sets BudgetType field to given value.


### GetScope

`func (o *ScanBudgetDefinition) GetScope() ScanBudgetScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ScanBudgetDefinition) GetScopeOk() (*ScanBudgetScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ScanBudgetDefinition) SetScope(v ScanBudgetScope)`

SetScope sets Scope field to given value.


### GetWindow

`func (o *ScanBudgetDefinition) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ScanBudgetDefinition) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ScanBudgetDefinition) SetWindow(v string)`

SetWindow sets Window field to given value.


### GetApplicableOn

`func (o *ScanBudgetDefinition) GetApplicableOn() string`

GetApplicableOn returns the ApplicableOn field if non-nil, zero value otherwise.

### GetApplicableOnOk

`func (o *ScanBudgetDefinition) GetApplicableOnOk() (*string, bool)`

GetApplicableOnOk returns a tuple with the ApplicableOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableOn

`func (o *ScanBudgetDefinition) SetApplicableOn(v string)`

SetApplicableOn sets ApplicableOn field to given value.


### GetGroupBy

`func (o *ScanBudgetDefinition) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ScanBudgetDefinition) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ScanBudgetDefinition) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.


### GetAction

`func (o *ScanBudgetDefinition) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ScanBudgetDefinition) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ScanBudgetDefinition) SetAction(v string)`

SetAction sets Action field to given value.


### GetStatus

`func (o *ScanBudgetDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScanBudgetDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScanBudgetDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScanBudgetDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


