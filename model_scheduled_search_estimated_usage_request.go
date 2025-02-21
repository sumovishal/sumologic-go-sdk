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

// checks if the ScheduledSearchEstimatedUsageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledSearchEstimatedUsageRequest{}

// ScheduledSearchEstimatedUsageRequest struct for ScheduledSearchEstimatedUsageRequest
type ScheduledSearchEstimatedUsageRequest struct {
	// The text of a logs search query.
	QueryString string `json:"queryString"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// Cron-like expression specifying the search's schedule. Field scheduleType must be set to \"Custom\", otherwise, scheduleType takes precedence over cronSchedule.
	CronSchedule *string `json:"cronSchedule,omitempty"`
	// Run schedule of the scheduled search. Set to \"Custom\" to specify the schedule with a CRON expression. Please note that  with Custom, 1Day and 1Week schedule types you need to provide the corresponding cron expression to determine when to  actually run the search. e.g. Sample Valid Cron for 1Day is \"0 0 16 ? * 2-6 *\". Possible schedule types are:   - `RealTime`   - `15Minutes`   - `1Hour`   - `2Hours`   - `4Hours`   - `6Hours`   - `8Hours`   - `12Hours`   - `1Day`   - `1Week`   - `Custom`
	ScheduleType string `json:"scheduleType" validate:"regexp=^(RealTime|15Minutes|1Hour|2Hours|4Hours|6Hours|8Hours|12Hours|1Day|1Week|Custom)$"`
	// Set it to true to run the search using receipt time. By default, searches do not run by receipt time.
	ByReceiptTime *bool `json:"byReceiptTime,omitempty"`
	// An array of search query parameter objects.
	QueryParameters []QueryParameterSyncDefinition `json:"queryParameters,omitempty"`
	// Time zone identifier for the estimates. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	TimeZone string `json:"timeZone"`
	AdditionalProperties map[string]interface{}
}

type _ScheduledSearchEstimatedUsageRequest ScheduledSearchEstimatedUsageRequest

// NewScheduledSearchEstimatedUsageRequest instantiates a new ScheduledSearchEstimatedUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledSearchEstimatedUsageRequest(queryString string, timeRange ResolvableTimeRange, scheduleType string, timeZone string) *ScheduledSearchEstimatedUsageRequest {
	this := ScheduledSearchEstimatedUsageRequest{}
	this.QueryString = queryString
	this.TimeRange = timeRange
	this.ScheduleType = scheduleType
	var byReceiptTime bool = false
	this.ByReceiptTime = &byReceiptTime
	this.TimeZone = timeZone
	return &this
}

// NewScheduledSearchEstimatedUsageRequestWithDefaults instantiates a new ScheduledSearchEstimatedUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledSearchEstimatedUsageRequestWithDefaults() *ScheduledSearchEstimatedUsageRequest {
	this := ScheduledSearchEstimatedUsageRequest{}
	var byReceiptTime bool = false
	this.ByReceiptTime = &byReceiptTime
	return &this
}

// GetQueryString returns the QueryString field value
func (o *ScheduledSearchEstimatedUsageRequest) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearchEstimatedUsageRequest) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *ScheduledSearchEstimatedUsageRequest) SetQueryString(v string) {
	o.QueryString = v
}

// GetTimeRange returns the TimeRange field value
func (o *ScheduledSearchEstimatedUsageRequest) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearchEstimatedUsageRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *ScheduledSearchEstimatedUsageRequest) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetCronSchedule returns the CronSchedule field value if set, zero value otherwise.
func (o *ScheduledSearchEstimatedUsageRequest) GetCronSchedule() string {
	if o == nil || IsNil(o.CronSchedule) {
		var ret string
		return ret
	}
	return *o.CronSchedule
}

// GetCronScheduleOk returns a tuple with the CronSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchEstimatedUsageRequest) GetCronScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.CronSchedule) {
		return nil, false
	}
	return o.CronSchedule, true
}

// HasCronSchedule returns a boolean if a field has been set.
func (o *ScheduledSearchEstimatedUsageRequest) HasCronSchedule() bool {
	if o != nil && !IsNil(o.CronSchedule) {
		return true
	}

	return false
}

// SetCronSchedule gets a reference to the given string and assigns it to the CronSchedule field.
func (o *ScheduledSearchEstimatedUsageRequest) SetCronSchedule(v string) {
	o.CronSchedule = &v
}

// GetScheduleType returns the ScheduleType field value
func (o *ScheduledSearchEstimatedUsageRequest) GetScheduleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearchEstimatedUsageRequest) GetScheduleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleType, true
}

// SetScheduleType sets field value
func (o *ScheduledSearchEstimatedUsageRequest) SetScheduleType(v string) {
	o.ScheduleType = v
}

// GetByReceiptTime returns the ByReceiptTime field value if set, zero value otherwise.
func (o *ScheduledSearchEstimatedUsageRequest) GetByReceiptTime() bool {
	if o == nil || IsNil(o.ByReceiptTime) {
		var ret bool
		return ret
	}
	return *o.ByReceiptTime
}

// GetByReceiptTimeOk returns a tuple with the ByReceiptTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchEstimatedUsageRequest) GetByReceiptTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.ByReceiptTime) {
		return nil, false
	}
	return o.ByReceiptTime, true
}

// HasByReceiptTime returns a boolean if a field has been set.
func (o *ScheduledSearchEstimatedUsageRequest) HasByReceiptTime() bool {
	if o != nil && !IsNil(o.ByReceiptTime) {
		return true
	}

	return false
}

// SetByReceiptTime gets a reference to the given bool and assigns it to the ByReceiptTime field.
func (o *ScheduledSearchEstimatedUsageRequest) SetByReceiptTime(v bool) {
	o.ByReceiptTime = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *ScheduledSearchEstimatedUsageRequest) GetQueryParameters() []QueryParameterSyncDefinition {
	if o == nil || IsNil(o.QueryParameters) {
		var ret []QueryParameterSyncDefinition
		return ret
	}
	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchEstimatedUsageRequest) GetQueryParametersOk() ([]QueryParameterSyncDefinition, bool) {
	if o == nil || IsNil(o.QueryParameters) {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *ScheduledSearchEstimatedUsageRequest) HasQueryParameters() bool {
	if o != nil && !IsNil(o.QueryParameters) {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given []QueryParameterSyncDefinition and assigns it to the QueryParameters field.
func (o *ScheduledSearchEstimatedUsageRequest) SetQueryParameters(v []QueryParameterSyncDefinition) {
	o.QueryParameters = v
}

// GetTimeZone returns the TimeZone field value
func (o *ScheduledSearchEstimatedUsageRequest) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearchEstimatedUsageRequest) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *ScheduledSearchEstimatedUsageRequest) SetTimeZone(v string) {
	o.TimeZone = v
}

func (o ScheduledSearchEstimatedUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledSearchEstimatedUsageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryString"] = o.QueryString
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.CronSchedule) {
		toSerialize["cronSchedule"] = o.CronSchedule
	}
	toSerialize["scheduleType"] = o.ScheduleType
	if !IsNil(o.ByReceiptTime) {
		toSerialize["byReceiptTime"] = o.ByReceiptTime
	}
	if !IsNil(o.QueryParameters) {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	toSerialize["timeZone"] = o.TimeZone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduledSearchEstimatedUsageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryString",
		"timeRange",
		"scheduleType",
		"timeZone",
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

	varScheduledSearchEstimatedUsageRequest := _ScheduledSearchEstimatedUsageRequest{}

	err = json.Unmarshal(data, &varScheduledSearchEstimatedUsageRequest)

	if err != nil {
		return err
	}

	*o = ScheduledSearchEstimatedUsageRequest(varScheduledSearchEstimatedUsageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "queryString")
		delete(additionalProperties, "timeRange")
		delete(additionalProperties, "cronSchedule")
		delete(additionalProperties, "scheduleType")
		delete(additionalProperties, "byReceiptTime")
		delete(additionalProperties, "queryParameters")
		delete(additionalProperties, "timeZone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduledSearchEstimatedUsageRequest struct {
	value *ScheduledSearchEstimatedUsageRequest
	isSet bool
}

func (v NullableScheduledSearchEstimatedUsageRequest) Get() *ScheduledSearchEstimatedUsageRequest {
	return v.value
}

func (v *NullableScheduledSearchEstimatedUsageRequest) Set(val *ScheduledSearchEstimatedUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledSearchEstimatedUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledSearchEstimatedUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledSearchEstimatedUsageRequest(val *ScheduledSearchEstimatedUsageRequest) *NullableScheduledSearchEstimatedUsageRequest {
	return &NullableScheduledSearchEstimatedUsageRequest{value: val, isSet: true}
}

func (v NullableScheduledSearchEstimatedUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledSearchEstimatedUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


