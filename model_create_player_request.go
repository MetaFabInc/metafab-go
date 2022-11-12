/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.2.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// CreatePlayerRequest struct for CreatePlayerRequest
type CreatePlayerRequest struct {
	// The players username, used to authenticate the player and if desired represent them in game. Usernames are unique. There cannot be 2 users with the same username created for a game.
	Username string `json:"username"`
	// The password to authenticate as the player. Additionally, this password is used to encrypt/decrypt a player's primary wallet and must be provided anytime this player makes blockchain interactions through various endpoints.
	Password string `json:"password"`
}

// NewCreatePlayerRequest instantiates a new CreatePlayerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePlayerRequest(username string, password string) *CreatePlayerRequest {
	this := CreatePlayerRequest{}
	this.Username = username
	this.Password = password
	return &this
}

// NewCreatePlayerRequestWithDefaults instantiates a new CreatePlayerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePlayerRequestWithDefaults() *CreatePlayerRequest {
	this := CreatePlayerRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *CreatePlayerRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreatePlayerRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreatePlayerRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *CreatePlayerRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreatePlayerRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreatePlayerRequest) SetPassword(v string) {
	o.Password = v
}

func (o CreatePlayerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePlayerRequest struct {
	value *CreatePlayerRequest
	isSet bool
}

func (v NullableCreatePlayerRequest) Get() *CreatePlayerRequest {
	return v.value
}

func (v *NullableCreatePlayerRequest) Set(val *CreatePlayerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePlayerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePlayerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePlayerRequest(val *CreatePlayerRequest) *NullableCreatePlayerRequest {
	return &NullableCreatePlayerRequest{value: val, isSet: true}
}

func (v NullableCreatePlayerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePlayerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


