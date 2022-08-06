# CriticalPathServiceBreakdownElementDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumSpans** | **int32** | Number of spans that are part of this service. | 
**LongestSegmentDuration** | **int64** | Number of nanoseconds the longest span segment in the critical path lasted. | 

## Methods

### NewCriticalPathServiceBreakdownElementDetailAllOf

`func NewCriticalPathServiceBreakdownElementDetailAllOf(numSpans int32, longestSegmentDuration int64, ) *CriticalPathServiceBreakdownElementDetailAllOf`

NewCriticalPathServiceBreakdownElementDetailAllOf instantiates a new CriticalPathServiceBreakdownElementDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriticalPathServiceBreakdownElementDetailAllOfWithDefaults

`func NewCriticalPathServiceBreakdownElementDetailAllOfWithDefaults() *CriticalPathServiceBreakdownElementDetailAllOf`

NewCriticalPathServiceBreakdownElementDetailAllOfWithDefaults instantiates a new CriticalPathServiceBreakdownElementDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumSpans

`func (o *CriticalPathServiceBreakdownElementDetailAllOf) GetNumSpans() int32`

GetNumSpans returns the NumSpans field if non-nil, zero value otherwise.

### GetNumSpansOk

`func (o *CriticalPathServiceBreakdownElementDetailAllOf) GetNumSpansOk() (*int32, bool)`

GetNumSpansOk returns a tuple with the NumSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSpans

`func (o *CriticalPathServiceBreakdownElementDetailAllOf) SetNumSpans(v int32)`

SetNumSpans sets NumSpans field to given value.


### GetLongestSegmentDuration

`func (o *CriticalPathServiceBreakdownElementDetailAllOf) GetLongestSegmentDuration() int64`

GetLongestSegmentDuration returns the LongestSegmentDuration field if non-nil, zero value otherwise.

### GetLongestSegmentDurationOk

`func (o *CriticalPathServiceBreakdownElementDetailAllOf) GetLongestSegmentDurationOk() (*int64, bool)`

GetLongestSegmentDurationOk returns a tuple with the LongestSegmentDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestSegmentDuration

`func (o *CriticalPathServiceBreakdownElementDetailAllOf) SetLongestSegmentDuration(v int64)`

SetLongestSegmentDuration sets LongestSegmentDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


