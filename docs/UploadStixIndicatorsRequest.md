# UploadStixIndicatorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | User-provided text to identify the source of the indicator | 
**Indicators** | [**[]StixIndicator**](StixIndicator.md) | The list of stix threat intel indicators to upload. | 

## Methods

### NewUploadStixIndicatorsRequest

`func NewUploadStixIndicatorsRequest(source string, indicators []StixIndicator, ) *UploadStixIndicatorsRequest`

NewUploadStixIndicatorsRequest instantiates a new UploadStixIndicatorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadStixIndicatorsRequestWithDefaults

`func NewUploadStixIndicatorsRequestWithDefaults() *UploadStixIndicatorsRequest`

NewUploadStixIndicatorsRequestWithDefaults instantiates a new UploadStixIndicatorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *UploadStixIndicatorsRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UploadStixIndicatorsRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UploadStixIndicatorsRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetIndicators

`func (o *UploadStixIndicatorsRequest) GetIndicators() []StixIndicator`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *UploadStixIndicatorsRequest) GetIndicatorsOk() (*[]StixIndicator, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *UploadStixIndicatorsRequest) SetIndicators(v []StixIndicator)`

SetIndicators sets Indicators field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


