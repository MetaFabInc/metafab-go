# BatchTransferCurrencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | An array of valid EVM based addresses to transfer currency to. | [optional] 
**WalletIds** | Pointer to **[]string** | An array of wallet ids within the MetaFab ecosystem to transfer currency to. | [optional] 
**Amounts** | **[]float32** | An array of currency amounts to transfer. Ordering corresponds to the ordering of provided &#x60;addresses&#x60; and/or &#x60;walletIds&#x60;. If both &#x60;addresses&#x60; and &#x60;walletIds&#x60; are provided, &#x60;addresses&#x60; are first in the order. | 
**References** | Pointer to **[]float32** | An optional array of uint256 numbers to reference each transfer in the batch. Ordering corresponds to the ordering of provided &#x60;addresses&#x60; or &#x60;walletIds&#x60;. If both &#x60;addresses&#x60; and &#x60;walletIds&#x60; are provided, &#x60;addresses&#x60; are first in the order. | [optional] 

## Methods

### NewBatchTransferCurrencyRequest

`func NewBatchTransferCurrencyRequest(amounts []float32, ) *BatchTransferCurrencyRequest`

NewBatchTransferCurrencyRequest instantiates a new BatchTransferCurrencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchTransferCurrencyRequestWithDefaults

`func NewBatchTransferCurrencyRequestWithDefaults() *BatchTransferCurrencyRequest`

NewBatchTransferCurrencyRequestWithDefaults instantiates a new BatchTransferCurrencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *BatchTransferCurrencyRequest) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BatchTransferCurrencyRequest) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BatchTransferCurrencyRequest) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BatchTransferCurrencyRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetWalletIds

`func (o *BatchTransferCurrencyRequest) GetWalletIds() []string`

GetWalletIds returns the WalletIds field if non-nil, zero value otherwise.

### GetWalletIdsOk

`func (o *BatchTransferCurrencyRequest) GetWalletIdsOk() (*[]string, bool)`

GetWalletIdsOk returns a tuple with the WalletIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletIds

`func (o *BatchTransferCurrencyRequest) SetWalletIds(v []string)`

SetWalletIds sets WalletIds field to given value.

### HasWalletIds

`func (o *BatchTransferCurrencyRequest) HasWalletIds() bool`

HasWalletIds returns a boolean if a field has been set.

### GetAmounts

`func (o *BatchTransferCurrencyRequest) GetAmounts() []float32`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *BatchTransferCurrencyRequest) GetAmountsOk() (*[]float32, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *BatchTransferCurrencyRequest) SetAmounts(v []float32)`

SetAmounts sets Amounts field to given value.


### GetReferences

`func (o *BatchTransferCurrencyRequest) GetReferences() []float32`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *BatchTransferCurrencyRequest) GetReferencesOk() (*[]float32, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *BatchTransferCurrencyRequest) SetReferences(v []float32)`

SetReferences sets References field to given value.

### HasReferences

`func (o *BatchTransferCurrencyRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


