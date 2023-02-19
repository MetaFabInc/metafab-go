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

// GetCollections200ResponseInnerAllOf struct for GetCollections200ResponseInnerAllOf
type GetCollections200ResponseInnerAllOf struct {
	Contract *ContractModel `json:"contract,omitempty"`
}

// NewGetCollections200ResponseInnerAllOf instantiates a new GetCollections200ResponseInnerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollections200ResponseInnerAllOf() *GetCollections200ResponseInnerAllOf {
	this := GetCollections200ResponseInnerAllOf{}
	return &this
}

// NewGetCollections200ResponseInnerAllOfWithDefaults instantiates a new GetCollections200ResponseInnerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollections200ResponseInnerAllOfWithDefaults() *GetCollections200ResponseInnerAllOf {
	this := GetCollections200ResponseInnerAllOf{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *GetCollections200ResponseInnerAllOf) GetContract() ContractModel {
	if o == nil || isNil(o.Contract) {
		var ret ContractModel
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollections200ResponseInnerAllOf) GetContractOk() (*ContractModel, bool) {
	if o == nil || isNil(o.Contract) {
    return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *GetCollections200ResponseInnerAllOf) HasContract() bool {
	if o != nil && !isNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given ContractModel and assigns it to the Contract field.
func (o *GetCollections200ResponseInnerAllOf) SetContract(v ContractModel) {
	o.Contract = &v
}

func (o GetCollections200ResponseInnerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Contract) {
		toSerialize["contract"] = o.Contract
	}
	return json.Marshal(toSerialize)
}

type NullableGetCollections200ResponseInnerAllOf struct {
	value *GetCollections200ResponseInnerAllOf
	isSet bool
}

func (v NullableGetCollections200ResponseInnerAllOf) Get() *GetCollections200ResponseInnerAllOf {
	return v.value
}

func (v *NullableGetCollections200ResponseInnerAllOf) Set(val *GetCollections200ResponseInnerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollections200ResponseInnerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollections200ResponseInnerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollections200ResponseInnerAllOf(val *GetCollections200ResponseInnerAllOf) *NullableGetCollections200ResponseInnerAllOf {
	return &NullableGetCollections200ResponseInnerAllOf{value: val, isSet: true}
}

func (v NullableGetCollections200ResponseInnerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCollections200ResponseInnerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


