# CollectorVersionRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinVersion** | Pointer to **string** | Minimum compatible version of otCollector. | [optional] 
**MaxVersion** | Pointer to **NullableString** | Maximum compatible version of the otcollector. if this is null, then latest otcollector is also compatible. | [optional] 

## Methods

### NewCollectorVersionRange

`func NewCollectorVersionRange() *CollectorVersionRange`

NewCollectorVersionRange instantiates a new CollectorVersionRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectorVersionRangeWithDefaults

`func NewCollectorVersionRangeWithDefaults() *CollectorVersionRange`

NewCollectorVersionRangeWithDefaults instantiates a new CollectorVersionRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinVersion

`func (o *CollectorVersionRange) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *CollectorVersionRange) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *CollectorVersionRange) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *CollectorVersionRange) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetMaxVersion

`func (o *CollectorVersionRange) GetMaxVersion() string`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *CollectorVersionRange) GetMaxVersionOk() (*string, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *CollectorVersionRange) SetMaxVersion(v string)`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *CollectorVersionRange) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### SetMaxVersionNil

`func (o *CollectorVersionRange) SetMaxVersionNil(b bool)`

 SetMaxVersionNil sets the value for MaxVersion to be an explicit nil

### UnsetMaxVersion
`func (o *CollectorVersionRange) UnsetMaxVersion()`

UnsetMaxVersion ensures that no value is present for MaxVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


