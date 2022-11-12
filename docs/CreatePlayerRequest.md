# CreatePlayerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The players username, used to authenticate the player and if desired represent them in game. Usernames are unique. There cannot be 2 users with the same username created for a game. | 
**Password** | **string** | The password to authenticate as the player. Additionally, this password is used to encrypt/decrypt a player&#39;s primary wallet and must be provided anytime this player makes blockchain interactions through various endpoints. | 

## Methods

### NewCreatePlayerRequest

`func NewCreatePlayerRequest(username string, password string, ) *CreatePlayerRequest`

NewCreatePlayerRequest instantiates a new CreatePlayerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlayerRequestWithDefaults

`func NewCreatePlayerRequestWithDefaults() *CreatePlayerRequest`

NewCreatePlayerRequestWithDefaults instantiates a new CreatePlayerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreatePlayerRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreatePlayerRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreatePlayerRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreatePlayerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreatePlayerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreatePlayerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


