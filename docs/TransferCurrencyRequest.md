# TransferCurrencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | A valid EVM based address to transfer currency to. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | [optional] 
**WalletId** | Pointer to **string** | Any wallet id within the MetaFab ecosystem to transfer currency to. | [optional] 
**Amount** | **float32** | The amount of currency to transfer. | 
**Reference** | Pointer to **float32** | An optional uint256 number to reference the transfer. Commonly used to identify transfers intended to pay for game items or services. | [optional] 

## Methods

### NewTransferCurrencyRequest

`func NewTransferCurrencyRequest(amount float32, ) *TransferCurrencyRequest`

NewTransferCurrencyRequest instantiates a new TransferCurrencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferCurrencyRequestWithDefaults

`func NewTransferCurrencyRequestWithDefaults() *TransferCurrencyRequest`

NewTransferCurrencyRequestWithDefaults instantiates a new TransferCurrencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransferCurrencyRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransferCurrencyRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransferCurrencyRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransferCurrencyRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetWalletId

`func (o *TransferCurrencyRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransferCurrencyRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransferCurrencyRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *TransferCurrencyRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetAmount

`func (o *TransferCurrencyRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferCurrencyRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferCurrencyRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetReference

`func (o *TransferCurrencyRequest) GetReference() float32`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TransferCurrencyRequest) GetReferenceOk() (*float32, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TransferCurrencyRequest) SetReference(v float32)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TransferCurrencyRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


