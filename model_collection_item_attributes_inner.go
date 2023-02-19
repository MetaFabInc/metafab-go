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

// CollectionItemAttributesInner struct for CollectionItemAttributesInner
type CollectionItemAttributesInner struct {
	// This field has not had a description added.
	TraitType *string `json:"trait_type,omitempty"`
	Value *CollectionItemAttributesInnerValue `json:"value,omitempty"`
}

// NewCollectionItemAttributesInner instantiates a new CollectionItemAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionItemAttributesInner() *CollectionItemAttributesInner {
	this := CollectionItemAttributesInner{}
	return &this
}

// NewCollectionItemAttributesInnerWithDefaults instantiates a new CollectionItemAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionItemAttributesInnerWithDefaults() *CollectionItemAttributesInner {
	this := CollectionItemAttributesInner{}
	return &this
}

// GetTraitType returns the TraitType field value if set, zero value otherwise.
func (o *CollectionItemAttributesInner) GetTraitType() string {
	if o == nil || isNil(o.TraitType) {
		var ret string
		return ret
	}
	return *o.TraitType
}

// GetTraitTypeOk returns a tuple with the TraitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItemAttributesInner) GetTraitTypeOk() (*string, bool) {
	if o == nil || isNil(o.TraitType) {
    return nil, false
	}
	return o.TraitType, true
}

// HasTraitType returns a boolean if a field has been set.
func (o *CollectionItemAttributesInner) HasTraitType() bool {
	if o != nil && !isNil(o.TraitType) {
		return true
	}

	return false
}

// SetTraitType gets a reference to the given string and assigns it to the TraitType field.
func (o *CollectionItemAttributesInner) SetTraitType(v string) {
	o.TraitType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionItemAttributesInner) GetValue() CollectionItemAttributesInnerValue {
	if o == nil || isNil(o.Value) {
		var ret CollectionItemAttributesInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItemAttributesInner) GetValueOk() (*CollectionItemAttributesInnerValue, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionItemAttributesInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CollectionItemAttributesInnerValue and assigns it to the Value field.
func (o *CollectionItemAttributesInner) SetValue(v CollectionItemAttributesInnerValue) {
	o.Value = &v
}

func (o CollectionItemAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TraitType) {
		toSerialize["trait_type"] = o.TraitType
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionItemAttributesInner struct {
	value *CollectionItemAttributesInner
	isSet bool
}

func (v NullableCollectionItemAttributesInner) Get() *CollectionItemAttributesInner {
	return v.value
}

func (v *NullableCollectionItemAttributesInner) Set(val *CollectionItemAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionItemAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionItemAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionItemAttributesInner(val *CollectionItemAttributesInner) *NullableCollectionItemAttributesInner {
	return &NullableCollectionItemAttributesInner{value: val, isSet: true}
}

func (v NullableCollectionItemAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionItemAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


