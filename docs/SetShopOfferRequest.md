# SetShopOfferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | A unique offer id to use for this offer for the shop. If an existing offer id is used, the current offer will be updated but the existing number of uses will be kept. If you want to reset the number of uses for an existing offer, first remove it using the remove offer endpoint, then set it. | 
**InputCollectionAddress** | Pointer to **string** | A valid EVM based ERC1155 or MetaFab game items contract address that represents the collection for input items required by this offer. &#x60;inputCollectionAddress&#x60; or &#x60;inputCollectionId&#x60; can optionally be provided. | [optional] 
**InputCollectionId** | Pointer to **string** | A valid MetaFab collection id that represents the collection for input items required by this offer. &#x60;inputCollectionAddress&#x60; or &#x60;inputCollectionId&#x60; can optionally be provided. | [optional] 
**InputCollectionItemIds** | Pointer to **[]int32** | An array of item ids from the provided input collection that are required to use this offer. Input items are transferred from the wallet to the shop contract upon using an offer. | [optional] 
**InputCollectionItemAmounts** | Pointer to **[]int32** | An array of amounts for each item id from the provided input collection that are required to use this offer. Item amounts array indices are reflective of the amount required for a given item id at the same index. | [optional] 
**InputCurrencyAddress** | Pointer to **string** | A valid EVM based ERC20 or MetaFab game currency contract address that for the currency required by this offer. | [optional] 
**InputCurrencyId** | Pointer to **string** | A valid MetaFab currency id that represents the currency required by this offer. | [optional] 
**InputCurrencyAmount** | Pointer to **float32** | The amount of currency required by this offer. If an inputCurrencyAmount is provided without in input currency address or id, the native chain currency is used as the required currency. | [optional] 
**OutputCollectionAddress** | Pointer to **string** | A valid EVM based ERC1155 or MetaFab game items contract address that represents the collection for output items given by this offer. &#x60;outputCollectionAddress&#x60; or &#x60;outputCollectionId&#x60; can optionally be provided. | [optional] 
**OutputCollectionId** | Pointer to **string** | A valid MetaFab collection id that represents the collection for output items given by this offer. &#x60;outputCollectionAddress&#x60; or &#x60;outputCollectionId&#x60; can optionally be provided. | [optional] 
**OutputCollectionItemIds** | Pointer to **[]int32** | An array of item ids from the provided output collection that are given by this offer. Output items are automatically minted if the shop contract has the &#x60;minter&#x60; role for the output collection contract. Otherwise, they are transferred from the item balance held by the shop contract. | [optional] 
**OutputCollectionItemAmounts** | Pointer to **[]int32** | An array of amounts for each item id from the provided output collection that are given by this offer. Item amounts array indices are reflective of the amount required for a given item id at the same index. | [optional] 
**OutputCurrencyAddress** | Pointer to **string** | A valid EVM based ERC20 or MetaFab game currency contract address that for the currency given by this offer. The output currency amount is automatically minted if the shop contract has the &#x60;minter&#x60; role for the output currency contract. Otherwise, they are transferred from the currency balance held by the shop contract. | [optional] 
**OutputCurrencyId** | Pointer to **string** | A valid MetaFab currency id for the currency given by this offer. | [optional] 
**OutputCurrencyAmount** | Pointer to **float32** | The amount of currency given by this offer. If an outputCurrencyAmount is provided without an output currency address or id, the native chain currency is used as the given currency. | [optional] 
**MaxUses** | Pointer to **int32** | The maximum number of times this offer can be used in total. maxUses is collective across all uses of the offer. If 5 unique players use an offer, that counts as 5 offer uses. Exclude this or use 0 to allow unlimited uses. | [optional] 

## Methods

### NewSetShopOfferRequest

`func NewSetShopOfferRequest(id int32, ) *SetShopOfferRequest`

NewSetShopOfferRequest instantiates a new SetShopOfferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetShopOfferRequestWithDefaults

`func NewSetShopOfferRequestWithDefaults() *SetShopOfferRequest`

NewSetShopOfferRequestWithDefaults instantiates a new SetShopOfferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetShopOfferRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetShopOfferRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetShopOfferRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetInputCollectionAddress

`func (o *SetShopOfferRequest) GetInputCollectionAddress() string`

GetInputCollectionAddress returns the InputCollectionAddress field if non-nil, zero value otherwise.

### GetInputCollectionAddressOk

`func (o *SetShopOfferRequest) GetInputCollectionAddressOk() (*string, bool)`

GetInputCollectionAddressOk returns a tuple with the InputCollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionAddress

`func (o *SetShopOfferRequest) SetInputCollectionAddress(v string)`

SetInputCollectionAddress sets InputCollectionAddress field to given value.

### HasInputCollectionAddress

`func (o *SetShopOfferRequest) HasInputCollectionAddress() bool`

HasInputCollectionAddress returns a boolean if a field has been set.

### GetInputCollectionId

`func (o *SetShopOfferRequest) GetInputCollectionId() string`

GetInputCollectionId returns the InputCollectionId field if non-nil, zero value otherwise.

### GetInputCollectionIdOk

`func (o *SetShopOfferRequest) GetInputCollectionIdOk() (*string, bool)`

GetInputCollectionIdOk returns a tuple with the InputCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionId

`func (o *SetShopOfferRequest) SetInputCollectionId(v string)`

SetInputCollectionId sets InputCollectionId field to given value.

### HasInputCollectionId

`func (o *SetShopOfferRequest) HasInputCollectionId() bool`

HasInputCollectionId returns a boolean if a field has been set.

### GetInputCollectionItemIds

`func (o *SetShopOfferRequest) GetInputCollectionItemIds() []int32`

GetInputCollectionItemIds returns the InputCollectionItemIds field if non-nil, zero value otherwise.

### GetInputCollectionItemIdsOk

`func (o *SetShopOfferRequest) GetInputCollectionItemIdsOk() (*[]int32, bool)`

GetInputCollectionItemIdsOk returns a tuple with the InputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemIds

`func (o *SetShopOfferRequest) SetInputCollectionItemIds(v []int32)`

SetInputCollectionItemIds sets InputCollectionItemIds field to given value.

### HasInputCollectionItemIds

`func (o *SetShopOfferRequest) HasInputCollectionItemIds() bool`

HasInputCollectionItemIds returns a boolean if a field has been set.

### GetInputCollectionItemAmounts

`func (o *SetShopOfferRequest) GetInputCollectionItemAmounts() []int32`

GetInputCollectionItemAmounts returns the InputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetInputCollectionItemAmountsOk

`func (o *SetShopOfferRequest) GetInputCollectionItemAmountsOk() (*[]int32, bool)`

GetInputCollectionItemAmountsOk returns a tuple with the InputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCollectionItemAmounts

`func (o *SetShopOfferRequest) SetInputCollectionItemAmounts(v []int32)`

SetInputCollectionItemAmounts sets InputCollectionItemAmounts field to given value.

### HasInputCollectionItemAmounts

`func (o *SetShopOfferRequest) HasInputCollectionItemAmounts() bool`

HasInputCollectionItemAmounts returns a boolean if a field has been set.

### GetInputCurrencyAddress

`func (o *SetShopOfferRequest) GetInputCurrencyAddress() string`

GetInputCurrencyAddress returns the InputCurrencyAddress field if non-nil, zero value otherwise.

### GetInputCurrencyAddressOk

`func (o *SetShopOfferRequest) GetInputCurrencyAddressOk() (*string, bool)`

GetInputCurrencyAddressOk returns a tuple with the InputCurrencyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCurrencyAddress

`func (o *SetShopOfferRequest) SetInputCurrencyAddress(v string)`

SetInputCurrencyAddress sets InputCurrencyAddress field to given value.

### HasInputCurrencyAddress

`func (o *SetShopOfferRequest) HasInputCurrencyAddress() bool`

HasInputCurrencyAddress returns a boolean if a field has been set.

### GetInputCurrencyId

`func (o *SetShopOfferRequest) GetInputCurrencyId() string`

GetInputCurrencyId returns the InputCurrencyId field if non-nil, zero value otherwise.

### GetInputCurrencyIdOk

`func (o *SetShopOfferRequest) GetInputCurrencyIdOk() (*string, bool)`

GetInputCurrencyIdOk returns a tuple with the InputCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCurrencyId

`func (o *SetShopOfferRequest) SetInputCurrencyId(v string)`

SetInputCurrencyId sets InputCurrencyId field to given value.

### HasInputCurrencyId

`func (o *SetShopOfferRequest) HasInputCurrencyId() bool`

HasInputCurrencyId returns a boolean if a field has been set.

### GetInputCurrencyAmount

`func (o *SetShopOfferRequest) GetInputCurrencyAmount() float32`

GetInputCurrencyAmount returns the InputCurrencyAmount field if non-nil, zero value otherwise.

### GetInputCurrencyAmountOk

`func (o *SetShopOfferRequest) GetInputCurrencyAmountOk() (*float32, bool)`

GetInputCurrencyAmountOk returns a tuple with the InputCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCurrencyAmount

`func (o *SetShopOfferRequest) SetInputCurrencyAmount(v float32)`

SetInputCurrencyAmount sets InputCurrencyAmount field to given value.

### HasInputCurrencyAmount

`func (o *SetShopOfferRequest) HasInputCurrencyAmount() bool`

HasInputCurrencyAmount returns a boolean if a field has been set.

### GetOutputCollectionAddress

`func (o *SetShopOfferRequest) GetOutputCollectionAddress() string`

GetOutputCollectionAddress returns the OutputCollectionAddress field if non-nil, zero value otherwise.

### GetOutputCollectionAddressOk

`func (o *SetShopOfferRequest) GetOutputCollectionAddressOk() (*string, bool)`

GetOutputCollectionAddressOk returns a tuple with the OutputCollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionAddress

`func (o *SetShopOfferRequest) SetOutputCollectionAddress(v string)`

SetOutputCollectionAddress sets OutputCollectionAddress field to given value.

### HasOutputCollectionAddress

`func (o *SetShopOfferRequest) HasOutputCollectionAddress() bool`

HasOutputCollectionAddress returns a boolean if a field has been set.

### GetOutputCollectionId

`func (o *SetShopOfferRequest) GetOutputCollectionId() string`

GetOutputCollectionId returns the OutputCollectionId field if non-nil, zero value otherwise.

### GetOutputCollectionIdOk

`func (o *SetShopOfferRequest) GetOutputCollectionIdOk() (*string, bool)`

GetOutputCollectionIdOk returns a tuple with the OutputCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionId

`func (o *SetShopOfferRequest) SetOutputCollectionId(v string)`

SetOutputCollectionId sets OutputCollectionId field to given value.

### HasOutputCollectionId

`func (o *SetShopOfferRequest) HasOutputCollectionId() bool`

HasOutputCollectionId returns a boolean if a field has been set.

### GetOutputCollectionItemIds

`func (o *SetShopOfferRequest) GetOutputCollectionItemIds() []int32`

GetOutputCollectionItemIds returns the OutputCollectionItemIds field if non-nil, zero value otherwise.

### GetOutputCollectionItemIdsOk

`func (o *SetShopOfferRequest) GetOutputCollectionItemIdsOk() (*[]int32, bool)`

GetOutputCollectionItemIdsOk returns a tuple with the OutputCollectionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemIds

`func (o *SetShopOfferRequest) SetOutputCollectionItemIds(v []int32)`

SetOutputCollectionItemIds sets OutputCollectionItemIds field to given value.

### HasOutputCollectionItemIds

`func (o *SetShopOfferRequest) HasOutputCollectionItemIds() bool`

HasOutputCollectionItemIds returns a boolean if a field has been set.

### GetOutputCollectionItemAmounts

`func (o *SetShopOfferRequest) GetOutputCollectionItemAmounts() []int32`

GetOutputCollectionItemAmounts returns the OutputCollectionItemAmounts field if non-nil, zero value otherwise.

### GetOutputCollectionItemAmountsOk

`func (o *SetShopOfferRequest) GetOutputCollectionItemAmountsOk() (*[]int32, bool)`

GetOutputCollectionItemAmountsOk returns a tuple with the OutputCollectionItemAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCollectionItemAmounts

`func (o *SetShopOfferRequest) SetOutputCollectionItemAmounts(v []int32)`

SetOutputCollectionItemAmounts sets OutputCollectionItemAmounts field to given value.

### HasOutputCollectionItemAmounts

`func (o *SetShopOfferRequest) HasOutputCollectionItemAmounts() bool`

HasOutputCollectionItemAmounts returns a boolean if a field has been set.

### GetOutputCurrencyAddress

`func (o *SetShopOfferRequest) GetOutputCurrencyAddress() string`

GetOutputCurrencyAddress returns the OutputCurrencyAddress field if non-nil, zero value otherwise.

### GetOutputCurrencyAddressOk

`func (o *SetShopOfferRequest) GetOutputCurrencyAddressOk() (*string, bool)`

GetOutputCurrencyAddressOk returns a tuple with the OutputCurrencyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCurrencyAddress

`func (o *SetShopOfferRequest) SetOutputCurrencyAddress(v string)`

SetOutputCurrencyAddress sets OutputCurrencyAddress field to given value.

### HasOutputCurrencyAddress

`func (o *SetShopOfferRequest) HasOutputCurrencyAddress() bool`

HasOutputCurrencyAddress returns a boolean if a field has been set.

### GetOutputCurrencyId

`func (o *SetShopOfferRequest) GetOutputCurrencyId() string`

GetOutputCurrencyId returns the OutputCurrencyId field if non-nil, zero value otherwise.

### GetOutputCurrencyIdOk

`func (o *SetShopOfferRequest) GetOutputCurrencyIdOk() (*string, bool)`

GetOutputCurrencyIdOk returns a tuple with the OutputCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCurrencyId

`func (o *SetShopOfferRequest) SetOutputCurrencyId(v string)`

SetOutputCurrencyId sets OutputCurrencyId field to given value.

### HasOutputCurrencyId

`func (o *SetShopOfferRequest) HasOutputCurrencyId() bool`

HasOutputCurrencyId returns a boolean if a field has been set.

### GetOutputCurrencyAmount

`func (o *SetShopOfferRequest) GetOutputCurrencyAmount() float32`

GetOutputCurrencyAmount returns the OutputCurrencyAmount field if non-nil, zero value otherwise.

### GetOutputCurrencyAmountOk

`func (o *SetShopOfferRequest) GetOutputCurrencyAmountOk() (*float32, bool)`

GetOutputCurrencyAmountOk returns a tuple with the OutputCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCurrencyAmount

`func (o *SetShopOfferRequest) SetOutputCurrencyAmount(v float32)`

SetOutputCurrencyAmount sets OutputCurrencyAmount field to given value.

### HasOutputCurrencyAmount

`func (o *SetShopOfferRequest) HasOutputCurrencyAmount() bool`

HasOutputCurrencyAmount returns a boolean if a field has been set.

### GetMaxUses

`func (o *SetShopOfferRequest) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *SetShopOfferRequest) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *SetShopOfferRequest) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *SetShopOfferRequest) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


