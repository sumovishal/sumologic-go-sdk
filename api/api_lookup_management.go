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
	"os"
)


// LookupManagementApiService LookupManagementApi service
type LookupManagementApiService service

type ApiCreateTableRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	lookupTableDefinition *LookupTableDefinition
}

// The schema and configuration for the lookup table.
func (r ApiCreateTableRequest) LookupTableDefinition(lookupTableDefinition LookupTableDefinition) ApiCreateTableRequest {
	r.lookupTableDefinition = &lookupTableDefinition
	return r
}

func (r ApiCreateTableRequest) Execute() (*LookupTable, *http.Response, error) {
	return r.ApiService.CreateTableExecute(r)
}

/*
CreateTable Create a lookup table.

Create a new lookup table by providing a schema and specifying its configuration. Providing parentFolderId
 is mandatory. Use the [getItemByPath](#operation/getItemByPath) endpoint to get content id of a path.
Please check [Content management API](#tag/contentManagement) and [Folder management API](#tag/folderManagement) for all available options.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTableRequest
*/
func (a *LookupManagementApiService) CreateTable(ctx context.Context) ApiCreateTableRequest {
	return ApiCreateTableRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LookupTable
func (a *LookupManagementApiService) CreateTableExecute(r ApiCreateTableRequest) (*LookupTable, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LookupTable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.CreateTable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lookupTableDefinition == nil {
		return localVarReturnValue, nil, reportError("lookupTableDefinition is required and must be specified")
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
	localVarPostBody = r.lookupTableDefinition
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

type ApiDeleteTableRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	id string
}

func (r ApiDeleteTableRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTableExecute(r)
}

/*
DeleteTable Delete a lookup table.

Delete a lookup table completely. <br> **Warning:** `This operation cannot be undone`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the lookup table.
 @return ApiDeleteTableRequest
*/
func (a *LookupManagementApiService) DeleteTable(ctx context.Context, id string) ApiDeleteTableRequest {
	return ApiDeleteTableRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *LookupManagementApiService) DeleteTableExecute(r ApiDeleteTableRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.DeleteTable")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables/{id}"
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

type ApiDeleteTableRowRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	id string
	rowDeleteDefinition *RowDeleteDefinition
}

// Lookup table row delete definition.
func (r ApiDeleteTableRowRequest) RowDeleteDefinition(rowDeleteDefinition RowDeleteDefinition) ApiDeleteTableRowRequest {
	r.rowDeleteDefinition = &rowDeleteDefinition
	return r
}

func (r ApiDeleteTableRowRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTableRowExecute(r)
}

/*
DeleteTableRow Delete a lookup table row.

Delete a row from lookup table by providing the row's primary keys' values. The complete set of primary key fields of the lookup table should be provided.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the lookup table.
 @return ApiDeleteTableRowRequest
*/
func (a *LookupManagementApiService) DeleteTableRow(ctx context.Context, id string) ApiDeleteTableRowRequest {
	return ApiDeleteTableRowRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *LookupManagementApiService) DeleteTableRowExecute(r ApiDeleteTableRowRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.DeleteTableRow")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables/{id}/deleteTableRow"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rowDeleteDefinition == nil {
		return nil, reportError("rowDeleteDefinition is required and must be specified")
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
	localVarPostBody = r.rowDeleteDefinition
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

type ApiLookupTableByIdRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	id string
}

func (r ApiLookupTableByIdRequest) Execute() (*LookupTable, *http.Response, error) {
	return r.ApiService.LookupTableByIdExecute(r)
}

/*
LookupTableById Get a lookup table.

Get a lookup table for the given identifier.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the lookup table.
 @return ApiLookupTableByIdRequest
*/
func (a *LookupManagementApiService) LookupTableById(ctx context.Context, id string) ApiLookupTableByIdRequest {
	return ApiLookupTableByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return LookupTable
func (a *LookupManagementApiService) LookupTableByIdExecute(r ApiLookupTableByIdRequest) (*LookupTable, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LookupTable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.LookupTableById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables/{id}"
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

type ApiRequestJobStatusRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	jobId string
}

func (r ApiRequestJobStatusRequest) Execute() (*LookupAsyncJobStatus, *http.Response, error) {
	return r.ApiService.RequestJobStatusExecute(r)
}

/*
RequestJobStatus Get the status of an async job.

Retrieve the status of a previously made request. If the request was successful, the status of the response object will be `Success`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId An identifier returned in response to an asynchronous request.
 @return ApiRequestJobStatusRequest
*/
func (a *LookupManagementApiService) RequestJobStatus(ctx context.Context, jobId string) ApiRequestJobStatusRequest {
	return ApiRequestJobStatusRequest{
		ApiService: a,
		ctx: ctx,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return LookupAsyncJobStatus
func (a *LookupManagementApiService) RequestJobStatusExecute(r ApiRequestJobStatusRequest) (*LookupAsyncJobStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LookupAsyncJobStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.RequestJobStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables/jobs/{jobId}/status"
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

type ApiTruncateTableRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	id string
}

func (r ApiTruncateTableRequest) Execute() (*LookupRequestToken, *http.Response, error) {
	return r.ApiService.TruncateTableExecute(r)
}

/*
TruncateTable Empty a lookup table.

Delete all data from a lookup table.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the table to clear.
 @return ApiTruncateTableRequest
*/
func (a *LookupManagementApiService) TruncateTable(ctx context.Context, id string) ApiTruncateTableRequest {
	return ApiTruncateTableRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return LookupRequestToken
func (a *LookupManagementApiService) TruncateTableExecute(r ApiTruncateTableRequest) (*LookupRequestToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LookupRequestToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.TruncateTable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables/{id}/truncate"
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

type ApiUpdateTableRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	id string
	lookupUpdateDefinition *LookupUpdateDefinition
}

// The configuration changes for the lookup table.
func (r ApiUpdateTableRequest) LookupUpdateDefinition(lookupUpdateDefinition LookupUpdateDefinition) ApiUpdateTableRequest {
	r.lookupUpdateDefinition = &lookupUpdateDefinition
	return r
}

func (r ApiUpdateTableRequest) Execute() (*LookupTable, *http.Response, error) {
	return r.ApiService.UpdateTableExecute(r)
}

/*
UpdateTable Edit a lookup table.

Edit the lookup table data. All the fields are mandatory in the request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the lookup table.
 @return ApiUpdateTableRequest
*/
func (a *LookupManagementApiService) UpdateTable(ctx context.Context, id string) ApiUpdateTableRequest {
	return ApiUpdateTableRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return LookupTable
func (a *LookupManagementApiService) UpdateTableExecute(r ApiUpdateTableRequest) (*LookupTable, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LookupTable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.UpdateTable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lookupUpdateDefinition == nil {
		return localVarReturnValue, nil, reportError("lookupUpdateDefinition is required and must be specified")
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
	localVarPostBody = r.lookupUpdateDefinition
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

type ApiUpdateTableRowRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	id string
	rowUpdateDefinition *RowUpdateDefinition
}

// Lookup table row update definition.
func (r ApiUpdateTableRowRequest) RowUpdateDefinition(rowUpdateDefinition RowUpdateDefinition) ApiUpdateTableRowRequest {
	r.rowUpdateDefinition = &rowUpdateDefinition
	return r
}

func (r ApiUpdateTableRowRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateTableRowExecute(r)
}

/*
UpdateTableRow Insert or Update a lookup table row.

Insert or update a row of a lookup table with the given identifier. A new row is inserted if the primary key does not exist already, otherwise the existing row with the specified primary key is updated. All the fields of the lookup table are required and will be updated to the given values. In case a field is not specified then it will be assumed to be set to null. If the table size exceeds the maximum limit of 100MB then based on the size limit action of the table the update will be processed or discarded.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the lookup table.
 @return ApiUpdateTableRowRequest
*/
func (a *LookupManagementApiService) UpdateTableRow(ctx context.Context, id string) ApiUpdateTableRowRequest {
	return ApiUpdateTableRowRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *LookupManagementApiService) UpdateTableRowExecute(r ApiUpdateTableRowRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.UpdateTableRow")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables/{id}/row"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rowUpdateDefinition == nil {
		return nil, reportError("rowUpdateDefinition is required and must be specified")
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
	localVarPostBody = r.rowUpdateDefinition
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

type ApiUploadFileRequest struct {
	ctx context.Context
	ApiService *LookupManagementApiService
	id string
	file *os.File
	merge *bool
	fileEncoding *string
}

// The CSV file to upload.   - The size limit for the CSV file is 100MB.   - Use Unix format, with newlines (\\\&quot;\\\\n\\\&quot;) separating rows.   - The first row should contain headers that match the lookup table schema. Matching is     case-insensitive.
func (r ApiUploadFileRequest) File(file *os.File) ApiUploadFileRequest {
	r.file = file
	return r
}

// This indicates whether the file contents will be merged with existing data in the lookup table or not. If this is true then data with the same primary keys will be updated while the rest of the rows will be appended. By default, merge is false. The response includes a request identifier that you need to use in the [Request Status API](#operation/requestStatus) to track the status of the upload request.
func (r ApiUploadFileRequest) Merge(merge bool) ApiUploadFileRequest {
	r.merge = &merge
	return r
}

// File encoding of file being uploaded.
func (r ApiUploadFileRequest) FileEncoding(fileEncoding string) ApiUploadFileRequest {
	r.fileEncoding = &fileEncoding
	return r
}

func (r ApiUploadFileRequest) Execute() (*LookupRequestToken, *http.Response, error) {
	return r.ApiService.UploadFileExecute(r)
}

/*
UploadFile Upload a CSV file.

Create a request to populate a lookup table with a CSV file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifier of the lookup table to populate.
 @return ApiUploadFileRequest
*/
func (a *LookupManagementApiService) UploadFile(ctx context.Context, id string) ApiUploadFileRequest {
	return ApiUploadFileRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return LookupRequestToken
func (a *LookupManagementApiService) UploadFileExecute(r ApiUploadFileRequest) (*LookupRequestToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LookupRequestToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LookupManagementApiService.UploadFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lookupTables/{id}/upload"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	if r.merge != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "merge", r.merge, "")
	}
	if r.fileEncoding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fileEncoding", r.fileEncoding, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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
