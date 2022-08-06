# SlosLibrarySloResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalType** | **string** | Type of SLI Signal (latency, error, throughput, availability or other). | 
**Compliance** | [**Compliance**](Compliance.md) |  | 
**Indicator** | [**Sli**](Sli.md) |  | 
**Service** | Pointer to **string** | Name of the service. | [optional] 
**Application** | Pointer to **string** | Name of the application. | [optional] 
**SloVersion** | Pointer to **int64** | Current SLO Version. This is incremented on every change of a critical field of the SLO (i.e, SLI or Compliance period timezone), that requires a recompute of the SLI values over the compliance period. | [optional] 

## Methods

### NewSlosLibrarySloResponseAllOf

`func NewSlosLibrarySloResponseAllOf(signalType string, compliance Compliance, indicator Sli, ) *SlosLibrarySloResponseAllOf`

NewSlosLibrarySloResponseAllOf instantiates a new SlosLibrarySloResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlosLibrarySloResponseAllOfWithDefaults

`func NewSlosLibrarySloResponseAllOfWithDefaults() *SlosLibrarySloResponseAllOf`

NewSlosLibrarySloResponseAllOfWithDefaults instantiates a new SlosLibrarySloResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalType

`func (o *SlosLibrarySloResponseAllOf) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *SlosLibrarySloResponseAllOf) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *SlosLibrarySloResponseAllOf) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetCompliance

`func (o *SlosLibrarySloResponseAllOf) GetCompliance() Compliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *SlosLibrarySloResponseAllOf) GetComplianceOk() (*Compliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *SlosLibrarySloResponseAllOf) SetCompliance(v Compliance)`

SetCompliance sets Compliance field to given value.


### GetIndicator

`func (o *SlosLibrarySloResponseAllOf) GetIndicator() Sli`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *SlosLibrarySloResponseAllOf) GetIndicatorOk() (*Sli, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *SlosLibrarySloResponseAllOf) SetIndicator(v Sli)`

SetIndicator sets Indicator field to given value.


### GetService

`func (o *SlosLibrarySloResponseAllOf) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SlosLibrarySloResponseAllOf) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SlosLibrarySloResponseAllOf) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SlosLibrarySloResponseAllOf) HasService() bool`

HasService returns a boolean if a field has been set.

### GetApplication

`func (o *SlosLibrarySloResponseAllOf) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *SlosLibrarySloResponseAllOf) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *SlosLibrarySloResponseAllOf) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *SlosLibrarySloResponseAllOf) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetSloVersion

`func (o *SlosLibrarySloResponseAllOf) GetSloVersion() int64`

GetSloVersion returns the SloVersion field if non-nil, zero value otherwise.

### GetSloVersionOk

`func (o *SlosLibrarySloResponseAllOf) GetSloVersionOk() (*int64, bool)`

GetSloVersionOk returns a tuple with the SloVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloVersion

`func (o *SlosLibrarySloResponseAllOf) SetSloVersion(v int64)`

SetSloVersion sets SloVersion field to given value.

### HasSloVersion

`func (o *SlosLibrarySloResponseAllOf) HasSloVersion() bool`

HasSloVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


