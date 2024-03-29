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

// TransferCollectionItemRequest struct for TransferCollectionItemRequest
type TransferCollectionItemRequest struct {
	// A valid EVM based addresses to transfer items to.
	Address *string `json:"address,omitempty"`
	// A wallet id within the MetaFab ecosystem to transfer items to.
	WalletId []string `json:"walletId,omitempty"`
	// The quantity of the collectionItemId to transfer.
	Quantity int32 `json:"quantity"`
}

// NewTransferCollectionItemRequest instantiates a new TransferCollectionItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCollectionItemRequest(quantity int32) *TransferCollectionItemRequest {
	this := TransferCollectionItemRequest{}
	this.Quantity = quantity
	return &this
}

// NewTransferCollectionItemRequestWithDefaults instantiates a new TransferCollectionItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCollectionItemRequestWithDefaults() *TransferCollectionItemRequest {
	this := TransferCollectionItemRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransferCollectionItemRequest) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCollectionItemRequest) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransferCollectionItemRequest) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TransferCollectionItemRequest) SetAddress(v string) {
	o.Address = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *TransferCollectionItemRequest) GetWalletId() []string {
	if o == nil || isNil(o.WalletId) {
		var ret []string
		return ret
	}
	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCollectionItemRequest) GetWalletIdOk() ([]string, bool) {
	if o == nil || isNil(o.WalletId) {
    return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *TransferCollectionItemRequest) HasWalletId() bool {
	if o != nil && !isNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given []string and assigns it to the WalletId field.
func (o *TransferCollectionItemRequest) SetWalletId(v []string) {
	o.WalletId = v
}

// GetQuantity returns the Quantity field value
func (o *TransferCollectionItemRequest) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *TransferCollectionItemRequest) GetQuantityOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *TransferCollectionItemRequest) SetQuantity(v int32) {
	o.Quantity = v
}

func (o TransferCollectionItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.WalletId) {
		toSerialize["walletId"] = o.WalletId
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableTransferCollectionItemRequest struct {
	value *TransferCollectionItemRequest
	isSet bool
}

func (v NullableTransferCollectionItemRequest) Get() *TransferCollectionItemRequest {
	return v.value
}

func (v *NullableTransferCollectionItemRequest) Set(val *TransferCollectionItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCollectionItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCollectionItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCollectionItemRequest(val *TransferCollectionItemRequest) *NullableTransferCollectionItemRequest {
	return &NullableTransferCollectionItemRequest{value: val, isSet: true}
}

func (v NullableTransferCollectionItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCollectionItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


