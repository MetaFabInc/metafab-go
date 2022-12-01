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

// CreateCollection200ResponseAllOfContractAllOf struct for CreateCollection200ResponseAllOfContractAllOf
type CreateCollection200ResponseAllOfContractAllOf struct {
	Transactions []TransactionModel `json:"transactions,omitempty"`
}

// NewCreateCollection200ResponseAllOfContractAllOf instantiates a new CreateCollection200ResponseAllOfContractAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollection200ResponseAllOfContractAllOf() *CreateCollection200ResponseAllOfContractAllOf {
	this := CreateCollection200ResponseAllOfContractAllOf{}
	return &this
}

// NewCreateCollection200ResponseAllOfContractAllOfWithDefaults instantiates a new CreateCollection200ResponseAllOfContractAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollection200ResponseAllOfContractAllOfWithDefaults() *CreateCollection200ResponseAllOfContractAllOf {
	this := CreateCollection200ResponseAllOfContractAllOf{}
	return &this
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContractAllOf) GetTransactions() []TransactionModel {
	if o == nil || o.Transactions == nil {
		var ret []TransactionModel
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContractAllOf) GetTransactionsOk() ([]TransactionModel, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContractAllOf) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []TransactionModel and assigns it to the Transactions field.
func (o *CreateCollection200ResponseAllOfContractAllOf) SetTransactions(v []TransactionModel) {
	o.Transactions = v
}

func (o CreateCollection200ResponseAllOfContractAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCollection200ResponseAllOfContractAllOf struct {
	value *CreateCollection200ResponseAllOfContractAllOf
	isSet bool
}

func (v NullableCreateCollection200ResponseAllOfContractAllOf) Get() *CreateCollection200ResponseAllOfContractAllOf {
	return v.value
}

func (v *NullableCreateCollection200ResponseAllOfContractAllOf) Set(val *CreateCollection200ResponseAllOfContractAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollection200ResponseAllOfContractAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollection200ResponseAllOfContractAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollection200ResponseAllOfContractAllOf(val *CreateCollection200ResponseAllOfContractAllOf) *NullableCreateCollection200ResponseAllOfContractAllOf {
	return &NullableCreateCollection200ResponseAllOfContractAllOf{value: val, isSet: true}
}

func (v NullableCreateCollection200ResponseAllOfContractAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollection200ResponseAllOfContractAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


