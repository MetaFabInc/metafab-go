# MintCollectionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | A valid EVM based address to create (mint) the item(s) for. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | [optional] 
**Quantity** | **float32** | The quantity of the specified item id to create (mint). | 
**WalletId** | Pointer to **string** | Any wallet id within the MetaFab ecosystem to create (mint) the item(s) for. | [optional] 

## Methods

### NewMintCollectionItemRequest

`func NewMintCollectionItemRequest(quantity float32, ) *MintCollectionItemRequest`

NewMintCollectionItemRequest instantiates a new MintCollectionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintCollectionItemRequestWithDefaults

`func NewMintCollectionItemRequestWithDefaults() *MintCollectionItemRequest`

NewMintCollectionItemRequestWithDefaults instantiates a new MintCollectionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MintCollectionItemRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MintCollectionItemRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MintCollectionItemRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MintCollectionItemRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetQuantity

`func (o *MintCollectionItemRequest) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MintCollectionItemRequest) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MintCollectionItemRequest) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetWalletId

`func (o *MintCollectionItemRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MintCollectionItemRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MintCollectionItemRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *MintCollectionItemRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


