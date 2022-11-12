# \ExchangesApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExchange**](ExchangesApi.md#CreateExchange) | **Post** /v1/exchanges | Create exchange
[**GetExchangeOffer**](ExchangesApi.md#GetExchangeOffer) | **Get** /v1/exchanges/{exchangeId}/items/{exchangeOfferId} | Get exchange offer
[**GetExchangeOffers**](ExchangesApi.md#GetExchangeOffers) | **Get** /v1/exchanges/{exchangeId}/offers | Get exchange offers
[**GetExchanges**](ExchangesApi.md#GetExchanges) | **Get** /v1/exchanges | Get exchanges
[**RemoveExchangeOffer**](ExchangesApi.md#RemoveExchangeOffer) | **Delete** /v1/exchanges/{exchangeId}/offers/{exchangeOfferId} | Remove exchange offer
[**SetExchangeOffer**](ExchangesApi.md#SetExchangeOffer) | **Post** /v1/exchanges/{exchangeId}/offers | Set exchange offer
[**UseExchangeOffer**](ExchangesApi.md#UseExchangeOffer) | **Post** /v1/exchanges/{exchangeId}/offers/{exchangeOfferId}/uses | Use exchange offer
[**WithdrawFromExchange**](ExchangesApi.md#WithdrawFromExchange) | **Post** /v1/exchanges/{exchangeId}/withdrawals | Withdraw from exchange



## CreateExchange

> CreateExchange200Response CreateExchange(ctx).XAuthorization(xAuthorization).XPassword(xPassword).CreateExchangeRequest(createExchangeRequest).Execute()

Create exchange



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
    createExchangeRequest := *openapiclient.NewCreateExchangeRequest("SELECT ONE") // CreateExchangeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.CreateExchange(context.Background()).XAuthorization(xAuthorization).XPassword(xPassword).CreateExchangeRequest(createExchangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.CreateExchange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExchange`: CreateExchange200Response
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.CreateExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **createExchangeRequest** | [**CreateExchangeRequest**](CreateExchangeRequest.md) |  | 

### Return type

[**CreateExchange200Response**](CreateExchange200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeOffer

> ExchangeOffer GetExchangeOffer(ctx, exchangeId, exchangeOfferId).Execute()

Get exchange offer



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
    exchangeId := "exchangeId_example" // string | Any exchange id within the MetaFab ecosystem.
    exchangeOfferId := "exchangeOfferId_example" // string | Any offer id for the exchange. Zero, or a positive integer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.GetExchangeOffer(context.Background(), exchangeId, exchangeOfferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchangeOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeOffer`: ExchangeOffer
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchangeOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeId** | **string** | Any exchange id within the MetaFab ecosystem. | 
**exchangeOfferId** | **string** | Any offer id for the exchange. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExchangeOffer**](ExchangeOffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeOffers

> []ExchangeOffer GetExchangeOffers(ctx, exchangeId).Execute()

Get exchange offers



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
    exchangeId := "exchangeId_example" // string | Any exchange id within the MetaFab ecosystem.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.GetExchangeOffers(context.Background(), exchangeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchangeOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeOffers`: []ExchangeOffer
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchangeOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeId** | **string** | Any exchange id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ExchangeOffer**](ExchangeOffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchanges

> []GetExchanges200ResponseInner GetExchanges(ctx).XGameKey(xGameKey).Execute()

Get exchanges



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
    resp, r, err := apiClient.ExchangesApi.GetExchanges(context.Background()).XGameKey(xGameKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchanges`: []GetExchanges200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGameKey** | **string** | The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc. | 

### Return type

[**[]GetExchanges200ResponseInner**](GetExchanges200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExchangeOffer

> TransactionModel RemoveExchangeOffer(ctx, exchangeId, exchangeOfferId).XAuthorization(xAuthorization).XPassword(xPassword).Execute()

Remove exchange offer



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
    exchangeId := "exchangeId_example" // string | Any exchange id within the MetaFab ecosystem.
    exchangeOfferId := "exchangeOfferId_example" // string | Any offer id for the exchange. Zero, or a positive integer.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.RemoveExchangeOffer(context.Background(), exchangeId, exchangeOfferId).XAuthorization(xAuthorization).XPassword(xPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.RemoveExchangeOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExchangeOffer`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.RemoveExchangeOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeId** | **string** | Any exchange id within the MetaFab ecosystem. | 
**exchangeOfferId** | **string** | Any offer id for the exchange. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExchangeOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 

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


## SetExchangeOffer

> TransactionModel SetExchangeOffer(ctx, exchangeId).XAuthorization(xAuthorization).XPassword(xPassword).SetExchangeOfferRequest(setExchangeOfferRequest).Execute()

Set exchange offer



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
    exchangeId := "exchangeId_example" // string | Any exchange id within the MetaFab ecosystem.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    setExchangeOfferRequest := *openapiclient.NewSetExchangeOfferRequest(float32(1337)) // SetExchangeOfferRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.SetExchangeOffer(context.Background(), exchangeId).XAuthorization(xAuthorization).XPassword(xPassword).SetExchangeOfferRequest(setExchangeOfferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.SetExchangeOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetExchangeOffer`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.SetExchangeOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeId** | **string** | Any exchange id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetExchangeOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **setExchangeOfferRequest** | [**SetExchangeOfferRequest**](SetExchangeOfferRequest.md) |  | 

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


## UseExchangeOffer

> TransactionModel UseExchangeOffer(ctx, exchangeId, exchangeOfferId).XAuthorization(xAuthorization).XPassword(xPassword).Execute()

Use exchange offer



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
    exchangeId := "exchangeId_example" // string | Any exchange id within the MetaFab ecosystem.
    exchangeOfferId := "exchangeOfferId_example" // string | Any offer id for the exchange. Zero, or a positive integer.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.UseExchangeOffer(context.Background(), exchangeId, exchangeOfferId).XAuthorization(xAuthorization).XPassword(xPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.UseExchangeOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UseExchangeOffer`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.UseExchangeOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeId** | **string** | Any exchange id within the MetaFab ecosystem. | 
**exchangeOfferId** | **string** | Any offer id for the exchange. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUseExchangeOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 

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


## WithdrawFromExchange

> TransactionModel WithdrawFromExchange(ctx, exchangeId).XAuthorization(xAuthorization).XPassword(xPassword).WithdrawFromExchangeRequest(withdrawFromExchangeRequest).Execute()

Withdraw from exchange



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
    exchangeId := "exchangeId_example" // string | Any exchange id within the MetaFab ecosystem.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    withdrawFromExchangeRequest := *openapiclient.NewWithdrawFromExchangeRequest() // WithdrawFromExchangeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.WithdrawFromExchange(context.Background(), exchangeId).XAuthorization(xAuthorization).XPassword(xPassword).WithdrawFromExchangeRequest(withdrawFromExchangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.WithdrawFromExchange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WithdrawFromExchange`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.WithdrawFromExchange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeId** | **string** | Any exchange id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawFromExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **withdrawFromExchangeRequest** | [**WithdrawFromExchangeRequest**](WithdrawFromExchangeRequest.md) |  | 

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

