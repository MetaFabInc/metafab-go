# BatchTransferCollectionItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | An array of valid EVM based addresses to transfer items to. | [optional] 
**WalletIds** | Pointer to **[]string** | An array of wallet ids within the MetaFab ecosystem to transfer items to. | [optional] 
**ItemIds** | **[]float32** | An array of unique itemIds to transfer. Each recipient will receive the same set of items provided. | 
**Quantities** | **[]float32** | The quantities of each unique itemId to transfer. Each recipient will receive the same quantities of items provided. | 

## Methods

### NewBatchTransferCollectionItemsRequest

`func NewBatchTransferCollectionItemsRequest(itemIds []float32, quantities []float32, ) *BatchTransferCollectionItemsRequest`

NewBatchTransferCollectionItemsRequest instantiates a new BatchTransferCollectionItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchTransferCollectionItemsRequestWithDefaults

`func NewBatchTransferCollectionItemsRequestWithDefaults() *BatchTransferCollectionItemsRequest`

NewBatchTransferCollectionItemsRequestWithDefaults instantiates a new BatchTransferCollectionItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *BatchTransferCollectionItemsRequest) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BatchTransferCollectionItemsRequest) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BatchTransferCollectionItemsRequest) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BatchTransferCollectionItemsRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetWalletIds

`func (o *BatchTransferCollectionItemsRequest) GetWalletIds() []string`

GetWalletIds returns the WalletIds field if non-nil, zero value otherwise.

### GetWalletIdsOk

`func (o *BatchTransferCollectionItemsRequest) GetWalletIdsOk() (*[]string, bool)`

GetWalletIdsOk returns a tuple with the WalletIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletIds

`func (o *BatchTransferCollectionItemsRequest) SetWalletIds(v []string)`

SetWalletIds sets WalletIds field to given value.

### HasWalletIds

`func (o *BatchTransferCollectionItemsRequest) HasWalletIds() bool`

HasWalletIds returns a boolean if a field has been set.

### GetItemIds

`func (o *BatchTransferCollectionItemsRequest) GetItemIds() []float32`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *BatchTransferCollectionItemsRequest) GetItemIdsOk() (*[]float32, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *BatchTransferCollectionItemsRequest) SetItemIds(v []float32)`

SetItemIds sets ItemIds field to given value.


### GetQuantities

`func (o *BatchTransferCollectionItemsRequest) GetQuantities() []float32`

GetQuantities returns the Quantities field if non-nil, zero value otherwise.

### GetQuantitiesOk

`func (o *BatchTransferCollectionItemsRequest) GetQuantitiesOk() (*[]float32, bool)`

GetQuantitiesOk returns a tuple with the Quantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantities

`func (o *BatchTransferCollectionItemsRequest) SetQuantities(v []float32)`

SetQuantities sets Quantities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


