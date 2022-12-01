# LootboxManagerLootbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of this lootbox. | [optional] 
**InputCollection** | Pointer to **string** | The address of the ERC1155 or MetaFab game items contract for input items required by this lootbox. | [optional] 
**InputCollectionItemIds** | Pointer to **[]int32** | An array of item ids from the input collection that are required for this lootbox. | [optional] 
**InputCollectionItemAmounts** | Pointer to **[]int32** | An array of amounts for each item id for the input collection that are required to open this lootbox. | [optional] 
**OutputCollection** | Pointer to **string** | The address of the ERC1155 of MetaFab game items contract for possible output items given by this lootbox. | [optional] 
**OutputCollectionItemIds** | Pointer to **[]int32** | An array of item ids from the output collection that are possibly given by this lootbox. | [optional] 
**OutputCollectionItemAmounts** | Pointer to **[]int32** | An array of amounts for each item id for the output collection that are possibly given by this lootbox. | [optional] 
**OutputCollectionItemWeights** | Pointer to **[]int32** | An array of weights for each item id for the output collection that are possibly given by this lootbox. | [optional] 
**OutputTotalItems** | Pointer to **int32** | The number of items randomly selected when this lootbox is opened. | [optional] 
**LastUpdatedAt** | Pointer to **int32** | A unix timestamp in seconds that represents the last time this offer was set or updated. | [optional] 

## Methods

### NewLootboxManagerLootbox

`func NewLootboxManagerLootbox() *LootboxManagerLootbox`

NewLootboxManagerLootbox instantiates a new LootboxManagerLootbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLootboxManagerLootboxWithDefaults

`func NewLootboxManagerLootboxWithDefaults() *LootboxManagerLootbox`

NewLootboxManagerLootboxWithDefaults instantiates a new LootboxManagerLootbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LootboxManagerLootbox) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LootboxManagerLootbox) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LootboxManagerLootbox) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LootboxManagerLootbox) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputCollection

`func (o *LootboxManagerLootbox) GetInputCollection() string`

GetInputCollection returns the InputCollection field if non-nil, zero value otherwise.

### GetInputCollectionOk

`func (o *LootboxManagerLootbox) GetInputCollectionOk() (*string, bool)`

GetInputCollectionOk returns a tuple with the InputCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollection

`func (o *LootboxManagerLootbox) SetInputCollection(v string)`

SetInputCollection sets InputCollection field to given value.

### HasInputCollection

`func (o *LootboxManagerLootbox) HasInputCollection() bool`

HasInputCollection returns a boolean if a field has been set.

### GetInputCollectionItemIds

`func (o *LootboxManagerLootbox) GetInputCollectionItemIds() []int32`

GetInputCollectionItemIds returns the InputCollectionItemIds field if non-nil, zero value otherwise.

### GetInputCollectionItemIdsOk

`func (o *LootboxManagerLootbox) GetInputCollectionItemIdsOk() (*[]int32, bool)`

GetInputCollectionItemIdsOk returns a tuple with the InputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemIds

`func (o *LootboxManagerLootbox) SetInputCollectionItemIds(v []int32)`

SetInputCollectionItemIds sets InputCollectionItemIds field to given value.

### HasInputCollectionItemIds

`func (o *LootboxManagerLootbox) HasInputCollectionItemIds() bool`

HasInputCollectionItemIds returns a boolean if a field has been set.

### GetInputCollectionItemAmounts

`func (o *LootboxManagerLootbox) GetInputCollectionItemAmounts() []int32`

GetInputCollectionItemAmounts returns the InputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetInputCollectionItemAmountsOk

`func (o *LootboxManagerLootbox) GetInputCollectionItemAmountsOk() (*[]int32, bool)`

GetInputCollectionItemAmountsOk returns a tuple with the InputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemAmounts

`func (o *LootboxManagerLootbox) SetInputCollectionItemAmounts(v []int32)`

SetInputCollectionItemAmounts sets InputCollectionItemAmounts field to given value.

### HasInputCollectionItemAmounts

`func (o *LootboxManagerLootbox) HasInputCollectionItemAmounts() bool`

HasInputCollectionItemAmounts returns a boolean if a field has been set.

### GetOutputCollection

`func (o *LootboxManagerLootbox) GetOutputCollection() string`

GetOutputCollection returns the OutputCollection field if non-nil, zero value otherwise.

### GetOutputCollectionOk

`func (o *LootboxManagerLootbox) GetOutputCollectionOk() (*string, bool)`

GetOutputCollectionOk returns a tuple with the OutputCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollection

`func (o *LootboxManagerLootbox) SetOutputCollection(v string)`

SetOutputCollection sets OutputCollection field to given value.

### HasOutputCollection

`func (o *LootboxManagerLootbox) HasOutputCollection() bool`

HasOutputCollection returns a boolean if a field has been set.

### GetOutputCollectionItemIds

`func (o *LootboxManagerLootbox) GetOutputCollectionItemIds() []int32`

GetOutputCollectionItemIds returns the OutputCollectionItemIds field if non-nil, zero value otherwise.

### GetOutputCollectionItemIdsOk

`func (o *LootboxManagerLootbox) GetOutputCollectionItemIdsOk() (*[]int32, bool)`

GetOutputCollectionItemIdsOk returns a tuple with the OutputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemIds

`func (o *LootboxManagerLootbox) SetOutputCollectionItemIds(v []int32)`

SetOutputCollectionItemIds sets OutputCollectionItemIds field to given value.

### HasOutputCollectionItemIds

`func (o *LootboxManagerLootbox) HasOutputCollectionItemIds() bool`

HasOutputCollectionItemIds returns a boolean if a field has been set.

### GetOutputCollectionItemAmounts

`func (o *LootboxManagerLootbox) GetOutputCollectionItemAmounts() []int32`

GetOutputCollectionItemAmounts returns the OutputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetOutputCollectionItemAmountsOk

`func (o *LootboxManagerLootbox) GetOutputCollectionItemAmountsOk() (*[]int32, bool)`

GetOutputCollectionItemAmountsOk returns a tuple with the OutputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemAmounts

`func (o *LootboxManagerLootbox) SetOutputCollectionItemAmounts(v []int32)`

SetOutputCollectionItemAmounts sets OutputCollectionItemAmounts field to given value.

### HasOutputCollectionItemAmounts

`func (o *LootboxManagerLootbox) HasOutputCollectionItemAmounts() bool`

HasOutputCollectionItemAmounts returns a boolean if a field has been set.

### GetOutputCollectionItemWeights

`func (o *LootboxManagerLootbox) GetOutputCollectionItemWeights() []int32`

GetOutputCollectionItemWeights returns the OutputCollectionItemWeights field if non-nil, zero value otherwise.

### GetOutputCollectionItemWeightsOk

`func (o *LootboxManagerLootbox) GetOutputCollectionItemWeightsOk() (*[]int32, bool)`

GetOutputCollectionItemWeightsOk returns a tuple with the OutputCollectionItemWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemWeights

`func (o *LootboxManagerLootbox) SetOutputCollectionItemWeights(v []int32)`

SetOutputCollectionItemWeights sets OutputCollectionItemWeights field to given value.

### HasOutputCollectionItemWeights

`func (o *LootboxManagerLootbox) HasOutputCollectionItemWeights() bool`

HasOutputCollectionItemWeights returns a boolean if a field has been set.

### GetOutputTotalItems

`func (o *LootboxManagerLootbox) GetOutputTotalItems() int32`

GetOutputTotalItems returns the OutputTotalItems field if non-nil, zero value otherwise.

### GetOutputTotalItemsOk

`func (o *LootboxManagerLootbox) GetOutputTotalItemsOk() (*int32, bool)`

GetOutputTotalItemsOk returns a tuple with the OutputTotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTotalItems

`func (o *LootboxManagerLootbox) SetOutputTotalItems(v int32)`

SetOutputTotalItems sets OutputTotalItems field to given value.

### HasOutputTotalItems

`func (o *LootboxManagerLootbox) HasOutputTotalItems() bool`

HasOutputTotalItems returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *LootboxManagerLootbox) GetLastUpdatedAt() int32`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *LootboxManagerLootbox) GetLastUpdatedAtOk() (*int32, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *LootboxManagerLootbox) SetLastUpdatedAt(v int32)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *LootboxManagerLootbox) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


