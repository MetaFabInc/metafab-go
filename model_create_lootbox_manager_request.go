/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.4.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// CreateLootboxManagerRequest struct for CreateLootboxManagerRequest
type CreateLootboxManagerRequest struct {
	// The blockchain you want to deploy this lootbox manager on. Support for new blockchains are added over time.
	Chain string `json:"chain"`
}

// NewCreateLootboxManagerRequest instantiates a new CreateLootboxManagerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLootboxManagerRequest(chain string) *CreateLootboxManagerRequest {
	this := CreateLootboxManagerRequest{}
	this.Chain = chain
	return &this
}

// NewCreateLootboxManagerRequestWithDefaults instantiates a new CreateLootboxManagerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLootboxManagerRequestWithDefaults() *CreateLootboxManagerRequest {
	this := CreateLootboxManagerRequest{}
	return &this
}

// GetChain returns the Chain field value
func (o *CreateLootboxManagerRequest) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *CreateLootboxManagerRequest) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *CreateLootboxManagerRequest) SetChain(v string) {
	o.Chain = v
}

func (o CreateLootboxManagerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chain"] = o.Chain
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLootboxManagerRequest struct {
	value *CreateLootboxManagerRequest
	isSet bool
}

func (v NullableCreateLootboxManagerRequest) Get() *CreateLootboxManagerRequest {
	return v.value
}

func (v *NullableCreateLootboxManagerRequest) Set(val *CreateLootboxManagerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLootboxManagerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLootboxManagerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLootboxManagerRequest(val *CreateLootboxManagerRequest) *NullableCreateLootboxManagerRequest {
	return &NullableCreateLootboxManagerRequest{value: val, isSet: true}
}

func (v NullableCreateLootboxManagerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLootboxManagerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


