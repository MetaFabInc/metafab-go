# \CurrenciesApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchTransferCurrency**](CurrenciesApi.md#BatchTransferCurrency) | **Post** /v1/currencies/{currencyId}/batchTransfers | Batch transfer currency
[**BurnCurrency**](CurrenciesApi.md#BurnCurrency) | **Post** /v1/currencies/{currencyId}/burns | Burn currency
[**CreateCurrency**](CurrenciesApi.md#CreateCurrency) | **Post** /v1/currencies | Create currency
[**GetCurrencies**](CurrenciesApi.md#GetCurrencies) | **Get** /v1/currencies | Get currencies
[**GetCurrencyBalance**](CurrenciesApi.md#GetCurrencyBalance) | **Get** /v1/currencies/{currencyId}/balances | Get currency balance
[**GetCurrencyFees**](CurrenciesApi.md#GetCurrencyFees) | **Get** /v1/currencies/{currencyId}/fees | Get currency fees
[**GetCurrencyRole**](CurrenciesApi.md#GetCurrencyRole) | **Get** /v1/currencies/{currencyId}/roles | Get currency role
[**GrantCurrencyRole**](CurrenciesApi.md#GrantCurrencyRole) | **Post** /v1/currencies/{currencyId}/roles | Grant currency role
[**MintCurrency**](CurrenciesApi.md#MintCurrency) | **Post** /v1/currencies/{currencyId}/mints | Mint currency
[**RevokeCurrencyRole**](CurrenciesApi.md#RevokeCurrencyRole) | **Delete** /v1/currencies/{currencyId}/roles | Revoke currency role
[**SetCurrencyFees**](CurrenciesApi.md#SetCurrencyFees) | **Post** /v1/currencies/{currencyId}/fees | Set currency fees
[**TransferCurrency**](CurrenciesApi.md#TransferCurrency) | **Post** /v1/currencies/{currencyId}/transfers | Transfer currency



## BatchTransferCurrency

> TransactionModel BatchTransferCurrency(ctx, currencyId).XAuthorization(xAuthorization).XPassword(xPassword).BatchTransferCurrencyRequest(batchTransferCurrencyRequest).Execute()

Batch transfer currency



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    batchTransferCurrencyRequest := *openapiclient.NewBatchTransferCurrencyRequest([]float32{float32(10)}) // BatchTransferCurrencyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.BatchTransferCurrency(context.Background(), currencyId).XAuthorization(xAuthorization).XPassword(xPassword).BatchTransferCurrencyRequest(batchTransferCurrencyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.BatchTransferCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchTransferCurrency`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.BatchTransferCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchTransferCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **batchTransferCurrencyRequest** | [**BatchTransferCurrencyRequest**](BatchTransferCurrencyRequest.md) |  | 

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


## BurnCurrency

> TransactionModel BurnCurrency(ctx, currencyId).XAuthorization(xAuthorization).XPassword(xPassword).BurnCurrencyRequest(burnCurrencyRequest).Execute()

Burn currency



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    burnCurrencyRequest := *openapiclient.NewBurnCurrencyRequest(float32(133.7)) // BurnCurrencyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.BurnCurrency(context.Background(), currencyId).XAuthorization(xAuthorization).XPassword(xPassword).BurnCurrencyRequest(burnCurrencyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.BurnCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BurnCurrency`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.BurnCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBurnCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **burnCurrencyRequest** | [**BurnCurrencyRequest**](BurnCurrencyRequest.md) |  | 

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


## CreateCurrency

> CreateCurrency200Response CreateCurrency(ctx).XAuthorization(xAuthorization).XPassword(xPassword).CreateCurrencyRequest(createCurrencyRequest).Execute()

Create currency



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
    createCurrencyRequest := *openapiclient.NewCreateCurrencyRequest("Bright Gems", "BGEM", float32(15000.5), "SELECT ONE") // CreateCurrencyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.CreateCurrency(context.Background()).XAuthorization(xAuthorization).XPassword(xPassword).CreateCurrencyRequest(createCurrencyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.CreateCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCurrency`: CreateCurrency200Response
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.CreateCurrency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **createCurrencyRequest** | [**CreateCurrencyRequest**](CreateCurrencyRequest.md) |  | 

### Return type

[**CreateCurrency200Response**](CreateCurrency200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrencies

> []GetCurrencies200ResponseInner GetCurrencies(ctx).XGameKey(xGameKey).Execute()

Get currencies



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
    resp, r, err := apiClient.CurrenciesApi.GetCurrencies(context.Background()).XGameKey(xGameKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetCurrencies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrencies`: []GetCurrencies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetCurrencies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGameKey** | **string** | The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc. | 

### Return type

[**[]GetCurrencies200ResponseInner**](GetCurrencies200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrencyBalance

> float32 GetCurrencyBalance(ctx, currencyId).Address(address).WalletId(walletId).Execute()

Get currency balance



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    address := "0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D" // string | A valid EVM based address. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`. (optional)
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.GetCurrencyBalance(context.Background(), currencyId).Address(address).WalletId(walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetCurrencyBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrencyBalance`: float32
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetCurrencyBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrencyBalanceRequest struct via the builder pattern


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


## GetCurrencyFees

> GetCurrencyFees200Response GetCurrencyFees(ctx, currencyId).Execute()

Get currency fees



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.GetCurrencyFees(context.Background(), currencyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetCurrencyFees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrencyFees`: GetCurrencyFees200Response
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetCurrencyFees`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrencyFeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCurrencyFees200Response**](GetCurrencyFees200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrencyRole

> bool GetCurrencyRole(ctx, currencyId).Role(role).Address(address).WalletId(walletId).Execute()

Get currency role



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    role := "minter" // string | A valid MetaFab role or bytes string representing a role, such as `0xc9eb32e43bf5ecbceacf00b32281dfc5d6d700a0db676ea26ccf938a385ac3b7`
    address := "0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D" // string | A valid EVM based address. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`. (optional)
    walletId := "walletId_example" // string | Any wallet id within the MetaFab ecosystem. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.GetCurrencyRole(context.Background(), currencyId).Role(role).Address(address).WalletId(walletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetCurrencyRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrencyRole`: bool
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetCurrencyRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrencyRoleRequest struct via the builder pattern


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


## GrantCurrencyRole

> TransactionModel GrantCurrencyRole(ctx, currencyId).XAuthorization(xAuthorization).XPassword(xPassword).GrantCurrencyRoleRequest(grantCurrencyRoleRequest).Execute()

Grant currency role



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    grantCurrencyRoleRequest := *openapiclient.NewGrantCurrencyRoleRequest("Role_example") // GrantCurrencyRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.GrantCurrencyRole(context.Background(), currencyId).XAuthorization(xAuthorization).XPassword(xPassword).GrantCurrencyRoleRequest(grantCurrencyRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GrantCurrencyRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrantCurrencyRole`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GrantCurrencyRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantCurrencyRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **grantCurrencyRoleRequest** | [**GrantCurrencyRoleRequest**](GrantCurrencyRoleRequest.md) |  | 

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


## MintCurrency

> TransactionModel MintCurrency(ctx, currencyId).XAuthorization(xAuthorization).XPassword(xPassword).MintCurrencyRequest(mintCurrencyRequest).Execute()

Mint currency



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    mintCurrencyRequest := *openapiclient.NewMintCurrencyRequest(float32(133.7)) // MintCurrencyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.MintCurrency(context.Background(), currencyId).XAuthorization(xAuthorization).XPassword(xPassword).MintCurrencyRequest(mintCurrencyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.MintCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MintCurrency`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.MintCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMintCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **mintCurrencyRequest** | [**MintCurrencyRequest**](MintCurrencyRequest.md) |  | 

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


## RevokeCurrencyRole

> TransactionModel RevokeCurrencyRole(ctx, currencyId).XAuthorization(xAuthorization).XPassword(xPassword).RevokeCollectionRoleRequest(revokeCollectionRoleRequest).Execute()

Revoke currency role



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    revokeCollectionRoleRequest := *openapiclient.NewRevokeCollectionRoleRequest("Role_example") // RevokeCollectionRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.RevokeCurrencyRole(context.Background(), currencyId).XAuthorization(xAuthorization).XPassword(xPassword).RevokeCollectionRoleRequest(revokeCollectionRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.RevokeCurrencyRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeCurrencyRole`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.RevokeCurrencyRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCurrencyRoleRequest struct via the builder pattern


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


## SetCurrencyFees

> TransactionModel SetCurrencyFees(ctx, currencyId).XAuthorization(xAuthorization).XPassword(xPassword).SetCurrencyFeesRequest(setCurrencyFeesRequest).Execute()

Set currency fees



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    setCurrencyFeesRequest := *openapiclient.NewSetCurrencyFeesRequest("RecipientAddress_example", float32(123), float32(123), float32(123)) // SetCurrencyFeesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.SetCurrencyFees(context.Background(), currencyId).XAuthorization(xAuthorization).XPassword(xPassword).SetCurrencyFeesRequest(setCurrencyFeesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.SetCurrencyFees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCurrencyFees`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.SetCurrencyFees`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCurrencyFeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **setCurrencyFeesRequest** | [**SetCurrencyFeesRequest**](SetCurrencyFeesRequest.md) |  | 

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


## TransferCurrency

> TransactionModel TransferCurrency(ctx, currencyId).XAuthorization(xAuthorization).XPassword(xPassword).TransferCurrencyRequest(transferCurrencyRequest).Execute()

Transfer currency



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
    currencyId := "currencyId_example" // string | Any currency id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    transferCurrencyRequest := *openapiclient.NewTransferCurrencyRequest(float32(133.7)) // TransferCurrencyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.TransferCurrency(context.Background(), currencyId).XAuthorization(xAuthorization).XPassword(xPassword).TransferCurrencyRequest(transferCurrencyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.TransferCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferCurrency`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.TransferCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Any currency id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **transferCurrencyRequest** | [**TransferCurrencyRequest**](TransferCurrencyRequest.md) |  | 

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

