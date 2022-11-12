# CreateContractRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address of the existing contract. | 
**Abi** | **string** | JSON of the abi. | 
**Chain** | **string** | The blockchain you want to deploy this currency on. Support for new blockchains are added over time. | 

## Methods

### NewCreateContractRequest

`func NewCreateContractRequest(address string, abi string, chain string, ) *CreateContractRequest`

NewCreateContractRequest instantiates a new CreateContractRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContractRequestWithDefaults

`func NewCreateContractRequestWithDefaults() *CreateContractRequest`

NewCreateContractRequestWithDefaults instantiates a new CreateContractRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateContractRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateContractRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateContractRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAbi

`func (o *CreateContractRequest) GetAbi() string`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *CreateContractRequest) GetAbiOk() (*string, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *CreateContractRequest) SetAbi(v string)`

SetAbi sets Abi field to given value.


### GetChain

`func (o *CreateContractRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateContractRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CreateContractRequest) SetChain(v string)`

SetChain sets Chain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


