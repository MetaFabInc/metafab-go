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

// GetPlayerData200Response struct for GetPlayerData200Response
type GetPlayerData200Response struct {
	ProtectedData map[string]interface{} `json:"protectedData,omitempty"`
	PublicData map[string]interface{} `json:"publicData,omitempty"`
}

// NewGetPlayerData200Response instantiates a new GetPlayerData200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPlayerData200Response() *GetPlayerData200Response {
	this := GetPlayerData200Response{}
	return &this
}

// NewGetPlayerData200ResponseWithDefaults instantiates a new GetPlayerData200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPlayerData200ResponseWithDefaults() *GetPlayerData200Response {
	this := GetPlayerData200Response{}
	return &this
}

// GetProtectedData returns the ProtectedData field value if set, zero value otherwise.
func (o *GetPlayerData200Response) GetProtectedData() map[string]interface{} {
	if o == nil || o.ProtectedData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ProtectedData
}

// GetProtectedDataOk returns a tuple with the ProtectedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlayerData200Response) GetProtectedDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ProtectedData == nil {
		return nil, false
	}
	return o.ProtectedData, true
}

// HasProtectedData returns a boolean if a field has been set.
func (o *GetPlayerData200Response) HasProtectedData() bool {
	if o != nil && o.ProtectedData != nil {
		return true
	}

	return false
}

// SetProtectedData gets a reference to the given map[string]interface{} and assigns it to the ProtectedData field.
func (o *GetPlayerData200Response) SetProtectedData(v map[string]interface{}) {
	o.ProtectedData = v
}

// GetPublicData returns the PublicData field value if set, zero value otherwise.
func (o *GetPlayerData200Response) GetPublicData() map[string]interface{} {
	if o == nil || o.PublicData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PublicData
}

// GetPublicDataOk returns a tuple with the PublicData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlayerData200Response) GetPublicDataOk() (map[string]interface{}, bool) {
	if o == nil || o.PublicData == nil {
		return nil, false
	}
	return o.PublicData, true
}

// HasPublicData returns a boolean if a field has been set.
func (o *GetPlayerData200Response) HasPublicData() bool {
	if o != nil && o.PublicData != nil {
		return true
	}

	return false
}

// SetPublicData gets a reference to the given map[string]interface{} and assigns it to the PublicData field.
func (o *GetPlayerData200Response) SetPublicData(v map[string]interface{}) {
	o.PublicData = v
}

func (o GetPlayerData200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectedData != nil {
		toSerialize["protectedData"] = o.ProtectedData
	}
	if o.PublicData != nil {
		toSerialize["publicData"] = o.PublicData
	}
	return json.Marshal(toSerialize)
}

type NullableGetPlayerData200Response struct {
	value *GetPlayerData200Response
	isSet bool
}

func (v NullableGetPlayerData200Response) Get() *GetPlayerData200Response {
	return v.value
}

func (v *NullableGetPlayerData200Response) Set(val *GetPlayerData200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPlayerData200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPlayerData200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPlayerData200Response(val *GetPlayerData200Response) *NullableGetPlayerData200Response {
	return &NullableGetPlayerData200Response{value: val, isSet: true}
}

func (v NullableGetPlayerData200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPlayerData200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


