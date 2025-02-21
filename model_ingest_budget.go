/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on Collector and Search Job APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> KR </td>       <td> https://api.kr.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the IngestBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestBudget{}

// IngestBudget struct for IngestBudget
type IngestBudget struct {
	// Display name of the ingest budget.
	Name string `json:"name"`
	// Custom field value that is used to assign Collectors to the ingest budget.
	FieldValue string `json:"fieldValue"`
	// Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit.
	CapacityBytes int64 `json:"capacityBytes"`
	// Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
	// Reset time of the ingest budget in HH:MM format.
	ResetTime string `json:"resetTime"`
	// Description of the ingest budget.
	Description *string `json:"description,omitempty"`
	// Action to take when ingest budget's capacity is reached. All actions are audited. Supported values are:   * `stopCollecting`   * `keepCollecting`
	Action string `json:"action" validate:"regexp=^(keepCollecting|stopCollecting)$"`
	// The threshold as a percentage of when an ingest budget's capacity usage is logged in the Audit Index.
	AuditThreshold *int32 `json:"auditThreshold,omitempty"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt NullableTime `json:"createdAt"`
	CreatedByUser UserInfo `json:"createdByUser"`
	// Last modification timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	ModifiedAt NullableTime `json:"modifiedAt"`
	ModifiedByUser UserInfo `json:"modifiedByUser"`
	// Unique identifier for the ingest budget.
	Id string `json:"id"`
	// Current usage since the last reset, in bytes.
	UsageBytes *int64 `json:"usageBytes,omitempty"`
	// Status of the current usage. Can be `Normal`, `Approaching`, `Exceeded`, or `Unknown` (unable to retrieve usage).
	UsageStatus *string `json:"usageStatus,omitempty"`
	// Number of collectors assigned to the ingest budget.
	NumberOfCollectors *int64 `json:"numberOfCollectors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IngestBudget IngestBudget

// NewIngestBudget instantiates a new IngestBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestBudget(name string, fieldValue string, capacityBytes int64, timezone string, resetTime string, action string, createdAt NullableTime, createdByUser UserInfo, modifiedAt NullableTime, modifiedByUser UserInfo, id string) *IngestBudget {
	this := IngestBudget{}
	this.Name = name
	this.FieldValue = fieldValue
	this.CapacityBytes = capacityBytes
	this.Timezone = timezone
	this.ResetTime = resetTime
	this.Action = action
	this.CreatedAt = createdAt
	this.CreatedByUser = createdByUser
	this.ModifiedAt = modifiedAt
	this.ModifiedByUser = modifiedByUser
	this.Id = id
	return &this
}

// NewIngestBudgetWithDefaults instantiates a new IngestBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestBudgetWithDefaults() *IngestBudget {
	this := IngestBudget{}
	return &this
}

// GetName returns the Name field value
func (o *IngestBudget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IngestBudget) SetName(v string) {
	o.Name = v
}

// GetFieldValue returns the FieldValue field value
func (o *IngestBudget) GetFieldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetFieldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldValue, true
}

// SetFieldValue sets field value
func (o *IngestBudget) SetFieldValue(v string) {
	o.FieldValue = v
}

// GetCapacityBytes returns the CapacityBytes field value
func (o *IngestBudget) GetCapacityBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CapacityBytes
}

// GetCapacityBytesOk returns a tuple with the CapacityBytes field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetCapacityBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityBytes, true
}

// SetCapacityBytes sets field value
func (o *IngestBudget) SetCapacityBytes(v int64) {
	o.CapacityBytes = v
}

// GetTimezone returns the Timezone field value
func (o *IngestBudget) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *IngestBudget) SetTimezone(v string) {
	o.Timezone = v
}

// GetResetTime returns the ResetTime field value
func (o *IngestBudget) GetResetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetTime
}

// GetResetTimeOk returns a tuple with the ResetTime field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetResetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetTime, true
}

// SetResetTime sets field value
func (o *IngestBudget) SetResetTime(v string) {
	o.ResetTime = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestBudget) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestBudget) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestBudget) SetDescription(v string) {
	o.Description = &v
}

// GetAction returns the Action field value
func (o *IngestBudget) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *IngestBudget) SetAction(v string) {
	o.Action = v
}

// GetAuditThreshold returns the AuditThreshold field value if set, zero value otherwise.
func (o *IngestBudget) GetAuditThreshold() int32 {
	if o == nil || IsNil(o.AuditThreshold) {
		var ret int32
		return ret
	}
	return *o.AuditThreshold
}

// GetAuditThresholdOk returns a tuple with the AuditThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetAuditThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditThreshold) {
		return nil, false
	}
	return o.AuditThreshold, true
}

// HasAuditThreshold returns a boolean if a field has been set.
func (o *IngestBudget) HasAuditThreshold() bool {
	if o != nil && !IsNil(o.AuditThreshold) {
		return true
	}

	return false
}

// SetAuditThreshold gets a reference to the given int32 and assigns it to the AuditThreshold field.
func (o *IngestBudget) SetAuditThreshold(v int32) {
	o.AuditThreshold = &v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IngestBudget) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestBudget) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *IngestBudget) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// GetCreatedByUser returns the CreatedByUser field value
func (o *IngestBudget) GetCreatedByUser() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}

	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetCreatedByUserOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// SetCreatedByUser sets field value
func (o *IngestBudget) SetCreatedByUser(v UserInfo) {
	o.CreatedByUser = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IngestBudget) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestBudget) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *IngestBudget) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// GetModifiedByUser returns the ModifiedByUser field value
func (o *IngestBudget) GetModifiedByUser() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}

	return o.ModifiedByUser
}

// GetModifiedByUserOk returns a tuple with the ModifiedByUser field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetModifiedByUserOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedByUser, true
}

// SetModifiedByUser sets field value
func (o *IngestBudget) SetModifiedByUser(v UserInfo) {
	o.ModifiedByUser = v
}

// GetId returns the Id field value
func (o *IngestBudget) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IngestBudget) SetId(v string) {
	o.Id = v
}

// GetUsageBytes returns the UsageBytes field value if set, zero value otherwise.
func (o *IngestBudget) GetUsageBytes() int64 {
	if o == nil || IsNil(o.UsageBytes) {
		var ret int64
		return ret
	}
	return *o.UsageBytes
}

// GetUsageBytesOk returns a tuple with the UsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetUsageBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.UsageBytes) {
		return nil, false
	}
	return o.UsageBytes, true
}

// HasUsageBytes returns a boolean if a field has been set.
func (o *IngestBudget) HasUsageBytes() bool {
	if o != nil && !IsNil(o.UsageBytes) {
		return true
	}

	return false
}

// SetUsageBytes gets a reference to the given int64 and assigns it to the UsageBytes field.
func (o *IngestBudget) SetUsageBytes(v int64) {
	o.UsageBytes = &v
}

// GetUsageStatus returns the UsageStatus field value if set, zero value otherwise.
func (o *IngestBudget) GetUsageStatus() string {
	if o == nil || IsNil(o.UsageStatus) {
		var ret string
		return ret
	}
	return *o.UsageStatus
}

// GetUsageStatusOk returns a tuple with the UsageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetUsageStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UsageStatus) {
		return nil, false
	}
	return o.UsageStatus, true
}

// HasUsageStatus returns a boolean if a field has been set.
func (o *IngestBudget) HasUsageStatus() bool {
	if o != nil && !IsNil(o.UsageStatus) {
		return true
	}

	return false
}

// SetUsageStatus gets a reference to the given string and assigns it to the UsageStatus field.
func (o *IngestBudget) SetUsageStatus(v string) {
	o.UsageStatus = &v
}

// GetNumberOfCollectors returns the NumberOfCollectors field value if set, zero value otherwise.
func (o *IngestBudget) GetNumberOfCollectors() int64 {
	if o == nil || IsNil(o.NumberOfCollectors) {
		var ret int64
		return ret
	}
	return *o.NumberOfCollectors
}

// GetNumberOfCollectorsOk returns a tuple with the NumberOfCollectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudget) GetNumberOfCollectorsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfCollectors) {
		return nil, false
	}
	return o.NumberOfCollectors, true
}

// HasNumberOfCollectors returns a boolean if a field has been set.
func (o *IngestBudget) HasNumberOfCollectors() bool {
	if o != nil && !IsNil(o.NumberOfCollectors) {
		return true
	}

	return false
}

// SetNumberOfCollectors gets a reference to the given int64 and assigns it to the NumberOfCollectors field.
func (o *IngestBudget) SetNumberOfCollectors(v int64) {
	o.NumberOfCollectors = &v
}

func (o IngestBudget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["fieldValue"] = o.FieldValue
	toSerialize["capacityBytes"] = o.CapacityBytes
	toSerialize["timezone"] = o.Timezone
	toSerialize["resetTime"] = o.ResetTime
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["action"] = o.Action
	if !IsNil(o.AuditThreshold) {
		toSerialize["auditThreshold"] = o.AuditThreshold
	}
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["createdByUser"] = o.CreatedByUser
	toSerialize["modifiedAt"] = o.ModifiedAt.Get()
	toSerialize["modifiedByUser"] = o.ModifiedByUser
	toSerialize["id"] = o.Id
	if !IsNil(o.UsageBytes) {
		toSerialize["usageBytes"] = o.UsageBytes
	}
	if !IsNil(o.UsageStatus) {
		toSerialize["usageStatus"] = o.UsageStatus
	}
	if !IsNil(o.NumberOfCollectors) {
		toSerialize["numberOfCollectors"] = o.NumberOfCollectors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IngestBudget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fieldValue",
		"capacityBytes",
		"timezone",
		"resetTime",
		"action",
		"createdAt",
		"createdByUser",
		"modifiedAt",
		"modifiedByUser",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIngestBudget := _IngestBudget{}

	err = json.Unmarshal(data, &varIngestBudget)

	if err != nil {
		return err
	}

	*o = IngestBudget(varIngestBudget)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldValue")
		delete(additionalProperties, "capacityBytes")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "resetTime")
		delete(additionalProperties, "description")
		delete(additionalProperties, "action")
		delete(additionalProperties, "auditThreshold")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "createdByUser")
		delete(additionalProperties, "modifiedAt")
		delete(additionalProperties, "modifiedByUser")
		delete(additionalProperties, "id")
		delete(additionalProperties, "usageBytes")
		delete(additionalProperties, "usageStatus")
		delete(additionalProperties, "numberOfCollectors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIngestBudget struct {
	value *IngestBudget
	isSet bool
}

func (v NullableIngestBudget) Get() *IngestBudget {
	return v.value
}

func (v *NullableIngestBudget) Set(val *IngestBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestBudget(val *IngestBudget) *NullableIngestBudget {
	return &NullableIngestBudget{value: val, isSet: true}
}

func (v NullableIngestBudget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


