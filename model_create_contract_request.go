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

// CreateContractRequest struct for CreateContractRequest
type CreateContractRequest struct {
	// The address of the existing contract.
	Address string `json:"address"`
	// The address of the ERC2771 forwarding contract trusted by the contract.
	ForwarderAddress *string `json:"forwarderAddress,omitempty"`
	// JSON of the abi.
	Abi string `json:"abi"`
	// The blockchain you want to deploy this currency on. Support for new blockchains are added over time.
	Chain string `json:"chain"`
}

// NewCreateContractRequest instantiates a new CreateContractRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContractRequest(address string, abi string, chain string) *CreateContractRequest {
	this := CreateContractRequest{}
	this.Address = address
	this.Abi = abi
	this.Chain = chain
	return &this
}

// NewCreateContractRequestWithDefaults instantiates a new CreateContractRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContractRequestWithDefaults() *CreateContractRequest {
	this := CreateContractRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateContractRequest) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateContractRequest) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateContractRequest) SetAddress(v string) {
	o.Address = v
}

// GetForwarderAddress returns the ForwarderAddress field value if set, zero value otherwise.
func (o *CreateContractRequest) GetForwarderAddress() string {
	if o == nil || o.ForwarderAddress == nil {
		var ret string
		return ret
	}
	return *o.ForwarderAddress
}

// GetForwarderAddressOk returns a tuple with the ForwarderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContractRequest) GetForwarderAddressOk() (*string, bool) {
	if o == nil || o.ForwarderAddress == nil {
		return nil, false
	}
	return o.ForwarderAddress, true
}

// HasForwarderAddress returns a boolean if a field has been set.
func (o *CreateContractRequest) HasForwarderAddress() bool {
	if o != nil && o.ForwarderAddress != nil {
		return true
	}

	return false
}

// SetForwarderAddress gets a reference to the given string and assigns it to the ForwarderAddress field.
func (o *CreateContractRequest) SetForwarderAddress(v string) {
	o.ForwarderAddress = &v
}

// GetAbi returns the Abi field value
func (o *CreateContractRequest) GetAbi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value
// and a boolean to check if the value has been set.
func (o *CreateContractRequest) GetAbiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Abi, true
}

// SetAbi sets field value
func (o *CreateContractRequest) SetAbi(v string) {
	o.Abi = v
}

// GetChain returns the Chain field value
func (o *CreateContractRequest) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *CreateContractRequest) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *CreateContractRequest) SetChain(v string) {
	o.Chain = v
}

func (o CreateContractRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if o.ForwarderAddress != nil {
		toSerialize["forwarderAddress"] = o.ForwarderAddress
	}
	if true {
		toSerialize["abi"] = o.Abi
	}
	if true {
		toSerialize["chain"] = o.Chain
	}
	return json.Marshal(toSerialize)
}

type NullableCreateContractRequest struct {
	value *CreateContractRequest
	isSet bool
}

func (v NullableCreateContractRequest) Get() *CreateContractRequest {
	return v.value
}

func (v *NullableCreateContractRequest) Set(val *CreateContractRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContractRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContractRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContractRequest(val *CreateContractRequest) *NullableCreateContractRequest {
	return &NullableCreateContractRequest{value: val, isSet: true}
}

func (v NullableCreateContractRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContractRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


