# EventExtractionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of event extraction rule. | 
**Description** | Pointer to **string** | Description of event extraction rule. | [optional] 
**Query** | **string** | Query String. | 
**CorrelationExpression** | [**CorrelationExpression**](CorrelationExpression.md) |  | 
**Retention** | Pointer to **int64** | Retention period in days for which the events are available in Sumologic. | [optional] [default to 180]

## Methods

### NewEventExtractionRule

`func NewEventExtractionRule(name string, query string, correlationExpression CorrelationExpression, ) *EventExtractionRule`

NewEventExtractionRule instantiates a new EventExtractionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventExtractionRuleWithDefaults

`func NewEventExtractionRuleWithDefaults() *EventExtractionRule`

NewEventExtractionRuleWithDefaults instantiates a new EventExtractionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventExtractionRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventExtractionRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventExtractionRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EventExtractionRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventExtractionRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventExtractionRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventExtractionRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuery

`func (o *EventExtractionRule) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EventExtractionRule) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EventExtractionRule) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetCorrelationExpression

`func (o *EventExtractionRule) GetCorrelationExpression() CorrelationExpression`

GetCorrelationExpression returns the CorrelationExpression field if non-nil, zero value otherwise.

### GetCorrelationExpressionOk

`func (o *EventExtractionRule) GetCorrelationExpressionOk() (*CorrelationExpression, bool)`

GetCorrelationExpressionOk returns a tuple with the CorrelationExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationExpression

`func (o *EventExtractionRule) SetCorrelationExpression(v CorrelationExpression)`

SetCorrelationExpression sets CorrelationExpression field to given value.


### GetRetention

`func (o *EventExtractionRule) GetRetention() int64`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *EventExtractionRule) GetRetentionOk() (*int64, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *EventExtractionRule) SetRetention(v int64)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *EventExtractionRule) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


