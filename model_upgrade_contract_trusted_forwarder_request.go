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

// UpgradeContractTrustedForwarderRequest struct for UpgradeContractTrustedForwarderRequest
type UpgradeContractTrustedForwarderRequest struct {
	// A ERC2771 forwarder smart contract address to assign as the new trusted forwarder of the target smart contract.
	ForwarderAddress string `json:"forwarderAddress"`
}

// NewUpgradeContractTrustedForwarderRequest instantiates a new UpgradeContractTrustedForwarderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeContractTrustedForwarderRequest(forwarderAddress string) *UpgradeContractTrustedForwarderRequest {
	this := UpgradeContractTrustedForwarderRequest{}
	this.ForwarderAddress = forwarderAddress
	return &this
}

// NewUpgradeContractTrustedForwarderRequestWithDefaults instantiates a new UpgradeContractTrustedForwarderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeContractTrustedForwarderRequestWithDefaults() *UpgradeContractTrustedForwarderRequest {
	this := UpgradeContractTrustedForwarderRequest{}
	return &this
}

// GetForwarderAddress returns the ForwarderAddress field value
func (o *UpgradeContractTrustedForwarderRequest) GetForwarderAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForwarderAddress
}

// GetForwarderAddressOk returns a tuple with the ForwarderAddress field value
// and a boolean to check if the value has been set.
func (o *UpgradeContractTrustedForwarderRequest) GetForwarderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForwarderAddress, true
}

// SetForwarderAddress sets field value
func (o *UpgradeContractTrustedForwarderRequest) SetForwarderAddress(v string) {
	o.ForwarderAddress = v
}

func (o UpgradeContractTrustedForwarderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["forwarderAddress"] = o.ForwarderAddress
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradeContractTrustedForwarderRequest struct {
	value *UpgradeContractTrustedForwarderRequest
	isSet bool
}

func (v NullableUpgradeContractTrustedForwarderRequest) Get() *UpgradeContractTrustedForwarderRequest {
	return v.value
}

func (v *NullableUpgradeContractTrustedForwarderRequest) Set(val *UpgradeContractTrustedForwarderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeContractTrustedForwarderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeContractTrustedForwarderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeContractTrustedForwarderRequest(val *UpgradeContractTrustedForwarderRequest) *NullableUpgradeContractTrustedForwarderRequest {
	return &NullableUpgradeContractTrustedForwarderRequest{value: val, isSet: true}
}

func (v NullableUpgradeContractTrustedForwarderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeContractTrustedForwarderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

