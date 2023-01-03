/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.4.0
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// GetCurrencies200ResponseInner struct for GetCurrencies200ResponseInner
type GetCurrencies200ResponseInner struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	GameId *string `json:"gameId,omitempty"`
	// This field has not had a description added.
	ContractId *string `json:"contractId,omitempty"`
	// This field has not had a description added.
	Name *string `json:"name,omitempty"`
	// This field has not had a description added.
	Symbol *string `json:"symbol,omitempty"`
	// This field has not had a description added.
	SupplyCap *int32 `json:"supplyCap,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
	Contract *ContractModel `json:"contract,omitempty"`
}

// NewGetCurrencies200ResponseInner instantiates a new GetCurrencies200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrencies200ResponseInner() *GetCurrencies200ResponseInner {
	this := GetCurrencies200ResponseInner{}
	return &this
}

// NewGetCurrencies200ResponseInnerWithDefaults instantiates a new GetCurrencies200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrencies200ResponseInnerWithDefaults() *GetCurrencies200ResponseInner {
	this := GetCurrencies200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetCurrencies200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetGameId returns the GameId field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetGameId() string {
	if o == nil || o.GameId == nil {
		var ret string
		return ret
	}
	return *o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetGameIdOk() (*string, bool) {
	if o == nil || o.GameId == nil {
		return nil, false
	}
	return o.GameId, true
}

// HasGameId returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasGameId() bool {
	if o != nil && o.GameId != nil {
		return true
	}

	return false
}

// SetGameId gets a reference to the given string and assigns it to the GameId field.
func (o *GetCurrencies200ResponseInner) SetGameId(v string) {
	o.GameId = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetContractId() string {
	if o == nil || o.ContractId == nil {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetContractIdOk() (*string, bool) {
	if o == nil || o.ContractId == nil {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasContractId() bool {
	if o != nil && o.ContractId != nil {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *GetCurrencies200ResponseInner) SetContractId(v string) {
	o.ContractId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetCurrencies200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetCurrencies200ResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSupplyCap returns the SupplyCap field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetSupplyCap() int32 {
	if o == nil || o.SupplyCap == nil {
		var ret int32
		return ret
	}
	return *o.SupplyCap
}

// GetSupplyCapOk returns a tuple with the SupplyCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetSupplyCapOk() (*int32, bool) {
	if o == nil || o.SupplyCap == nil {
		return nil, false
	}
	return o.SupplyCap, true
}

// HasSupplyCap returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasSupplyCap() bool {
	if o != nil && o.SupplyCap != nil {
		return true
	}

	return false
}

// SetSupplyCap gets a reference to the given int32 and assigns it to the SupplyCap field.
func (o *GetCurrencies200ResponseInner) SetSupplyCap(v int32) {
	o.SupplyCap = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GetCurrencies200ResponseInner) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetCurrencies200ResponseInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *GetCurrencies200ResponseInner) GetContract() ContractModel {
	if o == nil || o.Contract == nil {
		var ret ContractModel
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencies200ResponseInner) GetContractOk() (*ContractModel, bool) {
	if o == nil || o.Contract == nil {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *GetCurrencies200ResponseInner) HasContract() bool {
	if o != nil && o.Contract != nil {
		return true
	}

	return false
}

// SetContract gets a reference to the given ContractModel and assigns it to the Contract field.
func (o *GetCurrencies200ResponseInner) SetContract(v ContractModel) {
	o.Contract = &v
}

func (o GetCurrencies200ResponseInner) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.SupplyCap != nil {
		toSerialize["supplyCap"] = o.SupplyCap
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Contract != nil {
		toSerialize["contract"] = o.Contract
	}
	return json.Marshal(toSerialize)
}

type NullableGetCurrencies200ResponseInner struct {
	value *GetCurrencies200ResponseInner
	isSet bool
}

func (v NullableGetCurrencies200ResponseInner) Get() *GetCurrencies200ResponseInner {
	return v.value
}

func (v *NullableGetCurrencies200ResponseInner) Set(val *GetCurrencies200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrencies200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrencies200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrencies200ResponseInner(val *GetCurrencies200ResponseInner) *NullableGetCurrencies200ResponseInner {
	return &NullableGetCurrencies200ResponseInner{value: val, isSet: true}
}

func (v NullableGetCurrencies200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrencies200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


