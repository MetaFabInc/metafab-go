# \ItemsApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchMintCollectionItems**](ItemsApi.md#BatchMintCollectionItems) | **Post** /v1/collections/{collectionId}/batchMints | Batch mint collection items
[**BatchTransferCollectionItems**](ItemsApi.md#BatchTransferCollectionItems) | **Post** /v1/collections/{collectionId}/batchTransfers | Batch transfer collection items
[**BurnCollectionItem**](ItemsApi.md#BurnCollectionItem) | **Post** /v1/collections/{collectionId}/items/{collectionItemId}/burns | Burn collection item
[**CreateCollection**](ItemsApi.md#CreateCollection) | **Post** /v1/collections | Create collection
[**CreateCollectionItem**](ItemsApi.md#CreateCollectionItem) | **Post** /v1/collections/{collectionId}/items | Create collection item
[**GetCollectionApproval**](ItemsApi.md#GetCollectionApproval) | **Get** /v1/collections/{collectionId}/approvals | Get collection approval
[**GetCollectionItem**](ItemsApi.md#GetCollectionItem) | **Get** /v1/collections/{collectionId}/items/{collectionItemId} | Get collection item
[**GetCollectionItemBalance**](ItemsApi.md#GetCollectionItemBalance) | **Get** /v1/collections/{collectionId}/items/{collectionItemId}/balances | Get collection item balance
[**GetCollectionItemBalances**](ItemsApi.md#GetCollectionItemBalances) | **Get** /v1/collections/{collectionId}/balances | Get collection item balances
[**GetCollectionItemSupplies**](ItemsApi.md#GetCollectionItemSupplies) | **Get** /v1/collections/{collectionId}/supplies | Get collection item supplies
[**GetCollectionItemSupply**](ItemsApi.md#GetCollectionItemSupply) | **Get** /v1/collections/{collectionId}/items/{collectionItemId}/supplies | Get collection item supply
[**GetCollectionItemTimelock**](ItemsApi.md#GetCollectionItemTimelock) | **Get** /v1/collections/{collectionId}/items/{collectionItemId}/timelocks | Get collection item timelock
[**GetCollectionItems**](ItemsApi.md#GetCollectionItems) | **Get** /v1/collections/{collectionId}/items | Get collection items
[**GetCollectionRole**](ItemsApi.md#GetCollectionRole) | **Get** /v1/collections/{collectionId}/roles | Get collection role
[**GetCollections**](ItemsApi.md#GetCollections) | **Get** /v1/collections | Get collections
[**GrantCollectionRole**](ItemsApi.md#GrantCollectionRole) | **Post** /v1/collections/{collectionId}/roles | Grant collection role
[**MintCollectionItem**](ItemsApi.md#MintCollectionItem) | **Post** /v1/collections/{collectionId}/items/{collectionItemId}/mints | Mint collection item
[**RevokeCollectionRole**](ItemsApi.md#RevokeCollectionRole) | **Delete** /v1/collections/{collectionId}/roles | Revoke collection role
[**SetCollectionApproval**](ItemsApi.md#SetCollectionApproval) | **Post** /v1/collections/{collectionId}/approvals | Set collection approval
[**SetCollectionItemTimelock**](ItemsApi.md#SetCollectionItemTimelock) | **Post** /v1/collections/{collectionId}/items/{collectionItemId}/timelocks | Set collection item timelock
[**TransferCollectionItem**](ItemsApi.md#TransferCollectionItem) | **Post** /v1/collections/{collectionId}/items/{collectionItemId}/transfers | Transfer collection item



## BatchMintCollectionItems

> TransactionModel BatchMintCollectionItems(ctx, collectionId).XAuthorization(xAuthorization).XPassword(xPassword).BatchMintCollectionItemsRequest(batchMintCollectionItemsRequest).Execute()

Batch mint collection items



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    batchMintCollectionItemsRequest := *openapiclient.NewBatchMintCollectionItemsRequest([]float32{float32(123)}, []float32{float32(123)}) // BatchMintCollectionItemsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.BatchMintCollectionItems(context.Background(), collectionId).XAuthorization(xAuthorization).XPassword(xPassword).BatchMintCollectionItemsRequest(batchMintCollectionItemsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.BatchMintCollectionItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchMintCollectionItems`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.BatchMintCollectionItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchMintCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **batchMintCollectionItemsRequest** | [**BatchMintCollectionItemsRequest**](BatchMintCollectionItemsRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchTransferCollectionItems

> TransactionModel BatchTransferCollectionItems(ctx, collectionId).XAuthorization(xAuthorization).XPassword(xPassword).BatchTransferCollectionItemsRequest(batchTransferCollectionItemsRequest).Execute()

Batch transfer collection items



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    batchTransferCollectionItemsRequest := *openapiclient.NewBatchTransferCollectionItemsRequest([]float32{float32(12)}, []float32{float32(1)}) // BatchTransferCollectionItemsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.BatchTransferCollectionItems(context.Background(), collectionId).XAuthorization(xAuthorization).XPassword(xPassword).BatchTransferCollectionItemsRequest(batchTransferCollectionItemsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.BatchTransferCollectionItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchTransferCollectionItems`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.BatchTransferCollectionItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchTransferCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **batchTransferCollectionItemsRequest** | [**BatchTransferCollectionItemsRequest**](BatchTransferCollectionItemsRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BurnCollectionItem

> TransactionModel BurnCollectionItem(ctx, collectionId, collectionItemId).XAuthorization(xAuthorization).XPassword(xPassword).BurnCollectionItemRequest(burnCollectionItemRequest).Execute()

Burn collection item



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    collectionItemId := float32(8.14) // float32 | Any item id for the collection. Zero, or a positive integer.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    burnCollectionItemRequest := *openapiclient.NewBurnCollectionItemRequest(float32(123)) // BurnCollectionItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.BurnCollectionItem(context.Background(), collectionId, collectionItemId).XAuthorization(xAuthorization).XPassword(xPassword).BurnCollectionItemRequest(burnCollectionItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.BurnCollectionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BurnCollectionItem`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.BurnCollectionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 
**collectionItemId** | **float32** | Any item id for the collection. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBurnCollectionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **burnCollectionItemRequest** | [**BurnCollectionItemRequest**](BurnCollectionItemRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollection

> CreateCollection200Response CreateCollection(ctx).XAuthorization(xAuthorization).XPassword(xPassword).CreateCollectionRequest(createCollectionRequest).Execute()

Create collection



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
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    createCollectionRequest := *openapiclient.NewCreateCollectionRequest("SELECT ONE") // CreateCollectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.CreateCollection(context.Background()).XAuthorization(xAuthorization).XPassword(xPassword).CreateCollectionRequest(createCollectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.CreateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCollection`: CreateCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.CreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **createCollectionRequest** | [**CreateCollectionRequest**](CreateCollectionRequest.md) |  | 

### Return type

[**CreateCollection200Response**](CreateCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollectionItem

> TransactionModel CreateCollectionItem(ctx, collectionId).XAuthorization(xAuthorization).XPassword(xPassword).CreateCollectionItemRequest(createCollectionItemRequest).Execute()

Create collection item



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    createCollectionItemRequest := *openapiclient.NewCreateCollectionItemRequest(float32(1337), "Fire Dagger", "Description_example") // CreateCollectionItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.CreateCollectionItem(context.Background(), collectionId).XAuthorization(xAuthorization).XPassword(xPassword).CreateCollectionItemRequest(createCollectionItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.CreateCollectionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCollectionItem`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.CreateCollectionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **createCollectionItemRequest** | [**CreateCollectionItemRequest**](CreateCollectionItemRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionApproval

> bool GetCollectionApproval(ctx, collectionId).OperatorAddress(operatorAddress).Address(address).WalletId(walletId).Execute()

Get collection approval



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    operatorAddress := "0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D" // string | A valid EVM based address. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`.
    address := "0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D" // string | A valid EVM based address. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`. (optional)
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionApproval(context.Background(), collectionId).OperatorAddress(operatorAddress).Address(address).WalletId(walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionApproval`: bool
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operatorAddress** | **string** | A valid EVM based address. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | 
 **address** | **string** | A valid EVM based address. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | 
 **walletId** | **string** | Any wallet id within the MetaFab ecosystem. | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItem

> map[string]interface{} GetCollectionItem(ctx, collectionId, collectionItemId).Execute()

Get collection item



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    collectionItemId := float32(8.14) // float32 | Any item id for the collection. Zero, or a positive integer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionItem(context.Background(), collectionId, collectionItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionItem`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 
**collectionItemId** | **float32** | Any item id for the collection. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItemBalance

> float32 GetCollectionItemBalance(ctx, collectionId, collectionItemId).Address(address).WalletId(walletId).Execute()

Get collection item balance



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    collectionItemId := float32(8.14) // float32 | Any item id for the collection. Zero, or a positive integer.
    address := "0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D" // string | A valid EVM based address. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`. (optional)
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionItemBalance(context.Background(), collectionId, collectionItemId).Address(address).WalletId(walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionItemBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionItemBalance`: float32
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionItemBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 
**collectionItemId** | **float32** | Any item id for the collection. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **address** | **string** | A valid EVM based address. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | 
 **walletId** | **string** | Any wallet id within the MetaFab ecosystem. | 

### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItemBalances

> map[string]float32 GetCollectionItemBalances(ctx, collectionId).Address(address).WalletId(walletId).Execute()

Get collection item balances



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    address := "0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D" // string | A valid EVM based address. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`. (optional)
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionItemBalances(context.Background(), collectionId).Address(address).WalletId(walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionItemBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionItemBalances`: map[string]float32
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionItemBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **string** | A valid EVM based address. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | 
 **walletId** | **string** | Any wallet id within the MetaFab ecosystem. | 

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


## GetCollectionItemSupplies

> map[string]float32 GetCollectionItemSupplies(ctx, collectionId).Execute()

Get collection item supplies



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionItemSupplies(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionItemSupplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionItemSupplies`: map[string]float32
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionItemSupplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemSuppliesRequest struct via the builder pattern


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


## GetCollectionItemSupply

> float32 GetCollectionItemSupply(ctx, collectionId, collectionItemId).Address(address).WalletId(walletId).Execute()

Get collection item supply



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    collectionItemId := float32(8.14) // float32 | Any item id for the collection. Zero, or a positive integer.
    address := "0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D" // string | A valid EVM based address. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`. (optional)
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionItemSupply(context.Background(), collectionId, collectionItemId).Address(address).WalletId(walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionItemSupply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionItemSupply`: float32
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionItemSupply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 
**collectionItemId** | **float32** | Any item id for the collection. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemSupplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **address** | **string** | A valid EVM based address. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | 
 **walletId** | **string** | Any wallet id within the MetaFab ecosystem. | 

### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItemTimelock

> float32 GetCollectionItemTimelock(ctx, collectionId, collectionItemId).Execute()

Get collection item timelock



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    collectionItemId := float32(8.14) // float32 | Any item id for the collection. Zero, or a positive integer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionItemTimelock(context.Background(), collectionId, collectionItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionItemTimelock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionItemTimelock`: float32
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionItemTimelock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 
**collectionItemId** | **float32** | Any item id for the collection. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemTimelockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItems

> []map[string]interface{} GetCollectionItems(ctx, collectionId).Execute()

Get collection items



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionItems(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionItems`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionRole

> bool GetCollectionRole(ctx, collectionId).Role(role).Address(address).WalletId(walletId).Execute()

Get collection role



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    role := "minter" // string | A valid MetaFab role or bytes string representing a role, such as `0xc9eb32e43bf5ecbceacf00b32281dfc5d6d700a0db676ea26ccf938a385ac3b7`
    address := "0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D" // string | A valid EVM based address. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`. (optional)
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollectionRole(context.Background(), collectionId).Role(role).Address(address).WalletId(walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollectionRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionRole`: bool
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollectionRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | A valid MetaFab role or bytes string representing a role, such as &#x60;0xc9eb32e43bf5ecbceacf00b32281dfc5d6d700a0db676ea26ccf938a385ac3b7&#x60; | 
 **address** | **string** | A valid EVM based address. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | 
 **walletId** | **string** | Any wallet id within the MetaFab ecosystem. | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollections

> []GetCollections200ResponseInner GetCollections(ctx).XGameKey(xGameKey).Execute()

Get collections



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
    xGameKey := "game_pk_4SOqpDi8pQdnQgfCOBW29qR8YmwOhxVPz5iHoMgUEJLDdPXgwLuHqZf8ewo2GajZ" // string | The `publishedKey` of a specific game. This can be shared or included in game clients, websites, etc.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GetCollections(context.Background()).XGameKey(xGameKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GetCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollections`: []GetCollections200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GetCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGameKey** | **string** | The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc. | 

### Return type

[**[]GetCollections200ResponseInner**](GetCollections200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantCollectionRole

> TransactionModel GrantCollectionRole(ctx, collectionId).XAuthorization(xAuthorization).XPassword(xPassword).GrantCollectionRoleRequest(grantCollectionRoleRequest).Execute()

Grant collection role



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    grantCollectionRoleRequest := *openapiclient.NewGrantCollectionRoleRequest("Role_example") // GrantCollectionRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.GrantCollectionRole(context.Background(), collectionId).XAuthorization(xAuthorization).XPassword(xPassword).GrantCollectionRoleRequest(grantCollectionRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.GrantCollectionRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrantCollectionRole`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.GrantCollectionRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantCollectionRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **grantCollectionRoleRequest** | [**GrantCollectionRoleRequest**](GrantCollectionRoleRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MintCollectionItem

> TransactionModel MintCollectionItem(ctx, collectionId, collectionItemId).XAuthorization(xAuthorization).XPassword(xPassword).MintCollectionItemRequest(mintCollectionItemRequest).Execute()

Mint collection item



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    collectionItemId := float32(8.14) // float32 | Any item id for the collection. Zero, or a positive integer.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    mintCollectionItemRequest := *openapiclient.NewMintCollectionItemRequest(float32(123)) // MintCollectionItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.MintCollectionItem(context.Background(), collectionId, collectionItemId).XAuthorization(xAuthorization).XPassword(xPassword).MintCollectionItemRequest(mintCollectionItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.MintCollectionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MintCollectionItem`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.MintCollectionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 
**collectionItemId** | **float32** | Any item id for the collection. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMintCollectionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **mintCollectionItemRequest** | [**MintCollectionItemRequest**](MintCollectionItemRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeCollectionRole

> TransactionModel RevokeCollectionRole(ctx, collectionId).XAuthorization(xAuthorization).XPassword(xPassword).RevokeCollectionRoleRequest(revokeCollectionRoleRequest).Execute()

Revoke collection role



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    revokeCollectionRoleRequest := *openapiclient.NewRevokeCollectionRoleRequest("Role_example") // RevokeCollectionRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.RevokeCollectionRole(context.Background(), collectionId).XAuthorization(xAuthorization).XPassword(xPassword).RevokeCollectionRoleRequest(revokeCollectionRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.RevokeCollectionRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeCollectionRole`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.RevokeCollectionRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCollectionRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **revokeCollectionRoleRequest** | [**RevokeCollectionRoleRequest**](RevokeCollectionRoleRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCollectionApproval

> TransactionModel SetCollectionApproval(ctx, collectionId).XAuthorization(xAuthorization).XPassword(xPassword).SetCollectionApprovalRequest(setCollectionApprovalRequest).Execute()

Set collection approval



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    setCollectionApprovalRequest := *openapiclient.NewSetCollectionApprovalRequest(false) // SetCollectionApprovalRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.SetCollectionApproval(context.Background(), collectionId).XAuthorization(xAuthorization).XPassword(xPassword).SetCollectionApprovalRequest(setCollectionApprovalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.SetCollectionApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCollectionApproval`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.SetCollectionApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCollectionApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **setCollectionApprovalRequest** | [**SetCollectionApprovalRequest**](SetCollectionApprovalRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCollectionItemTimelock

> TransactionModel SetCollectionItemTimelock(ctx, collectionId, collectionItemId).XAuthorization(xAuthorization).XPassword(xPassword).SetCollectionItemTimelockRequest(setCollectionItemTimelockRequest).Execute()

Set collection item timelock



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    collectionItemId := float32(8.14) // float32 | Any item id for the collection. Zero, or a positive integer.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    setCollectionItemTimelockRequest := *openapiclient.NewSetCollectionItemTimelockRequest(float32(1665786026)) // SetCollectionItemTimelockRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.SetCollectionItemTimelock(context.Background(), collectionId, collectionItemId).XAuthorization(xAuthorization).XPassword(xPassword).SetCollectionItemTimelockRequest(setCollectionItemTimelockRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.SetCollectionItemTimelock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCollectionItemTimelock`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.SetCollectionItemTimelock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 
**collectionItemId** | **float32** | Any item id for the collection. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCollectionItemTimelockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **setCollectionItemTimelockRequest** | [**SetCollectionItemTimelockRequest**](SetCollectionItemTimelockRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferCollectionItem

> TransactionModel TransferCollectionItem(ctx, collectionId, collectionItemId).XAuthorization(xAuthorization).XPassword(xPassword).TransferCollectionItemRequest(transferCollectionItemRequest).Execute()

Transfer collection item



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
    collectionId := "collectionId_example" // string | Any collection id within the MetaFab ecosystem.
    collectionItemId := float32(8.14) // float32 | Any item id for the collection. Zero, or a positive integer.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    transferCollectionItemRequest := *openapiclient.NewTransferCollectionItemRequest(float32(123)) // TransferCollectionItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ItemsApi.TransferCollectionItem(context.Background(), collectionId, collectionItemId).XAuthorization(xAuthorization).XPassword(xPassword).TransferCollectionItemRequest(transferCollectionItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemsApi.TransferCollectionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferCollectionItem`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ItemsApi.TransferCollectionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Any collection id within the MetaFab ecosystem. | 
**collectionItemId** | **float32** | Any item id for the collection. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferCollectionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **transferCollectionItemRequest** | [**TransferCollectionItemRequest**](TransferCollectionItemRequest.md) |  | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

