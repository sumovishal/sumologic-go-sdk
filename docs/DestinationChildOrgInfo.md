# DestinationChildOrgInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | [**[]ChildOrgInfo**](ChildOrgInfo.md) | Organization Info which needs to be included in Destination Organisation List. | 
**Excluded** | [**[]ChildOrgInfo**](ChildOrgInfo.md) | Organization Info which needs to be excluded from Destination Organisation List. | 

## Methods

### NewDestinationChildOrgInfo

`func NewDestinationChildOrgInfo(included []ChildOrgInfo, excluded []ChildOrgInfo, ) *DestinationChildOrgInfo`

NewDestinationChildOrgInfo instantiates a new DestinationChildOrgInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationChildOrgInfoWithDefaults

`func NewDestinationChildOrgInfoWithDefaults() *DestinationChildOrgInfo`

NewDestinationChildOrgInfoWithDefaults instantiates a new DestinationChildOrgInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *DestinationChildOrgInfo) GetIncluded() []ChildOrgInfo`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *DestinationChildOrgInfo) GetIncludedOk() (*[]ChildOrgInfo, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *DestinationChildOrgInfo) SetIncluded(v []ChildOrgInfo)`

SetIncluded sets Included field to given value.


### GetExcluded

`func (o *DestinationChildOrgInfo) GetExcluded() []ChildOrgInfo`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *DestinationChildOrgInfo) GetExcludedOk() (*[]ChildOrgInfo, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *DestinationChildOrgInfo) SetExcluded(v []ChildOrgInfo)`

SetExcluded sets Excluded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


