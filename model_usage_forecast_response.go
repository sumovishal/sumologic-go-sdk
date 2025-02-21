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

// checks if the UsageForecastResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageForecastResponse{}

// UsageForecastResponse Usage forecast for the organization.
type UsageForecastResponse struct {
	// Average credit usage per day till now.
	AverageUsage *float64 `json:"averageUsage,omitempty"`
	// Percentage of total credits used till date.
	UsagePercentage *float64 `json:"usagePercentage,omitempty"`
	// Total expected usage by the end of contract period.
	ForecastedUsage *float64 `json:"forecastedUsage,omitempty"`
	// Percentage of allocated credits that will be used in the contract period.
	ForecastedUsagePercentage *float64 `json:"forecastedUsagePercentage,omitempty"`
	// Days remaining till all the credits are consumed.
	RemainingDays *float64 `json:"remainingDays,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsageForecastResponse UsageForecastResponse

// NewUsageForecastResponse instantiates a new UsageForecastResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageForecastResponse() *UsageForecastResponse {
	this := UsageForecastResponse{}
	return &this
}

// NewUsageForecastResponseWithDefaults instantiates a new UsageForecastResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageForecastResponseWithDefaults() *UsageForecastResponse {
	this := UsageForecastResponse{}
	return &this
}

// GetAverageUsage returns the AverageUsage field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetAverageUsage() float64 {
	if o == nil || IsNil(o.AverageUsage) {
		var ret float64
		return ret
	}
	return *o.AverageUsage
}

// GetAverageUsageOk returns a tuple with the AverageUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetAverageUsageOk() (*float64, bool) {
	if o == nil || IsNil(o.AverageUsage) {
		return nil, false
	}
	return o.AverageUsage, true
}

// HasAverageUsage returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasAverageUsage() bool {
	if o != nil && !IsNil(o.AverageUsage) {
		return true
	}

	return false
}

// SetAverageUsage gets a reference to the given float64 and assigns it to the AverageUsage field.
func (o *UsageForecastResponse) SetAverageUsage(v float64) {
	o.AverageUsage = &v
}

// GetUsagePercentage returns the UsagePercentage field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetUsagePercentage() float64 {
	if o == nil || IsNil(o.UsagePercentage) {
		var ret float64
		return ret
	}
	return *o.UsagePercentage
}

// GetUsagePercentageOk returns a tuple with the UsagePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetUsagePercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.UsagePercentage) {
		return nil, false
	}
	return o.UsagePercentage, true
}

// HasUsagePercentage returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasUsagePercentage() bool {
	if o != nil && !IsNil(o.UsagePercentage) {
		return true
	}

	return false
}

// SetUsagePercentage gets a reference to the given float64 and assigns it to the UsagePercentage field.
func (o *UsageForecastResponse) SetUsagePercentage(v float64) {
	o.UsagePercentage = &v
}

// GetForecastedUsage returns the ForecastedUsage field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetForecastedUsage() float64 {
	if o == nil || IsNil(o.ForecastedUsage) {
		var ret float64
		return ret
	}
	return *o.ForecastedUsage
}

// GetForecastedUsageOk returns a tuple with the ForecastedUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetForecastedUsageOk() (*float64, bool) {
	if o == nil || IsNil(o.ForecastedUsage) {
		return nil, false
	}
	return o.ForecastedUsage, true
}

// HasForecastedUsage returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasForecastedUsage() bool {
	if o != nil && !IsNil(o.ForecastedUsage) {
		return true
	}

	return false
}

// SetForecastedUsage gets a reference to the given float64 and assigns it to the ForecastedUsage field.
func (o *UsageForecastResponse) SetForecastedUsage(v float64) {
	o.ForecastedUsage = &v
}

// GetForecastedUsagePercentage returns the ForecastedUsagePercentage field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetForecastedUsagePercentage() float64 {
	if o == nil || IsNil(o.ForecastedUsagePercentage) {
		var ret float64
		return ret
	}
	return *o.ForecastedUsagePercentage
}

// GetForecastedUsagePercentageOk returns a tuple with the ForecastedUsagePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetForecastedUsagePercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.ForecastedUsagePercentage) {
		return nil, false
	}
	return o.ForecastedUsagePercentage, true
}

// HasForecastedUsagePercentage returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasForecastedUsagePercentage() bool {
	if o != nil && !IsNil(o.ForecastedUsagePercentage) {
		return true
	}

	return false
}

// SetForecastedUsagePercentage gets a reference to the given float64 and assigns it to the ForecastedUsagePercentage field.
func (o *UsageForecastResponse) SetForecastedUsagePercentage(v float64) {
	o.ForecastedUsagePercentage = &v
}

// GetRemainingDays returns the RemainingDays field value if set, zero value otherwise.
func (o *UsageForecastResponse) GetRemainingDays() float64 {
	if o == nil || IsNil(o.RemainingDays) {
		var ret float64
		return ret
	}
	return *o.RemainingDays
}

// GetRemainingDaysOk returns a tuple with the RemainingDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageForecastResponse) GetRemainingDaysOk() (*float64, bool) {
	if o == nil || IsNil(o.RemainingDays) {
		return nil, false
	}
	return o.RemainingDays, true
}

// HasRemainingDays returns a boolean if a field has been set.
func (o *UsageForecastResponse) HasRemainingDays() bool {
	if o != nil && !IsNil(o.RemainingDays) {
		return true
	}

	return false
}

// SetRemainingDays gets a reference to the given float64 and assigns it to the RemainingDays field.
func (o *UsageForecastResponse) SetRemainingDays(v float64) {
	o.RemainingDays = &v
}

func (o UsageForecastResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageForecastResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AverageUsage) {
		toSerialize["averageUsage"] = o.AverageUsage
	}
	if !IsNil(o.UsagePercentage) {
		toSerialize["usagePercentage"] = o.UsagePercentage
	}
	if !IsNil(o.ForecastedUsage) {
		toSerialize["forecastedUsage"] = o.ForecastedUsage
	}
	if !IsNil(o.ForecastedUsagePercentage) {
		toSerialize["forecastedUsagePercentage"] = o.ForecastedUsagePercentage
	}
	if !IsNil(o.RemainingDays) {
		toSerialize["remainingDays"] = o.RemainingDays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsageForecastResponse) UnmarshalJSON(data []byte) (err error) {
	varUsageForecastResponse := _UsageForecastResponse{}

	err = json.Unmarshal(data, &varUsageForecastResponse)

	if err != nil {
		return err
	}

	*o = UsageForecastResponse(varUsageForecastResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "averageUsage")
		delete(additionalProperties, "usagePercentage")
		delete(additionalProperties, "forecastedUsage")
		delete(additionalProperties, "forecastedUsagePercentage")
		delete(additionalProperties, "remainingDays")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsageForecastResponse struct {
	value *UsageForecastResponse
	isSet bool
}

func (v NullableUsageForecastResponse) Get() *UsageForecastResponse {
	return v.value
}

func (v *NullableUsageForecastResponse) Set(val *UsageForecastResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageForecastResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageForecastResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageForecastResponse(val *UsageForecastResponse) *NullableUsageForecastResponse {
	return &NullableUsageForecastResponse{value: val, isSet: true}
}

func (v NullableUsageForecastResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageForecastResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


