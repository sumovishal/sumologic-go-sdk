# SearchQueryPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]Field**](Field.md) | List of all the fields defined for each of the messages returned. | 
**Messages** | **[]map[string]string** | Map of the field names to the field values. | 

## Methods

### NewSearchQueryPaginatedResponse

`func NewSearchQueryPaginatedResponse(fields []Field, messages []map[string]string, ) *SearchQueryPaginatedResponse`

NewSearchQueryPaginatedResponse instantiates a new SearchQueryPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryPaginatedResponseWithDefaults

`func NewSearchQueryPaginatedResponseWithDefaults() *SearchQueryPaginatedResponse`

NewSearchQueryPaginatedResponseWithDefaults instantiates a new SearchQueryPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *SearchQueryPaginatedResponse) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchQueryPaginatedResponse) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchQueryPaginatedResponse) SetFields(v []Field)`

SetFields sets Fields field to given value.


### GetMessages

`func (o *SearchQueryPaginatedResponse) GetMessages() []map[string]string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *SearchQueryPaginatedResponse) GetMessagesOk() (*[]map[string]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *SearchQueryPaginatedResponse) SetMessages(v []map[string]string)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


