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

// GrantCurrencyRoleRequest struct for GrantCurrencyRoleRequest
type GrantCurrencyRoleRequest struct {
	// A valid MetaFab role or bytes string representing a role, such as `minter` or `0xc9eb32e43bf5ecbceacf00b32281dfc5d6d700a0db676ea26ccf938a385ac3b7`
	Role string `json:"role"`
	// A valid EVM based address to grant the role to.
	Address *string `json:"address,omitempty"`
	// A wallet id within the MetaFab ecosystem to grant the role to.
	WalletId *string `json:"walletId,omitempty"`
}

// NewGrantCurrencyRoleRequest instantiates a new GrantCurrencyRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantCurrencyRoleRequest(role string) *GrantCurrencyRoleRequest {
	this := GrantCurrencyRoleRequest{}
	this.Role = role
	return &this
}

// NewGrantCurrencyRoleRequestWithDefaults instantiates a new GrantCurrencyRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantCurrencyRoleRequestWithDefaults() *GrantCurrencyRoleRequest {
	this := GrantCurrencyRoleRequest{}
	return &this
}

// GetRole returns the Role field value
func (o *GrantCurrencyRoleRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *GrantCurrencyRoleRequest) GetRoleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *GrantCurrencyRoleRequest) SetRole(v string) {
	o.Role = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GrantCurrencyRoleRequest) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantCurrencyRoleRequest) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GrantCurrencyRoleRequest) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GrantCurrencyRoleRequest) SetAddress(v string) {
	o.Address = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *GrantCurrencyRoleRequest) GetWalletId() string {
	if o == nil || isNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantCurrencyRoleRequest) GetWalletIdOk() (*string, bool) {
	if o == nil || isNil(o.WalletId) {
    return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *GrantCurrencyRoleRequest) HasWalletId() bool {
	if o != nil && !isNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *GrantCurrencyRoleRequest) SetWalletId(v string) {
	o.WalletId = &v
}

func (o GrantCurrencyRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.WalletId) {
		toSerialize["walletId"] = o.WalletId
	}
	return json.Marshal(toSerialize)
}

type NullableGrantCurrencyRoleRequest struct {
	value *GrantCurrencyRoleRequest
	isSet bool
}

func (v NullableGrantCurrencyRoleRequest) Get() *GrantCurrencyRoleRequest {
	return v.value
}

func (v *NullableGrantCurrencyRoleRequest) Set(val *GrantCurrencyRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantCurrencyRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantCurrencyRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantCurrencyRoleRequest(val *GrantCurrencyRoleRequest) *NullableGrantCurrencyRoleRequest {
	return &NullableGrantCurrencyRoleRequest{value: val, isSet: true}
}

func (v NullableGrantCurrencyRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantCurrencyRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


