# TransferCollectionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | A valid EVM based addresses to transfer items to. | [optional] 
**WalletId** | Pointer to **[]string** | A wallet id within the MetaFab ecosystem to transfer items to. | [optional] 
**Quantity** | **float32** | The quantity of the collectionItemId to transfer. | 

## Methods

### NewTransferCollectionItemRequest

`func NewTransferCollectionItemRequest(quantity float32, ) *TransferCollectionItemRequest`

NewTransferCollectionItemRequest instantiates a new TransferCollectionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferCollectionItemRequestWithDefaults

`func NewTransferCollectionItemRequestWithDefaults() *TransferCollectionItemRequest`

NewTransferCollectionItemRequestWithDefaults instantiates a new TransferCollectionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransferCollectionItemRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransferCollectionItemRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransferCollectionItemRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransferCollectionItemRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetWalletId

`func (o *TransferCollectionItemRequest) GetWalletId() []string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *TransferCollectionItemRequest) GetWalletIdOk() (*[]string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *TransferCollectionItemRequest) SetWalletId(v []string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *TransferCollectionItemRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetQuantity

`func (o *TransferCollectionItemRequest) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TransferCollectionItemRequest) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TransferCollectionItemRequest) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


