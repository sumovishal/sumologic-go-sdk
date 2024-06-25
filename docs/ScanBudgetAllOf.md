# ScanBudgetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the budget. | 
**OrgId** | **string** | Org Id of the org for the budget. | 
**LastModified** | **time.Time** | Date &amp; time when budget was last modified. | 

## Methods

### NewScanBudgetAllOf

`func NewScanBudgetAllOf(id string, orgId string, lastModified time.Time, ) *ScanBudgetAllOf`

NewScanBudgetAllOf instantiates a new ScanBudgetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanBudgetAllOfWithDefaults

`func NewScanBudgetAllOfWithDefaults() *ScanBudgetAllOf`

NewScanBudgetAllOfWithDefaults instantiates a new ScanBudgetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScanBudgetAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScanBudgetAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScanBudgetAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *ScanBudgetAllOf) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ScanBudgetAllOf) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ScanBudgetAllOf) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetLastModified

`func (o *ScanBudgetAllOf) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ScanBudgetAllOf) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ScanBudgetAllOf) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


