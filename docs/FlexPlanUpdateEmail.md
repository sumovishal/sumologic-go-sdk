# FlexPlanUpdateEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailId** | **string** | email id on which support team will contact on | 
**PhoneNumber** | Pointer to **string** | contact number on which support team can call user | [optional] 
**Details** | Pointer to **string** | option details the user might want to inform | [optional] 

## Methods

### NewFlexPlanUpdateEmail

`func NewFlexPlanUpdateEmail(emailId string, ) *FlexPlanUpdateEmail`

NewFlexPlanUpdateEmail instantiates a new FlexPlanUpdateEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexPlanUpdateEmailWithDefaults

`func NewFlexPlanUpdateEmailWithDefaults() *FlexPlanUpdateEmail`

NewFlexPlanUpdateEmailWithDefaults instantiates a new FlexPlanUpdateEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailId

`func (o *FlexPlanUpdateEmail) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *FlexPlanUpdateEmail) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *FlexPlanUpdateEmail) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.


### GetPhoneNumber

`func (o *FlexPlanUpdateEmail) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *FlexPlanUpdateEmail) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *FlexPlanUpdateEmail) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *FlexPlanUpdateEmail) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetDetails

`func (o *FlexPlanUpdateEmail) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FlexPlanUpdateEmail) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FlexPlanUpdateEmail) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FlexPlanUpdateEmail) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


