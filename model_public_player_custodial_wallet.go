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

// PublicPlayerCustodialWallet struct for PublicPlayerCustodialWallet
type PublicPlayerCustodialWallet struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	Address *string `json:"address,omitempty"`
}

// NewPublicPlayerCustodialWallet instantiates a new PublicPlayerCustodialWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicPlayerCustodialWallet() *PublicPlayerCustodialWallet {
	this := PublicPlayerCustodialWallet{}
	return &this
}

// NewPublicPlayerCustodialWalletWithDefaults instantiates a new PublicPlayerCustodialWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicPlayerCustodialWalletWithDefaults() *PublicPlayerCustodialWallet {
	this := PublicPlayerCustodialWallet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicPlayerCustodialWallet) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayerCustodialWallet) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicPlayerCustodialWallet) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicPlayerCustodialWallet) SetId(v string) {
	o.Id = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PublicPlayerCustodialWallet) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayerCustodialWallet) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PublicPlayerCustodialWallet) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *PublicPlayerCustodialWallet) SetAddress(v string) {
	o.Address = &v
}

func (o PublicPlayerCustodialWallet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullablePublicPlayerCustodialWallet struct {
	value *PublicPlayerCustodialWallet
	isSet bool
}

func (v NullablePublicPlayerCustodialWallet) Get() *PublicPlayerCustodialWallet {
	return v.value
}

func (v *NullablePublicPlayerCustodialWallet) Set(val *PublicPlayerCustodialWallet) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPlayerCustodialWallet) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPlayerCustodialWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPlayerCustodialWallet(val *PublicPlayerCustodialWallet) *NullablePublicPlayerCustodialWallet {
	return &NullablePublicPlayerCustodialWallet{value: val, isSet: true}
}

func (v NullablePublicPlayerCustodialWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPlayerCustodialWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


