# \SchemaBaseManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchemaIdentitiesGrouped**](SchemaBaseManagementAPI.md#GetSchemaIdentitiesGrouped) | **Get** /v1/schemaIdentitiesGrouped | Get schema base identities grouped by type and sorted by version.



## GetSchemaIdentitiesGrouped

> ListSchemaBaseTypeToVersionsResponse GetSchemaIdentitiesGrouped(ctx).Execute()

Get schema base identities grouped by type and sorted by version.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaBaseManagementAPI.GetSchemaIdentitiesGrouped(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaBaseManagementAPI.GetSchemaIdentitiesGrouped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchemaIdentitiesGrouped`: ListSchemaBaseTypeToVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SchemaBaseManagementAPI.GetSchemaIdentitiesGrouped`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaIdentitiesGroupedRequest struct via the builder pattern


### Return type

[**ListSchemaBaseTypeToVersionsResponse**](ListSchemaBaseTypeToVersionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

