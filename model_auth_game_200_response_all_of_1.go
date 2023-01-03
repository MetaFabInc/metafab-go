/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.4.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// AuthGame200ResponseAllOf1 struct for AuthGame200ResponseAllOf1
type AuthGame200ResponseAllOf1 struct {
	FundingWallet *WalletModel `json:"fundingWallet,omitempty"`
}

// NewAuthGame200ResponseAllOf1 instantiates a new AuthGame200ResponseAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthGame200ResponseAllOf1() *AuthGame200ResponseAllOf1 {
	this := AuthGame200ResponseAllOf1{}
	return &this
}

// NewAuthGame200ResponseAllOf1WithDefaults instantiates a new AuthGame200ResponseAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthGame200ResponseAllOf1WithDefaults() *AuthGame200ResponseAllOf1 {
	this := AuthGame200ResponseAllOf1{}
	return &this
}

// GetFundingWallet returns the FundingWallet field value if set, zero value otherwise.
func (o *AuthGame200ResponseAllOf1) GetFundingWallet() WalletModel {
	if o == nil || o.FundingWallet == nil {
		var ret WalletModel
		return ret
	}
	return *o.FundingWallet
}

// GetFundingWalletOk returns a tuple with the FundingWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200ResponseAllOf1) GetFundingWalletOk() (*WalletModel, bool) {
	if o == nil || o.FundingWallet == nil {
		return nil, false
	}
	return o.FundingWallet, true
}

// HasFundingWallet returns a boolean if a field has been set.
func (o *AuthGame200ResponseAllOf1) HasFundingWallet() bool {
	if o != nil && o.FundingWallet != nil {
		return true
	}

	return false
}

// SetFundingWallet gets a reference to the given WalletModel and assigns it to the FundingWallet field.
func (o *AuthGame200ResponseAllOf1) SetFundingWallet(v WalletModel) {
	o.FundingWallet = &v
}

func (o AuthGame200ResponseAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FundingWallet != nil {
		toSerialize["fundingWallet"] = o.FundingWallet
	}
	return json.Marshal(toSerialize)
}

type NullableAuthGame200ResponseAllOf1 struct {
	value *AuthGame200ResponseAllOf1
	isSet bool
}

func (v NullableAuthGame200ResponseAllOf1) Get() *AuthGame200ResponseAllOf1 {
	return v.value
}

func (v *NullableAuthGame200ResponseAllOf1) Set(val *AuthGame200ResponseAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthGame200ResponseAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthGame200ResponseAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthGame200ResponseAllOf1(val *AuthGame200ResponseAllOf1) *NullableAuthGame200ResponseAllOf1 {
	return &NullableAuthGame200ResponseAllOf1{value: val, isSet: true}
}

func (v NullableAuthGame200ResponseAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthGame200ResponseAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


