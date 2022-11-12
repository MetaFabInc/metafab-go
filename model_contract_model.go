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

// ContractModel struct for ContractModel
type ContractModel struct {
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
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewContractModel instantiates a new ContractModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractModel() *ContractModel {
	this := ContractModel{}
	return &this
}

// NewContractModelWithDefaults instantiates a new ContractModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractModelWithDefaults() *ContractModel {
	this := ContractModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContractModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContractModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContractModel) SetId(v string) {
	o.Id = &v
}

// GetGameId returns the GameId field value if set, zero value otherwise.
func (o *ContractModel) GetGameId() string {
	if o == nil || o.GameId == nil {
		var ret string
		return ret
	}
	return *o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractModel) GetGameIdOk() (*string, bool) {
	if o == nil || o.GameId == nil {
		return nil, false
	}
	return o.GameId, true
}

// HasGameId returns a boolean if a field has been set.
func (o *ContractModel) HasGameId() bool {
	if o != nil && o.GameId != nil {
		return true
	}

	return false
}

// SetGameId gets a reference to the given string and assigns it to the GameId field.
func (o *ContractModel) SetGameId(v string) {
	o.GameId = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *ContractModel) GetChain() string {
	if o == nil || o.Chain == nil {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractModel) GetChainOk() (*string, bool) {
	if o == nil || o.Chain == nil {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *ContractModel) HasChain() bool {
	if o != nil && o.Chain != nil {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *ContractModel) SetChain(v string) {
	o.Chain = &v
}

// GetAbi returns the Abi field value if set, zero value otherwise.
func (o *ContractModel) GetAbi() map[string]interface{} {
	if o == nil || o.Abi == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractModel) GetAbiOk() (map[string]interface{}, bool) {
	if o == nil || o.Abi == nil {
		return nil, false
	}
	return o.Abi, true
}

// HasAbi returns a boolean if a field has been set.
func (o *ContractModel) HasAbi() bool {
	if o != nil && o.Abi != nil {
		return true
	}

	return false
}

// SetAbi gets a reference to the given map[string]interface{} and assigns it to the Abi field.
func (o *ContractModel) SetAbi(v map[string]interface{}) {
	o.Abi = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContractModel) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractModel) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContractModel) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContractModel) SetType(v string) {
	o.Type = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ContractModel) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractModel) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ContractModel) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ContractModel) SetAddress(v string) {
	o.Address = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ContractModel) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractModel) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ContractModel) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ContractModel) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ContractModel) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ContractModel) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ContractModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o ContractModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.GameId != nil {
		toSerialize["gameId"] = o.GameId
	}
	if o.Chain != nil {
		toSerialize["chain"] = o.Chain
	}
	if o.Abi != nil {
		toSerialize["abi"] = o.Abi
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableContractModel struct {
	value *ContractModel
	isSet bool
}

func (v NullableContractModel) Get() *ContractModel {
	return v.value
}

func (v *NullableContractModel) Set(val *ContractModel) {
	v.value = val
	v.isSet = true
}

func (v NullableContractModel) IsSet() bool {
	return v.isSet
}

func (v *NullableContractModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractModel(val *ContractModel) *NullableContractModel {
	return &NullableContractModel{value: val, isSet: true}
}

func (v NullableContractModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

