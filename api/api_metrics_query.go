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
)


// MetricsQueryApiService MetricsQueryApi service
type MetricsQueryApiService service

type ApiRunMetricsQueriesRequest struct {
	ctx context.Context
	ApiService *MetricsQueryApiService
	metricsQueryRequest *MetricsQueryRequest
}

// The parameters for the metrics query.
func (r ApiRunMetricsQueriesRequest) MetricsQueryRequest(metricsQueryRequest MetricsQueryRequest) ApiRunMetricsQueriesRequest {
	r.metricsQueryRequest = &metricsQueryRequest
	return r
}

func (r ApiRunMetricsQueriesRequest) Execute() (*MetricsQueryResponse, *http.Response, error) {
	return r.ApiService.RunMetricsQueriesExecute(r)
}

/*
RunMetricsQueries Run metrics queries

Execute up to six metrics queries. If you specify multiple queries, each is returned as a separate set of time series. A metric query returns a maximum of 300 data points per metric. A metric query will process a maximum of 15K unique time series to calculate the query results. Query results are limited to 1000 unique time series.
For more information see [Metrics Queries](https://help.sumologic.com/?cid=10144).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRunMetricsQueriesRequest
*/
func (a *MetricsQueryApiService) RunMetricsQueries(ctx context.Context) ApiRunMetricsQueriesRequest {
	return ApiRunMetricsQueriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MetricsQueryResponse
func (a *MetricsQueryApiService) RunMetricsQueriesExecute(r ApiRunMetricsQueriesRequest) (*MetricsQueryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MetricsQueryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsQueryApiService.RunMetricsQueries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metricsQueries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.metricsQueryRequest == nil {
		return localVarReturnValue, nil, reportError("metricsQueryRequest is required and must be specified")
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
	localVarPostBody = r.metricsQueryRequest
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
