# ScheduledSearchEstimatedUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanEstimates** | Pointer to [**[]ScanEstimateDetails**](ScanEstimateDetails.md) | Scan estimate detail for a particular tier. | [optional] 

## Methods

### NewScheduledSearchEstimatedUsageResponse

`func NewScheduledSearchEstimatedUsageResponse() *ScheduledSearchEstimatedUsageResponse`

NewScheduledSearchEstimatedUsageResponse instantiates a new ScheduledSearchEstimatedUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledSearchEstimatedUsageResponseWithDefaults

`func NewScheduledSearchEstimatedUsageResponseWithDefaults() *ScheduledSearchEstimatedUsageResponse`

NewScheduledSearchEstimatedUsageResponseWithDefaults instantiates a new ScheduledSearchEstimatedUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanEstimates

`func (o *ScheduledSearchEstimatedUsageResponse) GetScanEstimates() []ScanEstimateDetails`

GetScanEstimates returns the ScanEstimates field if non-nil, zero value otherwise.

### GetScanEstimatesOk

`func (o *ScheduledSearchEstimatedUsageResponse) GetScanEstimatesOk() (*[]ScanEstimateDetails, bool)`

GetScanEstimatesOk returns a tuple with the ScanEstimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanEstimates

`func (o *ScheduledSearchEstimatedUsageResponse) SetScanEstimates(v []ScanEstimateDetails)`

SetScanEstimates sets ScanEstimates field to given value.

### HasScanEstimates

`func (o *ScheduledSearchEstimatedUsageResponse) HasScanEstimates() bool`

HasScanEstimates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


