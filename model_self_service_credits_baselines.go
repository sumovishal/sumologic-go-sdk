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

// checks if the SelfServiceCreditsBaselines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceCreditsBaselines{}

// SelfServiceCreditsBaselines Details of product variables and its quantity as required for credits.
type SelfServiceCreditsBaselines struct {
	// The amount of continuous logs ingest to allocate to the organization, in GBs.
	ContinuousIngest *int64 `json:"continuousIngest,omitempty"`
	// Number of days of continuous logs storage to allocate to the organization, in Days.
	ContinuousStorage *int64 `json:"continuousStorage,omitempty"`
	// The amount of Metrics usage to allocate to the organization, in DPMs (Data Points per Minute).
	Metrics *int64 `json:"metrics,omitempty"`
	// The amount of tracing data ingest to allocate to the organization, in GBs.
	TracingIngest *int64 `json:"tracingIngest,omitempty"`
	// The amount of flex logs ingest to allocate to the organization, in GBs.
	FlexIngest *int64 `json:"flexIngest,omitempty"`
	// Number of days of flex logs storage to allocate to the organization, in Days.
	FlexStorage *int64 `json:"flexStorage,omitempty"`
	// The amount of flex logs ingest scan ratio.
	FlexScanRatio *int64 `json:"flexScanRatio,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceCreditsBaselines SelfServiceCreditsBaselines

// NewSelfServiceCreditsBaselines instantiates a new SelfServiceCreditsBaselines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceCreditsBaselines() *SelfServiceCreditsBaselines {
	this := SelfServiceCreditsBaselines{}
	var continuousIngest int64 = 0
	this.ContinuousIngest = &continuousIngest
	var continuousStorage int64 = 0
	this.ContinuousStorage = &continuousStorage
	var metrics int64 = 0
	this.Metrics = &metrics
	var tracingIngest int64 = 0
	this.TracingIngest = &tracingIngest
	var flexIngest int64 = 0
	this.FlexIngest = &flexIngest
	var flexStorage int64 = 0
	this.FlexStorage = &flexStorage
	var flexScanRatio int64 = 0
	this.FlexScanRatio = &flexScanRatio
	return &this
}

// NewSelfServiceCreditsBaselinesWithDefaults instantiates a new SelfServiceCreditsBaselines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceCreditsBaselinesWithDefaults() *SelfServiceCreditsBaselines {
	this := SelfServiceCreditsBaselines{}
	var continuousIngest int64 = 0
	this.ContinuousIngest = &continuousIngest
	var continuousStorage int64 = 0
	this.ContinuousStorage = &continuousStorage
	var metrics int64 = 0
	this.Metrics = &metrics
	var tracingIngest int64 = 0
	this.TracingIngest = &tracingIngest
	var flexIngest int64 = 0
	this.FlexIngest = &flexIngest
	var flexStorage int64 = 0
	this.FlexStorage = &flexStorage
	var flexScanRatio int64 = 0
	this.FlexScanRatio = &flexScanRatio
	return &this
}

// GetContinuousIngest returns the ContinuousIngest field value if set, zero value otherwise.
func (o *SelfServiceCreditsBaselines) GetContinuousIngest() int64 {
	if o == nil || IsNil(o.ContinuousIngest) {
		var ret int64
		return ret
	}
	return *o.ContinuousIngest
}

// GetContinuousIngestOk returns a tuple with the ContinuousIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceCreditsBaselines) GetContinuousIngestOk() (*int64, bool) {
	if o == nil || IsNil(o.ContinuousIngest) {
		return nil, false
	}
	return o.ContinuousIngest, true
}

// HasContinuousIngest returns a boolean if a field has been set.
func (o *SelfServiceCreditsBaselines) HasContinuousIngest() bool {
	if o != nil && !IsNil(o.ContinuousIngest) {
		return true
	}

	return false
}

// SetContinuousIngest gets a reference to the given int64 and assigns it to the ContinuousIngest field.
func (o *SelfServiceCreditsBaselines) SetContinuousIngest(v int64) {
	o.ContinuousIngest = &v
}

// GetContinuousStorage returns the ContinuousStorage field value if set, zero value otherwise.
func (o *SelfServiceCreditsBaselines) GetContinuousStorage() int64 {
	if o == nil || IsNil(o.ContinuousStorage) {
		var ret int64
		return ret
	}
	return *o.ContinuousStorage
}

// GetContinuousStorageOk returns a tuple with the ContinuousStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceCreditsBaselines) GetContinuousStorageOk() (*int64, bool) {
	if o == nil || IsNil(o.ContinuousStorage) {
		return nil, false
	}
	return o.ContinuousStorage, true
}

// HasContinuousStorage returns a boolean if a field has been set.
func (o *SelfServiceCreditsBaselines) HasContinuousStorage() bool {
	if o != nil && !IsNil(o.ContinuousStorage) {
		return true
	}

	return false
}

// SetContinuousStorage gets a reference to the given int64 and assigns it to the ContinuousStorage field.
func (o *SelfServiceCreditsBaselines) SetContinuousStorage(v int64) {
	o.ContinuousStorage = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *SelfServiceCreditsBaselines) GetMetrics() int64 {
	if o == nil || IsNil(o.Metrics) {
		var ret int64
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceCreditsBaselines) GetMetricsOk() (*int64, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *SelfServiceCreditsBaselines) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given int64 and assigns it to the Metrics field.
func (o *SelfServiceCreditsBaselines) SetMetrics(v int64) {
	o.Metrics = &v
}

// GetTracingIngest returns the TracingIngest field value if set, zero value otherwise.
func (o *SelfServiceCreditsBaselines) GetTracingIngest() int64 {
	if o == nil || IsNil(o.TracingIngest) {
		var ret int64
		return ret
	}
	return *o.TracingIngest
}

// GetTracingIngestOk returns a tuple with the TracingIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceCreditsBaselines) GetTracingIngestOk() (*int64, bool) {
	if o == nil || IsNil(o.TracingIngest) {
		return nil, false
	}
	return o.TracingIngest, true
}

// HasTracingIngest returns a boolean if a field has been set.
func (o *SelfServiceCreditsBaselines) HasTracingIngest() bool {
	if o != nil && !IsNil(o.TracingIngest) {
		return true
	}

	return false
}

// SetTracingIngest gets a reference to the given int64 and assigns it to the TracingIngest field.
func (o *SelfServiceCreditsBaselines) SetTracingIngest(v int64) {
	o.TracingIngest = &v
}

// GetFlexIngest returns the FlexIngest field value if set, zero value otherwise.
func (o *SelfServiceCreditsBaselines) GetFlexIngest() int64 {
	if o == nil || IsNil(o.FlexIngest) {
		var ret int64
		return ret
	}
	return *o.FlexIngest
}

// GetFlexIngestOk returns a tuple with the FlexIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceCreditsBaselines) GetFlexIngestOk() (*int64, bool) {
	if o == nil || IsNil(o.FlexIngest) {
		return nil, false
	}
	return o.FlexIngest, true
}

// HasFlexIngest returns a boolean if a field has been set.
func (o *SelfServiceCreditsBaselines) HasFlexIngest() bool {
	if o != nil && !IsNil(o.FlexIngest) {
		return true
	}

	return false
}

// SetFlexIngest gets a reference to the given int64 and assigns it to the FlexIngest field.
func (o *SelfServiceCreditsBaselines) SetFlexIngest(v int64) {
	o.FlexIngest = &v
}

// GetFlexStorage returns the FlexStorage field value if set, zero value otherwise.
func (o *SelfServiceCreditsBaselines) GetFlexStorage() int64 {
	if o == nil || IsNil(o.FlexStorage) {
		var ret int64
		return ret
	}
	return *o.FlexStorage
}

// GetFlexStorageOk returns a tuple with the FlexStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceCreditsBaselines) GetFlexStorageOk() (*int64, bool) {
	if o == nil || IsNil(o.FlexStorage) {
		return nil, false
	}
	return o.FlexStorage, true
}

// HasFlexStorage returns a boolean if a field has been set.
func (o *SelfServiceCreditsBaselines) HasFlexStorage() bool {
	if o != nil && !IsNil(o.FlexStorage) {
		return true
	}

	return false
}

// SetFlexStorage gets a reference to the given int64 and assigns it to the FlexStorage field.
func (o *SelfServiceCreditsBaselines) SetFlexStorage(v int64) {
	o.FlexStorage = &v
}

// GetFlexScanRatio returns the FlexScanRatio field value if set, zero value otherwise.
func (o *SelfServiceCreditsBaselines) GetFlexScanRatio() int64 {
	if o == nil || IsNil(o.FlexScanRatio) {
		var ret int64
		return ret
	}
	return *o.FlexScanRatio
}

// GetFlexScanRatioOk returns a tuple with the FlexScanRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceCreditsBaselines) GetFlexScanRatioOk() (*int64, bool) {
	if o == nil || IsNil(o.FlexScanRatio) {
		return nil, false
	}
	return o.FlexScanRatio, true
}

// HasFlexScanRatio returns a boolean if a field has been set.
func (o *SelfServiceCreditsBaselines) HasFlexScanRatio() bool {
	if o != nil && !IsNil(o.FlexScanRatio) {
		return true
	}

	return false
}

// SetFlexScanRatio gets a reference to the given int64 and assigns it to the FlexScanRatio field.
func (o *SelfServiceCreditsBaselines) SetFlexScanRatio(v int64) {
	o.FlexScanRatio = &v
}

func (o SelfServiceCreditsBaselines) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceCreditsBaselines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContinuousIngest) {
		toSerialize["continuousIngest"] = o.ContinuousIngest
	}
	if !IsNil(o.ContinuousStorage) {
		toSerialize["continuousStorage"] = o.ContinuousStorage
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.TracingIngest) {
		toSerialize["tracingIngest"] = o.TracingIngest
	}
	if !IsNil(o.FlexIngest) {
		toSerialize["flexIngest"] = o.FlexIngest
	}
	if !IsNil(o.FlexStorage) {
		toSerialize["flexStorage"] = o.FlexStorage
	}
	if !IsNil(o.FlexScanRatio) {
		toSerialize["flexScanRatio"] = o.FlexScanRatio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServiceCreditsBaselines) UnmarshalJSON(data []byte) (err error) {
	varSelfServiceCreditsBaselines := _SelfServiceCreditsBaselines{}

	err = json.Unmarshal(data, &varSelfServiceCreditsBaselines)

	if err != nil {
		return err
	}

	*o = SelfServiceCreditsBaselines(varSelfServiceCreditsBaselines)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "continuousIngest")
		delete(additionalProperties, "continuousStorage")
		delete(additionalProperties, "metrics")
		delete(additionalProperties, "tracingIngest")
		delete(additionalProperties, "flexIngest")
		delete(additionalProperties, "flexStorage")
		delete(additionalProperties, "flexScanRatio")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServiceCreditsBaselines struct {
	value *SelfServiceCreditsBaselines
	isSet bool
}

func (v NullableSelfServiceCreditsBaselines) Get() *SelfServiceCreditsBaselines {
	return v.value
}

func (v *NullableSelfServiceCreditsBaselines) Set(val *SelfServiceCreditsBaselines) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceCreditsBaselines) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceCreditsBaselines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceCreditsBaselines(val *SelfServiceCreditsBaselines) *NullableSelfServiceCreditsBaselines {
	return &NullableSelfServiceCreditsBaselines{value: val, isSet: true}
}

func (v NullableSelfServiceCreditsBaselines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceCreditsBaselines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


