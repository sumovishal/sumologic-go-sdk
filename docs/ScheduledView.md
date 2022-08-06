# ScheduledView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The query that defines the data to be included in the scheduled view. | 
**IndexName** | **string** | Name of the index for the scheduled view. | 
**StartTime** | **time.Time** | Start timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**RetentionPeriod** | Pointer to **int32** | The number of days to retain data in the scheduled view, or -1 to use the default value for your account. Only relevant if your account has multi-retention enabled. | [optional] [default to -1]
**DataForwardingId** | Pointer to **string** | An optional ID of a data forwarding configuration to be used by the scheduled view. | [optional] 
**ParsingMode** | Pointer to **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to "Manual"]
**NewRetentionPeriod** | Pointer to **int32** | If the retention period is scheduled to be updated in the future (i.e., if retention period is previously reduced with value of reduceRetentionPeriodImmediately as false), this property gives the future value of retention period while retentionPeriod gives the current value. retentionPeriod will take up the value of newRetentionPeriod after the scheduled time. | [optional] 
**RetentionEffectiveAt** | Pointer to **time.Time** | When the newRetentionPeriod will become effective in UTC format. | [optional] 
**Id** | **string** | Identifier for the scheduled view. | 
**IndexId** | Pointer to **string** | The &#x60;id&#x60; of the Index where the output from Scheduled view is stored. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp in UTC. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Last modification timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] 
**CreatedByOptimizeIt** | Pointer to **bool** | If the scheduled view is created by OptimizeIt. | [optional] 
**Error** | Pointer to **string** | Errors related to the scheduled view. | [optional] 
**Status** | Pointer to **string** | Status of the scheduled view. | [optional] 
**TotalBytes** | Pointer to **int64** | Total storage consumed by the scheduled view. | [optional] 
**TotalMessageCount** | Pointer to **int64** | Total number of messages for the scheduled view. | [optional] 
**CreatedBy** | Pointer to **string** | Identifier of the user who created the scheduled view. | [optional] 
**ModifiedBy** | Pointer to **string** | Identifier of the user who last modified the resource. | [optional] 
**FilledRanges** | Pointer to [**[]FilledRange**](FilledRange.md) | List of the different units of filled ranges since the autoview has been created. | [optional] 

## Methods

### NewScheduledView

`func NewScheduledView(query string, indexName string, startTime time.Time, id string, ) *ScheduledView`

NewScheduledView instantiates a new ScheduledView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledViewWithDefaults

`func NewScheduledViewWithDefaults() *ScheduledView`

NewScheduledViewWithDefaults instantiates a new ScheduledView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ScheduledView) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ScheduledView) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ScheduledView) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetIndexName

`func (o *ScheduledView) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *ScheduledView) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *ScheduledView) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.


### GetStartTime

`func (o *ScheduledView) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduledView) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduledView) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetRetentionPeriod

`func (o *ScheduledView) GetRetentionPeriod() int32`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *ScheduledView) GetRetentionPeriodOk() (*int32, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *ScheduledView) SetRetentionPeriod(v int32)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *ScheduledView) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetDataForwardingId

`func (o *ScheduledView) GetDataForwardingId() string`

GetDataForwardingId returns the DataForwardingId field if non-nil, zero value otherwise.

### GetDataForwardingIdOk

`func (o *ScheduledView) GetDataForwardingIdOk() (*string, bool)`

GetDataForwardingIdOk returns a tuple with the DataForwardingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwardingId

`func (o *ScheduledView) SetDataForwardingId(v string)`

SetDataForwardingId sets DataForwardingId field to given value.

### HasDataForwardingId

`func (o *ScheduledView) HasDataForwardingId() bool`

HasDataForwardingId returns a boolean if a field has been set.

### GetParsingMode

`func (o *ScheduledView) GetParsingMode() string`

GetParsingMode returns the ParsingMode field if non-nil, zero value otherwise.

### GetParsingModeOk

`func (o *ScheduledView) GetParsingModeOk() (*string, bool)`

GetParsingModeOk returns a tuple with the ParsingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingMode

`func (o *ScheduledView) SetParsingMode(v string)`

SetParsingMode sets ParsingMode field to given value.

### HasParsingMode

`func (o *ScheduledView) HasParsingMode() bool`

HasParsingMode returns a boolean if a field has been set.

### GetNewRetentionPeriod

`func (o *ScheduledView) GetNewRetentionPeriod() int32`

GetNewRetentionPeriod returns the NewRetentionPeriod field if non-nil, zero value otherwise.

### GetNewRetentionPeriodOk

`func (o *ScheduledView) GetNewRetentionPeriodOk() (*int32, bool)`

GetNewRetentionPeriodOk returns a tuple with the NewRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRetentionPeriod

`func (o *ScheduledView) SetNewRetentionPeriod(v int32)`

SetNewRetentionPeriod sets NewRetentionPeriod field to given value.

### HasNewRetentionPeriod

`func (o *ScheduledView) HasNewRetentionPeriod() bool`

HasNewRetentionPeriod returns a boolean if a field has been set.

### GetRetentionEffectiveAt

`func (o *ScheduledView) GetRetentionEffectiveAt() time.Time`

GetRetentionEffectiveAt returns the RetentionEffectiveAt field if non-nil, zero value otherwise.

### GetRetentionEffectiveAtOk

`func (o *ScheduledView) GetRetentionEffectiveAtOk() (*time.Time, bool)`

GetRetentionEffectiveAtOk returns a tuple with the RetentionEffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionEffectiveAt

`func (o *ScheduledView) SetRetentionEffectiveAt(v time.Time)`

SetRetentionEffectiveAt sets RetentionEffectiveAt field to given value.

### HasRetentionEffectiveAt

`func (o *ScheduledView) HasRetentionEffectiveAt() bool`

HasRetentionEffectiveAt returns a boolean if a field has been set.

### GetId

`func (o *ScheduledView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledView) SetId(v string)`

SetId sets Id field to given value.


### GetIndexId

`func (o *ScheduledView) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *ScheduledView) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *ScheduledView) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.

### HasIndexId

`func (o *ScheduledView) HasIndexId() bool`

HasIndexId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ScheduledView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduledView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduledView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScheduledView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ScheduledView) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ScheduledView) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ScheduledView) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ScheduledView) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCreatedByOptimizeIt

`func (o *ScheduledView) GetCreatedByOptimizeIt() bool`

GetCreatedByOptimizeIt returns the CreatedByOptimizeIt field if non-nil, zero value otherwise.

### GetCreatedByOptimizeItOk

`func (o *ScheduledView) GetCreatedByOptimizeItOk() (*bool, bool)`

GetCreatedByOptimizeItOk returns a tuple with the CreatedByOptimizeIt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByOptimizeIt

`func (o *ScheduledView) SetCreatedByOptimizeIt(v bool)`

SetCreatedByOptimizeIt sets CreatedByOptimizeIt field to given value.

### HasCreatedByOptimizeIt

`func (o *ScheduledView) HasCreatedByOptimizeIt() bool`

HasCreatedByOptimizeIt returns a boolean if a field has been set.

### GetError

`func (o *ScheduledView) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScheduledView) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScheduledView) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ScheduledView) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduledView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduledView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalBytes

`func (o *ScheduledView) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *ScheduledView) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *ScheduledView) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *ScheduledView) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetTotalMessageCount

`func (o *ScheduledView) GetTotalMessageCount() int64`

GetTotalMessageCount returns the TotalMessageCount field if non-nil, zero value otherwise.

### GetTotalMessageCountOk

`func (o *ScheduledView) GetTotalMessageCountOk() (*int64, bool)`

GetTotalMessageCountOk returns a tuple with the TotalMessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessageCount

`func (o *ScheduledView) SetTotalMessageCount(v int64)`

SetTotalMessageCount sets TotalMessageCount field to given value.

### HasTotalMessageCount

`func (o *ScheduledView) HasTotalMessageCount() bool`

HasTotalMessageCount returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ScheduledView) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScheduledView) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScheduledView) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ScheduledView) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *ScheduledView) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ScheduledView) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ScheduledView) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ScheduledView) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetFilledRanges

`func (o *ScheduledView) GetFilledRanges() []FilledRange`

GetFilledRanges returns the FilledRanges field if non-nil, zero value otherwise.

### GetFilledRangesOk

`func (o *ScheduledView) GetFilledRangesOk() (*[]FilledRange, bool)`

GetFilledRangesOk returns a tuple with the FilledRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledRanges

`func (o *ScheduledView) SetFilledRanges(v []FilledRange)`

SetFilledRanges sets FilledRanges field to given value.

### HasFilledRanges

`func (o *ScheduledView) HasFilledRanges() bool`

HasFilledRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


