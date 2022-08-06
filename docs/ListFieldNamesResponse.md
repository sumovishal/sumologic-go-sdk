# ListFieldNamesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FieldName**](FieldName.md) | List of all built-in and custom field names. | 

## Methods

### NewListFieldNamesResponse

`func NewListFieldNamesResponse(data []FieldName, ) *ListFieldNamesResponse`

NewListFieldNamesResponse instantiates a new ListFieldNamesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFieldNamesResponseWithDefaults

`func NewListFieldNamesResponseWithDefaults() *ListFieldNamesResponse`

NewListFieldNamesResponseWithDefaults instantiates a new ListFieldNamesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFieldNamesResponse) GetData() []FieldName`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFieldNamesResponse) GetDataOk() (*[]FieldName, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFieldNamesResponse) SetData(v []FieldName)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


