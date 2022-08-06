# AndTracingExpressionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expressions** | [**[]TraceQueryExpression**](TraceQueryExpression.md) | Evaluates to true, if (and only if) all expressions evaluate to true, otherwise evaluates to false. | 

## Methods

### NewAndTracingExpressionAllOf

`func NewAndTracingExpressionAllOf(expressions []TraceQueryExpression, ) *AndTracingExpressionAllOf`

NewAndTracingExpressionAllOf instantiates a new AndTracingExpressionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndTracingExpressionAllOfWithDefaults

`func NewAndTracingExpressionAllOfWithDefaults() *AndTracingExpressionAllOf`

NewAndTracingExpressionAllOfWithDefaults instantiates a new AndTracingExpressionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpressions

`func (o *AndTracingExpressionAllOf) GetExpressions() []TraceQueryExpression`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *AndTracingExpressionAllOf) GetExpressionsOk() (*[]TraceQueryExpression, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *AndTracingExpressionAllOf) SetExpressions(v []TraceQueryExpression)`

SetExpressions sets Expressions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


