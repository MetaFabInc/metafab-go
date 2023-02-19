# \EcosystemsApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveEcosystemGame**](EcosystemsApi.md#ApproveEcosystemGame) | **Post** /v1/ecosystems/{ecosystemId}/games | Approve ecosystem game
[**AuthEcosystem**](EcosystemsApi.md#AuthEcosystem) | **Get** /v1/ecosystems/auth | Authenticate ecosystem
[**AuthProfile**](EcosystemsApi.md#AuthProfile) | **Get** /v1/profiles/auth | Authenticate profile
[**AuthProfilePlayer**](EcosystemsApi.md#AuthProfilePlayer) | **Get** /v1/profiles/{profileId}/games/{gameId}/players/auth | Authenticate profile player
[**CreateEcosystem**](EcosystemsApi.md#CreateEcosystem) | **Post** /v1/ecosystems | Create ecosystem
[**CreateProfile**](EcosystemsApi.md#CreateProfile) | **Post** /v1/profiles | Create profile
[**CreateProfilePlayer**](EcosystemsApi.md#CreateProfilePlayer) | **Post** /v1/profiles/{profileId}/games/{gameId}/players | Create profile player
[**GetEcosystem**](EcosystemsApi.md#GetEcosystem) | **Get** /v1/ecosystems/{ecosystemId} | Get ecosystem
[**GetEcosystemGame**](EcosystemsApi.md#GetEcosystemGame) | **Get** /v1/ecosystems/{ecosystemId}/games/{gameId} | Get ecosystem game
[**GetEcosystemGames**](EcosystemsApi.md#GetEcosystemGames) | **Get** /v1/ecosystems/{ecosystemId}/games | Get ecosystem games
[**GetProfileGame**](EcosystemsApi.md#GetProfileGame) | **Get** /v1/profiles/{profileId}/games/{gameId} | Get profile game
[**GetProfileGames**](EcosystemsApi.md#GetProfileGames) | **Get** /v1/profiles/{profileId}/games | Get profile games
[**UnapproveEcosystemGame**](EcosystemsApi.md#UnapproveEcosystemGame) | **Delete** /v1/ecosystems/{ecosystemId}/games/{gameId} | Unapprove ecosystem game
[**UpdateEcosystem**](EcosystemsApi.md#UpdateEcosystem) | **Patch** /v1/ecosystems/{ecosystemId} | Update ecosystem
[**UpdateProfile**](EcosystemsApi.md#UpdateProfile) | **Patch** /v1/profiles/{profileId} | Update profile
[**UpdateProfilePlayer**](EcosystemsApi.md#UpdateProfilePlayer) | **Patch** /v1/profiles/{profileId}/games/{gameId}/players/{playerId} | Update profile player



## ApproveEcosystemGame

> ApproveEcosystemGame(ctx, ecosystemId).XAuthorization(xAuthorization).ApproveEcosystemGameRequest(approveEcosystemGameRequest).Execute()

Approve ecosystem game



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
    ecosystemId := "ecosystemId_example" // string | The ecosystem id of the authenticating ecosystem.
    xAuthorization := "ecosystem_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating ecosystem.
    approveEcosystemGameRequest := *openapiclient.NewApproveEcosystemGameRequest("GameId_example") // ApproveEcosystemGameRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.ApproveEcosystemGame(context.Background(), ecosystemId).XAuthorization(xAuthorization).ApproveEcosystemGameRequest(approveEcosystemGameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.ApproveEcosystemGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ecosystemId** | **string** | The ecosystem id of the authenticating ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveEcosystemGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating ecosystem. | 
 **approveEcosystemGameRequest** | [**ApproveEcosystemGameRequest**](ApproveEcosystemGameRequest.md) |  | 

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


## AuthEcosystem

> EcosystemModel AuthEcosystem(ctx).Execute()

Authenticate ecosystem



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
    resp, r, err := apiClient.EcosystemsApi.AuthEcosystem(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.AuthEcosystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthEcosystem`: EcosystemModel
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.AuthEcosystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthEcosystemRequest struct via the builder pattern


### Return type

[**EcosystemModel**](EcosystemModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthProfile

> AuthProfile200Response AuthProfile(ctx).XEcosystemKey(xEcosystemKey).Execute()

Authenticate profile



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
    xEcosystemKey := "ecosystem_pk_a5sFpDi8pQdnQgfCOBW29qR8YmwOhxVPz5iHoMgUEJLDdPXgwLuHqZf8ewo2GajZ" // string | The `publishedKey` of a specific ecosystem. This can be shared or included in clients, websites, etc.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.AuthProfile(context.Background()).XEcosystemKey(xEcosystemKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.AuthProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthProfile`: AuthProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.AuthProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEcosystemKey** | **string** | The &#x60;publishedKey&#x60; of a specific ecosystem. This can be shared or included in clients, websites, etc. | 

### Return type

[**AuthProfile200Response**](AuthProfile200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthProfilePlayer

> AuthPlayer200Response AuthProfilePlayer(ctx, profileId, gameId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).XUsername(xUsername).Execute()

Authenticate profile player



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
    profileId := "profileId_example" // string | The profile id of the authenticating profile.
    gameId := "gameId_example" // string | Any game id within the MetaFab platform.
    xAuthorization := "profile_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `accessToken` of the authenticating profile.
    xWalletDecryptKey := "AXNP8MKb+5SbBtHWrZu5KHh5/BomXY/dMRG/BDUn7a4=" // string | The `walletDecryptKey` of the authenticating profile. Required to decrypt and perform blockchain transactions with the profile wallet.
    xUsername := "arkdev" // string | The username of a player.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.AuthProfilePlayer(context.Background(), profileId, gameId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).XUsername(xUsername).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.AuthProfilePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthProfilePlayer`: AuthPlayer200Response
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.AuthProfilePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile id of the authenticating profile. | 
**gameId** | **string** | Any game id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthProfilePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;accessToken&#x60; of the authenticating profile. | 
 **xWalletDecryptKey** | **string** | The &#x60;walletDecryptKey&#x60; of the authenticating profile. Required to decrypt and perform blockchain transactions with the profile wallet. | 
 **xUsername** | **string** | The username of a player. | 

### Return type

[**AuthPlayer200Response**](AuthPlayer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEcosystem

> EcosystemModel CreateEcosystem(ctx).CreateEcosystemRequest(createEcosystemRequest).Execute()

Create ecosystem



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
    createEcosystemRequest := *openapiclient.NewCreateEcosystemRequest("NFT Worlds", "dev@nftworlds.com", "aReallyStrongPassword123!") // CreateEcosystemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.CreateEcosystem(context.Background()).CreateEcosystemRequest(createEcosystemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.CreateEcosystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEcosystem`: EcosystemModel
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.CreateEcosystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEcosystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEcosystemRequest** | [**CreateEcosystemRequest**](CreateEcosystemRequest.md) |  | 

### Return type

[**EcosystemModel**](EcosystemModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProfile

> AuthProfile200Response CreateProfile(ctx).XEcosystemKey(xEcosystemKey).CreateProfileRequest(createProfileRequest).Execute()

Create profile



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
    xEcosystemKey := "ecosystem_pk_a5sFpDi8pQdnQgfCOBW29qR8YmwOhxVPz5iHoMgUEJLDdPXgwLuHqZf8ewo2GajZ" // string | The `publishedKey` of a specific ecosystem. This can be shared or included in clients, websites, etc.
    createProfileRequest := *openapiclient.NewCreateProfileRequest("Username_example", "aReallyStrongPassword123") // CreateProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.CreateProfile(context.Background()).XEcosystemKey(xEcosystemKey).CreateProfileRequest(createProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.CreateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfile`: AuthProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.CreateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEcosystemKey** | **string** | The &#x60;publishedKey&#x60; of a specific ecosystem. This can be shared or included in clients, websites, etc. | 
 **createProfileRequest** | [**CreateProfileRequest**](CreateProfileRequest.md) |  | 

### Return type

[**AuthProfile200Response**](AuthProfile200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProfilePlayer

> AuthPlayer200Response CreateProfilePlayer(ctx, profileId, gameId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).CreateProfilePlayerRequest(createProfilePlayerRequest).Execute()

Create profile player



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
    profileId := "profileId_example" // string | The profile id of the authenticating profile.
    gameId := "gameId_example" // string | Any game id within the MetaFab platform.
    xAuthorization := "profile_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `accessToken` of the authenticating profile.
    xWalletDecryptKey := "AXNP8MKb+5SbBtHWrZu5KHh5/BomXY/dMRG/BDUn7a4=" // string | The `walletDecryptKey` of the authenticating profile. Required to decrypt and perform blockchain transactions with the profile wallet.
    createProfilePlayerRequest := *openapiclient.NewCreateProfilePlayerRequest("Username_example") // CreateProfilePlayerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.CreateProfilePlayer(context.Background(), profileId, gameId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).CreateProfilePlayerRequest(createProfilePlayerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.CreateProfilePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfilePlayer`: AuthPlayer200Response
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.CreateProfilePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile id of the authenticating profile. | 
**gameId** | **string** | Any game id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfilePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;accessToken&#x60; of the authenticating profile. | 
 **xWalletDecryptKey** | **string** | The &#x60;walletDecryptKey&#x60; of the authenticating profile. Required to decrypt and perform blockchain transactions with the profile wallet. | 
 **createProfilePlayerRequest** | [**CreateProfilePlayerRequest**](CreateProfilePlayerRequest.md) |  | 

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


## GetEcosystem

> PublicEcosystem GetEcosystem(ctx, ecosystemId).Execute()

Get ecosystem



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
    ecosystemId := "ecosystemId_example" // string | Any ecosystem id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.GetEcosystem(context.Background(), ecosystemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.GetEcosystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEcosystem`: PublicEcosystem
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.GetEcosystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ecosystemId** | **string** | Any ecosystem id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcosystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicEcosystem**](PublicEcosystem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcosystemGame

> PublicGame GetEcosystemGame(ctx, ecosystemId, gameId).Execute()

Get ecosystem game



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
    ecosystemId := "ecosystemId_example" // string | Any ecosystem id within the MetaFab platform.
    gameId := "gameId_example" // string | Any game id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.GetEcosystemGame(context.Background(), ecosystemId, gameId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.GetEcosystemGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEcosystemGame`: PublicGame
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.GetEcosystemGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ecosystemId** | **string** | Any ecosystem id within the MetaFab platform. | 
**gameId** | **string** | Any game id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcosystemGameRequest struct via the builder pattern


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


## GetEcosystemGames

> []PublicGame GetEcosystemGames(ctx, ecosystemId).Execute()

Get ecosystem games



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
    ecosystemId := "ecosystemId_example" // string | Any ecosystem id within the MetaFab platform.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.GetEcosystemGames(context.Background(), ecosystemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.GetEcosystemGames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEcosystemGames`: []PublicGame
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.GetEcosystemGames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ecosystemId** | **string** | Any ecosystem id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcosystemGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PublicGame**](PublicGame.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileGame

> GetProfileGames200ResponseInner GetProfileGame(ctx, profileId, gameId).XAuthorization(xAuthorization).Execute()

Get profile game



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
    profileId := "profileId_example" // string | The profile id of the authenticating profile.
    gameId := "gameId_example" // string | Any game id within the MetaFab platform.
    xAuthorization := "profile_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `accessToken` of the authenticating profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.GetProfileGame(context.Background(), profileId, gameId).XAuthorization(xAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.GetProfileGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileGame`: GetProfileGames200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.GetProfileGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile id of the authenticating profile. | 
**gameId** | **string** | Any game id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;accessToken&#x60; of the authenticating profile. | 

### Return type

[**GetProfileGames200ResponseInner**](GetProfileGames200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileGames

> []GetProfileGames200ResponseInner GetProfileGames(ctx, profileId).XAuthorization(xAuthorization).Execute()

Get profile games



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
    profileId := "profileId_example" // string | The profile id of the authenticating profile.
    xAuthorization := "profile_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `accessToken` of the authenticating profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.GetProfileGames(context.Background(), profileId).XAuthorization(xAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.GetProfileGames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileGames`: []GetProfileGames200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.GetProfileGames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile id of the authenticating profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;accessToken&#x60; of the authenticating profile. | 

### Return type

[**[]GetProfileGames200ResponseInner**](GetProfileGames200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnapproveEcosystemGame

> UnapproveEcosystemGame(ctx, ecosystemId, gameId).XAuthorization(xAuthorization).Execute()

Unapprove ecosystem game



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
    ecosystemId := "ecosystemId_example" // string | The ecosystem id of the authenticating ecosystem.
    gameId := "gameId_example" // string | Any game id within the MetaFab platform.
    xAuthorization := "ecosystem_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating ecosystem.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.UnapproveEcosystemGame(context.Background(), ecosystemId, gameId).XAuthorization(xAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.UnapproveEcosystemGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ecosystemId** | **string** | The ecosystem id of the authenticating ecosystem. | 
**gameId** | **string** | Any game id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnapproveEcosystemGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating ecosystem. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEcosystem

> EcosystemModel UpdateEcosystem(ctx, ecosystemId).XAuthorization(xAuthorization).UpdateEcosystemRequest(updateEcosystemRequest).Execute()

Update ecosystem



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
    ecosystemId := "ecosystemId_example" // string | The ecosystem id of the authenticating ecosystem.
    xAuthorization := "ecosystem_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `secretKey` of the authenticating ecosystem.
    updateEcosystemRequest := *openapiclient.NewUpdateEcosystemRequest() // UpdateEcosystemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.UpdateEcosystem(context.Background(), ecosystemId).XAuthorization(xAuthorization).UpdateEcosystemRequest(updateEcosystemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.UpdateEcosystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEcosystem`: EcosystemModel
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.UpdateEcosystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ecosystemId** | **string** | The ecosystem id of the authenticating ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEcosystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating ecosystem. | 
 **updateEcosystemRequest** | [**UpdateEcosystemRequest**](UpdateEcosystemRequest.md) |  | 

### Return type

[**EcosystemModel**](EcosystemModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> ProfileModel UpdateProfile(ctx, profileId).XAuthorization(xAuthorization).UpdateProfileRequest(updateProfileRequest).Execute()

Update profile



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
    profileId := "profileId_example" // string | The profile id of the authenticating profile.
    xAuthorization := "profile_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `accessToken` of the authenticating profile.
    updateProfileRequest := *openapiclient.NewUpdateProfileRequest() // UpdateProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.UpdateProfile(context.Background(), profileId).XAuthorization(xAuthorization).UpdateProfileRequest(updateProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.UpdateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfile`: ProfileModel
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile id of the authenticating profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;accessToken&#x60; of the authenticating profile. | 
 **updateProfileRequest** | [**UpdateProfileRequest**](UpdateProfileRequest.md) |  | 

### Return type

[**ProfileModel**](ProfileModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfilePlayer

> UpdateProfilePlayer200Response UpdateProfilePlayer(ctx, profileId, gameId, playerId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).UpdateProfilePlayerRequest(updateProfilePlayerRequest).Execute()

Update profile player



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
    profileId := "profileId_example" // string | The profile id of the authenticating profile.
    gameId := "gameId_example" // string | Any game id within the MetaFab platform.
    playerId := "playerId_example" // string | Any player id within the MetaFab platform.
    xAuthorization := "profile_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP" // string | The `accessToken` of the authenticating profile.
    xWalletDecryptKey := "AXNP8MKb+5SbBtHWrZu5KHh5/BomXY/dMRG/BDUn7a4=" // string | The `walletDecryptKey` of the authenticating profile. Required to decrypt and perform blockchain transactions with the profile wallet.
    updateProfilePlayerRequest := *openapiclient.NewUpdateProfilePlayerRequest() // UpdateProfilePlayerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EcosystemsApi.UpdateProfilePlayer(context.Background(), profileId, gameId, playerId).XAuthorization(xAuthorization).XWalletDecryptKey(xWalletDecryptKey).UpdateProfilePlayerRequest(updateProfilePlayerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EcosystemsApi.UpdateProfilePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfilePlayer`: UpdateProfilePlayer200Response
    fmt.Fprintf(os.Stdout, "Response from `EcosystemsApi.UpdateProfilePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile id of the authenticating profile. | 
**gameId** | **string** | Any game id within the MetaFab platform. | 
**playerId** | **string** | Any player id within the MetaFab platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfilePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAuthorization** | **string** | The &#x60;accessToken&#x60; of the authenticating profile. | 
 **xWalletDecryptKey** | **string** | The &#x60;walletDecryptKey&#x60; of the authenticating profile. Required to decrypt and perform blockchain transactions with the profile wallet. | 
 **updateProfilePlayerRequest** | [**UpdateProfilePlayerRequest**](UpdateProfilePlayerRequest.md) |  | 

### Return type

[**UpdateProfilePlayer200Response**](UpdateProfilePlayer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

