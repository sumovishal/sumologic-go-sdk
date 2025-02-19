# PaginatedOTCollectorsRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[][]OtTag**]([]OtTag.md) | tags associated with the OT collector | [optional] [default to []]
**Os** | Pointer to **NullableString** | Name of the Operating System. | [optional] 
**CollectorVersionRange** | Pointer to [**VersionRange**](VersionRange.md) |  | [optional] 
**Alive** | Pointer to **NullableBool** | alive Status of the OT Collector based on heartbeat. | [optional] 
**IsRemotelyManaged** | Pointer to **NullableBool** | Management Status of the OT Collector based on if it is remotely or locally managed. | [optional] 

## Methods

### NewPaginatedOTCollectorsRequestFilters

`func NewPaginatedOTCollectorsRequestFilters() *PaginatedOTCollectorsRequestFilters`

NewPaginatedOTCollectorsRequestFilters instantiates a new PaginatedOTCollectorsRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOTCollectorsRequestFiltersWithDefaults

`func NewPaginatedOTCollectorsRequestFiltersWithDefaults() *PaginatedOTCollectorsRequestFilters`

NewPaginatedOTCollectorsRequestFiltersWithDefaults instantiates a new PaginatedOTCollectorsRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PaginatedOTCollectorsRequestFilters) GetTags() [][]OtTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PaginatedOTCollectorsRequestFilters) GetTagsOk() (*[][]OtTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PaginatedOTCollectorsRequestFilters) SetTags(v [][]OtTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PaginatedOTCollectorsRequestFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOs

`func (o *PaginatedOTCollectorsRequestFilters) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *PaginatedOTCollectorsRequestFilters) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *PaginatedOTCollectorsRequestFilters) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *PaginatedOTCollectorsRequestFilters) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *PaginatedOTCollectorsRequestFilters) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *PaginatedOTCollectorsRequestFilters) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetCollectorVersionRange

`func (o *PaginatedOTCollectorsRequestFilters) GetCollectorVersionRange() VersionRange`

GetCollectorVersionRange returns the CollectorVersionRange field if non-nil, zero value otherwise.

### GetCollectorVersionRangeOk

`func (o *PaginatedOTCollectorsRequestFilters) GetCollectorVersionRangeOk() (*VersionRange, bool)`

GetCollectorVersionRangeOk returns a tuple with the CollectorVersionRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorVersionRange

`func (o *PaginatedOTCollectorsRequestFilters) SetCollectorVersionRange(v VersionRange)`

SetCollectorVersionRange sets CollectorVersionRange field to given value.

### HasCollectorVersionRange

`func (o *PaginatedOTCollectorsRequestFilters) HasCollectorVersionRange() bool`

HasCollectorVersionRange returns a boolean if a field has been set.

### GetAlive

`func (o *PaginatedOTCollectorsRequestFilters) GetAlive() bool`

GetAlive returns the Alive field if non-nil, zero value otherwise.

### GetAliveOk

`func (o *PaginatedOTCollectorsRequestFilters) GetAliveOk() (*bool, bool)`

GetAliveOk returns a tuple with the Alive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlive

`func (o *PaginatedOTCollectorsRequestFilters) SetAlive(v bool)`

SetAlive sets Alive field to given value.

### HasAlive

`func (o *PaginatedOTCollectorsRequestFilters) HasAlive() bool`

HasAlive returns a boolean if a field has been set.

### SetAliveNil

`func (o *PaginatedOTCollectorsRequestFilters) SetAliveNil(b bool)`

 SetAliveNil sets the value for Alive to be an explicit nil

### UnsetAlive
`func (o *PaginatedOTCollectorsRequestFilters) UnsetAlive()`

UnsetAlive ensures that no value is present for Alive, not even an explicit nil
### GetIsRemotelyManaged

`func (o *PaginatedOTCollectorsRequestFilters) GetIsRemotelyManaged() bool`

GetIsRemotelyManaged returns the IsRemotelyManaged field if non-nil, zero value otherwise.

### GetIsRemotelyManagedOk

`func (o *PaginatedOTCollectorsRequestFilters) GetIsRemotelyManagedOk() (*bool, bool)`

GetIsRemotelyManagedOk returns a tuple with the IsRemotelyManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemotelyManaged

`func (o *PaginatedOTCollectorsRequestFilters) SetIsRemotelyManaged(v bool)`

SetIsRemotelyManaged sets IsRemotelyManaged field to given value.

### HasIsRemotelyManaged

`func (o *PaginatedOTCollectorsRequestFilters) HasIsRemotelyManaged() bool`

HasIsRemotelyManaged returns a boolean if a field has been set.

### SetIsRemotelyManagedNil

`func (o *PaginatedOTCollectorsRequestFilters) SetIsRemotelyManagedNil(b bool)`

 SetIsRemotelyManagedNil sets the value for IsRemotelyManaged to be an explicit nil

### UnsetIsRemotelyManaged
`func (o *PaginatedOTCollectorsRequestFilters) UnsetIsRemotelyManaged()`

UnsetIsRemotelyManaged ensures that no value is present for IsRemotelyManaged, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


