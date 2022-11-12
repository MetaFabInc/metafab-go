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

// CreateCollection200Response struct for CreateCollection200Response
type CreateCollection200Response struct {
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
	Contract *CreateCollection200ResponseAllOfContract `json:"contract,omitempty"`
}

// NewCreateCollection200Response instantiates a new CreateCollection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollection200Response() *CreateCollection200Response {
	this := CreateCollection200Response{}
	return &this
}

// NewCreateCollection200ResponseWithDefaults instantiates a new CreateCollection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollection200ResponseWithDefaults() *CreateCollection200Response {
	this := CreateCollection200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateCollection200Response) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200Response) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateCollection200Response) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateCollection200Response) SetId(v string) {
	o.Id = &v
}

// GetGameId returns the GameId field value if set, zero value otherwise.
func (o *CreateCollection200Response) GetGameId() string {
	if o == nil || o.GameId == nil {
		var ret string
		return ret
	}
	return *o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200Response) GetGameIdOk() (*string, bool) {
	if o == nil || o.GameId == nil {
		return nil, false
	}
	return o.GameId, true
}

// HasGameId returns a boolean if a field has been set.
func (o *CreateCollection200Response) HasGameId() bool {
	if o != nil && o.GameId != nil {
		return true
	}

	return false
}

// SetGameId gets a reference to the given string and assigns it to the GameId field.
func (o *CreateCollection200Response) SetGameId(v string) {
	o.GameId = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *CreateCollection200Response) GetContractId() string {
	if o == nil || o.ContractId == nil {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200Response) GetContractIdOk() (*string, bool) {
	if o == nil || o.ContractId == nil {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *CreateCollection200Response) HasContractId() bool {
	if o != nil && o.ContractId != nil {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *CreateCollection200Response) SetContractId(v string) {
	o.ContractId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateCollection200Response) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200Response) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateCollection200Response) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CreateCollection200Response) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateCollection200Response) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200Response) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateCollection200Response) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CreateCollection200Response) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *CreateCollection200Response) GetContract() CreateCollection200ResponseAllOfContract {
	if o == nil || o.Contract == nil {
		var ret CreateCollection200ResponseAllOfContract
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200Response) GetContractOk() (*CreateCollection200ResponseAllOfContract, bool) {
	if o == nil || o.Contract == nil {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *CreateCollection200Response) HasContract() bool {
	if o != nil && o.Contract != nil {
		return true
	}

	return false
}

// SetContract gets a reference to the given CreateCollection200ResponseAllOfContract and assigns it to the Contract field.
func (o *CreateCollection200Response) SetContract(v CreateCollection200ResponseAllOfContract) {
	o.Contract = &v
}

func (o CreateCollection200Response) MarshalJSON() ([]byte, error) {
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
	if o.Contract != nil {
		toSerialize["contract"] = o.Contract
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCollection200Response struct {
	value *CreateCollection200Response
	isSet bool
}

func (v NullableCreateCollection200Response) Get() *CreateCollection200Response {
	return v.value
}

func (v *NullableCreateCollection200Response) Set(val *CreateCollection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollection200Response(val *CreateCollection200Response) *NullableCreateCollection200Response {
	return &NullableCreateCollection200Response{value: val, isSet: true}
}

func (v NullableCreateCollection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


