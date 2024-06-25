# StringMatchCorrelationExpressionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryFieldName** | **string** | Name of the query field returned by a log search query. | 
**EventFieldName** | **string** | Name of the field from event query output. | 
**StringMatchingAlgorithm** | **string** | Type of string matching algorithm which tells how to match eventFieldName and queryFieldName. | 

## Methods

### NewStringMatchCorrelationExpressionAllOf

`func NewStringMatchCorrelationExpressionAllOf(queryFieldName string, eventFieldName string, stringMatchingAlgorithm string, ) *StringMatchCorrelationExpressionAllOf`

NewStringMatchCorrelationExpressionAllOf instantiates a new StringMatchCorrelationExpressionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringMatchCorrelationExpressionAllOfWithDefaults

`func NewStringMatchCorrelationExpressionAllOfWithDefaults() *StringMatchCorrelationExpressionAllOf`

NewStringMatchCorrelationExpressionAllOfWithDefaults instantiates a new StringMatchCorrelationExpressionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryFieldName

`func (o *StringMatchCorrelationExpressionAllOf) GetQueryFieldName() string`

GetQueryFieldName returns the QueryFieldName field if non-nil, zero value otherwise.

### GetQueryFieldNameOk

`func (o *StringMatchCorrelationExpressionAllOf) GetQueryFieldNameOk() (*string, bool)`

GetQueryFieldNameOk returns a tuple with the QueryFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFieldName

`func (o *StringMatchCorrelationExpressionAllOf) SetQueryFieldName(v string)`

SetQueryFieldName sets QueryFieldName field to given value.


### GetEventFieldName

`func (o *StringMatchCorrelationExpressionAllOf) GetEventFieldName() string`

GetEventFieldName returns the EventFieldName field if non-nil, zero value otherwise.

### GetEventFieldNameOk

`func (o *StringMatchCorrelationExpressionAllOf) GetEventFieldNameOk() (*string, bool)`

GetEventFieldNameOk returns a tuple with the EventFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFieldName

`func (o *StringMatchCorrelationExpressionAllOf) SetEventFieldName(v string)`

SetEventFieldName sets EventFieldName field to given value.


### GetStringMatchingAlgorithm

`func (o *StringMatchCorrelationExpressionAllOf) GetStringMatchingAlgorithm() string`

GetStringMatchingAlgorithm returns the StringMatchingAlgorithm field if non-nil, zero value otherwise.

### GetStringMatchingAlgorithmOk

`func (o *StringMatchCorrelationExpressionAllOf) GetStringMatchingAlgorithmOk() (*string, bool)`

GetStringMatchingAlgorithmOk returns a tuple with the StringMatchingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringMatchingAlgorithm

`func (o *StringMatchCorrelationExpressionAllOf) SetStringMatchingAlgorithm(v string)`

SetStringMatchingAlgorithm sets StringMatchingAlgorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


