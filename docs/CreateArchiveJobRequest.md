# CreateArchiveJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the ingestion job. | 
**StartTime** | **time.Time** | The starting timestamp of the ingestion job. | 
**EndTime** | **time.Time** | The ending timestamp of the ingestion job. | 

## Methods

### NewCreateArchiveJobRequest

`func NewCreateArchiveJobRequest(name string, startTime time.Time, endTime time.Time, ) *CreateArchiveJobRequest`

NewCreateArchiveJobRequest instantiates a new CreateArchiveJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateArchiveJobRequestWithDefaults

`func NewCreateArchiveJobRequestWithDefaults() *CreateArchiveJobRequest`

NewCreateArchiveJobRequestWithDefaults instantiates a new CreateArchiveJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateArchiveJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateArchiveJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateArchiveJobRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *CreateArchiveJobRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateArchiveJobRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateArchiveJobRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CreateArchiveJobRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateArchiveJobRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateArchiveJobRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


