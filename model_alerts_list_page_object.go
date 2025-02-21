/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on Collector and Search Job APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> KR </td>       <td> https://api.kr.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// checks if the AlertsListPageObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsListPageObject{}

// AlertsListPageObject Alert list page object.
type AlertsListPageObject struct {
	// Identifier of the alert.
	Id *string `json:"id,omitempty"`
	// Name of the alert.
	Name *string `json:"name,omitempty"`
	// The severity of the Alert. Valid values:   1. `Critical`   2. `Warning`   3. `MissingData`
	Severity *string `json:"severity,omitempty" validate:"regexp=^(Critical|Warning|MissingData)$"`
	// The status of the Alert. Valid values:   1. `Active`   2. `Resolved`
	Status *string `json:"status,omitempty" validate:"regexp=^(Active|Resolved)$"`
	// List of AlertEntityInfo for primary entities. The primary entity is the most concrete entity  (e.g. k8s container) that can be assigned per time series or log group,  secondary entities are the less specific ones (e.g. k8s cluster or EC2 host). 
	EntitiesInfo []AlertEntityInfo `json:"entitiesInfo,omitempty"`
	// List of secondary AlertEntityInfo for primary entities. Primary/secondary entities are explained in description for `entitiesInfo`. 
	SecondaryEntitiesInfo []AlertEntityInfo `json:"secondaryEntitiesInfo,omitempty"`
	// The number of unique result groups that have met the alert condition.
	ViolationCount *string `json:"violationCount,omitempty"`
	// The condition from the last alert violation.
	LastViolation *string `json:"lastViolation,omitempty"`
	// The current duration of the alert.
	Duration *string `json:"duration,omitempty"`
	// The creation time of the alert.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time when this alert was updated with the most recent violation.
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// True if the ARP was created while the monitor was muted
	IsMuted *bool `json:"isMuted,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AlertsListPageObject AlertsListPageObject

// NewAlertsListPageObject instantiates a new AlertsListPageObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsListPageObject() *AlertsListPageObject {
	this := AlertsListPageObject{}
	var isMuted bool = false
	this.IsMuted = &isMuted
	return &this
}

// NewAlertsListPageObjectWithDefaults instantiates a new AlertsListPageObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsListPageObjectWithDefaults() *AlertsListPageObject {
	this := AlertsListPageObject{}
	var isMuted bool = false
	this.IsMuted = &isMuted
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertsListPageObject) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertsListPageObject) SetName(v string) {
	o.Name = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AlertsListPageObject) SetSeverity(v string) {
	o.Severity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlertsListPageObject) SetStatus(v string) {
	o.Status = &v
}

// GetEntitiesInfo returns the EntitiesInfo field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetEntitiesInfo() []AlertEntityInfo {
	if o == nil || IsNil(o.EntitiesInfo) {
		var ret []AlertEntityInfo
		return ret
	}
	return o.EntitiesInfo
}

// GetEntitiesInfoOk returns a tuple with the EntitiesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetEntitiesInfoOk() ([]AlertEntityInfo, bool) {
	if o == nil || IsNil(o.EntitiesInfo) {
		return nil, false
	}
	return o.EntitiesInfo, true
}

// HasEntitiesInfo returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasEntitiesInfo() bool {
	if o != nil && !IsNil(o.EntitiesInfo) {
		return true
	}

	return false
}

// SetEntitiesInfo gets a reference to the given []AlertEntityInfo and assigns it to the EntitiesInfo field.
func (o *AlertsListPageObject) SetEntitiesInfo(v []AlertEntityInfo) {
	o.EntitiesInfo = v
}

// GetSecondaryEntitiesInfo returns the SecondaryEntitiesInfo field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetSecondaryEntitiesInfo() []AlertEntityInfo {
	if o == nil || IsNil(o.SecondaryEntitiesInfo) {
		var ret []AlertEntityInfo
		return ret
	}
	return o.SecondaryEntitiesInfo
}

// GetSecondaryEntitiesInfoOk returns a tuple with the SecondaryEntitiesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetSecondaryEntitiesInfoOk() ([]AlertEntityInfo, bool) {
	if o == nil || IsNil(o.SecondaryEntitiesInfo) {
		return nil, false
	}
	return o.SecondaryEntitiesInfo, true
}

// HasSecondaryEntitiesInfo returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasSecondaryEntitiesInfo() bool {
	if o != nil && !IsNil(o.SecondaryEntitiesInfo) {
		return true
	}

	return false
}

// SetSecondaryEntitiesInfo gets a reference to the given []AlertEntityInfo and assigns it to the SecondaryEntitiesInfo field.
func (o *AlertsListPageObject) SetSecondaryEntitiesInfo(v []AlertEntityInfo) {
	o.SecondaryEntitiesInfo = v
}

// GetViolationCount returns the ViolationCount field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetViolationCount() string {
	if o == nil || IsNil(o.ViolationCount) {
		var ret string
		return ret
	}
	return *o.ViolationCount
}

// GetViolationCountOk returns a tuple with the ViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetViolationCountOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationCount) {
		return nil, false
	}
	return o.ViolationCount, true
}

// HasViolationCount returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasViolationCount() bool {
	if o != nil && !IsNil(o.ViolationCount) {
		return true
	}

	return false
}

// SetViolationCount gets a reference to the given string and assigns it to the ViolationCount field.
func (o *AlertsListPageObject) SetViolationCount(v string) {
	o.ViolationCount = &v
}

// GetLastViolation returns the LastViolation field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetLastViolation() string {
	if o == nil || IsNil(o.LastViolation) {
		var ret string
		return ret
	}
	return *o.LastViolation
}

// GetLastViolationOk returns a tuple with the LastViolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetLastViolationOk() (*string, bool) {
	if o == nil || IsNil(o.LastViolation) {
		return nil, false
	}
	return o.LastViolation, true
}

// HasLastViolation returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasLastViolation() bool {
	if o != nil && !IsNil(o.LastViolation) {
		return true
	}

	return false
}

// SetLastViolation gets a reference to the given string and assigns it to the LastViolation field.
func (o *AlertsListPageObject) SetLastViolation(v string) {
	o.LastViolation = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *AlertsListPageObject) SetDuration(v string) {
	o.Duration = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AlertsListPageObject) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *AlertsListPageObject) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise.
func (o *AlertsListPageObject) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted) {
		var ret bool
		return ret
	}
	return *o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsListPageObject) GetIsMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMuted) {
		return nil, false
	}
	return o.IsMuted, true
}

// HasIsMuted returns a boolean if a field has been set.
func (o *AlertsListPageObject) HasIsMuted() bool {
	if o != nil && !IsNil(o.IsMuted) {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.
func (o *AlertsListPageObject) SetIsMuted(v bool) {
	o.IsMuted = &v
}

func (o AlertsListPageObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsListPageObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.EntitiesInfo) {
		toSerialize["entitiesInfo"] = o.EntitiesInfo
	}
	if !IsNil(o.SecondaryEntitiesInfo) {
		toSerialize["secondaryEntitiesInfo"] = o.SecondaryEntitiesInfo
	}
	if !IsNil(o.ViolationCount) {
		toSerialize["violationCount"] = o.ViolationCount
	}
	if !IsNil(o.LastViolation) {
		toSerialize["lastViolation"] = o.LastViolation
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.IsMuted) {
		toSerialize["isMuted"] = o.IsMuted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AlertsListPageObject) UnmarshalJSON(data []byte) (err error) {
	varAlertsListPageObject := _AlertsListPageObject{}

	err = json.Unmarshal(data, &varAlertsListPageObject)

	if err != nil {
		return err
	}

	*o = AlertsListPageObject(varAlertsListPageObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "status")
		delete(additionalProperties, "entitiesInfo")
		delete(additionalProperties, "secondaryEntitiesInfo")
		delete(additionalProperties, "violationCount")
		delete(additionalProperties, "lastViolation")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "isMuted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAlertsListPageObject struct {
	value *AlertsListPageObject
	isSet bool
}

func (v NullableAlertsListPageObject) Get() *AlertsListPageObject {
	return v.value
}

func (v *NullableAlertsListPageObject) Set(val *AlertsListPageObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsListPageObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsListPageObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsListPageObject(val *AlertsListPageObject) *NullableAlertsListPageObject {
	return &NullableAlertsListPageObject{value: val, isSet: true}
}

func (v NullableAlertsListPageObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsListPageObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


