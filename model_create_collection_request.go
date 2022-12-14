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

// CreateCollectionRequest struct for CreateCollectionRequest
type CreateCollectionRequest struct {
	// The blockchain you want to deploy this item collection on. Support for new blockchains are added over time.
	Chain string `json:"chain"`
}

// NewCreateCollectionRequest instantiates a new CreateCollectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollectionRequest(chain string) *CreateCollectionRequest {
	this := CreateCollectionRequest{}
	this.Chain = chain
	return &this
}

// NewCreateCollectionRequestWithDefaults instantiates a new CreateCollectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollectionRequestWithDefaults() *CreateCollectionRequest {
	this := CreateCollectionRequest{}
	return &this
}

// GetChain returns the Chain field value
func (o *CreateCollectionRequest) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionRequest) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *CreateCollectionRequest) SetChain(v string) {
	o.Chain = v
}

func (o CreateCollectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chain"] = o.Chain
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCollectionRequest struct {
	value *CreateCollectionRequest
	isSet bool
}

func (v NullableCreateCollectionRequest) Get() *CreateCollectionRequest {
	return v.value
}

func (v *NullableCreateCollectionRequest) Set(val *CreateCollectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollectionRequest(val *CreateCollectionRequest) *NullableCreateCollectionRequest {
	return &NullableCreateCollectionRequest{value: val, isSet: true}
}

func (v NullableCreateCollectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


