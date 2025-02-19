# HistogramBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | **int64** | Length is in milliseconds, tells the width of the bucket. | 
**Count** | **int32** | Count of messages in this bucket. | 
**StartTimestamp** | **int64** | Start time of the bucket. | 

## Methods

### NewHistogramBucket

`func NewHistogramBucket(length int64, count int32, startTimestamp int64, ) *HistogramBucket`

NewHistogramBucket instantiates a new HistogramBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistogramBucketWithDefaults

`func NewHistogramBucketWithDefaults() *HistogramBucket`

NewHistogramBucketWithDefaults instantiates a new HistogramBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *HistogramBucket) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *HistogramBucket) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *HistogramBucket) SetLength(v int64)`

SetLength sets Length field to given value.


### GetCount

`func (o *HistogramBucket) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HistogramBucket) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HistogramBucket) SetCount(v int32)`

SetCount sets Count field to given value.


### GetStartTimestamp

`func (o *HistogramBucket) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *HistogramBucket) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *HistogramBucket) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


