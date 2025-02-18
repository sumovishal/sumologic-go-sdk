# TierEstimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | Pointer to **string** | Name of the data tier | [optional] 
**PerScanInBytes** | Pointer to **int64** | estimate data scanned per monitor scan in bytes | [optional] 
**PerDayInBytes** | Pointer to **int64** | estimate data scanned per day in bytes | [optional] 
**PerYearInBytes** | Pointer to **int64** | estimate data scanned per year in bytes | [optional] 
**TrainingScanInBytes** | Pointer to **int64** | one-time scan for log anomaly monitor | [optional] 

## Methods

### NewTierEstimate

`func NewTierEstimate() *TierEstimate`

NewTierEstimate instantiates a new TierEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierEstimateWithDefaults

`func NewTierEstimateWithDefaults() *TierEstimate`

NewTierEstimateWithDefaults instantiates a new TierEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *TierEstimate) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *TierEstimate) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *TierEstimate) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *TierEstimate) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetPerScanInBytes

`func (o *TierEstimate) GetPerScanInBytes() int64`

GetPerScanInBytes returns the PerScanInBytes field if non-nil, zero value otherwise.

### GetPerScanInBytesOk

`func (o *TierEstimate) GetPerScanInBytesOk() (*int64, bool)`

GetPerScanInBytesOk returns a tuple with the PerScanInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerScanInBytes

`func (o *TierEstimate) SetPerScanInBytes(v int64)`

SetPerScanInBytes sets PerScanInBytes field to given value.

### HasPerScanInBytes

`func (o *TierEstimate) HasPerScanInBytes() bool`

HasPerScanInBytes returns a boolean if a field has been set.

### GetPerDayInBytes

`func (o *TierEstimate) GetPerDayInBytes() int64`

GetPerDayInBytes returns the PerDayInBytes field if non-nil, zero value otherwise.

### GetPerDayInBytesOk

`func (o *TierEstimate) GetPerDayInBytesOk() (*int64, bool)`

GetPerDayInBytesOk returns a tuple with the PerDayInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerDayInBytes

`func (o *TierEstimate) SetPerDayInBytes(v int64)`

SetPerDayInBytes sets PerDayInBytes field to given value.

### HasPerDayInBytes

`func (o *TierEstimate) HasPerDayInBytes() bool`

HasPerDayInBytes returns a boolean if a field has been set.

### GetPerYearInBytes

`func (o *TierEstimate) GetPerYearInBytes() int64`

GetPerYearInBytes returns the PerYearInBytes field if non-nil, zero value otherwise.

### GetPerYearInBytesOk

`func (o *TierEstimate) GetPerYearInBytesOk() (*int64, bool)`

GetPerYearInBytesOk returns a tuple with the PerYearInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerYearInBytes

`func (o *TierEstimate) SetPerYearInBytes(v int64)`

SetPerYearInBytes sets PerYearInBytes field to given value.

### HasPerYearInBytes

`func (o *TierEstimate) HasPerYearInBytes() bool`

HasPerYearInBytes returns a boolean if a field has been set.

### GetTrainingScanInBytes

`func (o *TierEstimate) GetTrainingScanInBytes() int64`

GetTrainingScanInBytes returns the TrainingScanInBytes field if non-nil, zero value otherwise.

### GetTrainingScanInBytesOk

`func (o *TierEstimate) GetTrainingScanInBytesOk() (*int64, bool)`

GetTrainingScanInBytesOk returns a tuple with the TrainingScanInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingScanInBytes

`func (o *TierEstimate) SetTrainingScanInBytes(v int64)`

SetTrainingScanInBytes sets TrainingScanInBytes field to given value.

### HasTrainingScanInBytes

`func (o *TierEstimate) HasTrainingScanInBytes() bool`

HasTrainingScanInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


