# \GamesApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthGame**](GamesApi.md#AuthGame) | **Get** /v1/games/auth | Authenticate game
[**CreateGame**](GamesApi.md#CreateGame) | **Post** /v1/games | Create game
[**GetGame**](GamesApi.md#GetGame) | **Get** /v1/games/{gameId} | Get game
[**UpdateGame**](GamesApi.md#UpdateGame) | **Patch** /v1/games/{gameId} | Update game



## AuthGame

> AuthGame200Response AuthGame(ctx).Execute()

Authenticate game



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.AuthGame(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.AuthGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthGame`: AuthGame200Response
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.AuthGame`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthGameRequest struct via the builder pattern


### Return type

[**AuthGame200Response**](AuthGame200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGame

> AuthGame200Response CreateGame(ctx).CreateGameRequest(createGameRequest).Execute()

Create game



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
    createGameRequest := *openapiclient.NewCreateGameRequest("NFT Worlds", "dev@nftworlds.com", "aReallyStrongPassword123!") // CreateGameRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.CreateGame(context.Background()).CreateGameRequest(createGameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.CreateGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGame`: AuthGame200Response
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.CreateGame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGameRequest** | [**CreateGameRequest**](CreateGameRequest.md) |  | 

### Return type

[**AuthGame200Response**](AuthGame200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGame

> PublicGame GetGame(ctx, gameId).Execute()

Get game



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
    gameId := "gameId_example" // string | Any game id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.GetGame(context.Background(), gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.GetGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGame`: PublicGame
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.GetGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | Any game id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicGame**](PublicGame.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGame

> UpdateGame200Response UpdateGame(ctx, gameId).XAuthorization(xAuthorization).UpdateGameRequest(updateGameRequest).Execute()

Update game



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
    gameId := "gameId_example" // string | The game id of the authenticating game.
    xAuthorization := "game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating game.
    updateGameRequest := *openapiclient.NewUpdateGameRequest() // UpdateGameRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GamesApi.UpdateGame(context.Background(), gameId).XAuthorization(xAuthorization).UpdateGameRequest(updateGameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GamesApi.UpdateGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGame`: UpdateGame200Response
    fmt.Fprintf(os.Stdout, "Response from `GamesApi.UpdateGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The game id of the authenticating game. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **updateGameRequest** | [**UpdateGameRequest**](UpdateGameRequest.md) |  | 

### Return type

[**UpdateGame200Response**](UpdateGame200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

