# StringMatchCorrelationExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryFieldName** | **string** | Name of the query field returned by a log search query. | 
**EventFieldName** | **string** | Name of the field from event query output. | 
**StringMatchingAlgorithm** | **string** | Type of string matching algorithm which tells how to match eventFieldName and queryFieldName. | 

## Methods

### NewStringMatchCorrelationExpression

`func NewStringMatchCorrelationExpression(queryFieldName string, eventFieldName string, stringMatchingAlgorithm string, ) *StringMatchCorrelationExpression`

NewStringMatchCorrelationExpression instantiates a new StringMatchCorrelationExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringMatchCorrelationExpressionWithDefaults

`func NewStringMatchCorrelationExpressionWithDefaults() *StringMatchCorrelationExpression`

NewStringMatchCorrelationExpressionWithDefaults instantiates a new StringMatchCorrelationExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryFieldName

`func (o *StringMatchCorrelationExpression) GetQueryFieldName() string`

GetQueryFieldName returns the QueryFieldName field if non-nil, zero value otherwise.

### GetQueryFieldNameOk

`func (o *StringMatchCorrelationExpression) GetQueryFieldNameOk() (*string, bool)`

GetQueryFieldNameOk returns a tuple with the QueryFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFieldName

`func (o *StringMatchCorrelationExpression) SetQueryFieldName(v string)`

SetQueryFieldName sets QueryFieldName field to given value.


### GetEventFieldName

`func (o *StringMatchCorrelationExpression) GetEventFieldName() string`

GetEventFieldName returns the EventFieldName field if non-nil, zero value otherwise.

### GetEventFieldNameOk

`func (o *StringMatchCorrelationExpression) GetEventFieldNameOk() (*string, bool)`

GetEventFieldNameOk returns a tuple with the EventFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFieldName

`func (o *StringMatchCorrelationExpression) SetEventFieldName(v string)`

SetEventFieldName sets EventFieldName field to given value.


### GetStringMatchingAlgorithm

`func (o *StringMatchCorrelationExpression) GetStringMatchingAlgorithm() string`

GetStringMatchingAlgorithm returns the StringMatchingAlgorithm field if non-nil, zero value otherwise.

### GetStringMatchingAlgorithmOk

`func (o *StringMatchCorrelationExpression) GetStringMatchingAlgorithmOk() (*string, bool)`

GetStringMatchingAlgorithmOk returns a tuple with the StringMatchingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringMatchingAlgorithm

`func (o *StringMatchCorrelationExpression) SetStringMatchingAlgorithm(v string)`

SetStringMatchingAlgorithm sets StringMatchingAlgorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


