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

// UpdateGame200ResponseAllOf struct for UpdateGame200ResponseAllOf
type UpdateGame200ResponseAllOf struct {
	// This field has not had a description added.
	WalletDecryptKey *string `json:"walletDecryptKey,omitempty"`
}

// NewUpdateGame200ResponseAllOf instantiates a new UpdateGame200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGame200ResponseAllOf() *UpdateGame200ResponseAllOf {
	this := UpdateGame200ResponseAllOf{}
	return &this
}

// NewUpdateGame200ResponseAllOfWithDefaults instantiates a new UpdateGame200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGame200ResponseAllOfWithDefaults() *UpdateGame200ResponseAllOf {
	this := UpdateGame200ResponseAllOf{}
	return &this
}

// GetWalletDecryptKey returns the WalletDecryptKey field value if set, zero value otherwise.
func (o *UpdateGame200ResponseAllOf) GetWalletDecryptKey() string {
	if o == nil || isNil(o.WalletDecryptKey) {
		var ret string
		return ret
	}
	return *o.WalletDecryptKey
}

// GetWalletDecryptKeyOk returns a tuple with the WalletDecryptKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGame200ResponseAllOf) GetWalletDecryptKeyOk() (*string, bool) {
	if o == nil || isNil(o.WalletDecryptKey) {
    return nil, false
	}
	return o.WalletDecryptKey, true
}

// HasWalletDecryptKey returns a boolean if a field has been set.
func (o *UpdateGame200ResponseAllOf) HasWalletDecryptKey() bool {
	if o != nil && !isNil(o.WalletDecryptKey) {
		return true
	}

	return false
}

// SetWalletDecryptKey gets a reference to the given string and assigns it to the WalletDecryptKey field.
func (o *UpdateGame200ResponseAllOf) SetWalletDecryptKey(v string) {
	o.WalletDecryptKey = &v
}

func (o UpdateGame200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WalletDecryptKey) {
		toSerialize["walletDecryptKey"] = o.WalletDecryptKey
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateGame200ResponseAllOf struct {
	value *UpdateGame200ResponseAllOf
	isSet bool
}

func (v NullableUpdateGame200ResponseAllOf) Get() *UpdateGame200ResponseAllOf {
	return v.value
}

func (v *NullableUpdateGame200ResponseAllOf) Set(val *UpdateGame200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGame200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGame200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGame200ResponseAllOf(val *UpdateGame200ResponseAllOf) *NullableUpdateGame200ResponseAllOf {
	return &NullableUpdateGame200ResponseAllOf{value: val, isSet: true}
}

func (v NullableUpdateGame200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGame200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


