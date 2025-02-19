# ListAppsV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | [**[]AppV2**](AppV2.md) | An array of apps. | 

## Methods

### NewListAppsV2Response

`func NewListAppsV2Response(apps []AppV2, ) *ListAppsV2Response`

NewListAppsV2Response instantiates a new ListAppsV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAppsV2ResponseWithDefaults

`func NewListAppsV2ResponseWithDefaults() *ListAppsV2Response`

NewListAppsV2ResponseWithDefaults instantiates a new ListAppsV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *ListAppsV2Response) GetApps() []AppV2`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ListAppsV2Response) GetAppsOk() (*[]AppV2, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ListAppsV2Response) SetApps(v []AppV2)`

SetApps sets Apps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


