# UpdateDataForwardingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | Pointer to **string** | Data forwarding destination id. | [optional] 
**Enabled** | Pointer to **bool** | True when the data forwarding rule is enabled. | [optional] 
**FileFormat** | Pointer to **string** | Specify the path prefix to a directory in the S3 bucket and how to format the file name. | [optional] 
**PayloadSchema** | Pointer to **string** | Schema for the payload. Default value of the payload schema is \&quot;allFields\&quot; for scheduled view, and \&quot;builtInFields\&quot; for partition. \&quot;raw\&quot; payloadSchema should be used in conjunction with \&quot;text\&quot; format and vice-versa. | [optional] 
**Format** | Pointer to **string** | Format of the payload. Default format will be \&quot;csv\&quot;. \&quot;text\&quot; format should be used in conjunction with \&quot;raw\&quot; payloadSchema and vice-versa. | [optional] 

## Methods

### NewUpdateDataForwardingRule

`func NewUpdateDataForwardingRule() *UpdateDataForwardingRule`

NewUpdateDataForwardingRule instantiates a new UpdateDataForwardingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataForwardingRuleWithDefaults

`func NewUpdateDataForwardingRuleWithDefaults() *UpdateDataForwardingRule`

NewUpdateDataForwardingRuleWithDefaults instantiates a new UpdateDataForwardingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *UpdateDataForwardingRule) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *UpdateDataForwardingRule) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *UpdateDataForwardingRule) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *UpdateDataForwardingRule) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateDataForwardingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateDataForwardingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateDataForwardingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateDataForwardingRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFileFormat

`func (o *UpdateDataForwardingRule) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *UpdateDataForwardingRule) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *UpdateDataForwardingRule) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *UpdateDataForwardingRule) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetPayloadSchema

`func (o *UpdateDataForwardingRule) GetPayloadSchema() string`

GetPayloadSchema returns the PayloadSchema field if non-nil, zero value otherwise.

### GetPayloadSchemaOk

`func (o *UpdateDataForwardingRule) GetPayloadSchemaOk() (*string, bool)`

GetPayloadSchemaOk returns a tuple with the PayloadSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadSchema

`func (o *UpdateDataForwardingRule) SetPayloadSchema(v string)`

SetPayloadSchema sets PayloadSchema field to given value.

### HasPayloadSchema

`func (o *UpdateDataForwardingRule) HasPayloadSchema() bool`

HasPayloadSchema returns a boolean if a field has been set.

### GetFormat

`func (o *UpdateDataForwardingRule) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UpdateDataForwardingRule) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UpdateDataForwardingRule) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UpdateDataForwardingRule) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


