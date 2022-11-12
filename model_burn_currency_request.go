/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.2.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// BurnCurrencyRequest struct for BurnCurrencyRequest
type BurnCurrencyRequest struct {
	// The amount of currency to remove (burn). The currency balance of the authenticating game or player's wallet must be equal to or greater than this amount.
	Amount float32 `json:"amount"`
}

// NewBurnCurrencyRequest instantiates a new BurnCurrencyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnCurrencyRequest(amount float32) *BurnCurrencyRequest {
	this := BurnCurrencyRequest{}
	this.Amount = amount
	return &this
}

// NewBurnCurrencyRequestWithDefaults instantiates a new BurnCurrencyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnCurrencyRequestWithDefaults() *BurnCurrencyRequest {
	this := BurnCurrencyRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *BurnCurrencyRequest) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BurnCurrencyRequest) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BurnCurrencyRequest) SetAmount(v float32) {
	o.Amount = v
}

func (o BurnCurrencyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableBurnCurrencyRequest struct {
	value *BurnCurrencyRequest
	isSet bool
}

func (v NullableBurnCurrencyRequest) Get() *BurnCurrencyRequest {
	return v.value
}

func (v *NullableBurnCurrencyRequest) Set(val *BurnCurrencyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnCurrencyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnCurrencyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnCurrencyRequest(val *BurnCurrencyRequest) *NullableBurnCurrencyRequest {
	return &NullableBurnCurrencyRequest{value: val, isSet: true}
}

func (v NullableBurnCurrencyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnCurrencyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


