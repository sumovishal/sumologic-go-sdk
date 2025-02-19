# ScanBudget

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
**Id** | **string** | Id of the budget. | 
**OrgId** | **string** | Org Id of the org for the budget. | 
**ResetTime** | **string** | Reset time of the time based scan budget in HH:MM format | [default to "00:00"]
**ResetTimeZone** | **string** | Time zone of the reset time for the time based scan budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | [default to "Etc/UTC"]
**ResetDayOfWeek** | **string** | The day of the week when the budget resets, applicable for time based budgets with a Weekly window. Must be a valid day of the week. | [default to "MONDAY"]
**ResetDateOfMonth** | **int32** | The date of the month when the budget resets, applicable for time based budgets with a Monthly window. Must be a valid day of the month (1-28). | [default to 1]
**CreatedAt** | **time.Time** | Date &amp; time when budget was created. | 
**CreatedBy** | **string** | Id of the user who created the budget. | 
**ModifiedAt** | **time.Time** | Date &amp; time when budget was last modified. | 
**ModifiedBy** | **string** | Id of the user who last modified the budget. | 

## Methods

### NewScanBudget

`func NewScanBudget(name string, capacity int64, unit string, budgetType string, scope ScanBudgetScope, window string, applicableOn string, groupBy string, action string, id string, orgId string, resetTime string, resetTimeZone string, resetDayOfWeek string, resetDateOfMonth int32, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, ) *ScanBudget`

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


### GetAction

`func (o *ScanBudget) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ScanBudget) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ScanBudget) SetAction(v string)`

SetAction sets Action field to given value.


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


### GetResetTime

`func (o *ScanBudget) GetResetTime() string`

GetResetTime returns the ResetTime field if non-nil, zero value otherwise.

### GetResetTimeOk

`func (o *ScanBudget) GetResetTimeOk() (*string, bool)`

GetResetTimeOk returns a tuple with the ResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTime

`func (o *ScanBudget) SetResetTime(v string)`

SetResetTime sets ResetTime field to given value.


### GetResetTimeZone

`func (o *ScanBudget) GetResetTimeZone() string`

GetResetTimeZone returns the ResetTimeZone field if non-nil, zero value otherwise.

### GetResetTimeZoneOk

`func (o *ScanBudget) GetResetTimeZoneOk() (*string, bool)`

GetResetTimeZoneOk returns a tuple with the ResetTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTimeZone

`func (o *ScanBudget) SetResetTimeZone(v string)`

SetResetTimeZone sets ResetTimeZone field to given value.


### GetResetDayOfWeek

`func (o *ScanBudget) GetResetDayOfWeek() string`

GetResetDayOfWeek returns the ResetDayOfWeek field if non-nil, zero value otherwise.

### GetResetDayOfWeekOk

`func (o *ScanBudget) GetResetDayOfWeekOk() (*string, bool)`

GetResetDayOfWeekOk returns a tuple with the ResetDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDayOfWeek

`func (o *ScanBudget) SetResetDayOfWeek(v string)`

SetResetDayOfWeek sets ResetDayOfWeek field to given value.


### GetResetDateOfMonth

`func (o *ScanBudget) GetResetDateOfMonth() int32`

GetResetDateOfMonth returns the ResetDateOfMonth field if non-nil, zero value otherwise.

### GetResetDateOfMonthOk

`func (o *ScanBudget) GetResetDateOfMonthOk() (*int32, bool)`

GetResetDateOfMonthOk returns a tuple with the ResetDateOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDateOfMonth

`func (o *ScanBudget) SetResetDateOfMonth(v int32)`

SetResetDateOfMonth sets ResetDateOfMonth field to given value.


### GetCreatedAt

`func (o *ScanBudget) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScanBudget) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScanBudget) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ScanBudget) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScanBudget) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScanBudget) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *ScanBudget) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ScanBudget) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ScanBudget) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ScanBudget) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ScanBudget) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ScanBudget) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


