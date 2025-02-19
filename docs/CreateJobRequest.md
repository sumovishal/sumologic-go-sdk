# CreateJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The actual search expression. | 
**From** | **string** | The ISO 8601 date and time of the time range to start the search.  | 
**To** | **string** | The ISO 8601 date and time of the time range to end the search.  | 
**TimeZone** | Pointer to **string** | The time zone of the from and to time. | [optional] [default to "UTC"]
**ByReceiptTime** | Pointer to **string** | Flag to order the search results in the order collector received it. | [optional] [default to "false"]

## Methods

### NewCreateJobRequest

`func NewCreateJobRequest(query string, from string, to string, ) *CreateJobRequest`

NewCreateJobRequest instantiates a new CreateJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobRequestWithDefaults

`func NewCreateJobRequestWithDefaults() *CreateJobRequest`

NewCreateJobRequestWithDefaults instantiates a new CreateJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CreateJobRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateJobRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateJobRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFrom

`func (o *CreateJobRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateJobRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateJobRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *CreateJobRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateJobRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateJobRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetTimeZone

`func (o *CreateJobRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CreateJobRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CreateJobRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CreateJobRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetByReceiptTime

`func (o *CreateJobRequest) GetByReceiptTime() string`

GetByReceiptTime returns the ByReceiptTime field if non-nil, zero value otherwise.

### GetByReceiptTimeOk

`func (o *CreateJobRequest) GetByReceiptTimeOk() (*string, bool)`

GetByReceiptTimeOk returns a tuple with the ByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByReceiptTime

`func (o *CreateJobRequest) SetByReceiptTime(v string)`

SetByReceiptTime sets ByReceiptTime field to given value.

### HasByReceiptTime

`func (o *CreateJobRequest) HasByReceiptTime() bool`

HasByReceiptTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


