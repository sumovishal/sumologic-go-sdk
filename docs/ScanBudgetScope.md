# ScanBudgetScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedUsers** | **[]string** | List of userIds included in the budget. | 
**ExcludedUsers** | **[]string** | List of userIds excluded in the budget. | 
**IncludedRoles** | **[]string** | List of roleIds included in the budget. | 
**ExcludedRoles** | **[]string** | List of roleIds excluded in the budget. | 

## Methods

### NewScanBudgetScope

`func NewScanBudgetScope(includedUsers []string, excludedUsers []string, includedRoles []string, excludedRoles []string, ) *ScanBudgetScope`

NewScanBudgetScope instantiates a new ScanBudgetScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanBudgetScopeWithDefaults

`func NewScanBudgetScopeWithDefaults() *ScanBudgetScope`

NewScanBudgetScopeWithDefaults instantiates a new ScanBudgetScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedUsers

`func (o *ScanBudgetScope) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *ScanBudgetScope) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *ScanBudgetScope) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.


### GetExcludedUsers

`func (o *ScanBudgetScope) GetExcludedUsers() []string`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *ScanBudgetScope) GetExcludedUsersOk() (*[]string, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *ScanBudgetScope) SetExcludedUsers(v []string)`

SetExcludedUsers sets ExcludedUsers field to given value.


### GetIncludedRoles

`func (o *ScanBudgetScope) GetIncludedRoles() []string`

GetIncludedRoles returns the IncludedRoles field if non-nil, zero value otherwise.

### GetIncludedRolesOk

`func (o *ScanBudgetScope) GetIncludedRolesOk() (*[]string, bool)`

GetIncludedRolesOk returns a tuple with the IncludedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedRoles

`func (o *ScanBudgetScope) SetIncludedRoles(v []string)`

SetIncludedRoles sets IncludedRoles field to given value.


### GetExcludedRoles

`func (o *ScanBudgetScope) GetExcludedRoles() []string`

GetExcludedRoles returns the ExcludedRoles field if non-nil, zero value otherwise.

### GetExcludedRolesOk

`func (o *ScanBudgetScope) GetExcludedRolesOk() (*[]string, bool)`

GetExcludedRolesOk returns a tuple with the ExcludedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedRoles

`func (o *ScanBudgetScope) SetExcludedRoles(v []string)`

SetExcludedRoles sets ExcludedRoles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


