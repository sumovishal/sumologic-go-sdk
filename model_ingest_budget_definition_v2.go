/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on Collector and Search Job APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> KR </td>       <td> https://api.kr.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"fmt"
)

// checks if the IngestBudgetDefinitionV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestBudgetDefinitionV2{}

// IngestBudgetDefinitionV2 struct for IngestBudgetDefinitionV2
type IngestBudgetDefinitionV2 struct {
	// Display name of the ingest budget.
	Name string `json:"name"`
	// A scope is a constraint that will be used to identify the messages on which budget needs to be applied. A scope is consists of key and value separated by =. The field must be enabled in the fields table. Value supports wildcard. e.g. _sourceCategory=*prod*payment*, cluster=kafka. If the scope is defined _sourceCategory=*nginx* in this budget will be applied on messages having fields _sourceCategory=prod/nginx, _sourceCategory=dev/nginx, or _sourceCategory=dev/nginx/error
	Scope string `json:"scope"`
	// Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit. The capacity bytes unit varies based on the budgetType field. For `dailyVolume` budgetType the capacity specified is in bytes/day whereas for `minuteVolume` budgetType its bytes/min.
	CapacityBytes int64 `json:"capacityBytes"`
	// Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone *string `json:"timezone,omitempty"`
	// Reset time of the ingest budget in HH:MM format.
	ResetTime *string `json:"resetTime,omitempty"`
	// Description of the ingest budget.
	Description *string `json:"description,omitempty"`
	// Action to take when ingest budget's capacity is reached. All actions are audited. Supported values are:   * `stopCollecting`   * `keepCollecting`
	Action string `json:"action" validate:"regexp=^(keepCollecting|stopCollecting)$"`
	// The threshold as a percentage of when an ingest budget's capacity usage is logged in the Audit Index.
	AuditThreshold *int32 `json:"auditThreshold,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IngestBudgetDefinitionV2 IngestBudgetDefinitionV2

// NewIngestBudgetDefinitionV2 instantiates a new IngestBudgetDefinitionV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestBudgetDefinitionV2(name string, scope string, capacityBytes int64, action string) *IngestBudgetDefinitionV2 {
	this := IngestBudgetDefinitionV2{}
	this.Name = name
	this.Scope = scope
	this.CapacityBytes = capacityBytes
	var timezone string = "Etc/UTC"
	this.Timezone = &timezone
	var resetTime string = "00:00"
	this.ResetTime = &resetTime
	this.Action = action
	return &this
}

// NewIngestBudgetDefinitionV2WithDefaults instantiates a new IngestBudgetDefinitionV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestBudgetDefinitionV2WithDefaults() *IngestBudgetDefinitionV2 {
	this := IngestBudgetDefinitionV2{}
	var timezone string = "Etc/UTC"
	this.Timezone = &timezone
	var resetTime string = "00:00"
	this.ResetTime = &resetTime
	return &this
}

// GetName returns the Name field value
func (o *IngestBudgetDefinitionV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetDefinitionV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IngestBudgetDefinitionV2) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *IngestBudgetDefinitionV2) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetDefinitionV2) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *IngestBudgetDefinitionV2) SetScope(v string) {
	o.Scope = v
}

// GetCapacityBytes returns the CapacityBytes field value
func (o *IngestBudgetDefinitionV2) GetCapacityBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CapacityBytes
}

// GetCapacityBytesOk returns a tuple with the CapacityBytes field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetDefinitionV2) GetCapacityBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityBytes, true
}

// SetCapacityBytes sets field value
func (o *IngestBudgetDefinitionV2) SetCapacityBytes(v int64) {
	o.CapacityBytes = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *IngestBudgetDefinitionV2) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetDefinitionV2) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *IngestBudgetDefinitionV2) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *IngestBudgetDefinitionV2) SetTimezone(v string) {
	o.Timezone = &v
}

// GetResetTime returns the ResetTime field value if set, zero value otherwise.
func (o *IngestBudgetDefinitionV2) GetResetTime() string {
	if o == nil || IsNil(o.ResetTime) {
		var ret string
		return ret
	}
	return *o.ResetTime
}

// GetResetTimeOk returns a tuple with the ResetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetDefinitionV2) GetResetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ResetTime) {
		return nil, false
	}
	return o.ResetTime, true
}

// HasResetTime returns a boolean if a field has been set.
func (o *IngestBudgetDefinitionV2) HasResetTime() bool {
	if o != nil && !IsNil(o.ResetTime) {
		return true
	}

	return false
}

// SetResetTime gets a reference to the given string and assigns it to the ResetTime field.
func (o *IngestBudgetDefinitionV2) SetResetTime(v string) {
	o.ResetTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestBudgetDefinitionV2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetDefinitionV2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestBudgetDefinitionV2) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestBudgetDefinitionV2) SetDescription(v string) {
	o.Description = &v
}

// GetAction returns the Action field value
func (o *IngestBudgetDefinitionV2) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetDefinitionV2) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *IngestBudgetDefinitionV2) SetAction(v string) {
	o.Action = v
}

// GetAuditThreshold returns the AuditThreshold field value if set, zero value otherwise.
func (o *IngestBudgetDefinitionV2) GetAuditThreshold() int32 {
	if o == nil || IsNil(o.AuditThreshold) {
		var ret int32
		return ret
	}
	return *o.AuditThreshold
}

// GetAuditThresholdOk returns a tuple with the AuditThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetDefinitionV2) GetAuditThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditThreshold) {
		return nil, false
	}
	return o.AuditThreshold, true
}

// HasAuditThreshold returns a boolean if a field has been set.
func (o *IngestBudgetDefinitionV2) HasAuditThreshold() bool {
	if o != nil && !IsNil(o.AuditThreshold) {
		return true
	}

	return false
}

// SetAuditThreshold gets a reference to the given int32 and assigns it to the AuditThreshold field.
func (o *IngestBudgetDefinitionV2) SetAuditThreshold(v int32) {
	o.AuditThreshold = &v
}

func (o IngestBudgetDefinitionV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestBudgetDefinitionV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["scope"] = o.Scope
	toSerialize["capacityBytes"] = o.CapacityBytes
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.ResetTime) {
		toSerialize["resetTime"] = o.ResetTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["action"] = o.Action
	if !IsNil(o.AuditThreshold) {
		toSerialize["auditThreshold"] = o.AuditThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IngestBudgetDefinitionV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"scope",
		"capacityBytes",
		"action",
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

	varIngestBudgetDefinitionV2 := _IngestBudgetDefinitionV2{}

	err = json.Unmarshal(data, &varIngestBudgetDefinitionV2)

	if err != nil {
		return err
	}

	*o = IngestBudgetDefinitionV2(varIngestBudgetDefinitionV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "capacityBytes")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "resetTime")
		delete(additionalProperties, "description")
		delete(additionalProperties, "action")
		delete(additionalProperties, "auditThreshold")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIngestBudgetDefinitionV2 struct {
	value *IngestBudgetDefinitionV2
	isSet bool
}

func (v NullableIngestBudgetDefinitionV2) Get() *IngestBudgetDefinitionV2 {
	return v.value
}

func (v *NullableIngestBudgetDefinitionV2) Set(val *IngestBudgetDefinitionV2) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestBudgetDefinitionV2) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestBudgetDefinitionV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestBudgetDefinitionV2(val *IngestBudgetDefinitionV2) *NullableIngestBudgetDefinitionV2 {
	return &NullableIngestBudgetDefinitionV2{value: val, isSet: true}
}

func (v NullableIngestBudgetDefinitionV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestBudgetDefinitionV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


