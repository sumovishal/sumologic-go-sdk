/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on Collector and Search Job APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> KR </td>       <td> https://api.kr.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ContentManagementAPIService ContentManagementAPI service
type ContentManagementAPIService service

type ApiAsyncCopyStatusRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	id string
	jobId string
	isAdminMode *string
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiAsyncCopyStatusRequest) IsAdminMode(isAdminMode string) ApiAsyncCopyStatusRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiAsyncCopyStatusRequest) Execute() (*AsyncJobStatus, *http.Response, error) {
	return r.ApiService.AsyncCopyStatusExecute(r)
}

/*
AsyncCopyStatus Content copy job status.

Get the status of the copy request with the given job identifier. On success, field `statusMessage` will contain identifier of the newly copied content in format: `id: {hexIdentifier}`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The identifier of the content which was copied.
 @param jobId The identifier of the asynchronous copy request job.
 @return ApiAsyncCopyStatusRequest
*/
func (a *ContentManagementAPIService) AsyncCopyStatus(ctx context.Context, id string, jobId string) ApiAsyncCopyStatusRequest {
	return ApiAsyncCopyStatusRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return AsyncJobStatus
func (a *ContentManagementAPIService) AsyncCopyStatusExecute(r ApiAsyncCopyStatusRequest) (*AsyncJobStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncJobStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.AsyncCopyStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{id}/copy/{jobId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBeginAsyncCopyRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	id string
	destinationFolder *string
	isAdminMode *string
}

// The identifier of the destination folder.
func (r ApiBeginAsyncCopyRequest) DestinationFolder(destinationFolder string) ApiBeginAsyncCopyRequest {
	r.destinationFolder = &destinationFolder
	return r
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiBeginAsyncCopyRequest) IsAdminMode(isAdminMode string) ApiBeginAsyncCopyRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiBeginAsyncCopyRequest) Execute() (*BeginAsyncJobResponse, *http.Response, error) {
	return r.ApiService.BeginAsyncCopyExecute(r)
}

/*
BeginAsyncCopy Start a content copy job.

Start an asynchronous content copy job with the given identifier to the destination folder. If the content item is a folder, everything under the folder is copied recursively.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The identifier of the content item to copy. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format.
 @return ApiBeginAsyncCopyRequest
*/
func (a *ContentManagementAPIService) BeginAsyncCopy(ctx context.Context, id string) ApiBeginAsyncCopyRequest {
	return ApiBeginAsyncCopyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BeginAsyncJobResponse
func (a *ContentManagementAPIService) BeginAsyncCopyExecute(r ApiBeginAsyncCopyRequest) (*BeginAsyncJobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BeginAsyncJobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.BeginAsyncCopy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{id}/copy"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.destinationFolder == nil {
		return localVarReturnValue, nil, reportError("destinationFolder is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "destinationFolder", r.destinationFolder, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBeginAsyncDeleteRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	id string
	isAdminMode *string
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiBeginAsyncDeleteRequest) IsAdminMode(isAdminMode string) ApiBeginAsyncDeleteRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiBeginAsyncDeleteRequest) Execute() (*BeginAsyncJobResponse, *http.Response, error) {
	return r.ApiService.BeginAsyncDeleteExecute(r)
}

/*
BeginAsyncDelete Start a content deletion job.

Start an asynchronous content deletion job with the given identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the content to delete. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format.
 @return ApiBeginAsyncDeleteRequest
*/
func (a *ContentManagementAPIService) BeginAsyncDelete(ctx context.Context, id string) ApiBeginAsyncDeleteRequest {
	return ApiBeginAsyncDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BeginAsyncJobResponse
func (a *ContentManagementAPIService) BeginAsyncDeleteExecute(r ApiBeginAsyncDeleteRequest) (*BeginAsyncJobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BeginAsyncJobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.BeginAsyncDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{id}/delete"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBeginAsyncExportRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	id string
	isAdminMode *string
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiBeginAsyncExportRequest) IsAdminMode(isAdminMode string) ApiBeginAsyncExportRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiBeginAsyncExportRequest) Execute() (*BeginAsyncJobResponse, *http.Response, error) {
	return r.ApiService.BeginAsyncExportExecute(r)
}

/*
BeginAsyncExport Start a content export job.

Schedule an _asynchronous_ export of content with the given identifier. You will get back an asynchronous job identifier on success. Use the [getAsyncExportStatus](#operation/getAsyncExportStatus) endpoint and the job identifier you got back in the response to track the status of an asynchronous export job.
If the content item is a folder, everything under the folder is exported recursively. Keep in mind when exporting large folders that there is a limit of 1000 content objects that can be exported at once. If you want to import more than 1000 content objects, then be sure to split the import into batches of 1000 objects or less.
The results from the export are compatible with the Library import feature in the Sumo Logic user interface as well as the API content import job.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The identifier of the content item to export. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format.
 @return ApiBeginAsyncExportRequest
*/
func (a *ContentManagementAPIService) BeginAsyncExport(ctx context.Context, id string) ApiBeginAsyncExportRequest {
	return ApiBeginAsyncExportRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BeginAsyncJobResponse
func (a *ContentManagementAPIService) BeginAsyncExportExecute(r ApiBeginAsyncExportRequest) (*BeginAsyncJobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BeginAsyncJobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.BeginAsyncExport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{id}/export"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBeginAsyncImportRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	folderId string
	contentSyncDefinition *ContentSyncDefinition
	isAdminMode *string
	overwrite *bool
}

// The content to import.
func (r ApiBeginAsyncImportRequest) ContentSyncDefinition(contentSyncDefinition ContentSyncDefinition) ApiBeginAsyncImportRequest {
	r.contentSyncDefinition = &contentSyncDefinition
	return r
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiBeginAsyncImportRequest) IsAdminMode(isAdminMode string) ApiBeginAsyncImportRequest {
	r.isAdminMode = &isAdminMode
	return r
}

// Set this to \&quot;true\&quot; to overwrite a content item if the name already exists.
func (r ApiBeginAsyncImportRequest) Overwrite(overwrite bool) ApiBeginAsyncImportRequest {
	r.overwrite = &overwrite
	return r
}

func (r ApiBeginAsyncImportRequest) Execute() (*BeginAsyncJobResponse, *http.Response, error) {
	return r.ApiService.BeginAsyncImportExecute(r)
}

/*
BeginAsyncImport Start a content import job.

Schedule an asynchronous import of content inside an existing folder with the given identifier. Import requests can be used to create or update content within a folder. Content items need to have a unique name within their folder. If there is already a content item with the same name in the folder, you can set the `overwrite` parameter to `true` to overwrite existing content items. By default, the `overwrite` parameter is set to `false`, where the import will fail if a content item with the same name already exist. Keep in mind when importing large folders that there is a limit of 1000 content objects that can be imported at once. If you want to import more than 1000 content objects, then be sure to split the import into batches of 1000 objects or less.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId The identifier of the folder to import into. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format.
 @return ApiBeginAsyncImportRequest
*/
func (a *ContentManagementAPIService) BeginAsyncImport(ctx context.Context, folderId string) ApiBeginAsyncImportRequest {
	return ApiBeginAsyncImportRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return BeginAsyncJobResponse
func (a *ContentManagementAPIService) BeginAsyncImportExecute(r ApiBeginAsyncImportRequest) (*BeginAsyncJobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BeginAsyncJobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.BeginAsyncImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/folders/{folderId}/import"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentSyncDefinition == nil {
		return localVarReturnValue, nil, reportError("contentSyncDefinition is required and must be specified")
	}

	if r.overwrite != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "overwrite", r.overwrite, "form", "")
	} else {
		var defaultValue bool = false
		r.overwrite = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	// body params
	localVarPostBody = r.contentSyncDefinition
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAsyncDeleteStatusRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	id string
	jobId string
	isAdminMode *string
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiGetAsyncDeleteStatusRequest) IsAdminMode(isAdminMode string) ApiGetAsyncDeleteStatusRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiGetAsyncDeleteStatusRequest) Execute() (*AsyncJobStatus, *http.Response, error) {
	return r.ApiService.GetAsyncDeleteStatusExecute(r)
}

/*
GetAsyncDeleteStatus Content deletion job status.

Get the status of an asynchronous content deletion job request for the given job identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the content to delete.
 @param jobId The identifier of the asynchronous job.
 @return ApiGetAsyncDeleteStatusRequest
*/
func (a *ContentManagementAPIService) GetAsyncDeleteStatus(ctx context.Context, id string, jobId string) ApiGetAsyncDeleteStatusRequest {
	return ApiGetAsyncDeleteStatusRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return AsyncJobStatus
func (a *ContentManagementAPIService) GetAsyncDeleteStatusExecute(r ApiGetAsyncDeleteStatusRequest) (*AsyncJobStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncJobStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.GetAsyncDeleteStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{id}/delete/{jobId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAsyncExportResultRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	contentId string
	jobId string
	isAdminMode *string
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiGetAsyncExportResultRequest) IsAdminMode(isAdminMode string) ApiGetAsyncExportResultRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiGetAsyncExportResultRequest) Execute() (*ContentSyncDefinition, *http.Response, error) {
	return r.ApiService.GetAsyncExportResultExecute(r)
}

/*
GetAsyncExportResult Content export job result.

Get results from content export job for the given job identifier. The results from this export are incompatible with the Library import feature in the Sumo user interface.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contentId The identifier of the exported content item.
 @param jobId The identifier of the asynchronous job.
 @return ApiGetAsyncExportResultRequest
*/
func (a *ContentManagementAPIService) GetAsyncExportResult(ctx context.Context, contentId string, jobId string) ApiGetAsyncExportResultRequest {
	return ApiGetAsyncExportResultRequest{
		ApiService: a,
		ctx: ctx,
		contentId: contentId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return ContentSyncDefinition
func (a *ContentManagementAPIService) GetAsyncExportResultExecute(r ApiGetAsyncExportResultRequest) (*ContentSyncDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContentSyncDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.GetAsyncExportResult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{contentId}/export/{jobId}/result"
	localVarPath = strings.Replace(localVarPath, "{"+"contentId"+"}", url.PathEscape(parameterValueToString(r.contentId, "contentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAsyncExportStatusRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	contentId string
	jobId string
	isAdminMode *string
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiGetAsyncExportStatusRequest) IsAdminMode(isAdminMode string) ApiGetAsyncExportStatusRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiGetAsyncExportStatusRequest) Execute() (*AsyncJobStatus, *http.Response, error) {
	return r.ApiService.GetAsyncExportStatusExecute(r)
}

/*
GetAsyncExportStatus Content export job status.

Get the status of an asynchronous content export request for the given job identifier. On success, use the [getExportResult](#operation/getAsyncExportResult) endpoint to get the result of the export job.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contentId The identifier of the exported content item.
 @param jobId The identifier of the asynchronous export job.
 @return ApiGetAsyncExportStatusRequest
*/
func (a *ContentManagementAPIService) GetAsyncExportStatus(ctx context.Context, contentId string, jobId string) ApiGetAsyncExportStatusRequest {
	return ApiGetAsyncExportStatusRequest{
		ApiService: a,
		ctx: ctx,
		contentId: contentId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return AsyncJobStatus
func (a *ContentManagementAPIService) GetAsyncExportStatusExecute(r ApiGetAsyncExportStatusRequest) (*AsyncJobStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncJobStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.GetAsyncExportStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{contentId}/export/{jobId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"contentId"+"}", url.PathEscape(parameterValueToString(r.contentId, "contentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAsyncImportStatusRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	folderId string
	jobId string
	isAdminMode *string
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiGetAsyncImportStatusRequest) IsAdminMode(isAdminMode string) ApiGetAsyncImportStatusRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiGetAsyncImportStatusRequest) Execute() (*AsyncJobStatus, *http.Response, error) {
	return r.ApiService.GetAsyncImportStatusExecute(r)
}

/*
GetAsyncImportStatus Content import job status.

Get the status of a content import job for the given job identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId The identifier of the folder to import into.
 @param jobId The identifier of the import request.
 @return ApiGetAsyncImportStatusRequest
*/
func (a *ContentManagementAPIService) GetAsyncImportStatus(ctx context.Context, folderId string, jobId string) ApiGetAsyncImportStatusRequest {
	return ApiGetAsyncImportStatusRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return AsyncJobStatus
func (a *ContentManagementAPIService) GetAsyncImportStatusExecute(r ApiGetAsyncImportStatusRequest) (*AsyncJobStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncJobStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.GetAsyncImportStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/folders/{folderId}/import/{jobId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetItemByPathRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	path *string
}

// Path of the content item to retrieve.
func (r ApiGetItemByPathRequest) Path(path string) ApiGetItemByPathRequest {
	r.path = &path
	return r
}

func (r ApiGetItemByPathRequest) Execute() (*Content, *http.Response, error) {
	return r.ApiService.GetItemByPathExecute(r)
}

/*
GetItemByPath Get content item by path.

Get a content item corresponding to the given path.

_Path is specified in the required query parameter `path`. The path should be URL encoded._ For example, to get "Acme Corp" folder of a user "user@sumo.com" you can use the following curl command:
  ```bash
  curl https://api.sumologic.com/api/v2/content/path?path=/Library/Users/user%40sumo.com/Acme%20Corp
  ```


The absolute path to a content item should be specified to get the item. The content library has "Library" folder at the root level. For items in "Personal" folder, the base path is "/Library/Users/user@sumo.com" where "user@sumo.com" is the email address of the user. For example if a user with email address `wile@acme.com` has `Rockets` folder inside Personal folder, the path of Rockets folder will be `/Library/Users/wile@acme.com/Rockets`.

For items in "Admin Recommended" folder, the base path is "/Library/Admin Recommended". For example, given a folder `Acme` in Admin Recommended folder, the path will be `/Library/Admin Recommended/Acme`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetItemByPathRequest
*/
func (a *ContentManagementAPIService) GetItemByPath(ctx context.Context) ApiGetItemByPathRequest {
	return ApiGetItemByPathRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Content
func (a *ContentManagementAPIService) GetItemByPathExecute(r ApiGetItemByPathRequest) (*Content, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Content
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.GetItemByPath")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/path"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.path == nil {
		return localVarReturnValue, nil, reportError("path is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "path", r.path, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPathByIdRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	contentId string
}

func (r ApiGetPathByIdRequest) Execute() (*ContentPath, *http.Response, error) {
	return r.ApiService.GetPathByIdExecute(r)
}

/*
GetPathById Get path of an item.

Get full path of a content item with the given identifier.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contentId Identifier of the content item to get the path.
 @return ApiGetPathByIdRequest
*/
func (a *ContentManagementAPIService) GetPathById(ctx context.Context, contentId string) ApiGetPathByIdRequest {
	return ApiGetPathByIdRequest{
		ApiService: a,
		ctx: ctx,
		contentId: contentId,
	}
}

// Execute executes the request
//  @return ContentPath
func (a *ContentManagementAPIService) GetPathByIdExecute(r ApiGetPathByIdRequest) (*ContentPath, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContentPath
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.GetPathById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{contentId}/path"
	localVarPath = strings.Replace(localVarPath, "{"+"contentId"+"}", url.PathEscape(parameterValueToString(r.contentId, "contentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMoveItemRequest struct {
	ctx context.Context
	ApiService *ContentManagementAPIService
	destinationFolderId *string
	id string
	isAdminMode *string
}

// Identifier of the destination folder.
func (r ApiMoveItemRequest) DestinationFolderId(destinationFolderId string) ApiMoveItemRequest {
	r.destinationFolderId = &destinationFolderId
	return r
}

// Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator.
func (r ApiMoveItemRequest) IsAdminMode(isAdminMode string) ApiMoveItemRequest {
	r.isAdminMode = &isAdminMode
	return r
}

func (r ApiMoveItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.MoveItemExecute(r)
}

/*
MoveItem Move an item.

Moves an item from its current location to another folder.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the item the user wants to move.
 @return ApiMoveItemRequest
*/
func (a *ContentManagementAPIService) MoveItem(ctx context.Context, id string) ApiMoveItemRequest {
	return ApiMoveItemRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ContentManagementAPIService) MoveItemExecute(r ApiMoveItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentManagementAPIService.MoveItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/content/{id}/move"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.destinationFolderId == nil {
		return nil, reportError("destinationFolderId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "destinationFolderId", r.destinationFolderId, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.isAdminMode != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "isAdminMode", r.isAdminMode, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
