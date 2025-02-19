# AccessKeyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | **bool** | Indicates whether the access key is disabled or not. | 
**CorsHeaders** | Pointer to **[]string** | An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don&#39;t match any entry in the allowlist. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes assigned to the key. &lt;br&gt;&lt;br&gt; Note: Updates to scopes will take up to 5m to reflect due to caching in the system. ### Alerting   - adminMonitorsV2   - viewMonitorsV2   - manageMonitorsV2  ### Data Management   - manageApps   - viewCollectors   - manageCollectors   - viewConnections   - manageConnections   - contentAdmin   - viewFieldExtractionRules   - manageFieldExtractionRules               - viewFields   - manageFields   - manageBudgets   - viewLibrary   - manageLibrary   - viewPartitions   - managePartitions   - manageS3DataForwarding   - viewScheduledViews   - manageScheduledViews   - manageTokens  ### Logs   - runLogSearch  ### Metrics   - runMetricsQuery   ### Reliability Management   - viewSlos   - manageSlos  ### Security   - manageAccessKeys   - viewPersonalAccessKeys   - managePersonalAccessKeys  ### UserManagement   - viewUsersAndRoles   - manageUsersAndRoles | [optional] 

## Methods

### NewAccessKeyUpdateRequest

`func NewAccessKeyUpdateRequest(disabled bool, ) *AccessKeyUpdateRequest`

NewAccessKeyUpdateRequest instantiates a new AccessKeyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyUpdateRequestWithDefaults

`func NewAccessKeyUpdateRequestWithDefaults() *AccessKeyUpdateRequest`

NewAccessKeyUpdateRequestWithDefaults instantiates a new AccessKeyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *AccessKeyUpdateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AccessKeyUpdateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AccessKeyUpdateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetCorsHeaders

`func (o *AccessKeyUpdateRequest) GetCorsHeaders() []string`

GetCorsHeaders returns the CorsHeaders field if non-nil, zero value otherwise.

### GetCorsHeadersOk

`func (o *AccessKeyUpdateRequest) GetCorsHeadersOk() (*[]string, bool)`

GetCorsHeadersOk returns a tuple with the CorsHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsHeaders

`func (o *AccessKeyUpdateRequest) SetCorsHeaders(v []string)`

SetCorsHeaders sets CorsHeaders field to given value.

### HasCorsHeaders

`func (o *AccessKeyUpdateRequest) HasCorsHeaders() bool`

HasCorsHeaders returns a boolean if a field has been set.

### GetScopes

`func (o *AccessKeyUpdateRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessKeyUpdateRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessKeyUpdateRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AccessKeyUpdateRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


