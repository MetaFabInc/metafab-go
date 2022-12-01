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

// WalletModel struct for WalletModel
type WalletModel struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	Address *string `json:"address,omitempty"`
}

// NewWalletModel instantiates a new WalletModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletModel() *WalletModel {
	this := WalletModel{}
	return &this
}

// NewWalletModelWithDefaults instantiates a new WalletModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletModelWithDefaults() *WalletModel {
	this := WalletModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WalletModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WalletModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WalletModel) SetId(v string) {
	o.Id = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *WalletModel) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletModel) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *WalletModel) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *WalletModel) SetAddress(v string) {
	o.Address = &v
}

func (o WalletModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableWalletModel struct {
	value *WalletModel
	isSet bool
}

func (v NullableWalletModel) Get() *WalletModel {
	return v.value
}

func (v *NullableWalletModel) Set(val *WalletModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletModel(val *WalletModel) *NullableWalletModel {
	return &NullableWalletModel{value: val, isSet: true}
}

func (v NullableWalletModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


