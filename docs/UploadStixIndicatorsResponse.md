# UploadStixIndicatorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidIndicators** | **[]string** | A list of invalid indicator IDs that were not ingested | 

## Methods

### NewUploadStixIndicatorsResponse

`func NewUploadStixIndicatorsResponse(invalidIndicators []string, ) *UploadStixIndicatorsResponse`

NewUploadStixIndicatorsResponse instantiates a new UploadStixIndicatorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadStixIndicatorsResponseWithDefaults

`func NewUploadStixIndicatorsResponseWithDefaults() *UploadStixIndicatorsResponse`

NewUploadStixIndicatorsResponseWithDefaults instantiates a new UploadStixIndicatorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidIndicators

`func (o *UploadStixIndicatorsResponse) GetInvalidIndicators() []string`

GetInvalidIndicators returns the InvalidIndicators field if non-nil, zero value otherwise.

### GetInvalidIndicatorsOk

`func (o *UploadStixIndicatorsResponse) GetInvalidIndicatorsOk() (*[]string, bool)`

GetInvalidIndicatorsOk returns a tuple with the InvalidIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidIndicators

`func (o *UploadStixIndicatorsResponse) SetInvalidIndicators(v []string)`

SetInvalidIndicators sets InvalidIndicators field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


