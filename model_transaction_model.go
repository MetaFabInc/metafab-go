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

// TransactionModel struct for TransactionModel
type TransactionModel struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	ContractId *string `json:"contractId,omitempty"`
	// This field has not had a description added.
	WalletId *string `json:"walletId,omitempty"`
	// This field has not had a description added.
	Function *string `json:"function,omitempty"`
	// This field has not had a description added.
	Args map[string]interface{} `json:"args,omitempty"`
	// This field has not had a description added.
	Hash *string `json:"hash,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewTransactionModel instantiates a new TransactionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionModel() *TransactionModel {
	this := TransactionModel{}
	return &this
}

// NewTransactionModelWithDefaults instantiates a new TransactionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionModelWithDefaults() *TransactionModel {
	this := TransactionModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransactionModel) SetId(v string) {
	o.Id = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *TransactionModel) GetContractId() string {
	if o == nil || o.ContractId == nil {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionModel) GetContractIdOk() (*string, bool) {
	if o == nil || o.ContractId == nil {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *TransactionModel) HasContractId() bool {
	if o != nil && o.ContractId != nil {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *TransactionModel) SetContractId(v string) {
	o.ContractId = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *TransactionModel) GetWalletId() string {
	if o == nil || o.WalletId == nil {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionModel) GetWalletIdOk() (*string, bool) {
	if o == nil || o.WalletId == nil {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *TransactionModel) HasWalletId() bool {
	if o != nil && o.WalletId != nil {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *TransactionModel) SetWalletId(v string) {
	o.WalletId = &v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *TransactionModel) GetFunction() string {
	if o == nil || o.Function == nil {
		var ret string
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionModel) GetFunctionOk() (*string, bool) {
	if o == nil || o.Function == nil {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *TransactionModel) HasFunction() bool {
	if o != nil && o.Function != nil {
		return true
	}

	return false
}

// SetFunction gets a reference to the given string and assigns it to the Function field.
func (o *TransactionModel) SetFunction(v string) {
	o.Function = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *TransactionModel) GetArgs() map[string]interface{} {
	if o == nil || o.Args == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionModel) GetArgsOk() (map[string]interface{}, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *TransactionModel) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given map[string]interface{} and assigns it to the Args field.
func (o *TransactionModel) SetArgs(v map[string]interface{}) {
	o.Args = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *TransactionModel) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionModel) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *TransactionModel) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *TransactionModel) SetHash(v string) {
	o.Hash = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TransactionModel) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionModel) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TransactionModel) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TransactionModel) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransactionModel) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransactionModel) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TransactionModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o TransactionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ContractId != nil {
		toSerialize["contractId"] = o.ContractId
	}
	if o.WalletId != nil {
		toSerialize["walletId"] = o.WalletId
	}
	if o.Function != nil {
		toSerialize["function"] = o.Function
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionModel struct {
	value *TransactionModel
	isSet bool
}

func (v NullableTransactionModel) Get() *TransactionModel {
	return v.value
}

func (v *NullableTransactionModel) Set(val *TransactionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionModel(val *TransactionModel) *NullableTransactionModel {
	return &NullableTransactionModel{value: val, isSet: true}
}

func (v NullableTransactionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

