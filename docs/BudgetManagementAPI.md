# \BudgetManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBudget**](BudgetManagementAPI.md#CreateBudget) | **Post** /v1/budgets | Creates a budget definition
[**DeleteBudget**](BudgetManagementAPI.md#DeleteBudget) | **Delete** /v1/budgets/{budgetId} | Delete budget
[**GetBudget**](BudgetManagementAPI.md#GetBudget) | **Get** /v1/budgets/{budgetId} | Get budget
[**GetBudgetUsage**](BudgetManagementAPI.md#GetBudgetUsage) | **Get** /v1/budgets/{budgetId}/usage | Get budget usage
[**GetBudgetUsages**](BudgetManagementAPI.md#GetBudgetUsages) | **Get** /v1/budgets/usage | Get budget usages
[**GetBudgets**](BudgetManagementAPI.md#GetBudgets) | **Get** /v1/budgets | Get budgets
[**UpdateBudget**](BudgetManagementAPI.md#UpdateBudget) | **Put** /v1/budgets/{budgetId} | Update budget



## CreateBudget

> ScanBudget CreateBudget(ctx).ScanBudgetDefinition(scanBudgetDefinition).Execute()

Creates a budget definition



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
	scanBudgetDefinition := *openapiclient.NewScanBudgetDefinition("Name_example", int64(123), "GB", "ScanBudget", *openapiclient.NewScanBudgetScope([]string{"IncludedUsers_example"}, []string{"ExcludedUsers_example"}, []string{"IncludedRoles_example"}, []string{"ExcludedRoles_example"}), "Query", "PerEntity", "User", "Warn") // ScanBudgetDefinition | Information about the new budget.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetManagementAPI.CreateBudget(context.Background()).ScanBudgetDefinition(scanBudgetDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetManagementAPI.CreateBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBudget`: ScanBudget
	fmt.Fprintf(os.Stdout, "Response from `BudgetManagementAPI.CreateBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scanBudgetDefinition** | [**ScanBudgetDefinition**](ScanBudgetDefinition.md) | Information about the new budget. | 

### Return type

[**ScanBudget**](ScanBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBudget

> DeleteBudget(ctx, budgetId).Execute()

Delete budget



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
	budgetId := "budgetId_example" // string | The id of the budget.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BudgetManagementAPI.DeleteBudget(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetManagementAPI.DeleteBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** | The id of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBudgetRequest struct via the builder pattern


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


## GetBudget

> ScanBudget GetBudget(ctx, budgetId).Execute()

Get budget



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
	budgetId := "budgetId_example" // string | The id of the budget.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetManagementAPI.GetBudget(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetManagementAPI.GetBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudget`: ScanBudget
	fmt.Fprintf(os.Stdout, "Response from `BudgetManagementAPI.GetBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** | The id of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScanBudget**](ScanBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetUsage

> ScanBudgetUsage GetBudgetUsage(ctx, budgetId).Execute()

Get budget usage



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
	budgetId := "budgetId_example" // string | The id of the budget.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetManagementAPI.GetBudgetUsage(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetManagementAPI.GetBudgetUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetUsage`: ScanBudgetUsage
	fmt.Fprintf(os.Stdout, "Response from `BudgetManagementAPI.GetBudgetUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** | The id of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScanBudgetUsage**](ScanBudgetUsage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetUsages

> ScanBudgetUsageList GetBudgetUsages(ctx).Limit(limit).Token(token).Execute()

Get budget usages



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
	limit := int32(56) // int32 | Limit the number of budget usages returned in the response. The number of budget usages returned may be less than the `limit`. (optional) (default to 100)
	token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetManagementAPI.GetBudgetUsages(context.Background()).Limit(limit).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetManagementAPI.GetBudgetUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetUsages`: ScanBudgetUsageList
	fmt.Fprintf(os.Stdout, "Response from `BudgetManagementAPI.GetBudgetUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of budget usages returned in the response. The number of budget usages returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ScanBudgetUsageList**](ScanBudgetUsageList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgets

> ScanBudgetList GetBudgets(ctx).Limit(limit).Token(token).Execute()

Get budgets



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
	limit := int32(56) // int32 | Limit the number of budgets returned in the response. The number of budgets returned may be less than the `limit`. (optional) (default to 100)
	token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetManagementAPI.GetBudgets(context.Background()).Limit(limit).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetManagementAPI.GetBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgets`: ScanBudgetList
	fmt.Fprintf(os.Stdout, "Response from `BudgetManagementAPI.GetBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of budgets returned in the response. The number of budgets returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ScanBudgetList**](ScanBudgetList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBudget

> ScanBudget UpdateBudget(ctx, budgetId).ScanBudgetDefinition(scanBudgetDefinition).Execute()

Update budget



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
	budgetId := "budgetId_example" // string | The id of the budget.
	scanBudgetDefinition := *openapiclient.NewScanBudgetDefinition("Name_example", int64(123), "GB", "ScanBudget", *openapiclient.NewScanBudgetScope([]string{"IncludedUsers_example"}, []string{"ExcludedUsers_example"}, []string{"IncludedRoles_example"}, []string{"ExcludedRoles_example"}), "Query", "PerEntity", "User", "Warn") // ScanBudgetDefinition | Updated budget.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetManagementAPI.UpdateBudget(context.Background(), budgetId).ScanBudgetDefinition(scanBudgetDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetManagementAPI.UpdateBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBudget`: ScanBudget
	fmt.Fprintf(os.Stdout, "Response from `BudgetManagementAPI.UpdateBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** | The id of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scanBudgetDefinition** | [**ScanBudgetDefinition**](ScanBudgetDefinition.md) | Updated budget. | 

### Return type

[**ScanBudget**](ScanBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

