# ShopOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of this offer. | [optional] 
**InputCollection** | Pointer to **string** | The address of the ERC1155 or MetaFab game items contract for input items required by this offer. | [optional] 
**InputCollectionItemIds** | Pointer to **[]int32** | An array of item ids from the input collection that are required for this offer. | [optional] 
**InputCollectionItemAmounts** | Pointer to **[]int32** | An array of amounts for each item id for the input collection that are required to use this offer. | [optional] 
**InputCurrency** | Pointer to **string** | The address of the ERC20 or MetaFab game currency for the currency required by this offer. | [optional] 
**InputCurrencyAmount** | Pointer to **float32** | The amount of currency required by this offer. | [optional] 
**OutputCollection** | Pointer to **string** | The address of the ERC1155 or MetaFab game items contract for output items given by this offer. | [optional] 
**OutputCollectionItemIds** | Pointer to **[]int32** | An array of item ids from the output collection that are given for this offer. | [optional] 
**OutputCollectionItemAmounts** | Pointer to **[]int32** | An array of amounts for each item id for the output collection that are given by this offer. | [optional] 
**OutputCurrency** | Pointer to **string** | The address of the ERC20 or MetaFab game currency for the output currency given by this offer. | [optional] 
**OutputCurrencyAmount** | Pointer to **float32** | The amount of currency given by this offer. | [optional] 
**Uses** | Pointer to **int32** | The number of times this offer has been used. | [optional] 
**MaxUses** | Pointer to **int32** | The maximum number of times this offer can be used. A value of &#x60;0&#x60; means there is no limit on how many times this offer can be used. | [optional] 
**LastUpdatedAt** | Pointer to **int32** | A unix timestamp in seconds that represents the last time this offer was set or updated. | [optional] 

## Methods

### NewShopOffer

`func NewShopOffer() *ShopOffer`

NewShopOffer instantiates a new ShopOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopOfferWithDefaults

`func NewShopOfferWithDefaults() *ShopOffer`

NewShopOfferWithDefaults instantiates a new ShopOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShopOffer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShopOffer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShopOffer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShopOffer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputCollection

`func (o *ShopOffer) GetInputCollection() string`

GetInputCollection returns the InputCollection field if non-nil, zero value otherwise.

### GetInputCollectionOk

`func (o *ShopOffer) GetInputCollectionOk() (*string, bool)`

GetInputCollectionOk returns a tuple with the InputCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollection

`func (o *ShopOffer) SetInputCollection(v string)`

SetInputCollection sets InputCollection field to given value.

### HasInputCollection

`func (o *ShopOffer) HasInputCollection() bool`

HasInputCollection returns a boolean if a field has been set.

### GetInputCollectionItemIds

`func (o *ShopOffer) GetInputCollectionItemIds() []int32`

GetInputCollectionItemIds returns the InputCollectionItemIds field if non-nil, zero value otherwise.

### GetInputCollectionItemIdsOk

`func (o *ShopOffer) GetInputCollectionItemIdsOk() (*[]int32, bool)`

GetInputCollectionItemIdsOk returns a tuple with the InputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemIds

`func (o *ShopOffer) SetInputCollectionItemIds(v []int32)`

SetInputCollectionItemIds sets InputCollectionItemIds field to given value.

### HasInputCollectionItemIds

`func (o *ShopOffer) HasInputCollectionItemIds() bool`

HasInputCollectionItemIds returns a boolean if a field has been set.

### GetInputCollectionItemAmounts

`func (o *ShopOffer) GetInputCollectionItemAmounts() []int32`

GetInputCollectionItemAmounts returns the InputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetInputCollectionItemAmountsOk

`func (o *ShopOffer) GetInputCollectionItemAmountsOk() (*[]int32, bool)`

GetInputCollectionItemAmountsOk returns a tuple with the InputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemAmounts

`func (o *ShopOffer) SetInputCollectionItemAmounts(v []int32)`

SetInputCollectionItemAmounts sets InputCollectionItemAmounts field to given value.

### HasInputCollectionItemAmounts

`func (o *ShopOffer) HasInputCollectionItemAmounts() bool`

HasInputCollectionItemAmounts returns a boolean if a field has been set.

### GetInputCurrency

`func (o *ShopOffer) GetInputCurrency() string`

GetInputCurrency returns the InputCurrency field if non-nil, zero value otherwise.

### GetInputCurrencyOk

`func (o *ShopOffer) GetInputCurrencyOk() (*string, bool)`

GetInputCurrencyOk returns a tuple with the InputCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCurrency

`func (o *ShopOffer) SetInputCurrency(v string)`

SetInputCurrency sets InputCurrency field to given value.

### HasInputCurrency

`func (o *ShopOffer) HasInputCurrency() bool`

HasInputCurrency returns a boolean if a field has been set.

### GetInputCurrencyAmount

`func (o *ShopOffer) GetInputCurrencyAmount() float32`

GetInputCurrencyAmount returns the InputCurrencyAmount field if non-nil, zero value otherwise.

### GetInputCurrencyAmountOk

`func (o *ShopOffer) GetInputCurrencyAmountOk() (*float32, bool)`

GetInputCurrencyAmountOk returns a tuple with the InputCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCurrencyAmount

`func (o *ShopOffer) SetInputCurrencyAmount(v float32)`

SetInputCurrencyAmount sets InputCurrencyAmount field to given value.

### HasInputCurrencyAmount

`func (o *ShopOffer) HasInputCurrencyAmount() bool`

HasInputCurrencyAmount returns a boolean if a field has been set.

### GetOutputCollection

`func (o *ShopOffer) GetOutputCollection() string`

GetOutputCollection returns the OutputCollection field if non-nil, zero value otherwise.

### GetOutputCollectionOk

`func (o *ShopOffer) GetOutputCollectionOk() (*string, bool)`

GetOutputCollectionOk returns a tuple with the OutputCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollection

`func (o *ShopOffer) SetOutputCollection(v string)`

SetOutputCollection sets OutputCollection field to given value.

### HasOutputCollection

`func (o *ShopOffer) HasOutputCollection() bool`

HasOutputCollection returns a boolean if a field has been set.

### GetOutputCollectionItemIds

`func (o *ShopOffer) GetOutputCollectionItemIds() []int32`

GetOutputCollectionItemIds returns the OutputCollectionItemIds field if non-nil, zero value otherwise.

### GetOutputCollectionItemIdsOk

`func (o *ShopOffer) GetOutputCollectionItemIdsOk() (*[]int32, bool)`

GetOutputCollectionItemIdsOk returns a tuple with the OutputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemIds

`func (o *ShopOffer) SetOutputCollectionItemIds(v []int32)`

SetOutputCollectionItemIds sets OutputCollectionItemIds field to given value.

### HasOutputCollectionItemIds

`func (o *ShopOffer) HasOutputCollectionItemIds() bool`

HasOutputCollectionItemIds returns a boolean if a field has been set.

### GetOutputCollectionItemAmounts

`func (o *ShopOffer) GetOutputCollectionItemAmounts() []int32`

GetOutputCollectionItemAmounts returns the OutputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetOutputCollectionItemAmountsOk

`func (o *ShopOffer) GetOutputCollectionItemAmountsOk() (*[]int32, bool)`

GetOutputCollectionItemAmountsOk returns a tuple with the OutputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemAmounts

`func (o *ShopOffer) SetOutputCollectionItemAmounts(v []int32)`

SetOutputCollectionItemAmounts sets OutputCollectionItemAmounts field to given value.

### HasOutputCollectionItemAmounts

`func (o *ShopOffer) HasOutputCollectionItemAmounts() bool`

HasOutputCollectionItemAmounts returns a boolean if a field has been set.

### GetOutputCurrency

`func (o *ShopOffer) GetOutputCurrency() string`

GetOutputCurrency returns the OutputCurrency field if non-nil, zero value otherwise.

### GetOutputCurrencyOk

`func (o *ShopOffer) GetOutputCurrencyOk() (*string, bool)`

GetOutputCurrencyOk returns a tuple with the OutputCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCurrency

`func (o *ShopOffer) SetOutputCurrency(v string)`

SetOutputCurrency sets OutputCurrency field to given value.

### HasOutputCurrency

`func (o *ShopOffer) HasOutputCurrency() bool`

HasOutputCurrency returns a boolean if a field has been set.

### GetOutputCurrencyAmount

`func (o *ShopOffer) GetOutputCurrencyAmount() float32`

GetOutputCurrencyAmount returns the OutputCurrencyAmount field if non-nil, zero value otherwise.

### GetOutputCurrencyAmountOk

`func (o *ShopOffer) GetOutputCurrencyAmountOk() (*float32, bool)`

GetOutputCurrencyAmountOk returns a tuple with the OutputCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCurrencyAmount

`func (o *ShopOffer) SetOutputCurrencyAmount(v float32)`

SetOutputCurrencyAmount sets OutputCurrencyAmount field to given value.

### HasOutputCurrencyAmount

`func (o *ShopOffer) HasOutputCurrencyAmount() bool`

HasOutputCurrencyAmount returns a boolean if a field has been set.

### GetUses

`func (o *ShopOffer) GetUses() int32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *ShopOffer) GetUsesOk() (*int32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *ShopOffer) SetUses(v int32)`

SetUses sets Uses field to given value.

### HasUses

`func (o *ShopOffer) HasUses() bool`

HasUses returns a boolean if a field has been set.

### GetMaxUses

`func (o *ShopOffer) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *ShopOffer) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *ShopOffer) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *ShopOffer) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *ShopOffer) GetLastUpdatedAt() int32`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ShopOffer) GetLastUpdatedAtOk() (*int32, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ShopOffer) SetLastUpdatedAt(v int32)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ShopOffer) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


