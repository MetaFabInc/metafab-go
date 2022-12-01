# \ShopsApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShop**](ShopsApi.md#CreateShop) | **Post** /v1/shops | Create shop
[**GetShopOffer**](ShopsApi.md#GetShopOffer) | **Get** /v1/shops/{shopId}/items/{shopOfferId} | Get shop offer
[**GetShopOffers**](ShopsApi.md#GetShopOffers) | **Get** /v1/shops/{shopId}/offers | Get shop offers
[**GetShops**](ShopsApi.md#GetShops) | **Get** /v1/shops | Get shops
[**RemoveShopOffer**](ShopsApi.md#RemoveShopOffer) | **Delete** /v1/shops/{shopId}/offers/{shopOfferId} | Remove shop offer
[**SetShopOffer**](ShopsApi.md#SetShopOffer) | **Post** /v1/shops/{shopId}/offers | Set shop offer
[**UseShopOffer**](ShopsApi.md#UseShopOffer) | **Post** /v1/shops/{shopId}/offers/{shopOfferId}/uses | Use shop offer
[**WithdrawFromShop**](ShopsApi.md#WithdrawFromShop) | **Post** /v1/shops/{shopId}/withdrawals | Withdraw from shop



## CreateShop

> CreateShop200Response CreateShop(ctx).XAuthorization(xAuthorization).XPassword(xPassword).CreateShopRequest(createShopRequest).Execute()

Create shop



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
    createShopRequest := *openapiclient.NewCreateShopRequest("SELECT ONE") // CreateShopRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopsApi.CreateShop(context.Background()).XAuthorization(xAuthorization).XPassword(xPassword).CreateShopRequest(createShopRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopsApi.CreateShop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShop`: CreateShop200Response
    fmt.Fprintf(os.Stdout, "Response from `ShopsApi.CreateShop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **createShopRequest** | [**CreateShopRequest**](CreateShopRequest.md) |  | 

### Return type

[**CreateShop200Response**](CreateShop200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopOffer

> ShopOffer GetShopOffer(ctx, shopId, shopOfferId).Execute()

Get shop offer



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
    shopId := "shopId_example" // string | Any shop id within the MetaFab ecosystem.
    shopOfferId := "shopOfferId_example" // string | Any offer id for the shop. Zero, or a positive integer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopsApi.GetShopOffer(context.Background(), shopId, shopOfferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopsApi.GetShopOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShopOffer`: ShopOffer
    fmt.Fprintf(os.Stdout, "Response from `ShopsApi.GetShopOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **string** | Any shop id within the MetaFab ecosystem. | 
**shopOfferId** | **string** | Any offer id for the shop. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ShopOffer**](ShopOffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShopOffers

> []ShopOffer GetShopOffers(ctx, shopId).Execute()

Get shop offers



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
    shopId := "shopId_example" // string | Any shop id within the MetaFab ecosystem.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopsApi.GetShopOffers(context.Background(), shopId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopsApi.GetShopOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShopOffers`: []ShopOffer
    fmt.Fprintf(os.Stdout, "Response from `ShopsApi.GetShopOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **string** | Any shop id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ShopOffer**](ShopOffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShops

> []GetShops200ResponseInner GetShops(ctx).XGameKey(xGameKey).Execute()

Get shops



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
    resp, r, err := apiClient.ShopsApi.GetShops(context.Background()).XGameKey(xGameKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopsApi.GetShops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShops`: []GetShops200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ShopsApi.GetShops`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGameKey** | **string** | The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc. | 

### Return type

[**[]GetShops200ResponseInner**](GetShops200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveShopOffer

> TransactionModel RemoveShopOffer(ctx, shopId, shopOfferId).XAuthorization(xAuthorization).XPassword(xPassword).Execute()

Remove shop offer



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
    shopId := "shopId_example" // string | Any shop id within the MetaFab ecosystem.
    shopOfferId := "shopOfferId_example" // string | Any offer id for the shop. Zero, or a positive integer.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopsApi.RemoveShopOffer(context.Background(), shopId, shopOfferId).XAuthorization(xAuthorization).XPassword(xPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopsApi.RemoveShopOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveShopOffer`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ShopsApi.RemoveShopOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **string** | Any shop id within the MetaFab ecosystem. | 
**shopOfferId** | **string** | Any offer id for the shop. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveShopOfferRequest struct via the builder pattern


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


## SetShopOffer

> TransactionModel SetShopOffer(ctx, shopId).XAuthorization(xAuthorization).XPassword(xPassword).SetShopOfferRequest(setShopOfferRequest).Execute()

Set shop offer



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
    shopId := "shopId_example" // string | Any shop id within the MetaFab ecosystem.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    setShopOfferRequest := *openapiclient.NewSetShopOfferRequest(int32(1337)) // SetShopOfferRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopsApi.SetShopOffer(context.Background(), shopId).XAuthorization(xAuthorization).XPassword(xPassword).SetShopOfferRequest(setShopOfferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopsApi.SetShopOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetShopOffer`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ShopsApi.SetShopOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **string** | Any shop id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetShopOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **setShopOfferRequest** | [**SetShopOfferRequest**](SetShopOfferRequest.md) |  | 

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


## UseShopOffer

> TransactionModel UseShopOffer(ctx, shopId, shopOfferId).XAuthorization(xAuthorization).XPassword(xPassword).Execute()

Use shop offer



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
    shopId := "shopId_example" // string | Any shop id within the MetaFab ecosystem.
    shopOfferId := "shopOfferId_example" // string | Any offer id for the shop. Zero, or a positive integer.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopsApi.UseShopOffer(context.Background(), shopId, shopOfferId).XAuthorization(xAuthorization).XPassword(xPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopsApi.UseShopOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UseShopOffer`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ShopsApi.UseShopOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **string** | Any shop id within the MetaFab ecosystem. | 
**shopOfferId** | **string** | Any offer id for the shop. Zero, or a positive integer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUseShopOfferRequest struct via the builder pattern


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


## WithdrawFromShop

> TransactionModel WithdrawFromShop(ctx, shopId).XAuthorization(xAuthorization).XPassword(xPassword).WithdrawFromShopRequest(withdrawFromShopRequest).Execute()

Withdraw from shop



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
    shopId := "shopId_example" // string | Any shop id within the MetaFab ecosystem.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    xPassword := "mySecurePassword" // string | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
    withdrawFromShopRequest := *openapiclient.NewWithdrawFromShopRequest() // WithdrawFromShopRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShopsApi.WithdrawFromShop(context.Background(), shopId).XAuthorization(xAuthorization).XPassword(xPassword).WithdrawFromShopRequest(withdrawFromShopRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShopsApi.WithdrawFromShop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WithdrawFromShop`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ShopsApi.WithdrawFromShop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shopId** | **string** | Any shop id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawFromShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **xPassword** | **string** | The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet. | 
 **withdrawFromShopRequest** | [**WithdrawFromShopRequest**](WithdrawFromShopRequest.md) |  | 

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

