# SchemaBaseTypeToVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the schema. | 
**Versions** | [**[]SchemaBaseComplete**](SchemaBaseComplete.md) | List of schema base identities sorted by latest version for a specific schema type. | 

## Methods

### NewSchemaBaseTypeToVersionsResponse

`func NewSchemaBaseTypeToVersionsResponse(type_ string, versions []SchemaBaseComplete, ) *SchemaBaseTypeToVersionsResponse`

NewSchemaBaseTypeToVersionsResponse instantiates a new SchemaBaseTypeToVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaBaseTypeToVersionsResponseWithDefaults

`func NewSchemaBaseTypeToVersionsResponseWithDefaults() *SchemaBaseTypeToVersionsResponse`

NewSchemaBaseTypeToVersionsResponseWithDefaults instantiates a new SchemaBaseTypeToVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemaBaseTypeToVersionsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaBaseTypeToVersionsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaBaseTypeToVersionsResponse) SetType(v string)`

SetType sets Type field to given value.


### GetVersions

`func (o *SchemaBaseTypeToVersionsResponse) GetVersions() []SchemaBaseComplete`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *SchemaBaseTypeToVersionsResponse) GetVersionsOk() (*[]SchemaBaseComplete, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *SchemaBaseTypeToVersionsResponse) SetVersions(v []SchemaBaseComplete)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


