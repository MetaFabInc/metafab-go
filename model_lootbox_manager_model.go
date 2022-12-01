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

// LootboxManagerModel struct for LootboxManagerModel
type LootboxManagerModel struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	GameId *string `json:"gameId,omitempty"`
	// This field has not had a description added.
	ContractId *string `json:"contractId,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewLootboxManagerModel instantiates a new LootboxManagerModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLootboxManagerModel() *LootboxManagerModel {
	this := LootboxManagerModel{}
	return &this
}

// NewLootboxManagerModelWithDefaults instantiates a new LootboxManagerModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLootboxManagerModelWithDefaults() *LootboxManagerModel {
	this := LootboxManagerModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LootboxManagerModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LootboxManagerModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LootboxManagerModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LootboxManagerModel) SetId(v string) {
	o.Id = &v
}

// GetGameId returns the GameId field value if set, zero value otherwise.
func (o *LootboxManagerModel) GetGameId() string {
	if o == nil || o.GameId == nil {
		var ret string
		return ret
	}
	return *o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LootboxManagerModel) GetGameIdOk() (*string, bool) {
	if o == nil || o.GameId == nil {
		return nil, false
	}
	return o.GameId, true
}

// HasGameId returns a boolean if a field has been set.
func (o *LootboxManagerModel) HasGameId() bool {
	if o != nil && o.GameId != nil {
		return true
	}

	return false
}

// SetGameId gets a reference to the given string and assigns it to the GameId field.
func (o *LootboxManagerModel) SetGameId(v string) {
	o.GameId = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *LootboxManagerModel) GetContractId() string {
	if o == nil || o.ContractId == nil {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LootboxManagerModel) GetContractIdOk() (*string, bool) {
	if o == nil || o.ContractId == nil {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *LootboxManagerModel) HasContractId() bool {
	if o != nil && o.ContractId != nil {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *LootboxManagerModel) SetContractId(v string) {
	o.ContractId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LootboxManagerModel) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LootboxManagerModel) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LootboxManagerModel) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LootboxManagerModel) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LootboxManagerModel) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LootboxManagerModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LootboxManagerModel) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LootboxManagerModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o LootboxManagerModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.GameId != nil {
		toSerialize["gameId"] = o.GameId
	}
	if o.ContractId != nil {
		toSerialize["contractId"] = o.ContractId
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableLootboxManagerModel struct {
	value *LootboxManagerModel
	isSet bool
}

func (v NullableLootboxManagerModel) Get() *LootboxManagerModel {
	return v.value
}

func (v *NullableLootboxManagerModel) Set(val *LootboxManagerModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLootboxManagerModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLootboxManagerModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLootboxManagerModel(val *LootboxManagerModel) *NullableLootboxManagerModel {
	return &NullableLootboxManagerModel{value: val, isSet: true}
}

func (v NullableLootboxManagerModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLootboxManagerModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

