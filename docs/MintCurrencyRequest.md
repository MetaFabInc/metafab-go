# MintCurrencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | The amount of currency to create (mint). | 
**Address** | Pointer to **string** | A valid EVM based address to create (mint) currency for. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | [optional] 
**WalletId** | Pointer to **string** | Any wallet id within the MetaFab ecosystem to create (mint) currency for. | [optional] 

## Methods

### NewMintCurrencyRequest

`func NewMintCurrencyRequest(amount float32, ) *MintCurrencyRequest`

NewMintCurrencyRequest instantiates a new MintCurrencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintCurrencyRequestWithDefaults

`func NewMintCurrencyRequestWithDefaults() *MintCurrencyRequest`

NewMintCurrencyRequestWithDefaults instantiates a new MintCurrencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *MintCurrencyRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MintCurrencyRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MintCurrencyRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAddress

`func (o *MintCurrencyRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MintCurrencyRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MintCurrencyRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MintCurrencyRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetWalletId

`func (o *MintCurrencyRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MintCurrencyRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MintCurrencyRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *MintCurrencyRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


