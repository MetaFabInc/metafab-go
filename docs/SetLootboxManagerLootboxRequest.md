# SetLootboxManagerLootboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | A unique lootbox id to use for this lootbox for the lootbox manager. If an existing lootbox id is used, the current lootbox will be updated but the existing number of opens will be kept. If you want to reset the number of opens for an existing lootbox, first remove it using the remove lootbox endpoint, then set it. | 
**InputCollectionAddress** | Pointer to **string** | A valid EVM based ERC1155 or MetaFab game items contract address that represents the collection for input items required by this lootbox. &#x60;inputCollectionAddress&#x60; or &#x60;inputCollectionId&#x60; can optionally be provided. | [optional] 
**InputCollectionId** | Pointer to **string** | A valid MetaFab collection id that represents the collection for input items required by this lootbox. &#x60;inputCollectionAddress&#x60; or &#x60;inputCollectionId&#x60; can optionally be provided. | [optional] 
**InputCollectionItemIds** | Pointer to **[]int32** | An array of item ids from the provided input collection that are required to open this lootbox. Input items are burn upon opening a lootbox. | [optional] 
**InputCollectionItemAmounts** | Pointer to **[]int32** | An array of amounts for each item id from the provided input collection that are required to open this lootbox. Item amounts array indices are reflective of the amount required for a given item id at the same index. | [optional] 
**OutputCollectionAddress** | Pointer to **string** | A valid EVM based ERC1155 or MetaFab game items contract address that represents the collection for possible output items given by this lootbox. &#x60;outputCollectionAddress&#x60; or &#x60;outputCollectionId&#x60; can optionally be provided. | [optional] 
**OutputCollectionId** | Pointer to **string** | A valid MetaFab collection id that represents the collection for possible output items given by this lootbox. &#x60;outputCollectionAddress&#x60; or &#x60;outputCollectionId&#x60; can optionally be provided. | [optional] 
**OutputCollectionItemIds** | Pointer to **[]int32** | An array of item ids from the provided output collection that are possibly given by this lootbox. Randomly selected output items are automatically minted if the lootbox manager contract has the &#x60;minter&#x60; role for the output collection contract. Otherwise, they are transferred from the item balance held by the lootbox manager contract. | [optional] 
**OutputCollectionItemAmounts** | Pointer to **[]int32** | An array of amounts for each item id that can be randomly selected from the provided output collection that are given by this lootbox. Item amounts array indices are reflective of the amount required for a given item id at the same index. | [optional] 
**OutputCollectionItemWeights** | Pointer to **[]int32** | An array of weights for each item id that can be randomly selected from the provided output collection that are given by this lootbox. Any positive integer for an item&#39;s weight can be provided. The weight for an item relative to the sum of all possible item weights determines the probability that an item will be picked upon a lootbox being opened. Item weights array indices are reflective of the probability weight for a given item id at the same index. | [optional] 
**OutputTotalItems** | Pointer to **int32** | The number of items randomly selected from the possible output items when this lootbox is open. If you provide a value greater than 1, it is possible for the same item to be selected more than once, giving the opener more than one of that item&#39;s output from the lootbox. | [optional] 

## Methods

### NewSetLootboxManagerLootboxRequest

`func NewSetLootboxManagerLootboxRequest(id int32, ) *SetLootboxManagerLootboxRequest`

NewSetLootboxManagerLootboxRequest instantiates a new SetLootboxManagerLootboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetLootboxManagerLootboxRequestWithDefaults

`func NewSetLootboxManagerLootboxRequestWithDefaults() *SetLootboxManagerLootboxRequest`

NewSetLootboxManagerLootboxRequestWithDefaults instantiates a new SetLootboxManagerLootboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetLootboxManagerLootboxRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetLootboxManagerLootboxRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetLootboxManagerLootboxRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetInputCollectionAddress

`func (o *SetLootboxManagerLootboxRequest) GetInputCollectionAddress() string`

GetInputCollectionAddress returns the InputCollectionAddress field if non-nil, zero value otherwise.

### GetInputCollectionAddressOk

`func (o *SetLootboxManagerLootboxRequest) GetInputCollectionAddressOk() (*string, bool)`

GetInputCollectionAddressOk returns a tuple with the InputCollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionAddress

`func (o *SetLootboxManagerLootboxRequest) SetInputCollectionAddress(v string)`

SetInputCollectionAddress sets InputCollectionAddress field to given value.

### HasInputCollectionAddress

`func (o *SetLootboxManagerLootboxRequest) HasInputCollectionAddress() bool`

HasInputCollectionAddress returns a boolean if a field has been set.

### GetInputCollectionId

`func (o *SetLootboxManagerLootboxRequest) GetInputCollectionId() string`

GetInputCollectionId returns the InputCollectionId field if non-nil, zero value otherwise.

### GetInputCollectionIdOk

`func (o *SetLootboxManagerLootboxRequest) GetInputCollectionIdOk() (*string, bool)`

GetInputCollectionIdOk returns a tuple with the InputCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionId

`func (o *SetLootboxManagerLootboxRequest) SetInputCollectionId(v string)`

SetInputCollectionId sets InputCollectionId field to given value.

### HasInputCollectionId

`func (o *SetLootboxManagerLootboxRequest) HasInputCollectionId() bool`

HasInputCollectionId returns a boolean if a field has been set.

### GetInputCollectionItemIds

`func (o *SetLootboxManagerLootboxRequest) GetInputCollectionItemIds() []int32`

GetInputCollectionItemIds returns the InputCollectionItemIds field if non-nil, zero value otherwise.

### GetInputCollectionItemIdsOk

`func (o *SetLootboxManagerLootboxRequest) GetInputCollectionItemIdsOk() (*[]int32, bool)`

GetInputCollectionItemIdsOk returns a tuple with the InputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemIds

`func (o *SetLootboxManagerLootboxRequest) SetInputCollectionItemIds(v []int32)`

SetInputCollectionItemIds sets InputCollectionItemIds field to given value.

### HasInputCollectionItemIds

`func (o *SetLootboxManagerLootboxRequest) HasInputCollectionItemIds() bool`

HasInputCollectionItemIds returns a boolean if a field has been set.

### GetInputCollectionItemAmounts

`func (o *SetLootboxManagerLootboxRequest) GetInputCollectionItemAmounts() []int32`

GetInputCollectionItemAmounts returns the InputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetInputCollectionItemAmountsOk

`func (o *SetLootboxManagerLootboxRequest) GetInputCollectionItemAmountsOk() (*[]int32, bool)`

GetInputCollectionItemAmountsOk returns a tuple with the InputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemAmounts

`func (o *SetLootboxManagerLootboxRequest) SetInputCollectionItemAmounts(v []int32)`

SetInputCollectionItemAmounts sets InputCollectionItemAmounts field to given value.

### HasInputCollectionItemAmounts

`func (o *SetLootboxManagerLootboxRequest) HasInputCollectionItemAmounts() bool`

HasInputCollectionItemAmounts returns a boolean if a field has been set.

### GetOutputCollectionAddress

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionAddress() string`

GetOutputCollectionAddress returns the OutputCollectionAddress field if non-nil, zero value otherwise.

### GetOutputCollectionAddressOk

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionAddressOk() (*string, bool)`

GetOutputCollectionAddressOk returns a tuple with the OutputCollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionAddress

`func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionAddress(v string)`

SetOutputCollectionAddress sets OutputCollectionAddress field to given value.

### HasOutputCollectionAddress

`func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionAddress() bool`

HasOutputCollectionAddress returns a boolean if a field has been set.

### GetOutputCollectionId

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionId() string`

GetOutputCollectionId returns the OutputCollectionId field if non-nil, zero value otherwise.

### GetOutputCollectionIdOk

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionIdOk() (*string, bool)`

GetOutputCollectionIdOk returns a tuple with the OutputCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionId

`func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionId(v string)`

SetOutputCollectionId sets OutputCollectionId field to given value.

### HasOutputCollectionId

`func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionId() bool`

HasOutputCollectionId returns a boolean if a field has been set.

### GetOutputCollectionItemIds

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemIds() []int32`

GetOutputCollectionItemIds returns the OutputCollectionItemIds field if non-nil, zero value otherwise.

### GetOutputCollectionItemIdsOk

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemIdsOk() (*[]int32, bool)`

GetOutputCollectionItemIdsOk returns a tuple with the OutputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemIds

`func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionItemIds(v []int32)`

SetOutputCollectionItemIds sets OutputCollectionItemIds field to given value.

### HasOutputCollectionItemIds

`func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionItemIds() bool`

HasOutputCollectionItemIds returns a boolean if a field has been set.

### GetOutputCollectionItemAmounts

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemAmounts() []int32`

GetOutputCollectionItemAmounts returns the OutputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetOutputCollectionItemAmountsOk

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemAmountsOk() (*[]int32, bool)`

GetOutputCollectionItemAmountsOk returns a tuple with the OutputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemAmounts

`func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionItemAmounts(v []int32)`

SetOutputCollectionItemAmounts sets OutputCollectionItemAmounts field to given value.

### HasOutputCollectionItemAmounts

`func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionItemAmounts() bool`

HasOutputCollectionItemAmounts returns a boolean if a field has been set.

### GetOutputCollectionItemWeights

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemWeights() []int32`

GetOutputCollectionItemWeights returns the OutputCollectionItemWeights field if non-nil, zero value otherwise.

### GetOutputCollectionItemWeightsOk

`func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemWeightsOk() (*[]int32, bool)`

GetOutputCollectionItemWeightsOk returns a tuple with the OutputCollectionItemWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemWeights

`func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionItemWeights(v []int32)`

SetOutputCollectionItemWeights sets OutputCollectionItemWeights field to given value.

### HasOutputCollectionItemWeights

`func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionItemWeights() bool`

HasOutputCollectionItemWeights returns a boolean if a field has been set.

### GetOutputTotalItems

`func (o *SetLootboxManagerLootboxRequest) GetOutputTotalItems() int32`

GetOutputTotalItems returns the OutputTotalItems field if non-nil, zero value otherwise.

### GetOutputTotalItemsOk

`func (o *SetLootboxManagerLootboxRequest) GetOutputTotalItemsOk() (*int32, bool)`

GetOutputTotalItemsOk returns a tuple with the OutputTotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTotalItems

`func (o *SetLootboxManagerLootboxRequest) SetOutputTotalItems(v int32)`

SetOutputTotalItems sets OutputTotalItems field to given value.

### HasOutputTotalItems

`func (o *SetLootboxManagerLootboxRequest) HasOutputTotalItems() bool`

HasOutputTotalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


