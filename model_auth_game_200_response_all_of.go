/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.4.0
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// AuthGame200ResponseAllOf struct for AuthGame200ResponseAllOf
type AuthGame200ResponseAllOf struct {
	Wallet *WalletModel `json:"wallet,omitempty"`
}

// NewAuthGame200ResponseAllOf instantiates a new AuthGame200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthGame200ResponseAllOf() *AuthGame200ResponseAllOf {
	this := AuthGame200ResponseAllOf{}
	return &this
}

// NewAuthGame200ResponseAllOfWithDefaults instantiates a new AuthGame200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthGame200ResponseAllOfWithDefaults() *AuthGame200ResponseAllOf {
	this := AuthGame200ResponseAllOf{}
	return &this
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *AuthGame200ResponseAllOf) GetWallet() WalletModel {
	if o == nil || o.Wallet == nil {
		var ret WalletModel
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200ResponseAllOf) GetWalletOk() (*WalletModel, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *AuthGame200ResponseAllOf) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given WalletModel and assigns it to the Wallet field.
func (o *AuthGame200ResponseAllOf) SetWallet(v WalletModel) {
	o.Wallet = &v
}

func (o AuthGame200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	return json.Marshal(toSerialize)
}

type NullableAuthGame200ResponseAllOf struct {
	value *AuthGame200ResponseAllOf
	isSet bool
}

func (v NullableAuthGame200ResponseAllOf) Get() *AuthGame200ResponseAllOf {
	return v.value
}

func (v *NullableAuthGame200ResponseAllOf) Set(val *AuthGame200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthGame200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthGame200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthGame200ResponseAllOf(val *AuthGame200ResponseAllOf) *NullableAuthGame200ResponseAllOf {
	return &NullableAuthGame200ResponseAllOf{value: val, isSet: true}
}

func (v NullableAuthGame200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthGame200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


