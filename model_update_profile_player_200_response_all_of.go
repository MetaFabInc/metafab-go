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

// UpdateProfilePlayer200ResponseAllOf struct for UpdateProfilePlayer200ResponseAllOf
type UpdateProfilePlayer200ResponseAllOf struct {
	Wallet *WalletModel `json:"wallet,omitempty"`
	CustodialWallet *WalletModel `json:"custodialWallet,omitempty"`
}

// NewUpdateProfilePlayer200ResponseAllOf instantiates a new UpdateProfilePlayer200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfilePlayer200ResponseAllOf() *UpdateProfilePlayer200ResponseAllOf {
	this := UpdateProfilePlayer200ResponseAllOf{}
	return &this
}

// NewUpdateProfilePlayer200ResponseAllOfWithDefaults instantiates a new UpdateProfilePlayer200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfilePlayer200ResponseAllOfWithDefaults() *UpdateProfilePlayer200ResponseAllOf {
	this := UpdateProfilePlayer200ResponseAllOf{}
	return &this
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200ResponseAllOf) GetWallet() WalletModel {
	if o == nil || isNil(o.Wallet) {
		var ret WalletModel
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200ResponseAllOf) GetWalletOk() (*WalletModel, bool) {
	if o == nil || isNil(o.Wallet) {
    return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200ResponseAllOf) HasWallet() bool {
	if o != nil && !isNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given WalletModel and assigns it to the Wallet field.
func (o *UpdateProfilePlayer200ResponseAllOf) SetWallet(v WalletModel) {
	o.Wallet = &v
}

// GetCustodialWallet returns the CustodialWallet field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200ResponseAllOf) GetCustodialWallet() WalletModel {
	if o == nil || isNil(o.CustodialWallet) {
		var ret WalletModel
		return ret
	}
	return *o.CustodialWallet
}

// GetCustodialWalletOk returns a tuple with the CustodialWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200ResponseAllOf) GetCustodialWalletOk() (*WalletModel, bool) {
	if o == nil || isNil(o.CustodialWallet) {
    return nil, false
	}
	return o.CustodialWallet, true
}

// HasCustodialWallet returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200ResponseAllOf) HasCustodialWallet() bool {
	if o != nil && !isNil(o.CustodialWallet) {
		return true
	}

	return false
}

// SetCustodialWallet gets a reference to the given WalletModel and assigns it to the CustodialWallet field.
func (o *UpdateProfilePlayer200ResponseAllOf) SetCustodialWallet(v WalletModel) {
	o.CustodialWallet = &v
}

func (o UpdateProfilePlayer200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wallet) {
		toSerialize["wallet"] = o.Wallet
	}
	if !isNil(o.CustodialWallet) {
		toSerialize["custodialWallet"] = o.CustodialWallet
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProfilePlayer200ResponseAllOf struct {
	value *UpdateProfilePlayer200ResponseAllOf
	isSet bool
}

func (v NullableUpdateProfilePlayer200ResponseAllOf) Get() *UpdateProfilePlayer200ResponseAllOf {
	return v.value
}

func (v *NullableUpdateProfilePlayer200ResponseAllOf) Set(val *UpdateProfilePlayer200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfilePlayer200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfilePlayer200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfilePlayer200ResponseAllOf(val *UpdateProfilePlayer200ResponseAllOf) *NullableUpdateProfilePlayer200ResponseAllOf {
	return &NullableUpdateProfilePlayer200ResponseAllOf{value: val, isSet: true}
}

func (v NullableUpdateProfilePlayer200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfilePlayer200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


