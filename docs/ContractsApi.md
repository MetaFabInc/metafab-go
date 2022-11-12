# \ContractsApi

All URIs are relative to *https://api.trymetafab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContract**](ContractsApi.md#CreateContract) | **Post** /v1/contracts | Create custom contract
[**GetContracts**](ContractsApi.md#GetContracts) | **Get** /v1/contracts | Get contracts
[**ReadContract**](ContractsApi.md#ReadContract) | **Get** /v1/contracts/{contractId}/reads | Read contract data
[**WriteContract**](ContractsApi.md#WriteContract) | **Post** /v1/contracts/{contractId}/writes | Write contract data



## CreateContract

> ContractModel CreateContract(ctx).XAuthorization(xAuthorization).CreateContractRequest(createContractRequest).Execute()

Create custom contract



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
    createContractRequest := *openapiclient.NewCreateContractRequest("Address_example", "Abi_example", "SELECT ONE") // CreateContractRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.CreateContract(context.Background()).XAuthorization(xAuthorization).CreateContractRequest(createContractRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.CreateContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContract`: ContractModel
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.CreateContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of the authenticating game. | 
 **createContractRequest** | [**CreateContractRequest**](CreateContractRequest.md) |  | 

### Return type

[**ContractModel**](ContractModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContracts

> []ContractModel GetContracts(ctx).XGameKey(xGameKey).Execute()

Get contracts



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
    resp, r, err := apiClient.ContractsApi.GetContracts(context.Background()).XGameKey(xGameKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContracts`: []ContractModel
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGameKey** | **string** | The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc. | 

### Return type

[**[]ContractModel**](ContractModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadContract

> interface{} ReadContract(ctx, contractId).Func_(func_).Args(args).Execute()

Read contract data



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
    contractId := "contractId_example" // string | Any contract id within the MetaFab ecosystem.
    func_ := "func__example" // string | A contract function name. This can be any valid function from the the ABI of the contract you are interacting with. For example, `balanceOf`.
    args := "123,"Hello",false" // string | A comma seperated list of basic data type arguments. This is optional and only necessary if the function being invoked requires arguments per the contract ABI. For example, `123,\"Hello\",false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.ReadContract(context.Background(), contractId).Func_(func_).Args(args).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ReadContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadContract`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.ReadContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | Any contract id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **func_** | **string** | A contract function name. This can be any valid function from the the ABI of the contract you are interacting with. For example, &#x60;balanceOf&#x60;. | 
 **args** | **string** | A comma seperated list of basic data type arguments. This is optional and only necessary if the function being invoked requires arguments per the contract ABI. For example, &#x60;123,\&quot;Hello\&quot;,false&#x60;. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteContract

> TransactionModel WriteContract(ctx, contractId).XAuthorization(xAuthorization).XPassword(xPassword).WriteContractRequest(writeContractRequest).Execute()

Write contract data



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
    contractId := "contractId_example" // string | Any contract id within the MetaFab ecosystem.
    xAuthorization := "["game_sk_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP","player_at_02z4Mv3c85Ig0gNowY9Dq0N2kjb1xwzr27ArLE0669RrRI6dLf822iPO26K1p1FP"]" // string | The `secretKey` of a specific game or the `accessToken` of a specific player.
    xPassword := "mySecurePassword" // string | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
    writeContractRequest := *openapiclient.NewWriteContractRequest("Func_example") // WriteContractRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.WriteContract(context.Background(), contractId).XAuthorization(xAuthorization).XPassword(xPassword).WriteContractRequest(writeContractRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.WriteContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WriteContract`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.WriteContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | Any contract id within the MetaFab ecosystem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWriteContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthorization** | **string** | The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player. | 
 **xPassword** | **string** | The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet. | 
 **writeContractRequest** | [**WriteContractRequest**](WriteContractRequest.md) |  | 

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

