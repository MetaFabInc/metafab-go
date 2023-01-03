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

// SetPlayerDataRequest struct for SetPlayerDataRequest
type SetPlayerDataRequest struct {
	// protectedData can only be set if `X-Authorization` includes credentials for the game the target player is a part of. Expects an arbitrary object allowed to contain any set of properties and nested data within those properties, including arrays.  protectedData is great for storing sensitive player data like tracking experience points, off-chain inventories, save states, and more - things that players shouldn't have the ability to directly change themselves.
	ProtectedData map[string]interface{} `json:"protectedData,omitempty"`
	// publicData can be set if `X-Authorization` includes credentials for the target player or game the player is a part of. Expects an arbitrary object allowed to contain any set of properties and nested data within those properties, including arrays.  publicData is great for storing player preferences like in-game settings, non-sensitive data and more. Anything that a player should have the ability to directly change themselves without client or server verification can be stored in publicData.
	PublicData map[string]interface{} `json:"publicData,omitempty"`
}

// NewSetPlayerDataRequest instantiates a new SetPlayerDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPlayerDataRequest() *SetPlayerDataRequest {
	this := SetPlayerDataRequest{}
	return &this
}

// NewSetPlayerDataRequestWithDefaults instantiates a new SetPlayerDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPlayerDataRequestWithDefaults() *SetPlayerDataRequest {
	this := SetPlayerDataRequest{}
	return &this
}

// GetProtectedData returns the ProtectedData field value if set, zero value otherwise.
func (o *SetPlayerDataRequest) GetProtectedData() map[string]interface{} {
	if o == nil || o.ProtectedData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ProtectedData
}

// GetProtectedDataOk returns a tuple with the ProtectedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetPlayerDataRequest) GetProtectedDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ProtectedData == nil {
		return nil, false
	}
	return o.ProtectedData, true
}

// HasProtectedData returns a boolean if a field has been set.
func (o *SetPlayerDataRequest) HasProtectedData() bool {
	if o != nil && o.ProtectedData != nil {
		return true
	}

	return false
}

// SetProtectedData gets a reference to the given map[string]interface{} and assigns it to the ProtectedData field.
func (o *SetPlayerDataRequest) SetProtectedData(v map[string]interface{}) {
	o.ProtectedData = v
}

// GetPublicData returns the PublicData field value if set, zero value otherwise.
func (o *SetPlayerDataRequest) GetPublicData() map[string]interface{} {
	if o == nil || o.PublicData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PublicData
}

// GetPublicDataOk returns a tuple with the PublicData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetPlayerDataRequest) GetPublicDataOk() (map[string]interface{}, bool) {
	if o == nil || o.PublicData == nil {
		return nil, false
	}
	return o.PublicData, true
}

// HasPublicData returns a boolean if a field has been set.
func (o *SetPlayerDataRequest) HasPublicData() bool {
	if o != nil && o.PublicData != nil {
		return true
	}

	return false
}

// SetPublicData gets a reference to the given map[string]interface{} and assigns it to the PublicData field.
func (o *SetPlayerDataRequest) SetPublicData(v map[string]interface{}) {
	o.PublicData = v
}

func (o SetPlayerDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectedData != nil {
		toSerialize["protectedData"] = o.ProtectedData
	}
	if o.PublicData != nil {
		toSerialize["publicData"] = o.PublicData
	}
	return json.Marshal(toSerialize)
}

type NullableSetPlayerDataRequest struct {
	value *SetPlayerDataRequest
	isSet bool
}

func (v NullableSetPlayerDataRequest) Get() *SetPlayerDataRequest {
	return v.value
}

func (v *NullableSetPlayerDataRequest) Set(val *SetPlayerDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPlayerDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPlayerDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPlayerDataRequest(val *SetPlayerDataRequest) *NullableSetPlayerDataRequest {
	return &NullableSetPlayerDataRequest{value: val, isSet: true}
}

func (v NullableSetPlayerDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetPlayerDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


