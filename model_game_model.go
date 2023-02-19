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

// GameModel struct for GameModel
type GameModel struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	WalletId *string `json:"walletId,omitempty"`
	// This field has not had a description added.
	FundingWalletId *string `json:"fundingWalletId,omitempty"`
	// This field has not had a description added.
	Email *string `json:"email,omitempty"`
	// This field has not had a description added.
	Name *string `json:"name,omitempty"`
	// This field has not had a description added.
	Rpcs map[string]interface{} `json:"rpcs,omitempty"`
	// This field has not had a description added.
	RedirectUris map[string]interface{} `json:"redirectUris,omitempty"`
	// This field has not had a description added.
	IconImageUrl *string `json:"iconImageUrl,omitempty"`
	// This field has not had a description added.
	CoverImageUrl *string `json:"coverImageUrl,omitempty"`
	// This field has not had a description added.
	PrimaryColorHex *string `json:"primaryColorHex,omitempty"`
	// This field has not had a description added.
	PublishedKey *string `json:"publishedKey,omitempty"`
	// This field has not had a description added.
	SecretKey *string `json:"secretKey,omitempty"`
	// This field has not had a description added.
	Verified *bool `json:"verified,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewGameModel instantiates a new GameModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameModel() *GameModel {
	this := GameModel{}
	return &this
}

// NewGameModelWithDefaults instantiates a new GameModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameModelWithDefaults() *GameModel {
	this := GameModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GameModel) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GameModel) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GameModel) SetId(v string) {
	o.Id = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *GameModel) GetWalletId() string {
	if o == nil || isNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetWalletIdOk() (*string, bool) {
	if o == nil || isNil(o.WalletId) {
    return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *GameModel) HasWalletId() bool {
	if o != nil && !isNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *GameModel) SetWalletId(v string) {
	o.WalletId = &v
}

// GetFundingWalletId returns the FundingWalletId field value if set, zero value otherwise.
func (o *GameModel) GetFundingWalletId() string {
	if o == nil || isNil(o.FundingWalletId) {
		var ret string
		return ret
	}
	return *o.FundingWalletId
}

// GetFundingWalletIdOk returns a tuple with the FundingWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetFundingWalletIdOk() (*string, bool) {
	if o == nil || isNil(o.FundingWalletId) {
    return nil, false
	}
	return o.FundingWalletId, true
}

// HasFundingWalletId returns a boolean if a field has been set.
func (o *GameModel) HasFundingWalletId() bool {
	if o != nil && !isNil(o.FundingWalletId) {
		return true
	}

	return false
}

// SetFundingWalletId gets a reference to the given string and assigns it to the FundingWalletId field.
func (o *GameModel) SetFundingWalletId(v string) {
	o.FundingWalletId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GameModel) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GameModel) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GameModel) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GameModel) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GameModel) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GameModel) SetName(v string) {
	o.Name = &v
}

// GetRpcs returns the Rpcs field value if set, zero value otherwise.
func (o *GameModel) GetRpcs() map[string]interface{} {
	if o == nil || isNil(o.Rpcs) {
		var ret map[string]interface{}
		return ret
	}
	return o.Rpcs
}

// GetRpcsOk returns a tuple with the Rpcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetRpcsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Rpcs) {
    return map[string]interface{}{}, false
	}
	return o.Rpcs, true
}

// HasRpcs returns a boolean if a field has been set.
func (o *GameModel) HasRpcs() bool {
	if o != nil && !isNil(o.Rpcs) {
		return true
	}

	return false
}

// SetRpcs gets a reference to the given map[string]interface{} and assigns it to the Rpcs field.
func (o *GameModel) SetRpcs(v map[string]interface{}) {
	o.Rpcs = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *GameModel) GetRedirectUris() map[string]interface{} {
	if o == nil || isNil(o.RedirectUris) {
		var ret map[string]interface{}
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetRedirectUrisOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.RedirectUris) {
    return map[string]interface{}{}, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *GameModel) HasRedirectUris() bool {
	if o != nil && !isNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given map[string]interface{} and assigns it to the RedirectUris field.
func (o *GameModel) SetRedirectUris(v map[string]interface{}) {
	o.RedirectUris = v
}

// GetIconImageUrl returns the IconImageUrl field value if set, zero value otherwise.
func (o *GameModel) GetIconImageUrl() string {
	if o == nil || isNil(o.IconImageUrl) {
		var ret string
		return ret
	}
	return *o.IconImageUrl
}

// GetIconImageUrlOk returns a tuple with the IconImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetIconImageUrlOk() (*string, bool) {
	if o == nil || isNil(o.IconImageUrl) {
    return nil, false
	}
	return o.IconImageUrl, true
}

// HasIconImageUrl returns a boolean if a field has been set.
func (o *GameModel) HasIconImageUrl() bool {
	if o != nil && !isNil(o.IconImageUrl) {
		return true
	}

	return false
}

// SetIconImageUrl gets a reference to the given string and assigns it to the IconImageUrl field.
func (o *GameModel) SetIconImageUrl(v string) {
	o.IconImageUrl = &v
}

// GetCoverImageUrl returns the CoverImageUrl field value if set, zero value otherwise.
func (o *GameModel) GetCoverImageUrl() string {
	if o == nil || isNil(o.CoverImageUrl) {
		var ret string
		return ret
	}
	return *o.CoverImageUrl
}

// GetCoverImageUrlOk returns a tuple with the CoverImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetCoverImageUrlOk() (*string, bool) {
	if o == nil || isNil(o.CoverImageUrl) {
    return nil, false
	}
	return o.CoverImageUrl, true
}

// HasCoverImageUrl returns a boolean if a field has been set.
func (o *GameModel) HasCoverImageUrl() bool {
	if o != nil && !isNil(o.CoverImageUrl) {
		return true
	}

	return false
}

// SetCoverImageUrl gets a reference to the given string and assigns it to the CoverImageUrl field.
func (o *GameModel) SetCoverImageUrl(v string) {
	o.CoverImageUrl = &v
}

// GetPrimaryColorHex returns the PrimaryColorHex field value if set, zero value otherwise.
func (o *GameModel) GetPrimaryColorHex() string {
	if o == nil || isNil(o.PrimaryColorHex) {
		var ret string
		return ret
	}
	return *o.PrimaryColorHex
}

// GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetPrimaryColorHexOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryColorHex) {
    return nil, false
	}
	return o.PrimaryColorHex, true
}

// HasPrimaryColorHex returns a boolean if a field has been set.
func (o *GameModel) HasPrimaryColorHex() bool {
	if o != nil && !isNil(o.PrimaryColorHex) {
		return true
	}

	return false
}

// SetPrimaryColorHex gets a reference to the given string and assigns it to the PrimaryColorHex field.
func (o *GameModel) SetPrimaryColorHex(v string) {
	o.PrimaryColorHex = &v
}

// GetPublishedKey returns the PublishedKey field value if set, zero value otherwise.
func (o *GameModel) GetPublishedKey() string {
	if o == nil || isNil(o.PublishedKey) {
		var ret string
		return ret
	}
	return *o.PublishedKey
}

// GetPublishedKeyOk returns a tuple with the PublishedKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetPublishedKeyOk() (*string, bool) {
	if o == nil || isNil(o.PublishedKey) {
    return nil, false
	}
	return o.PublishedKey, true
}

// HasPublishedKey returns a boolean if a field has been set.
func (o *GameModel) HasPublishedKey() bool {
	if o != nil && !isNil(o.PublishedKey) {
		return true
	}

	return false
}

// SetPublishedKey gets a reference to the given string and assigns it to the PublishedKey field.
func (o *GameModel) SetPublishedKey(v string) {
	o.PublishedKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *GameModel) GetSecretKey() string {
	if o == nil || isNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetSecretKeyOk() (*string, bool) {
	if o == nil || isNil(o.SecretKey) {
    return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *GameModel) HasSecretKey() bool {
	if o != nil && !isNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *GameModel) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *GameModel) GetVerified() bool {
	if o == nil || isNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetVerifiedOk() (*bool, bool) {
	if o == nil || isNil(o.Verified) {
    return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *GameModel) HasVerified() bool {
	if o != nil && !isNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *GameModel) SetVerified(v bool) {
	o.Verified = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GameModel) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GameModel) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GameModel) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GameModel) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GameModel) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GameModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o GameModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.WalletId) {
		toSerialize["walletId"] = o.WalletId
	}
	if !isNil(o.FundingWalletId) {
		toSerialize["fundingWalletId"] = o.FundingWalletId
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Rpcs) {
		toSerialize["rpcs"] = o.Rpcs
	}
	if !isNil(o.RedirectUris) {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if !isNil(o.IconImageUrl) {
		toSerialize["iconImageUrl"] = o.IconImageUrl
	}
	if !isNil(o.CoverImageUrl) {
		toSerialize["coverImageUrl"] = o.CoverImageUrl
	}
	if !isNil(o.PrimaryColorHex) {
		toSerialize["primaryColorHex"] = o.PrimaryColorHex
	}
	if !isNil(o.PublishedKey) {
		toSerialize["publishedKey"] = o.PublishedKey
	}
	if !isNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	if !isNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableGameModel struct {
	value *GameModel
	isSet bool
}

func (v NullableGameModel) Get() *GameModel {
	return v.value
}

func (v *NullableGameModel) Set(val *GameModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGameModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGameModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameModel(val *GameModel) *NullableGameModel {
	return &NullableGameModel{value: val, isSet: true}
}

func (v NullableGameModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


