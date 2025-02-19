# CollectorTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of the given tag. | 
**Value** | Pointer to **string** | Values of the given tag. | [optional] 

## Methods

### NewCollectorTag

`func NewCollectorTag(key string, ) *CollectorTag`

NewCollectorTag instantiates a new CollectorTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectorTagWithDefaults

`func NewCollectorTagWithDefaults() *CollectorTag`

NewCollectorTagWithDefaults instantiates a new CollectorTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CollectorTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CollectorTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CollectorTag) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *CollectorTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CollectorTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CollectorTag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CollectorTag) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


