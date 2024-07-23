# VersionRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinVersion** | Pointer to **string** | Minimum version of otCollector. | [optional] 
**MaxVersion** | Pointer to **string** | Maximum version of the collector. | [optional] 
**RangeType** | Pointer to **string** | Specifies how filtering should be applied when &#x60;minVersion&#x60; and &#x60;maxVersion&#x60; are defined. - &#x60;Within&#x60;: Filtering includes the specified range. - &#x60;Outside&#x60;: Filtering excludes the specified range. By default, filtering includes the specified range. | [optional] 

## Methods

### NewVersionRange

`func NewVersionRange() *VersionRange`

NewVersionRange instantiates a new VersionRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionRangeWithDefaults

`func NewVersionRangeWithDefaults() *VersionRange`

NewVersionRangeWithDefaults instantiates a new VersionRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinVersion

`func (o *VersionRange) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *VersionRange) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *VersionRange) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *VersionRange) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetMaxVersion

`func (o *VersionRange) GetMaxVersion() string`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *VersionRange) GetMaxVersionOk() (*string, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *VersionRange) SetMaxVersion(v string)`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *VersionRange) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### GetRangeType

`func (o *VersionRange) GetRangeType() string`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *VersionRange) GetRangeTypeOk() (*string, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *VersionRange) SetRangeType(v string)`

SetRangeType sets RangeType field to given value.

### HasRangeType

`func (o *VersionRange) HasRangeType() bool`

HasRangeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


