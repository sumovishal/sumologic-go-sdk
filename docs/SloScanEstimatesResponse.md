# SloScanEstimatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanEstimates** | Pointer to [**[]TierEstimate1**](TierEstimate1.md) | Scan estimates by data tier for the SLO configuration. | [optional] 

## Methods

### NewSloScanEstimatesResponse

`func NewSloScanEstimatesResponse() *SloScanEstimatesResponse`

NewSloScanEstimatesResponse instantiates a new SloScanEstimatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloScanEstimatesResponseWithDefaults

`func NewSloScanEstimatesResponseWithDefaults() *SloScanEstimatesResponse`

NewSloScanEstimatesResponseWithDefaults instantiates a new SloScanEstimatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanEstimates

`func (o *SloScanEstimatesResponse) GetScanEstimates() []TierEstimate1`

GetScanEstimates returns the ScanEstimates field if non-nil, zero value otherwise.

### GetScanEstimatesOk

`func (o *SloScanEstimatesResponse) GetScanEstimatesOk() (*[]TierEstimate1, bool)`

GetScanEstimatesOk returns a tuple with the ScanEstimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanEstimates

`func (o *SloScanEstimatesResponse) SetScanEstimates(v []TierEstimate1)`

SetScanEstimates sets ScanEstimates field to given value.

### HasScanEstimates

`func (o *SloScanEstimatesResponse) HasScanEstimates() bool`

HasScanEstimates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


