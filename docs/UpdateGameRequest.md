# UpdateGameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A new name. Replaces the game&#39;s current name. | [optional] 
**Email** | Pointer to **string** | A new email address. The game&#39;s old email will no longer be valid for account authentication. &#x60;currentPassword&#x60; must also be provided. | [optional] 
**CurrentPassword** | Pointer to **string** | The game&#39;s current password. Must be provided if setting &#x60;newPassword&#x60; or &#x60;email&#x60;. | [optional] 
**NewPassword** | Pointer to **string** | A new password. The game&#39;s old password will no longer be valid. | [optional] 
**Rpcs** | Pointer to **map[string]string** | Sets a custom RPC for your game to use instead of MetaFab&#39;s default RPCs for the chain(s) you specify.  Expects a JSON object containing key value pairs of supported &#x60;chain&#x60; -&gt; &#x60;rpc url&#x60;. Only the chain names provided as keys in the object will be explicitly overriden. To delete a custom RPC for your game, provide the chain name to delete as a key in the provided object and &#x60;null&#x60; as the value.  Set RPC example, &#x60;{ MATIC: &#39;https://polygon-rpc.com&#39; }&#x60; Delete RPC example, &#x60;{ MATIC: null }&#x60; | [optional] 
**RedirectUris** | Pointer to **[]string** | An array of valid base redirect uris or exact uris that can be used for the redirect uri of various MetaFab features such as player login/registration and wallet connection.  Expects base or exact uris. For example, you could use include a uri of &#x60;https://trymetafab.com&#x60; and it would allow redirection to any valid uri on the domain, such as &#x60;https://trymetafab.com/play/game&#x60;. | [optional] 
**IconImageBase64** | Pointer to **string** | A base64 string of the icon image for this game. Supported image formats are &#x60;jpg&#x60;, &#x60;jpeg&#x60;, &#x60;png&#x60;, &#x60;gif&#x60; Recommended size is 512x512 pixels, or 1:1 aspect ratio. This image is used for your auth/connect wallet flow and other MetaFab features for your game. | [optional] 
**CoverImageBase64** | Pointer to **string** | A base64 string of the cover image for this game. Supported image formats are &#x60;jpg&#x60;, &#x60;jpeg&#x60;, &#x60;png&#x60;, &#x60;gif&#x60;. Recommended size is 1600x1000 pixels, or 16:10 aspect ratio.  This image is used as the background image for your auth/connect wallet flow and other MetaFab features for your game. | [optional] 
**PrimaryColorHex** | Pointer to **string** | A valid hex color code. This color is used for your auth/connect wallet flow to control the color of buttons and other brandable MetaFab features for your game. | [optional] 
**ResetPublishedKey** | Pointer to **bool** | Revokes the game&#39;s previous published key and returns a new one if true. | [optional] 
**ResetSecretKey** | Pointer to **bool** | Revokes the game&#39;s previous secret key and returns a new on if true. | [optional] 

## Methods

### NewUpdateGameRequest

`func NewUpdateGameRequest() *UpdateGameRequest`

NewUpdateGameRequest instantiates a new UpdateGameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGameRequestWithDefaults

`func NewUpdateGameRequestWithDefaults() *UpdateGameRequest`

NewUpdateGameRequestWithDefaults instantiates a new UpdateGameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGameRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGameRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGameRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGameRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateGameRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateGameRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateGameRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateGameRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCurrentPassword

`func (o *UpdateGameRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdateGameRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdateGameRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdateGameRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UpdateGameRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdateGameRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdateGameRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdateGameRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetRpcs

`func (o *UpdateGameRequest) GetRpcs() map[string]string`

GetRpcs returns the Rpcs field if non-nil, zero value otherwise.

### GetRpcsOk

`func (o *UpdateGameRequest) GetRpcsOk() (*map[string]string, bool)`

GetRpcsOk returns a tuple with the Rpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcs

`func (o *UpdateGameRequest) SetRpcs(v map[string]string)`

SetRpcs sets Rpcs field to given value.

### HasRpcs

`func (o *UpdateGameRequest) HasRpcs() bool`

HasRpcs returns a boolean if a field has been set.

### GetRedirectUris

`func (o *UpdateGameRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *UpdateGameRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *UpdateGameRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *UpdateGameRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetIconImageBase64

`func (o *UpdateGameRequest) GetIconImageBase64() string`

GetIconImageBase64 returns the IconImageBase64 field if non-nil, zero value otherwise.

### GetIconImageBase64Ok

`func (o *UpdateGameRequest) GetIconImageBase64Ok() (*string, bool)`

GetIconImageBase64Ok returns a tuple with the IconImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconImageBase64

`func (o *UpdateGameRequest) SetIconImageBase64(v string)`

SetIconImageBase64 sets IconImageBase64 field to given value.

### HasIconImageBase64

`func (o *UpdateGameRequest) HasIconImageBase64() bool`

HasIconImageBase64 returns a boolean if a field has been set.

### GetCoverImageBase64

`func (o *UpdateGameRequest) GetCoverImageBase64() string`

GetCoverImageBase64 returns the CoverImageBase64 field if non-nil, zero value otherwise.

### GetCoverImageBase64Ok

`func (o *UpdateGameRequest) GetCoverImageBase64Ok() (*string, bool)`

GetCoverImageBase64Ok returns a tuple with the CoverImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageBase64

`func (o *UpdateGameRequest) SetCoverImageBase64(v string)`

SetCoverImageBase64 sets CoverImageBase64 field to given value.

### HasCoverImageBase64

`func (o *UpdateGameRequest) HasCoverImageBase64() bool`

HasCoverImageBase64 returns a boolean if a field has been set.

### GetPrimaryColorHex

`func (o *UpdateGameRequest) GetPrimaryColorHex() string`

GetPrimaryColorHex returns the PrimaryColorHex field if non-nil, zero value otherwise.

### GetPrimaryColorHexOk

`func (o *UpdateGameRequest) GetPrimaryColorHexOk() (*string, bool)`

GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorHex

`func (o *UpdateGameRequest) SetPrimaryColorHex(v string)`

SetPrimaryColorHex sets PrimaryColorHex field to given value.

### HasPrimaryColorHex

`func (o *UpdateGameRequest) HasPrimaryColorHex() bool`

HasPrimaryColorHex returns a boolean if a field has been set.

### GetResetPublishedKey

`func (o *UpdateGameRequest) GetResetPublishedKey() bool`

GetResetPublishedKey returns the ResetPublishedKey field if non-nil, zero value otherwise.

### GetResetPublishedKeyOk

`func (o *UpdateGameRequest) GetResetPublishedKeyOk() (*bool, bool)`

GetResetPublishedKeyOk returns a tuple with the ResetPublishedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPublishedKey

`func (o *UpdateGameRequest) SetResetPublishedKey(v bool)`

SetResetPublishedKey sets ResetPublishedKey field to given value.

### HasResetPublishedKey

`func (o *UpdateGameRequest) HasResetPublishedKey() bool`

HasResetPublishedKey returns a boolean if a field has been set.

### GetResetSecretKey

`func (o *UpdateGameRequest) GetResetSecretKey() bool`

GetResetSecretKey returns the ResetSecretKey field if non-nil, zero value otherwise.

### GetResetSecretKeyOk

`func (o *UpdateGameRequest) GetResetSecretKeyOk() (*bool, bool)`

GetResetSecretKeyOk returns a tuple with the ResetSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetSecretKey

`func (o *UpdateGameRequest) SetResetSecretKey(v bool)`

SetResetSecretKey sets ResetSecretKey field to given value.

### HasResetSecretKey

`func (o *UpdateGameRequest) HasResetSecretKey() bool`

HasResetSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


