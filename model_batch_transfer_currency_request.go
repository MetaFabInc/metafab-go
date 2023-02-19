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

// BatchTransferCurrencyRequest struct for BatchTransferCurrencyRequest
type BatchTransferCurrencyRequest struct {
	// An array of valid EVM based addresses to transfer currency to.
	Addresses []string `json:"addresses,omitempty"`
	// An array of wallet ids within the MetaFab ecosystem to transfer currency to.
	WalletIds []string `json:"walletIds,omitempty"`
	// An array of currency amounts to transfer. Ordering corresponds to the ordering of provided `addresses` and/or `walletIds`. If both `addresses` and `walletIds` are provided, `addresses` are first in the order.
	Amounts []float32 `json:"amounts"`
	// An optional array of uint256 numbers to reference each transfer in the batch. Ordering corresponds to the ordering of provided `addresses` or `walletIds`. If both `addresses` and `walletIds` are provided, `addresses` are first in the order.
	References []float32 `json:"references,omitempty"`
}

// NewBatchTransferCurrencyRequest instantiates a new BatchTransferCurrencyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchTransferCurrencyRequest(amounts []float32) *BatchTransferCurrencyRequest {
	this := BatchTransferCurrencyRequest{}
	this.Amounts = amounts
	return &this
}

// NewBatchTransferCurrencyRequestWithDefaults instantiates a new BatchTransferCurrencyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchTransferCurrencyRequestWithDefaults() *BatchTransferCurrencyRequest {
	this := BatchTransferCurrencyRequest{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *BatchTransferCurrencyRequest) GetAddresses() []string {
	if o == nil || isNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchTransferCurrencyRequest) GetAddressesOk() ([]string, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *BatchTransferCurrencyRequest) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *BatchTransferCurrencyRequest) SetAddresses(v []string) {
	o.Addresses = v
}

// GetWalletIds returns the WalletIds field value if set, zero value otherwise.
func (o *BatchTransferCurrencyRequest) GetWalletIds() []string {
	if o == nil || isNil(o.WalletIds) {
		var ret []string
		return ret
	}
	return o.WalletIds
}

// GetWalletIdsOk returns a tuple with the WalletIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchTransferCurrencyRequest) GetWalletIdsOk() ([]string, bool) {
	if o == nil || isNil(o.WalletIds) {
    return nil, false
	}
	return o.WalletIds, true
}

// HasWalletIds returns a boolean if a field has been set.
func (o *BatchTransferCurrencyRequest) HasWalletIds() bool {
	if o != nil && !isNil(o.WalletIds) {
		return true
	}

	return false
}

// SetWalletIds gets a reference to the given []string and assigns it to the WalletIds field.
func (o *BatchTransferCurrencyRequest) SetWalletIds(v []string) {
	o.WalletIds = v
}

// GetAmounts returns the Amounts field value
func (o *BatchTransferCurrencyRequest) GetAmounts() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value
// and a boolean to check if the value has been set.
func (o *BatchTransferCurrencyRequest) GetAmountsOk() ([]float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Amounts, true
}

// SetAmounts sets field value
func (o *BatchTransferCurrencyRequest) SetAmounts(v []float32) {
	o.Amounts = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *BatchTransferCurrencyRequest) GetReferences() []float32 {
	if o == nil || isNil(o.References) {
		var ret []float32
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchTransferCurrencyRequest) GetReferencesOk() ([]float32, bool) {
	if o == nil || isNil(o.References) {
    return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *BatchTransferCurrencyRequest) HasReferences() bool {
	if o != nil && !isNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []float32 and assigns it to the References field.
func (o *BatchTransferCurrencyRequest) SetReferences(v []float32) {
	o.References = v
}

func (o BatchTransferCurrencyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !isNil(o.WalletIds) {
		toSerialize["walletIds"] = o.WalletIds
	}
	if true {
		toSerialize["amounts"] = o.Amounts
	}
	if !isNil(o.References) {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableBatchTransferCurrencyRequest struct {
	value *BatchTransferCurrencyRequest
	isSet bool
}

func (v NullableBatchTransferCurrencyRequest) Get() *BatchTransferCurrencyRequest {
	return v.value
}

func (v *NullableBatchTransferCurrencyRequest) Set(val *BatchTransferCurrencyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchTransferCurrencyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchTransferCurrencyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchTransferCurrencyRequest(val *BatchTransferCurrencyRequest) *NullableBatchTransferCurrencyRequest {
	return &NullableBatchTransferCurrencyRequest{value: val, isSet: true}
}

func (v NullableBatchTransferCurrencyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchTransferCurrencyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


