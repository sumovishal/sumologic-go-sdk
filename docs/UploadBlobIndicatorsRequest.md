# UploadBlobIndicatorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | User-provided text to identify the source of the indicator | 
**Format** | **string** | The input format. STIX_2_JSON includes support for minor versions (i.e. 2.0, 2.1) | 
**Blob** | **string** | The blob containing indicators. | 

## Methods

### NewUploadBlobIndicatorsRequest

`func NewUploadBlobIndicatorsRequest(source string, format string, blob string, ) *UploadBlobIndicatorsRequest`

NewUploadBlobIndicatorsRequest instantiates a new UploadBlobIndicatorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadBlobIndicatorsRequestWithDefaults

`func NewUploadBlobIndicatorsRequestWithDefaults() *UploadBlobIndicatorsRequest`

NewUploadBlobIndicatorsRequestWithDefaults instantiates a new UploadBlobIndicatorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *UploadBlobIndicatorsRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UploadBlobIndicatorsRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UploadBlobIndicatorsRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetFormat

`func (o *UploadBlobIndicatorsRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UploadBlobIndicatorsRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UploadBlobIndicatorsRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetBlob

`func (o *UploadBlobIndicatorsRequest) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *UploadBlobIndicatorsRequest) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *UploadBlobIndicatorsRequest) SetBlob(v string)`

SetBlob sets Blob field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


