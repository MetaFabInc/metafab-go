# ExchangeOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | The id of this offer. | [optional] 
**InputCollection** | Pointer to **string** | The address of the ERC1155 of MetaFab game items contract for input items required by this offer. | [optional] 
**InputCollectionItemIds** | Pointer to **[]float32** | An array of item ids from the input collection that are required for this offer. | [optional] 
**InputCollectionItemAmounts** | Pointer to **[]float32** | An array of amounts for each item id for the input collection that are required to use this offer. | [optional] 
**InputCurrency** | Pointer to **string** | The address of the ERC20 or MetaFab game currency for the currency required by this offer. | [optional] 
**InputCurrencyAmount** | Pointer to **float32** | The amount of currency required by this offer. | [optional] 
**OutputCollection** | Pointer to **string** | The address of the ERC1155 of MetaFab game items contract for output items given by this offer. | [optional] 
**OutputCollectionItemIds** | Pointer to **[]float32** | An array of item ids from the output collection that are given for this offer. | [optional] 
**OutputCollectionItemAmounts** | Pointer to **[]float32** | An array of amounts for each item id for the output collection that are given by this offer. | [optional] 
**OutputCurrency** | Pointer to **string** | The address of the ERC20 or MetaFab game currency for the output currency given by this offer. | [optional] 
**OutputCurrencyAmount** | Pointer to **float32** | The amount of currency given by this offer. | [optional] 
**Uses** | Pointer to **float32** | The number of times this offer has been used. | [optional] 
**MaxUses** | Pointer to **float32** | The maximum number of times this offer can be used. A value of &#x60;0&#x60; means there is no limit on how many times this offer can be used. | [optional] 
**LastUpdatedAt** | Pointer to **float32** | A unix timestamp in seconds that represents the last time this offer was set or updated. | [optional] 

## Methods

### NewExchangeOffer

`func NewExchangeOffer() *ExchangeOffer`

NewExchangeOffer instantiates a new ExchangeOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeOfferWithDefaults

`func NewExchangeOfferWithDefaults() *ExchangeOffer`

NewExchangeOfferWithDefaults instantiates a new ExchangeOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExchangeOffer) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExchangeOffer) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExchangeOffer) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ExchangeOffer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputCollection

`func (o *ExchangeOffer) GetInputCollection() string`

GetInputCollection returns the InputCollection field if non-nil, zero value otherwise.

### GetInputCollectionOk

`func (o *ExchangeOffer) GetInputCollectionOk() (*string, bool)`

GetInputCollectionOk returns a tuple with the InputCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollection

`func (o *ExchangeOffer) SetInputCollection(v string)`

SetInputCollection sets InputCollection field to given value.

### HasInputCollection

`func (o *ExchangeOffer) HasInputCollection() bool`

HasInputCollection returns a boolean if a field has been set.

### GetInputCollectionItemIds

`func (o *ExchangeOffer) GetInputCollectionItemIds() []float32`

GetInputCollectionItemIds returns the InputCollectionItemIds field if non-nil, zero value otherwise.

### GetInputCollectionItemIdsOk

`func (o *ExchangeOffer) GetInputCollectionItemIdsOk() (*[]float32, bool)`

GetInputCollectionItemIdsOk returns a tuple with the InputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemIds

`func (o *ExchangeOffer) SetInputCollectionItemIds(v []float32)`

SetInputCollectionItemIds sets InputCollectionItemIds field to given value.

### HasInputCollectionItemIds

`func (o *ExchangeOffer) HasInputCollectionItemIds() bool`

HasInputCollectionItemIds returns a boolean if a field has been set.

### GetInputCollectionItemAmounts

`func (o *ExchangeOffer) GetInputCollectionItemAmounts() []float32`

GetInputCollectionItemAmounts returns the InputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetInputCollectionItemAmountsOk

`func (o *ExchangeOffer) GetInputCollectionItemAmountsOk() (*[]float32, bool)`

GetInputCollectionItemAmountsOk returns a tuple with the InputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemAmounts

`func (o *ExchangeOffer) SetInputCollectionItemAmounts(v []float32)`

SetInputCollectionItemAmounts sets InputCollectionItemAmounts field to given value.

### HasInputCollectionItemAmounts

`func (o *ExchangeOffer) HasInputCollectionItemAmounts() bool`

HasInputCollectionItemAmounts returns a boolean if a field has been set.

### GetInputCurrency

`func (o *ExchangeOffer) GetInputCurrency() string`

GetInputCurrency returns the InputCurrency field if non-nil, zero value otherwise.

### GetInputCurrencyOk

`func (o *ExchangeOffer) GetInputCurrencyOk() (*string, bool)`

GetInputCurrencyOk returns a tuple with the InputCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCurrency

`func (o *ExchangeOffer) SetInputCurrency(v string)`

SetInputCurrency sets InputCurrency field to given value.

### HasInputCurrency

`func (o *ExchangeOffer) HasInputCurrency() bool`

HasInputCurrency returns a boolean if a field has been set.

### GetInputCurrencyAmount

`func (o *ExchangeOffer) GetInputCurrencyAmount() float32`

GetInputCurrencyAmount returns the InputCurrencyAmount field if non-nil, zero value otherwise.

### GetInputCurrencyAmountOk

`func (o *ExchangeOffer) GetInputCurrencyAmountOk() (*float32, bool)`

GetInputCurrencyAmountOk returns a tuple with the InputCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCurrencyAmount

`func (o *ExchangeOffer) SetInputCurrencyAmount(v float32)`

SetInputCurrencyAmount sets InputCurrencyAmount field to given value.

### HasInputCurrencyAmount

`func (o *ExchangeOffer) HasInputCurrencyAmount() bool`

HasInputCurrencyAmount returns a boolean if a field has been set.

### GetOutputCollection

`func (o *ExchangeOffer) GetOutputCollection() string`

GetOutputCollection returns the OutputCollection field if non-nil, zero value otherwise.

### GetOutputCollectionOk

`func (o *ExchangeOffer) GetOutputCollectionOk() (*string, bool)`

GetOutputCollectionOk returns a tuple with the OutputCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollection

`func (o *ExchangeOffer) SetOutputCollection(v string)`

SetOutputCollection sets OutputCollection field to given value.

### HasOutputCollection

`func (o *ExchangeOffer) HasOutputCollection() bool`

HasOutputCollection returns a boolean if a field has been set.

### GetOutputCollectionItemIds

`func (o *ExchangeOffer) GetOutputCollectionItemIds() []float32`

GetOutputCollectionItemIds returns the OutputCollectionItemIds field if non-nil, zero value otherwise.

### GetOutputCollectionItemIdsOk

`func (o *ExchangeOffer) GetOutputCollectionItemIdsOk() (*[]float32, bool)`

GetOutputCollectionItemIdsOk returns a tuple with the OutputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemIds

`func (o *ExchangeOffer) SetOutputCollectionItemIds(v []float32)`

SetOutputCollectionItemIds sets OutputCollectionItemIds field to given value.

### HasOutputCollectionItemIds

`func (o *ExchangeOffer) HasOutputCollectionItemIds() bool`

HasOutputCollectionItemIds returns a boolean if a field has been set.

### GetOutputCollectionItemAmounts

`func (o *ExchangeOffer) GetOutputCollectionItemAmounts() []float32`

GetOutputCollectionItemAmounts returns the OutputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetOutputCollectionItemAmountsOk

`func (o *ExchangeOffer) GetOutputCollectionItemAmountsOk() (*[]float32, bool)`

GetOutputCollectionItemAmountsOk returns a tuple with the OutputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemAmounts

`func (o *ExchangeOffer) SetOutputCollectionItemAmounts(v []float32)`

SetOutputCollectionItemAmounts sets OutputCollectionItemAmounts field to given value.

### HasOutputCollectionItemAmounts

`func (o *ExchangeOffer) HasOutputCollectionItemAmounts() bool`

HasOutputCollectionItemAmounts returns a boolean if a field has been set.

### GetOutputCurrency

`func (o *ExchangeOffer) GetOutputCurrency() string`

GetOutputCurrency returns the OutputCurrency field if non-nil, zero value otherwise.

### GetOutputCurrencyOk

`func (o *ExchangeOffer) GetOutputCurrencyOk() (*string, bool)`

GetOutputCurrencyOk returns a tuple with the OutputCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCurrency

`func (o *ExchangeOffer) SetOutputCurrency(v string)`

SetOutputCurrency sets OutputCurrency field to given value.

### HasOutputCurrency

`func (o *ExchangeOffer) HasOutputCurrency() bool`

HasOutputCurrency returns a boolean if a field has been set.

### GetOutputCurrencyAmount

`func (o *ExchangeOffer) GetOutputCurrencyAmount() float32`

GetOutputCurrencyAmount returns the OutputCurrencyAmount field if non-nil, zero value otherwise.

### GetOutputCurrencyAmountOk

`func (o *ExchangeOffer) GetOutputCurrencyAmountOk() (*float32, bool)`

GetOutputCurrencyAmountOk returns a tuple with the OutputCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCurrencyAmount

`func (o *ExchangeOffer) SetOutputCurrencyAmount(v float32)`

SetOutputCurrencyAmount sets OutputCurrencyAmount field to given value.

### HasOutputCurrencyAmount

`func (o *ExchangeOffer) HasOutputCurrencyAmount() bool`

HasOutputCurrencyAmount returns a boolean if a field has been set.

### GetUses

`func (o *ExchangeOffer) GetUses() float32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *ExchangeOffer) GetUsesOk() (*float32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *ExchangeOffer) SetUses(v float32)`

SetUses sets Uses field to given value.

### HasUses

`func (o *ExchangeOffer) HasUses() bool`

HasUses returns a boolean if a field has been set.

### GetMaxUses

`func (o *ExchangeOffer) GetMaxUses() float32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *ExchangeOffer) GetMaxUsesOk() (*float32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *ExchangeOffer) SetMaxUses(v float32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *ExchangeOffer) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *ExchangeOffer) GetLastUpdatedAt() float32`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ExchangeOffer) GetLastUpdatedAtOk() (*float32, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ExchangeOffer) SetLastUpdatedAt(v float32)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ExchangeOffer) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


