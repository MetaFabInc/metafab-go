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

// CreateCollection200ResponseAllOfContract struct for CreateCollection200ResponseAllOfContract
type CreateCollection200ResponseAllOfContract struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	GameId *string `json:"gameId,omitempty"`
	// This field has not had a description added.
	Chain *string `json:"chain,omitempty"`
	// This field has not had a description added.
	Abi map[string]interface{} `json:"abi,omitempty"`
	// This field has not had a description added.
	Type *string `json:"type,omitempty"`
	// This field has not had a description added.
	Address *string `json:"address,omitempty"`
	// This field has not had a description added.
	ForwarderAddress *string `json:"forwarderAddress,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
	Transactions []TransactionModel `json:"transactions,omitempty"`
}

// NewCreateCollection200ResponseAllOfContract instantiates a new CreateCollection200ResponseAllOfContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollection200ResponseAllOfContract() *CreateCollection200ResponseAllOfContract {
	this := CreateCollection200ResponseAllOfContract{}
	return &this
}

// NewCreateCollection200ResponseAllOfContractWithDefaults instantiates a new CreateCollection200ResponseAllOfContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollection200ResponseAllOfContractWithDefaults() *CreateCollection200ResponseAllOfContract {
	this := CreateCollection200ResponseAllOfContract{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateCollection200ResponseAllOfContract) SetId(v string) {
	o.Id = &v
}

// GetGameId returns the GameId field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetGameId() string {
	if o == nil || isNil(o.GameId) {
		var ret string
		return ret
	}
	return *o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetGameIdOk() (*string, bool) {
	if o == nil || isNil(o.GameId) {
    return nil, false
	}
	return o.GameId, true
}

// HasGameId returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasGameId() bool {
	if o != nil && !isNil(o.GameId) {
		return true
	}

	return false
}

// SetGameId gets a reference to the given string and assigns it to the GameId field.
func (o *CreateCollection200ResponseAllOfContract) SetGameId(v string) {
	o.GameId = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetChain() string {
	if o == nil || isNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetChainOk() (*string, bool) {
	if o == nil || isNil(o.Chain) {
    return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasChain() bool {
	if o != nil && !isNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *CreateCollection200ResponseAllOfContract) SetChain(v string) {
	o.Chain = &v
}

// GetAbi returns the Abi field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetAbi() map[string]interface{} {
	if o == nil || isNil(o.Abi) {
		var ret map[string]interface{}
		return ret
	}
	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetAbiOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Abi) {
    return map[string]interface{}{}, false
	}
	return o.Abi, true
}

// HasAbi returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasAbi() bool {
	if o != nil && !isNil(o.Abi) {
		return true
	}

	return false
}

// SetAbi gets a reference to the given map[string]interface{} and assigns it to the Abi field.
func (o *CreateCollection200ResponseAllOfContract) SetAbi(v map[string]interface{}) {
	o.Abi = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateCollection200ResponseAllOfContract) SetType(v string) {
	o.Type = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CreateCollection200ResponseAllOfContract) SetAddress(v string) {
	o.Address = &v
}

// GetForwarderAddress returns the ForwarderAddress field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetForwarderAddress() string {
	if o == nil || isNil(o.ForwarderAddress) {
		var ret string
		return ret
	}
	return *o.ForwarderAddress
}

// GetForwarderAddressOk returns a tuple with the ForwarderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetForwarderAddressOk() (*string, bool) {
	if o == nil || isNil(o.ForwarderAddress) {
    return nil, false
	}
	return o.ForwarderAddress, true
}

// HasForwarderAddress returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasForwarderAddress() bool {
	if o != nil && !isNil(o.ForwarderAddress) {
		return true
	}

	return false
}

// SetForwarderAddress gets a reference to the given string and assigns it to the ForwarderAddress field.
func (o *CreateCollection200ResponseAllOfContract) SetForwarderAddress(v string) {
	o.ForwarderAddress = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CreateCollection200ResponseAllOfContract) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CreateCollection200ResponseAllOfContract) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOfContract) GetTransactions() []TransactionModel {
	if o == nil || isNil(o.Transactions) {
		var ret []TransactionModel
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOfContract) GetTransactionsOk() ([]TransactionModel, bool) {
	if o == nil || isNil(o.Transactions) {
    return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOfContract) HasTransactions() bool {
	if o != nil && !isNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []TransactionModel and assigns it to the Transactions field.
func (o *CreateCollection200ResponseAllOfContract) SetTransactions(v []TransactionModel) {
	o.Transactions = v
}

func (o CreateCollection200ResponseAllOfContract) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.GameId) {
		toSerialize["gameId"] = o.GameId
	}
	if !isNil(o.Chain) {
		toSerialize["chain"] = o.Chain
	}
	if !isNil(o.Abi) {
		toSerialize["abi"] = o.Abi
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.ForwarderAddress) {
		toSerialize["forwarderAddress"] = o.ForwarderAddress
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCollection200ResponseAllOfContract struct {
	value *CreateCollection200ResponseAllOfContract
	isSet bool
}

func (v NullableCreateCollection200ResponseAllOfContract) Get() *CreateCollection200ResponseAllOfContract {
	return v.value
}

func (v *NullableCreateCollection200ResponseAllOfContract) Set(val *CreateCollection200ResponseAllOfContract) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollection200ResponseAllOfContract) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollection200ResponseAllOfContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollection200ResponseAllOfContract(val *CreateCollection200ResponseAllOfContract) *NullableCreateCollection200ResponseAllOfContract {
	return &NullableCreateCollection200ResponseAllOfContract{value: val, isSet: true}
}

func (v NullableCreateCollection200ResponseAllOfContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollection200ResponseAllOfContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


