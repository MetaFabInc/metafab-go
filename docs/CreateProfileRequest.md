# CreateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The profile&#39;s username, used to authenticate the profile. Profile usernames are globally unique across MetaFab. There cannot be 2 profiles with the same username created. | 
**Password** | **string** | The password to authenticate as the profile. Additionally, this password is used to encrypt/decrypt a profile&#39;s primary wallet and must be provided anytime this profile makes blockchain interactions through various endpoints. | 

## Methods

### NewCreateProfileRequest

`func NewCreateProfileRequest(username string, password string, ) *CreateProfileRequest`

NewCreateProfileRequest instantiates a new CreateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileRequestWithDefaults

`func NewCreateProfileRequestWithDefaults() *CreateProfileRequest`

NewCreateProfileRequestWithDefaults instantiates a new CreateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateProfileRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateProfileRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateProfileRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreateProfileRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateProfileRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateProfileRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


