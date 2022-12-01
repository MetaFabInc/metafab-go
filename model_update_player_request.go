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

// UpdatePlayerRequest struct for UpdatePlayerRequest
type UpdatePlayerRequest struct {
	// The player's current password. Must be provided if setting `newPassword`.
	CurrentPassword *string `json:"currentPassword,omitempty"`
	// A new password. The player's old password will no longer be valid.
	NewPassword *string `json:"newPassword,omitempty"`
	// Revokes the player's previous access token and returns a new one if true.
	ResetAccessToken *bool `json:"resetAccessToken,omitempty"`
}

// NewUpdatePlayerRequest instantiates a new UpdatePlayerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePlayerRequest() *UpdatePlayerRequest {
	this := UpdatePlayerRequest{}
	return &this
}

// NewUpdatePlayerRequestWithDefaults instantiates a new UpdatePlayerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePlayerRequestWithDefaults() *UpdatePlayerRequest {
	this := UpdatePlayerRequest{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *UpdatePlayerRequest) GetCurrentPassword() string {
	if o == nil || o.CurrentPassword == nil {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlayerRequest) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || o.CurrentPassword == nil {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *UpdatePlayerRequest) HasCurrentPassword() bool {
	if o != nil && o.CurrentPassword != nil {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *UpdatePlayerRequest) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *UpdatePlayerRequest) GetNewPassword() string {
	if o == nil || o.NewPassword == nil {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlayerRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil || o.NewPassword == nil {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *UpdatePlayerRequest) HasNewPassword() bool {
	if o != nil && o.NewPassword != nil {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *UpdatePlayerRequest) SetNewPassword(v string) {
	o.NewPassword = &v
}

// GetResetAccessToken returns the ResetAccessToken field value if set, zero value otherwise.
func (o *UpdatePlayerRequest) GetResetAccessToken() bool {
	if o == nil || o.ResetAccessToken == nil {
		var ret bool
		return ret
	}
	return *o.ResetAccessToken
}

// GetResetAccessTokenOk returns a tuple with the ResetAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlayerRequest) GetResetAccessTokenOk() (*bool, bool) {
	if o == nil || o.ResetAccessToken == nil {
		return nil, false
	}
	return o.ResetAccessToken, true
}

// HasResetAccessToken returns a boolean if a field has been set.
func (o *UpdatePlayerRequest) HasResetAccessToken() bool {
	if o != nil && o.ResetAccessToken != nil {
		return true
	}

	return false
}

// SetResetAccessToken gets a reference to the given bool and assigns it to the ResetAccessToken field.
func (o *UpdatePlayerRequest) SetResetAccessToken(v bool) {
	o.ResetAccessToken = &v
}

func (o UpdatePlayerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPassword != nil {
		toSerialize["currentPassword"] = o.CurrentPassword
	}
	if o.NewPassword != nil {
		toSerialize["newPassword"] = o.NewPassword
	}
	if o.ResetAccessToken != nil {
		toSerialize["resetAccessToken"] = o.ResetAccessToken
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePlayerRequest struct {
	value *UpdatePlayerRequest
	isSet bool
}

func (v NullableUpdatePlayerRequest) Get() *UpdatePlayerRequest {
	return v.value
}

func (v *NullableUpdatePlayerRequest) Set(val *UpdatePlayerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePlayerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePlayerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePlayerRequest(val *UpdatePlayerRequest) *NullableUpdatePlayerRequest {
	return &NullableUpdatePlayerRequest{value: val, isSet: true}
}

func (v NullableUpdatePlayerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePlayerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


