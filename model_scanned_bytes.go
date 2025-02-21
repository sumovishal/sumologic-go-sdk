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

// checks if the ScannedBytes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannedBytes{}

// ScannedBytes The total number of scanned bytes from tiered data sources (ex. infrequent, continuous, frequent). See https://help.sumologic.com/docs/manage/partitions-data-tiers/data-tiers/ for a more detailed explaination. 
type ScannedBytes struct {
	// The total number of scanned bytes from infrequent tier data for the query in bytes.
	Infrequent *int64 `json:"infrequent,omitempty"`
	// The total number of scanned bytes from continuous tier data for the query in bytes.
	Continuous *int64 `json:"continuous,omitempty"`
	// The total number of scanned bytes from frequent tier data for the query in bytes.
	Frequent *int64 `json:"frequent,omitempty"`
	// The total number of scanned bytes from security tier data for the query in bytes.
	Security *int64 `json:"security,omitempty"`
	// The total number of scanned bytes from tracing tier data for the query in bytes.
	Tracing *int64 `json:"tracing,omitempty"`
	// The total number of scanned bytes from upfront tier data for the query in bytes.
	Upfront *int64 `json:"upfront,omitempty"`
	// The total number of scanned bytes from metered tier data for the query in bytes.
	Metered *int64 `json:"metered,omitempty"`
	// The total number of scanned bytes from rce tier data for the query in bytes.
	Rce *int64 `json:"rce,omitempty"`
	// The total number of scanned bytes from flex tier data for the query in bytes.
	Flex *int64 `json:"flex,omitempty"`
	// The total number of scanned bytes from continuous security tier data for the query in bytes.
	ContinuousSecurity *int64 `json:"continuousSecurity,omitempty"`
	// The total number of scanned bytes from flex security tier data for the query in bytes.
	FlexSecurity *int64 `json:"flexSecurity,omitempty"`
	// The total number of scanned bytes from flex upfront tier data for the query in bytes.
	FlexUpfront *int64 `json:"flexUpfront,omitempty"`
	// The total number of scanned bytes from flex metered tier data for the query in bytes.
	FlexMetered *int64 `json:"flexMetered,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScannedBytes ScannedBytes

// NewScannedBytes instantiates a new ScannedBytes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannedBytes() *ScannedBytes {
	this := ScannedBytes{}
	return &this
}

// NewScannedBytesWithDefaults instantiates a new ScannedBytes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannedBytesWithDefaults() *ScannedBytes {
	this := ScannedBytes{}
	return &this
}

// GetInfrequent returns the Infrequent field value if set, zero value otherwise.
func (o *ScannedBytes) GetInfrequent() int64 {
	if o == nil || IsNil(o.Infrequent) {
		var ret int64
		return ret
	}
	return *o.Infrequent
}

// GetInfrequentOk returns a tuple with the Infrequent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetInfrequentOk() (*int64, bool) {
	if o == nil || IsNil(o.Infrequent) {
		return nil, false
	}
	return o.Infrequent, true
}

// HasInfrequent returns a boolean if a field has been set.
func (o *ScannedBytes) HasInfrequent() bool {
	if o != nil && !IsNil(o.Infrequent) {
		return true
	}

	return false
}

// SetInfrequent gets a reference to the given int64 and assigns it to the Infrequent field.
func (o *ScannedBytes) SetInfrequent(v int64) {
	o.Infrequent = &v
}

// GetContinuous returns the Continuous field value if set, zero value otherwise.
func (o *ScannedBytes) GetContinuous() int64 {
	if o == nil || IsNil(o.Continuous) {
		var ret int64
		return ret
	}
	return *o.Continuous
}

// GetContinuousOk returns a tuple with the Continuous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetContinuousOk() (*int64, bool) {
	if o == nil || IsNil(o.Continuous) {
		return nil, false
	}
	return o.Continuous, true
}

// HasContinuous returns a boolean if a field has been set.
func (o *ScannedBytes) HasContinuous() bool {
	if o != nil && !IsNil(o.Continuous) {
		return true
	}

	return false
}

// SetContinuous gets a reference to the given int64 and assigns it to the Continuous field.
func (o *ScannedBytes) SetContinuous(v int64) {
	o.Continuous = &v
}

// GetFrequent returns the Frequent field value if set, zero value otherwise.
func (o *ScannedBytes) GetFrequent() int64 {
	if o == nil || IsNil(o.Frequent) {
		var ret int64
		return ret
	}
	return *o.Frequent
}

// GetFrequentOk returns a tuple with the Frequent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetFrequentOk() (*int64, bool) {
	if o == nil || IsNil(o.Frequent) {
		return nil, false
	}
	return o.Frequent, true
}

// HasFrequent returns a boolean if a field has been set.
func (o *ScannedBytes) HasFrequent() bool {
	if o != nil && !IsNil(o.Frequent) {
		return true
	}

	return false
}

// SetFrequent gets a reference to the given int64 and assigns it to the Frequent field.
func (o *ScannedBytes) SetFrequent(v int64) {
	o.Frequent = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *ScannedBytes) GetSecurity() int64 {
	if o == nil || IsNil(o.Security) {
		var ret int64
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetSecurityOk() (*int64, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *ScannedBytes) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given int64 and assigns it to the Security field.
func (o *ScannedBytes) SetSecurity(v int64) {
	o.Security = &v
}

// GetTracing returns the Tracing field value if set, zero value otherwise.
func (o *ScannedBytes) GetTracing() int64 {
	if o == nil || IsNil(o.Tracing) {
		var ret int64
		return ret
	}
	return *o.Tracing
}

// GetTracingOk returns a tuple with the Tracing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetTracingOk() (*int64, bool) {
	if o == nil || IsNil(o.Tracing) {
		return nil, false
	}
	return o.Tracing, true
}

// HasTracing returns a boolean if a field has been set.
func (o *ScannedBytes) HasTracing() bool {
	if o != nil && !IsNil(o.Tracing) {
		return true
	}

	return false
}

// SetTracing gets a reference to the given int64 and assigns it to the Tracing field.
func (o *ScannedBytes) SetTracing(v int64) {
	o.Tracing = &v
}

// GetUpfront returns the Upfront field value if set, zero value otherwise.
func (o *ScannedBytes) GetUpfront() int64 {
	if o == nil || IsNil(o.Upfront) {
		var ret int64
		return ret
	}
	return *o.Upfront
}

// GetUpfrontOk returns a tuple with the Upfront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetUpfrontOk() (*int64, bool) {
	if o == nil || IsNil(o.Upfront) {
		return nil, false
	}
	return o.Upfront, true
}

// HasUpfront returns a boolean if a field has been set.
func (o *ScannedBytes) HasUpfront() bool {
	if o != nil && !IsNil(o.Upfront) {
		return true
	}

	return false
}

// SetUpfront gets a reference to the given int64 and assigns it to the Upfront field.
func (o *ScannedBytes) SetUpfront(v int64) {
	o.Upfront = &v
}

// GetMetered returns the Metered field value if set, zero value otherwise.
func (o *ScannedBytes) GetMetered() int64 {
	if o == nil || IsNil(o.Metered) {
		var ret int64
		return ret
	}
	return *o.Metered
}

// GetMeteredOk returns a tuple with the Metered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetMeteredOk() (*int64, bool) {
	if o == nil || IsNil(o.Metered) {
		return nil, false
	}
	return o.Metered, true
}

// HasMetered returns a boolean if a field has been set.
func (o *ScannedBytes) HasMetered() bool {
	if o != nil && !IsNil(o.Metered) {
		return true
	}

	return false
}

// SetMetered gets a reference to the given int64 and assigns it to the Metered field.
func (o *ScannedBytes) SetMetered(v int64) {
	o.Metered = &v
}

// GetRce returns the Rce field value if set, zero value otherwise.
func (o *ScannedBytes) GetRce() int64 {
	if o == nil || IsNil(o.Rce) {
		var ret int64
		return ret
	}
	return *o.Rce
}

// GetRceOk returns a tuple with the Rce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetRceOk() (*int64, bool) {
	if o == nil || IsNil(o.Rce) {
		return nil, false
	}
	return o.Rce, true
}

// HasRce returns a boolean if a field has been set.
func (o *ScannedBytes) HasRce() bool {
	if o != nil && !IsNil(o.Rce) {
		return true
	}

	return false
}

// SetRce gets a reference to the given int64 and assigns it to the Rce field.
func (o *ScannedBytes) SetRce(v int64) {
	o.Rce = &v
}

// GetFlex returns the Flex field value if set, zero value otherwise.
func (o *ScannedBytes) GetFlex() int64 {
	if o == nil || IsNil(o.Flex) {
		var ret int64
		return ret
	}
	return *o.Flex
}

// GetFlexOk returns a tuple with the Flex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetFlexOk() (*int64, bool) {
	if o == nil || IsNil(o.Flex) {
		return nil, false
	}
	return o.Flex, true
}

// HasFlex returns a boolean if a field has been set.
func (o *ScannedBytes) HasFlex() bool {
	if o != nil && !IsNil(o.Flex) {
		return true
	}

	return false
}

// SetFlex gets a reference to the given int64 and assigns it to the Flex field.
func (o *ScannedBytes) SetFlex(v int64) {
	o.Flex = &v
}

// GetContinuousSecurity returns the ContinuousSecurity field value if set, zero value otherwise.
func (o *ScannedBytes) GetContinuousSecurity() int64 {
	if o == nil || IsNil(o.ContinuousSecurity) {
		var ret int64
		return ret
	}
	return *o.ContinuousSecurity
}

// GetContinuousSecurityOk returns a tuple with the ContinuousSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetContinuousSecurityOk() (*int64, bool) {
	if o == nil || IsNil(o.ContinuousSecurity) {
		return nil, false
	}
	return o.ContinuousSecurity, true
}

// HasContinuousSecurity returns a boolean if a field has been set.
func (o *ScannedBytes) HasContinuousSecurity() bool {
	if o != nil && !IsNil(o.ContinuousSecurity) {
		return true
	}

	return false
}

// SetContinuousSecurity gets a reference to the given int64 and assigns it to the ContinuousSecurity field.
func (o *ScannedBytes) SetContinuousSecurity(v int64) {
	o.ContinuousSecurity = &v
}

// GetFlexSecurity returns the FlexSecurity field value if set, zero value otherwise.
func (o *ScannedBytes) GetFlexSecurity() int64 {
	if o == nil || IsNil(o.FlexSecurity) {
		var ret int64
		return ret
	}
	return *o.FlexSecurity
}

// GetFlexSecurityOk returns a tuple with the FlexSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetFlexSecurityOk() (*int64, bool) {
	if o == nil || IsNil(o.FlexSecurity) {
		return nil, false
	}
	return o.FlexSecurity, true
}

// HasFlexSecurity returns a boolean if a field has been set.
func (o *ScannedBytes) HasFlexSecurity() bool {
	if o != nil && !IsNil(o.FlexSecurity) {
		return true
	}

	return false
}

// SetFlexSecurity gets a reference to the given int64 and assigns it to the FlexSecurity field.
func (o *ScannedBytes) SetFlexSecurity(v int64) {
	o.FlexSecurity = &v
}

// GetFlexUpfront returns the FlexUpfront field value if set, zero value otherwise.
func (o *ScannedBytes) GetFlexUpfront() int64 {
	if o == nil || IsNil(o.FlexUpfront) {
		var ret int64
		return ret
	}
	return *o.FlexUpfront
}

// GetFlexUpfrontOk returns a tuple with the FlexUpfront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetFlexUpfrontOk() (*int64, bool) {
	if o == nil || IsNil(o.FlexUpfront) {
		return nil, false
	}
	return o.FlexUpfront, true
}

// HasFlexUpfront returns a boolean if a field has been set.
func (o *ScannedBytes) HasFlexUpfront() bool {
	if o != nil && !IsNil(o.FlexUpfront) {
		return true
	}

	return false
}

// SetFlexUpfront gets a reference to the given int64 and assigns it to the FlexUpfront field.
func (o *ScannedBytes) SetFlexUpfront(v int64) {
	o.FlexUpfront = &v
}

// GetFlexMetered returns the FlexMetered field value if set, zero value otherwise.
func (o *ScannedBytes) GetFlexMetered() int64 {
	if o == nil || IsNil(o.FlexMetered) {
		var ret int64
		return ret
	}
	return *o.FlexMetered
}

// GetFlexMeteredOk returns a tuple with the FlexMetered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedBytes) GetFlexMeteredOk() (*int64, bool) {
	if o == nil || IsNil(o.FlexMetered) {
		return nil, false
	}
	return o.FlexMetered, true
}

// HasFlexMetered returns a boolean if a field has been set.
func (o *ScannedBytes) HasFlexMetered() bool {
	if o != nil && !IsNil(o.FlexMetered) {
		return true
	}

	return false
}

// SetFlexMetered gets a reference to the given int64 and assigns it to the FlexMetered field.
func (o *ScannedBytes) SetFlexMetered(v int64) {
	o.FlexMetered = &v
}

func (o ScannedBytes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannedBytes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Infrequent) {
		toSerialize["infrequent"] = o.Infrequent
	}
	if !IsNil(o.Continuous) {
		toSerialize["continuous"] = o.Continuous
	}
	if !IsNil(o.Frequent) {
		toSerialize["frequent"] = o.Frequent
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.Tracing) {
		toSerialize["tracing"] = o.Tracing
	}
	if !IsNil(o.Upfront) {
		toSerialize["upfront"] = o.Upfront
	}
	if !IsNil(o.Metered) {
		toSerialize["metered"] = o.Metered
	}
	if !IsNil(o.Rce) {
		toSerialize["rce"] = o.Rce
	}
	if !IsNil(o.Flex) {
		toSerialize["flex"] = o.Flex
	}
	if !IsNil(o.ContinuousSecurity) {
		toSerialize["continuousSecurity"] = o.ContinuousSecurity
	}
	if !IsNil(o.FlexSecurity) {
		toSerialize["flexSecurity"] = o.FlexSecurity
	}
	if !IsNil(o.FlexUpfront) {
		toSerialize["flexUpfront"] = o.FlexUpfront
	}
	if !IsNil(o.FlexMetered) {
		toSerialize["flexMetered"] = o.FlexMetered
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScannedBytes) UnmarshalJSON(data []byte) (err error) {
	varScannedBytes := _ScannedBytes{}

	err = json.Unmarshal(data, &varScannedBytes)

	if err != nil {
		return err
	}

	*o = ScannedBytes(varScannedBytes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "infrequent")
		delete(additionalProperties, "continuous")
		delete(additionalProperties, "frequent")
		delete(additionalProperties, "security")
		delete(additionalProperties, "tracing")
		delete(additionalProperties, "upfront")
		delete(additionalProperties, "metered")
		delete(additionalProperties, "rce")
		delete(additionalProperties, "flex")
		delete(additionalProperties, "continuousSecurity")
		delete(additionalProperties, "flexSecurity")
		delete(additionalProperties, "flexUpfront")
		delete(additionalProperties, "flexMetered")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScannedBytes struct {
	value *ScannedBytes
	isSet bool
}

func (v NullableScannedBytes) Get() *ScannedBytes {
	return v.value
}

func (v *NullableScannedBytes) Set(val *ScannedBytes) {
	v.value = val
	v.isSet = true
}

func (v NullableScannedBytes) IsSet() bool {
	return v.isSet
}

func (v *NullableScannedBytes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannedBytes(val *ScannedBytes) *NullableScannedBytes {
	return &NullableScannedBytes{value: val, isSet: true}
}

func (v NullableScannedBytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannedBytes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


