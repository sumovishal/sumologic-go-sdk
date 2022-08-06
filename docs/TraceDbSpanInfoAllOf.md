# TraceDbSpanInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbType** | Pointer to **string** | Database type. | [optional] 
**Instance** | Pointer to **string** | Database instance name, e.g. in java, if jdbc.url&#x3D;\&quot;jdbc:mysql://127.0.0.1:3306/customers\&quot;, the instance name is \&quot;customers\&quot;. | [optional] 
**Statement** | Pointer to **string** | Database statement for the given database type. | [optional] 

## Methods

### NewTraceDbSpanInfoAllOf

`func NewTraceDbSpanInfoAllOf() *TraceDbSpanInfoAllOf`

NewTraceDbSpanInfoAllOf instantiates a new TraceDbSpanInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceDbSpanInfoAllOfWithDefaults

`func NewTraceDbSpanInfoAllOfWithDefaults() *TraceDbSpanInfoAllOf`

NewTraceDbSpanInfoAllOfWithDefaults instantiates a new TraceDbSpanInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbType

`func (o *TraceDbSpanInfoAllOf) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *TraceDbSpanInfoAllOf) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *TraceDbSpanInfoAllOf) SetDbType(v string)`

SetDbType sets DbType field to given value.

### HasDbType

`func (o *TraceDbSpanInfoAllOf) HasDbType() bool`

HasDbType returns a boolean if a field has been set.

### GetInstance

`func (o *TraceDbSpanInfoAllOf) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *TraceDbSpanInfoAllOf) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *TraceDbSpanInfoAllOf) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *TraceDbSpanInfoAllOf) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetStatement

`func (o *TraceDbSpanInfoAllOf) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *TraceDbSpanInfoAllOf) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *TraceDbSpanInfoAllOf) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *TraceDbSpanInfoAllOf) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


