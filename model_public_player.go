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

// PublicPlayer struct for PublicPlayer
type PublicPlayer struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	GameId *string `json:"gameId,omitempty"`
	// This field has not had a description added.
	WalletId *string `json:"walletId,omitempty"`
	// This field has not had a description added.
	Username *string `json:"username,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
	Wallet *PublicPlayerWallet `json:"wallet,omitempty"`
}

// NewPublicPlayer instantiates a new PublicPlayer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicPlayer() *PublicPlayer {
	this := PublicPlayer{}
	return &this
}

// NewPublicPlayerWithDefaults instantiates a new PublicPlayer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicPlayerWithDefaults() *PublicPlayer {
	this := PublicPlayer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicPlayer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicPlayer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicPlayer) SetId(v string) {
	o.Id = &v
}

// GetGameId returns the GameId field value if set, zero value otherwise.
func (o *PublicPlayer) GetGameId() string {
	if o == nil || o.GameId == nil {
		var ret string
		return ret
	}
	return *o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayer) GetGameIdOk() (*string, bool) {
	if o == nil || o.GameId == nil {
		return nil, false
	}
	return o.GameId, true
}

// HasGameId returns a boolean if a field has been set.
func (o *PublicPlayer) HasGameId() bool {
	if o != nil && o.GameId != nil {
		return true
	}

	return false
}

// SetGameId gets a reference to the given string and assigns it to the GameId field.
func (o *PublicPlayer) SetGameId(v string) {
	o.GameId = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *PublicPlayer) GetWalletId() string {
	if o == nil || o.WalletId == nil {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayer) GetWalletIdOk() (*string, bool) {
	if o == nil || o.WalletId == nil {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *PublicPlayer) HasWalletId() bool {
	if o != nil && o.WalletId != nil {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *PublicPlayer) SetWalletId(v string) {
	o.WalletId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PublicPlayer) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayer) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PublicPlayer) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PublicPlayer) SetUsername(v string) {
	o.Username = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PublicPlayer) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayer) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PublicPlayer) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PublicPlayer) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PublicPlayer) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayer) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PublicPlayer) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PublicPlayer) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *PublicPlayer) GetWallet() PublicPlayerWallet {
	if o == nil || o.Wallet == nil {
		var ret PublicPlayerWallet
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPlayer) GetWalletOk() (*PublicPlayerWallet, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *PublicPlayer) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given PublicPlayerWallet and assigns it to the Wallet field.
func (o *PublicPlayer) SetWallet(v PublicPlayerWallet) {
	o.Wallet = &v
}

func (o PublicPlayer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.GameId != nil {
		toSerialize["gameId"] = o.GameId
	}
	if o.WalletId != nil {
		toSerialize["walletId"] = o.WalletId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	return json.Marshal(toSerialize)
}

type NullablePublicPlayer struct {
	value *PublicPlayer
	isSet bool
}

func (v NullablePublicPlayer) Get() *PublicPlayer {
	return v.value
}

func (v *NullablePublicPlayer) Set(val *PublicPlayer) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPlayer) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPlayer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPlayer(val *PublicPlayer) *NullablePublicPlayer {
	return &NullablePublicPlayer{value: val, isSet: true}
}

func (v NullablePublicPlayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPlayer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


