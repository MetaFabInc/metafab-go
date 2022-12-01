# CreateShopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | **string** | The blockchain you want to deploy this shop on. Support for new blockchains are added over time. | 

## Methods

### NewCreateShopRequest

`func NewCreateShopRequest(chain string, ) *CreateShopRequest`

NewCreateShopRequest instantiates a new CreateShopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShopRequestWithDefaults

`func NewCreateShopRequestWithDefaults() *CreateShopRequest`

NewCreateShopRequestWithDefaults instantiates a new CreateShopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *CreateShopRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateShopRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CreateShopRequest) SetChain(v string)`

SetChain sets Chain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


