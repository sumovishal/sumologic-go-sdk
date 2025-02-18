# EventExtractionRuleWithDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of event extraction rule. | 
**Description** | Pointer to **string** | Description of event extraction rule. | [optional] 
**Query** | **string** | Query String. | 
**CorrelationExpression** | [**CorrelationExpression**](CorrelationExpression.md) |  | 
**Id** | **string** | Id of the event extraction rule. | 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] 
**CreatedBy** | Pointer to **string** | Identifier of the user who created the resource. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Last modification timestamp in UTC. | [optional] 
**ModifiedBy** | Pointer to **string** | Identifier of the user who last modified the resource. | [optional] 

## Methods

### NewEventExtractionRuleWithDetails

`func NewEventExtractionRuleWithDetails(name string, query string, correlationExpression CorrelationExpression, id string, ) *EventExtractionRuleWithDetails`

NewEventExtractionRuleWithDetails instantiates a new EventExtractionRuleWithDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventExtractionRuleWithDetailsWithDefaults

`func NewEventExtractionRuleWithDetailsWithDefaults() *EventExtractionRuleWithDetails`

NewEventExtractionRuleWithDetailsWithDefaults instantiates a new EventExtractionRuleWithDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventExtractionRuleWithDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventExtractionRuleWithDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventExtractionRuleWithDetails) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EventExtractionRuleWithDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventExtractionRuleWithDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventExtractionRuleWithDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventExtractionRuleWithDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuery

`func (o *EventExtractionRuleWithDetails) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EventExtractionRuleWithDetails) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EventExtractionRuleWithDetails) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetCorrelationExpression

`func (o *EventExtractionRuleWithDetails) GetCorrelationExpression() CorrelationExpression`

GetCorrelationExpression returns the CorrelationExpression field if non-nil, zero value otherwise.

### GetCorrelationExpressionOk

`func (o *EventExtractionRuleWithDetails) GetCorrelationExpressionOk() (*CorrelationExpression, bool)`

GetCorrelationExpressionOk returns a tuple with the CorrelationExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationExpression

`func (o *EventExtractionRuleWithDetails) SetCorrelationExpression(v CorrelationExpression)`

SetCorrelationExpression sets CorrelationExpression field to given value.


### GetId

`func (o *EventExtractionRuleWithDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventExtractionRuleWithDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventExtractionRuleWithDetails) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EventExtractionRuleWithDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventExtractionRuleWithDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventExtractionRuleWithDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventExtractionRuleWithDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EventExtractionRuleWithDetails) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EventExtractionRuleWithDetails) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EventExtractionRuleWithDetails) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EventExtractionRuleWithDetails) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedAt

`func (o *EventExtractionRuleWithDetails) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EventExtractionRuleWithDetails) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EventExtractionRuleWithDetails) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EventExtractionRuleWithDetails) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *EventExtractionRuleWithDetails) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *EventExtractionRuleWithDetails) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *EventExtractionRuleWithDetails) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *EventExtractionRuleWithDetails) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


