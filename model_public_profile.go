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

// PublicProfile struct for PublicProfile
type PublicProfile struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	WalletId *string `json:"walletId,omitempty"`
	// This field has not had a description added.
	Username *string `json:"username,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
	CustodialWallet *PublicPlayerCustodialWallet `json:"custodialWallet,omitempty"`
	Wallet *PublicPlayerCustodialWallet `json:"wallet,omitempty"`
}

// NewPublicProfile instantiates a new PublicProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicProfile() *PublicProfile {
	this := PublicProfile{}
	return &this
}

// NewPublicProfileWithDefaults instantiates a new PublicProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicProfileWithDefaults() *PublicProfile {
	this := PublicProfile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicProfile) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfile) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicProfile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicProfile) SetId(v string) {
	o.Id = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *PublicProfile) GetWalletId() string {
	if o == nil || isNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfile) GetWalletIdOk() (*string, bool) {
	if o == nil || isNil(o.WalletId) {
    return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *PublicProfile) HasWalletId() bool {
	if o != nil && !isNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *PublicProfile) SetWalletId(v string) {
	o.WalletId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PublicProfile) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfile) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PublicProfile) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PublicProfile) SetUsername(v string) {
	o.Username = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PublicProfile) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfile) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PublicProfile) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PublicProfile) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PublicProfile) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfile) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PublicProfile) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PublicProfile) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCustodialWallet returns the CustodialWallet field value if set, zero value otherwise.
func (o *PublicProfile) GetCustodialWallet() PublicPlayerCustodialWallet {
	if o == nil || isNil(o.CustodialWallet) {
		var ret PublicPlayerCustodialWallet
		return ret
	}
	return *o.CustodialWallet
}

// GetCustodialWalletOk returns a tuple with the CustodialWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfile) GetCustodialWalletOk() (*PublicPlayerCustodialWallet, bool) {
	if o == nil || isNil(o.CustodialWallet) {
    return nil, false
	}
	return o.CustodialWallet, true
}

// HasCustodialWallet returns a boolean if a field has been set.
func (o *PublicProfile) HasCustodialWallet() bool {
	if o != nil && !isNil(o.CustodialWallet) {
		return true
	}

	return false
}

// SetCustodialWallet gets a reference to the given PublicPlayerCustodialWallet and assigns it to the CustodialWallet field.
func (o *PublicProfile) SetCustodialWallet(v PublicPlayerCustodialWallet) {
	o.CustodialWallet = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *PublicProfile) GetWallet() PublicPlayerCustodialWallet {
	if o == nil || isNil(o.Wallet) {
		var ret PublicPlayerCustodialWallet
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicProfile) GetWalletOk() (*PublicPlayerCustodialWallet, bool) {
	if o == nil || isNil(o.Wallet) {
    return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *PublicProfile) HasWallet() bool {
	if o != nil && !isNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given PublicPlayerCustodialWallet and assigns it to the Wallet field.
func (o *PublicProfile) SetWallet(v PublicPlayerCustodialWallet) {
	o.Wallet = &v
}

func (o PublicProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.WalletId) {
		toSerialize["walletId"] = o.WalletId
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CustodialWallet) {
		toSerialize["custodialWallet"] = o.CustodialWallet
	}
	if !isNil(o.Wallet) {
		toSerialize["wallet"] = o.Wallet
	}
	return json.Marshal(toSerialize)
}

type NullablePublicProfile struct {
	value *PublicProfile
	isSet bool
}

func (v NullablePublicProfile) Get() *PublicProfile {
	return v.value
}

func (v *NullablePublicProfile) Set(val *PublicProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicProfile(val *PublicProfile) *NullablePublicProfile {
	return &NullablePublicProfile{value: val, isSet: true}
}

func (v NullablePublicProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

