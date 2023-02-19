# \WalletsApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWalletSignature**](WalletsApi.md#CreateWalletSignature) | **Post** /v1/wallets/{walletId}/signatures | Create wallet signature
[**GetWallet**](WalletsApi.md#GetWallet) | **Get** /v1/wallets/{walletId} | Get wallet
[**GetWalletBalances**](WalletsApi.md#GetWalletBalances) | **Get** /v1/wallets/{walletId}/balances | Get wallet balances
[**GetWalletTransactions**](WalletsApi.md#GetWalletTransactions) | **Get** /v1/wallets/{walletId}/transactions | Get wallet transactions



## CreateWalletSignature

> string CreateWalletSignature(ctx, walletId).XWalletDecryptKey(xWalletDecryptKey).CreateWalletSignatureRequest(createWalletSignatureRequest).Execute()

Create wallet signature



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
    walletId := "walletId_example" // string | Any wallet id within the MetaFab platform.
    xWalletDecryptKey := "AXNP8MKb+5SbBtHWrZu5KHh5/BomXY/dMRG/BDUn7a4=" // string | The `walletDecryptKey` for the provided `walletId`.
    createWalletSignatureRequest := *openapiclient.NewCreateWalletSignatureRequest("Message_example") // CreateWalletSignatureRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.CreateWalletSignature(context.Background(), walletId).XWalletDecryptKey(xWalletDecryptKey).CreateWalletSignatureRequest(createWalletSignatureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.CreateWalletSignature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWalletSignature`: string
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.CreateWalletSignature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Any wallet id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletSignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xWalletDecryptKey** | **string** | The &#x60;walletDecryptKey&#x60; for the provided &#x60;walletId&#x60;. | 
 **createWalletSignatureRequest** | [**CreateWalletSignatureRequest**](CreateWalletSignatureRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallet

> WalletModel GetWallet(ctx, walletId).Execute()

Get wallet



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
    walletId := "walletId_example" // string | Any wallet id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWallet(context.Background(), walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWallet`: WalletModel
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** | Any wallet id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WalletModel**](WalletModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    walletId := "walletId_example" // string | Any wallet id within the MetaFab platform.

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
**walletId** | **string** | Any wallet id within the MetaFab platform. | 

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
    walletId := "walletId_example" // string | Any wallet id within the MetaFab platform.

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
**walletId** | **string** | Any wallet id within the MetaFab platform. | 

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

