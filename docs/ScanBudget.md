# ScanBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the budget. | 
**Capacity** | **int64** | Capacity of the budget. | 
**Unit** | **string** | Unit of the budget. | 
**BudgetType** | **string** | Type of the budget. | 
**Scope** | [**ScanBudgetScope**](ScanBudgetScope.md) |  | 
**Window** | **string** | Window of the budget. | 
**ApplicableOn** | **string** | Grouping of the budget. | 
**GroupBy** | **string** | Grouping Entity of the budget. | 
**Actions** | **[]string** | Actions allowed in the budget. | 
**Status** | Pointer to **string** | Signifies the state of the budget. (Active/Inactive) | [optional] 
**Id** | **string** | Id of the budget. | 
**OrgId** | **string** | Org Id of the org for the budget. | 
**LastModified** | **time.Time** | Date &amp; time when budget was last modified. | 

## Methods

### NewScanBudget

`func NewScanBudget(name string, capacity int64, unit string, budgetType string, scope ScanBudgetScope, window string, applicableOn string, groupBy string, actions []string, id string, orgId string, lastModified time.Time, ) *ScanBudget`

NewScanBudget instantiates a new ScanBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanBudgetWithDefaults

`func NewScanBudgetWithDefaults() *ScanBudget`

NewScanBudgetWithDefaults instantiates a new ScanBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScanBudget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScanBudget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScanBudget) SetName(v string)`

SetName sets Name field to given value.


### GetCapacity

`func (o *ScanBudget) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ScanBudget) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ScanBudget) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.


### GetUnit

`func (o *ScanBudget) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ScanBudget) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ScanBudget) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetBudgetType

`func (o *ScanBudget) GetBudgetType() string`

GetBudgetType returns the BudgetType field if non-nil, zero value otherwise.

### GetBudgetTypeOk

`func (o *ScanBudget) GetBudgetTypeOk() (*string, bool)`

GetBudgetTypeOk returns a tuple with the BudgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetType

`func (o *ScanBudget) SetBudgetType(v string)`

SetBudgetType sets BudgetType field to given value.


### GetScope

`func (o *ScanBudget) GetScope() ScanBudgetScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ScanBudget) GetScopeOk() (*ScanBudgetScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ScanBudget) SetScope(v ScanBudgetScope)`

SetScope sets Scope field to given value.


### GetWindow

`func (o *ScanBudget) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ScanBudget) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ScanBudget) SetWindow(v string)`

SetWindow sets Window field to given value.


### GetApplicableOn

`func (o *ScanBudget) GetApplicableOn() string`

GetApplicableOn returns the ApplicableOn field if non-nil, zero value otherwise.

### GetApplicableOnOk

`func (o *ScanBudget) GetApplicableOnOk() (*string, bool)`

GetApplicableOnOk returns a tuple with the ApplicableOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableOn

`func (o *ScanBudget) SetApplicableOn(v string)`

SetApplicableOn sets ApplicableOn field to given value.


### GetGroupBy

`func (o *ScanBudget) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ScanBudget) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ScanBudget) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.


### GetActions

`func (o *ScanBudget) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ScanBudget) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ScanBudget) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetStatus

`func (o *ScanBudget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScanBudget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScanBudget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScanBudget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *ScanBudget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScanBudget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScanBudget) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *ScanBudget) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ScanBudget) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ScanBudget) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetLastModified

`func (o *ScanBudget) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ScanBudget) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ScanBudget) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


