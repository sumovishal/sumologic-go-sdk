# RequestBasedSliAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | Pointer to **float32** | Compared against threshold query&#39;s raw data points to determine success. | [optional] 
**Op** | Pointer to **string** | Comparison function with threshold (LessThan/GreaterThan/LessThanOrEqual/GreaterThanOrEqual). | [optional] 

## Methods

### NewRequestBasedSliAllOf

`func NewRequestBasedSliAllOf() *RequestBasedSliAllOf`

NewRequestBasedSliAllOf instantiates a new RequestBasedSliAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestBasedSliAllOfWithDefaults

`func NewRequestBasedSliAllOfWithDefaults() *RequestBasedSliAllOf`

NewRequestBasedSliAllOfWithDefaults instantiates a new RequestBasedSliAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *RequestBasedSliAllOf) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RequestBasedSliAllOf) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RequestBasedSliAllOf) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *RequestBasedSliAllOf) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetOp

`func (o *RequestBasedSliAllOf) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RequestBasedSliAllOf) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RequestBasedSliAllOf) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RequestBasedSliAllOf) HasOp() bool`

HasOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


