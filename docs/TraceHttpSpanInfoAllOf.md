# TraceHttpSpanInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | HTTP method of the request for the associated span. | [optional] 
**Url** | Pointer to **string** | URL of the request being handled in this span, in the standard URI format. | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code for the associated span. | [optional] 

## Methods

### NewTraceHttpSpanInfoAllOf

`func NewTraceHttpSpanInfoAllOf() *TraceHttpSpanInfoAllOf`

NewTraceHttpSpanInfoAllOf instantiates a new TraceHttpSpanInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceHttpSpanInfoAllOfWithDefaults

`func NewTraceHttpSpanInfoAllOfWithDefaults() *TraceHttpSpanInfoAllOf`

NewTraceHttpSpanInfoAllOfWithDefaults instantiates a new TraceHttpSpanInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TraceHttpSpanInfoAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TraceHttpSpanInfoAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TraceHttpSpanInfoAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TraceHttpSpanInfoAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *TraceHttpSpanInfoAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TraceHttpSpanInfoAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TraceHttpSpanInfoAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TraceHttpSpanInfoAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatusCode

`func (o *TraceHttpSpanInfoAllOf) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TraceHttpSpanInfoAllOf) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TraceHttpSpanInfoAllOf) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *TraceHttpSpanInfoAllOf) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


