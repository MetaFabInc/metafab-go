# UpdatePlayerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** | The player&#39;s current password. Must be provided if setting &#x60;newPassword&#x60;. | [optional] 
**NewPassword** | Pointer to **string** | A new password. The player&#39;s old password will no longer be valid. | [optional] 
**ResetAccessToken** | Pointer to **bool** | Revokes the player&#39;s previous access token and returns a new one if true. | [optional] 

## Methods

### NewUpdatePlayerRequest

`func NewUpdatePlayerRequest() *UpdatePlayerRequest`

NewUpdatePlayerRequest instantiates a new UpdatePlayerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlayerRequestWithDefaults

`func NewUpdatePlayerRequestWithDefaults() *UpdatePlayerRequest`

NewUpdatePlayerRequestWithDefaults instantiates a new UpdatePlayerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UpdatePlayerRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdatePlayerRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdatePlayerRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdatePlayerRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UpdatePlayerRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdatePlayerRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdatePlayerRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdatePlayerRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetResetAccessToken

`func (o *UpdatePlayerRequest) GetResetAccessToken() bool`

GetResetAccessToken returns the ResetAccessToken field if non-nil, zero value otherwise.

### GetResetAccessTokenOk

`func (o *UpdatePlayerRequest) GetResetAccessTokenOk() (*bool, bool)`

GetResetAccessTokenOk returns a tuple with the ResetAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAccessToken

`func (o *UpdatePlayerRequest) SetResetAccessToken(v bool)`

SetResetAccessToken sets ResetAccessToken field to given value.

### HasResetAccessToken

`func (o *UpdatePlayerRequest) HasResetAccessToken() bool`

HasResetAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


