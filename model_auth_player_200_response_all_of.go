/*
MetaFab API

Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.5.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// AuthPlayer200ResponseAllOf struct for AuthPlayer200ResponseAllOf
type AuthPlayer200ResponseAllOf struct {
	// This field has not had a description added.
	WalletDecryptKey *string `json:"walletDecryptKey,omitempty"`
	Wallet *WalletModel `json:"wallet,omitempty"`
	CustodialWallet *WalletModel `json:"custodialWallet,omitempty"`
}

// NewAuthPlayer200ResponseAllOf instantiates a new AuthPlayer200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthPlayer200ResponseAllOf() *AuthPlayer200ResponseAllOf {
	this := AuthPlayer200ResponseAllOf{}
	return &this
}

// NewAuthPlayer200ResponseAllOfWithDefaults instantiates a new AuthPlayer200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthPlayer200ResponseAllOfWithDefaults() *AuthPlayer200ResponseAllOf {
	this := AuthPlayer200ResponseAllOf{}
	return &this
}

// GetWalletDecryptKey returns the WalletDecryptKey field value if set, zero value otherwise.
func (o *AuthPlayer200ResponseAllOf) GetWalletDecryptKey() string {
	if o == nil || isNil(o.WalletDecryptKey) {
		var ret string
		return ret
	}
	return *o.WalletDecryptKey
}

// GetWalletDecryptKeyOk returns a tuple with the WalletDecryptKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthPlayer200ResponseAllOf) GetWalletDecryptKeyOk() (*string, bool) {
	if o == nil || isNil(o.WalletDecryptKey) {
    return nil, false
	}
	return o.WalletDecryptKey, true
}

// HasWalletDecryptKey returns a boolean if a field has been set.
func (o *AuthPlayer200ResponseAllOf) HasWalletDecryptKey() bool {
	if o != nil && !isNil(o.WalletDecryptKey) {
		return true
	}

	return false
}

// SetWalletDecryptKey gets a reference to the given string and assigns it to the WalletDecryptKey field.
func (o *AuthPlayer200ResponseAllOf) SetWalletDecryptKey(v string) {
	o.WalletDecryptKey = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *AuthPlayer200ResponseAllOf) GetWallet() WalletModel {
	if o == nil || isNil(o.Wallet) {
		var ret WalletModel
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthPlayer200ResponseAllOf) GetWalletOk() (*WalletModel, bool) {
	if o == nil || isNil(o.Wallet) {
    return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *AuthPlayer200ResponseAllOf) HasWallet() bool {
	if o != nil && !isNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given WalletModel and assigns it to the Wallet field.
func (o *AuthPlayer200ResponseAllOf) SetWallet(v WalletModel) {
	o.Wallet = &v
}

// GetCustodialWallet returns the CustodialWallet field value if set, zero value otherwise.
func (o *AuthPlayer200ResponseAllOf) GetCustodialWallet() WalletModel {
	if o == nil || isNil(o.CustodialWallet) {
		var ret WalletModel
		return ret
	}
	return *o.CustodialWallet
}

// GetCustodialWalletOk returns a tuple with the CustodialWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthPlayer200ResponseAllOf) GetCustodialWalletOk() (*WalletModel, bool) {
	if o == nil || isNil(o.CustodialWallet) {
    return nil, false
	}
	return o.CustodialWallet, true
}

// HasCustodialWallet returns a boolean if a field has been set.
func (o *AuthPlayer200ResponseAllOf) HasCustodialWallet() bool {
	if o != nil && !isNil(o.CustodialWallet) {
		return true
	}

	return false
}

// SetCustodialWallet gets a reference to the given WalletModel and assigns it to the CustodialWallet field.
func (o *AuthPlayer200ResponseAllOf) SetCustodialWallet(v WalletModel) {
	o.CustodialWallet = &v
}

func (o AuthPlayer200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WalletDecryptKey) {
		toSerialize["walletDecryptKey"] = o.WalletDecryptKey
	}
	if !isNil(o.Wallet) {
		toSerialize["wallet"] = o.Wallet
	}
	if !isNil(o.CustodialWallet) {
		toSerialize["custodialWallet"] = o.CustodialWallet
	}
	return json.Marshal(toSerialize)
}

type NullableAuthPlayer200ResponseAllOf struct {
	value *AuthPlayer200ResponseAllOf
	isSet bool
}

func (v NullableAuthPlayer200ResponseAllOf) Get() *AuthPlayer200ResponseAllOf {
	return v.value
}

func (v *NullableAuthPlayer200ResponseAllOf) Set(val *AuthPlayer200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthPlayer200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthPlayer200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthPlayer200ResponseAllOf(val *AuthPlayer200ResponseAllOf) *NullableAuthPlayer200ResponseAllOf {
	return &NullableAuthPlayer200ResponseAllOf{value: val, isSet: true}
}

func (v NullableAuthPlayer200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthPlayer200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

