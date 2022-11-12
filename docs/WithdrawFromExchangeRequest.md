# WithdrawFromExchangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | A valid EVM based address to withdraw to. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | [optional] 
**WalletId** | Pointer to **string** | Any wallet id within the MetaFab ecosystem to withdraw to. | [optional] 
**CurrencyAddress** | Pointer to **string** | The address of the currency (ERC20) token to withdraw from the exchange. If no currencyAddress or currencyId, and no collectionAddress or collectionId are provided, the native token held by the exchange will be withdrawn. | [optional] 
**CurrencyId** | Pointer to **string** | A valid MetaFab currency id that represents the currency token to withdraw from the exchange. &#x60;currencyAddress&#x60; or &#x60;currencyId&#x60; can be provided when withdrawing currency. | [optional] 
**CollectionAddress** | Pointer to **string** | The address of the collection (ERC1155) for the items to withdraw from the exchange. If no currencyAddress and no collectionAddress is provided, the native token held by the exchange will be withdrawn. | [optional] 
**CollectionId** | Pointer to **string** | A valid MetaFab collection id that represents the collection for the items to withdraw from the exchange. &#x60;collectionAddress&#x60; or &#x60;collectionId&#x60; can be provided when withdrawing items. | [optional] 
**ItemIds** | Pointer to **[]float32** | The specific itemIds of the provided collection to withdraw from the exchange. | [optional] 

## Methods

### NewWithdrawFromExchangeRequest

`func NewWithdrawFromExchangeRequest() *WithdrawFromExchangeRequest`

NewWithdrawFromExchangeRequest instantiates a new WithdrawFromExchangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawFromExchangeRequestWithDefaults

`func NewWithdrawFromExchangeRequestWithDefaults() *WithdrawFromExchangeRequest`

NewWithdrawFromExchangeRequestWithDefaults instantiates a new WithdrawFromExchangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *WithdrawFromExchangeRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WithdrawFromExchangeRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WithdrawFromExchangeRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WithdrawFromExchangeRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetWalletId

`func (o *WithdrawFromExchangeRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WithdrawFromExchangeRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WithdrawFromExchangeRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *WithdrawFromExchangeRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetCurrencyAddress

`func (o *WithdrawFromExchangeRequest) GetCurrencyAddress() string`

GetCurrencyAddress returns the CurrencyAddress field if non-nil, zero value otherwise.

### GetCurrencyAddressOk

`func (o *WithdrawFromExchangeRequest) GetCurrencyAddressOk() (*string, bool)`

GetCurrencyAddressOk returns a tuple with the CurrencyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyAddress

`func (o *WithdrawFromExchangeRequest) SetCurrencyAddress(v string)`

SetCurrencyAddress sets CurrencyAddress field to given value.

### HasCurrencyAddress

`func (o *WithdrawFromExchangeRequest) HasCurrencyAddress() bool`

HasCurrencyAddress returns a boolean if a field has been set.

### GetCurrencyId

`func (o *WithdrawFromExchangeRequest) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WithdrawFromExchangeRequest) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WithdrawFromExchangeRequest) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *WithdrawFromExchangeRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCollectionAddress

`func (o *WithdrawFromExchangeRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *WithdrawFromExchangeRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *WithdrawFromExchangeRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.

### HasCollectionAddress

`func (o *WithdrawFromExchangeRequest) HasCollectionAddress() bool`

HasCollectionAddress returns a boolean if a field has been set.

### GetCollectionId

`func (o *WithdrawFromExchangeRequest) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *WithdrawFromExchangeRequest) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *WithdrawFromExchangeRequest) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *WithdrawFromExchangeRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetItemIds

`func (o *WithdrawFromExchangeRequest) GetItemIds() []float32`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *WithdrawFromExchangeRequest) GetItemIdsOk() (*[]float32, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *WithdrawFromExchangeRequest) SetItemIds(v []float32)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *WithdrawFromExchangeRequest) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


