# CollectorCompatibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | Pointer to **string** | Name of the Operating System. | [optional] 
**CollectorVersionRange** | Pointer to [**CollectorVersionRange**](CollectorVersionRange.md) |  | [optional] 

## Methods

### NewCollectorCompatibility

`func NewCollectorCompatibility() *CollectorCompatibility`

NewCollectorCompatibility instantiates a new CollectorCompatibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectorCompatibilityWithDefaults

`func NewCollectorCompatibilityWithDefaults() *CollectorCompatibility`

NewCollectorCompatibilityWithDefaults instantiates a new CollectorCompatibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *CollectorCompatibility) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CollectorCompatibility) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CollectorCompatibility) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CollectorCompatibility) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCollectorVersionRange

`func (o *CollectorCompatibility) GetCollectorVersionRange() CollectorVersionRange`

GetCollectorVersionRange returns the CollectorVersionRange field if non-nil, zero value otherwise.

### GetCollectorVersionRangeOk

`func (o *CollectorCompatibility) GetCollectorVersionRangeOk() (*CollectorVersionRange, bool)`

GetCollectorVersionRangeOk returns a tuple with the CollectorVersionRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorVersionRange

`func (o *CollectorCompatibility) SetCollectorVersionRange(v CollectorVersionRange)`

SetCollectorVersionRange sets CollectorVersionRange field to given value.

### HasCollectorVersionRange

`func (o *CollectorCompatibility) HasCollectorVersionRange() bool`

HasCollectorVersionRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


