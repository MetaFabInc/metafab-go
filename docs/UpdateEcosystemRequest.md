# UpdateEcosystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A new name. Replaces the ecosystem&#39;s current name. | [optional] 
**Email** | Pointer to **string** | A new email address. The ecosystem&#39;s old email will no longer be valid for account authentication. &#x60;currentPassword&#x60; must also be provided. | [optional] 
**CurrentPassword** | Pointer to **string** | The ecosystem&#39;s current password. Must be provided if setting &#x60;newPassword&#x60; or &#x60;email&#x60;. | [optional] 
**NewPassword** | Pointer to **string** | A new password. The ecosystem&#39;s old password will no longer be valid. | [optional] 
**IconImageBase64** | Pointer to **string** | A base64 string of the icon image for this ecosystem. Supported image formats are &#x60;jpg&#x60;, &#x60;jpeg&#x60;, &#x60;png&#x60;, &#x60;gif&#x60; Recommended size is 512x512 pixels, or 1:1 aspect ratio. This image is used for your profile authorization flow and other MetaFab features for your ecosystem. | [optional] 
**CoverImageBase64** | Pointer to **string** | A base64 string of the cover image for this ecosystem. Supported image formats are &#x60;jpg&#x60;, &#x60;jpeg&#x60;, &#x60;png&#x60;, &#x60;gif&#x60;. Recommended size is 1600x1000 pixels, or 16:10 aspect ratio.  This image is used as the background image for your profile authorization flow and other MetaFab features for your ecosystem. | [optional] 
**PrimaryColorHex** | Pointer to **string** | A valid hex color code. This color is used for your profile authorization flow to control the color of buttons and other brandable MetaFab features for your ecosystem. | [optional] 
**ResetPublishedKey** | Pointer to **bool** | Revokes the ecosystem&#39;s previous published key and returns a new one if true. | [optional] 
**ResetSecretKey** | Pointer to **bool** | Revokes the ecosystem&#39;s previous secret key and returns a new on if true. | [optional] 

## Methods

### NewUpdateEcosystemRequest

`func NewUpdateEcosystemRequest() *UpdateEcosystemRequest`

NewUpdateEcosystemRequest instantiates a new UpdateEcosystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEcosystemRequestWithDefaults

`func NewUpdateEcosystemRequestWithDefaults() *UpdateEcosystemRequest`

NewUpdateEcosystemRequestWithDefaults instantiates a new UpdateEcosystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEcosystemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEcosystemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEcosystemRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEcosystemRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateEcosystemRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateEcosystemRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateEcosystemRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateEcosystemRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCurrentPassword

`func (o *UpdateEcosystemRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdateEcosystemRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdateEcosystemRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdateEcosystemRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UpdateEcosystemRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdateEcosystemRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdateEcosystemRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdateEcosystemRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetIconImageBase64

`func (o *UpdateEcosystemRequest) GetIconImageBase64() string`

GetIconImageBase64 returns the IconImageBase64 field if non-nil, zero value otherwise.

### GetIconImageBase64Ok

`func (o *UpdateEcosystemRequest) GetIconImageBase64Ok() (*string, bool)`

GetIconImageBase64Ok returns a tuple with the IconImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconImageBase64

`func (o *UpdateEcosystemRequest) SetIconImageBase64(v string)`

SetIconImageBase64 sets IconImageBase64 field to given value.

### HasIconImageBase64

`func (o *UpdateEcosystemRequest) HasIconImageBase64() bool`

HasIconImageBase64 returns a boolean if a field has been set.

### GetCoverImageBase64

`func (o *UpdateEcosystemRequest) GetCoverImageBase64() string`

GetCoverImageBase64 returns the CoverImageBase64 field if non-nil, zero value otherwise.

### GetCoverImageBase64Ok

`func (o *UpdateEcosystemRequest) GetCoverImageBase64Ok() (*string, bool)`

GetCoverImageBase64Ok returns a tuple with the CoverImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageBase64

`func (o *UpdateEcosystemRequest) SetCoverImageBase64(v string)`

SetCoverImageBase64 sets CoverImageBase64 field to given value.

### HasCoverImageBase64

`func (o *UpdateEcosystemRequest) HasCoverImageBase64() bool`

HasCoverImageBase64 returns a boolean if a field has been set.

### GetPrimaryColorHex

`func (o *UpdateEcosystemRequest) GetPrimaryColorHex() string`

GetPrimaryColorHex returns the PrimaryColorHex field if non-nil, zero value otherwise.

### GetPrimaryColorHexOk

`func (o *UpdateEcosystemRequest) GetPrimaryColorHexOk() (*string, bool)`

GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorHex

`func (o *UpdateEcosystemRequest) SetPrimaryColorHex(v string)`

SetPrimaryColorHex sets PrimaryColorHex field to given value.

### HasPrimaryColorHex

`func (o *UpdateEcosystemRequest) HasPrimaryColorHex() bool`

HasPrimaryColorHex returns a boolean if a field has been set.

### GetResetPublishedKey

`func (o *UpdateEcosystemRequest) GetResetPublishedKey() bool`

GetResetPublishedKey returns the ResetPublishedKey field if non-nil, zero value otherwise.

### GetResetPublishedKeyOk

`func (o *UpdateEcosystemRequest) GetResetPublishedKeyOk() (*bool, bool)`

GetResetPublishedKeyOk returns a tuple with the ResetPublishedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPublishedKey

`func (o *UpdateEcosystemRequest) SetResetPublishedKey(v bool)`

SetResetPublishedKey sets ResetPublishedKey field to given value.

### HasResetPublishedKey

`func (o *UpdateEcosystemRequest) HasResetPublishedKey() bool`

HasResetPublishedKey returns a boolean if a field has been set.

### GetResetSecretKey

`func (o *UpdateEcosystemRequest) GetResetSecretKey() bool`

GetResetSecretKey returns the ResetSecretKey field if non-nil, zero value otherwise.

### GetResetSecretKeyOk

`func (o *UpdateEcosystemRequest) GetResetSecretKeyOk() (*bool, bool)`

GetResetSecretKeyOk returns a tuple with the ResetSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetSecretKey

`func (o *UpdateEcosystemRequest) SetResetSecretKey(v bool)`

SetResetSecretKey sets ResetSecretKey field to given value.

### HasResetSecretKey

`func (o *UpdateEcosystemRequest) HasResetSecretKey() bool`

HasResetSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


