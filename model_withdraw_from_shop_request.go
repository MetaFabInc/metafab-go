/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.3.0
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// WithdrawFromShopRequest struct for WithdrawFromShopRequest
type WithdrawFromShopRequest struct {
	// A valid EVM based address to withdraw to. For example, `0x39cb70F972E0EE920088AeF97Dbe5c6251a9c25D`.
	Address *string `json:"address,omitempty"`
	// Any wallet id within the MetaFab ecosystem to withdraw to.
	WalletId *string `json:"walletId,omitempty"`
	// The address of the currency (ERC20) token to withdraw from the shop. If no currencyAddress or currencyId, and no collectionAddress or collectionId are provided, the native token held by the shop will be withdrawn.
	CurrencyAddress *string `json:"currencyAddress,omitempty"`
	// A valid MetaFab currency id that represents the currency token to withdraw from the shop. `currencyAddress` or `currencyId` can be provided when withdrawing currency.
	CurrencyId *string `json:"currencyId,omitempty"`
	// The address of the collection (ERC1155) for the items to withdraw from the shop. If no currencyAddress and no collectionAddress is provided, the native token held by the shop will be withdrawn.
	CollectionAddress *string `json:"collectionAddress,omitempty"`
	// A valid MetaFab collection id that represents the collection for the items to withdraw from the shop. `collectionAddress` or `collectionId` can be provided when withdrawing items.
	CollectionId *string `json:"collectionId,omitempty"`
	// The specific itemIds of the provided collection to withdraw from the shop.
	ItemIds []int32 `json:"itemIds,omitempty"`
}

// NewWithdrawFromShopRequest instantiates a new WithdrawFromShopRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawFromShopRequest() *WithdrawFromShopRequest {
	this := WithdrawFromShopRequest{}
	return &this
}

// NewWithdrawFromShopRequestWithDefaults instantiates a new WithdrawFromShopRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawFromShopRequestWithDefaults() *WithdrawFromShopRequest {
	this := WithdrawFromShopRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *WithdrawFromShopRequest) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawFromShopRequest) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *WithdrawFromShopRequest) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *WithdrawFromShopRequest) SetAddress(v string) {
	o.Address = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *WithdrawFromShopRequest) GetWalletId() string {
	if o == nil || o.WalletId == nil {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawFromShopRequest) GetWalletIdOk() (*string, bool) {
	if o == nil || o.WalletId == nil {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *WithdrawFromShopRequest) HasWalletId() bool {
	if o != nil && o.WalletId != nil {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *WithdrawFromShopRequest) SetWalletId(v string) {
	o.WalletId = &v
}

// GetCurrencyAddress returns the CurrencyAddress field value if set, zero value otherwise.
func (o *WithdrawFromShopRequest) GetCurrencyAddress() string {
	if o == nil || o.CurrencyAddress == nil {
		var ret string
		return ret
	}
	return *o.CurrencyAddress
}

// GetCurrencyAddressOk returns a tuple with the CurrencyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawFromShopRequest) GetCurrencyAddressOk() (*string, bool) {
	if o == nil || o.CurrencyAddress == nil {
		return nil, false
	}
	return o.CurrencyAddress, true
}

// HasCurrencyAddress returns a boolean if a field has been set.
func (o *WithdrawFromShopRequest) HasCurrencyAddress() bool {
	if o != nil && o.CurrencyAddress != nil {
		return true
	}

	return false
}

// SetCurrencyAddress gets a reference to the given string and assigns it to the CurrencyAddress field.
func (o *WithdrawFromShopRequest) SetCurrencyAddress(v string) {
	o.CurrencyAddress = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *WithdrawFromShopRequest) GetCurrencyId() string {
	if o == nil || o.CurrencyId == nil {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawFromShopRequest) GetCurrencyIdOk() (*string, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *WithdrawFromShopRequest) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *WithdrawFromShopRequest) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCollectionAddress returns the CollectionAddress field value if set, zero value otherwise.
func (o *WithdrawFromShopRequest) GetCollectionAddress() string {
	if o == nil || o.CollectionAddress == nil {
		var ret string
		return ret
	}
	return *o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawFromShopRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil || o.CollectionAddress == nil {
		return nil, false
	}
	return o.CollectionAddress, true
}

// HasCollectionAddress returns a boolean if a field has been set.
func (o *WithdrawFromShopRequest) HasCollectionAddress() bool {
	if o != nil && o.CollectionAddress != nil {
		return true
	}

	return false
}

// SetCollectionAddress gets a reference to the given string and assigns it to the CollectionAddress field.
func (o *WithdrawFromShopRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *WithdrawFromShopRequest) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawFromShopRequest) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *WithdrawFromShopRequest) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *WithdrawFromShopRequest) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise.
func (o *WithdrawFromShopRequest) GetItemIds() []int32 {
	if o == nil || o.ItemIds == nil {
		var ret []int32
		return ret
	}
	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawFromShopRequest) GetItemIdsOk() ([]int32, bool) {
	if o == nil || o.ItemIds == nil {
		return nil, false
	}
	return o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *WithdrawFromShopRequest) HasItemIds() bool {
	if o != nil && o.ItemIds != nil {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []int32 and assigns it to the ItemIds field.
func (o *WithdrawFromShopRequest) SetItemIds(v []int32) {
	o.ItemIds = v
}

func (o WithdrawFromShopRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.WalletId != nil {
		toSerialize["walletId"] = o.WalletId
	}
	if o.CurrencyAddress != nil {
		toSerialize["currencyAddress"] = o.CurrencyAddress
	}
	if o.CurrencyId != nil {
		toSerialize["currencyId"] = o.CurrencyId
	}
	if o.CollectionAddress != nil {
		toSerialize["collectionAddress"] = o.CollectionAddress
	}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.ItemIds != nil {
		toSerialize["itemIds"] = o.ItemIds
	}
	return json.Marshal(toSerialize)
}

type NullableWithdrawFromShopRequest struct {
	value *WithdrawFromShopRequest
	isSet bool
}

func (v NullableWithdrawFromShopRequest) Get() *WithdrawFromShopRequest {
	return v.value
}

func (v *NullableWithdrawFromShopRequest) Set(val *WithdrawFromShopRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawFromShopRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawFromShopRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawFromShopRequest(val *WithdrawFromShopRequest) *NullableWithdrawFromShopRequest {
	return &NullableWithdrawFromShopRequest{value: val, isSet: true}
}

func (v NullableWithdrawFromShopRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawFromShopRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

