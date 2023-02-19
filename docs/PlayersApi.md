# \PlayersApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthPlayer**](PlayersApi.md#AuthPlayer) | **Get** /v1/players/auth | Authenticate player
[**CreatePlayer**](PlayersApi.md#CreatePlayer) | **Post** /v1/players | Create player
[**GetPlayer**](PlayersApi.md#GetPlayer) | **Get** /v1/players/{playerId} | Get player
[**GetPlayerData**](PlayersApi.md#GetPlayerData) | **Get** /v1/players/{playerId}/data | Get player data
[**GetPlayers**](PlayersApi.md#GetPlayers) | **Get** /v1/players | Get players
[**RemovePlayerConnectedWallet**](PlayersApi.md#RemovePlayerConnectedWallet) | **Delete** /v1/players/{playerId}/wallets/{playerWalletId} | Remove player connected wallet
[**SetPlayerConnectedWallet**](PlayersApi.md#SetPlayerConnectedWallet) | **Post** /v1/players/{playerId}/wallets | Set player connected wallet
[**SetPlayerData**](PlayersApi.md#SetPlayerData) | **Post** /v1/players/{playerId}/data | Set player data
[**UpdatePlayer**](PlayersApi.md#UpdatePlayer) | **Patch** /v1/players/{playerId} | Update player



## AuthPlayer

> AuthPlayer200Response AuthPlayer(ctx).XGameKey(xGameKey).Execute()

Authenticate player



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
    resp, r, err := apiClient.PlayersApi.AuthPlayer(context.Background()).XGameKey(xGameKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.AuthPlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthPlayer`: AuthPlayer200Response
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.AuthPlayer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthPlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGameKey** | **string** | The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc. | 

### Return type

[**AuthPlayer200Response**](AuthPlayer200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlayer

> AuthPlayer200Response CreatePlayer(ctx).XGameKey(xGameKey).CreatePlayerRequest(createPlayerRequest).Execute()

Create player



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
    createPlayerRequest := *openapiclient.NewCreatePlayerRequest("Username_example", "aReallyStrongPassword123") // CreatePlayerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayersApi.CreatePlayer(context.Background()).XGameKey(xGameKey).CreatePlayerRequest(createPlayerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.CreatePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlayer`: AuthPlayer200Response
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.CreatePlayer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGameKey** | **string** | The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc. | 
 **createPlayerRequest** | [**CreatePlayerRequest**](CreatePlayerRequest.md) |  | 

### Return type

[**AuthPlayer200Response**](AuthPlayer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayer

> PublicPlayer GetPlayer(ctx, playerId).Execute()

Get player



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
    playerId := "playerId_example" // string | Any player id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayersApi.GetPlayer(context.Background(), playerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.GetPlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayer`: PublicPlayer
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.GetPlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | Any player id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicPlayer**](PublicPlayer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayerData

> GetPlayerData200Response GetPlayerData(ctx, playerId).Execute()

Get player data



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
    playerId := "playerId_example" // string | Any player id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayersApi.GetPlayerData(context.Background(), playerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.GetPlayerData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayerData`: GetPlayerData200Response
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.GetPlayerData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | Any player id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPlayerData200Response**](GetPlayerData200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayers

> []PublicPlayer GetPlayers(ctx).XAuthorization(xAuthorization).Execute()

Get players



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayersApi.GetPlayers(context.Background()).XAuthorization(xAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.GetPlayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayers`: []PublicPlayer
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.GetPlayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 

### Return type

[**[]PublicPlayer**](PublicPlayer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePlayerConnectedWallet

> RemovePlayerConnectedWallet(ctx, playerId, playerWalletId).RemovePlayerConnectedWalletRequest(removePlayerConnectedWalletRequest).Execute()

Remove player connected wallet



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
    playerId := "playerId_example" // string | Any player id within the MetaFab platform.
    playerWalletId := "playerWalletId_example" // string | Any player wallet id within the MetaFab platform.
    removePlayerConnectedWalletRequest := *openapiclient.NewRemovePlayerConnectedWalletRequest("Address_example", int32(123), "Signature_example", "SELECT ONE") // RemovePlayerConnectedWalletRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayersApi.RemovePlayerConnectedWallet(context.Background(), playerId, playerWalletId).RemovePlayerConnectedWalletRequest(removePlayerConnectedWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.RemovePlayerConnectedWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | Any player id within the MetaFab platform. | 
**playerWalletId** | **string** | Any player wallet id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePlayerConnectedWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removePlayerConnectedWalletRequest** | [**RemovePlayerConnectedWalletRequest**](RemovePlayerConnectedWalletRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPlayerConnectedWallet

> SetPlayerConnectedWallet200Response SetPlayerConnectedWallet(ctx, playerId).XAuthorization(xAuthorization).SetPlayerConnectedWalletRequest(setPlayerConnectedWalletRequest).Execute()

Set player connected wallet



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
    playerId := "playerId_example" // string | The player id of the authenticating player.
    xAuthorization := "player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `accessToken` of the authenticating player.
    setPlayerConnectedWalletRequest := *openapiclient.NewSetPlayerConnectedWalletRequest("Address_example", int32(123), "Signature_example", "SELECT ONE") // SetPlayerConnectedWalletRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayersApi.SetPlayerConnectedWallet(context.Background(), playerId).XAuthorization(xAuthorization).SetPlayerConnectedWalletRequest(setPlayerConnectedWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.SetPlayerConnectedWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPlayerConnectedWallet`: SetPlayerConnectedWallet200Response
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.SetPlayerConnectedWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | The player id of the authenticating player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPlayerConnectedWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;accessToken&#x60; of the authenticating player. | 
 **setPlayerConnectedWalletRequest** | [**SetPlayerConnectedWalletRequest**](SetPlayerConnectedWalletRequest.md) |  | 

### Return type

[**SetPlayerConnectedWallet200Response**](SetPlayerConnectedWallet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPlayerData

> GetPlayerData200Response SetPlayerData(ctx, playerId).XAuthorization(xAuthorization).SetPlayerDataRequest(setPlayerDataRequest).Execute()

Set player data



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
    playerId := "playerId_example" // string | Any player id within the MetaFab platform.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    setPlayerDataRequest := *openapiclient.NewSetPlayerDataRequest() // SetPlayerDataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayersApi.SetPlayerData(context.Background(), playerId).XAuthorization(xAuthorization).SetPlayerDataRequest(setPlayerDataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.SetPlayerData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPlayerData`: GetPlayerData200Response
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.SetPlayerData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | Any player id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPlayerDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **setPlayerDataRequest** | [**SetPlayerDataRequest**](SetPlayerDataRequest.md) |  | 

### Return type

[**GetPlayerData200Response**](GetPlayerData200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlayer

> UpdatePlayer200Response UpdatePlayer(ctx, playerId).XAuthorization(xAuthorization).UpdatePlayerRequest(updatePlayerRequest).Execute()

Update player



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
    playerId := "playerId_example" // string | The player id of the authenticating player.
    xAuthorization := "player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `accessToken` of the authenticating player.
    updatePlayerRequest := *openapiclient.NewUpdatePlayerRequest() // UpdatePlayerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayersApi.UpdatePlayer(context.Background(), playerId).XAuthorization(xAuthorization).UpdatePlayerRequest(updatePlayerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.UpdatePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlayer`: UpdatePlayer200Response
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.UpdatePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | The player id of the authenticating player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;accessToken&#x60; of the authenticating player. | 
 **updatePlayerRequest** | [**UpdatePlayerRequest**](UpdatePlayerRequest.md) |  | 

### Return type

[**UpdatePlayer200Response**](UpdatePlayer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

