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

// RevokeCollectionRoleRequest struct for RevokeCollectionRoleRequest
type RevokeCollectionRoleRequest struct {
	// A valid MetaFab role or bytes string representing a role, such as `minter` or `0xc9eb32e43bf5ecbceacf00b32281dfc5d6d700a0db676ea26ccf938a385ac3b7`
	Role string `json:"role"`
	// A valid EVM based address to revoke the role from.
	Address *string `json:"address,omitempty"`
	// A wallet id within the MetaFab ecosystem to revoke the role from.
	WalletId []string `json:"walletId,omitempty"`
}

// NewRevokeCollectionRoleRequest instantiates a new RevokeCollectionRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeCollectionRoleRequest(role string) *RevokeCollectionRoleRequest {
	this := RevokeCollectionRoleRequest{}
	this.Role = role
	return &this
}

// NewRevokeCollectionRoleRequestWithDefaults instantiates a new RevokeCollectionRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeCollectionRoleRequestWithDefaults() *RevokeCollectionRoleRequest {
	this := RevokeCollectionRoleRequest{}
	return &this
}

// GetRole returns the Role field value
func (o *RevokeCollectionRoleRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RevokeCollectionRoleRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RevokeCollectionRoleRequest) SetRole(v string) {
	o.Role = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RevokeCollectionRoleRequest) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokeCollectionRoleRequest) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RevokeCollectionRoleRequest) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *RevokeCollectionRoleRequest) SetAddress(v string) {
	o.Address = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *RevokeCollectionRoleRequest) GetWalletId() []string {
	if o == nil || o.WalletId == nil {
		var ret []string
		return ret
	}
	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokeCollectionRoleRequest) GetWalletIdOk() ([]string, bool) {
	if o == nil || o.WalletId == nil {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *RevokeCollectionRoleRequest) HasWalletId() bool {
	if o != nil && o.WalletId != nil {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given []string and assigns it to the WalletId field.
func (o *RevokeCollectionRoleRequest) SetWalletId(v []string) {
	o.WalletId = v
}

func (o RevokeCollectionRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.WalletId != nil {
		toSerialize["walletId"] = o.WalletId
	}
	return json.Marshal(toSerialize)
}

type NullableRevokeCollectionRoleRequest struct {
	value *RevokeCollectionRoleRequest
	isSet bool
}

func (v NullableRevokeCollectionRoleRequest) Get() *RevokeCollectionRoleRequest {
	return v.value
}

func (v *NullableRevokeCollectionRoleRequest) Set(val *RevokeCollectionRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeCollectionRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeCollectionRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeCollectionRoleRequest(val *RevokeCollectionRoleRequest) *NullableRevokeCollectionRoleRequest {
	return &NullableRevokeCollectionRoleRequest{value: val, isSet: true}
}

func (v NullableRevokeCollectionRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeCollectionRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

