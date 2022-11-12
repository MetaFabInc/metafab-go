# \WalletsApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWalletBalances**](WalletsApi.md#GetWalletBalances) | **Get** /v1/wallets/{walletId}/balances | Get wallet balances
[**GetWalletTransactions**](WalletsApi.md#GetWalletTransactions) | **Get** /v1/wallets/{walletId}/transactions | Get wallet transactions



## GetWalletBalances

> map[string]float32 GetWalletBalances(ctx, walletId).Execute()

Get wallet balances



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
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWalletBalances(context.Background(), walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWalletBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletBalances`: map[string]float32
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWalletBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Any wallet id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletTransactions

> []TransactionModel GetWalletTransactions(ctx, walletId).Execute()

Get wallet transactions



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
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWalletTransactions(context.Background(), walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWalletTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletTransactions`: []TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWalletTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Any wallet id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

