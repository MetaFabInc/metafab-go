# \LootboxesApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLootboxManager**](LootboxesApi.md#CreateLootboxManager) | **Post** /v1/lootboxManagers | Create lootbox manager
[**GetLootboxManagerLootbox**](LootboxesApi.md#GetLootboxManagerLootbox) | **Get** /v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId} | Get lootbox manager lootbox
[**GetLootboxManagerLootboxes**](LootboxesApi.md#GetLootboxManagerLootboxes) | **Get** /v1/lootboxManagers/{lootboxManagerId}/lootboxes | Get lootbox manager lootboxes
[**GetLootboxManagers**](LootboxesApi.md#GetLootboxManagers) | **Get** /v1/lootboxManagers | Get lootbox managers
[**OpenLootboxManagerLootbox**](LootboxesApi.md#OpenLootboxManagerLootbox) | **Post** /v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId}/opens | Open lootbox manager lootbox
[**RemoveLootboxManagerLootbox**](LootboxesApi.md#RemoveLootboxManagerLootbox) | **Delete** /v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId} | Remove lootbox manager lootbox
[**SetLootboxManagerLootbox**](LootboxesApi.md#SetLootboxManagerLootbox) | **Post** /v1/lootboxManagers/{lootboxManagerId}/lootboxes | Set lootbox manager lootbox



## CreateLootboxManager

> CreateLootboxManager200Response CreateLootboxManager(ctx).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).CreateLootboxManagerRequest(createLootboxManagerRequest).Execute()

Create lootbox manager



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
    xWalletDecryptKey := "AXNP8MKb+5SbBtHWrZu5KHh5/BomXY/dMRG/BDUn7a4=" // string | The `walletDecryptKey` of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    createLootboxManagerRequest := *openapiclient.NewCreateLootboxManagerRequest("SELECT ONE") // CreateLootboxManagerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LootboxesApi.CreateLootboxManager(context.Background()).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).CreateLootboxManagerRequest(createLootboxManagerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LootboxesApi.CreateLootboxManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLootboxManager`: CreateLootboxManager200Response
    fmt.Fprintf(os.Stdout, "Response from `LootboxesApi.CreateLootboxManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLootboxManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xWalletDecryptKey** | **string** | The &#x60;walletDecryptKey&#x60; of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **createLootboxManagerRequest** | [**CreateLootboxManagerRequest**](CreateLootboxManagerRequest.md) |  | 

### Return type

[**CreateLootboxManager200Response**](CreateLootboxManager200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLootboxManagerLootbox

> LootboxManagerLootbox GetLootboxManagerLootbox(ctx, lootboxManagerId, lootboxManagerLootboxId).Execute()

Get lootbox manager lootbox



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
    lootboxManagerId := "lootboxManagerId_example" // string | Any lootbox manager id within the MetaFab platform.
    lootboxManagerLootboxId := "lootboxManagerLootboxId_example" // string | Any lootbox manager lootbox id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LootboxesApi.GetLootboxManagerLootbox(context.Background(), lootboxManagerId, lootboxManagerLootboxId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LootboxesApi.GetLootboxManagerLootbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLootboxManagerLootbox`: LootboxManagerLootbox
    fmt.Fprintf(os.Stdout, "Response from `LootboxesApi.GetLootboxManagerLootbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lootboxManagerId** | **string** | Any lootbox manager id within the MetaFab platform. | 
**lootboxManagerLootboxId** | **string** | Any lootbox manager lootbox id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLootboxManagerLootboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LootboxManagerLootbox**](LootboxManagerLootbox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLootboxManagerLootboxes

> []LootboxManagerLootbox GetLootboxManagerLootboxes(ctx, lootboxManagerId).Execute()

Get lootbox manager lootboxes



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
    lootboxManagerId := "lootboxManagerId_example" // string | Any lootbox manager id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LootboxesApi.GetLootboxManagerLootboxes(context.Background(), lootboxManagerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LootboxesApi.GetLootboxManagerLootboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLootboxManagerLootboxes`: []LootboxManagerLootbox
    fmt.Fprintf(os.Stdout, "Response from `LootboxesApi.GetLootboxManagerLootboxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lootboxManagerId** | **string** | Any lootbox manager id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLootboxManagerLootboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LootboxManagerLootbox**](LootboxManagerLootbox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLootboxManagers

> []GetLootboxManagers200ResponseInner GetLootboxManagers(ctx).XGameKey(xGameKey).Execute()

Get lootbox managers



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
    resp, r, err := apiClient.LootboxesApi.GetLootboxManagers(context.Background()).XGameKey(xGameKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LootboxesApi.GetLootboxManagers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLootboxManagers`: []GetLootboxManagers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LootboxesApi.GetLootboxManagers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLootboxManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGameKey** | **string** | The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc. | 

### Return type

[**[]GetLootboxManagers200ResponseInner**](GetLootboxManagers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenLootboxManagerLootbox

> []TransactionModel OpenLootboxManagerLootbox(ctx, lootboxManagerId, lootboxManagerLootboxId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).Execute()

Open lootbox manager lootbox



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
    lootboxManagerId := "lootboxManagerId_example" // string | Any lootbox manager id within the MetaFab platform.
    lootboxManagerLootboxId := "lootboxManagerLootboxId_example" // string | Any lootbox manager lootbox id within the MetaFab platform.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xWalletDecryptKey := "AXNP8MKb+5SbBtHWrZu5KHh5/BomXY/dMRG/BDUn7a4=" // string | The `walletDecryptKey` of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LootboxesApi.OpenLootboxManagerLootbox(context.Background(), lootboxManagerId, lootboxManagerLootboxId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LootboxesApi.OpenLootboxManagerLootbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenLootboxManagerLootbox`: []TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `LootboxesApi.OpenLootboxManagerLootbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lootboxManagerId** | **string** | Any lootbox manager id within the MetaFab platform. | 
**lootboxManagerLootboxId** | **string** | Any lootbox manager lootbox id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenLootboxManagerLootboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xWalletDecryptKey** | **string** | The &#x60;walletDecryptKey&#x60; of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 

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


## RemoveLootboxManagerLootbox

> TransactionModel RemoveLootboxManagerLootbox(ctx, lootboxManagerId, lootboxManagerLootboxId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).Execute()

Remove lootbox manager lootbox



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
    lootboxManagerId := "lootboxManagerId_example" // string | Any lootbox manager id within the MetaFab platform.
    lootboxManagerLootboxId := "lootboxManagerLootboxId_example" // string | Any lootbox manager lootbox id within the MetaFab platform.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xWalletDecryptKey := "AXNP8MKb+5SbBtHWrZu5KHh5/BomXY/dMRG/BDUn7a4=" // string | The `walletDecryptKey` of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LootboxesApi.RemoveLootboxManagerLootbox(context.Background(), lootboxManagerId, lootboxManagerLootboxId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LootboxesApi.RemoveLootboxManagerLootbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveLootboxManagerLootbox`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `LootboxesApi.RemoveLootboxManagerLootbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lootboxManagerId** | **string** | Any lootbox manager id within the MetaFab platform. | 
**lootboxManagerLootboxId** | **string** | Any lootbox manager lootbox id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLootboxManagerLootboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xWalletDecryptKey** | **string** | The &#x60;walletDecryptKey&#x60; of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 

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


## SetLootboxManagerLootbox

> TransactionModel SetLootboxManagerLootbox(ctx, lootboxManagerId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).SetLootboxManagerLootboxRequest(setLootboxManagerLootboxRequest).Execute()

Set lootbox manager lootbox



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
    lootboxManagerId := "lootboxManagerId_example" // string | Any lootbox manager id within the MetaFab platform.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xWalletDecryptKey := "AXNP8MKb+5SbBtHWrZu5KHh5/BomXY/dMRG/BDUn7a4=" // string | The `walletDecryptKey` of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    setLootboxManagerLootboxRequest := *openapiclient.NewSetLootboxManagerLootboxRequest(int32(1337)) // SetLootboxManagerLootboxRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LootboxesApi.SetLootboxManagerLootbox(context.Background(), lootboxManagerId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).SetLootboxManagerLootboxRequest(setLootboxManagerLootboxRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LootboxesApi.SetLootboxManagerLootbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLootboxManagerLootbox`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `LootboxesApi.SetLootboxManagerLootbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lootboxManagerId** | **string** | Any lootbox manager id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLootboxManagerLootboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xWalletDecryptKey** | **string** | The &#x60;walletDecryptKey&#x60; of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **setLootboxManagerLootboxRequest** | [**SetLootboxManagerLootboxRequest**](SetLootboxManagerLootboxRequest.md) |  | 

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

