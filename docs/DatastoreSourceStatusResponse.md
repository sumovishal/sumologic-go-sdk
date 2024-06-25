# DatastoreSourceStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The source name | 
**Description** | Pointer to **string** | The source description | [optional] 
**DiskSize** | Pointer to **int64** | Disk utilization in bytes estimate for the indicator source | [optional] 
**IndicatorCount** | Pointer to **int64** | Number of indicators for the indicator source | [optional] 
**SumoProvided** | Pointer to **bool** | True if sumo provided source | [optional] 
**SupportsCat** | Pointer to **bool** | True if can be used in cat operator | [optional] 
**Enabled** | Pointer to **bool** | True if enabled | [optional] 

## Methods

### NewDatastoreSourceStatusResponse

`func NewDatastoreSourceStatusResponse(source string, ) *DatastoreSourceStatusResponse`

NewDatastoreSourceStatusResponse instantiates a new DatastoreSourceStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreSourceStatusResponseWithDefaults

`func NewDatastoreSourceStatusResponseWithDefaults() *DatastoreSourceStatusResponse`

NewDatastoreSourceStatusResponseWithDefaults instantiates a new DatastoreSourceStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *DatastoreSourceStatusResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DatastoreSourceStatusResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DatastoreSourceStatusResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetDescription

`func (o *DatastoreSourceStatusResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatastoreSourceStatusResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatastoreSourceStatusResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatastoreSourceStatusResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskSize

`func (o *DatastoreSourceStatusResponse) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *DatastoreSourceStatusResponse) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *DatastoreSourceStatusResponse) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *DatastoreSourceStatusResponse) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetIndicatorCount

`func (o *DatastoreSourceStatusResponse) GetIndicatorCount() int64`

GetIndicatorCount returns the IndicatorCount field if non-nil, zero value otherwise.

### GetIndicatorCountOk

`func (o *DatastoreSourceStatusResponse) GetIndicatorCountOk() (*int64, bool)`

GetIndicatorCountOk returns a tuple with the IndicatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorCount

`func (o *DatastoreSourceStatusResponse) SetIndicatorCount(v int64)`

SetIndicatorCount sets IndicatorCount field to given value.

### HasIndicatorCount

`func (o *DatastoreSourceStatusResponse) HasIndicatorCount() bool`

HasIndicatorCount returns a boolean if a field has been set.

### GetSumoProvided

`func (o *DatastoreSourceStatusResponse) GetSumoProvided() bool`

GetSumoProvided returns the SumoProvided field if non-nil, zero value otherwise.

### GetSumoProvidedOk

`func (o *DatastoreSourceStatusResponse) GetSumoProvidedOk() (*bool, bool)`

GetSumoProvidedOk returns a tuple with the SumoProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumoProvided

`func (o *DatastoreSourceStatusResponse) SetSumoProvided(v bool)`

SetSumoProvided sets SumoProvided field to given value.

### HasSumoProvided

`func (o *DatastoreSourceStatusResponse) HasSumoProvided() bool`

HasSumoProvided returns a boolean if a field has been set.

### GetSupportsCat

`func (o *DatastoreSourceStatusResponse) GetSupportsCat() bool`

GetSupportsCat returns the SupportsCat field if non-nil, zero value otherwise.

### GetSupportsCatOk

`func (o *DatastoreSourceStatusResponse) GetSupportsCatOk() (*bool, bool)`

GetSupportsCatOk returns a tuple with the SupportsCat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCat

`func (o *DatastoreSourceStatusResponse) SetSupportsCat(v bool)`

SetSupportsCat sets SupportsCat field to given value.

### HasSupportsCat

`func (o *DatastoreSourceStatusResponse) HasSupportsCat() bool`

HasSupportsCat returns a boolean if a field has been set.

### GetEnabled

`func (o *DatastoreSourceStatusResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DatastoreSourceStatusResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DatastoreSourceStatusResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DatastoreSourceStatusResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


