# TierEstimate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | Pointer to **string** | Name of the data tier | [optional] 
**PerDayInBytes** | Pointer to **int64** | estimate data scanned per day in bytes | [optional] 
**PerYearInBytes** | Pointer to **int64** | estimate data scanned per year in bytes | [optional] 

## Methods

### NewTierEstimate1

`func NewTierEstimate1() *TierEstimate1`

NewTierEstimate1 instantiates a new TierEstimate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierEstimate1WithDefaults

`func NewTierEstimate1WithDefaults() *TierEstimate1`

NewTierEstimate1WithDefaults instantiates a new TierEstimate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *TierEstimate1) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *TierEstimate1) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *TierEstimate1) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *TierEstimate1) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetPerDayInBytes

`func (o *TierEstimate1) GetPerDayInBytes() int64`

GetPerDayInBytes returns the PerDayInBytes field if non-nil, zero value otherwise.

### GetPerDayInBytesOk

`func (o *TierEstimate1) GetPerDayInBytesOk() (*int64, bool)`

GetPerDayInBytesOk returns a tuple with the PerDayInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerDayInBytes

`func (o *TierEstimate1) SetPerDayInBytes(v int64)`

SetPerDayInBytes sets PerDayInBytes field to given value.

### HasPerDayInBytes

`func (o *TierEstimate1) HasPerDayInBytes() bool`

HasPerDayInBytes returns a boolean if a field has been set.

### GetPerYearInBytes

`func (o *TierEstimate1) GetPerYearInBytes() int64`

GetPerYearInBytes returns the PerYearInBytes field if non-nil, zero value otherwise.

### GetPerYearInBytesOk

`func (o *TierEstimate1) GetPerYearInBytesOk() (*int64, bool)`

GetPerYearInBytesOk returns a tuple with the PerYearInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerYearInBytes

`func (o *TierEstimate1) SetPerYearInBytes(v int64)`

SetPerYearInBytes sets PerYearInBytes field to given value.

### HasPerYearInBytes

`func (o *TierEstimate1) HasPerYearInBytes() bool`

HasPerYearInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


