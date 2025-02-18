/*
Sumo Logic API

Go client for Sumo Logic API. 

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


// ArchiveManagementAPIService ArchiveManagementAPI service
type ArchiveManagementAPIService service

type ApiCreateArchiveJobRequest struct {
	ctx context.Context
	ApiService *ArchiveManagementAPIService
	sourceId string
	createArchiveJobRequest *CreateArchiveJobRequest
}

// The definition of the ingestion job to create.
func (r ApiCreateArchiveJobRequest) CreateArchiveJobRequest(createArchiveJobRequest CreateArchiveJobRequest) ApiCreateArchiveJobRequest {
	r.createArchiveJobRequest = &createArchiveJobRequest
	return r
}

func (r ApiCreateArchiveJobRequest) Execute() (*ArchiveJob, *http.Response, error) {
	return r.ApiService.CreateArchiveJobExecute(r)
}

/*
CreateArchiveJob Create an ingestion job.

Create an ingestion job to pull data from your S3 bucket.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sourceId The identifier of the Archive Source for which the job is to be added.
 @return ApiCreateArchiveJobRequest
*/
func (a *ArchiveManagementAPIService) CreateArchiveJob(ctx context.Context, sourceId string) ApiCreateArchiveJobRequest {
	return ApiCreateArchiveJobRequest{
		ApiService: a,
		ctx: ctx,
		sourceId: sourceId,
	}
}

// Execute executes the request
//  @return ArchiveJob
func (a *ArchiveManagementAPIService) CreateArchiveJobExecute(r ApiCreateArchiveJobRequest) (*ArchiveJob, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArchiveJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveManagementAPIService.CreateArchiveJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/archive/{sourceId}/jobs"
	localVarPath = strings.Replace(localVarPath, "{"+"sourceId"+"}", url.PathEscape(parameterValueToString(r.sourceId, "sourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createArchiveJobRequest == nil {
		return localVarReturnValue, nil, reportError("createArchiveJobRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.createArchiveJobRequest
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

type ApiDeleteArchiveJobRequest struct {
	ctx context.Context
	ApiService *ArchiveManagementAPIService
	sourceId string
	id string
}

func (r ApiDeleteArchiveJobRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteArchiveJobExecute(r)
}

/*
DeleteArchiveJob Delete an ingestion job.

Delete an ingestion job with the given identifier from the organization. The delete operation is only possible for jobs with a Succeeded or Failed status.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sourceId The identifier of the Archive Source.
 @param id The identifier of the ingestion job to delete.
 @return ApiDeleteArchiveJobRequest
*/
func (a *ArchiveManagementAPIService) DeleteArchiveJob(ctx context.Context, sourceId string, id string) ApiDeleteArchiveJobRequest {
	return ApiDeleteArchiveJobRequest{
		ApiService: a,
		ctx: ctx,
		sourceId: sourceId,
		id: id,
	}
}

// Execute executes the request
func (a *ArchiveManagementAPIService) DeleteArchiveJobExecute(r ApiDeleteArchiveJobRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveManagementAPIService.DeleteArchiveJob")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/archive/{sourceId}/jobs/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"sourceId"+"}", url.PathEscape(parameterValueToString(r.sourceId, "sourceId")), -1)
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

type ApiListArchiveJobsBySourceIdRequest struct {
	ctx context.Context
	ApiService *ArchiveManagementAPIService
	sourceId string
	limit *int32
	token *string
}

// Limit the number of jobs returned in the response. The number of jobs returned may be less than the &#x60;limit&#x60;.
func (r ApiListArchiveJobsBySourceIdRequest) Limit(limit int32) ApiListArchiveJobsBySourceIdRequest {
	r.limit = &limit
	return r
}

// Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left.
func (r ApiListArchiveJobsBySourceIdRequest) Token(token string) ApiListArchiveJobsBySourceIdRequest {
	r.token = &token
	return r
}

func (r ApiListArchiveJobsBySourceIdRequest) Execute() (*ListArchiveJobsResponse, *http.Response, error) {
	return r.ApiService.ListArchiveJobsBySourceIdExecute(r)
}

/*
ListArchiveJobsBySourceId Get ingestion jobs for an Archive Source.

Get a list of all the ingestion jobs created on an Archive Source. The response is paginated with a default limit of 10 jobs per page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sourceId The identifier of an Archive Source.
 @return ApiListArchiveJobsBySourceIdRequest
*/
func (a *ArchiveManagementAPIService) ListArchiveJobsBySourceId(ctx context.Context, sourceId string) ApiListArchiveJobsBySourceIdRequest {
	return ApiListArchiveJobsBySourceIdRequest{
		ApiService: a,
		ctx: ctx,
		sourceId: sourceId,
	}
}

// Execute executes the request
//  @return ListArchiveJobsResponse
func (a *ArchiveManagementAPIService) ListArchiveJobsBySourceIdExecute(r ApiListArchiveJobsBySourceIdRequest) (*ListArchiveJobsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListArchiveJobsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveManagementAPIService.ListArchiveJobsBySourceId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/archive/{sourceId}/jobs"
	localVarPath = strings.Replace(localVarPath, "{"+"sourceId"+"}", url.PathEscape(parameterValueToString(r.sourceId, "sourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "")
	}
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

type ApiListArchiveJobsCountPerSourceRequest struct {
	ctx context.Context
	ApiService *ArchiveManagementAPIService
}

func (r ApiListArchiveJobsCountPerSourceRequest) Execute() (*ListArchiveJobsCount, *http.Response, error) {
	return r.ApiService.ListArchiveJobsCountPerSourceExecute(r)
}

/*
ListArchiveJobsCountPerSource List ingestion jobs for all Archive Sources.

Get a list of all Archive Sources with the count and status of ingestion jobs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListArchiveJobsCountPerSourceRequest
*/
func (a *ArchiveManagementAPIService) ListArchiveJobsCountPerSource(ctx context.Context) ApiListArchiveJobsCountPerSourceRequest {
	return ApiListArchiveJobsCountPerSourceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListArchiveJobsCount
func (a *ArchiveManagementAPIService) ListArchiveJobsCountPerSourceExecute(r ApiListArchiveJobsCountPerSourceRequest) (*ListArchiveJobsCount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListArchiveJobsCount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchiveManagementAPIService.ListArchiveJobsCountPerSource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/archive/jobs/count"

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
