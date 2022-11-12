# CreateGameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the game you&#39;re creating. | 
**Email** | **string** | The email address associated with this game and used to login/authenticate as the game. | 
**Password** | **string** | The password to authenticate as the game. Additionally, this password is used to encrypt/decrypt your game&#39;s primary wallet and must be provided anytime this game makes blockchain interactions through various endpoints. | 

## Methods

### NewCreateGameRequest

`func NewCreateGameRequest(name string, email string, password string, ) *CreateGameRequest`

NewCreateGameRequest instantiates a new CreateGameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGameRequestWithDefaults

`func NewCreateGameRequestWithDefaults() *CreateGameRequest`

NewCreateGameRequestWithDefaults instantiates a new CreateGameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGameRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGameRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGameRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateGameRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateGameRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateGameRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CreateGameRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateGameRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateGameRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


