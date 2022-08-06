/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// AccessKeyManagementApiService AccessKeyManagementApi service
type AccessKeyManagementApiService service

type ApiCreateAccessKeyRequest struct {
	ctx context.Context
	ApiService *AccessKeyManagementApiService
	accessKeyCreateRequest *AccessKeyCreateRequest
}

func (r ApiCreateAccessKeyRequest) AccessKeyCreateRequest(accessKeyCreateRequest AccessKeyCreateRequest) ApiCreateAccessKeyRequest {
	r.accessKeyCreateRequest = &accessKeyCreateRequest
	return r
}

func (r ApiCreateAccessKeyRequest) Execute() (*AccessKey, *http.Response, error) {
	return r.ApiService.CreateAccessKeyExecute(r)
}

/*
CreateAccessKey Create an access key.

Creates a new access ID and key pair. The new access key can be used from the domains specified in corsHeaders field. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:
  1. Requests with an ORIGIN header but the allowlist is empty.
  2. Requests with an ORIGIN header that don't match any entry in the allowlist.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAccessKeyRequest
*/
func (a *AccessKeyManagementApiService) CreateAccessKey(ctx context.Context) ApiCreateAccessKeyRequest {
	return ApiCreateAccessKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccessKey
func (a *AccessKeyManagementApiService) CreateAccessKeyExecute(r ApiCreateAccessKeyRequest) (*AccessKey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessKeyManagementApiService.CreateAccessKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accessKeys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessKeyCreateRequest == nil {
		return localVarReturnValue, nil, reportError("accessKeyCreateRequest is required and must be specified")
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
	localVarPostBody = r.accessKeyCreateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteAccessKeyRequest struct {
	ctx context.Context
	ApiService *AccessKeyManagementApiService
	id string
}

func (r ApiDeleteAccessKeyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAccessKeyExecute(r)
}

/*
DeleteAccessKey Delete an access key.

Deletes the access key with the given accessId.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The accessId of the access key to delete.
 @return ApiDeleteAccessKeyRequest
*/
func (a *AccessKeyManagementApiService) DeleteAccessKey(ctx context.Context, id string) ApiDeleteAccessKeyRequest {
	return ApiDeleteAccessKeyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AccessKeyManagementApiService) DeleteAccessKeyExecute(r ApiDeleteAccessKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessKeyManagementApiService.DeleteAccessKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accessKeys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListAccessKeysRequest struct {
	ctx context.Context
	ApiService *AccessKeyManagementApiService
	limit *int32
	token *string
}

// Limit the number of access keys returned in the response. The number of access keys returned may be less than the &#x60;limit&#x60;.
func (r ApiListAccessKeysRequest) Limit(limit int32) ApiListAccessKeysRequest {
	r.limit = &limit
	return r
}

// Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left.
func (r ApiListAccessKeysRequest) Token(token string) ApiListAccessKeysRequest {
	r.token = &token
	return r
}

func (r ApiListAccessKeysRequest) Execute() (*PaginatedListAccessKeysResult, *http.Response, error) {
	return r.ApiService.ListAccessKeysExecute(r)
}

/*
ListAccessKeys List all access keys.

List all access keys in your account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAccessKeysRequest
*/
func (a *AccessKeyManagementApiService) ListAccessKeys(ctx context.Context) ApiListAccessKeysRequest {
	return ApiListAccessKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedListAccessKeysResult
func (a *AccessKeyManagementApiService) ListAccessKeysExecute(r ApiListAccessKeysRequest) (*PaginatedListAccessKeysResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedListAccessKeysResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessKeyManagementApiService.ListAccessKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accessKeys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListPersonalAccessKeysRequest struct {
	ctx context.Context
	ApiService *AccessKeyManagementApiService
}

func (r ApiListPersonalAccessKeysRequest) Execute() (*ListAccessKeysResult, *http.Response, error) {
	return r.ApiService.ListPersonalAccessKeysExecute(r)
}

/*
ListPersonalAccessKeys List personal keys.

List all access keys that belong to your user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPersonalAccessKeysRequest
*/
func (a *AccessKeyManagementApiService) ListPersonalAccessKeys(ctx context.Context) ApiListPersonalAccessKeysRequest {
	return ApiListPersonalAccessKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListAccessKeysResult
func (a *AccessKeyManagementApiService) ListPersonalAccessKeysExecute(r ApiListPersonalAccessKeysRequest) (*ListAccessKeysResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAccessKeysResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessKeyManagementApiService.ListPersonalAccessKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accessKeys/personal"

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateAccessKeyRequest struct {
	ctx context.Context
	ApiService *AccessKeyManagementApiService
	id string
	accessKeyUpdateRequest *AccessKeyUpdateRequest
}

func (r ApiUpdateAccessKeyRequest) AccessKeyUpdateRequest(accessKeyUpdateRequest AccessKeyUpdateRequest) ApiUpdateAccessKeyRequest {
	r.accessKeyUpdateRequest = &accessKeyUpdateRequest
	return r
}

func (r ApiUpdateAccessKeyRequest) Execute() (*AccessKeyPublic, *http.Response, error) {
	return r.ApiService.UpdateAccessKeyExecute(r)
}

/*
UpdateAccessKey Update an access key.

Updates the properties of existing accessKey by accessId. It can be used to enable or disable the access key and to update the corsHeaders list.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The accessId of the access key to update.
 @return ApiUpdateAccessKeyRequest
*/
func (a *AccessKeyManagementApiService) UpdateAccessKey(ctx context.Context, id string) ApiUpdateAccessKeyRequest {
	return ApiUpdateAccessKeyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccessKeyPublic
func (a *AccessKeyManagementApiService) UpdateAccessKeyExecute(r ApiUpdateAccessKeyRequest) (*AccessKeyPublic, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessKeyPublic
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessKeyManagementApiService.UpdateAccessKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accessKeys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessKeyUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("accessKeyUpdateRequest is required and must be specified")
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
	localVarPostBody = r.accessKeyUpdateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
