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
)


// LogSearchesEstimatedUsageAPIService LogSearchesEstimatedUsageAPI service
type LogSearchesEstimatedUsageAPIService service

type ApiGetLogSearchEstimatedUsageRequest struct {
	ctx context.Context
	ApiService *LogSearchesEstimatedUsageAPIService
	logSearchEstimatedUsageRequest *LogSearchEstimatedUsageRequest
}

// The definition of the log search estimated usage.
func (r ApiGetLogSearchEstimatedUsageRequest) LogSearchEstimatedUsageRequest(logSearchEstimatedUsageRequest LogSearchEstimatedUsageRequest) ApiGetLogSearchEstimatedUsageRequest {
	r.logSearchEstimatedUsageRequest = &logSearchEstimatedUsageRequest
	return r
}

func (r ApiGetLogSearchEstimatedUsageRequest) Execute() (*LogSearchEstimatedUsageDefinition, *http.Response, error) {
	return r.ApiService.GetLogSearchEstimatedUsageExecute(r)
}

/*
GetLogSearchEstimatedUsage Gets estimated usage details.

Gets the estimated volume of data that would be scanned for a given log search in the Infrequent data tier.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLogSearchEstimatedUsageRequest
*/
func (a *LogSearchesEstimatedUsageAPIService) GetLogSearchEstimatedUsage(ctx context.Context) ApiGetLogSearchEstimatedUsageRequest {
	return ApiGetLogSearchEstimatedUsageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LogSearchEstimatedUsageDefinition
func (a *LogSearchesEstimatedUsageAPIService) GetLogSearchEstimatedUsageExecute(r ApiGetLogSearchEstimatedUsageRequest) (*LogSearchEstimatedUsageDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogSearchEstimatedUsageDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogSearchesEstimatedUsageAPIService.GetLogSearchEstimatedUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/logSearches/estimatedUsage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.logSearchEstimatedUsageRequest == nil {
		return localVarReturnValue, nil, reportError("logSearchEstimatedUsageRequest is required and must be specified")
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
	localVarPostBody = r.logSearchEstimatedUsageRequest
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

type ApiGetLogSearchEstimatedUsageByTierRequest struct {
	ctx context.Context
	ApiService *LogSearchesEstimatedUsageAPIService
	logSearchEstimatedUsageRequestV2 *LogSearchEstimatedUsageRequestV2
}

// The definition of the log search estimated usage.
func (r ApiGetLogSearchEstimatedUsageByTierRequest) LogSearchEstimatedUsageRequestV2(logSearchEstimatedUsageRequestV2 LogSearchEstimatedUsageRequestV2) ApiGetLogSearchEstimatedUsageByTierRequest {
	r.logSearchEstimatedUsageRequestV2 = &logSearchEstimatedUsageRequestV2
	return r
}

func (r ApiGetLogSearchEstimatedUsageByTierRequest) Execute() (*LogSearchEstimatedUsageByTierDefinition, *http.Response, error) {
	return r.ApiService.GetLogSearchEstimatedUsageByTierExecute(r)
}

/*
GetLogSearchEstimatedUsageByTier Gets Tier Wise estimated usage details.

Gets the estimated volume of data that would be scanned for a given log search per data tier.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLogSearchEstimatedUsageByTierRequest
*/
func (a *LogSearchesEstimatedUsageAPIService) GetLogSearchEstimatedUsageByTier(ctx context.Context) ApiGetLogSearchEstimatedUsageByTierRequest {
	return ApiGetLogSearchEstimatedUsageByTierRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LogSearchEstimatedUsageByTierDefinition
func (a *LogSearchesEstimatedUsageAPIService) GetLogSearchEstimatedUsageByTierExecute(r ApiGetLogSearchEstimatedUsageByTierRequest) (*LogSearchEstimatedUsageByTierDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogSearchEstimatedUsageByTierDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogSearchesEstimatedUsageAPIService.GetLogSearchEstimatedUsageByTier")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/logSearches/estimatedUsageByTier"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.logSearchEstimatedUsageRequestV2 == nil {
		return localVarReturnValue, nil, reportError("logSearchEstimatedUsageRequestV2 is required and must be specified")
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
	localVarPostBody = r.logSearchEstimatedUsageRequestV2
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
