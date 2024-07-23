# ScanEstimateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | Pointer to **string** | Name of the tier for which usage is estimated. | [optional] 
**PerScanInBytes** | Pointer to **int64** | Amount of data scanned in bytes, to run the schedule search once. | [optional] 
**PerDayInBytes** | Pointer to **int64** | Amount of data scanned in bytes, to run the schedule search each day. | [optional] 
**PerYearInBytes** | Pointer to **int64** | Amount of data scanned in bytes, to run the schedule search each year. | [optional] 

## Methods

### NewScanEstimateDetails

`func NewScanEstimateDetails() *ScanEstimateDetails`

NewScanEstimateDetails instantiates a new ScanEstimateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanEstimateDetailsWithDefaults

`func NewScanEstimateDetailsWithDefaults() *ScanEstimateDetails`

NewScanEstimateDetailsWithDefaults instantiates a new ScanEstimateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *ScanEstimateDetails) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ScanEstimateDetails) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ScanEstimateDetails) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ScanEstimateDetails) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetPerScanInBytes

`func (o *ScanEstimateDetails) GetPerScanInBytes() int64`

GetPerScanInBytes returns the PerScanInBytes field if non-nil, zero value otherwise.

### GetPerScanInBytesOk

`func (o *ScanEstimateDetails) GetPerScanInBytesOk() (*int64, bool)`

GetPerScanInBytesOk returns a tuple with the PerScanInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerScanInBytes

`func (o *ScanEstimateDetails) SetPerScanInBytes(v int64)`

SetPerScanInBytes sets PerScanInBytes field to given value.

### HasPerScanInBytes

`func (o *ScanEstimateDetails) HasPerScanInBytes() bool`

HasPerScanInBytes returns a boolean if a field has been set.

### GetPerDayInBytes

`func (o *ScanEstimateDetails) GetPerDayInBytes() int64`

GetPerDayInBytes returns the PerDayInBytes field if non-nil, zero value otherwise.

### GetPerDayInBytesOk

`func (o *ScanEstimateDetails) GetPerDayInBytesOk() (*int64, bool)`

GetPerDayInBytesOk returns a tuple with the PerDayInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerDayInBytes

`func (o *ScanEstimateDetails) SetPerDayInBytes(v int64)`

SetPerDayInBytes sets PerDayInBytes field to given value.

### HasPerDayInBytes

`func (o *ScanEstimateDetails) HasPerDayInBytes() bool`

HasPerDayInBytes returns a boolean if a field has been set.

### GetPerYearInBytes

`func (o *ScanEstimateDetails) GetPerYearInBytes() int64`

GetPerYearInBytes returns the PerYearInBytes field if non-nil, zero value otherwise.

### GetPerYearInBytesOk

`func (o *ScanEstimateDetails) GetPerYearInBytesOk() (*int64, bool)`

GetPerYearInBytesOk returns a tuple with the PerYearInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerYearInBytes

`func (o *ScanEstimateDetails) SetPerYearInBytes(v int64)`

SetPerYearInBytes sets PerYearInBytes field to given value.

### HasPerYearInBytes

`func (o *ScanEstimateDetails) HasPerYearInBytes() bool`

HasPerYearInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


