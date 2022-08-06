# ServiceMapPanelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | Filter services by the application custom tag. | [optional] 
**Service** | Pointer to **string** | Show only the specific service and its connections to other services. | [optional] 
**ShowRemoteServices** | Pointer to **bool** | Show remote services, like databases or external calls, automatically detected in client traffic. | [optional] 

## Methods

### NewServiceMapPanelAllOf

`func NewServiceMapPanelAllOf() *ServiceMapPanelAllOf`

NewServiceMapPanelAllOf instantiates a new ServiceMapPanelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMapPanelAllOfWithDefaults

`func NewServiceMapPanelAllOfWithDefaults() *ServiceMapPanelAllOf`

NewServiceMapPanelAllOfWithDefaults instantiates a new ServiceMapPanelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ServiceMapPanelAllOf) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ServiceMapPanelAllOf) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ServiceMapPanelAllOf) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ServiceMapPanelAllOf) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetService

`func (o *ServiceMapPanelAllOf) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceMapPanelAllOf) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceMapPanelAllOf) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceMapPanelAllOf) HasService() bool`

HasService returns a boolean if a field has been set.

### GetShowRemoteServices

`func (o *ServiceMapPanelAllOf) GetShowRemoteServices() bool`

GetShowRemoteServices returns the ShowRemoteServices field if non-nil, zero value otherwise.

### GetShowRemoteServicesOk

`func (o *ServiceMapPanelAllOf) GetShowRemoteServicesOk() (*bool, bool)`

GetShowRemoteServicesOk returns a tuple with the ShowRemoteServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRemoteServices

`func (o *ServiceMapPanelAllOf) SetShowRemoteServices(v bool)`

SetShowRemoteServices sets ShowRemoteServices field to given value.

### HasShowRemoteServices

`func (o *ServiceMapPanelAllOf) HasShowRemoteServices() bool`

HasShowRemoteServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


