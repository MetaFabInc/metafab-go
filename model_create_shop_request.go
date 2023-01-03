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

// CreateShopRequest struct for CreateShopRequest
type CreateShopRequest struct {
	// The blockchain you want to deploy this shop on. Support for new blockchains are added over time.
	Chain string `json:"chain"`
}

// NewCreateShopRequest instantiates a new CreateShopRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShopRequest(chain string) *CreateShopRequest {
	this := CreateShopRequest{}
	this.Chain = chain
	return &this
}

// NewCreateShopRequestWithDefaults instantiates a new CreateShopRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShopRequestWithDefaults() *CreateShopRequest {
	this := CreateShopRequest{}
	return &this
}

// GetChain returns the Chain field value
func (o *CreateShopRequest) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *CreateShopRequest) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *CreateShopRequest) SetChain(v string) {
	o.Chain = v
}

func (o CreateShopRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chain"] = o.Chain
	}
	return json.Marshal(toSerialize)
}

type NullableCreateShopRequest struct {
	value *CreateShopRequest
	isSet bool
}

func (v NullableCreateShopRequest) Get() *CreateShopRequest {
	return v.value
}

func (v *NullableCreateShopRequest) Set(val *CreateShopRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShopRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShopRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShopRequest(val *CreateShopRequest) *NullableCreateShopRequest {
	return &NullableCreateShopRequest{value: val, isSet: true}
}

func (v NullableCreateShopRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShopRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


