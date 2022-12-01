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

// AuthGame200Response struct for AuthGame200Response
type AuthGame200Response struct {
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
	PublishedKey *string `json:"publishedKey,omitempty"`
	// This field has not had a description added.
	SecretKey *string `json:"secretKey,omitempty"`
	// This field has not had a description added.
	Verified *bool `json:"verified,omitempty"`
	// This field has not had a description added.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// This field has not had a description added.
	CreatedAt *string `json:"createdAt,omitempty"`
	Wallet *WalletModel `json:"wallet,omitempty"`
	FundingWallet *WalletModel `json:"fundingWallet,omitempty"`
}

// NewAuthGame200Response instantiates a new AuthGame200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthGame200Response() *AuthGame200Response {
	this := AuthGame200Response{}
	return &this
}

// NewAuthGame200ResponseWithDefaults instantiates a new AuthGame200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthGame200ResponseWithDefaults() *AuthGame200Response {
	this := AuthGame200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthGame200Response) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthGame200Response) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthGame200Response) SetId(v string) {
	o.Id = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *AuthGame200Response) GetWalletId() string {
	if o == nil || o.WalletId == nil {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetWalletIdOk() (*string, bool) {
	if o == nil || o.WalletId == nil {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *AuthGame200Response) HasWalletId() bool {
	if o != nil && o.WalletId != nil {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *AuthGame200Response) SetWalletId(v string) {
	o.WalletId = &v
}

// GetFundingWalletId returns the FundingWalletId field value if set, zero value otherwise.
func (o *AuthGame200Response) GetFundingWalletId() string {
	if o == nil || o.FundingWalletId == nil {
		var ret string
		return ret
	}
	return *o.FundingWalletId
}

// GetFundingWalletIdOk returns a tuple with the FundingWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetFundingWalletIdOk() (*string, bool) {
	if o == nil || o.FundingWalletId == nil {
		return nil, false
	}
	return o.FundingWalletId, true
}

// HasFundingWalletId returns a boolean if a field has been set.
func (o *AuthGame200Response) HasFundingWalletId() bool {
	if o != nil && o.FundingWalletId != nil {
		return true
	}

	return false
}

// SetFundingWalletId gets a reference to the given string and assigns it to the FundingWalletId field.
func (o *AuthGame200Response) SetFundingWalletId(v string) {
	o.FundingWalletId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthGame200Response) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthGame200Response) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthGame200Response) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthGame200Response) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthGame200Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthGame200Response) SetName(v string) {
	o.Name = &v
}

// GetRpcs returns the Rpcs field value if set, zero value otherwise.
func (o *AuthGame200Response) GetRpcs() map[string]interface{} {
	if o == nil || o.Rpcs == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Rpcs
}

// GetRpcsOk returns a tuple with the Rpcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetRpcsOk() (map[string]interface{}, bool) {
	if o == nil || o.Rpcs == nil {
		return nil, false
	}
	return o.Rpcs, true
}

// HasRpcs returns a boolean if a field has been set.
func (o *AuthGame200Response) HasRpcs() bool {
	if o != nil && o.Rpcs != nil {
		return true
	}

	return false
}

// SetRpcs gets a reference to the given map[string]interface{} and assigns it to the Rpcs field.
func (o *AuthGame200Response) SetRpcs(v map[string]interface{}) {
	o.Rpcs = v
}

// GetPublishedKey returns the PublishedKey field value if set, zero value otherwise.
func (o *AuthGame200Response) GetPublishedKey() string {
	if o == nil || o.PublishedKey == nil {
		var ret string
		return ret
	}
	return *o.PublishedKey
}

// GetPublishedKeyOk returns a tuple with the PublishedKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetPublishedKeyOk() (*string, bool) {
	if o == nil || o.PublishedKey == nil {
		return nil, false
	}
	return o.PublishedKey, true
}

// HasPublishedKey returns a boolean if a field has been set.
func (o *AuthGame200Response) HasPublishedKey() bool {
	if o != nil && o.PublishedKey != nil {
		return true
	}

	return false
}

// SetPublishedKey gets a reference to the given string and assigns it to the PublishedKey field.
func (o *AuthGame200Response) SetPublishedKey(v string) {
	o.PublishedKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *AuthGame200Response) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *AuthGame200Response) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *AuthGame200Response) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *AuthGame200Response) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *AuthGame200Response) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *AuthGame200Response) SetVerified(v bool) {
	o.Verified = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuthGame200Response) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuthGame200Response) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuthGame200Response) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuthGame200Response) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuthGame200Response) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AuthGame200Response) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *AuthGame200Response) GetWallet() WalletModel {
	if o == nil || o.Wallet == nil {
		var ret WalletModel
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetWalletOk() (*WalletModel, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *AuthGame200Response) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given WalletModel and assigns it to the Wallet field.
func (o *AuthGame200Response) SetWallet(v WalletModel) {
	o.Wallet = &v
}

// GetFundingWallet returns the FundingWallet field value if set, zero value otherwise.
func (o *AuthGame200Response) GetFundingWallet() WalletModel {
	if o == nil || o.FundingWallet == nil {
		var ret WalletModel
		return ret
	}
	return *o.FundingWallet
}

// GetFundingWalletOk returns a tuple with the FundingWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthGame200Response) GetFundingWalletOk() (*WalletModel, bool) {
	if o == nil || o.FundingWallet == nil {
		return nil, false
	}
	return o.FundingWallet, true
}

// HasFundingWallet returns a boolean if a field has been set.
func (o *AuthGame200Response) HasFundingWallet() bool {
	if o != nil && o.FundingWallet != nil {
		return true
	}

	return false
}

// SetFundingWallet gets a reference to the given WalletModel and assigns it to the FundingWallet field.
func (o *AuthGame200Response) SetFundingWallet(v WalletModel) {
	o.FundingWallet = &v
}

func (o AuthGame200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.WalletId != nil {
		toSerialize["walletId"] = o.WalletId
	}
	if o.FundingWalletId != nil {
		toSerialize["fundingWalletId"] = o.FundingWalletId
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rpcs != nil {
		toSerialize["rpcs"] = o.Rpcs
	}
	if o.PublishedKey != nil {
		toSerialize["publishedKey"] = o.PublishedKey
	}
	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
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
	if o.FundingWallet != nil {
		toSerialize["fundingWallet"] = o.FundingWallet
	}
	return json.Marshal(toSerialize)
}

type NullableAuthGame200Response struct {
	value *AuthGame200Response
	isSet bool
}

func (v NullableAuthGame200Response) Get() *AuthGame200Response {
	return v.value
}

func (v *NullableAuthGame200Response) Set(val *AuthGame200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthGame200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthGame200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthGame200Response(val *AuthGame200Response) *NullableAuthGame200Response {
	return &NullableAuthGame200Response{value: val, isSet: true}
}

func (v NullableAuthGame200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthGame200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


