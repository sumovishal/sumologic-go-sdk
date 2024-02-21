/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ThreatIntelIngestApiService ThreatIntelIngestApi service
type ThreatIntelIngestApiService service

type ApiDatastoreGetRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
}

func (r ApiDatastoreGetRequest) Execute() (*DatastoreStatusResponse, *http.Response, error) {
	return r.ApiService.DatastoreGetExecute(r)
}

/*
DatastoreGet Get threat intel indicators DB information

Get threat intel indicators DB information, such as storage utilization and indicator counts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDatastoreGetRequest
*/
func (a *ThreatIntelIngestApiService) DatastoreGet(ctx context.Context) ApiDatastoreGetRequest {
	return ApiDatastoreGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DatastoreStatusResponse
func (a *ThreatIntelIngestApiService) DatastoreGetExecute(r ApiDatastoreGetRequest) (*DatastoreStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DatastoreStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.DatastoreGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/db"

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

type ApiRemoveDatastoreRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
}

func (r ApiRemoveDatastoreRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveDatastoreExecute(r)
}

/*
RemoveDatastore Remove the threat intel indicators DB

Removes the entire database and all indicators associated with this tenant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveDatastoreRequest
*/
func (a *ThreatIntelIngestApiService) RemoveDatastore(ctx context.Context) ApiRemoveDatastoreRequest {
	return ApiRemoveDatastoreRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ThreatIntelIngestApiService) RemoveDatastoreExecute(r ApiRemoveDatastoreRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.RemoveDatastore")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/db"

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

type ApiRemoveIndicatorsRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
	removeIndicatorsRequest *RemoveIndicatorsRequest
}

// The list of indicator IDs to remove
func (r ApiRemoveIndicatorsRequest) RemoveIndicatorsRequest(removeIndicatorsRequest RemoveIndicatorsRequest) ApiRemoveIndicatorsRequest {
	r.removeIndicatorsRequest = &removeIndicatorsRequest
	return r
}

func (r ApiRemoveIndicatorsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveIndicatorsExecute(r)
}

/*
RemoveIndicators Removes indicators by their IDS

Removes indicators by specifying a list of indicator IDs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveIndicatorsRequest
*/
func (a *ThreatIntelIngestApiService) RemoveIndicators(ctx context.Context) ApiRemoveIndicatorsRequest {
	return ApiRemoveIndicatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ThreatIntelIngestApiService) RemoveIndicatorsExecute(r ApiRemoveIndicatorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.RemoveIndicators")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/indicators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.removeIndicatorsRequest == nil {
		return nil, reportError("removeIndicatorsRequest is required and must be specified")
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
	localVarPostBody = r.removeIndicatorsRequest
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

type ApiRetentionPeriodRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
}

func (r ApiRetentionPeriodRequest) Execute() (*DatastoreRetentionPeriod, *http.Response, error) {
	return r.ApiService.RetentionPeriodExecute(r)
}

/*
RetentionPeriod Get threat intel indicators store retention period in terms of days.

Get the threat intel indicators store retention period in terms of days.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetentionPeriodRequest
*/
func (a *ThreatIntelIngestApiService) RetentionPeriod(ctx context.Context) ApiRetentionPeriodRequest {
	return ApiRetentionPeriodRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DatastoreRetentionPeriod
func (a *ThreatIntelIngestApiService) RetentionPeriodExecute(r ApiRetentionPeriodRequest) (*DatastoreRetentionPeriod, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DatastoreRetentionPeriod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.RetentionPeriod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/retentionPeriod"

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

type ApiSetRetentionPeriodRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
	datastoreRetentionPeriod *DatastoreRetentionPeriod
}

// The threat intel indicators store retention period in terms of days.
func (r ApiSetRetentionPeriodRequest) DatastoreRetentionPeriod(datastoreRetentionPeriod DatastoreRetentionPeriod) ApiSetRetentionPeriodRequest {
	r.datastoreRetentionPeriod = &datastoreRetentionPeriod
	return r
}

func (r ApiSetRetentionPeriodRequest) Execute() (*DatastoreRetentionPeriod, *http.Response, error) {
	return r.ApiService.SetRetentionPeriodExecute(r)
}

/*
SetRetentionPeriod Set the threat intel indicators store retention period in terms of days.

Sets the threat intel indicators store retention period in terms of days.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSetRetentionPeriodRequest
*/
func (a *ThreatIntelIngestApiService) SetRetentionPeriod(ctx context.Context) ApiSetRetentionPeriodRequest {
	return ApiSetRetentionPeriodRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DatastoreRetentionPeriod
func (a *ThreatIntelIngestApiService) SetRetentionPeriodExecute(r ApiSetRetentionPeriodRequest) (*DatastoreRetentionPeriod, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DatastoreRetentionPeriod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.SetRetentionPeriod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/retentionPeriod"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.datastoreRetentionPeriod == nil {
		return localVarReturnValue, nil, reportError("datastoreRetentionPeriod is required and must be specified")
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
	localVarPostBody = r.datastoreRetentionPeriod
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

type ApiUploadBlobIndicatorsRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
	uploadBlobIndicatorsRequest *UploadBlobIndicatorsRequest
}

// Upload blob indicators request body.
func (r ApiUploadBlobIndicatorsRequest) UploadBlobIndicatorsRequest(uploadBlobIndicatorsRequest UploadBlobIndicatorsRequest) ApiUploadBlobIndicatorsRequest {
	r.uploadBlobIndicatorsRequest = &uploadBlobIndicatorsRequest
	return r
}

func (r ApiUploadBlobIndicatorsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadBlobIndicatorsExecute(r)
}

/*
UploadBlobIndicators Uploads indicators in a blob format to be parsed (CSV or JSON).

Uploads indicators in a blob format to be parsed (CSV or JSON).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadBlobIndicatorsRequest
*/
func (a *ThreatIntelIngestApiService) UploadBlobIndicators(ctx context.Context) ApiUploadBlobIndicatorsRequest {
	return ApiUploadBlobIndicatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ThreatIntelIngestApiService) UploadBlobIndicatorsExecute(r ApiUploadBlobIndicatorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.UploadBlobIndicators")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/indicators/blob"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadBlobIndicatorsRequest == nil {
		return nil, reportError("uploadBlobIndicatorsRequest is required and must be specified")
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
	localVarPostBody = r.uploadBlobIndicatorsRequest
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

type ApiUploadCsvIndicatorsRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
	uploadCsvIndicatorsRequest *UploadCsvIndicatorsRequest
}

// Upload CSV indicators request body.
func (r ApiUploadCsvIndicatorsRequest) UploadCsvIndicatorsRequest(uploadCsvIndicatorsRequest UploadCsvIndicatorsRequest) ApiUploadCsvIndicatorsRequest {
	r.uploadCsvIndicatorsRequest = &uploadCsvIndicatorsRequest
	return r
}

func (r ApiUploadCsvIndicatorsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadCsvIndicatorsExecute(r)
}

/*
UploadCsvIndicators Uploads indicators in CSV format

Uploads indicators in a CSV format, where the fields are:
  1. ID
  2. Value
  3. Indicator Type
  4. Source
  5. From
  6. Until
  7. Confidence
  8. Threat Type
  9. Actors

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadCsvIndicatorsRequest
*/
func (a *ThreatIntelIngestApiService) UploadCsvIndicators(ctx context.Context) ApiUploadCsvIndicatorsRequest {
	return ApiUploadCsvIndicatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ThreatIntelIngestApiService) UploadCsvIndicatorsExecute(r ApiUploadCsvIndicatorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.UploadCsvIndicators")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/indicators/csv"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadCsvIndicatorsRequest == nil {
		return nil, reportError("uploadCsvIndicatorsRequest is required and must be specified")
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
	localVarPostBody = r.uploadCsvIndicatorsRequest
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

type ApiUploadNormalizedIndicatorsRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
	uploadNormalizedIndicatorRequest *UploadNormalizedIndicatorRequest
}

// The list of normalized threat intel indicators to upload.
func (r ApiUploadNormalizedIndicatorsRequest) UploadNormalizedIndicatorRequest(uploadNormalizedIndicatorRequest UploadNormalizedIndicatorRequest) ApiUploadNormalizedIndicatorsRequest {
	r.uploadNormalizedIndicatorRequest = &uploadNormalizedIndicatorRequest
	return r
}

func (r ApiUploadNormalizedIndicatorsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadNormalizedIndicatorsExecute(r)
}

/*
UploadNormalizedIndicators Uploads indicators in a Sumo normalized format.

Uploads a list indicators in a Sumo normalized format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadNormalizedIndicatorsRequest
*/
func (a *ThreatIntelIngestApiService) UploadNormalizedIndicators(ctx context.Context) ApiUploadNormalizedIndicatorsRequest {
	return ApiUploadNormalizedIndicatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ThreatIntelIngestApiService) UploadNormalizedIndicatorsExecute(r ApiUploadNormalizedIndicatorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.UploadNormalizedIndicators")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/indicators/normalized"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadNormalizedIndicatorRequest == nil {
		return nil, reportError("uploadNormalizedIndicatorRequest is required and must be specified")
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
	localVarPostBody = r.uploadNormalizedIndicatorRequest
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

type ApiUploadStixIndicatorsRequest struct {
	ctx context.Context
	ApiService *ThreatIntelIngestApiService
	uploadStixIndicatorsRequest *UploadStixIndicatorsRequest
}

// Upload stix indicators request body.
func (r ApiUploadStixIndicatorsRequest) UploadStixIndicatorsRequest(uploadStixIndicatorsRequest UploadStixIndicatorsRequest) ApiUploadStixIndicatorsRequest {
	r.uploadStixIndicatorsRequest = &uploadStixIndicatorsRequest
	return r
}

func (r ApiUploadStixIndicatorsRequest) Execute() (*UploadStixIndicatorsResponse, *http.Response, error) {
	return r.ApiService.UploadStixIndicatorsExecute(r)
}

/*
UploadStixIndicators Uploads indicators in a STIX 2.1 json format.

Uploads a list indicators in in a STIX 2.1 json format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadStixIndicatorsRequest
*/
func (a *ThreatIntelIngestApiService) UploadStixIndicators(ctx context.Context) ApiUploadStixIndicatorsRequest {
	return ApiUploadStixIndicatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UploadStixIndicatorsResponse
func (a *ThreatIntelIngestApiService) UploadStixIndicatorsExecute(r ApiUploadStixIndicatorsRequest) (*UploadStixIndicatorsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadStixIndicatorsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreatIntelIngestApiService.UploadStixIndicators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threatIntel/datastore/indicators/stix"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadStixIndicatorsRequest == nil {
		return localVarReturnValue, nil, reportError("uploadStixIndicatorsRequest is required and must be specified")
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
	localVarPostBody = r.uploadStixIndicatorsRequest
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
