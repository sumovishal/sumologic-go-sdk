# EstimatedUsageDetailsWithMeteringType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeteringType** | Pointer to **string** | Name of the metering type. Metering type indicates how the data scanned within a particular data tier is actually metered and billed. Supported Values are Continuous, Frequent, Infrequent, ContinuousSecurity and FlexSecurity.  | [optional] 
**DataScannedInBytes** | Pointer to **int64** | Amount of data scanned in bytes, to run the query. | [optional] 
**Tier** | Pointer to **string** | Name of the data tier. Supported Values are Continuous, Frequent, Infrequent and Flex. | [optional] 
**ScanCreditAccounted** | Pointer to **bool** | Whether particular metering type is accounted against a customer&#39;s credit on a per scan basis.  e.g Data belonging to \&quot;Flex\&quot; and \&quot;Infrequent\&quot; metering type is accounted for credits on per scan basis. For other metering types, eg. \&quot;Continuous\&quot; it&#39;s charged upfront during ingestion.  | [optional] 

## Methods

### NewEstimatedUsageDetailsWithMeteringType

`func NewEstimatedUsageDetailsWithMeteringType() *EstimatedUsageDetailsWithMeteringType`

NewEstimatedUsageDetailsWithMeteringType instantiates a new EstimatedUsageDetailsWithMeteringType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedUsageDetailsWithMeteringTypeWithDefaults

`func NewEstimatedUsageDetailsWithMeteringTypeWithDefaults() *EstimatedUsageDetailsWithMeteringType`

NewEstimatedUsageDetailsWithMeteringTypeWithDefaults instantiates a new EstimatedUsageDetailsWithMeteringType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeteringType

`func (o *EstimatedUsageDetailsWithMeteringType) GetMeteringType() string`

GetMeteringType returns the MeteringType field if non-nil, zero value otherwise.

### GetMeteringTypeOk

`func (o *EstimatedUsageDetailsWithMeteringType) GetMeteringTypeOk() (*string, bool)`

GetMeteringTypeOk returns a tuple with the MeteringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringType

`func (o *EstimatedUsageDetailsWithMeteringType) SetMeteringType(v string)`

SetMeteringType sets MeteringType field to given value.

### HasMeteringType

`func (o *EstimatedUsageDetailsWithMeteringType) HasMeteringType() bool`

HasMeteringType returns a boolean if a field has been set.

### GetDataScannedInBytes

`func (o *EstimatedUsageDetailsWithMeteringType) GetDataScannedInBytes() int64`

GetDataScannedInBytes returns the DataScannedInBytes field if non-nil, zero value otherwise.

### GetDataScannedInBytesOk

`func (o *EstimatedUsageDetailsWithMeteringType) GetDataScannedInBytesOk() (*int64, bool)`

GetDataScannedInBytesOk returns a tuple with the DataScannedInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataScannedInBytes

`func (o *EstimatedUsageDetailsWithMeteringType) SetDataScannedInBytes(v int64)`

SetDataScannedInBytes sets DataScannedInBytes field to given value.

### HasDataScannedInBytes

`func (o *EstimatedUsageDetailsWithMeteringType) HasDataScannedInBytes() bool`

HasDataScannedInBytes returns a boolean if a field has been set.

### GetTier

`func (o *EstimatedUsageDetailsWithMeteringType) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *EstimatedUsageDetailsWithMeteringType) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *EstimatedUsageDetailsWithMeteringType) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *EstimatedUsageDetailsWithMeteringType) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetScanCreditAccounted

`func (o *EstimatedUsageDetailsWithMeteringType) GetScanCreditAccounted() bool`

GetScanCreditAccounted returns the ScanCreditAccounted field if non-nil, zero value otherwise.

### GetScanCreditAccountedOk

`func (o *EstimatedUsageDetailsWithMeteringType) GetScanCreditAccountedOk() (*bool, bool)`

GetScanCreditAccountedOk returns a tuple with the ScanCreditAccounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanCreditAccounted

`func (o *EstimatedUsageDetailsWithMeteringType) SetScanCreditAccounted(v bool)`

SetScanCreditAccounted sets ScanCreditAccounted field to given value.

### HasScanCreditAccounted

`func (o *EstimatedUsageDetailsWithMeteringType) HasScanCreditAccounted() bool`

HasScanCreditAccounted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


