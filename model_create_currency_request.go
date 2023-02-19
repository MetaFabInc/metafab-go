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

// CreateCurrencyRequest struct for CreateCurrencyRequest
type CreateCurrencyRequest struct {
	// The name of this currency. This can be anything, such as `Bright Gems`, `Gold`, etc.
	Name string `json:"name"`
	// The shorthand symbol to represent this currency. This can be anything, such as `BGEM`, `GLD`, etc.
	Symbol string `json:"symbol"`
	// The maximum amount of this currency that can ever exist. Use `0` if you do not want this currency to have a maximum supply.
	SupplyCap float32 `json:"supplyCap"`
	// The blockchain you want to deploy this currency on. Support for new blockchains are added over time.
	Chain string `json:"chain"`
}

// NewCreateCurrencyRequest instantiates a new CreateCurrencyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCurrencyRequest(name string, symbol string, supplyCap float32, chain string) *CreateCurrencyRequest {
	this := CreateCurrencyRequest{}
	this.Name = name
	this.Symbol = symbol
	this.SupplyCap = supplyCap
	this.Chain = chain
	return &this
}

// NewCreateCurrencyRequestWithDefaults instantiates a new CreateCurrencyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCurrencyRequestWithDefaults() *CreateCurrencyRequest {
	this := CreateCurrencyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCurrencyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCurrencyRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCurrencyRequest) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *CreateCurrencyRequest) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *CreateCurrencyRequest) GetSymbolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *CreateCurrencyRequest) SetSymbol(v string) {
	o.Symbol = v
}

// GetSupplyCap returns the SupplyCap field value
func (o *CreateCurrencyRequest) GetSupplyCap() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SupplyCap
}

// GetSupplyCapOk returns a tuple with the SupplyCap field value
// and a boolean to check if the value has been set.
func (o *CreateCurrencyRequest) GetSupplyCapOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SupplyCap, true
}

// SetSupplyCap sets field value
func (o *CreateCurrencyRequest) SetSupplyCap(v float32) {
	o.SupplyCap = v
}

// GetChain returns the Chain field value
func (o *CreateCurrencyRequest) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *CreateCurrencyRequest) GetChainOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *CreateCurrencyRequest) SetChain(v string) {
	o.Chain = v
}

func (o CreateCurrencyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["supplyCap"] = o.SupplyCap
	}
	if true {
		toSerialize["chain"] = o.Chain
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCurrencyRequest struct {
	value *CreateCurrencyRequest
	isSet bool
}

func (v NullableCreateCurrencyRequest) Get() *CreateCurrencyRequest {
	return v.value
}

func (v *NullableCreateCurrencyRequest) Set(val *CreateCurrencyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCurrencyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCurrencyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCurrencyRequest(val *CreateCurrencyRequest) *NullableCreateCurrencyRequest {
	return &NullableCreateCurrencyRequest{value: val, isSet: true}
}

func (v NullableCreateCurrencyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCurrencyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


