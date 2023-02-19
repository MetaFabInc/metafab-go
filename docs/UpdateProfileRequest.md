# UpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** | The profile&#39;s current password. Must be provided if setting &#x60;newPassword&#x60;. | [optional] 
**NewPassword** | Pointer to **string** | A new password. The profile&#39;s old password will no longer be valid. | [optional] 
**ResetAccessToken** | Pointer to **bool** | Revokes the profile&#39;s previous access token and returns a new one if true. | [optional] 

## Methods

### NewUpdateProfileRequest

`func NewUpdateProfileRequest() *UpdateProfileRequest`

NewUpdateProfileRequest instantiates a new UpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileRequestWithDefaults

`func NewUpdateProfileRequestWithDefaults() *UpdateProfileRequest`

NewUpdateProfileRequestWithDefaults instantiates a new UpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UpdateProfileRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdateProfileRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdateProfileRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdateProfileRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UpdateProfileRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdateProfileRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdateProfileRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdateProfileRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetResetAccessToken

`func (o *UpdateProfileRequest) GetResetAccessToken() bool`

GetResetAccessToken returns the ResetAccessToken field if non-nil, zero value otherwise.

### GetResetAccessTokenOk

`func (o *UpdateProfileRequest) GetResetAccessTokenOk() (*bool, bool)`

GetResetAccessTokenOk returns a tuple with the ResetAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAccessToken

`func (o *UpdateProfileRequest) SetResetAccessToken(v bool)`

SetResetAccessToken sets ResetAccessToken field to given value.

### HasResetAccessToken

`func (o *UpdateProfileRequest) HasResetAccessToken() bool`

HasResetAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


