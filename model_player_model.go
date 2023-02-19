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

// PlayerModel struct for PlayerModel
type PlayerModel struct {
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
}

// NewPlayerModel instantiates a new PlayerModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerModel() *PlayerModel {
	this := PlayerModel{}
	return &this
}

// NewPlayerModelWithDefaults instantiates a new PlayerModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerModelWithDefaults() *PlayerModel {
	this := PlayerModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlayerModel) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlayerModel) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlayerModel) SetId(v string) {
	o.Id = &v
}

// GetGameId returns the GameId field value if set, zero value otherwise.
func (o *PlayerModel) GetGameId() string {
	if o == nil || isNil(o.GameId) {
		var ret string
		return ret
	}
	return *o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetGameIdOk() (*string, bool) {
	if o == nil || isNil(o.GameId) {
    return nil, false
	}
	return o.GameId, true
}

// HasGameId returns a boolean if a field has been set.
func (o *PlayerModel) HasGameId() bool {
	if o != nil && !isNil(o.GameId) {
		return true
	}

	return false
}

// SetGameId gets a reference to the given string and assigns it to the GameId field.
func (o *PlayerModel) SetGameId(v string) {
	o.GameId = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *PlayerModel) GetWalletId() string {
	if o == nil || isNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetWalletIdOk() (*string, bool) {
	if o == nil || isNil(o.WalletId) {
    return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *PlayerModel) HasWalletId() bool {
	if o != nil && !isNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *PlayerModel) SetWalletId(v string) {
	o.WalletId = &v
}

// GetConnectedWalletId returns the ConnectedWalletId field value if set, zero value otherwise.
func (o *PlayerModel) GetConnectedWalletId() string {
	if o == nil || isNil(o.ConnectedWalletId) {
		var ret string
		return ret
	}
	return *o.ConnectedWalletId
}

// GetConnectedWalletIdOk returns a tuple with the ConnectedWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetConnectedWalletIdOk() (*string, bool) {
	if o == nil || isNil(o.ConnectedWalletId) {
    return nil, false
	}
	return o.ConnectedWalletId, true
}

// HasConnectedWalletId returns a boolean if a field has been set.
func (o *PlayerModel) HasConnectedWalletId() bool {
	if o != nil && !isNil(o.ConnectedWalletId) {
		return true
	}

	return false
}

// SetConnectedWalletId gets a reference to the given string and assigns it to the ConnectedWalletId field.
func (o *PlayerModel) SetConnectedWalletId(v string) {
	o.ConnectedWalletId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *PlayerModel) GetProfileId() string {
	if o == nil || isNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *PlayerModel) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *PlayerModel) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PlayerModel) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PlayerModel) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PlayerModel) SetUsername(v string) {
	o.Username = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *PlayerModel) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
    return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *PlayerModel) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *PlayerModel) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetProfilePermissions returns the ProfilePermissions field value if set, zero value otherwise.
func (o *PlayerModel) GetProfilePermissions() map[string]interface{} {
	if o == nil || isNil(o.ProfilePermissions) {
		var ret map[string]interface{}
		return ret
	}
	return o.ProfilePermissions
}

// GetProfilePermissionsOk returns a tuple with the ProfilePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetProfilePermissionsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ProfilePermissions) {
    return map[string]interface{}{}, false
	}
	return o.ProfilePermissions, true
}

// HasProfilePermissions returns a boolean if a field has been set.
func (o *PlayerModel) HasProfilePermissions() bool {
	if o != nil && !isNil(o.ProfilePermissions) {
		return true
	}

	return false
}

// SetProfilePermissions gets a reference to the given map[string]interface{} and assigns it to the ProfilePermissions field.
func (o *PlayerModel) SetProfilePermissions(v map[string]interface{}) {
	o.ProfilePermissions = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PlayerModel) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PlayerModel) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PlayerModel) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PlayerModel) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PlayerModel) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PlayerModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o PlayerModel) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullablePlayerModel struct {
	value *PlayerModel
	isSet bool
}

func (v NullablePlayerModel) Get() *PlayerModel {
	return v.value
}

func (v *NullablePlayerModel) Set(val *PlayerModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerModel(val *PlayerModel) *NullablePlayerModel {
	return &NullablePlayerModel{value: val, isSet: true}
}

func (v NullablePlayerModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


