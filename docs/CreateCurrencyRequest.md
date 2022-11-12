# CreateCurrencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this currency. This can be anything, such as &#x60;Bright Gems&#x60;, &#x60;Gold&#x60;, etc. | 
**Symbol** | **string** | The shorthand symbol to represent this currency. This can be anything, such as &#x60;BGEM&#x60;, &#x60;GLD&#x60;, etc. | 
**SupplyCap** | **float32** | The maximum amount of this currency that can ever exist. Use &#x60;0&#x60; if you do not want this currency to have a maximum supply. | 
**Chain** | **string** | The blockchain you want to deploy this currency on. Support for new blockchains are added over time. | 

## Methods

### NewCreateCurrencyRequest

`func NewCreateCurrencyRequest(name string, symbol string, supplyCap float32, chain string, ) *CreateCurrencyRequest`

NewCreateCurrencyRequest instantiates a new CreateCurrencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCurrencyRequestWithDefaults

`func NewCreateCurrencyRequestWithDefaults() *CreateCurrencyRequest`

NewCreateCurrencyRequestWithDefaults instantiates a new CreateCurrencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCurrencyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCurrencyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCurrencyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *CreateCurrencyRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CreateCurrencyRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CreateCurrencyRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetSupplyCap

`func (o *CreateCurrencyRequest) GetSupplyCap() float32`

GetSupplyCap returns the SupplyCap field if non-nil, zero value otherwise.

### GetSupplyCapOk

`func (o *CreateCurrencyRequest) GetSupplyCapOk() (*float32, bool)`

GetSupplyCapOk returns a tuple with the SupplyCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplyCap

`func (o *CreateCurrencyRequest) SetSupplyCap(v float32)`

SetSupplyCap sets SupplyCap field to given value.


### GetChain

`func (o *CreateCurrencyRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateCurrencyRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CreateCurrencyRequest) SetChain(v string)`

SetChain sets Chain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


