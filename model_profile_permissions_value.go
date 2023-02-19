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

// ProfilePermissionsValue Key should be the contract address, value is the permissions object request for the contract.
type ProfilePermissionsValue struct {
	// The target chain for the contract and related permissions.
	Chain *string `json:"chain,omitempty"`
	// An optional array of valid permissioning scopes.
	Scopes []string `json:"scopes,omitempty"`
	// An optional array of contract functions to request permission for.
	Functions []string `json:"functions,omitempty"`
	// A maximum lifetime limit of erc20 that can be tranferred for this contract address.
	Erc20Limit *int32 `json:"erc20Limit,omitempty"`
	// An object mapping erc1155 ids to maximum lifetime transfer limits of each permitted item id supplied for this contract address.
	Erc1155Limits *map[string]int32 `json:"erc1155Limits,omitempty"`
}

// NewProfilePermissionsValue instantiates a new ProfilePermissionsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfilePermissionsValue() *ProfilePermissionsValue {
	this := ProfilePermissionsValue{}
	return &this
}

// NewProfilePermissionsValueWithDefaults instantiates a new ProfilePermissionsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfilePermissionsValueWithDefaults() *ProfilePermissionsValue {
	this := ProfilePermissionsValue{}
	return &this
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *ProfilePermissionsValue) GetChain() string {
	if o == nil || isNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePermissionsValue) GetChainOk() (*string, bool) {
	if o == nil || isNil(o.Chain) {
    return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *ProfilePermissionsValue) HasChain() bool {
	if o != nil && !isNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *ProfilePermissionsValue) SetChain(v string) {
	o.Chain = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ProfilePermissionsValue) GetScopes() []string {
	if o == nil || isNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePermissionsValue) GetScopesOk() ([]string, bool) {
	if o == nil || isNil(o.Scopes) {
    return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ProfilePermissionsValue) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ProfilePermissionsValue) SetScopes(v []string) {
	o.Scopes = v
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
func (o *ProfilePermissionsValue) GetFunctions() []string {
	if o == nil || isNil(o.Functions) {
		var ret []string
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePermissionsValue) GetFunctionsOk() ([]string, bool) {
	if o == nil || isNil(o.Functions) {
    return nil, false
	}
	return o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *ProfilePermissionsValue) HasFunctions() bool {
	if o != nil && !isNil(o.Functions) {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given []string and assigns it to the Functions field.
func (o *ProfilePermissionsValue) SetFunctions(v []string) {
	o.Functions = v
}

// GetErc20Limit returns the Erc20Limit field value if set, zero value otherwise.
func (o *ProfilePermissionsValue) GetErc20Limit() int32 {
	if o == nil || isNil(o.Erc20Limit) {
		var ret int32
		return ret
	}
	return *o.Erc20Limit
}

// GetErc20LimitOk returns a tuple with the Erc20Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePermissionsValue) GetErc20LimitOk() (*int32, bool) {
	if o == nil || isNil(o.Erc20Limit) {
    return nil, false
	}
	return o.Erc20Limit, true
}

// HasErc20Limit returns a boolean if a field has been set.
func (o *ProfilePermissionsValue) HasErc20Limit() bool {
	if o != nil && !isNil(o.Erc20Limit) {
		return true
	}

	return false
}

// SetErc20Limit gets a reference to the given int32 and assigns it to the Erc20Limit field.
func (o *ProfilePermissionsValue) SetErc20Limit(v int32) {
	o.Erc20Limit = &v
}

// GetErc1155Limits returns the Erc1155Limits field value if set, zero value otherwise.
func (o *ProfilePermissionsValue) GetErc1155Limits() map[string]int32 {
	if o == nil || isNil(o.Erc1155Limits) {
		var ret map[string]int32
		return ret
	}
	return *o.Erc1155Limits
}

// GetErc1155LimitsOk returns a tuple with the Erc1155Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePermissionsValue) GetErc1155LimitsOk() (*map[string]int32, bool) {
	if o == nil || isNil(o.Erc1155Limits) {
    return nil, false
	}
	return o.Erc1155Limits, true
}

// HasErc1155Limits returns a boolean if a field has been set.
func (o *ProfilePermissionsValue) HasErc1155Limits() bool {
	if o != nil && !isNil(o.Erc1155Limits) {
		return true
	}

	return false
}

// SetErc1155Limits gets a reference to the given map[string]int32 and assigns it to the Erc1155Limits field.
func (o *ProfilePermissionsValue) SetErc1155Limits(v map[string]int32) {
	o.Erc1155Limits = &v
}

func (o ProfilePermissionsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Chain) {
		toSerialize["chain"] = o.Chain
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.Functions) {
		toSerialize["functions"] = o.Functions
	}
	if !isNil(o.Erc20Limit) {
		toSerialize["erc20Limit"] = o.Erc20Limit
	}
	if !isNil(o.Erc1155Limits) {
		toSerialize["erc1155Limits"] = o.Erc1155Limits
	}
	return json.Marshal(toSerialize)
}

type NullableProfilePermissionsValue struct {
	value *ProfilePermissionsValue
	isSet bool
}

func (v NullableProfilePermissionsValue) Get() *ProfilePermissionsValue {
	return v.value
}

func (v *NullableProfilePermissionsValue) Set(val *ProfilePermissionsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableProfilePermissionsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableProfilePermissionsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfilePermissionsValue(val *ProfilePermissionsValue) *NullableProfilePermissionsValue {
	return &NullableProfilePermissionsValue{value: val, isSet: true}
}

func (v NullableProfilePermissionsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfilePermissionsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


