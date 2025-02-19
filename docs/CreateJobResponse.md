# CreateJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The search job identifier. | [optional] 
**Link** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewCreateJobResponse

`func NewCreateJobResponse() *CreateJobResponse`

NewCreateJobResponse instantiates a new CreateJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobResponseWithDefaults

`func NewCreateJobResponseWithDefaults() *CreateJobResponse`

NewCreateJobResponseWithDefaults instantiates a new CreateJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateJobResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateJobResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateJobResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateJobResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLink

`func (o *CreateJobResponse) GetLink() Link`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CreateJobResponse) GetLinkOk() (*Link, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CreateJobResponse) SetLink(v Link)`

SetLink sets Link field to given value.

### HasLink

`func (o *CreateJobResponse) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


