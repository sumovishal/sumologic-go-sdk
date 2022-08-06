# TraceSpanDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | Produced error message (could be a stack trace, database error code, ..) | [optional] 
**Fields** | Pointer to [**map[string]TracingValue**](TracingValue.md) | Fields attached to this span. | [optional] 
**CriticalPathContribution** | Pointer to [**TraceSpanCriticalPathContribution**](TraceSpanCriticalPathContribution.md) |  | [optional] 
**Logs** | Pointer to **[]string** | Logs attached to this span. | [optional] 
**Events** | Pointer to [**[]SpanEvent**](SpanEvent.md) | Events attached to this span. | [optional] 
**Links** | Pointer to [**[]SpanLink**](SpanLink.md) | List of casually related spans. | [optional] 

## Methods

### NewTraceSpanDetailAllOf

`func NewTraceSpanDetailAllOf() *TraceSpanDetailAllOf`

NewTraceSpanDetailAllOf instantiates a new TraceSpanDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSpanDetailAllOfWithDefaults

`func NewTraceSpanDetailAllOfWithDefaults() *TraceSpanDetailAllOf`

NewTraceSpanDetailAllOfWithDefaults instantiates a new TraceSpanDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *TraceSpanDetailAllOf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TraceSpanDetailAllOf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TraceSpanDetailAllOf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TraceSpanDetailAllOf) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFields

`func (o *TraceSpanDetailAllOf) GetFields() map[string]TracingValue`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TraceSpanDetailAllOf) GetFieldsOk() (*map[string]TracingValue, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TraceSpanDetailAllOf) SetFields(v map[string]TracingValue)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TraceSpanDetailAllOf) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetCriticalPathContribution

`func (o *TraceSpanDetailAllOf) GetCriticalPathContribution() TraceSpanCriticalPathContribution`

GetCriticalPathContribution returns the CriticalPathContribution field if non-nil, zero value otherwise.

### GetCriticalPathContributionOk

`func (o *TraceSpanDetailAllOf) GetCriticalPathContributionOk() (*TraceSpanCriticalPathContribution, bool)`

GetCriticalPathContributionOk returns a tuple with the CriticalPathContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalPathContribution

`func (o *TraceSpanDetailAllOf) SetCriticalPathContribution(v TraceSpanCriticalPathContribution)`

SetCriticalPathContribution sets CriticalPathContribution field to given value.

### HasCriticalPathContribution

`func (o *TraceSpanDetailAllOf) HasCriticalPathContribution() bool`

HasCriticalPathContribution returns a boolean if a field has been set.

### GetLogs

`func (o *TraceSpanDetailAllOf) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *TraceSpanDetailAllOf) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *TraceSpanDetailAllOf) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *TraceSpanDetailAllOf) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetEvents

`func (o *TraceSpanDetailAllOf) GetEvents() []SpanEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TraceSpanDetailAllOf) GetEventsOk() (*[]SpanEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TraceSpanDetailAllOf) SetEvents(v []SpanEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TraceSpanDetailAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLinks

`func (o *TraceSpanDetailAllOf) GetLinks() []SpanLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TraceSpanDetailAllOf) GetLinksOk() (*[]SpanLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TraceSpanDetailAllOf) SetLinks(v []SpanLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TraceSpanDetailAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


