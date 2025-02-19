# SearchQueryPaginatedRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]Field**](Field.md) | List of all the fields defined for each of the records returned. | 
**Records** | Pointer to **[]map[string]string** | Map of the field names to the field values. | [optional] 

## Methods

### NewSearchQueryPaginatedRecords

`func NewSearchQueryPaginatedRecords(fields []Field, ) *SearchQueryPaginatedRecords`

NewSearchQueryPaginatedRecords instantiates a new SearchQueryPaginatedRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryPaginatedRecordsWithDefaults

`func NewSearchQueryPaginatedRecordsWithDefaults() *SearchQueryPaginatedRecords`

NewSearchQueryPaginatedRecordsWithDefaults instantiates a new SearchQueryPaginatedRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *SearchQueryPaginatedRecords) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchQueryPaginatedRecords) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchQueryPaginatedRecords) SetFields(v []Field)`

SetFields sets Fields field to given value.


### GetRecords

`func (o *SearchQueryPaginatedRecords) GetRecords() []map[string]string`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SearchQueryPaginatedRecords) GetRecordsOk() (*[]map[string]string, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SearchQueryPaginatedRecords) SetRecords(v []map[string]string)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SearchQueryPaginatedRecords) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


