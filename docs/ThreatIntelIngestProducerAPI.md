# \ThreatIntelIngestProducerAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveIndicators**](ThreatIntelIngestProducerAPI.md#RemoveIndicators) | **Delete** /v1/threatIntel/datastore/indicators | Removes indicators by their IDS
[**UploadNormalizedIndicators**](ThreatIntelIngestProducerAPI.md#UploadNormalizedIndicators) | **Post** /v1/threatIntel/datastore/indicators/normalized | Uploads indicators in a Sumo normalized format.
[**UploadStixIndicators**](ThreatIntelIngestProducerAPI.md#UploadStixIndicators) | **Post** /v1/threatIntel/datastore/indicators/stix | Uploads indicators in a STIX 2.x json format.



## RemoveIndicators

> RemoveIndicators(ctx).RemoveIndicatorsRequest(removeIndicatorsRequest).Execute()

Removes indicators by their IDS



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	removeIndicatorsRequest := *openapiclient.NewRemoveIndicatorsRequest("Crowdstrike", []string{"IndicatorIds_example"}) // RemoveIndicatorsRequest | The list of indicator IDs to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreatIntelIngestProducerAPI.RemoveIndicators(context.Background()).RemoveIndicatorsRequest(removeIndicatorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestProducerAPI.RemoveIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeIndicatorsRequest** | [**RemoveIndicatorsRequest**](RemoveIndicatorsRequest.md) | The list of indicator IDs to remove | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadNormalizedIndicators

> UploadNormalizedIndicators(ctx).UploadNormalizedIndicatorRequest(uploadNormalizedIndicatorRequest).Execute()

Uploads indicators in a Sumo normalized format.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uploadNormalizedIndicatorRequest := *openapiclient.NewUploadNormalizedIndicatorRequest([]openapiclient.NormalizedIndicator{*openapiclient.NewNormalizedIndicator("indicator--d81f86b9-975b-4c0b-875e-810c5ad45a4f", "182.158.1.1", "ipv4-addr:value", "FreeTAXII", time.Now(), int32(123), "indicator")}) // UploadNormalizedIndicatorRequest | The list of normalized threat intel indicators to upload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreatIntelIngestProducerAPI.UploadNormalizedIndicators(context.Background()).UploadNormalizedIndicatorRequest(uploadNormalizedIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestProducerAPI.UploadNormalizedIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadNormalizedIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadNormalizedIndicatorRequest** | [**UploadNormalizedIndicatorRequest**](UploadNormalizedIndicatorRequest.md) | The list of normalized threat intel indicators to upload. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadStixIndicators

> UploadStixIndicatorsResponse UploadStixIndicators(ctx).UploadStixIndicatorsRequest(uploadStixIndicatorsRequest).Execute()

Uploads indicators in a STIX 2.x json format.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uploadStixIndicatorsRequest := *openapiclient.NewUploadStixIndicatorsRequest("FreeTAXII", []openapiclient.StixIndicator{*openapiclient.NewStixIndicator("indicator", "2.1", "acme:indicator-bf8bc5d5-c7e6-46b0-8d22-7500fea77196", time.Now(), time.Now(), "[ipv4-addr:value = '1.2.3.4']", "stix", time.Now())}) // UploadStixIndicatorsRequest | Upload stix indicators request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatIntelIngestProducerAPI.UploadStixIndicators(context.Background()).UploadStixIndicatorsRequest(uploadStixIndicatorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestProducerAPI.UploadStixIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadStixIndicators`: UploadStixIndicatorsResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatIntelIngestProducerAPI.UploadStixIndicators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadStixIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadStixIndicatorsRequest** | [**UploadStixIndicatorsRequest**](UploadStixIndicatorsRequest.md) | Upload stix indicators request body. | 

### Return type

[**UploadStixIndicatorsResponse**](UploadStixIndicatorsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

