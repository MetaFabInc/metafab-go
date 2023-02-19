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

// UpdateProfilePlayer200Response struct for UpdateProfilePlayer200Response
type UpdateProfilePlayer200Response struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	GameId *string `json:"gameId,omitempty"`
	// This field has not had a description added.
	WalletId *string `json:"walletId,omitempty"`
	// This field has not had a description added.
	ConnectedWalletId *string `json:"connectedWalletId,omitempty"`
	// This field has not had a description added.
	ProfileId *string `json:"profileId,omitempty"`
	// This field has not had a description added.
	Username *string `json:"username,omitempty"`
	// This field has not had a description added.
	AccessToken *string `json:"accessToken,omitempty"`
	// This field has not had a description added.
	ProfilePermissions map[string]interface{} `json:"profilePermissions,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
	Wallet *WalletModel `json:"wallet,omitempty"`
	CustodialWallet *WalletModel `json:"custodialWallet,omitempty"`
}

// NewUpdateProfilePlayer200Response instantiates a new UpdateProfilePlayer200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfilePlayer200Response() *UpdateProfilePlayer200Response {
	this := UpdateProfilePlayer200Response{}
	return &this
}

// NewUpdateProfilePlayer200ResponseWithDefaults instantiates a new UpdateProfilePlayer200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfilePlayer200ResponseWithDefaults() *UpdateProfilePlayer200Response {
	this := UpdateProfilePlayer200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateProfilePlayer200Response) SetId(v string) {
	o.Id = &v
}

// GetGameId returns the GameId field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetGameId() string {
	if o == nil || isNil(o.GameId) {
		var ret string
		return ret
	}
	return *o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetGameIdOk() (*string, bool) {
	if o == nil || isNil(o.GameId) {
    return nil, false
	}
	return o.GameId, true
}

// HasGameId returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasGameId() bool {
	if o != nil && !isNil(o.GameId) {
		return true
	}

	return false
}

// SetGameId gets a reference to the given string and assigns it to the GameId field.
func (o *UpdateProfilePlayer200Response) SetGameId(v string) {
	o.GameId = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetWalletId() string {
	if o == nil || isNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetWalletIdOk() (*string, bool) {
	if o == nil || isNil(o.WalletId) {
    return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasWalletId() bool {
	if o != nil && !isNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *UpdateProfilePlayer200Response) SetWalletId(v string) {
	o.WalletId = &v
}

// GetConnectedWalletId returns the ConnectedWalletId field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetConnectedWalletId() string {
	if o == nil || isNil(o.ConnectedWalletId) {
		var ret string
		return ret
	}
	return *o.ConnectedWalletId
}

// GetConnectedWalletIdOk returns a tuple with the ConnectedWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetConnectedWalletIdOk() (*string, bool) {
	if o == nil || isNil(o.ConnectedWalletId) {
    return nil, false
	}
	return o.ConnectedWalletId, true
}

// HasConnectedWalletId returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasConnectedWalletId() bool {
	if o != nil && !isNil(o.ConnectedWalletId) {
		return true
	}

	return false
}

// SetConnectedWalletId gets a reference to the given string and assigns it to the ConnectedWalletId field.
func (o *UpdateProfilePlayer200Response) SetConnectedWalletId(v string) {
	o.ConnectedWalletId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetProfileId() string {
	if o == nil || isNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *UpdateProfilePlayer200Response) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateProfilePlayer200Response) SetUsername(v string) {
	o.Username = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
    return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *UpdateProfilePlayer200Response) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetProfilePermissions returns the ProfilePermissions field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetProfilePermissions() map[string]interface{} {
	if o == nil || isNil(o.ProfilePermissions) {
		var ret map[string]interface{}
		return ret
	}
	return o.ProfilePermissions
}

// GetProfilePermissionsOk returns a tuple with the ProfilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetProfilePermissionsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ProfilePermissions) {
    return map[string]interface{}{}, false
	}
	return o.ProfilePermissions, true
}

// HasProfilePermissions returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasProfilePermissions() bool {
	if o != nil && !isNil(o.ProfilePermissions) {
		return true
	}

	return false
}

// SetProfilePermissions gets a reference to the given map[string]interface{} and assigns it to the ProfilePermissions field.
func (o *UpdateProfilePlayer200Response) SetProfilePermissions(v map[string]interface{}) {
	o.ProfilePermissions = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UpdateProfilePlayer200Response) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UpdateProfilePlayer200Response) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetWallet() WalletModel {
	if o == nil || isNil(o.Wallet) {
		var ret WalletModel
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetWalletOk() (*WalletModel, bool) {
	if o == nil || isNil(o.Wallet) {
    return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasWallet() bool {
	if o != nil && !isNil(o.Wallet) {
		return true
	}

	return false
}

// SetWallet gets a reference to the given WalletModel and assigns it to the Wallet field.
func (o *UpdateProfilePlayer200Response) SetWallet(v WalletModel) {
	o.Wallet = &v
}

// GetCustodialWallet returns the CustodialWallet field value if set, zero value otherwise.
func (o *UpdateProfilePlayer200Response) GetCustodialWallet() WalletModel {
	if o == nil || isNil(o.CustodialWallet) {
		var ret WalletModel
		return ret
	}
	return *o.CustodialWallet
}

// GetCustodialWalletOk returns a tuple with the CustodialWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfilePlayer200Response) GetCustodialWalletOk() (*WalletModel, bool) {
	if o == nil || isNil(o.CustodialWallet) {
    return nil, false
	}
	return o.CustodialWallet, true
}

// HasCustodialWallet returns a boolean if a field has been set.
func (o *UpdateProfilePlayer200Response) HasCustodialWallet() bool {
	if o != nil && !isNil(o.CustodialWallet) {
		return true
	}

	return false
}

// SetCustodialWallet gets a reference to the given WalletModel and assigns it to the CustodialWallet field.
func (o *UpdateProfilePlayer200Response) SetCustodialWallet(v WalletModel) {
	o.CustodialWallet = &v
}

func (o UpdateProfilePlayer200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.GameId) {
		toSerialize["gameId"] = o.GameId
	}
	if !isNil(o.WalletId) {
		toSerialize["walletId"] = o.WalletId
	}
	if !isNil(o.ConnectedWalletId) {
		toSerialize["connectedWalletId"] = o.ConnectedWalletId
	}
	if !isNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !isNil(o.ProfilePermissions) {
		toSerialize["profilePermissions"] = o.ProfilePermissions
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Wallet) {
		toSerialize["wallet"] = o.Wallet
	}
	if !isNil(o.CustodialWallet) {
		toSerialize["custodialWallet"] = o.CustodialWallet
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProfilePlayer200Response struct {
	value *UpdateProfilePlayer200Response
	isSet bool
}

func (v NullableUpdateProfilePlayer200Response) Get() *UpdateProfilePlayer200Response {
	return v.value
}

func (v *NullableUpdateProfilePlayer200Response) Set(val *UpdateProfilePlayer200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfilePlayer200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfilePlayer200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfilePlayer200Response(val *UpdateProfilePlayer200Response) *NullableUpdateProfilePlayer200Response {
	return &NullableUpdateProfilePlayer200Response{value: val, isSet: true}
}

func (v NullableUpdateProfilePlayer200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfilePlayer200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

