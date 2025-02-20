# \SourceTemplateManagementExternalAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSourceTemplate**](SourceTemplateManagementExternalAPI.md#CreateSourceTemplate) | **Post** /v1/sourceTemplate | Create source Template.
[**DeleteSourceTemplate**](SourceTemplateManagementExternalAPI.md#DeleteSourceTemplate) | **Delete** /v1/sourceTemplate/{id} | Delete a Source Template.
[**GetLinkedSourceTemplatesUpdate**](SourceTemplateManagementExternalAPI.md#GetLinkedSourceTemplatesUpdate) | **Post** /v1/sourceTemplate/getLinkedSourceTemplatesImpact | Get linked source templates update based on the ot-collector tags user is wants to update.
[**GetSourceTemplate**](SourceTemplateManagementExternalAPI.md#GetSourceTemplate) | **Get** /v1/sourceTemplate/{id} | get a Source Template by Id.
[**GetSourceTemplates**](SourceTemplateManagementExternalAPI.md#GetSourceTemplates) | **Get** /v1/sourceTemplate | Return all source templates of a customer.
[**UpdateSourceTemplate**](SourceTemplateManagementExternalAPI.md#UpdateSourceTemplate) | **Post** /v1/sourceTemplate/{id} | Update source Template.
[**UpdateSourceTemplateStatus**](SourceTemplateManagementExternalAPI.md#UpdateSourceTemplateStatus) | **Put** /v1/sourceTemplate/{id}/status | Update status of source template
[**UpgradeSourceTemplate**](SourceTemplateManagementExternalAPI.md#UpgradeSourceTemplate) | **Post** /v1/upgrade/sourceTemplate/{id} | Upgrade source Template.



## CreateSourceTemplate

> SourceTemplateDefinition CreateSourceTemplate(ctx).SourceTemplateRequest(sourceTemplateRequest).Execute()

Create source Template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	sourceTemplateRequest := *openapiclient.NewSourceTemplateRequest(*openapiclient.NewSchemaRef("Apache"), *openapiclient.NewSourceTemplateRequestInputJson("apache_test_source_template", map[string]interface{}({}))) // SourceTemplateRequest | Create source Template Details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceTemplateManagementExternalAPI.CreateSourceTemplate(context.Background()).SourceTemplateRequest(sourceTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTemplateManagementExternalAPI.CreateSourceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSourceTemplate`: SourceTemplateDefinition
	fmt.Fprintf(os.Stdout, "Response from `SourceTemplateManagementExternalAPI.CreateSourceTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceTemplateRequest** | [**SourceTemplateRequest**](SourceTemplateRequest.md) | Create source Template Details | 

### Return type

[**SourceTemplateDefinition**](SourceTemplateDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSourceTemplate

> DeleteSourceTemplate(ctx, id).Execute()

Delete a Source Template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	id := "id_example" // string | Identifier of the Source Template to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourceTemplateManagementExternalAPI.DeleteSourceTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTemplateManagementExternalAPI.DeleteSourceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Source Template to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedSourceTemplatesUpdate

> LinkedSourceTemplatesUpdateResponse GetLinkedSourceTemplatesUpdate(ctx).LinkedSourceTemplatesUpdateRequest(linkedSourceTemplatesUpdateRequest).Execute()

Get linked source templates update based on the ot-collector tags user is wants to update.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	linkedSourceTemplatesUpdateRequest := *openapiclient.NewLinkedSourceTemplatesUpdateRequest("00005AF3107BF0D6") // LinkedSourceTemplatesUpdateRequest | request body containing otCollector id and set of tags.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceTemplateManagementExternalAPI.GetLinkedSourceTemplatesUpdate(context.Background()).LinkedSourceTemplatesUpdateRequest(linkedSourceTemplatesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTemplateManagementExternalAPI.GetLinkedSourceTemplatesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedSourceTemplatesUpdate`: LinkedSourceTemplatesUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `SourceTemplateManagementExternalAPI.GetLinkedSourceTemplatesUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedSourceTemplatesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkedSourceTemplatesUpdateRequest** | [**LinkedSourceTemplatesUpdateRequest**](LinkedSourceTemplatesUpdateRequest.md) | request body containing otCollector id and set of tags. | 

### Return type

[**LinkedSourceTemplatesUpdateResponse**](LinkedSourceTemplatesUpdateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceTemplate

> SourceTemplateDefinition GetSourceTemplate(ctx, id).Execute()

get a Source Template by Id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	id := "id_example" // string | Identifier of the Source Template to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceTemplateManagementExternalAPI.GetSourceTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTemplateManagementExternalAPI.GetSourceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceTemplate`: SourceTemplateDefinition
	fmt.Fprintf(os.Stdout, "Response from `SourceTemplateManagementExternalAPI.GetSourceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Source Template to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceTemplateDefinition**](SourceTemplateDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceTemplates

> SourceTemplateListResponse GetSourceTemplates(ctx).ShowDisabled(showDisabled).Execute()

Return all source templates of a customer.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	showDisabled := true // bool | A boolean parameter to get all, including disabled source templates. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceTemplateManagementExternalAPI.GetSourceTemplates(context.Background()).ShowDisabled(showDisabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTemplateManagementExternalAPI.GetSourceTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceTemplates`: SourceTemplateListResponse
	fmt.Fprintf(os.Stdout, "Response from `SourceTemplateManagementExternalAPI.GetSourceTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showDisabled** | **bool** | A boolean parameter to get all, including disabled source templates. | [default to false]

### Return type

[**SourceTemplateListResponse**](SourceTemplateListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceTemplate

> SourceTemplateDefinition UpdateSourceTemplate(ctx, id).SourceTemplateRequest(sourceTemplateRequest).Execute()

Update source Template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	id := "id_example" // string | Identifier of the Source Template to update.
	sourceTemplateRequest := *openapiclient.NewSourceTemplateRequest(*openapiclient.NewSchemaRef("Apache"), *openapiclient.NewSourceTemplateRequestInputJson("apache_test_source_template", map[string]interface{}({}))) // SourceTemplateRequest | Source Template request details.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceTemplateManagementExternalAPI.UpdateSourceTemplate(context.Background(), id).SourceTemplateRequest(sourceTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTemplateManagementExternalAPI.UpdateSourceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSourceTemplate`: SourceTemplateDefinition
	fmt.Fprintf(os.Stdout, "Response from `SourceTemplateManagementExternalAPI.UpdateSourceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Source Template to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceTemplateRequest** | [**SourceTemplateRequest**](SourceTemplateRequest.md) | Source Template request details. | 

### Return type

[**SourceTemplateDefinition**](SourceTemplateDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceTemplateStatus

> SourceTemplateDefinition UpdateSourceTemplateStatus(ctx, id).SourceTemplateStatusUpdateRequest(sourceTemplateStatusUpdateRequest).Execute()

Update status of source template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	id := "id_example" // string | Identifier of the source template to update.
	sourceTemplateStatusUpdateRequest := *openapiclient.NewSourceTemplateStatusUpdateRequest("Status_example") // SourceTemplateStatusUpdateRequest | Status of source template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceTemplateManagementExternalAPI.UpdateSourceTemplateStatus(context.Background(), id).SourceTemplateStatusUpdateRequest(sourceTemplateStatusUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTemplateManagementExternalAPI.UpdateSourceTemplateStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSourceTemplateStatus`: SourceTemplateDefinition
	fmt.Fprintf(os.Stdout, "Response from `SourceTemplateManagementExternalAPI.UpdateSourceTemplateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the source template to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceTemplateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceTemplateStatusUpdateRequest** | [**SourceTemplateStatusUpdateRequest**](SourceTemplateStatusUpdateRequest.md) | Status of source template | 

### Return type

[**SourceTemplateDefinition**](SourceTemplateDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSourceTemplate

> SourceTemplateDefinition UpgradeSourceTemplate(ctx, id).SourceTemplateUpgradeRequest(sourceTemplateUpgradeRequest).Execute()

Upgrade source Template.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	id := "id_example" // string | Identifier of the Source Template to upgrade.
	sourceTemplateUpgradeRequest := *openapiclient.NewSourceTemplateUpgradeRequest(*openapiclient.NewUpgradeSchemaRef("Apache", "1.0.0"), *openapiclient.NewSourceTemplateUpgradeRequestInputJson("apache_test_source_template", map[string]interface{}({"hostmetrics":{"receiverType":"hostmetrics","collection_interval":"5m"}}))) // SourceTemplateUpgradeRequest | Source Template upgrade request details.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceTemplateManagementExternalAPI.UpgradeSourceTemplate(context.Background(), id).SourceTemplateUpgradeRequest(sourceTemplateUpgradeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceTemplateManagementExternalAPI.UpgradeSourceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeSourceTemplate`: SourceTemplateDefinition
	fmt.Fprintf(os.Stdout, "Response from `SourceTemplateManagementExternalAPI.UpgradeSourceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Source Template to upgrade. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSourceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceTemplateUpgradeRequest** | [**SourceTemplateUpgradeRequest**](SourceTemplateUpgradeRequest.md) | Source Template upgrade request details. | 

### Return type

[**SourceTemplateDefinition**](SourceTemplateDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

