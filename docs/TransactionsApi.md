# \TransactionsApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /v1/transactions/{transactionId} | Get transaction



## GetTransaction

> TransactionModel GetTransaction(ctx, transactionId).Execute()

Get transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionId := "transactionId_example" // string | Any transaction id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Any transaction id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

