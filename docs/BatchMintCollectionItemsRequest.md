# BatchMintCollectionItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | A valid EVM based address to create (mint) the items for. For example, &#x60;0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D&#x60;. | [optional] 
**ItemIds** | **[]int32** | An array of unique itemIds to create (mint). | 
**Quantities** | **[]int32** | An array of the quantities of each of the unique itemIds provided to create (mint). The quantity of each itemId in itemIds should be at the same index as the specific itemId in the itemIds array. For example, quantities[2] defines the quantity to mint for itemIds[2], etc. | 
**WalletId** | Pointer to **string** | Any wallet id within the MetaFab ecosystem to create (mint) the items for. | [optional] 

## Methods

### NewBatchMintCollectionItemsRequest

`func NewBatchMintCollectionItemsRequest(itemIds []int32, quantities []int32, ) *BatchMintCollectionItemsRequest`

NewBatchMintCollectionItemsRequest instantiates a new BatchMintCollectionItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchMintCollectionItemsRequestWithDefaults

`func NewBatchMintCollectionItemsRequestWithDefaults() *BatchMintCollectionItemsRequest`

NewBatchMintCollectionItemsRequestWithDefaults instantiates a new BatchMintCollectionItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BatchMintCollectionItemsRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BatchMintCollectionItemsRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BatchMintCollectionItemsRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BatchMintCollectionItemsRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetItemIds

`func (o *BatchMintCollectionItemsRequest) GetItemIds() []int32`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *BatchMintCollectionItemsRequest) GetItemIdsOk() (*[]int32, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *BatchMintCollectionItemsRequest) SetItemIds(v []int32)`

SetItemIds sets ItemIds field to given value.


### GetQuantities

`func (o *BatchMintCollectionItemsRequest) GetQuantities() []int32`

GetQuantities returns the Quantities field if non-nil, zero value otherwise.

### GetQuantitiesOk

`func (o *BatchMintCollectionItemsRequest) GetQuantitiesOk() (*[]int32, bool)`

GetQuantitiesOk returns a tuple with the Quantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantities

`func (o *BatchMintCollectionItemsRequest) SetQuantities(v []int32)`

SetQuantities sets Quantities field to given value.


### GetWalletId

`func (o *BatchMintCollectionItemsRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *BatchMintCollectionItemsRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *BatchMintCollectionItemsRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *BatchMintCollectionItemsRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


