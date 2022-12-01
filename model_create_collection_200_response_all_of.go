/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.3.0
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// CreateCollection200ResponseAllOf struct for CreateCollection200ResponseAllOf
type CreateCollection200ResponseAllOf struct {
	Contract *CreateCollection200ResponseAllOfContract `json:"contract,omitempty"`
}

// NewCreateCollection200ResponseAllOf instantiates a new CreateCollection200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollection200ResponseAllOf() *CreateCollection200ResponseAllOf {
	this := CreateCollection200ResponseAllOf{}
	return &this
}

// NewCreateCollection200ResponseAllOfWithDefaults instantiates a new CreateCollection200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollection200ResponseAllOfWithDefaults() *CreateCollection200ResponseAllOf {
	this := CreateCollection200ResponseAllOf{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *CreateCollection200ResponseAllOf) GetContract() CreateCollection200ResponseAllOfContract {
	if o == nil || o.Contract == nil {
		var ret CreateCollection200ResponseAllOfContract
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollection200ResponseAllOf) GetContractOk() (*CreateCollection200ResponseAllOfContract, bool) {
	if o == nil || o.Contract == nil {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *CreateCollection200ResponseAllOf) HasContract() bool {
	if o != nil && o.Contract != nil {
		return true
	}

	return false
}

// SetContract gets a reference to the given CreateCollection200ResponseAllOfContract and assigns it to the Contract field.
func (o *CreateCollection200ResponseAllOf) SetContract(v CreateCollection200ResponseAllOfContract) {
	o.Contract = &v
}

func (o CreateCollection200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contract != nil {
		toSerialize["contract"] = o.Contract
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCollection200ResponseAllOf struct {
	value *CreateCollection200ResponseAllOf
	isSet bool
}

func (v NullableCreateCollection200ResponseAllOf) Get() *CreateCollection200ResponseAllOf {
	return v.value
}

func (v *NullableCreateCollection200ResponseAllOf) Set(val *CreateCollection200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollection200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollection200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollection200ResponseAllOf(val *CreateCollection200ResponseAllOf) *NullableCreateCollection200ResponseAllOf {
	return &NullableCreateCollection200ResponseAllOf{value: val, isSet: true}
}

func (v NullableCreateCollection200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollection200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


