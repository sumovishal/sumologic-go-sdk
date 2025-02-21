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

// checks if the ScanBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScanBudget{}

// ScanBudget struct for ScanBudget
type ScanBudget struct {
	// Name of the budget.
	Name string `json:"name"`
	// Capacity of the budget.
	Capacity int64 `json:"capacity"`
	// Unit of the budget.
	Unit string `json:"unit" validate:"regexp=^(GB|MB|TB|KB)$"`
	// Type of the budget.
	BudgetType string `json:"budgetType" validate:"regexp=^(ScanBudget)$"`
	Scope ScanBudgetScope `json:"scope"`
	// Window of the budget. Use Daily/Weekly/Monthly for creating a time based budget (beta)
	Window string `json:"window" validate:"regexp=^(Query|Daily|Weekly|Monthly)$"`
	// Grouping of the budget.
	ApplicableOn string `json:"applicableOn" validate:"regexp=^(PerEntity|Sum)$"`
	// Grouping Entity of the budget.
	GroupBy string `json:"groupBy" validate:"regexp=^(User)$"`
	// Action to be taken if the budget is breached
	Action string `json:"action" validate:"regexp=^(StopForeGroundScan|Warn)$"`
	// Signifies the state of the budget. (Active/Inactive)
	Status *string `json:"status,omitempty" validate:"regexp=^(active|inactive)$"`
	// Id of the budget.
	Id string `json:"id"`
	// Org Id of the org for the budget.
	OrgId string `json:"orgId"`
	// Reset time of the time based scan budget in HH:MM format
	ResetTime string `json:"resetTime"`
	// Time zone of the reset time for the time based scan budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	ResetTimeZone string `json:"resetTimeZone"`
	// The day of the week when the budget resets, applicable for time based budgets with a Weekly window. Must be a valid day of the week.
	ResetDayOfWeek string `json:"resetDayOfWeek" validate:"regexp=^(MONDAY|TUESDAY|WEDNESDAY|THURSDAY|FRIDAY|SATURDAY|SUNDAY)$"`
	// The date of the month when the budget resets, applicable for time based budgets with a Monthly window. Must be a valid day of the month (1-28).
	ResetDateOfMonth int32 `json:"resetDateOfMonth"`
	// Date & time when budget was created.
	CreatedAt time.Time `json:"createdAt"`
	// Id of the user who created the budget.
	CreatedBy string `json:"createdBy"`
	// Date & time when budget was last modified.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Id of the user who last modified the budget.
	ModifiedBy string `json:"modifiedBy"`
	AdditionalProperties map[string]interface{}
}

type _ScanBudget ScanBudget

// NewScanBudget instantiates a new ScanBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScanBudget(name string, capacity int64, unit string, budgetType string, scope ScanBudgetScope, window string, applicableOn string, groupBy string, action string, id string, orgId string, resetTime string, resetTimeZone string, resetDayOfWeek string, resetDateOfMonth int32, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *ScanBudget {
	this := ScanBudget{}
	this.Name = name
	this.Capacity = capacity
	this.Unit = unit
	this.BudgetType = budgetType
	this.Scope = scope
	this.Window = window
	this.ApplicableOn = applicableOn
	this.GroupBy = groupBy
	this.Action = action
	this.Id = id
	this.OrgId = orgId
	this.ResetTime = resetTime
	this.ResetTimeZone = resetTimeZone
	this.ResetDayOfWeek = resetDayOfWeek
	this.ResetDateOfMonth = resetDateOfMonth
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	return &this
}

// NewScanBudgetWithDefaults instantiates a new ScanBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScanBudgetWithDefaults() *ScanBudget {
	this := ScanBudget{}
	var resetTime string = "00:00"
	this.ResetTime = resetTime
	var resetTimeZone string = "Etc/UTC"
	this.ResetTimeZone = resetTimeZone
	var resetDayOfWeek string = "MONDAY"
	this.ResetDayOfWeek = resetDayOfWeek
	var resetDateOfMonth int32 = 1
	this.ResetDateOfMonth = resetDateOfMonth
	return &this
}

// GetName returns the Name field value
func (o *ScanBudget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScanBudget) SetName(v string) {
	o.Name = v
}

// GetCapacity returns the Capacity field value
func (o *ScanBudget) GetCapacity() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetCapacityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *ScanBudget) SetCapacity(v int64) {
	o.Capacity = v
}

// GetUnit returns the Unit field value
func (o *ScanBudget) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ScanBudget) SetUnit(v string) {
	o.Unit = v
}

// GetBudgetType returns the BudgetType field value
func (o *ScanBudget) GetBudgetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetBudgetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *ScanBudget) SetBudgetType(v string) {
	o.BudgetType = v
}

// GetScope returns the Scope field value
func (o *ScanBudget) GetScope() ScanBudgetScope {
	if o == nil {
		var ret ScanBudgetScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetScopeOk() (*ScanBudgetScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *ScanBudget) SetScope(v ScanBudgetScope) {
	o.Scope = v
}

// GetWindow returns the Window field value
func (o *ScanBudget) GetWindow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetWindowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *ScanBudget) SetWindow(v string) {
	o.Window = v
}

// GetApplicableOn returns the ApplicableOn field value
func (o *ScanBudget) GetApplicableOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicableOn
}

// GetApplicableOnOk returns a tuple with the ApplicableOn field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetApplicableOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicableOn, true
}

// SetApplicableOn sets field value
func (o *ScanBudget) SetApplicableOn(v string) {
	o.ApplicableOn = v
}

// GetGroupBy returns the GroupBy field value
func (o *ScanBudget) GetGroupBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetGroupByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value
func (o *ScanBudget) SetGroupBy(v string) {
	o.GroupBy = v
}

// GetAction returns the Action field value
func (o *ScanBudget) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ScanBudget) SetAction(v string) {
	o.Action = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ScanBudget) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ScanBudget) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ScanBudget) SetStatus(v string) {
	o.Status = &v
}

// GetId returns the Id field value
func (o *ScanBudget) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScanBudget) SetId(v string) {
	o.Id = v
}

// GetOrgId returns the OrgId field value
func (o *ScanBudget) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ScanBudget) SetOrgId(v string) {
	o.OrgId = v
}

// GetResetTime returns the ResetTime field value
func (o *ScanBudget) GetResetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetTime
}

// GetResetTimeOk returns a tuple with the ResetTime field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetResetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetTime, true
}

// SetResetTime sets field value
func (o *ScanBudget) SetResetTime(v string) {
	o.ResetTime = v
}

// GetResetTimeZone returns the ResetTimeZone field value
func (o *ScanBudget) GetResetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetTimeZone
}

// GetResetTimeZoneOk returns a tuple with the ResetTimeZone field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetResetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetTimeZone, true
}

// SetResetTimeZone sets field value
func (o *ScanBudget) SetResetTimeZone(v string) {
	o.ResetTimeZone = v
}

// GetResetDayOfWeek returns the ResetDayOfWeek field value
func (o *ScanBudget) GetResetDayOfWeek() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetDayOfWeek
}

// GetResetDayOfWeekOk returns a tuple with the ResetDayOfWeek field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetResetDayOfWeekOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetDayOfWeek, true
}

// SetResetDayOfWeek sets field value
func (o *ScanBudget) SetResetDayOfWeek(v string) {
	o.ResetDayOfWeek = v
}

// GetResetDateOfMonth returns the ResetDateOfMonth field value
func (o *ScanBudget) GetResetDateOfMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResetDateOfMonth
}

// GetResetDateOfMonthOk returns a tuple with the ResetDateOfMonth field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetResetDateOfMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetDateOfMonth, true
}

// SetResetDateOfMonth sets field value
func (o *ScanBudget) SetResetDateOfMonth(v int32) {
	o.ResetDateOfMonth = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ScanBudget) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ScanBudget) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ScanBudget) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ScanBudget) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *ScanBudget) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *ScanBudget) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *ScanBudget) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *ScanBudget) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *ScanBudget) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

func (o ScanBudget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScanBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["capacity"] = o.Capacity
	toSerialize["unit"] = o.Unit
	toSerialize["budgetType"] = o.BudgetType
	toSerialize["scope"] = o.Scope
	toSerialize["window"] = o.Window
	toSerialize["applicableOn"] = o.ApplicableOn
	toSerialize["groupBy"] = o.GroupBy
	toSerialize["action"] = o.Action
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["id"] = o.Id
	toSerialize["orgId"] = o.OrgId
	toSerialize["resetTime"] = o.ResetTime
	toSerialize["resetTimeZone"] = o.ResetTimeZone
	toSerialize["resetDayOfWeek"] = o.ResetDayOfWeek
	toSerialize["resetDateOfMonth"] = o.ResetDateOfMonth
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScanBudget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"capacity",
		"unit",
		"budgetType",
		"scope",
		"window",
		"applicableOn",
		"groupBy",
		"action",
		"id",
		"orgId",
		"resetTime",
		"resetTimeZone",
		"resetDayOfWeek",
		"resetDateOfMonth",
		"createdAt",
		"createdBy",
		"modifiedAt",
		"modifiedBy",
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

	varScanBudget := _ScanBudget{}

	err = json.Unmarshal(data, &varScanBudget)

	if err != nil {
		return err
	}

	*o = ScanBudget(varScanBudget)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "capacity")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "budgetType")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "window")
		delete(additionalProperties, "applicableOn")
		delete(additionalProperties, "groupBy")
		delete(additionalProperties, "action")
		delete(additionalProperties, "status")
		delete(additionalProperties, "id")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "resetTime")
		delete(additionalProperties, "resetTimeZone")
		delete(additionalProperties, "resetDayOfWeek")
		delete(additionalProperties, "resetDateOfMonth")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "modifiedAt")
		delete(additionalProperties, "modifiedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScanBudget struct {
	value *ScanBudget
	isSet bool
}

func (v NullableScanBudget) Get() *ScanBudget {
	return v.value
}

func (v *NullableScanBudget) Set(val *ScanBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableScanBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableScanBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScanBudget(val *ScanBudget) *NullableScanBudget {
	return &NullableScanBudget{value: val, isSet: true}
}

func (v NullableScanBudget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScanBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


